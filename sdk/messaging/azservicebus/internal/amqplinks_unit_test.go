// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package internal

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/internal/log"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/amqpwrap"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/go-amqp"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/mock"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/mock/emulation"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/test"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/utils"
	"github.com/stretchr/testify/require"
)

func TestAMQPLinksRetriesUnit(t *testing.T) {
	tests := []struct {
		Err         error
		Attempts    []int32
		ExpectReset bool
	}{
		// nothing goes wrong, only need the one attempt
		{Err: nil, Attempts: []int32{0}},

		// connection related or unknown failures happen, all attempts exhausted
		{Err: &amqp.ConnError{}, Attempts: []int32{0, 1, 2, 3}},
		{Err: errors.New("unknown error"), Attempts: []int32{0, 1, 2, 3}},

		// fatal errors don't retry at all.
		{Err: NewErrNonRetriable("non retriable error"), Attempts: []int32{0}},

		// detach error happens - we have slightly special behavior here in that we do a quick
		// retry for attempt '0', to avoid sleeping if the error was stale. This mostly happens
		// in situations like sending, where you might have long times in between sends and your
		// link is closed due to idling.
		{Err: &amqp.LinkError{}, Attempts: []int32{0, 0, 1, 2, 3}, ExpectReset: true},
	}

	for _, testData := range tests {
		var testName string = ""

		if testData.Err != nil {
			testName = testData.Err.Error()
		}

		t.Run(testName, func(t *testing.T) {
			endLogging := test.CaptureLogsForTest()
			defer endLogging()

			receiver := &FakeAMQPReceiver{}
			sender := &FakeAMQPSender{}
			ns := &FakeNS{}

			links := NewAMQPLinks(NewAMQPLinksArgs{
				NS:         ns,
				EntityPath: "entityPath",
				CreateLinkFunc: func(ctx context.Context, session amqpwrap.AMQPSession) (amqpwrap.AMQPSenderCloser, amqpwrap.AMQPReceiverCloser, error) {
					return sender, receiver, nil
				},
				GetRecoveryKindFunc: GetRecoveryKind,
			})

			defer func() {
				err := links.Close(context.Background(), true)
				require.NoError(t, err)
			}()

			var attempts []int32

			err := links.Retry(context.Background(), log.Event("NotUsed"), "OverallOperation", func(ctx context.Context, lwid *LinksWithID, args *utils.RetryFnArgs) error {
				attempts = append(attempts, args.I)
				return testData.Err
			}, exported.RetryOptions{
				RetryDelay: time.Millisecond,
			})

			require.Equal(t, testData.Err, err)
			require.Equal(t, testData.Attempts, attempts)

			logMessages := endLogging()

			if testData.ExpectReset {
				require.Contains(t, logMessages, fmt.Sprintf("[azsb.Conn] (OverallOperation) Link was previously detached. Attempting quick reconnect to recover from error: %s", err.Error()))
			} else {
				for _, msg := range logMessages {
					require.NotContains(t, msg, "Link was previously detached")
				}
			}
		})
	}
}

func TestAMQPLinks_Logging(t *testing.T) {
	t.Run("link", func(t *testing.T) {
		receiver := &FakeAMQPReceiver{}
		ns := &FakeNS{}

		links := NewAMQPLinks(NewAMQPLinksArgs{
			NS:         ns,
			EntityPath: "entityPath",
			CreateLinkFunc: func(ctx context.Context, session amqpwrap.AMQPSession) (amqpwrap.AMQPSenderCloser, amqpwrap.AMQPReceiverCloser, error) {
				return nil, receiver, nil
			},
			GetRecoveryKindFunc: GetRecoveryKind,
		})

		defer func() {
			err := links.Close(context.Background(), true)
			require.NoError(t, err)
		}()

		endCapture := test.CaptureLogsForTest()
		defer endCapture()

		err := links.RecoverIfNeeded(context.Background(), LinkID{}, &amqp.LinkError{})
		require.NoError(t, err)

		messages := endCapture()

		require.Equal(t, []string{
			"[azsb.Conn] Recovering link for error amqp: link closed",
			"[azsb.Conn] Recovering link only",
			"[azsb.Conn] Recovered links",
		}, messages)
	})

	t.Run("connection", func(t *testing.T) {
		receiver := &FakeAMQPReceiver{}
		ns := &FakeNS{}

		links := NewAMQPLinks(NewAMQPLinksArgs{
			NS:         ns,
			EntityPath: "entityPath",
			CreateLinkFunc: func(ctx context.Context, session amqpwrap.AMQPSession) (amqpwrap.AMQPSenderCloser, amqpwrap.AMQPReceiverCloser, error) {
				return nil, receiver, nil
			}, GetRecoveryKindFunc: GetRecoveryKind,
		})

		defer func() {
			err := links.Close(context.Background(), true)
			require.NoError(t, err)
		}()

		endCapture := test.CaptureLogsForTest()
		defer endCapture()

		err := links.RecoverIfNeeded(context.Background(), LinkID{}, &amqp.ConnError{})
		require.NoError(t, err)

		messages := endCapture()

		require.Equal(t, []string{
			"[azsb.Conn] Recovering link for error amqp: connection closed",
			"[azsb.Conn] Recovering connection (and links)",
			"[azsb.Conn] closing old link: current:{0 0}, old:{0 0}",
			"[azsb.Conn] recreating link: c: true, current:{0 0}, old:{0 0}",
			"[azsb.Conn] Recovered connection and links",
		}, messages)
	})
}

