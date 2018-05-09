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

package databox

import original "github.com/blacknut/azure-sdk-for-go/services/databox/mgmt/2018-01-01/databox"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type JobsClient = original.JobsClient
type AddressType = original.AddressType

const (
	Commercial  AddressType = original.Commercial
	None        AddressType = original.None
	Residential AddressType = original.Residential
)

type AddressValidationStatus = original.AddressValidationStatus

const (
	Ambiguous AddressValidationStatus = original.Ambiguous
	Invalid   AddressValidationStatus = original.Invalid
	Valid     AddressValidationStatus = original.Valid
)

type CopyLogDetailsType = original.CopyLogDetailsType

const (
	CopyLogDetailsTypeCopyLogDetails CopyLogDetailsType = original.CopyLogDetailsTypeCopyLogDetails
	CopyLogDetailsTypeDisk           CopyLogDetailsType = original.CopyLogDetailsTypeDisk
	CopyLogDetailsTypePod            CopyLogDetailsType = original.CopyLogDetailsTypePod
)

type CopyStatus = original.CopyStatus

const (
	Completed           CopyStatus = original.Completed
	CompletedWithErrors CopyStatus = original.CompletedWithErrors
	Failed              CopyStatus = original.Failed
	InProgress          CopyStatus = original.InProgress
	NotStarted          CopyStatus = original.NotStarted
)

type DeviceIssueType = original.DeviceIssueType

const (
	DeviceHealthCheckShowFailures DeviceIssueType = original.DeviceHealthCheckShowFailures
	DeviceNotBootingUp            DeviceIssueType = original.DeviceNotBootingUp
	DeviceTampering               DeviceIssueType = original.DeviceTampering
	Misc                          DeviceIssueType = original.Misc
	NICsAreNotWorking             DeviceIssueType = original.NICsAreNotWorking
)

type DeviceType = original.DeviceType

const (
	Cabinet DeviceType = original.Cabinet
	Disk    DeviceType = original.Disk
	Pod     DeviceType = original.Pod
)

type IssueType = original.IssueType

const (
	CredentialNotWorking     IssueType = original.CredentialNotWorking
	DeviceFailure            IssueType = original.DeviceFailure
	DeviceMismatch           IssueType = original.DeviceMismatch
	ValidationStringMismatch IssueType = original.ValidationStringMismatch
)

type JobDetailsType = original.JobDetailsType

const (
	JobDetailsTypeDisk       JobDetailsType = original.JobDetailsTypeDisk
	JobDetailsTypeJobDetails JobDetailsType = original.JobDetailsTypeJobDetails
	JobDetailsTypePod        JobDetailsType = original.JobDetailsTypePod
)

type JobSecretsType = original.JobSecretsType

const (
	JobSecretsTypeCabinet    JobSecretsType = original.JobSecretsTypeCabinet
	JobSecretsTypeDisk       JobSecretsType = original.JobSecretsTypeDisk
	JobSecretsTypeJobSecrets JobSecretsType = original.JobSecretsTypeJobSecrets
	JobSecretsTypePod        JobSecretsType = original.JobSecretsTypePod
)

type NotificationStageName = original.NotificationStageName

const (
	AtAzureDC      NotificationStageName = original.AtAzureDC
	DataCopy       NotificationStageName = original.DataCopy
	Delivered      NotificationStageName = original.Delivered
	DevicePrepared NotificationStageName = original.DevicePrepared
	Dispatched     NotificationStageName = original.Dispatched
	PickedUp       NotificationStageName = original.PickedUp
)

type StageName = original.StageName

const (
	StageNameAborted                       StageName = original.StageNameAborted
	StageNameAtAzureDC                     StageName = original.StageNameAtAzureDC
	StageNameCancelled                     StageName = original.StageNameCancelled
	StageNameCompleted                     StageName = original.StageNameCompleted
	StageNameCompletedWithErrors           StageName = original.StageNameCompletedWithErrors
	StageNameDataCopy                      StageName = original.StageNameDataCopy
	StageNameDelivered                     StageName = original.StageNameDelivered
	StageNameDeviceOrdered                 StageName = original.StageNameDeviceOrdered
	StageNameDevicePrepared                StageName = original.StageNameDevicePrepared
	StageNameDispatched                    StageName = original.StageNameDispatched
	StageNameFailedIssueDetectedAtAzureDC  StageName = original.StageNameFailedIssueDetectedAtAzureDC
	StageNameFailedIssueReportedAtCustomer StageName = original.StageNameFailedIssueReportedAtCustomer
	StageNamePickedUp                      StageName = original.StageNamePickedUp
)

type StageStatus = original.StageStatus

const (
	StageStatusCancelled           StageStatus = original.StageStatusCancelled
	StageStatusCancelling          StageStatus = original.StageStatusCancelling
	StageStatusFailed              StageStatus = original.StageStatusFailed
	StageStatusInProgress          StageStatus = original.StageStatusInProgress
	StageStatusNone                StageStatus = original.StageStatusNone
	StageStatusSucceeded           StageStatus = original.StageStatusSucceeded
	StageStatusSucceededWithErrors StageStatus = original.StageStatusSucceededWithErrors
)

