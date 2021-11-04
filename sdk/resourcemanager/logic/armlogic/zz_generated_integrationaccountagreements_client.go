//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// IntegrationAccountAgreementsClient contains the methods for the IntegrationAccountAgreements group.
// Don't use this type directly, use NewIntegrationAccountAgreementsClient() instead.
type IntegrationAccountAgreementsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewIntegrationAccountAgreementsClient creates a new instance of IntegrationAccountAgreementsClient with the specified values.
func NewIntegrationAccountAgreementsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *IntegrationAccountAgreementsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &IntegrationAccountAgreementsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Creates or updates an integration account agreement.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationAccountAgreementsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string, agreement IntegrationAccountAgreement, options *IntegrationAccountAgreementsCreateOrUpdateOptions) (IntegrationAccountAgreementsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, integrationAccountName, agreementName, agreement, options)
	if err != nil {
		return IntegrationAccountAgreementsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountAgreementsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return IntegrationAccountAgreementsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IntegrationAccountAgreementsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string, agreement IntegrationAccountAgreement, options *IntegrationAccountAgreementsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements/{agreementName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if agreementName == "" {
		return nil, errors.New("parameter agreementName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agreementName}", url.PathEscape(agreementName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, agreement)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *IntegrationAccountAgreementsClient) createOrUpdateHandleResponse(resp *http.Response) (IntegrationAccountAgreementsCreateOrUpdateResponse, error) {
	result := IntegrationAccountAgreementsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountAgreement); err != nil {
		return IntegrationAccountAgreementsCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *IntegrationAccountAgreementsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes an integration account agreement.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationAccountAgreementsClient) Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string, options *IntegrationAccountAgreementsDeleteOptions) (IntegrationAccountAgreementsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, integrationAccountName, agreementName, options)
	if err != nil {
		return IntegrationAccountAgreementsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountAgreementsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return IntegrationAccountAgreementsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return IntegrationAccountAgreementsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IntegrationAccountAgreementsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string, options *IntegrationAccountAgreementsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements/{agreementName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if agreementName == "" {
		return nil, errors.New("parameter agreementName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agreementName}", url.PathEscape(agreementName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *IntegrationAccountAgreementsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets an integration account agreement.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationAccountAgreementsClient) Get(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string, options *IntegrationAccountAgreementsGetOptions) (IntegrationAccountAgreementsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, integrationAccountName, agreementName, options)
	if err != nil {
		return IntegrationAccountAgreementsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountAgreementsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationAccountAgreementsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IntegrationAccountAgreementsClient) getCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string, options *IntegrationAccountAgreementsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements/{agreementName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if agreementName == "" {
		return nil, errors.New("parameter agreementName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agreementName}", url.PathEscape(agreementName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IntegrationAccountAgreementsClient) getHandleResponse(resp *http.Response) (IntegrationAccountAgreementsGetResponse, error) {
	result := IntegrationAccountAgreementsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountAgreement); err != nil {
		return IntegrationAccountAgreementsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *IntegrationAccountAgreementsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Gets a list of integration account agreements.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationAccountAgreementsClient) List(resourceGroupName string, integrationAccountName string, options *IntegrationAccountAgreementsListOptions) *IntegrationAccountAgreementsListPager {
	return &IntegrationAccountAgreementsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, integrationAccountName, options)
		},
		advancer: func(ctx context.Context, resp IntegrationAccountAgreementsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.IntegrationAccountAgreementListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *IntegrationAccountAgreementsClient) listCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, options *IntegrationAccountAgreementsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IntegrationAccountAgreementsClient) listHandleResponse(resp *http.Response) (IntegrationAccountAgreementsListResponse, error) {
	result := IntegrationAccountAgreementsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationAccountAgreementListResult); err != nil {
		return IntegrationAccountAgreementsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *IntegrationAccountAgreementsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListContentCallbackURL - Get the content callback url.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationAccountAgreementsClient) ListContentCallbackURL(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string, listContentCallbackURL GetCallbackURLParameters, options *IntegrationAccountAgreementsListContentCallbackURLOptions) (IntegrationAccountAgreementsListContentCallbackURLResponse, error) {
	req, err := client.listContentCallbackURLCreateRequest(ctx, resourceGroupName, integrationAccountName, agreementName, listContentCallbackURL, options)
	if err != nil {
		return IntegrationAccountAgreementsListContentCallbackURLResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountAgreementsListContentCallbackURLResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationAccountAgreementsListContentCallbackURLResponse{}, client.listContentCallbackURLHandleError(resp)
	}
	return client.listContentCallbackURLHandleResponse(resp)
}

// listContentCallbackURLCreateRequest creates the ListContentCallbackURL request.
func (client *IntegrationAccountAgreementsClient) listContentCallbackURLCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, agreementName string, listContentCallbackURL GetCallbackURLParameters, options *IntegrationAccountAgreementsListContentCallbackURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements/{agreementName}/listContentCallbackUrl"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if agreementName == "" {
		return nil, errors.New("parameter agreementName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agreementName}", url.PathEscape(agreementName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, listContentCallbackURL)
}

// listContentCallbackURLHandleResponse handles the ListContentCallbackURL response.
func (client *IntegrationAccountAgreementsClient) listContentCallbackURLHandleResponse(resp *http.Response) (IntegrationAccountAgreementsListContentCallbackURLResponse, error) {
	result := IntegrationAccountAgreementsListContentCallbackURLResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowTriggerCallbackURL); err != nil {
		return IntegrationAccountAgreementsListContentCallbackURLResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listContentCallbackURLHandleError handles the ListContentCallbackURL error response.
func (client *IntegrationAccountAgreementsClient) listContentCallbackURLHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
