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

import original "github.com/blacknut/azure-sdk-for-go/services/datalake/analytics/mgmt/2016-11-01/account"

type AccountsClient = original.AccountsClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type ComputePoliciesClient = original.ComputePoliciesClient
type DataLakeStoreAccountsClient = original.DataLakeStoreAccountsClient
type FirewallRulesClient = original.FirewallRulesClient
type LocationsClient = original.LocationsClient
type AADObjectType = original.AADObjectType

const (
	Group            AADObjectType = original.Group
	ServicePrincipal AADObjectType = original.ServicePrincipal
	User             AADObjectType = original.User
)

type DataLakeAnalyticsAccountState = original.DataLakeAnalyticsAccountState

const (
	Active    DataLakeAnalyticsAccountState = original.Active
	Suspended DataLakeAnalyticsAccountState = original.Suspended
)

type DataLakeAnalyticsAccountStatus = original.DataLakeAnalyticsAccountStatus

const (
	Canceled   DataLakeAnalyticsAccountStatus = original.Canceled
	Creating   DataLakeAnalyticsAccountStatus = original.Creating
	Deleted    DataLakeAnalyticsAccountStatus = original.Deleted
	Deleting   DataLakeAnalyticsAccountStatus = original.Deleting
	Failed     DataLakeAnalyticsAccountStatus = original.Failed
	Patching   DataLakeAnalyticsAccountStatus = original.Patching
	Resuming   DataLakeAnalyticsAccountStatus = original.Resuming
	Running    DataLakeAnalyticsAccountStatus = original.Running
	Succeeded  DataLakeAnalyticsAccountStatus = original.Succeeded
	Suspending DataLakeAnalyticsAccountStatus = original.Suspending
	Undeleting DataLakeAnalyticsAccountStatus = original.Undeleting
)

type FirewallAllowAzureIpsState = original.FirewallAllowAzureIpsState

const (
	Disabled FirewallAllowAzureIpsState = original.Disabled
	Enabled  FirewallAllowAzureIpsState = original.Enabled
)

type FirewallState = original.FirewallState

const (
	FirewallStateDisabled FirewallState = original.FirewallStateDisabled
	FirewallStateEnabled  FirewallState = original.FirewallStateEnabled
)

type OperationOrigin = original.OperationOrigin

const (
	OperationOriginSystem     OperationOrigin = original.OperationOriginSystem
	OperationOriginUser       OperationOrigin = original.OperationOriginUser
	OperationOriginUsersystem OperationOrigin = original.OperationOriginUsersystem
)

type SubscriptionState = original.SubscriptionState

const (
	SubscriptionStateDeleted      SubscriptionState = original.SubscriptionStateDeleted
	SubscriptionStateRegistered   SubscriptionState = original.SubscriptionStateRegistered
	SubscriptionStateSuspended    SubscriptionState = original.SubscriptionStateSuspended
	SubscriptionStateUnregistered SubscriptionState = original.SubscriptionStateUnregistered
	SubscriptionStateWarned       SubscriptionState = original.SubscriptionStateWarned
)

type TierType = original.TierType

const (
	Commitment100000AUHours TierType = original.Commitment100000AUHours
	Commitment10000AUHours  TierType = original.Commitment10000AUHours
	Commitment1000AUHours   TierType = original.Commitment1000AUHours
	Commitment100AUHours    TierType = original.Commitment100AUHours
	Commitment500000AUHours TierType = original.Commitment500000AUHours
	Commitment50000AUHours  TierType = original.Commitment50000AUHours
	Commitment5000AUHours   TierType = original.Commitment5000AUHours
	Commitment500AUHours    TierType = original.Commitment500AUHours
	Consumption             TierType = original.Consumption
)

