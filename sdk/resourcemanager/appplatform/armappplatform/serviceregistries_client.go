//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappplatform

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ServiceRegistriesClient contains the methods for the ServiceRegistries group.
// Don't use this type directly, use NewServiceRegistriesClient() instead.
type ServiceRegistriesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServiceRegistriesClient creates a new instance of ServiceRegistriesClient with the specified values.
//   - subscriptionID - Gets subscription ID which uniquely identify the Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServiceRegistriesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServiceRegistriesClient, error) {
	cl, err := arm.NewClient(moduleName+".ServiceRegistriesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServiceRegistriesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create the default Service Registry or update the existing Service Registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - serviceRegistryName - The name of Service Registry.
//   - options - ServiceRegistriesClientBeginCreateOrUpdateOptions contains the optional parameters for the ServiceRegistriesClient.BeginCreateOrUpdate
//     method.
func (client *ServiceRegistriesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, serviceRegistryName string, options *ServiceRegistriesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ServiceRegistriesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceName, serviceRegistryName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ServiceRegistriesClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ServiceRegistriesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create the default Service Registry or update the existing Service Registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
func (client *ServiceRegistriesClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, serviceRegistryName string, options *ServiceRegistriesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, serviceRegistryName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServiceRegistriesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, serviceRegistryName string, options *ServiceRegistriesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/serviceRegistries/{serviceRegistryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if serviceRegistryName == "" {
		return nil, errors.New("parameter serviceRegistryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceRegistryName}", url.PathEscape(serviceRegistryName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDelete - Disable the default Service Registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - serviceRegistryName - The name of Service Registry.
//   - options - ServiceRegistriesClientBeginDeleteOptions contains the optional parameters for the ServiceRegistriesClient.BeginDelete
//     method.
func (client *ServiceRegistriesClient) BeginDelete(ctx context.Context, resourceGroupName string, serviceName string, serviceRegistryName string, options *ServiceRegistriesClientBeginDeleteOptions) (*runtime.Poller[ServiceRegistriesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serviceName, serviceRegistryName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ServiceRegistriesClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ServiceRegistriesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Disable the default Service Registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
func (client *ServiceRegistriesClient) deleteOperation(ctx context.Context, resourceGroupName string, serviceName string, serviceRegistryName string, options *ServiceRegistriesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, serviceRegistryName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServiceRegistriesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, serviceRegistryName string, options *ServiceRegistriesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/serviceRegistries/{serviceRegistryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if serviceRegistryName == "" {
		return nil, errors.New("parameter serviceRegistryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceRegistryName}", url.PathEscape(serviceRegistryName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the Service Registry and its properties.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - serviceRegistryName - The name of Service Registry.
//   - options - ServiceRegistriesClientGetOptions contains the optional parameters for the ServiceRegistriesClient.Get method.
func (client *ServiceRegistriesClient) Get(ctx context.Context, resourceGroupName string, serviceName string, serviceRegistryName string, options *ServiceRegistriesClientGetOptions) (ServiceRegistriesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, serviceRegistryName, options)
	if err != nil {
		return ServiceRegistriesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceRegistriesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceRegistriesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServiceRegistriesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, serviceRegistryName string, options *ServiceRegistriesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/serviceRegistries/{serviceRegistryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if serviceRegistryName == "" {
		return nil, errors.New("parameter serviceRegistryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceRegistryName}", url.PathEscape(serviceRegistryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServiceRegistriesClient) getHandleResponse(resp *http.Response) (ServiceRegistriesClientGetResponse, error) {
	result := ServiceRegistriesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceRegistryResource); err != nil {
		return ServiceRegistriesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Handles requests to list all resources in a Service.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - options - ServiceRegistriesClientListOptions contains the optional parameters for the ServiceRegistriesClient.NewListPager
//     method.
func (client *ServiceRegistriesClient) NewListPager(resourceGroupName string, serviceName string, options *ServiceRegistriesClientListOptions) *runtime.Pager[ServiceRegistriesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServiceRegistriesClientListResponse]{
		More: func(page ServiceRegistriesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServiceRegistriesClientListResponse) (ServiceRegistriesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServiceRegistriesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ServiceRegistriesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServiceRegistriesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ServiceRegistriesClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *ServiceRegistriesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/serviceRegistries"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServiceRegistriesClient) listHandleResponse(resp *http.Response) (ServiceRegistriesClientListResponse, error) {
	result := ServiceRegistriesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceRegistryResourceCollection); err != nil {
		return ServiceRegistriesClientListResponse{}, err
	}
	return result, nil
}
