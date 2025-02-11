//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/fileImports/GetFileImports.json
func ExampleFileImportsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFileImportsClient().NewListPager("myRg", "myWorkspace", &armsecurityinsights.FileImportsClientListOptions{Filter: nil,
		Orderby:   to.Ptr("properties/createdTimeUtc desc"),
		Top:       to.Ptr[int32](1),
		SkipToken: nil,
	})
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
		// page.FileImportList = armsecurityinsights.FileImportList{
		// 	Value: []*armsecurityinsights.FileImport{
		// 		{
		// 			Name: to.Ptr("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/FileImports"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/FileImports/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
		// 			Properties: &armsecurityinsights.FileImportProperties{
		// 				ContentType: to.Ptr(armsecurityinsights.FileImportContentTypeStixIndicator),
		// 				CreatedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-25T21:02:38.8350631Z"); return t}()),
		// 				FilesValidUntilTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-26T21:02:38.8350632Z"); return t}()),
		// 				ImportFile: &armsecurityinsights.FileMetadata{
		// 					DeleteStatus: to.Ptr(armsecurityinsights.DeleteStatusNotDeleted),
		// 					FileFormat: to.Ptr(armsecurityinsights.FileFormatJSON),
		// 					FileName: to.Ptr("fileName.json"),
		// 					FileSize: to.Ptr[int32](5146),
		// 				},
		// 				ImportValidUntilTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-24T21:02:38.8350636Z"); return t}()),
		// 				IngestedRecordCount: to.Ptr[int32](5),
		// 				IngestionMode: to.Ptr(armsecurityinsights.IngestionModeIngestAnyValidRecords),
		// 				Source: to.Ptr("mySource"),
		// 				State: to.Ptr(armsecurityinsights.FileImportStateIngested),
		// 				TotalRecordCount: to.Ptr[int32](5),
		// 				ValidRecordCount: to.Ptr[int32](5),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/fileImports/GetFileImportById.json
func ExampleFileImportsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileImportsClient().Get(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FileImport = armsecurityinsights.FileImport{
	// 	Name: to.Ptr("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 	Type: to.Ptr("Microsoft.SecurityInsights/FileImports"),
	// 	ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/FileImports/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 	Properties: &armsecurityinsights.FileImportProperties{
	// 		ContentType: to.Ptr(armsecurityinsights.FileImportContentTypeStixIndicator),
	// 		CreatedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-25T21:02:38.8350631Z"); return t}()),
	// 		FilesValidUntilTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-26T21:02:38.8350632Z"); return t}()),
	// 		ImportFile: &armsecurityinsights.FileMetadata{
	// 			DeleteStatus: to.Ptr(armsecurityinsights.DeleteStatusNotDeleted),
	// 			FileContentURI: to.Ptr("https://sentinelimportswus2.blob.core.windows.net/78c2e51a-3cd3-4ca0-a2d4-e7effb9a05fe/43967a5e-47a7-474e-afb8-2081e9b99ca1/myFile.json?skoid=40ca3ff4-ed1d-4c65-a409-8c6caff8a6d5&sktid=72f988bf-86f1-41af-91ab-2d7cd011db47&skt=2022-03-25T21%3A12%3A51Z&ske=2022-03-25T22%3A12%3A51Z&sks=b&skv=2020-10-02&sv=2020-08-04&st=2022-03-25T21%3A12%3A51Z&se=2022-03-25T22%3A12%3A51Z&sr=b&sp=c&sig=5n0D%2FERS6ZOQdfdO2adleeSVOM4b6mQeds%2FWYCGm9pU%3D"),
	// 			FileFormat: to.Ptr(armsecurityinsights.FileFormatJSON),
	// 			FileName: to.Ptr("myFile.json"),
	// 			FileSize: to.Ptr[int32](5146),
	// 		},
	// 		ImportValidUntilTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-04-24T21:02:38.8350636Z"); return t}()),
	// 		IngestedRecordCount: to.Ptr[int32](5),
	// 		IngestionMode: to.Ptr(armsecurityinsights.IngestionModeIngestAnyValidRecords),
	// 		Source: to.Ptr("mySource"),
	// 		State: to.Ptr(armsecurityinsights.FileImportStateIngested),
	// 		TotalRecordCount: to.Ptr[int32](5),
	// 		ValidRecordCount: to.Ptr[int32](5),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/fileImports/CreateFileImport.json
func ExampleFileImportsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewFileImportsClient().Create(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", armsecurityinsights.FileImport{
		Properties: &armsecurityinsights.FileImportProperties{
			ContentType: to.Ptr(armsecurityinsights.FileImportContentTypeStixIndicator),
			ImportFile: &armsecurityinsights.FileMetadata{
				FileFormat: to.Ptr(armsecurityinsights.FileFormatJSON),
				FileName:   to.Ptr("myFile.json"),
				FileSize:   to.Ptr[int32](4653),
			},
			IngestionMode: to.Ptr(armsecurityinsights.IngestionModeIngestAnyValidRecords),
			Source:        to.Ptr("mySource"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/fileImports/DeleteFileImport.json
func ExampleFileImportsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFileImportsClient().BeginDelete(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
