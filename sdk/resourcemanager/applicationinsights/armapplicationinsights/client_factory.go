//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armapplicationinsights

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential     azcore.TokenCredential
	options        *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

func (c *ClientFactory) NewAnnotationsClient() *AnnotationsClient {
	subClient, _ := NewAnnotationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAPIKeysClient() *APIKeysClient {
	subClient, _ := NewAPIKeysClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewExportConfigurationsClient() *ExportConfigurationsClient {
	subClient, _ := NewExportConfigurationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewComponentCurrentBillingFeaturesClient() *ComponentCurrentBillingFeaturesClient {
	subClient, _ := NewComponentCurrentBillingFeaturesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewComponentQuotaStatusClient() *ComponentQuotaStatusClient {
	subClient, _ := NewComponentQuotaStatusClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewComponentFeatureCapabilitiesClient() *ComponentFeatureCapabilitiesClient {
	subClient, _ := NewComponentFeatureCapabilitiesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewComponentAvailableFeaturesClient() *ComponentAvailableFeaturesClient {
	subClient, _ := NewComponentAvailableFeaturesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProactiveDetectionConfigurationsClient() *ProactiveDetectionConfigurationsClient {
	subClient, _ := NewProactiveDetectionConfigurationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkItemConfigurationsClient() *WorkItemConfigurationsClient {
	subClient, _ := NewWorkItemConfigurationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewFavoritesClient() *FavoritesClient {
	subClient, _ := NewFavoritesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWebTestLocationsClient() *WebTestLocationsClient {
	subClient, _ := NewWebTestLocationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWebTestsClient() *WebTestsClient {
	subClient, _ := NewWebTestsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAnalyticsItemsClient() *AnalyticsItemsClient {
	subClient, _ := NewAnalyticsItemsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkbookTemplatesClient() *WorkbookTemplatesClient {
	subClient, _ := NewWorkbookTemplatesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewMyWorkbooksClient() *MyWorkbooksClient {
	subClient, _ := NewMyWorkbooksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkbooksClient() *WorkbooksClient {
	subClient, _ := NewWorkbooksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewComponentsClient() *ComponentsClient {
	subClient, _ := NewComponentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewComponentLinkedStorageAccountsClient() *ComponentLinkedStorageAccountsClient {
	subClient, _ := NewComponentLinkedStorageAccountsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewLiveTokenClient() *LiveTokenClient {
	subClient, _ := NewLiveTokenClient(c.credential, c.options)
	return subClient
}
