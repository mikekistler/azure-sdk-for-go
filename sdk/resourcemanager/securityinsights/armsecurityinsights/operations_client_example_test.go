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

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/operations/ListOperations.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationsList = armsecurityinsights.OperationsList{
		// 	Value: []*armsecurityinsights.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/operations/read"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Gets operations"),
		// 				Operation: to.Ptr("Get Operations"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Operations"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/automationRules/read"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Gets an automation rule"),
		// 				Operation: to.Ptr("Get Automation Rules"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("AutomationRules"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/automationRules/write"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Updates an automation rule"),
		// 				Operation: to.Ptr("Update Automation Rules"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("AutomationRules"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/automationRules/delete"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Deletes an automation rule"),
		// 				Operation: to.Ptr("Delete Automation Rules"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("AutomationRules"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/Bookmarks/read"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Gets bookmarks"),
		// 				Operation: to.Ptr("Get Bookmarks"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Bookmarks"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/Bookmarks/write"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Updates bookmarks"),
		// 				Operation: to.Ptr("Update Bookmarks"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Bookmarks"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/Bookmarks/delete"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Deletes bookmarks"),
		// 				Operation: to.Ptr("Delete Bookmarks"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Bookmarks"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/Bookmarks/expand/action"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Gets related entities of an entity by a specific expansion"),
		// 				Operation: to.Ptr("Expand on entity"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Bookmarks"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/bookmarks/relations/read"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Gets a bookmark relation"),
		// 				Operation: to.Ptr("Get Bookmark Relations"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Bookmark Relations"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/bookmarks/relations/write"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Updates a bookmark relation"),
		// 				Operation: to.Ptr("Update Bookmark Relations"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Bookmark Relations"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/bookmarks/relations/delete"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Deletes a bookmark relation"),
		// 				Operation: to.Ptr("Delete Bookmark Relations"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Bookmark Relations"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/alertRules/read"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Gets the alert rules"),
		// 				Operation: to.Ptr("Get Alert Rules"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Alert Rules"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/alertRules/write"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Updates alert rules"),
		// 				Operation: to.Ptr("Update Alert Rules"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Alert Rules"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/alertRules/delete"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Deletes alert rules"),
		// 				Operation: to.Ptr("Delete Alert Rules"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Alert Rules"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/alertRules/actions/read"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Gets the response actions of an alert rule"),
		// 				Operation: to.Ptr("Get Alert Rule Response Actions"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Alert Rules Actions"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/alertRules/actions/write"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Updates the response actions of an alert rule"),
		// 				Operation: to.Ptr("Update Alert Rule Response Actions"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Alert Rules Actions"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/alertRules/actions/delete"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Deletes the response actions of an alert rule"),
		// 				Operation: to.Ptr("Delete Alert Rule Response Actions"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Alert Rules Actions"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/dataConnectors/read"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Gets the data connectors"),
		// 				Operation: to.Ptr("Get Data Connectors"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("DataConnectors"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/dataConnectors/write"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Updates a data connector"),
		// 				Operation: to.Ptr("Update Data Connectors"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("DataConnectors"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/dataConnectors/delete"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Deletes a data connector"),
		// 				Operation: to.Ptr("Delete a Data Connector"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("DataConnectors"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/dataConnectorsCheckRequirements/action"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Check user authorization and license"),
		// 				Operation: to.Ptr("Check user authorization and license"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("DataConnectorsCheckRequirements"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/incidents/read"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Gets an incident"),
		// 				Operation: to.Ptr("Get Incidents"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Incidents"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/incidents/write"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Updates an incident"),
		// 				Operation: to.Ptr("Update Incidents"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Incidents"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/incidents/delete"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Deletes an incident"),
		// 				Operation: to.Ptr("Delete Incidents"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Incidents"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/incidents/comments/read"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Gets the incident comments"),
		// 				Operation: to.Ptr("Get Incident Comments"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Incident Comments"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/incidents/comments/write"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Creates a comment on the incident"),
		// 				Operation: to.Ptr("Create Incident Comments"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Incident Comments"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/incidents/comments/delete"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Deletes a comment on the incident"),
		// 				Operation: to.Ptr("Delete Incident Comment"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Incident Comments"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/incidents/relations/read"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Gets a relation between the incident and related resources"),
		// 				Operation: to.Ptr("Get Incident Relations"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Incident Relations"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/incidents/relations/write"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Updates a relation between the incident and related resources"),
		// 				Operation: to.Ptr("Update Incident Relations"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Incident Relations"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/incidents/relations/delete"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Deletes a relation between the incident and related resources"),
		// 				Operation: to.Ptr("Delete Incident Relations"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Incident Relations"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/threatintelligence/read"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Gets Threat Intelligence"),
		// 				Operation: to.Ptr("Get Threat Intelligence"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("ThreatIntelligence"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/threatintelligence/write"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Updates Threat Intelligence"),
		// 				Operation: to.Ptr("Update Threat Intelligence"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("ThreatIntelligence"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/threatintelligence/delete"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Deletes Threat Intelligence"),
		// 				Operation: to.Ptr("Delete Threat Intelligence"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("ThreatIntelligence"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/threatintelligence/query/action"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Query Threat Intelligence"),
		// 				Operation: to.Ptr("Query Threat Intelligence"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("ThreatIntelligence"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/threatintelligence/metrics/action"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Collect Threat Intelligence Metrics"),
		// 				Operation: to.Ptr("Collect Threat Intelligence Metrics"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("ThreatIntelligence"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/threatintelligence/bulkDelete/action"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Bulk Delete Threat Intelligence"),
		// 				Operation: to.Ptr("Bulk Delete Threat Intelligence"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("ThreatIntelligence"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/threatintelligence/bulkTag/action"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Bulk Tags Threat Intelligence"),
		// 				Operation: to.Ptr("Bulk Tags Threat Intelligence"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("ThreatIntelligence"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/threatintelligence/indicators/write"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Updates Threat Intelligence Indicators"),
		// 				Operation: to.Ptr("Update Threat Intelligence Indicators"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("ThreatIntelligence"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/threatintelligence/indicators/delete"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Deletes Threat Intelligence Indicators"),
		// 				Operation: to.Ptr("Delete Threat Intelligence Indicators"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("ThreatIntelligence"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/threatintelligence/indicators/query/action"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Query Threat Intelligence Indicators"),
		// 				Operation: to.Ptr("Query Threat Intelligence Indicators"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("ThreatIntelligence"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/threatintelligence/indicators/metrics/action"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Get Threat Intelligence Indicator Metrics"),
		// 				Operation: to.Ptr("Get Threat Intelligence Indicator Metrics"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("ThreatIntelligence"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/threatintelligence/indicators/bulkDelete/action"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Bulk Delete Threat Intelligence Indicators"),
		// 				Operation: to.Ptr("Bulk Delete Threat Intelligence Indicators"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("ThreatIntelligence"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/threatintelligence/indicators/bulkTag/action"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Bulk Tags Threat Intelligence Indicators"),
		// 				Operation: to.Ptr("Bulk Tags Threat Intelligence Indicators"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("ThreatIntelligence"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/threatintelligence/indicators/read"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Gets Threat Intelligence Indicators"),
		// 				Operation: to.Ptr("Get Threat Intelligence Indicators"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("ThreatIntelligence"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/threatintelligence/metrics/read"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Collect Threat Intelligence Metrics"),
		// 				Operation: to.Ptr("Collect Threat Intelligence Metrics"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("ThreatIntelligence"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/threatintelligence/createIndicator/action"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Create Threat Intelligence Indicator"),
		// 				Operation: to.Ptr("Create Threat Intelligence Indicator"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("ThreatIntelligence"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/threatintelligence/indicators/appendTags/action"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Append tags to Threat Intelligence Indicator"),
		// 				Operation: to.Ptr("Append tags to Threat Intelligence Indicator"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("ThreatIntelligence"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/threatintelligence/indicators/replaceTags/action"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Replace Tags of Threat Intelligence Indicator"),
		// 				Operation: to.Ptr("Replace Tags of Threat Intelligence Indicator"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("ThreatIntelligence"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/threatintelligence/queryIndicators/action"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Query Threat Intelligence Indicators"),
		// 				Operation: to.Ptr("Query Threat Intelligence Indicators"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("ThreatIntelligence"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/Watchlists/read"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Gets Watchlists"),
		// 				Operation: to.Ptr("Get Watchlists"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Watchlists"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/Watchlists/write"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Create Watchlists"),
		// 				Operation: to.Ptr("Create Watchlists"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Watchlists"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/Watchlists/delete"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Deletes Watchlists"),
		// 				Operation: to.Ptr("Delete Watchlists"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Watchlists"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/onboardingStates/read"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Gets an onboarding state"),
		// 				Operation: to.Ptr("Get Onboarding States"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Onboarding States"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/onboardingStates/write"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Updates an onboarding state"),
		// 				Operation: to.Ptr("Update Onboarding States"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Onboarding States"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.SecurityInsights/onboardingStates/delete"),
		// 			Display: &armsecurityinsights.OperationDisplay{
		// 				Description: to.Ptr("Deletes an onboarding state"),
		// 				Operation: to.Ptr("Delete Onboarding States"),
		// 				Provider: to.Ptr("Microsoft Security Insights"),
		// 				Resource: to.Ptr("Onboarding States"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 	}},
		// }
	}
}
