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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRuleTemplates/GetAlertRuleTemplates.json
func ExampleAlertRuleTemplatesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAlertRuleTemplatesClient().NewListPager("myRg", "myWorkspace", nil)
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
		// page.AlertRuleTemplatesList = armsecurityinsights.AlertRuleTemplatesList{
		// 	Value: []armsecurityinsights.AlertRuleTemplateClassification{
		// 		&armsecurityinsights.ScheduledAlertRuleTemplate{
		// 			Name: to.Ptr("65360bb0-8986-4ade-a89d-af3cf44d28aa"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/AlertRuleTemplates"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/AlertRuleTemplates/65360bb0-8986-4ade-a89d-af3cf44d28aa"),
		// 			Kind: to.Ptr(armsecurityinsights.AlertRuleKindScheduled),
		// 			Properties: &armsecurityinsights.ScheduledAlertRuleTemplateProperties{
		// 				Description: to.Ptr("This alert monitors changes to Amazon VPC (Virtual Private Cloud) settings such as new ACL entries and routes in route tables.\nMore information: https://medium.com/@GorillaStack/the-most-important-aws-cloudtrail-security-events-to-track-a5b9873f8255 \nand https://aws.amazon.com/vpc/"),
		// 				AlertRulesCreatedByTemplateCount: to.Ptr[int32](0),
		// 				CreatedDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-02-27T00:00:00Z"); return t}()),
		// 				DisplayName: to.Ptr("Changes to Amazon VPC settings"),
		// 				LastUpdatedDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-27T00:00:00Z"); return t}()),
		// 				Query: to.Ptr("let timeframe = 1d;\nAWSCloudTrail\n| where TimeGenerated >= ago(timeframe)\n| where EventName == \"CreateNetworkAclEntry\"\n    or EventName == \"CreateRoute\"\n| project TimeGenerated, EventName, EventTypeName, UserIdentityAccountId, UserIdentityPrincipalid, UserAgent, UserIdentityUserName, SessionMfaAuthenticated, SourceIpAddress, AWSRegion, EventSource, AdditionalEventData, ResponseElements\n| extend AccountCustomEntity = UserIdentityUserName, IPCustomEntity = SourceIpAddress"),
		// 				QueryFrequency: to.Ptr("P1D"),
		// 				QueryPeriod: to.Ptr("P1D"),
		// 				RequiredDataConnectors: []*armsecurityinsights.AlertRuleTemplateDataSource{
		// 					{
		// 						ConnectorID: to.Ptr("AWS"),
		// 						DataTypes: []*string{
		// 							to.Ptr("AWSCloudTrail")},
		// 					}},
		// 					Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
		// 					Status: to.Ptr(armsecurityinsights.TemplateStatusAvailable),
		// 					Tactics: []*armsecurityinsights.AttackTactic{
		// 						to.Ptr(armsecurityinsights.AttackTacticPrivilegeEscalation),
		// 						to.Ptr(armsecurityinsights.AttackTacticLateralMovement)},
		// 						Techniques: []*string{
		// 							to.Ptr("T1037"),
		// 							to.Ptr("T1021")},
		// 							TriggerOperator: to.Ptr(armsecurityinsights.TriggerOperatorGreaterThan),
		// 							TriggerThreshold: to.Ptr[int32](0),
		// 							Version: to.Ptr("1.0.1"),
		// 						},
		// 					},
		// 					&armsecurityinsights.FusionAlertRuleTemplate{
		// 						Name: to.Ptr("f71aba3d-28fb-450b-b192-4e76a83015c8"),
		// 						Type: to.Ptr("Microsoft.SecurityInsights/AlertRuleTemplates"),
		// 						ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/AlertRuleTemplates/f71aba3d-28fb-450b-b192-4e76a83015c8"),
		// 						Kind: to.Ptr(armsecurityinsights.AlertRuleKindFusion),
		// 						Properties: &armsecurityinsights.FusionAlertRuleTemplateProperties{
		// 							Description: to.Ptr("Microsoft Sentinel uses Fusion, a correlation engine based on scalable machine learning algorithms, to automatically detect multistage attacks by identifying combinations of anomalous behaviors and suspicious activities that are observed at various stages of the kill chain. On the basis of these discoveries, Azure Sentinel generates incidents that would otherwise be very difficult to catch. By design, these incidents are low-volume, high-fidelity, and high-severity, which is why this detection is turned ON by default.\n\nSince Fusion correlates multiple signals from various products to detect advanced multistage attacks, successful Fusion detections are presented as Fusion incidents on the Microsoft Sentinel Incidents page. This rule covers the following detections:\n- Fusion for emerging threats\n- Fusion for ransomware\n- Scenario-based Fusion detections (122 scenarios)\n\nTo enable these detections, we recommend you configure the following data connectors for best results:\n- Out-of-the-box anomaly detections\n- Azure Active Directory Identity Protection\n- Azure Defender\n- Azure Defender for IoT\n- Microsoft 365 Defender\n- Microsoft Cloud App Security    \n- Microsoft Defender for Endpoint\n- Microsoft Defender for Identity\n- Microsoft Defender for Office 365\n- Palo Alto Networks\n- Scheduled analytics rules, both built-in and those created by your security analysts. Analytics rules must contain kill-chain (tactics) and entity mapping information in order to be used by Fusion.\n\nFor the full description of each detection that is supported by Fusion, go to https://aka.ms/SentinelFusion."),
		// 							AlertRulesCreatedByTemplateCount: to.Ptr[int32](0),
		// 							CreatedDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-25T00:00:00Z"); return t}()),
		// 							DisplayName: to.Ptr("Advanced Multi-Stage Attack Detection"),
		// 							LastUpdatedDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-09T00:00:00Z"); return t}()),
		// 							Severity: to.Ptr(armsecurityinsights.AlertSeverityHigh),
		// 							SourceSettings: []*armsecurityinsights.FusionTemplateSourceSetting{
		// 								{
		// 									SourceName: to.Ptr("Anomalies"),
		// 								},
		// 								{
		// 									SourceName: to.Ptr("Alert providers"),
		// 									SourceSubTypes: []*armsecurityinsights.FusionTemplateSourceSubType{
		// 										{
		// 											SeverityFilter: &armsecurityinsights.FusionTemplateSubTypeSeverityFilter{
		// 												IsSupported: to.Ptr(true),
		// 												SeverityFilters: []*armsecurityinsights.AlertSeverity{
		// 													to.Ptr(armsecurityinsights.AlertSeverityInformational),
		// 													to.Ptr(armsecurityinsights.AlertSeverityLow),
		// 													to.Ptr(armsecurityinsights.AlertSeverityMedium),
		// 													to.Ptr(armsecurityinsights.AlertSeverityHigh)},
		// 												},
		// 												SourceSubTypeName: to.Ptr("Azure Active Directory Identity Protection"),
		// 											},
		// 											{
		// 												SeverityFilter: &armsecurityinsights.FusionTemplateSubTypeSeverityFilter{
		// 													IsSupported: to.Ptr(true),
		// 													SeverityFilters: []*armsecurityinsights.AlertSeverity{
		// 														to.Ptr(armsecurityinsights.AlertSeverityInformational),
		// 														to.Ptr(armsecurityinsights.AlertSeverityLow),
		// 														to.Ptr(armsecurityinsights.AlertSeverityMedium),
		// 														to.Ptr(armsecurityinsights.AlertSeverityHigh)},
		// 													},
		// 													SourceSubTypeName: to.Ptr("Azure Defender"),
		// 												},
		// 												{
		// 													SeverityFilter: &armsecurityinsights.FusionTemplateSubTypeSeverityFilter{
		// 														IsSupported: to.Ptr(true),
		// 														SeverityFilters: []*armsecurityinsights.AlertSeverity{
		// 															to.Ptr(armsecurityinsights.AlertSeverityInformational),
		// 															to.Ptr(armsecurityinsights.AlertSeverityLow),
		// 															to.Ptr(armsecurityinsights.AlertSeverityMedium),
		// 															to.Ptr(armsecurityinsights.AlertSeverityHigh)},
		// 														},
		// 														SourceSubTypeName: to.Ptr("Azure Defender for IoT"),
		// 													},
		// 													{
		// 														SeverityFilter: &armsecurityinsights.FusionTemplateSubTypeSeverityFilter{
		// 															IsSupported: to.Ptr(true),
		// 															SeverityFilters: []*armsecurityinsights.AlertSeverity{
		// 																to.Ptr(armsecurityinsights.AlertSeverityInformational),
		// 																to.Ptr(armsecurityinsights.AlertSeverityLow),
		// 																to.Ptr(armsecurityinsights.AlertSeverityMedium),
		// 																to.Ptr(armsecurityinsights.AlertSeverityHigh)},
		// 															},
		// 															SourceSubTypeName: to.Ptr("Microsoft 365 Defender"),
		// 														},
		// 														{
		// 															SeverityFilter: &armsecurityinsights.FusionTemplateSubTypeSeverityFilter{
		// 																IsSupported: to.Ptr(true),
		// 																SeverityFilters: []*armsecurityinsights.AlertSeverity{
		// 																	to.Ptr(armsecurityinsights.AlertSeverityInformational),
		// 																	to.Ptr(armsecurityinsights.AlertSeverityLow),
		// 																	to.Ptr(armsecurityinsights.AlertSeverityMedium),
		// 																	to.Ptr(armsecurityinsights.AlertSeverityHigh)},
		// 																},
		// 																SourceSubTypeName: to.Ptr("Microsoft Cloud App Security"),
		// 															},
		// 															{
		// 																SeverityFilter: &armsecurityinsights.FusionTemplateSubTypeSeverityFilter{
		// 																	IsSupported: to.Ptr(true),
		// 																	SeverityFilters: []*armsecurityinsights.AlertSeverity{
		// 																		to.Ptr(armsecurityinsights.AlertSeverityInformational),
		// 																		to.Ptr(armsecurityinsights.AlertSeverityLow),
		// 																		to.Ptr(armsecurityinsights.AlertSeverityMedium),
		// 																		to.Ptr(armsecurityinsights.AlertSeverityHigh)},
		// 																	},
		// 																	SourceSubTypeName: to.Ptr("Microsoft Defender for Endpoint"),
		// 																},
		// 																{
		// 																	SeverityFilter: &armsecurityinsights.FusionTemplateSubTypeSeverityFilter{
		// 																		IsSupported: to.Ptr(true),
		// 																		SeverityFilters: []*armsecurityinsights.AlertSeverity{
		// 																			to.Ptr(armsecurityinsights.AlertSeverityInformational),
		// 																			to.Ptr(armsecurityinsights.AlertSeverityLow),
		// 																			to.Ptr(armsecurityinsights.AlertSeverityMedium),
		// 																			to.Ptr(armsecurityinsights.AlertSeverityHigh)},
		// 																		},
		// 																		SourceSubTypeName: to.Ptr("Microsoft Defender for Identity"),
		// 																	},
		// 																	{
		// 																		SeverityFilter: &armsecurityinsights.FusionTemplateSubTypeSeverityFilter{
		// 																			IsSupported: to.Ptr(true),
		// 																			SeverityFilters: []*armsecurityinsights.AlertSeverity{
		// 																				to.Ptr(armsecurityinsights.AlertSeverityInformational),
		// 																				to.Ptr(armsecurityinsights.AlertSeverityLow),
		// 																				to.Ptr(armsecurityinsights.AlertSeverityMedium),
		// 																				to.Ptr(armsecurityinsights.AlertSeverityHigh)},
		// 																			},
		// 																			SourceSubTypeName: to.Ptr("Microsoft Defender for Office 365"),
		// 																		},
		// 																		{
		// 																			SeverityFilter: &armsecurityinsights.FusionTemplateSubTypeSeverityFilter{
		// 																				IsSupported: to.Ptr(true),
		// 																				SeverityFilters: []*armsecurityinsights.AlertSeverity{
		// 																					to.Ptr(armsecurityinsights.AlertSeverityInformational),
		// 																					to.Ptr(armsecurityinsights.AlertSeverityLow),
		// 																					to.Ptr(armsecurityinsights.AlertSeverityMedium),
		// 																					to.Ptr(armsecurityinsights.AlertSeverityHigh)},
		// 																				},
		// 																				SourceSubTypeName: to.Ptr("Azure Sentinel scheduled analytics rules"),
		// 																		}},
		// 																	},
		// 																	{
		// 																		SourceName: to.Ptr("Raw logs from other sources"),
		// 																		SourceSubTypes: []*armsecurityinsights.FusionTemplateSourceSubType{
		// 																			{
		// 																				SeverityFilter: &armsecurityinsights.FusionTemplateSubTypeSeverityFilter{
		// 																					IsSupported: to.Ptr(false),
		// 																				},
		// 																				SourceSubTypeName: to.Ptr("Palo Alto Networks"),
		// 																		}},
		// 																}},
		// 																Status: to.Ptr(armsecurityinsights.TemplateStatusAvailable),
		// 																Tactics: []*armsecurityinsights.AttackTactic{
		// 																	to.Ptr(armsecurityinsights.AttackTacticCollection),
		// 																	to.Ptr(armsecurityinsights.AttackTacticCommandAndControl),
		// 																	to.Ptr(armsecurityinsights.AttackTacticCredentialAccess),
		// 																	to.Ptr(armsecurityinsights.AttackTacticDefenseEvasion),
		// 																	to.Ptr(armsecurityinsights.AttackTacticDiscovery),
		// 																	to.Ptr(armsecurityinsights.AttackTacticExecution),
		// 																	to.Ptr(armsecurityinsights.AttackTacticExfiltration),
		// 																	to.Ptr(armsecurityinsights.AttackTacticImpact),
		// 																	to.Ptr(armsecurityinsights.AttackTacticInitialAccess),
		// 																	to.Ptr(armsecurityinsights.AttackTacticLateralMovement),
		// 																	to.Ptr(armsecurityinsights.AttackTacticPersistence),
		// 																	to.Ptr(armsecurityinsights.AttackTacticPrivilegeEscalation)},
		// 																},
		// 															},
		// 															&armsecurityinsights.MicrosoftSecurityIncidentCreationAlertRuleTemplate{
		// 																Name: to.Ptr("b3cfc7c0-092c-481c-a55b-34a3979758cb"),
		// 																Type: to.Ptr("Microsoft.SecurityInsights/AlertRuleTemplates"),
		// 																ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/AlertRuleTemplates/b3cfc7c0-092c-481c-a55b-34a3979758cb"),
		// 																Kind: to.Ptr(armsecurityinsights.AlertRuleKindMicrosoftSecurityIncidentCreation),
		// 																Properties: &armsecurityinsights.MicrosoftSecurityIncidentCreationAlertRuleTemplateProperties{
		// 																	Description: to.Ptr("Create incidents based on all alerts generated in Microsoft Cloud App Security"),
		// 																	AlertRulesCreatedByTemplateCount: to.Ptr[int32](0),
		// 																	CreatedDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-16T00:00:00Z"); return t}()),
		// 																	DisplayName: to.Ptr("Create incidents based on Microsoft Cloud App Security alerts"),
		// 																	LastUpdatedDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-27T00:00:00Z"); return t}()),
		// 																	Status: to.Ptr(armsecurityinsights.TemplateStatusAvailable),
		// 																	ProductFilter: to.Ptr(armsecurityinsights.MicrosoftSecurityProductNameMicrosoftCloudAppSecurity),
		// 																},
		// 														}},
		// 													}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRuleTemplates/GetAlertRuleTemplateById.json
func ExampleAlertRuleTemplatesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAlertRuleTemplatesClient().Get(ctx, "myRg", "myWorkspace", "65360bb0-8986-4ade-a89d-af3cf44d28aa", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.AlertRuleTemplatesClientGetResponse{
	// 	                            AlertRuleTemplateClassification: &armsecurityinsights.ScheduledAlertRuleTemplate{
	// 		Name: to.Ptr("65360bb0-8986-4ade-a89d-af3cf44d28aa"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/AlertRuleTemplates"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRuleTemplates/65360bb0-8986-4ade-a89d-af3cf44d28aa"),
	// 		Kind: to.Ptr(armsecurityinsights.AlertRuleKindScheduled),
	// 		Properties: &armsecurityinsights.ScheduledAlertRuleTemplateProperties{
	// 			Description: to.Ptr("This alert monitors changes to Amazon VPC (Virtual Private Cloud) settings such as new ACL entries and routes in route tables.\nMore information: https://medium.com/@GorillaStack/the-most-important-aws-cloudtrail-security-events-to-track-a5b9873f8255 \nand https://aws.amazon.com/vpc/"),
	// 			AlertRulesCreatedByTemplateCount: to.Ptr[int32](0),
	// 			CreatedDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-02-27T00:00:00Z"); return t}()),
	// 			DisplayName: to.Ptr("Changes to Amazon VPC settings"),
	// 			EventGroupingSettings: &armsecurityinsights.EventGroupingSettings{
	// 				AggregationKind: to.Ptr(armsecurityinsights.EventGroupingAggregationKindAlertPerResult),
	// 			},
	// 			LastUpdatedDateUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-27T00:00:00Z"); return t}()),
	// 			Query: to.Ptr("let timeframe = 1d;\nAWSCloudTrail\n| where TimeGenerated >= ago(timeframe)\n| where EventName == \"CreateNetworkAclEntry\"\n    or EventName == \"CreateRoute\"\n| project TimeGenerated, EventName, EventTypeName, UserIdentityAccountId, UserIdentityPrincipalid, UserAgent, UserIdentityUserName, SessionMfaAuthenticated, SourceIpAddress, AWSRegion, EventSource, AdditionalEventData, ResponseElements\n| extend AccountCustomEntity = UserIdentityUserName, IPCustomEntity = SourceIpAddress"),
	// 			QueryFrequency: to.Ptr("P1D"),
	// 			QueryPeriod: to.Ptr("P1D"),
	// 			RequiredDataConnectors: []*armsecurityinsights.AlertRuleTemplateDataSource{
	// 				{
	// 					ConnectorID: to.Ptr("AWS"),
	// 					DataTypes: []*string{
	// 						to.Ptr("AWSCloudTrail")},
	// 				}},
	// 				Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
	// 				Status: to.Ptr(armsecurityinsights.TemplateStatusAvailable),
	// 				Tactics: []*armsecurityinsights.AttackTactic{
	// 					to.Ptr(armsecurityinsights.AttackTacticPrivilegeEscalation),
	// 					to.Ptr(armsecurityinsights.AttackTacticLateralMovement)},
	// 					Techniques: []*string{
	// 						to.Ptr("T1037"),
	// 						to.Ptr("T1021")},
	// 						TriggerOperator: to.Ptr(armsecurityinsights.TriggerOperatorGreaterThan),
	// 						TriggerThreshold: to.Ptr[int32](0),
	// 						Version: to.Ptr("1.0.2"),
	// 					},
	// 				},
	// 				                        }
}
