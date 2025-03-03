//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armvirtualmachineimagebuilder

const (
	moduleName    = "armvirtualmachineimagebuilder"
	moduleVersion = "v1.2.0"
)

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// ProvisioningErrorCode - Error code of the provisioning failure
type ProvisioningErrorCode string

const (
	ProvisioningErrorCodeBadCustomizerType           ProvisioningErrorCode = "BadCustomizerType"
	ProvisioningErrorCodeBadDistributeType           ProvisioningErrorCode = "BadDistributeType"
	ProvisioningErrorCodeBadManagedImageSource       ProvisioningErrorCode = "BadManagedImageSource"
	ProvisioningErrorCodeBadPIRSource                ProvisioningErrorCode = "BadPIRSource"
	ProvisioningErrorCodeBadSharedImageDistribute    ProvisioningErrorCode = "BadSharedImageDistribute"
	ProvisioningErrorCodeBadSharedImageVersionSource ProvisioningErrorCode = "BadSharedImageVersionSource"
	ProvisioningErrorCodeBadSourceType               ProvisioningErrorCode = "BadSourceType"
	ProvisioningErrorCodeBadStagingResourceGroup     ProvisioningErrorCode = "BadStagingResourceGroup"
	ProvisioningErrorCodeBadValidatorType            ProvisioningErrorCode = "BadValidatorType"
	ProvisioningErrorCodeNoCustomizerScript          ProvisioningErrorCode = "NoCustomizerScript"
	ProvisioningErrorCodeNoValidatorScript           ProvisioningErrorCode = "NoValidatorScript"
	ProvisioningErrorCodeOther                       ProvisioningErrorCode = "Other"
	ProvisioningErrorCodeServerError                 ProvisioningErrorCode = "ServerError"
	ProvisioningErrorCodeUnsupportedCustomizerType   ProvisioningErrorCode = "UnsupportedCustomizerType"
	ProvisioningErrorCodeUnsupportedValidatorType    ProvisioningErrorCode = "UnsupportedValidatorType"
)

// PossibleProvisioningErrorCodeValues returns the possible values for the ProvisioningErrorCode const type.
func PossibleProvisioningErrorCodeValues() []ProvisioningErrorCode {
	return []ProvisioningErrorCode{
		ProvisioningErrorCodeBadCustomizerType,
		ProvisioningErrorCodeBadDistributeType,
		ProvisioningErrorCodeBadManagedImageSource,
		ProvisioningErrorCodeBadPIRSource,
		ProvisioningErrorCodeBadSharedImageDistribute,
		ProvisioningErrorCodeBadSharedImageVersionSource,
		ProvisioningErrorCodeBadSourceType,
		ProvisioningErrorCodeBadStagingResourceGroup,
		ProvisioningErrorCodeBadValidatorType,
		ProvisioningErrorCodeNoCustomizerScript,
		ProvisioningErrorCodeNoValidatorScript,
		ProvisioningErrorCodeOther,
		ProvisioningErrorCodeServerError,
		ProvisioningErrorCodeUnsupportedCustomizerType,
		ProvisioningErrorCodeUnsupportedValidatorType,
	}
}

// ProvisioningState - Provisioning state of the resource
type ProvisioningState string

const (
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCreating,
		ProvisioningStateUpdating,
		ProvisioningStateSucceeded,
		ProvisioningStateFailed,
		ProvisioningStateDeleting,
	}
}

// ResourceIdentityType - The type of identity used for the image template. The type 'None' will remove any identities from
// the image template.
type ResourceIdentityType string

const (
	ResourceIdentityTypeUserAssigned ResourceIdentityType = "UserAssigned"
	ResourceIdentityTypeNone         ResourceIdentityType = "None"
)

// PossibleResourceIdentityTypeValues returns the possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{
		ResourceIdentityTypeUserAssigned,
		ResourceIdentityTypeNone,
	}
}

// RunState - State of the last run
type RunState string

const (
	RunStateRunning            RunState = "Running"
	RunStateCanceling          RunState = "Canceling"
	RunStateSucceeded          RunState = "Succeeded"
	RunStatePartiallySucceeded RunState = "PartiallySucceeded"
	RunStateFailed             RunState = "Failed"
	RunStateCanceled           RunState = "Canceled"
)

// PossibleRunStateValues returns the possible values for the RunState const type.
func PossibleRunStateValues() []RunState {
	return []RunState{
		RunStateRunning,
		RunStateCanceling,
		RunStateSucceeded,
		RunStatePartiallySucceeded,
		RunStateFailed,
		RunStateCanceled,
	}
}

// RunSubState - Sub-state of the last run
type RunSubState string

const (
	RunSubStateQueued       RunSubState = "Queued"
	RunSubStateBuilding     RunSubState = "Building"
	RunSubStateCustomizing  RunSubState = "Customizing"
	RunSubStateValidating   RunSubState = "Validating"
	RunSubStateDistributing RunSubState = "Distributing"
)

// PossibleRunSubStateValues returns the possible values for the RunSubState const type.
func PossibleRunSubStateValues() []RunSubState {
	return []RunSubState{
		RunSubStateQueued,
		RunSubStateBuilding,
		RunSubStateCustomizing,
		RunSubStateValidating,
		RunSubStateDistributing,
	}
}

// SharedImageStorageAccountType - Storage account type to be used to store the shared image. Omit to use the default (Standard_LRS).
type SharedImageStorageAccountType string

const (
	SharedImageStorageAccountTypeStandardLRS SharedImageStorageAccountType = "Standard_LRS"
	SharedImageStorageAccountTypeStandardZRS SharedImageStorageAccountType = "Standard_ZRS"
)

// PossibleSharedImageStorageAccountTypeValues returns the possible values for the SharedImageStorageAccountType const type.
func PossibleSharedImageStorageAccountTypeValues() []SharedImageStorageAccountType {
	return []SharedImageStorageAccountType{
		SharedImageStorageAccountTypeStandardLRS,
		SharedImageStorageAccountTypeStandardZRS,
	}
}
