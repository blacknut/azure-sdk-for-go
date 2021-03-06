package operationalinsights

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
)

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// DataSourceKind enumerates the values for data source kind.
type DataSourceKind string

const (
	// AzureActivityLog ...
	AzureActivityLog DataSourceKind = "AzureActivityLog"
	// ChangeTrackingCustomRegistry ...
	ChangeTrackingCustomRegistry DataSourceKind = "ChangeTrackingCustomRegistry"
	// ChangeTrackingDefaultPath ...
	ChangeTrackingDefaultPath DataSourceKind = "ChangeTrackingDefaultPath"
	// ChangeTrackingDefaultRegistry ...
	ChangeTrackingDefaultRegistry DataSourceKind = "ChangeTrackingDefaultRegistry"
	// ChangeTrackingPath ...
	ChangeTrackingPath DataSourceKind = "ChangeTrackingPath"
	// CustomLog ...
	CustomLog DataSourceKind = "CustomLog"
	// CustomLogCollection ...
	CustomLogCollection DataSourceKind = "CustomLogCollection"
	// GenericDataSource ...
	GenericDataSource DataSourceKind = "GenericDataSource"
	// IISLogs ...
	IISLogs DataSourceKind = "IISLogs"
	// LinuxPerformanceCollection ...
	LinuxPerformanceCollection DataSourceKind = "LinuxPerformanceCollection"
	// LinuxPerformanceObject ...
	LinuxPerformanceObject DataSourceKind = "LinuxPerformanceObject"
	// LinuxSyslog ...
	LinuxSyslog DataSourceKind = "LinuxSyslog"
	// LinuxSyslogCollection ...
	LinuxSyslogCollection DataSourceKind = "LinuxSyslogCollection"
	// WindowsEvent ...
	WindowsEvent DataSourceKind = "WindowsEvent"
	// WindowsPerformanceCounter ...
	WindowsPerformanceCounter DataSourceKind = "WindowsPerformanceCounter"
)

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// PossibleDataSourceKindValues returns an array of possible values for the DataSourceKind const type.
func PossibleDataSourceKindValues() []DataSourceKind {
	return []DataSourceKind{AzureActivityLog, ChangeTrackingCustomRegistry, ChangeTrackingDefaultPath, ChangeTrackingDefaultRegistry, ChangeTrackingPath, CustomLog, CustomLogCollection, GenericDataSource, IISLogs, LinuxPerformanceCollection, LinuxPerformanceObject, LinuxSyslog, LinuxSyslogCollection, WindowsEvent, WindowsPerformanceCounter}
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// EntityStatus enumerates the values for entity status.
type EntityStatus string

const (
	// Canceled ...
	Canceled EntityStatus = "Canceled"
	// Creating ...
	Creating EntityStatus = "Creating"
	// Deleting ...
	Deleting EntityStatus = "Deleting"
	// Failed ...
	Failed EntityStatus = "Failed"
	// ProvisioningAccount ...
	ProvisioningAccount EntityStatus = "ProvisioningAccount"
	// Succeeded ...
	Succeeded EntityStatus = "Succeeded"
)

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// PossibleEntityStatusValues returns an array of possible values for the EntityStatus const type.
func PossibleEntityStatusValues() []EntityStatus {
	return []EntityStatus{Canceled, Creating, Deleting, Failed, ProvisioningAccount, Succeeded}
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// SkuNameEnum enumerates the values for sku name enum.
type SkuNameEnum string

const (
	// Free ...
	Free SkuNameEnum = "Free"
	// PerGB2018 ...
	PerGB2018 SkuNameEnum = "PerGB2018"
	// PerNode ...
	PerNode SkuNameEnum = "PerNode"
	// Premium ...
	Premium SkuNameEnum = "Premium"
	// Standalone ...
	Standalone SkuNameEnum = "Standalone"
	// Standard ...
	Standard SkuNameEnum = "Standard"
	// Unlimited ...
	Unlimited SkuNameEnum = "Unlimited"
)

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// PossibleSkuNameEnumValues returns an array of possible values for the SkuNameEnum const type.
func PossibleSkuNameEnumValues() []SkuNameEnum {
	return []SkuNameEnum{Free, PerGB2018, PerNode, Premium, Standalone, Standard, Unlimited}
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// DataSource datasources under OMS Workspace.
type DataSource struct {
	autorest.Response `json:"-"`
	// Properties - The data source properties in raw json format, each kind of data source have it's own schema.
	Properties interface{} `json:"properties,omitempty"`
	// ETag - The ETag of the data source.
	ETag *string `json:"eTag,omitempty"`
	// Kind - Possible values include: 'AzureActivityLog', 'ChangeTrackingPath', 'ChangeTrackingDefaultPath', 'ChangeTrackingDefaultRegistry', 'ChangeTrackingCustomRegistry', 'CustomLog', 'CustomLogCollection', 'GenericDataSource', 'IISLogs', 'LinuxPerformanceObject', 'LinuxPerformanceCollection', 'LinuxSyslog', 'LinuxSyslogCollection', 'WindowsEvent', 'WindowsPerformanceCounter'
	Kind DataSourceKind `json:"kind,omitempty"`
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// MarshalJSON is the custom marshaler for DataSource.
func (ds DataSource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["properties"] = ds.Properties
	if ds.ETag != nil {
		objectMap["eTag"] = ds.ETag
	}
	if ds.Kind != "" {
		objectMap["kind"] = ds.Kind
	}
	if ds.ID != nil {
		objectMap["id"] = ds.ID
	}
	if ds.Name != nil {
		objectMap["name"] = ds.Name
	}
	if ds.Type != nil {
		objectMap["type"] = ds.Type
	}
	if ds.Tags != nil {
		objectMap["tags"] = ds.Tags
	}
	return json.Marshal(objectMap)
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// DataSourceFilter dataSource filter. Right now, only filter by kind is supported.
type DataSourceFilter struct {
	// Kind - Possible values include: 'AzureActivityLog', 'ChangeTrackingPath', 'ChangeTrackingDefaultPath', 'ChangeTrackingDefaultRegistry', 'ChangeTrackingCustomRegistry', 'CustomLog', 'CustomLogCollection', 'GenericDataSource', 'IISLogs', 'LinuxPerformanceObject', 'LinuxPerformanceCollection', 'LinuxSyslog', 'LinuxSyslogCollection', 'WindowsEvent', 'WindowsPerformanceCounter'
	Kind DataSourceKind `json:"kind,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// DataSourceListResult the list data source by workspace operation response.
type DataSourceListResult struct {
	autorest.Response `json:"-"`
	// Value - A list of datasources.
	Value *[]DataSource `json:"value,omitempty"`
	// NextLink - The link (url) to the next page of datasources.
	NextLink *string `json:"nextLink,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// DataSourceListResultIterator provides access to a complete listing of DataSource values.
type DataSourceListResultIterator struct {
	i    int
	page DataSourceListResultPage
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *DataSourceListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter DataSourceListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// Response returns the raw server response from the last page request.
func (iter DataSourceListResultIterator) Response() DataSourceListResult {
	return iter.page.Response()
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter DataSourceListResultIterator) Value() DataSource {
	if !iter.page.NotDone() {
		return DataSource{}
	}
	return iter.page.Values()[iter.i]
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// IsEmpty returns true if the ListResult contains no values.
func (dslr DataSourceListResult) IsEmpty() bool {
	return dslr.Value == nil || len(*dslr.Value) == 0
}

// dataSourceListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (dslr DataSourceListResult) dataSourceListResultPreparer() (*http.Request, error) {
	if dslr.NextLink == nil || len(to.String(dslr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(dslr.NextLink)))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// DataSourceListResultPage contains a page of DataSource values.
type DataSourceListResultPage struct {
	fn   func(DataSourceListResult) (DataSourceListResult, error)
	dslr DataSourceListResult
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *DataSourceListResultPage) Next() error {
	next, err := page.fn(page.dslr)
	if err != nil {
		return err
	}
	page.dslr = next
	return nil
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page DataSourceListResultPage) NotDone() bool {
	return !page.dslr.IsEmpty()
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// Response returns the raw server response from the last page request.
func (page DataSourceListResultPage) Response() DataSourceListResult {
	return page.dslr
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// Values returns the slice of values for the current page or nil if there are no values.
func (page DataSourceListResultPage) Values() []DataSource {
	if page.dslr.IsEmpty() {
		return nil
	}
	return *page.dslr.Value
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// IntelligencePack intelligence Pack containing a string name and boolean indicating if it's enabled.
type IntelligencePack struct {
	// Name - The name of the intelligence pack.
	Name *string `json:"name,omitempty"`
	// Enabled - The enabled boolean for the intelligence pack.
	Enabled *bool `json:"enabled,omitempty"`
	// DisplayName - The display name of the intelligence pack.
	DisplayName *string `json:"displayName,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// LinkedService the top level Linked service resource container.
type LinkedService struct {
	autorest.Response `json:"-"`
	// LinkedServiceProperties - The properties of the linked service.
	*LinkedServiceProperties `json:"properties,omitempty"`
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// MarshalJSON is the custom marshaler for LinkedService.
func (ls LinkedService) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ls.LinkedServiceProperties != nil {
		objectMap["properties"] = ls.LinkedServiceProperties
	}
	if ls.ID != nil {
		objectMap["id"] = ls.ID
	}
	if ls.Name != nil {
		objectMap["name"] = ls.Name
	}
	if ls.Type != nil {
		objectMap["type"] = ls.Type
	}
	if ls.Tags != nil {
		objectMap["tags"] = ls.Tags
	}
	return json.Marshal(objectMap)
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// UnmarshalJSON is the custom unmarshaler for LinkedService struct.
func (ls *LinkedService) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var linkedServiceProperties LinkedServiceProperties
				err = json.Unmarshal(*v, &linkedServiceProperties)
				if err != nil {
					return err
				}
				ls.LinkedServiceProperties = &linkedServiceProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				ls.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				ls.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				ls.Type = &typeVar
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				ls.Tags = tags
			}
		}
	}

	return nil
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// LinkedServiceListResult the list linked service operation response.
type LinkedServiceListResult struct {
	autorest.Response `json:"-"`
	// Value - Gets or sets a list of linked service instances.
	Value *[]LinkedService `json:"value,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// LinkedServiceProperties linked service properties.
type LinkedServiceProperties struct {
	// ResourceID - The resource id of the resource that will be linked to the workspace.
	ResourceID *string `json:"resourceId,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// ListIntelligencePack ...
type ListIntelligencePack struct {
	autorest.Response `json:"-"`
	Value             *[]IntelligencePack `json:"value,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// ManagementGroup a management group that is connected to a workspace
type ManagementGroup struct {
	// ManagementGroupProperties - The properties of the management group.
	*ManagementGroupProperties `json:"properties,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// MarshalJSON is the custom marshaler for ManagementGroup.
func (mg ManagementGroup) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if mg.ManagementGroupProperties != nil {
		objectMap["properties"] = mg.ManagementGroupProperties
	}
	return json.Marshal(objectMap)
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// UnmarshalJSON is the custom unmarshaler for ManagementGroup struct.
func (mg *ManagementGroup) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var managementGroupProperties ManagementGroupProperties
				err = json.Unmarshal(*v, &managementGroupProperties)
				if err != nil {
					return err
				}
				mg.ManagementGroupProperties = &managementGroupProperties
			}
		}
	}

	return nil
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// ManagementGroupProperties management group properties.
type ManagementGroupProperties struct {
	// ServerCount - The number of servers connected to the management group.
	ServerCount *int32 `json:"serverCount,omitempty"`
	// IsGateway - Gets or sets a value indicating whether the management group is a gateway.
	IsGateway *bool `json:"isGateway,omitempty"`
	// Name - The name of the management group.
	Name *string `json:"name,omitempty"`
	// ID - The unique ID of the management group.
	ID *string `json:"id,omitempty"`
	// Created - The datetime that the management group was created.
	Created *date.Time `json:"created,omitempty"`
	// DataReceived - The last datetime that the management group received data.
	DataReceived *date.Time `json:"dataReceived,omitempty"`
	// Version - The version of System Center that is managing the management group.
	Version *string `json:"version,omitempty"`
	// Sku - The SKU of System Center that is managing the management group.
	Sku *string `json:"sku,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// MetricName the name of a metric.
type MetricName struct {
	// Value - The system name of the metric.
	Value *string `json:"value,omitempty"`
	// LocalizedValue - The localized name of the metric.
	LocalizedValue *string `json:"localizedValue,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// Operation supported operation of OperationalInsights resource provider.
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - Display metadata associated with the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// OperationDisplay display metadata associated with the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft OperationsManagement.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Type of operation: get, read, delete, etc.
	Operation *string `json:"operation,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// OperationListResult result of the request to list solution operations.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of solution operations supported by the OperationsManagement resource provider.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// OperationListResultIterator provides access to a complete listing of Operation values.
type OperationListResultIterator struct {
	i    int
	page OperationListResultPage
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// Response returns the raw server response from the last page request.
func (iter OperationListResultIterator) Response() OperationListResult {
	return iter.page.Response()
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListResultIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// IsEmpty returns true if the ListResult contains no values.
func (olr OperationListResult) IsEmpty() bool {
	return olr.Value == nil || len(*olr.Value) == 0
}

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer() (*http.Request, error) {
	if olr.NextLink == nil || len(to.String(olr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(olr.NextLink)))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// OperationListResultPage contains a page of Operation values.
type OperationListResultPage struct {
	fn  func(OperationListResult) (OperationListResult, error)
	olr OperationListResult
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListResultPage) Next() error {
	next, err := page.fn(page.olr)
	if err != nil {
		return err
	}
	page.olr = next
	return nil
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListResultPage) NotDone() bool {
	return !page.olr.IsEmpty()
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// Response returns the raw server response from the last page request.
func (page OperationListResultPage) Response() OperationListResult {
	return page.olr
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListResultPage) Values() []Operation {
	if page.olr.IsEmpty() {
		return nil
	}
	return *page.olr.Value
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// ProxyResource common properties of proxy resource.
type ProxyResource struct {
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// MarshalJSON is the custom marshaler for ProxyResource.
func (pr ProxyResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if pr.ID != nil {
		objectMap["id"] = pr.ID
	}
	if pr.Name != nil {
		objectMap["name"] = pr.Name
	}
	if pr.Type != nil {
		objectMap["type"] = pr.Type
	}
	if pr.Tags != nil {
		objectMap["tags"] = pr.Tags
	}
	return json.Marshal(objectMap)
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// Resource the resource definition.
type Resource struct {
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
	// Name - Resource name
	Name *string `json:"name,omitempty"`
	// Type - Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.ID != nil {
		objectMap["id"] = r.ID
	}
	if r.Name != nil {
		objectMap["name"] = r.Name
	}
	if r.Type != nil {
		objectMap["type"] = r.Type
	}
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	return json.Marshal(objectMap)
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// SharedKeys the shared keys for a workspace.
type SharedKeys struct {
	autorest.Response `json:"-"`
	// PrimarySharedKey - The primary shared key of a workspace.
	PrimarySharedKey *string `json:"primarySharedKey,omitempty"`
	// SecondarySharedKey - The secondary shared key of a workspace.
	SecondarySharedKey *string `json:"secondarySharedKey,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// Sku the SKU (tier) of a workspace.
type Sku struct {
	// Name - The name of the SKU. Possible values include: 'Free', 'Standard', 'Premium', 'Unlimited', 'PerNode', 'PerGB2018', 'Standalone'
	Name SkuNameEnum `json:"name,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// UsageMetric a metric describing the usage of a resource.
type UsageMetric struct {
	// Name - The name of the metric.
	Name *MetricName `json:"name,omitempty"`
	// Unit - The units used for the metric.
	Unit *string `json:"unit,omitempty"`
	// CurrentValue - The current value of the metric.
	CurrentValue *float64 `json:"currentValue,omitempty"`
	// Limit - The quota limit for the metric.
	Limit *float64 `json:"limit,omitempty"`
	// NextResetTime - The time that the metric's value will reset.
	NextResetTime *date.Time `json:"nextResetTime,omitempty"`
	// QuotaPeriod - The quota period that determines the length of time between value resets.
	QuotaPeriod *string `json:"quotaPeriod,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// Workspace the top level Workspace resource container.
type Workspace struct {
	autorest.Response `json:"-"`
	// WorkspaceProperties - Workspace properties.
	*WorkspaceProperties `json:"properties,omitempty"`
	// ETag - The ETag of the workspace.
	ETag *string `json:"eTag,omitempty"`
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
	// Name - Resource name
	Name *string `json:"name,omitempty"`
	// Type - Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// MarshalJSON is the custom marshaler for Workspace.
func (w Workspace) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if w.WorkspaceProperties != nil {
		objectMap["properties"] = w.WorkspaceProperties
	}
	if w.ETag != nil {
		objectMap["eTag"] = w.ETag
	}
	if w.ID != nil {
		objectMap["id"] = w.ID
	}
	if w.Name != nil {
		objectMap["name"] = w.Name
	}
	if w.Type != nil {
		objectMap["type"] = w.Type
	}
	if w.Location != nil {
		objectMap["location"] = w.Location
	}
	if w.Tags != nil {
		objectMap["tags"] = w.Tags
	}
	return json.Marshal(objectMap)
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// UnmarshalJSON is the custom unmarshaler for Workspace struct.
func (w *Workspace) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var workspaceProperties WorkspaceProperties
				err = json.Unmarshal(*v, &workspaceProperties)
				if err != nil {
					return err
				}
				w.WorkspaceProperties = &workspaceProperties
			}
		case "eTag":
			if v != nil {
				var eTag string
				err = json.Unmarshal(*v, &eTag)
				if err != nil {
					return err
				}
				w.ETag = &eTag
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				w.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				w.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				w.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				w.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				w.Tags = tags
			}
		}
	}

	return nil
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// WorkspaceListManagementGroupsResult the list workspace managmement groups operation response.
type WorkspaceListManagementGroupsResult struct {
	autorest.Response `json:"-"`
	// Value - Gets or sets a list of management groups attached to the workspace.
	Value *[]ManagementGroup `json:"value,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// WorkspaceListResult the list workspaces operation response.
type WorkspaceListResult struct {
	autorest.Response `json:"-"`
	// Value - A list of workspaces.
	Value *[]Workspace `json:"value,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// WorkspaceListUsagesResult the list workspace usages operation response.
type WorkspaceListUsagesResult struct {
	autorest.Response `json:"-"`
	// Value - Gets or sets a list of usage metrics for a workspace.
	Value *[]UsageMetric `json:"value,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// WorkspaceProperties workspace properties.
type WorkspaceProperties struct {
	// ProvisioningState - The provisioning state of the workspace. Possible values include: 'Creating', 'Succeeded', 'Failed', 'Canceled', 'Deleting', 'ProvisioningAccount'
	ProvisioningState EntityStatus `json:"provisioningState,omitempty"`
	// Source - The source of the workspace.  Source defines where the workspace was created. 'Azure' implies it was created in Azure.  'External' implies it was created via the Operational Insights Portal. This value is set on the service side and read-only on the client side.
	Source *string `json:"source,omitempty"`
	// CustomerID - The ID associated with the workspace.  Setting this value at creation time allows the workspace being created to be linked to an existing workspace.
	CustomerID *string `json:"customerId,omitempty"`
	// PortalURL - The URL of the Operational Insights portal for this workspace.  This value is set on the service side and read-only on the client side.
	PortalURL *string `json:"portalUrl,omitempty"`
	// Sku - The SKU of the workspace.
	Sku *Sku `json:"sku,omitempty"`
	// RetentionInDays - The workspace data retention in days. -1 means Unlimited retention for the Unlimited Sku. 730 days is the maximum allowed for all other Skus.
	RetentionInDays *int32 `json:"retentionInDays,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// WorkspacesCreateOrUpdateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type WorkspacesCreateOrUpdateFuture struct {
	azure.Future
	req *http.Request
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights instead.
// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future WorkspacesCreateOrUpdateFuture) Result(client WorkspacesClient) (w Workspace, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacesCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		return w, azure.NewAsyncOpIncompleteError("operationalinsights.WorkspacesCreateOrUpdateFuture")
	}
	if future.PollingMethod() == azure.PollingLocation {
		w, err = client.CreateOrUpdateResponder(future.Response())
		if err != nil {
			err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacesCreateOrUpdateFuture", "Result", future.Response(), "Failure responding to request")
		}
		return
	}
	var req *http.Request
	var resp *http.Response
	if future.PollingURL() != "" {
		req, err = http.NewRequest(http.MethodGet, future.PollingURL(), nil)
		if err != nil {
			return
		}
	} else {
		req = autorest.ChangeToGet(future.req)
	}
	resp, err = autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacesCreateOrUpdateFuture", "Result", resp, "Failure sending request")
		return
	}
	w, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacesCreateOrUpdateFuture", "Result", resp, "Failure responding to request")
	}
	return
}
