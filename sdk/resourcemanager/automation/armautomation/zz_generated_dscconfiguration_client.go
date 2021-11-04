//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// DscConfigurationClient contains the methods for the DscConfiguration group.
// Don't use this type directly, use NewDscConfigurationClient() instead.
type DscConfigurationClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewDscConfigurationClient creates a new instance of DscConfigurationClient with the specified values.
func NewDscConfigurationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DscConfigurationClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &DscConfigurationClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Create the configuration identified by configuration name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DscConfigurationClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, parameters string, options *DscConfigurationCreateOrUpdateOptions) (DscConfigurationCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, automationAccountName, configurationName, parameters, options)
	if err != nil {
		return DscConfigurationCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DscConfigurationCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DscConfigurationCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DscConfigurationClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, parameters string, options *DscConfigurationCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations/{configurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	body := streaming.NopCloser(strings.NewReader(parameters))
	return req, req.SetBody(body, "text/plain; encoding=UTF-8")
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DscConfigurationClient) createOrUpdateHandleResponse(resp *http.Response) (DscConfigurationCreateOrUpdateResponse, error) {
	result := DscConfigurationCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscConfiguration); err != nil {
		return DscConfigurationCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DscConfigurationClient) createOrUpdateHandleError(resp *http.Response) error {
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

// CreateOrUpdateWithDscConfigurationCreateOrUpdateParameters - Create the configuration identified by configuration name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DscConfigurationClient) CreateOrUpdateWithDscConfigurationCreateOrUpdateParameters(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, parameters DscConfigurationCreateOrUpdateParameters, options *DscConfigurationCreateOrUpdateWithDscConfigurationCreateOrUpdateParametersOptions) (DscConfigurationCreateOrUpdateWithDscConfigurationCreateOrUpdateParametersResponse, error) {
	req, err := client.createOrUpdateWithDscConfigurationCreateOrUpdateParametersCreateRequest(ctx, resourceGroupName, automationAccountName, configurationName, parameters, options)
	if err != nil {
		return DscConfigurationCreateOrUpdateWithDscConfigurationCreateOrUpdateParametersResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DscConfigurationCreateOrUpdateWithDscConfigurationCreateOrUpdateParametersResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DscConfigurationCreateOrUpdateWithDscConfigurationCreateOrUpdateParametersResponse{}, client.createOrUpdateWithDscConfigurationCreateOrUpdateParametersHandleError(resp)
	}
	return client.createOrUpdateWithDscConfigurationCreateOrUpdateParametersHandleResponse(resp)
}

// createOrUpdateWithDscConfigurationCreateOrUpdateParametersCreateRequest creates the CreateOrUpdateWithDscConfigurationCreateOrUpdateParameters request.
func (client *DscConfigurationClient) createOrUpdateWithDscConfigurationCreateOrUpdateParametersCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, parameters DscConfigurationCreateOrUpdateParameters, options *DscConfigurationCreateOrUpdateWithDscConfigurationCreateOrUpdateParametersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations/{configurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateWithDscConfigurationCreateOrUpdateParametersHandleResponse handles the CreateOrUpdateWithDscConfigurationCreateOrUpdateParameters response.
func (client *DscConfigurationClient) createOrUpdateWithDscConfigurationCreateOrUpdateParametersHandleResponse(resp *http.Response) (DscConfigurationCreateOrUpdateWithDscConfigurationCreateOrUpdateParametersResponse, error) {
	result := DscConfigurationCreateOrUpdateWithDscConfigurationCreateOrUpdateParametersResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscConfiguration); err != nil {
		return DscConfigurationCreateOrUpdateWithDscConfigurationCreateOrUpdateParametersResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateWithDscConfigurationCreateOrUpdateParametersHandleError handles the CreateOrUpdateWithDscConfigurationCreateOrUpdateParameters error response.
func (client *DscConfigurationClient) createOrUpdateWithDscConfigurationCreateOrUpdateParametersHandleError(resp *http.Response) error {
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

// Delete - Delete the dsc configuration identified by configuration name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DscConfigurationClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, options *DscConfigurationDeleteOptions) (DscConfigurationDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, automationAccountName, configurationName, options)
	if err != nil {
		return DscConfigurationDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DscConfigurationDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DscConfigurationDeleteResponse{}, client.deleteHandleError(resp)
	}
	return DscConfigurationDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DscConfigurationClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, options *DscConfigurationDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations/{configurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DscConfigurationClient) deleteHandleError(resp *http.Response) error {
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

// Get - Retrieve the configuration identified by configuration name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DscConfigurationClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, options *DscConfigurationGetOptions) (DscConfigurationGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, configurationName, options)
	if err != nil {
		return DscConfigurationGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DscConfigurationGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DscConfigurationGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DscConfigurationClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, options *DscConfigurationGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations/{configurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DscConfigurationClient) getHandleResponse(resp *http.Response) (DscConfigurationGetResponse, error) {
	result := DscConfigurationGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscConfiguration); err != nil {
		return DscConfigurationGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DscConfigurationClient) getHandleError(resp *http.Response) error {
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

// GetContent - Retrieve the configuration script identified by configuration name.
// If the operation fails it returns a generic error.
func (client *DscConfigurationClient) GetContent(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, options *DscConfigurationGetContentOptions) (DscConfigurationGetContentResponse, error) {
	req, err := client.getContentCreateRequest(ctx, resourceGroupName, automationAccountName, configurationName, options)
	if err != nil {
		return DscConfigurationGetContentResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DscConfigurationGetContentResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DscConfigurationGetContentResponse{}, client.getContentHandleError(resp)
	}
	return client.getContentHandleResponse(resp)
}

// getContentCreateRequest creates the GetContent request.
func (client *DscConfigurationClient) getContentCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, options *DscConfigurationGetContentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations/{configurationName}/content"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "text/powershell")
	return req, nil
}

// getContentHandleResponse handles the GetContent response.
func (client *DscConfigurationClient) getContentHandleResponse(resp *http.Response) (DscConfigurationGetContentResponse, error) {
	result := DscConfigurationGetContentResponse{RawResponse: resp}
	return result, nil
}

// getContentHandleError handles the GetContent error response.
func (client *DscConfigurationClient) getContentHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByAutomationAccount - Retrieve a list of configurations.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DscConfigurationClient) ListByAutomationAccount(resourceGroupName string, automationAccountName string, options *DscConfigurationListByAutomationAccountOptions) *DscConfigurationListByAutomationAccountPager {
	return &DscConfigurationListByAutomationAccountPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
		},
		advancer: func(ctx context.Context, resp DscConfigurationListByAutomationAccountResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DscConfigurationListResult.NextLink)
		},
	}
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *DscConfigurationClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *DscConfigurationListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Inlinecount != nil {
		reqQP.Set("$inlinecount", *options.Inlinecount)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *DscConfigurationClient) listByAutomationAccountHandleResponse(resp *http.Response) (DscConfigurationListByAutomationAccountResponse, error) {
	result := DscConfigurationListByAutomationAccountResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscConfigurationListResult); err != nil {
		return DscConfigurationListByAutomationAccountResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByAutomationAccountHandleError handles the ListByAutomationAccount error response.
func (client *DscConfigurationClient) listByAutomationAccountHandleError(resp *http.Response) error {
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

// Update - Create the configuration identified by configuration name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DscConfigurationClient) Update(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, options *DscConfigurationUpdateOptions) (DscConfigurationUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, automationAccountName, configurationName, options)
	if err != nil {
		return DscConfigurationUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DscConfigurationUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DscConfigurationUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *DscConfigurationClient) updateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, options *DscConfigurationUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations/{configurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Parameters != nil {
		body := streaming.NopCloser(strings.NewReader(*options.Parameters))
		return req, req.SetBody(body, "text/plain; encoding=UTF-8")
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *DscConfigurationClient) updateHandleResponse(resp *http.Response) (DscConfigurationUpdateResponse, error) {
	result := DscConfigurationUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscConfiguration); err != nil {
		return DscConfigurationUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *DscConfigurationClient) updateHandleError(resp *http.Response) error {
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

// UpdateWithDscConfigurationUpdateParameters - Create the configuration identified by configuration name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DscConfigurationClient) UpdateWithDscConfigurationUpdateParameters(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, options *DscConfigurationUpdateWithDscConfigurationUpdateParametersOptions) (DscConfigurationUpdateWithDscConfigurationUpdateParametersResponse, error) {
	req, err := client.updateWithDscConfigurationUpdateParametersCreateRequest(ctx, resourceGroupName, automationAccountName, configurationName, options)
	if err != nil {
		return DscConfigurationUpdateWithDscConfigurationUpdateParametersResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DscConfigurationUpdateWithDscConfigurationUpdateParametersResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DscConfigurationUpdateWithDscConfigurationUpdateParametersResponse{}, client.updateWithDscConfigurationUpdateParametersHandleError(resp)
	}
	return client.updateWithDscConfigurationUpdateParametersHandleResponse(resp)
}

// updateWithDscConfigurationUpdateParametersCreateRequest creates the UpdateWithDscConfigurationUpdateParameters request.
func (client *DscConfigurationClient) updateWithDscConfigurationUpdateParametersCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, configurationName string, options *DscConfigurationUpdateWithDscConfigurationUpdateParametersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations/{configurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}

// updateWithDscConfigurationUpdateParametersHandleResponse handles the UpdateWithDscConfigurationUpdateParameters response.
func (client *DscConfigurationClient) updateWithDscConfigurationUpdateParametersHandleResponse(resp *http.Response) (DscConfigurationUpdateWithDscConfigurationUpdateParametersResponse, error) {
	result := DscConfigurationUpdateWithDscConfigurationUpdateParametersResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscConfiguration); err != nil {
		return DscConfigurationUpdateWithDscConfigurationUpdateParametersResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateWithDscConfigurationUpdateParametersHandleError handles the UpdateWithDscConfigurationUpdateParameters error response.
func (client *DscConfigurationClient) updateWithDscConfigurationUpdateParametersHandleError(resp *http.Response) error {
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
