//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armpostgresqlflexibleservers

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

func (c *ClientFactory) NewAdministratorsClient() *AdministratorsClient {
	subClient, _ := NewAdministratorsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBackupsClient() *BackupsClient {
	subClient, _ := NewBackupsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewLocationBasedCapabilitiesClient() *LocationBasedCapabilitiesClient {
	subClient, _ := NewLocationBasedCapabilitiesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewServerCapabilitiesClient() *ServerCapabilitiesClient {
	subClient, _ := NewServerCapabilitiesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCheckNameAvailabilityClient() *CheckNameAvailabilityClient {
	subClient, _ := NewCheckNameAvailabilityClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCheckNameAvailabilityWithLocationClient() *CheckNameAvailabilityWithLocationClient {
	subClient, _ := NewCheckNameAvailabilityWithLocationClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewConfigurationsClient() *ConfigurationsClient {
	subClient, _ := NewConfigurationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDatabasesClient() *DatabasesClient {
	subClient, _ := NewDatabasesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewFirewallRulesClient() *FirewallRulesClient {
	subClient, _ := NewFirewallRulesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewServersClient() *ServersClient {
	subClient, _ := NewServersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewMigrationsClient() *MigrationsClient {
	subClient, _ := NewMigrationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPostgreSQLManagementClient() *PostgreSQLManagementClient {
	subClient, _ := NewPostgreSQLManagementClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGetPrivateDNSZoneSuffixClient() *GetPrivateDNSZoneSuffixClient {
	subClient, _ := NewGetPrivateDNSZoneSuffixClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReplicasClient() *ReplicasClient {
	subClient, _ := NewReplicasClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewLogFilesClient() *LogFilesClient {
	subClient, _ := NewLogFilesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewVirtualNetworkSubnetUsageClient() *VirtualNetworkSubnetUsageClient {
	subClient, _ := NewVirtualNetworkSubnetUsageClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewFlexibleServerClient() *FlexibleServerClient {
	subClient, _ := NewFlexibleServerClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewLtrBackupOperationsClient() *LtrBackupOperationsClient {
	subClient, _ := NewLtrBackupOperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