type AccountsCreateFutureType = original.AccountsCreateFutureType
type AccountsDeleteFutureType = original.AccountsDeleteFutureType
type AccountsUpdateFutureType = original.AccountsUpdateFutureType
type AddDataLakeStoreParameters = original.AddDataLakeStoreParameters
type AddDataLakeStoreProperties = original.AddDataLakeStoreProperties
type AddDataLakeStoreWithAccountParameters = original.AddDataLakeStoreWithAccountParameters
type AddStorageAccountParameters = original.AddStorageAccountParameters
type AddStorageAccountProperties = original.AddStorageAccountProperties
type AddStorageAccountWithAccountParameters = original.AddStorageAccountWithAccountParameters
type CapabilityInformation = original.CapabilityInformation
type CheckNameAvailabilityParameters = original.CheckNameAvailabilityParameters
type ComputePolicy = original.ComputePolicy
type ComputePolicyListResult = original.ComputePolicyListResult
type ComputePolicyListResultIterator = original.ComputePolicyListResultIterator
type ComputePolicyListResultPage = original.ComputePolicyListResultPage
type ComputePolicyProperties = original.ComputePolicyProperties
type CreateComputePolicyWithAccountParameters = original.CreateComputePolicyWithAccountParameters
type CreateDataLakeAnalyticsAccountParameters = original.CreateDataLakeAnalyticsAccountParameters
type CreateDataLakeAnalyticsAccountProperties = original.CreateDataLakeAnalyticsAccountProperties
type CreateFirewallRuleWithAccountParameters = original.CreateFirewallRuleWithAccountParameters
type CreateOrUpdateComputePolicyParameters = original.CreateOrUpdateComputePolicyParameters
type CreateOrUpdateComputePolicyProperties = original.CreateOrUpdateComputePolicyProperties
type CreateOrUpdateFirewallRuleParameters = original.CreateOrUpdateFirewallRuleParameters
type CreateOrUpdateFirewallRuleProperties = original.CreateOrUpdateFirewallRuleProperties
type DataLakeAnalyticsAccount = original.DataLakeAnalyticsAccount
type DataLakeAnalyticsAccountBasic = original.DataLakeAnalyticsAccountBasic
type DataLakeAnalyticsAccountListResult = original.DataLakeAnalyticsAccountListResult
type DataLakeAnalyticsAccountListResultIterator = original.DataLakeAnalyticsAccountListResultIterator
type DataLakeAnalyticsAccountListResultPage = original.DataLakeAnalyticsAccountListResultPage
type DataLakeAnalyticsAccountProperties = original.DataLakeAnalyticsAccountProperties
type DataLakeAnalyticsAccountPropertiesBasic = original.DataLakeAnalyticsAccountPropertiesBasic
type DataLakeStoreAccountInformation = original.DataLakeStoreAccountInformation
type DataLakeStoreAccountInformationListResult = original.DataLakeStoreAccountInformationListResult
type DataLakeStoreAccountInformationListResultIterator = original.DataLakeStoreAccountInformationListResultIterator
type DataLakeStoreAccountInformationListResultPage = original.DataLakeStoreAccountInformationListResultPage
type DataLakeStoreAccountInformationProperties = original.DataLakeStoreAccountInformationProperties
type FirewallRule = original.FirewallRule
type FirewallRuleListResult = original.FirewallRuleListResult
type FirewallRuleListResultIterator = original.FirewallRuleListResultIterator
type FirewallRuleListResultPage = original.FirewallRuleListResultPage
type FirewallRuleProperties = original.FirewallRuleProperties
type NameAvailabilityInformation = original.NameAvailabilityInformation
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type Resource = original.Resource
type SasTokenInformation = original.SasTokenInformation
type SasTokenInformationListResult = original.SasTokenInformationListResult
type SasTokenInformationListResultIterator = original.SasTokenInformationListResultIterator
type SasTokenInformationListResultPage = original.SasTokenInformationListResultPage
type StorageAccountInformation = original.StorageAccountInformation
type StorageAccountInformationListResult = original.StorageAccountInformationListResult
type StorageAccountInformationListResultIterator = original.StorageAccountInformationListResultIterator
type StorageAccountInformationListResultPage = original.StorageAccountInformationListResultPage
type StorageAccountInformationProperties = original.StorageAccountInformationProperties
type StorageContainer = original.StorageContainer
type StorageContainerListResult = original.StorageContainerListResult
type StorageContainerListResultIterator = original.StorageContainerListResultIterator
type StorageContainerListResultPage = original.StorageContainerListResultPage
type StorageContainerProperties = original.StorageContainerProperties
type SubResource = original.SubResource
type UpdateComputePolicyParameters = original.UpdateComputePolicyParameters
type UpdateComputePolicyProperties = original.UpdateComputePolicyProperties
type UpdateComputePolicyWithAccountParameters = original.UpdateComputePolicyWithAccountParameters
type UpdateDataLakeAnalyticsAccountParameters = original.UpdateDataLakeAnalyticsAccountParameters
type UpdateDataLakeAnalyticsAccountProperties = original.UpdateDataLakeAnalyticsAccountProperties
type UpdateDataLakeStoreProperties = original.UpdateDataLakeStoreProperties
type UpdateDataLakeStoreWithAccountParameters = original.UpdateDataLakeStoreWithAccountParameters
type UpdateFirewallRuleParameters = original.UpdateFirewallRuleParameters
type UpdateFirewallRuleProperties = original.UpdateFirewallRuleProperties
type UpdateFirewallRuleWithAccountParameters = original.UpdateFirewallRuleWithAccountParameters
type UpdateStorageAccountParameters = original.UpdateStorageAccountParameters
type UpdateStorageAccountProperties = original.UpdateStorageAccountProperties
type UpdateStorageAccountWithAccountParameters = original.UpdateStorageAccountWithAccountParameters
type OperationsClient = original.OperationsClient
type StorageAccountsClient = original.StorageAccountsClient

