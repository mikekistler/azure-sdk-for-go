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

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/entities/timeline/PostTimelineEntity.json
func ExampleEntitiesGetTimelineClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEntitiesGetTimelineClient().List(ctx, "myRg", "myWorkspace", "e1d3d618-e11f-478b-98e3-bb381539a8e1", armsecurityinsights.EntityTimelineParameters{
		EndTime:        to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-01T00:00:00.000Z"); return t }()),
		NumberOfBucket: to.Ptr[int32](4),
		StartTime:      to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T00:00:00.000Z"); return t }()),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EntityTimelineResponse = armsecurityinsights.EntityTimelineResponse{
	// 	MetaData: &armsecurityinsights.TimelineResultsMetadata{
	// 		Aggregations: []*armsecurityinsights.TimelineAggregation{
	// 			{
	// 				Count: to.Ptr[int32](4),
	// 				Kind: to.Ptr(armsecurityinsights.EntityTimelineKindActivity),
	// 			},
	// 			{
	// 				Count: to.Ptr[int32](2),
	// 				Kind: to.Ptr(armsecurityinsights.EntityTimelineKindSecurityAlert),
	// 			},
	// 			{
	// 				Count: to.Ptr[int32](1),
	// 				Kind: to.Ptr(armsecurityinsights.EntityTimelineKindAnomaly),
	// 		}},
	// 		Errors: []*armsecurityinsights.TimelineError{
	// 			{
	// 				ErrorMessage: to.Ptr("syntax error"),
	// 				Kind: to.Ptr(armsecurityinsights.EntityTimelineKindActivity),
	// 				QueryID: to.Ptr("11067f9f-d6a7-4488-887f-0ba564268879"),
	// 			},
	// 			{
	// 				ErrorMessage: to.Ptr("internal server error"),
	// 				Kind: to.Ptr(armsecurityinsights.EntityTimelineKindSecurityAlert),
	// 		}},
	// 		TotalCount: to.Ptr[int32](6),
	// 	},
	// 	Value: []armsecurityinsights.EntityTimelineItemClassification{
	// 		&armsecurityinsights.SecurityAlertTimelineItem{
	// 			Kind: to.Ptr(armsecurityinsights.EntityTimelineKindSecurityAlert),
	// 			Description: to.Ptr("The alert description"),
	// 			AlertType: to.Ptr("4467341f-fb73-4f99-a9b3-29473532cf5a_c93bf33e-055e-4972-9e7d-f84fe3fb61ae"),
	// 			AzureResourceID: to.Ptr("4467341f-fb73-4f99-a9b3-29473532cf5a_bf7c3a2f-b743-6410-3ff0-ec64b5995d50"),
	// 			DisplayName: to.Ptr("Alert display name"),
	// 			EndTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T23:31:28.02Z"); return t}()),
	// 			ProductName: to.Ptr("Azure Sentinel"),
	// 			Severity: to.Ptr(armsecurityinsights.AlertSeverityMedium),
	// 			StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T23:32:28.01Z"); return t}()),
	// 			TimeGenerated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T23:37:25.8136594Z"); return t}()),
	// 		},
	// 		&armsecurityinsights.ActivityTimelineItem{
	// 			Kind: to.Ptr(armsecurityinsights.EntityTimelineKindActivity),
	// 			BucketEndTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T23:31:28.02Z"); return t}()),
	// 			BucketStartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T21:31:28.02Z"); return t}()),
	// 			Content: to.Ptr("he user has deleted the account 3 time(s)"),
	// 			FirstActivityTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T21:35:28.02Z"); return t}()),
	// 			LastActivityTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T21:35:28.02Z"); return t}()),
	// 			QueryID: to.Ptr("e0459780-ac9d-4b72-8bd4-fecf6b46a0a1"),
	// 			Title: to.Ptr("The user has deleted an account"),
	// 		},
	// 		&armsecurityinsights.AnomalyTimelineItem{
	// 			Kind: to.Ptr(armsecurityinsights.EntityTimelineKindAnomaly),
	// 			Description: to.Ptr("Anomalous private to public port scanning activity with high destination port count along with low port ratio. The ratios are normalized by multiplying them by 10,000 to get them to a more usable value between 0.0 and 1.0."),
	// 			AzureResourceID: to.Ptr("4467341f-fb73-4f99-a9b3-29473532cf5a_d56430ef-f421-2c9c-0b7d-d082285843c6"),
	// 			DisplayName: to.Ptr("(Preview) Anomalous scanning activity"),
	// 			EndTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T23:31:28.02Z"); return t}()),
	// 			Intent: to.Ptr("Discovery"),
	// 			ProductName: to.Ptr("Azure Sentinel"),
	// 			Reasons: []*string{
	// 				to.Ptr("High destination port count"),
	// 				to.Ptr("Low port ratio")},
	// 				StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T23:32:28.01Z"); return t}()),
	// 				Techniques: []*string{
	// 					to.Ptr("T1046")},
	// 					TimeGenerated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T23:37:25.8136594Z"); return t}()),
	// 					Vendor: to.Ptr("Microsoft"),
	// 			}},
	// 		}
}
