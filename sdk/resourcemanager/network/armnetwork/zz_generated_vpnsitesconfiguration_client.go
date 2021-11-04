//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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
	"strings"
)

// VPNSitesConfigurationClient contains the methods for the VPNSitesConfiguration group.
// Don't use this type directly, use NewVPNSitesConfigurationClient() instead.
type VPNSitesConfigurationClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewVPNSitesConfigurationClient creates a new instance of VPNSitesConfigurationClient with the specified values.
func NewVPNSitesConfigurationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VPNSitesConfigurationClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &VPNSitesConfigurationClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginDownload - Gives the sas-url to download the configurations for vpn-sites in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *VPNSitesConfigurationClient) BeginDownload(ctx context.Context, resourceGroupName string, virtualWANName string, request GetVPNSitesConfigurationRequest, options *VPNSitesConfigurationBeginDownloadOptions) (VPNSitesConfigurationDownloadPollerResponse, error) {
	resp, err := client.download(ctx, resourceGroupName, virtualWANName, request, options)
	if err != nil {
		return VPNSitesConfigurationDownloadPollerResponse{}, err
	}
	result := VPNSitesConfigurationDownloadPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VPNSitesConfigurationClient.Download", "location", resp, client.pl, client.downloadHandleError)
	if err != nil {
		return VPNSitesConfigurationDownloadPollerResponse{}, err
	}
	result.Poller = &VPNSitesConfigurationDownloadPoller{
		pt: pt,
	}
	return result, nil
}

// Download - Gives the sas-url to download the configurations for vpn-sites in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *VPNSitesConfigurationClient) download(ctx context.Context, resourceGroupName string, virtualWANName string, request GetVPNSitesConfigurationRequest, options *VPNSitesConfigurationBeginDownloadOptions) (*http.Response, error) {
	req, err := client.downloadCreateRequest(ctx, resourceGroupName, virtualWANName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.downloadHandleError(resp)
	}
	return resp, nil
}

// downloadCreateRequest creates the Download request.
func (client *VPNSitesConfigurationClient) downloadCreateRequest(ctx context.Context, resourceGroupName string, virtualWANName string, request GetVPNSitesConfigurationRequest, options *VPNSitesConfigurationBeginDownloadOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWANName}/vpnConfiguration"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualWANName == "" {
		return nil, errors.New("parameter virtualWANName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualWANName}", url.PathEscape(virtualWANName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, request)
}

// downloadHandleError handles the Download error response.
func (client *VPNSitesConfigurationClient) downloadHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
