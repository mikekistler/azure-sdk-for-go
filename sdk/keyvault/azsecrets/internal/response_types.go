//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package internal

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// HSMSecurityDomainDownloadPendingResponse contains the response from method HSMSecurityDomain.DownloadPending.
type HSMSecurityDomainDownloadPendingResponse struct {
	HSMSecurityDomainDownloadPendingResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HSMSecurityDomainDownloadPendingResult contains the result from method HSMSecurityDomain.DownloadPending.
type HSMSecurityDomainDownloadPendingResult struct {
	SecurityDomainOperationStatus
}

// HSMSecurityDomainDownloadPollerResponse contains the response from method HSMSecurityDomain.Download.
type HSMSecurityDomainDownloadPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *HSMSecurityDomainDownloadPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
func (l HSMSecurityDomainDownloadPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (HSMSecurityDomainDownloadResponse, error) {
	respType := HSMSecurityDomainDownloadResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.SecurityDomainObject)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a HSMSecurityDomainDownloadPollerResponse from the provided client and resume token.
func (l *HSMSecurityDomainDownloadPollerResponse) Resume(ctx context.Context, client *HSMSecurityDomainClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("HSMSecurityDomainClient.Download", token, client.con.Pipeline(), client.downloadHandleError)
	if err != nil {
		return err
	}
	poller := &HSMSecurityDomainDownloadPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// HSMSecurityDomainDownloadResponse contains the response from method HSMSecurityDomain.Download.
type HSMSecurityDomainDownloadResponse struct {
	HSMSecurityDomainDownloadResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HSMSecurityDomainDownloadResult contains the result from method HSMSecurityDomain.Download.
type HSMSecurityDomainDownloadResult struct {
	SecurityDomainObject
}

// HSMSecurityDomainTransferKeyResponse contains the response from method HSMSecurityDomain.TransferKey.
type HSMSecurityDomainTransferKeyResponse struct {
	HSMSecurityDomainTransferKeyResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HSMSecurityDomainTransferKeyResult contains the result from method HSMSecurityDomain.TransferKey.
type HSMSecurityDomainTransferKeyResult struct {
	TransferKey
}

// HSMSecurityDomainUploadPendingResponse contains the response from method HSMSecurityDomain.UploadPending.
type HSMSecurityDomainUploadPendingResponse struct {
	HSMSecurityDomainUploadPendingResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HSMSecurityDomainUploadPendingResult contains the result from method HSMSecurityDomain.UploadPending.
type HSMSecurityDomainUploadPendingResult struct {
	SecurityDomainOperationStatus
}

// HSMSecurityDomainUploadPollerResponse contains the response from method HSMSecurityDomain.Upload.
type HSMSecurityDomainUploadPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *HSMSecurityDomainUploadPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
func (l HSMSecurityDomainUploadPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (HSMSecurityDomainUploadResponse, error) {
	respType := HSMSecurityDomainUploadResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.SecurityDomainOperationStatus)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a HSMSecurityDomainUploadPollerResponse from the provided client and resume token.
func (l *HSMSecurityDomainUploadPollerResponse) Resume(ctx context.Context, client *HSMSecurityDomainClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("HSMSecurityDomainClient.Upload", token, client.con.Pipeline(), client.uploadHandleError)
	if err != nil {
		return err
	}
	poller := &HSMSecurityDomainUploadPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// HSMSecurityDomainUploadResponse contains the response from method HSMSecurityDomain.Upload.
type HSMSecurityDomainUploadResponse struct {
	HSMSecurityDomainUploadResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HSMSecurityDomainUploadResult contains the result from method HSMSecurityDomain.Upload.
type HSMSecurityDomainUploadResult struct {
	SecurityDomainOperationStatus
}

// KeyVaultClientBackupSecretResponse contains the response from method KeyVaultClient.BackupSecret.
type KeyVaultClientBackupSecretResponse struct {
	KeyVaultClientBackupSecretResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// KeyVaultClientBackupSecretResult contains the result from method KeyVaultClient.BackupSecret.
type KeyVaultClientBackupSecretResult struct {
	BackupSecretResult
}

// KeyVaultClientDeleteSecretResponse contains the response from method KeyVaultClient.DeleteSecret.
type KeyVaultClientDeleteSecretResponse struct {
	KeyVaultClientDeleteSecretResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// KeyVaultClientDeleteSecretResult contains the result from method KeyVaultClient.DeleteSecret.
type KeyVaultClientDeleteSecretResult struct {
	DeletedSecretBundle
}

// KeyVaultClientGetDeletedSecretResponse contains the response from method KeyVaultClient.GetDeletedSecret.
type KeyVaultClientGetDeletedSecretResponse struct {
	KeyVaultClientGetDeletedSecretResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// KeyVaultClientGetDeletedSecretResult contains the result from method KeyVaultClient.GetDeletedSecret.
type KeyVaultClientGetDeletedSecretResult struct {
	DeletedSecretBundle
}

// KeyVaultClientGetDeletedSecretsResponse contains the response from method KeyVaultClient.GetDeletedSecrets.
type KeyVaultClientGetDeletedSecretsResponse struct {
	KeyVaultClientGetDeletedSecretsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// KeyVaultClientGetDeletedSecretsResult contains the result from method KeyVaultClient.GetDeletedSecrets.
type KeyVaultClientGetDeletedSecretsResult struct {
	DeletedSecretListResult
}

// KeyVaultClientGetSecretResponse contains the response from method KeyVaultClient.GetSecret.
type KeyVaultClientGetSecretResponse struct {
	KeyVaultClientGetSecretResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// KeyVaultClientGetSecretResult contains the result from method KeyVaultClient.GetSecret.
type KeyVaultClientGetSecretResult struct {
	SecretBundle
}

// KeyVaultClientGetSecretVersionsResponse contains the response from method KeyVaultClient.GetSecretVersions.
type KeyVaultClientGetSecretVersionsResponse struct {
	KeyVaultClientGetSecretVersionsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// KeyVaultClientGetSecretVersionsResult contains the result from method KeyVaultClient.GetSecretVersions.
type KeyVaultClientGetSecretVersionsResult struct {
	SecretListResult
}

// KeyVaultClientGetSecretsResponse contains the response from method KeyVaultClient.GetSecrets.
type KeyVaultClientGetSecretsResponse struct {
	KeyVaultClientGetSecretsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// KeyVaultClientGetSecretsResult contains the result from method KeyVaultClient.GetSecrets.
type KeyVaultClientGetSecretsResult struct {
	SecretListResult
}

// KeyVaultClientPurgeDeletedSecretResponse contains the response from method KeyVaultClient.PurgeDeletedSecret.
type KeyVaultClientPurgeDeletedSecretResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// KeyVaultClientRecoverDeletedSecretResponse contains the response from method KeyVaultClient.RecoverDeletedSecret.
type KeyVaultClientRecoverDeletedSecretResponse struct {
	KeyVaultClientRecoverDeletedSecretResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// KeyVaultClientRecoverDeletedSecretResult contains the result from method KeyVaultClient.RecoverDeletedSecret.
type KeyVaultClientRecoverDeletedSecretResult struct {
	SecretBundle
}

// KeyVaultClientRestoreSecretResponse contains the response from method KeyVaultClient.RestoreSecret.
type KeyVaultClientRestoreSecretResponse struct {
	KeyVaultClientRestoreSecretResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// KeyVaultClientRestoreSecretResult contains the result from method KeyVaultClient.RestoreSecret.
type KeyVaultClientRestoreSecretResult struct {
	SecretBundle
}

// KeyVaultClientSetSecretResponse contains the response from method KeyVaultClient.SetSecret.
type KeyVaultClientSetSecretResponse struct {
	KeyVaultClientSetSecretResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// KeyVaultClientSetSecretResult contains the result from method KeyVaultClient.SetSecret.
type KeyVaultClientSetSecretResult struct {
	SecretBundle
}

// KeyVaultClientUpdateSecretResponse contains the response from method KeyVaultClient.UpdateSecret.
type KeyVaultClientUpdateSecretResponse struct {
	KeyVaultClientUpdateSecretResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// KeyVaultClientUpdateSecretResult contains the result from method KeyVaultClient.UpdateSecret.
type KeyVaultClientUpdateSecretResult struct {
	SecretBundle
}

// RoleAssignmentsCreateResponse contains the response from method RoleAssignments.Create.
type RoleAssignmentsCreateResponse struct {
	RoleAssignmentsCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RoleAssignmentsCreateResult contains the result from method RoleAssignments.Create.
type RoleAssignmentsCreateResult struct {
	RoleAssignment
}

// RoleAssignmentsDeleteResponse contains the response from method RoleAssignments.Delete.
type RoleAssignmentsDeleteResponse struct {
	RoleAssignmentsDeleteResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RoleAssignmentsDeleteResult contains the result from method RoleAssignments.Delete.
type RoleAssignmentsDeleteResult struct {
	RoleAssignment
}

// RoleAssignmentsGetResponse contains the response from method RoleAssignments.Get.
type RoleAssignmentsGetResponse struct {
	RoleAssignmentsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RoleAssignmentsGetResult contains the result from method RoleAssignments.Get.
type RoleAssignmentsGetResult struct {
	RoleAssignment
}

// RoleAssignmentsListForScopeResponse contains the response from method RoleAssignments.ListForScope.
type RoleAssignmentsListForScopeResponse struct {
	RoleAssignmentsListForScopeResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RoleAssignmentsListForScopeResult contains the result from method RoleAssignments.ListForScope.
type RoleAssignmentsListForScopeResult struct {
	RoleAssignmentListResult
}

// RoleDefinitionsCreateOrUpdateResponse contains the response from method RoleDefinitions.CreateOrUpdate.
type RoleDefinitionsCreateOrUpdateResponse struct {
	RoleDefinitionsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RoleDefinitionsCreateOrUpdateResult contains the result from method RoleDefinitions.CreateOrUpdate.
type RoleDefinitionsCreateOrUpdateResult struct {
	RoleDefinition
}

// RoleDefinitionsDeleteResponse contains the response from method RoleDefinitions.Delete.
type RoleDefinitionsDeleteResponse struct {
	RoleDefinitionsDeleteResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RoleDefinitionsDeleteResult contains the result from method RoleDefinitions.Delete.
type RoleDefinitionsDeleteResult struct {
	RoleDefinition
}

// RoleDefinitionsGetResponse contains the response from method RoleDefinitions.Get.
type RoleDefinitionsGetResponse struct {
	RoleDefinitionsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RoleDefinitionsGetResult contains the result from method RoleDefinitions.Get.
type RoleDefinitionsGetResult struct {
	RoleDefinition
}

// RoleDefinitionsListResponse contains the response from method RoleDefinitions.List.
type RoleDefinitionsListResponse struct {
	RoleDefinitionsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RoleDefinitionsListResult contains the result from method RoleDefinitions.List.
type RoleDefinitionsListResult struct {
	RoleDefinitionListResult
}
