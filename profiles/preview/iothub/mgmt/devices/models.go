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

package devices

import original "github.com/blacknut/azure-sdk-for-go/services/iothub/mgmt/2018-04-01/devices"

type CertificatesClient = original.CertificatesClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type IotHubResourceClient = original.IotHubResourceClient
type AccessRights = original.AccessRights

const (
	DeviceConnect                                        AccessRights = original.DeviceConnect
	RegistryRead                                         AccessRights = original.RegistryRead
	RegistryReadDeviceConnect                            AccessRights = original.RegistryReadDeviceConnect
	RegistryReadRegistryWrite                            AccessRights = original.RegistryReadRegistryWrite
	RegistryReadRegistryWriteDeviceConnect               AccessRights = original.RegistryReadRegistryWriteDeviceConnect
	RegistryReadRegistryWriteServiceConnect              AccessRights = original.RegistryReadRegistryWriteServiceConnect
	RegistryReadRegistryWriteServiceConnectDeviceConnect AccessRights = original.RegistryReadRegistryWriteServiceConnectDeviceConnect
	RegistryReadServiceConnect                           AccessRights = original.RegistryReadServiceConnect
	RegistryReadServiceConnectDeviceConnect              AccessRights = original.RegistryReadServiceConnectDeviceConnect
	RegistryWrite                                        AccessRights = original.RegistryWrite
	RegistryWriteDeviceConnect                           AccessRights = original.RegistryWriteDeviceConnect
	RegistryWriteServiceConnect                          AccessRights = original.RegistryWriteServiceConnect
	RegistryWriteServiceConnectDeviceConnect             AccessRights = original.RegistryWriteServiceConnectDeviceConnect
	ServiceConnect                                       AccessRights = original.ServiceConnect
	ServiceConnectDeviceConnect                          AccessRights = original.ServiceConnectDeviceConnect
)

type Capabilities = original.Capabilities

const (
	DeviceManagement Capabilities = original.DeviceManagement
	None             Capabilities = original.None
)

type IotHubNameUnavailabilityReason = original.IotHubNameUnavailabilityReason

const (
	AlreadyExists IotHubNameUnavailabilityReason = original.AlreadyExists
	Invalid       IotHubNameUnavailabilityReason = original.Invalid
)

type IotHubScaleType = original.IotHubScaleType

const (
	IotHubScaleTypeAutomatic IotHubScaleType = original.IotHubScaleTypeAutomatic
	IotHubScaleTypeManual    IotHubScaleType = original.IotHubScaleTypeManual
	IotHubScaleTypeNone      IotHubScaleType = original.IotHubScaleTypeNone
)

type IotHubSku = original.IotHubSku

const (
	B1 IotHubSku = original.B1
	B2 IotHubSku = original.B2
	B3 IotHubSku = original.B3
	F1 IotHubSku = original.F1
	S1 IotHubSku = original.S1
	S2 IotHubSku = original.S2
	S3 IotHubSku = original.S3
)

type IotHubSkuTier = original.IotHubSkuTier

const (
	Basic    IotHubSkuTier = original.Basic
	Free     IotHubSkuTier = original.Free
	Standard IotHubSkuTier = original.Standard
)

type IPFilterActionType = original.IPFilterActionType

const (
	Accept IPFilterActionType = original.Accept
	Reject IPFilterActionType = original.Reject
)

type JobStatus = original.JobStatus

const (
	Cancelled JobStatus = original.Cancelled
	Completed JobStatus = original.Completed
	Enqueued  JobStatus = original.Enqueued
	Failed    JobStatus = original.Failed
	Running   JobStatus = original.Running
	Unknown   JobStatus = original.Unknown
)

type JobType = original.JobType

