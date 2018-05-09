package hanaonazure

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
	"github.com/Azure/go-autorest/autorest/to"
)

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// HanaHardwareTypeNamesEnum enumerates the values for hana hardware type names enum.
type HanaHardwareTypeNamesEnum string

const (
	// CiscoUCS ...
	CiscoUCS HanaHardwareTypeNamesEnum = "Cisco_UCS"
)

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// PossibleHanaHardwareTypeNamesEnumValues returns an array of possible values for the HanaHardwareTypeNamesEnum const type.
func PossibleHanaHardwareTypeNamesEnumValues() []HanaHardwareTypeNamesEnum {
	return []HanaHardwareTypeNamesEnum{CiscoUCS}
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// HanaInstanceSizeNamesEnum enumerates the values for hana instance size names enum.
type HanaInstanceSizeNamesEnum string

const (
	// S144 ...
	S144 HanaInstanceSizeNamesEnum = "S144"
	// S144m ...
	S144m HanaInstanceSizeNamesEnum = "S144m"
	// S192 ...
	S192 HanaInstanceSizeNamesEnum = "S192"
	// S192m ...
	S192m HanaInstanceSizeNamesEnum = "S192m"
	// S72 ...
	S72 HanaInstanceSizeNamesEnum = "S72"
	// S72m ...
	S72m HanaInstanceSizeNamesEnum = "S72m"
)

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// PossibleHanaInstanceSizeNamesEnumValues returns an array of possible values for the HanaInstanceSizeNamesEnum const type.
func PossibleHanaInstanceSizeNamesEnumValues() []HanaInstanceSizeNamesEnum {
	return []HanaInstanceSizeNamesEnum{S144, S144m, S192, S192m, S72, S72m}
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// Disk specifies the disk information fo the HANA instance
type Disk struct {
	// Name - The disk name.
	Name *string `json:"name,omitempty"`
	// DiskSizeGB - Specifies the size of an empty data disk in gigabytes.
	DiskSizeGB *int32 `json:"diskSizeGB,omitempty"`
	// Lun - Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and therefore must be unique for each data disk attached to a VM.
	Lun *int32 `json:"lun,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// Display detailed HANA operation information
type Display struct {
	// Provider - The localized friendly form of the resource provider name. This form is also expected to include the publisher/company responsible. Use Title Casing. Begin with “Microsoft” for 1st party services.
	Provider *string `json:"provider,omitempty"`
	// Resource - The localized friendly form of the resource type related to this action/operation. This form should match the public documentation for the resource provider. Use Title Casing. For examples, refer to the “name” section.
	Resource *string `json:"resource,omitempty"`
	// Operation - The localized friendly name for the operation as shown to the user. This name should be concise (to fit in drop downs), but clear (self-documenting). Use Title Casing and include the entity/resource to which it applies.
	Operation *string `json:"operation,omitempty"`
	// Description - The localized friendly description for the operation as shown to the user. This description should be thorough, yet concise. It will be used in tool-tips and detailed views.
	Description *string `json:"description,omitempty"`
	// Origin - The intended executor of the operation; governs the display of the operation in the RBAC UX and the audit logs UX. Default value is 'user,system'
	Origin *string `json:"origin,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// ErrorResponse describes the format of Error response.
type ErrorResponse struct {
	// Code - Error code
	Code *string `json:"code,omitempty"`
	// Message - Error message indicating why the operation failed.
	Message *string `json:"message,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// HanaInstance HANA instance info on Azure (ARM properties and HANA properties)
type HanaInstance struct {
	autorest.Response `json:"-"`
	// HanaInstanceProperties - HANA instance properties
	*HanaInstanceProperties `json:"properties,omitempty"`
	// ID - Resource ID
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

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// MarshalJSON is the custom marshaler for HanaInstance.
func (hi HanaInstance) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if hi.HanaInstanceProperties != nil {
		objectMap["properties"] = hi.HanaInstanceProperties
	}
	if hi.ID != nil {
		objectMap["id"] = hi.ID
	}
	if hi.Name != nil {
		objectMap["name"] = hi.Name
	}
	if hi.Type != nil {
		objectMap["type"] = hi.Type
	}
	if hi.Location != nil {
		objectMap["location"] = hi.Location
	}
	if hi.Tags != nil {
		objectMap["tags"] = hi.Tags
	}
	return json.Marshal(objectMap)
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// UnmarshalJSON is the custom unmarshaler for HanaInstance struct.
func (hi *HanaInstance) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var hanaInstanceProperties HanaInstanceProperties
				err = json.Unmarshal(*v, &hanaInstanceProperties)
				if err != nil {
					return err
				}
				hi.HanaInstanceProperties = &hanaInstanceProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				hi.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				hi.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				hi.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				hi.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				hi.Tags = tags
			}
		}
	}

	return nil
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// HanaInstanceProperties describes the properties of a HANA instance.
type HanaInstanceProperties struct {
	// HardwareProfile - Specifies the hardware settings for the HANA instance.
	HardwareProfile *HardwareProfile `json:"hardwareProfile,omitempty"`
	// StorageProfile - Specifies the storage settings for the HANA instance disks.
	StorageProfile *StorageProfile `json:"storageProfile,omitempty"`
	// OsProfile - Specifies the operating system settings for the HANA instance.
	OsProfile *OSProfile `json:"osProfile,omitempty"`
	// NetworkProfile - Specifies the network settings for the HANA instance.
	NetworkProfile *NetworkProfile `json:"networkProfile,omitempty"`
	// HanaInstanceID - Specifies the HANA instance unique ID.
	HanaInstanceID *string `json:"hanaInstanceId,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// HanaInstancesListResult the response from the List HANA Instances operation.
type HanaInstancesListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of SAP HANA on Azure instances.
	Value *[]HanaInstance `json:"value,omitempty"`
	// NextLink - The URL to get the next set of HANA instances.
	NextLink *string `json:"nextLink,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// HanaInstancesListResultIterator provides access to a complete listing of HanaInstance values.
type HanaInstancesListResultIterator struct {
	i    int
	page HanaInstancesListResultPage
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *HanaInstancesListResultIterator) Next() error {
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

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter HanaInstancesListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// Response returns the raw server response from the last page request.
func (iter HanaInstancesListResultIterator) Response() HanaInstancesListResult {
	return iter.page.Response()
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter HanaInstancesListResultIterator) Value() HanaInstance {
	if !iter.page.NotDone() {
		return HanaInstance{}
	}
	return iter.page.Values()[iter.i]
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// IsEmpty returns true if the ListResult contains no values.
func (hilr HanaInstancesListResult) IsEmpty() bool {
	return hilr.Value == nil || len(*hilr.Value) == 0
}

// hanaInstancesListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (hilr HanaInstancesListResult) hanaInstancesListResultPreparer() (*http.Request, error) {
	if hilr.NextLink == nil || len(to.String(hilr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(hilr.NextLink)))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// HanaInstancesListResultPage contains a page of HanaInstance values.
type HanaInstancesListResultPage struct {
	fn   func(HanaInstancesListResult) (HanaInstancesListResult, error)
	hilr HanaInstancesListResult
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *HanaInstancesListResultPage) Next() error {
	next, err := page.fn(page.hilr)
	if err != nil {
		return err
	}
	page.hilr = next
	return nil
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page HanaInstancesListResultPage) NotDone() bool {
	return !page.hilr.IsEmpty()
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// Response returns the raw server response from the last page request.
func (page HanaInstancesListResultPage) Response() HanaInstancesListResult {
	return page.hilr
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// Values returns the slice of values for the current page or nil if there are no values.
func (page HanaInstancesListResultPage) Values() []HanaInstance {
	if page.hilr.IsEmpty() {
		return nil
	}
	return *page.hilr.Value
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// HardwareProfile specifies the hardware settings for the HANA instance.
type HardwareProfile struct {
	// HardwareType - Name of the hardware type (vendor and/or thrie product name). Possible values include: 'CiscoUCS'
	HardwareType HanaHardwareTypeNamesEnum `json:"hardwareType,omitempty"`
	// HanaInstanceSize - Specifies the HANA instance SKU. Possible values include: 'S72m', 'S144m', 'S72', 'S144', 'S192', 'S192m'
	HanaInstanceSize HanaInstanceSizeNamesEnum `json:"hanaInstanceSize,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// IPAddress specifies the IP address of the network interaface.
type IPAddress struct {
	// IPAddress - Specifies the IP address of the network interface.
	IPAddress *string `json:"ipAddress,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// NetworkProfile specifies the network settings for the HANA instance disks.
type NetworkProfile struct {
	// NetworkInterfaces - Specifies the network interfaces for the HANA instance.
	NetworkInterfaces *[]IPAddress `json:"networkInterfaces,omitempty"`
	// CircuitID - Specifies the circuit id for connecting to express route.
	CircuitID *string `json:"circuitId,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// Operation HANA operation information
type Operation struct {
	// Name - The name of the operation being performed on this particular object. This name should match the action name that appears in RBAC / the event service.
	Name *string `json:"name,omitempty"`
	// Display - Displayed HANA operation information
	Display *Display `json:"display,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// OperationList list of HANA operations
type OperationList struct {
	autorest.Response `json:"-"`
	// Value - List of HANA operations
	Value *[]Operation `json:"value,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// OSProfile specifies the operating system settings for the HANA instance.
type OSProfile struct {
	// ComputerName - Specifies the host OS name of the HANA instance.
	ComputerName *string `json:"computerName,omitempty"`
	// OsType - This property allows you to specify the type of the OS.
	OsType *string `json:"osType,omitempty"`
	// Version - Specifies version of operating system.
	Version *string `json:"version,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// Resource the resource model definition.
type Resource struct {
	// ID - Resource ID
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

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
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

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure instead.
// StorageProfile specifies the storage settings for the HANA instance disks.
type StorageProfile struct {
	// NfsIPAddress - IP Address to connect to storage.
	NfsIPAddress *string `json:"nfsIpAddress,omitempty"`
	// OsDisks - Specifies information about the operating system disk used by the hana instance.
	OsDisks *[]Disk `json:"osDisks,omitempty"`
}