type AccountCopyLogDetails = original.AccountCopyLogDetails
type AccountCredentialDetails = original.AccountCredentialDetails
type AddressValidationOutput = original.AddressValidationOutput
type AddressValidationProperties = original.AddressValidationProperties
type ArmBaseObject = original.ArmBaseObject
type AvailableSkuRequest = original.AvailableSkuRequest
type AvailableSkusResult = original.AvailableSkusResult
type AvailableSkusResultIterator = original.AvailableSkusResultIterator
type AvailableSkusResultPage = original.AvailableSkusResultPage
type CabinetJobSecrets = original.CabinetJobSecrets
type CabinetPodSecret = original.CabinetPodSecret
type CancellationReason = original.CancellationReason
type ContactDetails = original.ContactDetails
type BasicCopyLogDetails = original.BasicCopyLogDetails
type CopyLogDetails = original.CopyLogDetails
type CopyProgress = original.CopyProgress
type DestinationAccountDetails = original.DestinationAccountDetails
type DestinationToServiceLocationMap = original.DestinationToServiceLocationMap
type DiskCopyLogDetails = original.DiskCopyLogDetails
type DiskCopyProgress = original.DiskCopyProgress
type DiskJobDetails = original.DiskJobDetails
type DiskJobSecrets = original.DiskJobSecrets
type DiskSecret = original.DiskSecret
type Error = original.Error
type GetCopyLogsURIOutput = original.GetCopyLogsURIOutput
type BasicJobDetails = original.BasicJobDetails
type JobDetails = original.JobDetails
type JobErrorDetails = original.JobErrorDetails
type JobProperties = original.JobProperties
type JobResource = original.JobResource
type JobResourceList = original.JobResourceList
type JobResourceListIterator = original.JobResourceListIterator
type JobResourceListPage = original.JobResourceListPage
type JobResourceUpdateParameter = original.JobResourceUpdateParameter
type JobsCreateFuture = original.JobsCreateFuture
type BasicJobSecrets = original.BasicJobSecrets
type JobSecrets = original.JobSecrets
type JobStages = original.JobStages
type JobsUpdateFuture = original.JobsUpdateFuture
type NotificationPreference = original.NotificationPreference
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type PackageShippingDetails = original.PackageShippingDetails
type PodJobDetails = original.PodJobDetails
type PodJobSecrets = original.PodJobSecrets
type PodSecret = original.PodSecret
type ReportIssueDetails = original.ReportIssueDetails
type Resource = original.Resource
type ShareCredentialDetails = original.ShareCredentialDetails
type ShipmentPickUpRequest = original.ShipmentPickUpRequest
type ShipmentPickUpResponse = original.ShipmentPickUpResponse
type ShippingAddress = original.ShippingAddress
type ShippingLabelDetails = original.ShippingLabelDetails
type Sku = original.Sku
type SkuCapacity = original.SkuCapacity
type SkuCost = original.SkuCost
type SkuInformation = original.SkuInformation
type SkuProperties = original.SkuProperties
type UnencryptedSecrets = original.UnencryptedSecrets
type UpdateJobDetails = original.UpdateJobDetails
type UpdateJobProperties = original.UpdateJobProperties
type ValidateAddress = original.ValidateAddress
type OperationsClient = original.OperationsClient
type ServiceClient = original.ServiceClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewJobsClient(subscriptionID string) JobsClient {
	return original.NewJobsClient(subscriptionID)
}
func NewJobsClientWithBaseURI(baseURI string, subscriptionID string) JobsClient {
	return original.NewJobsClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleAddressTypeValues() []AddressType {
	return original.PossibleAddressTypeValues()
}
func PossibleAddressValidationStatusValues() []AddressValidationStatus {
	return original.PossibleAddressValidationStatusValues()
}
func PossibleCopyLogDetailsTypeValues() []CopyLogDetailsType {
	return original.PossibleCopyLogDetailsTypeValues()
}
func PossibleCopyStatusValues() []CopyStatus {
	return original.PossibleCopyStatusValues()
}
func PossibleDeviceIssueTypeValues() []DeviceIssueType {
	return original.PossibleDeviceIssueTypeValues()
}
func PossibleDeviceTypeValues() []DeviceType {
	return original.PossibleDeviceTypeValues()
}
func PossibleIssueTypeValues() []IssueType {
	return original.PossibleIssueTypeValues()
}
func PossibleJobDetailsTypeValues() []JobDetailsType {
	return original.PossibleJobDetailsTypeValues()
}
func PossibleJobSecretsTypeValues() []JobSecretsType {
	return original.PossibleJobSecretsTypeValues()
}
func PossibleNotificationStageNameValues() []NotificationStageName {
	return original.PossibleNotificationStageNameValues()
}
func PossibleStageNameValues() []StageName {
	return original.PossibleStageNameValues()
}
func PossibleStageStatusValues() []StageStatus {
	return original.PossibleStageStatusValues()
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewServiceClient(subscriptionID string) ServiceClient {
	return original.NewServiceClient(subscriptionID)
}
func NewServiceClientWithBaseURI(baseURI string, subscriptionID string) ServiceClient {
	return original.NewServiceClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