const (
	JobTypeBackup                    JobType = original.JobTypeBackup
	JobTypeExport                    JobType = original.JobTypeExport
	JobTypeFactoryResetDevice        JobType = original.JobTypeFactoryResetDevice
	JobTypeFirmwareUpdate            JobType = original.JobTypeFirmwareUpdate
	JobTypeImport                    JobType = original.JobTypeImport
	JobTypeReadDeviceProperties      JobType = original.JobTypeReadDeviceProperties
	JobTypeRebootDevice              JobType = original.JobTypeRebootDevice
	JobTypeUnknown                   JobType = original.JobTypeUnknown
	JobTypeUpdateDeviceConfiguration JobType = original.JobTypeUpdateDeviceConfiguration
	JobTypeWriteDeviceProperties     JobType = original.JobTypeWriteDeviceProperties
)

type OperationMonitoringLevel = original.OperationMonitoringLevel

const (
	OperationMonitoringLevelError            OperationMonitoringLevel = original.OperationMonitoringLevelError
	OperationMonitoringLevelErrorInformation OperationMonitoringLevel = original.OperationMonitoringLevelErrorInformation
	OperationMonitoringLevelInformation      OperationMonitoringLevel = original.OperationMonitoringLevelInformation
	OperationMonitoringLevelNone             OperationMonitoringLevel = original.OperationMonitoringLevelNone
)

type RoutingSource = original.RoutingSource

const (
	DeviceJobLifecycleEvents RoutingSource = original.DeviceJobLifecycleEvents
	DeviceLifecycleEvents    RoutingSource = original.DeviceLifecycleEvents
	DeviceMessages           RoutingSource = original.DeviceMessages
	TwinChangeEvents         RoutingSource = original.TwinChangeEvents
)

type CertificateBodyDescription = original.CertificateBodyDescription
type CertificateDescription = original.CertificateDescription
type CertificateListDescription = original.CertificateListDescription
type CertificateProperties = original.CertificateProperties
type CertificatePropertiesWithNonce = original.CertificatePropertiesWithNonce
type CertificateVerificationDescription = original.CertificateVerificationDescription
type CertificateWithNonceDescription = original.CertificateWithNonceDescription
type CloudToDeviceProperties = original.CloudToDeviceProperties
type ErrorDetails = original.ErrorDetails
type EventHubConsumerGroupInfo = original.EventHubConsumerGroupInfo
type EventHubConsumerGroupsListResult = original.EventHubConsumerGroupsListResult
type EventHubConsumerGroupsListResultIterator = original.EventHubConsumerGroupsListResultIterator
type EventHubConsumerGroupsListResultPage = original.EventHubConsumerGroupsListResultPage
type EventHubProperties = original.EventHubProperties
type ExportDevicesRequest = original.ExportDevicesRequest
type FallbackRouteProperties = original.FallbackRouteProperties
type FeedbackProperties = original.FeedbackProperties
type ImportDevicesRequest = original.ImportDevicesRequest
type IotHubCapacity = original.IotHubCapacity
type IotHubDescription = original.IotHubDescription
type IotHubDescriptionListResult = original.IotHubDescriptionListResult
type IotHubDescriptionListResultIterator = original.IotHubDescriptionListResultIterator
type IotHubDescriptionListResultPage = original.IotHubDescriptionListResultPage
type IotHubNameAvailabilityInfo = original.IotHubNameAvailabilityInfo
type IotHubProperties = original.IotHubProperties
type IotHubQuotaMetricInfo = original.IotHubQuotaMetricInfo
type IotHubQuotaMetricInfoListResult = original.IotHubQuotaMetricInfoListResult
type IotHubQuotaMetricInfoListResultIterator = original.IotHubQuotaMetricInfoListResultIterator
type IotHubQuotaMetricInfoListResultPage = original.IotHubQuotaMetricInfoListResultPage
type IotHubResourceCreateOrUpdateFuture = original.IotHubResourceCreateOrUpdateFuture
type IotHubResourceDeleteFuture = original.IotHubResourceDeleteFuture
type IotHubResourceUpdateFuture = original.IotHubResourceUpdateFuture
type IotHubSkuDescription = original.IotHubSkuDescription
type IotHubSkuDescriptionListResult = original.IotHubSkuDescriptionListResult
type IotHubSkuDescriptionListResultIterator = original.IotHubSkuDescriptionListResultIterator
type IotHubSkuDescriptionListResultPage = original.IotHubSkuDescriptionListResultPage
type IotHubSkuInfo = original.IotHubSkuInfo
type IPFilterRule = original.IPFilterRule
type JobResponse = original.JobResponse
type JobResponseListResult = original.JobResponseListResult
type JobResponseListResultIterator = original.JobResponseListResultIterator
type JobResponseListResultPage = original.JobResponseListResultPage
type MessagingEndpointProperties = original.MessagingEndpointProperties
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationInputs = original.OperationInputs
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsMonitoringProperties = original.OperationsMonitoringProperties
type RegistryStatistics = original.RegistryStatistics
type Resource = original.Resource
type RouteProperties = original.RouteProperties
type RoutingEndpoints = original.RoutingEndpoints
type RoutingEventHubProperties = original.RoutingEventHubProperties
type RoutingProperties = original.RoutingProperties
type RoutingServiceBusQueueEndpointProperties = original.RoutingServiceBusQueueEndpointProperties
type RoutingServiceBusTopicEndpointProperties = original.RoutingServiceBusTopicEndpointProperties
type RoutingStorageContainerProperties = original.RoutingStorageContainerProperties
type SetObject = original.SetObject
type SharedAccessSignatureAuthorizationRule = original.SharedAccessSignatureAuthorizationRule
type SharedAccessSignatureAuthorizationRuleListResult = original.SharedAccessSignatureAuthorizationRuleListResult
type SharedAccessSignatureAuthorizationRuleListResultIterator = original.SharedAccessSignatureAuthorizationRuleListResultIterator
type SharedAccessSignatureAuthorizationRuleListResultPage = original.SharedAccessSignatureAuthorizationRuleListResultPage
type StorageEndpointProperties = original.StorageEndpointProperties
type TagsResource = original.TagsResource
type OperationsClient = original.OperationsClient

