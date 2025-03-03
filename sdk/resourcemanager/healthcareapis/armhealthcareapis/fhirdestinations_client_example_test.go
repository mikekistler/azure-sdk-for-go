//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhealthcareapis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/2f2b633d940230cbbd5bcf1339a2e1c48674e4a2/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/preview/2022-01-31-preview/examples/iotconnectors/iotconnector_fhirdestination_List.json
func ExampleFhirDestinationsClient_NewListByIotConnectorPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthcareapis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFhirDestinationsClient().NewListByIotConnectorPager("testRG", "workspace1", "blue", nil)
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
		// page.IotFhirDestinationCollection = armhealthcareapis.IotFhirDestinationCollection{
		// 	Value: []*armhealthcareapis.IotFhirDestination{
		// 		{
		// 			Name: to.Ptr("dest1"),
		// 			Type: to.Ptr("Microsoft.HealthcareApis/workspaces/iotconnectors/fhirdestinations"),
		// 			Etag: to.Ptr("00000000-0000-0000-f5ac-912ca49e01d6"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/testRG/providers/Microsoft.HealthcareApis/workspaces/workspace1/iotconnectors/blue/fhirdestinations/dest1"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armhealthcareapis.IotFhirDestinationProperties{
		// 				ProvisioningState: to.Ptr(armhealthcareapis.ProvisioningStateSucceeded),
		// 				FhirMapping: &armhealthcareapis.IotMappingProperties{
		// 					Content: map[string]any{
		// 						"template":[]any{
		// 							map[string]any{
		// 								"template":map[string]any{
		// 									"codes":[]any{
		// 										map[string]any{
		// 											"code": "8867-4",
		// 											"display": "Heart rate",
		// 											"system": "http://loinc.org",
		// 										},
		// 									},
		// 									"periodInterval": float64(60),
		// 									"typeName": "heartrate",
		// 									"value":map[string]any{
		// 										"defaultPeriod": float64(5000),
		// 										"unit": "count/min",
		// 										"valueName": "hr",
		// 										"valueType": "SampledData",
		// 									},
		// 								},
		// 								"templateType": "CodeValueFhir",
		// 							},
		// 						},
		// 						"templateType": "CollectionFhirTemplate",
		// 					},
		// 				},
		// 				FhirServiceResourceID: to.Ptr("subscriptions/11111111-2222-3333-4444-555566667777/resourceGroups/myrg/providers/Microsoft.HealthcareApis/workspaces/myworkspace/fhirservices/myfhirservice"),
		// 				ResourceIdentityResolutionType: to.Ptr(armhealthcareapis.IotIdentityResolutionTypeCreate),
		// 			},
		// 			SystemData: &armhealthcareapis.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-28T19:26:24.072Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-28T19:26:24.072Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("dest2"),
		// 			Type: to.Ptr("Microsoft.HealthcareApis/workspaces/iotconnectors/fhirdestinations"),
		// 			Etag: to.Ptr("00000000-0000-0000-f6ac-912ca49e01d6"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/testRG/providers/Microsoft.HealthcareApis/workspaces/workspace1/iotconnectors/blue/fhirdestinations/dest2"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armhealthcareapis.IotFhirDestinationProperties{
		// 				ProvisioningState: to.Ptr(armhealthcareapis.ProvisioningStateSucceeded),
		// 				FhirMapping: &armhealthcareapis.IotMappingProperties{
		// 					Content: map[string]any{
		// 						"template":[]any{
		// 							map[string]any{
		// 								"template":map[string]any{
		// 									"codes":[]any{
		// 										map[string]any{
		// 											"code": "8867-4",
		// 											"display": "Heart rate",
		// 											"system": "http://loinc.org",
		// 										},
		// 									},
		// 									"periodInterval": float64(60),
		// 									"typeName": "heartrate",
		// 									"value":map[string]any{
		// 										"defaultPeriod": float64(5000),
		// 										"unit": "count/min",
		// 										"valueName": "hr",
		// 										"valueType": "SampledData",
		// 									},
		// 								},
		// 								"templateType": "CodeValueFhir",
		// 							},
		// 						},
		// 						"templateType": "CollectionFhirTemplate",
		// 					},
		// 				},
		// 				FhirServiceResourceID: to.Ptr("subscriptions/11111111-2222-3333-4444-555566667777/resourceGroups/myrg/providers/Microsoft.HealthcareApis/workspaces/myworkspace/fhirservices/myfhirservice"),
		// 				ResourceIdentityResolutionType: to.Ptr(armhealthcareapis.IotIdentityResolutionTypeLookup),
		// 			},
		// 			SystemData: &armhealthcareapis.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-27T19:26:24.072Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-27T19:26:24.072Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 			},
		// 	}},
		// }
	}
}