func TestAMQPCloseLinkTimeout_Receiver_ExternalCancellation(t *testing.T) {
	userCtx, cancelUserCtx := context.WithCancel(context.Background())
	defer cancelUserCtx()

	var md *emulation.MockData
	var links *AMQPLinksImpl

	preReceiverMock := func(orig *emulation.MockReceiver, ctx context.Context) error {
		if orig.Source == "entity path" {
			orig.EXPECT().Close(&mock.ContextCreatedForTest{}).DoAndReturn(func(ctx context.Context) error {
				md.Events.CloseLink(orig.LinkEvent())

				// this simulates as if the user cancelled their outer context
				cancelUserCtx()

				return ctx.Err()
			})
		}

		return nil
	}

	createLinkFn := func(ctx context.Context, session amqpwrap.AMQPSession) (amqpwrap.AMQPSenderCloser, amqpwrap.AMQPReceiverCloser, error) {
		receiver, err := session.NewReceiver(ctx, "entity path", &amqp.ReceiverOptions{
			SettlementMode:            amqp.ReceiverSettleModeFirst.Ptr(),
			Credit:                    -1,
			RequestedSenderSettleMode: amqp.SenderSettleModeSettled.Ptr(),
		})

		return nil, receiver, err
	}

	tempMD, tempLinks, ns, cleanup := newAMQPLinksForTest(t, emulation.MockDataOptions{
		PreReceiverMock: preReceiverMock,
	}, createLinkFn)
	defer cleanup()

	md = tempMD
	links = tempLinks

	var lwid *LinksWithID

	// create all the links for the first time.
	err := links.Retry(userCtx, exported.EventConn, "Test", func(ctx context.Context, tmpLWID *LinksWithID, args *utils.RetryFnArgs) error {
		lwid = tmpLWID
		return nil
	}, exported.RetryOptions{})

	require.NoError(t, err)
	require.NotNil(t, lwid)

	// we've initialized our links now.
	require.Equal(t, 3, len(md.Events.GetOpenLinks()), "mgmt link (sender and receiver) + receiver are open")
	require.Equal(t, 1, len(md.Events.GetOpenConns()), "connection is open")

	// capture logs just before we do the "link.Close() timeout causes the connection to close" behavior.
	getLogs := test.CaptureLogsForTest()

	// now close the links. We've made it so the receiver will cancel the context, as if the user
	// interrupted the close. This will end up closing the connection as well.
	rk := links.CloseIfNeeded(userCtx, &amqp.LinkError{})

	require.Contains(t, getLogs(), "[azsb.Conn] Connection closed instead. Link closing has timed out.")

	require.Equal(t, rk, RecoveryKindConn, "Link is upgraded to a connection error instead")
	emulation.RequireNoLeaks(t, md.Events)

	// check that we left ourselves into a correct position to recover.
	// TODO: it'd be nice to see if we "over-recover", which happened in Event Hubs.
	err = links.Retry(context.Background(), exported.EventConn, "Test", func(ctx context.Context, tmpLWID *LinksWithID, args *utils.RetryFnArgs) error {
		lwid = tmpLWID
		return nil
	}, exported.RetryOptions{})

	require.NoError(t, err)
	require.NotNil(t, lwid)

	require.Equal(t, 3, len(md.Events.GetOpenLinks()), "mgmt link (sender and receiver) + receiver are open")
	require.Equal(t, 1, len(md.Events.GetOpenConns()), "connection is open")

	err = links.Close(context.Background(), false)
	require.NoError(t, err)

	require.NoError(t, ns.Close(true))

	emulation.RequireNoLeaks(t, md.Events)
}