func NewCertificatesClient(subscriptionID string) CertificatesClient {
	return original.NewCertificatesClient(subscriptionID)
}
func NewCertificatesClientWithBaseURI(baseURI string, subscriptionID string) CertificatesClient {
	return original.NewCertificatesClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewIotHubResourceClient(subscriptionID string) IotHubResourceClient {
	return original.NewIotHubResourceClient(subscriptionID)
}
func NewIotHubResourceClientWithBaseURI(baseURI string, subscriptionID string) IotHubResourceClient {
	return original.NewIotHubResourceClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccessRightsValues() []AccessRights {
	return original.PossibleAccessRightsValues()
}
func PossibleCapabilitiesValues() []Capabilities {
	return original.PossibleCapabilitiesValues()
}
func PossibleIotHubNameUnavailabilityReasonValues() []IotHubNameUnavailabilityReason {
	return original.PossibleIotHubNameUnavailabilityReasonValues()
}
func PossibleIotHubScaleTypeValues() []IotHubScaleType {
	return original.PossibleIotHubScaleTypeValues()
}
func PossibleIotHubSkuValues() []IotHubSku {
	return original.PossibleIotHubSkuValues()
}
func PossibleIotHubSkuTierValues() []IotHubSkuTier {
	return original.PossibleIotHubSkuTierValues()
}
func PossibleIPFilterActionTypeValues() []IPFilterActionType {
	return original.PossibleIPFilterActionTypeValues()
}
func PossibleJobStatusValues() []JobStatus {
	return original.PossibleJobStatusValues()
}
func PossibleJobTypeValues() []JobType {
	return original.PossibleJobTypeValues()
}
func PossibleOperationMonitoringLevelValues() []OperationMonitoringLevel {
	return original.PossibleOperationMonitoringLevelValues()
}
func PossibleRoutingSourceValues() []RoutingSource {
	return original.PossibleRoutingSourceValues()
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