func NewAccountsClient(subscriptionID string) AccountsClient {
	return original.NewAccountsClient(subscriptionID)
}
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return original.NewAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewComputePoliciesClient(subscriptionID string) ComputePoliciesClient {
	return original.NewComputePoliciesClient(subscriptionID)
}
func NewComputePoliciesClientWithBaseURI(baseURI string, subscriptionID string) ComputePoliciesClient {
	return original.NewComputePoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewDataLakeStoreAccountsClient(subscriptionID string) DataLakeStoreAccountsClient {
	return original.NewDataLakeStoreAccountsClient(subscriptionID)
}
func NewDataLakeStoreAccountsClientWithBaseURI(baseURI string, subscriptionID string) DataLakeStoreAccountsClient {
	return original.NewDataLakeStoreAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewFirewallRulesClient(subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClient(subscriptionID)
}
func NewFirewallRulesClientWithBaseURI(baseURI string, subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewLocationsClient(subscriptionID string) LocationsClient {
	return original.NewLocationsClient(subscriptionID)
}
func NewLocationsClientWithBaseURI(baseURI string, subscriptionID string) LocationsClient {
	return original.NewLocationsClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleAADObjectTypeValues() []AADObjectType {
	return original.PossibleAADObjectTypeValues()
}
func PossibleDataLakeAnalyticsAccountStateValues() []DataLakeAnalyticsAccountState {
	return original.PossibleDataLakeAnalyticsAccountStateValues()
}
func PossibleDataLakeAnalyticsAccountStatusValues() []DataLakeAnalyticsAccountStatus {
	return original.PossibleDataLakeAnalyticsAccountStatusValues()
}
func PossibleFirewallAllowAzureIpsStateValues() []FirewallAllowAzureIpsState {
	return original.PossibleFirewallAllowAzureIpsStateValues()
}
func PossibleFirewallStateValues() []FirewallState {
	return original.PossibleFirewallStateValues()
}
func PossibleOperationOriginValues() []OperationOrigin {
	return original.PossibleOperationOriginValues()
}
func PossibleSubscriptionStateValues() []SubscriptionState {
	return original.PossibleSubscriptionStateValues()
}
func PossibleTierTypeValues() []TierType {
	return original.PossibleTierTypeValues()
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewStorageAccountsClient(subscriptionID string) StorageAccountsClient {
	return original.NewStorageAccountsClient(subscriptionID)
}
func NewStorageAccountsClientWithBaseURI(baseURI string, subscriptionID string) StorageAccountsClient {
	return original.NewStorageAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
