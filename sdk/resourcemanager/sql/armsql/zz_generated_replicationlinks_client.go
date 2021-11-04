//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ReplicationLinksClient contains the methods for the ReplicationLinks group.
// Don't use this type directly, use NewReplicationLinksClient() instead.
type ReplicationLinksClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewReplicationLinksClient creates a new instance of ReplicationLinksClient with the specified values.
func NewReplicationLinksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ReplicationLinksClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ReplicationLinksClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Delete - Deletes a database replication link. Cannot be done during failover.
// If the operation fails it returns a generic error.
func (client *ReplicationLinksClient) Delete(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string, options *ReplicationLinksDeleteOptions) (ReplicationLinksDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, databaseName, linkID, options)
	if err != nil {
		return ReplicationLinksDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReplicationLinksDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ReplicationLinksDeleteResponse{}, client.deleteHandleError(resp)
	}
	return ReplicationLinksDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ReplicationLinksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string, options *ReplicationLinksDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/replicationLinks/{linkId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if linkID == "" {
		return nil, errors.New("parameter linkID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkId}", url.PathEscape(linkID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ReplicationLinksClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginFailover - Sets which replica database is primary by failing over from the current primary replica database.
// If the operation fails it returns a generic error.
func (client *ReplicationLinksClient) BeginFailover(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string, options *ReplicationLinksBeginFailoverOptions) (ReplicationLinksFailoverPollerResponse, error) {
	resp, err := client.failover(ctx, resourceGroupName, serverName, databaseName, linkID, options)
	if err != nil {
		return ReplicationLinksFailoverPollerResponse{}, err
	}
	result := ReplicationLinksFailoverPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ReplicationLinksClient.Failover", "", resp, client.pl, client.failoverHandleError)
	if err != nil {
		return ReplicationLinksFailoverPollerResponse{}, err
	}
	result.Poller = &ReplicationLinksFailoverPoller{
		pt: pt,
	}
	return result, nil
}

// Failover - Sets which replica database is primary by failing over from the current primary replica database.
// If the operation fails it returns a generic error.
func (client *ReplicationLinksClient) failover(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string, options *ReplicationLinksBeginFailoverOptions) (*http.Response, error) {
	req, err := client.failoverCreateRequest(ctx, resourceGroupName, serverName, databaseName, linkID, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.failoverHandleError(resp)
	}
	return resp, nil
}

// failoverCreateRequest creates the Failover request.
func (client *ReplicationLinksClient) failoverCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string, options *ReplicationLinksBeginFailoverOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/replicationLinks/{linkId}/failover"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if linkID == "" {
		return nil, errors.New("parameter linkID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkId}", url.PathEscape(linkID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// failoverHandleError handles the Failover error response.
func (client *ReplicationLinksClient) failoverHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginFailoverAllowDataLoss - Sets which replica database is primary by failing over from the current primary replica database. This operation might result
// in data loss.
// If the operation fails it returns a generic error.
func (client *ReplicationLinksClient) BeginFailoverAllowDataLoss(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string, options *ReplicationLinksBeginFailoverAllowDataLossOptions) (ReplicationLinksFailoverAllowDataLossPollerResponse, error) {
	resp, err := client.failoverAllowDataLoss(ctx, resourceGroupName, serverName, databaseName, linkID, options)
	if err != nil {
		return ReplicationLinksFailoverAllowDataLossPollerResponse{}, err
	}
	result := ReplicationLinksFailoverAllowDataLossPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ReplicationLinksClient.FailoverAllowDataLoss", "", resp, client.pl, client.failoverAllowDataLossHandleError)
	if err != nil {
		return ReplicationLinksFailoverAllowDataLossPollerResponse{}, err
	}
	result.Poller = &ReplicationLinksFailoverAllowDataLossPoller{
		pt: pt,
	}
	return result, nil
}

// FailoverAllowDataLoss - Sets which replica database is primary by failing over from the current primary replica database. This operation might result
// in data loss.
// If the operation fails it returns a generic error.
func (client *ReplicationLinksClient) failoverAllowDataLoss(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string, options *ReplicationLinksBeginFailoverAllowDataLossOptions) (*http.Response, error) {
	req, err := client.failoverAllowDataLossCreateRequest(ctx, resourceGroupName, serverName, databaseName, linkID, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.failoverAllowDataLossHandleError(resp)
	}
	return resp, nil
}

// failoverAllowDataLossCreateRequest creates the FailoverAllowDataLoss request.
func (client *ReplicationLinksClient) failoverAllowDataLossCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string, options *ReplicationLinksBeginFailoverAllowDataLossOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/replicationLinks/{linkId}/forceFailoverAllowDataLoss"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if linkID == "" {
		return nil, errors.New("parameter linkID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkId}", url.PathEscape(linkID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// failoverAllowDataLossHandleError handles the FailoverAllowDataLoss error response.
func (client *ReplicationLinksClient) failoverAllowDataLossHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets a replication link.
// If the operation fails it returns a generic error.
func (client *ReplicationLinksClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string, options *ReplicationLinksGetOptions) (ReplicationLinksGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, databaseName, linkID, options)
	if err != nil {
		return ReplicationLinksGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReplicationLinksGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReplicationLinksGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ReplicationLinksClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string, options *ReplicationLinksGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/replicationLinks/{linkId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if linkID == "" {
		return nil, errors.New("parameter linkID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkId}", url.PathEscape(linkID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReplicationLinksClient) getHandleResponse(resp *http.Response) (ReplicationLinksGetResponse, error) {
	result := ReplicationLinksGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReplicationLink); err != nil {
		return ReplicationLinksGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ReplicationLinksClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByDatabase - Gets a list of replication links on database.
// If the operation fails it returns a generic error.
func (client *ReplicationLinksClient) ListByDatabase(resourceGroupName string, serverName string, databaseName string, options *ReplicationLinksListByDatabaseOptions) *ReplicationLinksListByDatabasePager {
	return &ReplicationLinksListByDatabasePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByDatabaseCreateRequest(ctx, resourceGroupName, serverName, databaseName, options)
		},
		advancer: func(ctx context.Context, resp ReplicationLinksListByDatabaseResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ReplicationLinkListResult.NextLink)
		},
	}
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *ReplicationLinksClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, options *ReplicationLinksListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/replicationLinks"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *ReplicationLinksClient) listByDatabaseHandleResponse(resp *http.Response) (ReplicationLinksListByDatabaseResponse, error) {
	result := ReplicationLinksListByDatabaseResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReplicationLinkListResult); err != nil {
		return ReplicationLinksListByDatabaseResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByDatabaseHandleError handles the ListByDatabase error response.
func (client *ReplicationLinksClient) listByDatabaseHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByServer - Gets a list of replication links.
// If the operation fails it returns a generic error.
func (client *ReplicationLinksClient) ListByServer(resourceGroupName string, serverName string, options *ReplicationLinksListByServerOptions) *ReplicationLinksListByServerPager {
	return &ReplicationLinksListByServerPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
		},
		advancer: func(ctx context.Context, resp ReplicationLinksListByServerResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ReplicationLinkListResult.NextLink)
		},
	}
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ReplicationLinksClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ReplicationLinksListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/replicationLinks"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ReplicationLinksClient) listByServerHandleResponse(resp *http.Response) (ReplicationLinksListByServerResponse, error) {
	result := ReplicationLinksListByServerResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReplicationLinkListResult); err != nil {
		return ReplicationLinksListByServerResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByServerHandleError handles the ListByServer error response.
func (client *ReplicationLinksClient) listByServerHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginUnlink - Deletes a database replication link in forced or friendly way.
// If the operation fails it returns a generic error.
func (client *ReplicationLinksClient) BeginUnlink(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string, parameters UnlinkParameters, options *ReplicationLinksBeginUnlinkOptions) (ReplicationLinksUnlinkPollerResponse, error) {
	resp, err := client.unlink(ctx, resourceGroupName, serverName, databaseName, linkID, parameters, options)
	if err != nil {
		return ReplicationLinksUnlinkPollerResponse{}, err
	}
	result := ReplicationLinksUnlinkPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ReplicationLinksClient.Unlink", "", resp, client.pl, client.unlinkHandleError)
	if err != nil {
		return ReplicationLinksUnlinkPollerResponse{}, err
	}
	result.Poller = &ReplicationLinksUnlinkPoller{
		pt: pt,
	}
	return result, nil
}

// Unlink - Deletes a database replication link in forced or friendly way.
// If the operation fails it returns a generic error.
func (client *ReplicationLinksClient) unlink(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string, parameters UnlinkParameters, options *ReplicationLinksBeginUnlinkOptions) (*http.Response, error) {
	req, err := client.unlinkCreateRequest(ctx, resourceGroupName, serverName, databaseName, linkID, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.unlinkHandleError(resp)
	}
	return resp, nil
}

// unlinkCreateRequest creates the Unlink request.
func (client *ReplicationLinksClient) unlinkCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string, parameters UnlinkParameters, options *ReplicationLinksBeginUnlinkOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/replicationLinks/{linkId}/unlink"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if linkID == "" {
		return nil, errors.New("parameter linkID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkId}", url.PathEscape(linkID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, runtime.MarshalAsJSON(req, parameters)
}

// unlinkHandleError handles the Unlink error response.
func (client *ReplicationLinksClient) unlinkHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
