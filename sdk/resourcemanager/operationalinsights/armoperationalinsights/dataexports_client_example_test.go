//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf3574813e15bb33b3cb610f44edfcbebd8b1b23/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/DataExportListByWorkspace.json
func ExampleDataExportsClient_NewListByWorkspacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataExportsClient().NewListByWorkspacePager("RgTest1", "DeWnTest1234", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.DataExportListResult = armoperationalinsights.DataExportListResult{
		// 	Value: []*armoperationalinsights.DataExport{
		// 		{
		// 			Name: to.Ptr("export1"),
		// 			Type: to.Ptr("Microsoft.OperationalInsights/workspaces/export"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourcegroups/RgTest1/providers/microsoft.operationalinsights/workspaces/DeWnTest1234/export/export1"),
		// 			Properties: &armoperationalinsights.DataExportProperties{
		// 				CreatedDate: to.Ptr("Sun, 12 Jan 2020 12:51:10 GMT"),
		// 				DataExportID: to.Ptr("d5233afc-7829-4b89-c594-08d7975e19a5"),
		// 				Destination: &armoperationalinsights.Destination{
		// 					Type: to.Ptr(armoperationalinsights.TypeEventHub),
		// 					ResourceID: to.Ptr("/subscriptions/192b9f85-a39a-4276-b96d-d5cd351703f9/resourceGroups/OIAutoRest1234/providers/Microsoft.EventHub/namespaces/test"),
		// 				},
		// 				Enable: to.Ptr(true),
		// 				LastModifiedDate: to.Ptr("Sun, 12 Jan 2020 12:51:10 GMT"),
		// 				TableNames: []*string{
		// 					to.Ptr("Heartbeat")},
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf3574813e15bb33b3cb610f44edfcbebd8b1b23/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/DataExportCreateOrUpdate.json
func ExampleDataExportsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataExportsClient().CreateOrUpdate(ctx, "RgTest1", "DeWnTest1234", "export1", armoperationalinsights.DataExport{
		Properties: &armoperationalinsights.DataExportProperties{
			Destination: &armoperationalinsights.Destination{
				ResourceID: to.Ptr("/subscriptions/192b9f85-a39a-4276-b96d-d5cd351703f9/resourceGroups/OIAutoRest1234/providers/Microsoft.EventHub/namespaces/test"),
			},
			TableNames: []*string{
				to.Ptr("Heartbeat")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataExport = armoperationalinsights.DataExport{
	// 	Name: to.Ptr("export1"),
	// 	Type: to.Ptr("Microsoft.OperationalInsights/workspaces/export"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourcegroups/RgTest1/providers/microsoft.operationalinsights/workspaces/DeWnTest1234/export/export1"),
	// 	Properties: &armoperationalinsights.DataExportProperties{
	// 		CreatedDate: to.Ptr("Sun, 12 Jan 2020 12:51:10 GMT"),
	// 		DataExportID: to.Ptr("d5233afc-7829-4b89-c594-08d7975e19a5"),
	// 		Destination: &armoperationalinsights.Destination{
	// 			Type: to.Ptr(armoperationalinsights.TypeEventHub),
	// 			ResourceID: to.Ptr("/subscriptions/192b9f85-a39a-4276-b96d-d5cd351703f9/resourceGroups/OIAutoRest1234/providers/Microsoft.EventHub/namespaces/test"),
	// 		},
	// 		Enable: to.Ptr(true),
	// 		LastModifiedDate: to.Ptr("Sun, 12 Jan 2020 12:51:10 GMT"),
	// 		TableNames: []*string{
	// 			to.Ptr("Heartbeat")},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf3574813e15bb33b3cb610f44edfcbebd8b1b23/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/DataExportGet.json
func ExampleDataExportsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataExportsClient().Get(ctx, "RgTest1", "DeWnTest1234", "export1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataExport = armoperationalinsights.DataExport{
	// 	Name: to.Ptr("export1"),
	// 	Type: to.Ptr("Microsoft.OperationalInsights/workspaces/export"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourcegroups/RgTest1/providers/microsoft.operationalinsights/workspaces/DeWnTest1234/export/export1"),
	// 	Properties: &armoperationalinsights.DataExportProperties{
	// 		CreatedDate: to.Ptr("Sun, 12 Jan 2020 12:51:10 GMT"),
	// 		DataExportID: to.Ptr("d5233afc-7829-4b89-c594-08d7975e19a5"),
	// 		Destination: &armoperationalinsights.Destination{
	// 			Type: to.Ptr(armoperationalinsights.TypeEventHub),
	// 			ResourceID: to.Ptr("/subscriptions/192b9f85-a39a-4276-b96d-d5cd351703f9/resourceGroups/OIAutoRest1234/providers/Microsoft.EventHub/namespaces/test"),
	// 		},
	// 		Enable: to.Ptr(true),
	// 		LastModifiedDate: to.Ptr("Sun, 12 Jan 2020 12:51:10 GMT"),
	// 		TableNames: []*string{
	// 			to.Ptr("Heartbeat")},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf3574813e15bb33b3cb610f44edfcbebd8b1b23/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/DataExportDelete.json
func ExampleDataExportsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDataExportsClient().Delete(ctx, "RgTest1", "DeWnTest1234", "export1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
