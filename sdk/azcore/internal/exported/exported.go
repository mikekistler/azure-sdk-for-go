//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package exported

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/shared"
)

type nopCloser struct {
	io.ReadSeeker
}

func (n nopCloser) Close() error {
	return nil
}

// NopCloser returns a ReadSeekCloser with a no-op close method wrapping the provided io.ReadSeeker.
// Exported as streaming.NopCloser().
func NopCloser(rs io.ReadSeeker) io.ReadSeekCloser {
	return nopCloser{rs}
}

// HasStatusCode returns true if the Response's status code is one of the specified values.
// Exported as runtime.HasStatusCode().
func HasStatusCode(resp *http.Response, statusCodes ...int) bool {
	if resp == nil {
		return false
	}
	for _, sc := range statusCodes {
		if resp.StatusCode == sc {
			return true
		}
	}
	return false
}

// Payload reads and returns the response body or an error.
// On a successful read, the response body is cached.
// Subsequent reads will access the cached value.
// Exported as runtime.Payload().
func Payload(resp *http.Response) ([]byte, error) {
	// r.Body won't be a nopClosingBytesReader if downloading was skipped
	if buf, ok := resp.Body.(*shared.NopClosingBytesReader); ok {
		return buf.Bytes(), nil
	}
	bytesBody, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	resp.Body = shared.NewNopClosingBytesReader(bytesBody)
	return bytesBody, nil
}

// AccessToken represents an Azure service bearer access token with expiry information.
// Exported as azcore.AccessToken.
type AccessToken struct {
	Token     string
	ExpiresOn time.Time
}

// TokenRequestOptions contain specific parameter that may be used by credentials types when attempting to get a token.
// Exported as policy.TokenRequestOptions.
type TokenRequestOptions struct {
	// Scopes contains the list of permission scopes required for the token.
	Scopes []string

	// TenantID identifies the tenant from which to request the token. azidentity credentials authenticate in
	// their configured default tenants when this field isn't set.
	TenantID string
}

// TokenCredential represents a credential capable of providing an OAuth token.
// Exported as azcore.TokenCredential.
type TokenCredential interface {
	// GetToken requests an access token for the specified set of scopes.
	GetToken(ctx context.Context, options TokenRequestOptions) (AccessToken, error)
}