func TestAMQPCloseLinkTimeout_Receiver_RecoverIfNeeded(t *testing.T) {
	userCtx, cancelUserCtx := context.WithCancel(context.Background())
	defer cancelUserCtx()

	var md *emulation.MockData
	var links *AMQPLinksImpl

	type keyType string

	preReceiverMock := func(orig *emulation.MockReceiver, ctx context.Context) error {
		if orig.Source == "entity path" {
			orig.EXPECT().Close(&mock.ContextCreatedForTest{}).DoAndReturn(func(ctx context.Context) error {
				if ctx.Value(keyType("close")) == "cancel" {
					md.Events.CloseLink(orig.LinkEvent())

					// this simulates as if the user cancelled their outer context
					cancelUserCtx()
				}

				select {
				case <-ctx.Done():
					return ctx.Err()
				default:
					return nil
				}
			})
		}

		return nil
	}

	createLinkFn := func(ctx context.Context, session amqpwrap.AMQPSession) (amqpwrap.AMQPSenderCloser, amqpwrap.AMQPReceiverCloser, error) {
		receiver, err := session.NewReceiver(ctx, "entity path", &amqp.ReceiverOptions{
			SettlementMode:            amqp.ReceiverSettleModeFirst.Ptr(),
			Credit:                    -1,
			RequestedSenderSettleMode: amqp.SenderSettleModeSettled.Ptr(),
		})

		return nil, receiver, err
	}

	tempMD, tempLinks, _, cleanup := newAMQPLinksForTest(t, emulation.MockDataOptions{
		PreReceiverMock: preReceiverMock,
	}, createLinkFn)
	defer cleanup()

	md = tempMD
	links = tempLinks

	var lwid *LinksWithID

	// create all the links for the first time.
	err := links.Retry(userCtx, exported.EventConn, "Test", func(ctx context.Context, tmpLWID *LinksWithID, args *utils.RetryFnArgs) error {
		lwid = tmpLWID
		return nil
	}, exported.RetryOptions{})

	require.NoError(t, err)
	require.NotNil(t, lwid)

	// we've initialized our links now.
	require.Equal(t, 3, len(md.Events.GetOpenLinks()), "mgmt link (sender and receiver) + receiver are open")
	require.Equal(t, 1, len(md.Events.GetOpenConns()), "connection is open")

	// capture logs just before we do the "link.Close() timeout causes the connection to close" behavior.
	getLogs := test.CaptureLogsForTest()

	// now close the links. We've made it so the receiver will cancel the context, as if the user
	// interrupted the close. This will end up closing the connection as well.
	recoveryErr := links.RecoverIfNeeded(context.WithValue(userCtx, keyType("close"), "cancel"), lwid.ID, &amqp.LinkError{})

	require.Contains(t, getLogs(), "[azsb.Conn] Upgrading to connection reset for recovery instead of link")
	require.ErrorIs(t, recoveryErr, context.Canceled)

	emulation.RequireNoLeaks(t, md.Events)

	nonCancelledCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	recoveryErr = links.RecoverIfNeeded(nonCancelledCtx, lwid.ID, &amqp.LinkError{})

	require.NotContains(t, getLogs(), "[azsb.Conn] Connection reset for recovery instead of link. Link closing has timed out.")
	require.NoError(t, recoveryErr)

	require.Equal(t, 3, len(md.Events.GetOpenLinks()), "mgmt link (sender and receiver) + receiver are open")
	require.Equal(t, 1, len(md.Events.GetOpenConns()), "connection is open")
}

