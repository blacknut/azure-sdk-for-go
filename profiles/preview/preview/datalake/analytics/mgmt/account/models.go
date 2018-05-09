// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/blacknut/azure-sdk-for-go/tools/profileBuilder

package account

import original "github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/mgmt/2015-10-01-preview/account"

type Client = original.Client

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type DataLakeAnalyticsAccountState = original.DataLakeAnalyticsAccountState

const (
	Active    DataLakeAnalyticsAccountState = original.Active
	Suspended DataLakeAnalyticsAccountState = original.Suspended
)

type DataLakeAnalyticsAccountStatus = original.DataLakeAnalyticsAccountStatus

const (
	Creating   DataLakeAnalyticsAccountStatus = original.Creating
	Deleted    DataLakeAnalyticsAccountStatus = original.Deleted
	Deleting   DataLakeAnalyticsAccountStatus = original.Deleting
	Failed     DataLakeAnalyticsAccountStatus = original.Failed
	Patching   DataLakeAnalyticsAccountStatus = original.Patching
	Resuming   DataLakeAnalyticsAccountStatus = original.Resuming
	Running    DataLakeAnalyticsAccountStatus = original.Running
	Succeeded  DataLakeAnalyticsAccountStatus = original.Succeeded
	Suspending DataLakeAnalyticsAccountStatus = original.Suspending
)

type OperationStatus = original.OperationStatus

const (
	OperationStatusFailed     OperationStatus = original.OperationStatusFailed
	OperationStatusInProgress OperationStatus = original.OperationStatusInProgress
	OperationStatusSucceeded  OperationStatus = original.OperationStatusSucceeded
)

type AddDataLakeStoreParameters = original.AddDataLakeStoreParameters
type AddStorageAccountParameters = original.AddStorageAccountParameters
type AzureAsyncOperationResult = original.AzureAsyncOperationResult
type BlobContainer = original.BlobContainer
type BlobContainerProperties = original.BlobContainerProperties
type CreateFuture = original.CreateFuture
type DataLakeAnalyticsAccount = original.DataLakeAnalyticsAccount
type DataLakeAnalyticsAccountListDataLakeStoreResult = original.DataLakeAnalyticsAccountListDataLakeStoreResult
type DataLakeAnalyticsAccountListDataLakeStoreResultIterator = original.DataLakeAnalyticsAccountListDataLakeStoreResultIterator
type DataLakeAnalyticsAccountListDataLakeStoreResultPage = original.DataLakeAnalyticsAccountListDataLakeStoreResultPage
type DataLakeAnalyticsAccountListResult = original.DataLakeAnalyticsAccountListResult
type DataLakeAnalyticsAccountListResultIterator = original.DataLakeAnalyticsAccountListResultIterator
type DataLakeAnalyticsAccountListResultPage = original.DataLakeAnalyticsAccountListResultPage
type DataLakeAnalyticsAccountListStorageAccountsResult = original.DataLakeAnalyticsAccountListStorageAccountsResult
type DataLakeAnalyticsAccountListStorageAccountsResultIterator = original.DataLakeAnalyticsAccountListStorageAccountsResultIterator
type DataLakeAnalyticsAccountListStorageAccountsResultPage = original.DataLakeAnalyticsAccountListStorageAccountsResultPage
type DataLakeAnalyticsAccountProperties = original.DataLakeAnalyticsAccountProperties
type DataLakeStoreAccountInfo = original.DataLakeStoreAccountInfo
type DataLakeStoreAccountInfoProperties = original.DataLakeStoreAccountInfoProperties
type DeleteFuture = original.DeleteFuture
type Error = original.Error
type ErrorDetails = original.ErrorDetails
type InnerError = original.InnerError
type ListBlobContainersResult = original.ListBlobContainersResult
type ListBlobContainersResultIterator = original.ListBlobContainersResultIterator
type ListBlobContainersResultPage = original.ListBlobContainersResultPage
type ListSasTokensResult = original.ListSasTokensResult
type ListSasTokensResultIterator = original.ListSasTokensResultIterator
type ListSasTokensResultPage = original.ListSasTokensResultPage
type SasTokenInfo = original.SasTokenInfo
type StorageAccountInfo = original.StorageAccountInfo
type StorageAccountProperties = original.StorageAccountProperties
type UpdateFuture = original.UpdateFuture

func NewClient(subscriptionID string) Client {
	return original.NewClient(subscriptionID)
}
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return original.NewClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleDataLakeAnalyticsAccountStateValues() []DataLakeAnalyticsAccountState {
	return original.PossibleDataLakeAnalyticsAccountStateValues()
}
func PossibleDataLakeAnalyticsAccountStatusValues() []DataLakeAnalyticsAccountStatus {
	return original.PossibleDataLakeAnalyticsAccountStatusValues()
}
func PossibleOperationStatusValues() []OperationStatus {
	return original.PossibleOperationStatusValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
