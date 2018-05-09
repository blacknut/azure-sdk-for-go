package policy

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

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// Mode enumerates the values for mode.
type Mode string

const (
	// All ...
	All Mode = "All"
	// Indexed ...
	Indexed Mode = "Indexed"
	// NotSpecified ...
	NotSpecified Mode = "NotSpecified"
)

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// PossibleModeValues returns an array of possible values for the Mode const type.
func PossibleModeValues() []Mode {
	return []Mode{All, Indexed, NotSpecified}
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// Type enumerates the values for type.
type Type string

const (
	// TypeBuiltIn ...
	TypeBuiltIn Type = "BuiltIn"
	// TypeCustom ...
	TypeCustom Type = "Custom"
	// TypeNotSpecified ...
	TypeNotSpecified Type = "NotSpecified"
)

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{TypeBuiltIn, TypeCustom, TypeNotSpecified}
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// Assignment the policy assignment.
type Assignment struct {
	autorest.Response `json:"-"`
	// AssignmentProperties - Properties for the policy assignment.
	*AssignmentProperties `json:"properties,omitempty"`
	// ID - The ID of the policy assignment.
	ID *string `json:"id,omitempty"`
	// Type - The type of the policy assignment.
	Type *string `json:"type,omitempty"`
	// Name - The name of the policy assignment.
	Name *string `json:"name,omitempty"`
	// Sku - The policy sku.
	Sku *Sku `json:"sku,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// MarshalJSON is the custom marshaler for Assignment.
func (a Assignment) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if a.AssignmentProperties != nil {
		objectMap["properties"] = a.AssignmentProperties
	}
	if a.ID != nil {
		objectMap["id"] = a.ID
	}
	if a.Type != nil {
		objectMap["type"] = a.Type
	}
	if a.Name != nil {
		objectMap["name"] = a.Name
	}
	if a.Sku != nil {
		objectMap["sku"] = a.Sku
	}
	return json.Marshal(objectMap)
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// UnmarshalJSON is the custom unmarshaler for Assignment struct.
func (a *Assignment) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var assignmentProperties AssignmentProperties
				err = json.Unmarshal(*v, &assignmentProperties)
				if err != nil {
					return err
				}
				a.AssignmentProperties = &assignmentProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				a.ID = &ID
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				a.Type = &typeVar
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				a.Name = &name
			}
		case "sku":
			if v != nil {
				var sku Sku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				a.Sku = &sku
			}
		}
	}

	return nil
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// AssignmentListResult list of policy assignments.
type AssignmentListResult struct {
	autorest.Response `json:"-"`
	// Value - An array of policy assignments.
	Value *[]Assignment `json:"value,omitempty"`
	// NextLink - The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// AssignmentListResultIterator provides access to a complete listing of Assignment values.
type AssignmentListResultIterator struct {
	i    int
	page AssignmentListResultPage
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *AssignmentListResultIterator) Next() error {
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

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter AssignmentListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// Response returns the raw server response from the last page request.
func (iter AssignmentListResultIterator) Response() AssignmentListResult {
	return iter.page.Response()
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter AssignmentListResultIterator) Value() Assignment {
	if !iter.page.NotDone() {
		return Assignment{}
	}
	return iter.page.Values()[iter.i]
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// IsEmpty returns true if the ListResult contains no values.
func (alr AssignmentListResult) IsEmpty() bool {
	return alr.Value == nil || len(*alr.Value) == 0
}

// assignmentListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (alr AssignmentListResult) assignmentListResultPreparer() (*http.Request, error) {
	if alr.NextLink == nil || len(to.String(alr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(alr.NextLink)))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// AssignmentListResultPage contains a page of Assignment values.
type AssignmentListResultPage struct {
	fn  func(AssignmentListResult) (AssignmentListResult, error)
	alr AssignmentListResult
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *AssignmentListResultPage) Next() error {
	next, err := page.fn(page.alr)
	if err != nil {
		return err
	}
	page.alr = next
	return nil
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page AssignmentListResultPage) NotDone() bool {
	return !page.alr.IsEmpty()
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// Response returns the raw server response from the last page request.
func (page AssignmentListResultPage) Response() AssignmentListResult {
	return page.alr
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// Values returns the slice of values for the current page or nil if there are no values.
func (page AssignmentListResultPage) Values() []Assignment {
	if page.alr.IsEmpty() {
		return nil
	}
	return *page.alr.Value
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// AssignmentProperties the policy assignment properties.
type AssignmentProperties struct {
	// DisplayName - The display name of the policy assignment.
	DisplayName *string `json:"displayName,omitempty"`
	// PolicyDefinitionID - The ID of the policy definition.
	PolicyDefinitionID *string `json:"policyDefinitionId,omitempty"`
	// Scope - The scope for the policy assignment.
	Scope *string `json:"scope,omitempty"`
	// NotScopes - The policy's excluded scopes.
	NotScopes *[]string `json:"notScopes,omitempty"`
	// Parameters - Required if a parameter is used in policy rule.
	Parameters interface{} `json:"parameters,omitempty"`
	// Description - This message will be part of response in case of policy violation.
	Description *string `json:"description,omitempty"`
	// Metadata - The policy assignment metadata.
	Metadata interface{} `json:"metadata,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// Definition the policy definition.
type Definition struct {
	autorest.Response `json:"-"`
	// DefinitionProperties - The policy definition properties.
	*DefinitionProperties `json:"properties,omitempty"`
	// ID - The ID of the policy definition.
	ID *string `json:"id,omitempty"`
	// Name - The name of the policy definition.
	Name *string `json:"name,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// MarshalJSON is the custom marshaler for Definition.
func (d Definition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if d.DefinitionProperties != nil {
		objectMap["properties"] = d.DefinitionProperties
	}
	if d.ID != nil {
		objectMap["id"] = d.ID
	}
	if d.Name != nil {
		objectMap["name"] = d.Name
	}
	return json.Marshal(objectMap)
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// UnmarshalJSON is the custom unmarshaler for Definition struct.
func (d *Definition) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var definitionProperties DefinitionProperties
				err = json.Unmarshal(*v, &definitionProperties)
				if err != nil {
					return err
				}
				d.DefinitionProperties = &definitionProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				d.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				d.Name = &name
			}
		}
	}

	return nil
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// DefinitionListResult list of policy definitions.
type DefinitionListResult struct {
	autorest.Response `json:"-"`
	// Value - An array of policy definitions.
	Value *[]Definition `json:"value,omitempty"`
	// NextLink - The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// DefinitionListResultIterator provides access to a complete listing of Definition values.
type DefinitionListResultIterator struct {
	i    int
	page DefinitionListResultPage
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *DefinitionListResultIterator) Next() error {
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

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter DefinitionListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// Response returns the raw server response from the last page request.
func (iter DefinitionListResultIterator) Response() DefinitionListResult {
	return iter.page.Response()
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter DefinitionListResultIterator) Value() Definition {
	if !iter.page.NotDone() {
		return Definition{}
	}
	return iter.page.Values()[iter.i]
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// IsEmpty returns true if the ListResult contains no values.
func (dlr DefinitionListResult) IsEmpty() bool {
	return dlr.Value == nil || len(*dlr.Value) == 0
}

// definitionListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (dlr DefinitionListResult) definitionListResultPreparer() (*http.Request, error) {
	if dlr.NextLink == nil || len(to.String(dlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(dlr.NextLink)))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// DefinitionListResultPage contains a page of Definition values.
type DefinitionListResultPage struct {
	fn  func(DefinitionListResult) (DefinitionListResult, error)
	dlr DefinitionListResult
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *DefinitionListResultPage) Next() error {
	next, err := page.fn(page.dlr)
	if err != nil {
		return err
	}
	page.dlr = next
	return nil
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page DefinitionListResultPage) NotDone() bool {
	return !page.dlr.IsEmpty()
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// Response returns the raw server response from the last page request.
func (page DefinitionListResultPage) Response() DefinitionListResult {
	return page.dlr
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// Values returns the slice of values for the current page or nil if there are no values.
func (page DefinitionListResultPage) Values() []Definition {
	if page.dlr.IsEmpty() {
		return nil
	}
	return *page.dlr.Value
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// DefinitionProperties the policy definition properties.
type DefinitionProperties struct {
	// PolicyType - The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom. Possible values include: 'TypeNotSpecified', 'TypeBuiltIn', 'TypeCustom'
	PolicyType Type `json:"policyType,omitempty"`
	// Mode - The policy definition mode. Possible values are NotSpecified, Indexed, and All. Possible values include: 'NotSpecified', 'Indexed', 'All'
	Mode Mode `json:"mode,omitempty"`
	// DisplayName - The display name of the policy definition.
	DisplayName *string `json:"displayName,omitempty"`
	// Description - The policy definition description.
	Description *string `json:"description,omitempty"`
	// PolicyRule - The policy rule.
	PolicyRule interface{} `json:"policyRule,omitempty"`
	// Metadata - The policy definition metadata.
	Metadata interface{} `json:"metadata,omitempty"`
	// Parameters - Required if a parameter is used in policy rule.
	Parameters interface{} `json:"parameters,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// DefinitionReference the policy definition reference.
type DefinitionReference struct {
	// PolicyDefinitionID - The ID of the policy definition or policy set definition.
	PolicyDefinitionID *string `json:"policyDefinitionId,omitempty"`
	// Parameters - Required if a parameter is used in policy rule.
	Parameters interface{} `json:"parameters,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// ErrorResponse error reponse indicates ARM is not able to process the incoming request. The reason is provided in
// the error message.
type ErrorResponse struct {
	// HTTPStatus - Http status code.
	HTTPStatus *string `json:"httpStatus,omitempty"`
	// ErrorCode - Error code.
	ErrorCode *string `json:"errorCode,omitempty"`
	// ErrorMessage - Error message indicating why the operation failed.
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// SetDefinition the policy set definition.
type SetDefinition struct {
	autorest.Response `json:"-"`
	// SetDefinitionProperties - The policy definition properties.
	*SetDefinitionProperties `json:"properties,omitempty"`
	// ID - The ID of the policy set definition.
	ID *string `json:"id,omitempty"`
	// Name - The name of the policy set definition.
	Name *string `json:"name,omitempty"`
	// Type - The type of the resource (Microsoft.Authorization/policySetDefinitions).
	Type *string `json:"type,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// MarshalJSON is the custom marshaler for SetDefinition.
func (sd SetDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sd.SetDefinitionProperties != nil {
		objectMap["properties"] = sd.SetDefinitionProperties
	}
	if sd.ID != nil {
		objectMap["id"] = sd.ID
	}
	if sd.Name != nil {
		objectMap["name"] = sd.Name
	}
	if sd.Type != nil {
		objectMap["type"] = sd.Type
	}
	return json.Marshal(objectMap)
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// UnmarshalJSON is the custom unmarshaler for SetDefinition struct.
func (sd *SetDefinition) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var setDefinitionProperties SetDefinitionProperties
				err = json.Unmarshal(*v, &setDefinitionProperties)
				if err != nil {
					return err
				}
				sd.SetDefinitionProperties = &setDefinitionProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				sd.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				sd.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				sd.Type = &typeVar
			}
		}
	}

	return nil
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// SetDefinitionListResult list of policy set definitions.
type SetDefinitionListResult struct {
	autorest.Response `json:"-"`
	// Value - An array of policy set definitions.
	Value *[]SetDefinition `json:"value,omitempty"`
	// NextLink - The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// SetDefinitionListResultIterator provides access to a complete listing of SetDefinition values.
type SetDefinitionListResultIterator struct {
	i    int
	page SetDefinitionListResultPage
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *SetDefinitionListResultIterator) Next() error {
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

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter SetDefinitionListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// Response returns the raw server response from the last page request.
func (iter SetDefinitionListResultIterator) Response() SetDefinitionListResult {
	return iter.page.Response()
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter SetDefinitionListResultIterator) Value() SetDefinition {
	if !iter.page.NotDone() {
		return SetDefinition{}
	}
	return iter.page.Values()[iter.i]
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// IsEmpty returns true if the ListResult contains no values.
func (sdlr SetDefinitionListResult) IsEmpty() bool {
	return sdlr.Value == nil || len(*sdlr.Value) == 0
}

// setDefinitionListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (sdlr SetDefinitionListResult) setDefinitionListResultPreparer() (*http.Request, error) {
	if sdlr.NextLink == nil || len(to.String(sdlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(sdlr.NextLink)))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// SetDefinitionListResultPage contains a page of SetDefinition values.
type SetDefinitionListResultPage struct {
	fn   func(SetDefinitionListResult) (SetDefinitionListResult, error)
	sdlr SetDefinitionListResult
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *SetDefinitionListResultPage) Next() error {
	next, err := page.fn(page.sdlr)
	if err != nil {
		return err
	}
	page.sdlr = next
	return nil
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page SetDefinitionListResultPage) NotDone() bool {
	return !page.sdlr.IsEmpty()
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// Response returns the raw server response from the last page request.
func (page SetDefinitionListResultPage) Response() SetDefinitionListResult {
	return page.sdlr
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// Values returns the slice of values for the current page or nil if there are no values.
func (page SetDefinitionListResultPage) Values() []SetDefinition {
	if page.sdlr.IsEmpty() {
		return nil
	}
	return *page.sdlr.Value
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// SetDefinitionProperties the policy set definition properties.
type SetDefinitionProperties struct {
	// PolicyType - The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom. Possible values include: 'TypeNotSpecified', 'TypeBuiltIn', 'TypeCustom'
	PolicyType Type `json:"policyType,omitempty"`
	// DisplayName - The display name of the policy set definition.
	DisplayName *string `json:"displayName,omitempty"`
	// Description - The policy set definition description.
	Description *string `json:"description,omitempty"`
	// Metadata - The policy set definition metadata.
	Metadata interface{} `json:"metadata,omitempty"`
	// Parameters - The policy set definition parameters that can be used in policy definition references.
	Parameters interface{} `json:"parameters,omitempty"`
	// PolicyDefinitions - An array of policy definition references.
	PolicyDefinitions *[]DefinitionReference `json:"policyDefinitions,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/resources/mgmt/2017-06-01-preview/policy instead.
// Sku the policy sku.
type Sku struct {
	// Name - The name of the policy sku. Possible values are A0 and A1.
	Name *string `json:"name,omitempty"`
	// Tier - The policy sku tier. Possible values are Free and Standard.
	Tier *string `json:"tier,omitempty"`
}