func TestAMQPCloseLinkTimeout_Sender(t *testing.T) {
	userCtx, cancelUserCtx := context.WithCancel(context.Background())
	defer cancelUserCtx()

	var md *emulation.MockData
	var links *AMQPLinksImpl
	var ns *Namespace

	preSenderMock := func(ms *emulation.MockSender, ctx context.Context) error {
		if ms.Target == "entity path" {
			// adjust the sender mock so when it closes it acts as if it were cancelled.
			// when the closing process is interrupted in our recovery it automatically upgrades
			// us to "connection level recovery" instead.
			ms.EXPECT().Close(&mock.ContextCreatedForTest{}).DoAndReturn(func(ctx context.Context) error {
				md.Events.CloseLink(ms.LinkEvent())

				// this simulates as if the user cancelled their outer context
				cancelUserCtx()

				return ctx.Err()
			})
		}

		return nil
	}

	createLinkFn := func(ctx context.Context, session amqpwrap.AMQPSession) (amqpwrap.AMQPSenderCloser, amqpwrap.AMQPReceiverCloser, error) {
		sender, err := session.NewSender(ctx, "entity path", &amqp.SenderOptions{SettlementMode: amqp.SenderSettleModeMixed.Ptr(), RequestedReceiverSettleMode: amqp.ReceiverSettleModeFirst.Ptr()})
		return sender, nil, err
	}

	md, tempLinks, tempNS, cleanup := newAMQPLinksForTest(t, emulation.MockDataOptions{PreSenderMock: preSenderMock}, createLinkFn)
	defer cleanup()

	links = tempLinks
	ns = tempNS

	var lwid *LinksWithID

	// create all the links for the first time.
	err := links.Retry(userCtx, exported.EventConn, "Test", func(ctx context.Context, tmpLWID *LinksWithID, args *utils.RetryFnArgs) error {
		lwid = tmpLWID
		return nil
	}, exported.RetryOptions{})

	require.NoError(t, err)
	require.NotNil(t, lwid)

	// we've initialized our links now.
	require.Equal(t, 3, len(md.Events.GetOpenLinks()), "mgmt link (sender and receiver) + receiver are open")
	require.Equal(t, 1, len(md.Events.GetOpenConns()), "connection is open")

	// capture logs just before we do the "link.Close() timeout causes the connection to close" behavior.
	getLogs := test.CaptureLogsForTest()

	// now close the links. We've made it so the receiver will cancel the context, as if the user
	// interrupted the close. This will end up closing the connection as well.
	rk := links.CloseIfNeeded(userCtx, &amqp.LinkError{})

	require.Contains(t, getLogs(), "[azsb.Conn] Connection closed instead. Link closing has timed out.")

	require.Equal(t, rk, RecoveryKindConn, "Link is upgraded to a connection error instead")
	emulation.RequireNoLeaks(t, md.Events)

	// check that we left ourselves into a correct position to recover.
	// TODO: it'd be nice to see if we "over-recover", which happened in Event Hubs.
	err = links.Retry(context.Background(), exported.EventConn, "Test", func(ctx context.Context, tmpLWID *LinksWithID, args *utils.RetryFnArgs) error {
		lwid = tmpLWID
		return nil
	}, exported.RetryOptions{})

	require.NoError(t, err)
	require.NotNil(t, lwid)

	require.Equal(t, 3, len(md.Events.GetOpenLinks()), "mgmt link (sender and receiver) + receiver are open")
	require.Equal(t, 1, len(md.Events.GetOpenConns()), "connection is open")

	err = links.Close(context.Background(), false)
	require.NoError(t, err)

	require.NoError(t, ns.Close(true))

	emulation.RequireNoLeaks(t, md.Events)
}

func newAMQPLinksForTest(t *testing.T, mockDataOptions emulation.MockDataOptions, createLinkFunc CreateLinkFunc) (*emulation.MockData, *AMQPLinksImpl, *Namespace, func()) {
	ns, err := NewNamespace(
		NamespaceWithConnectionString("Endpoint=sb://example.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=DEADBEEF"),
	)
	require.NoError(t, err)

	md := emulation.NewMockData(t, &mockDataOptions)
	ns.newClientFn = md.NewConnection

	tmpLinks := NewAMQPLinks(NewAMQPLinksArgs{
		NS:                  ns,
		EntityPath:          "entity path",
		CreateLinkFunc:      createLinkFunc,
		GetRecoveryKindFunc: GetRecoveryKind,
	})

	links := tmpLinks.(*AMQPLinksImpl)
	links.contextWithTimeoutFn = mock.NewContextWithTimeoutForTests

	return md, links, ns, func() {
		test.RequireLinksClose(t, links)
		test.RequireNSClose(t, ns)
		md.Close()
	}
}

// newLinksForAMQPLinksTest creates a amqpwrap.AMQPSenderCloser and a amqpwrap.AMQPReceiverCloser linkwith the same options
// we use when we create them with the azservicebus.Receiver/Sender.
func newLinksForAMQPLinksTest(entityPath string, session amqpwrap.AMQPSession) (amqpwrap.AMQPSenderCloser, amqpwrap.AMQPReceiverCloser, error) {
	receiverOpts := &amqp.ReceiverOptions{
		SettlementMode: amqp.ReceiverSettleModeSecond.Ptr(),
		Credit:         -1,
	}

	receiver, err := session.NewReceiver(context.Background(), entityPath, receiverOpts)

	if err != nil {
		return nil, nil, err
	}

	sender, err := session.NewSender(
		context.Background(),
		entityPath,
		&amqp.SenderOptions{
			SettlementMode:              amqp.SenderSettleModeMixed.Ptr(),
			RequestedReceiverSettleMode: amqp.ReceiverSettleModeFirst.Ptr(),
		})

	if err != nil {
		_ = receiver.Close(context.Background())
		return nil, nil, err
	}

	return sender, receiver, nil
}
