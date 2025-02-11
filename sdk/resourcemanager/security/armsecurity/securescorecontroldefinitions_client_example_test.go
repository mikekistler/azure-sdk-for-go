//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/af3f7994582c0cbd61a48b636907ad2ac95d332c/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/secureScoreControlDefinitions/ListSecureScoreControlDefinitions_example.json
func ExampleSecureScoreControlDefinitionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSecureScoreControlDefinitionsClient().NewListPager(nil)
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
		// page.SecureScoreControlDefinitionList = armsecurity.SecureScoreControlDefinitionList{
		// 	Value: []*armsecurity.SecureScoreControlDefinitionItem{
		// 		{
		// 			Name: to.Ptr("a9909064-42b4-4d34-8143-275477afe18b"),
		// 			Type: to.Ptr("Microsoft.Security/SecureScoreControlDefinitions"),
		// 			ID: to.Ptr("/providers/Microsoft.Security/SecureScoreControlDefinitions/a9909064-42b4-4d34-8143-275477afe18b"),
		// 			Properties: &armsecurity.SecureScoreControlDefinitionItemProperties{
		// 				Description: to.Ptr("This control contains recommendations regarding DDoS attack prevention"),
		// 				AssessmentDefinitions: []*armsecurity.AzureResourceLink{
		// 					{
		// 						ID: to.Ptr("/providers/Microsoft.Security/assessmentMetadata/e3de1cc0-f4dd-3b34-e496-8b5381ba2d70"),
		// 				}},
		// 				DisplayName: to.Ptr("Protect applications against DDoS attacks"),
		// 				MaxScore: to.Ptr[int32](2),
		// 				Source: &armsecurity.SecureScoreControlDefinitionSource{
		// 					SourceType: to.Ptr(armsecurity.ControlTypeBuiltIn),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("a000c66f-6da2-4f9d-826d-2364347d2588"),
		// 			Type: to.Ptr("Microsoft.Security/SecureScoreControlDefinitions"),
		// 			ID: to.Ptr("/providers/Microsoft.Security/SecureScoreControlDefinitions/a000c66f-6da2-4f9d-826d-2364347d2588"),
		// 			Properties: &armsecurity.SecureScoreControlDefinitionItemProperties{
		// 				Description: to.Ptr("This control contains recommendations regarding enabling adaptive application control"),
		// 				AssessmentDefinitions: []*armsecurity.AzureResourceLink{
		// 					{
		// 						ID: to.Ptr("/providers/Microsoft.Security/assessmentMetadata/35f45c95-27cf-4e52-891f-8390d1de5828"),
		// 					},
		// 					{
		// 						ID: to.Ptr("/providers/Microsoft.Security/assessmentMetadata/d1db3318-01ff-16de-29eb-28b344515626"),
		// 					},
		// 					{
		// 						ID: to.Ptr("/providers/Microsoft.Security/assessmentMetadata/e7ee30c4-bac9-2966-54bd-2023a4282872"),
		// 					},
		// 					{
		// 						ID: to.Ptr("/providers/Microsoft.Security/assessmentMetadata/8e2b96ff-3de2-289b-b5c1-3b9921a3441e"),
		// 				}},
		// 				DisplayName: to.Ptr("Apply adaptive application control"),
		// 				MaxScore: to.Ptr[int32](3),
		// 				Source: &armsecurity.SecureScoreControlDefinitionSource{
		// 					SourceType: to.Ptr(armsecurity.ControlTypeBuiltIn),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/af3f7994582c0cbd61a48b636907ad2ac95d332c/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/secureScoreControlDefinitions/ListSecureScoreControlDefinitions_subscription_example.json
func ExampleSecureScoreControlDefinitionsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSecureScoreControlDefinitionsClient().NewListBySubscriptionPager(nil)
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
		// page.SecureScoreControlDefinitionList = armsecurity.SecureScoreControlDefinitionList{
		// 	Value: []*armsecurity.SecureScoreControlDefinitionItem{
		// 		{
		// 			Name: to.Ptr("a9909064-42b4-4d34-8143-275477afe18b"),
		// 			Type: to.Ptr("Microsoft.Security/SecureScoreControlDefinitions"),
		// 			ID: to.Ptr("/providers/Microsoft.Security/SecureScoreControlDefinitions/a9909064-42b4-4d34-8143-275477afe18b"),
		// 			Properties: &armsecurity.SecureScoreControlDefinitionItemProperties{
		// 				Description: to.Ptr("This control contains recommendations regarding DDoS attack prevention"),
		// 				AssessmentDefinitions: []*armsecurity.AzureResourceLink{
		// 					{
		// 						ID: to.Ptr("/providers/Microsoft.Security/assessmentMetadata/e3de1cc0-f4dd-3b34-e496-8b5381ba2d70"),
		// 				}},
		// 				DisplayName: to.Ptr("Protect applications against DDoS attacks"),
		// 				MaxScore: to.Ptr[int32](2),
		// 				Source: &armsecurity.SecureScoreControlDefinitionSource{
		// 					SourceType: to.Ptr(armsecurity.ControlTypeBuiltIn),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("34a42fb3-e6db-409c-b56b-7b1db6b8aee0"),
		// 			Type: to.Ptr("Microsoft.Security/SecureScoreControlDefinitions"),
		// 			ID: to.Ptr("/providers/Microsoft.Security/SecureScoreControlDefinitions/34a42fb3-e6db-409c-b56b-7b1db6b8aee0"),
		// 			Properties: &armsecurity.SecureScoreControlDefinitionItemProperties{
		// 				Description: to.Ptr("This control contains recommendations regarding MFA"),
		// 				AssessmentDefinitions: []*armsecurity.AzureResourceLink{
		// 					{
		// 						ID: to.Ptr("/providers/Microsoft.Security/assessmentMetadata/151e82c5-5341-a74b-1eb0-bc38d2c84bb5"),
		// 					},
		// 					{
		// 						ID: to.Ptr("/providers/Microsoft.Security/assessmentMetadata/57e98606-6b1e-6193-0e3d-fe621387c16b"),
		// 					},
		// 					{
		// 						ID: to.Ptr("/providers/Microsoft.Security/assessmentMetadata/94290b00-4d0c-d7b4-7cea-064a9554e681"),
		// 				}},
		// 				DisplayName: to.Ptr("Enable MFA"),
		// 				MaxScore: to.Ptr[int32](10),
		// 				Source: &armsecurity.SecureScoreControlDefinitionSource{
		// 					SourceType: to.Ptr(armsecurity.ControlTypeBuiltIn),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("a000c66f-6da2-4f9d-826d-2364347d2588"),
		// 			Type: to.Ptr("Microsoft.Security/SecureScoreControlDefinitions"),
		// 			ID: to.Ptr("/providers/Microsoft.Security/SecureScoreControlDefinitions/a000c66f-6da2-4f9d-826d-2364347d2588"),
		// 			Properties: &armsecurity.SecureScoreControlDefinitionItemProperties{
		// 				Description: to.Ptr("This control contains recommendations regarding enabling adaptive application control"),
		// 				AssessmentDefinitions: []*armsecurity.AzureResourceLink{
		// 					{
		// 						ID: to.Ptr("/providers/Microsoft.Security/assessmentMetadata/35f45c95-27cf-4e52-891f-8390d1de5828"),
		// 					},
		// 					{
		// 						ID: to.Ptr("/providers/Microsoft.Security/assessmentMetadata/d1db3318-01ff-16de-29eb-28b344515626"),
		// 					},
		// 					{
		// 						ID: to.Ptr("/providers/Microsoft.Security/assessmentMetadata/e7ee30c4-bac9-2966-54bd-2023a4282872"),
		// 					},
		// 					{
		// 						ID: to.Ptr("/providers/Microsoft.Security/assessmentMetadata/8e2b96ff-3de2-289b-b5c1-3b9921a3441e"),
		// 				}},
		// 				DisplayName: to.Ptr("Apply adaptive application control"),
		// 				MaxScore: to.Ptr[int32](3),
		// 				Source: &armsecurity.SecureScoreControlDefinitionSource{
		// 					SourceType: to.Ptr(armsecurity.ControlTypeBuiltIn),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
