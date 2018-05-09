package policyinsights

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

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
)

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-10-17-preview/policyinsights instead.
// PolicyStatesResource enumerates the values for policy states resource.
type PolicyStatesResource string

const (
	// Default ...
	Default PolicyStatesResource = "default"
	// Latest ...
	Latest PolicyStatesResource = "latest"
)

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-10-17-preview/policyinsights instead.
// PossiblePolicyStatesResourceValues returns an array of possible values for the PolicyStatesResource const type.
func PossiblePolicyStatesResourceValues() []PolicyStatesResource {
	return []PolicyStatesResource{Default, Latest}
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-10-17-preview/policyinsights instead.
// Operation operation definition.
type Operation struct {
	// Name - Operation name.
	Name *string `json:"name,omitempty"`
	// Display - Display metadata associated with the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-10-17-preview/policyinsights instead.
// OperationDisplay display metadata associated with the operation.
type OperationDisplay struct {
	// Provider - Resource provider name.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource name on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation name.
	Operation *string `json:"operation,omitempty"`
	// Description - Operation description.
	Description *string `json:"description,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-10-17-preview/policyinsights instead.
// OperationsListResults list of available operations.
type OperationsListResults struct {
	autorest.Response `json:"-"`
	// OdataCount - OData entity count; represents the number of operations returned.
	OdataCount *int32 `json:"@odata.count,omitempty"`
	// Value - List of available operations.
	Value *[]Operation `json:"value,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-10-17-preview/policyinsights instead.
// PolicyEvent policy event record.
type PolicyEvent struct {
	// AdditionalProperties - Unmatched properties from the message are deserialized this collection
	AdditionalProperties map[string]interface{} `json:""`
	// OdataID - OData entity ID; always set to null since policy event records do not have an entity ID.
	OdataID *string `json:"@odata.id,omitempty"`
	// OdataContext - OData context string; used by OData clients to resolve type information based on metadata.
	OdataContext *string `json:"@odata.context,omitempty"`
	// Timestamp - Timestamp for the policy event record.
	Timestamp *date.Time `json:"timestamp,omitempty"`
	// ResourceID - Resource ID.
	ResourceID *string `json:"resourceId,omitempty"`
	// PolicyAssignmentID - Policy assignment ID.
	PolicyAssignmentID *string `json:"policyAssignmentId,omitempty"`
	// PolicyDefinitionID - Policy definition ID.
	PolicyDefinitionID *string `json:"policyDefinitionId,omitempty"`
	// EffectiveParameters - Effective parameters for the policy assignment.
	EffectiveParameters *string `json:"effectiveParameters,omitempty"`
	// IsCompliant - Flag which states whether the resource is compliant against the policy assignment it was evaluated against.
	IsCompliant *bool `json:"isCompliant,omitempty"`
	// SubscriptionID - Subscription ID.
	SubscriptionID *string `json:"subscriptionId,omitempty"`
	// ResourceType - Resource type.
	ResourceType *string `json:"resourceType,omitempty"`
	// ResourceLocation - Resource location.
	ResourceLocation *string `json:"resourceLocation,omitempty"`
	// ResourceGroup - Resource group name.
	ResourceGroup *string `json:"resourceGroup,omitempty"`
	// ResourceTags - List of resource tags.
	ResourceTags *string `json:"resourceTags,omitempty"`
	// PolicyAssignmentName - Policy assignment name.
	PolicyAssignmentName *string `json:"policyAssignmentName,omitempty"`
	// PolicyAssignmentOwner - Policy assignment owner.
	PolicyAssignmentOwner *string `json:"policyAssignmentOwner,omitempty"`
	// PolicyAssignmentParameters - Policy assignment parameters.
	PolicyAssignmentParameters *string `json:"policyAssignmentParameters,omitempty"`
	// PolicyAssignmentScope - Policy assignment scope.
	PolicyAssignmentScope *string `json:"policyAssignmentScope,omitempty"`
	// PolicyDefinitionName - Policy definition name.
	PolicyDefinitionName *string `json:"policyDefinitionName,omitempty"`
	// PolicyDefinitionAction - Policy definition action, i.e. effect.
	PolicyDefinitionAction *string `json:"policyDefinitionAction,omitempty"`
	// PolicyDefinitionCategory - Policy definition category.
	PolicyDefinitionCategory *string `json:"policyDefinitionCategory,omitempty"`
	// PolicySetDefinitionID - Policy set definition ID, if the policy assignment is for a policy set.
	PolicySetDefinitionID *string `json:"policySetDefinitionId,omitempty"`
	// PolicySetDefinitionName - Policy set definition name, if the policy assignment is for a policy set.
	PolicySetDefinitionName *string `json:"policySetDefinitionName,omitempty"`
	// PolicySetDefinitionOwner - Policy set definition owner, if the policy assignment is for a policy set.
	PolicySetDefinitionOwner *string `json:"policySetDefinitionOwner,omitempty"`
	// PolicySetDefinitionCategory - Policy set definition category, if the policy assignment is for a policy set.
	PolicySetDefinitionCategory *string `json:"policySetDefinitionCategory,omitempty"`
	// PolicySetDefinitionParameters - Policy set definition parameters, if the policy assignment is for a policy set.
	PolicySetDefinitionParameters *string `json:"policySetDefinitionParameters,omitempty"`
	// ManagementGroupIds - Comma seperated list of management group IDs, which represent the hierarchy of the management groups the resource is under.
	ManagementGroupIds *string `json:"managementGroupIds,omitempty"`
	// PolicyDefinitionReferenceID - Reference ID for the policy definition inside the policy set, if the policy assignment is for a policy set.
	PolicyDefinitionReferenceID *string `json:"policyDefinitionReferenceId,omitempty"`
	// TenantID - Tenant ID for the policy event record.
	TenantID *string `json:"tenantId,omitempty"`
	// PrincipalOid - Principal object ID for the user who initiated the resource operation that triggered the policy event.
	PrincipalOid *string `json:"principalOid,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-10-17-preview/policyinsights instead.
// MarshalJSON is the custom marshaler for PolicyEvent.
func (peVar PolicyEvent) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if peVar.OdataID != nil {
		objectMap["@odata.id"] = peVar.OdataID
	}
	if peVar.OdataContext != nil {
		objectMap["@odata.context"] = peVar.OdataContext
	}
	if peVar.Timestamp != nil {
		objectMap["timestamp"] = peVar.Timestamp
	}
	if peVar.ResourceID != nil {
		objectMap["resourceId"] = peVar.ResourceID
	}
	if peVar.PolicyAssignmentID != nil {
		objectMap["policyAssignmentId"] = peVar.PolicyAssignmentID
	}
	if peVar.PolicyDefinitionID != nil {
		objectMap["policyDefinitionId"] = peVar.PolicyDefinitionID
	}
	if peVar.EffectiveParameters != nil {
		objectMap["effectiveParameters"] = peVar.EffectiveParameters
	}
	if peVar.IsCompliant != nil {
		objectMap["isCompliant"] = peVar.IsCompliant
	}
	if peVar.SubscriptionID != nil {
		objectMap["subscriptionId"] = peVar.SubscriptionID
	}
	if peVar.ResourceType != nil {
		objectMap["resourceType"] = peVar.ResourceType
	}
	if peVar.ResourceLocation != nil {
		objectMap["resourceLocation"] = peVar.ResourceLocation
	}
	if peVar.ResourceGroup != nil {
		objectMap["resourceGroup"] = peVar.ResourceGroup
	}
	if peVar.ResourceTags != nil {
		objectMap["resourceTags"] = peVar.ResourceTags
	}
	if peVar.PolicyAssignmentName != nil {
		objectMap["policyAssignmentName"] = peVar.PolicyAssignmentName
	}
	if peVar.PolicyAssignmentOwner != nil {
		objectMap["policyAssignmentOwner"] = peVar.PolicyAssignmentOwner
	}
	if peVar.PolicyAssignmentParameters != nil {
		objectMap["policyAssignmentParameters"] = peVar.PolicyAssignmentParameters
	}
	if peVar.PolicyAssignmentScope != nil {
		objectMap["policyAssignmentScope"] = peVar.PolicyAssignmentScope
	}
	if peVar.PolicyDefinitionName != nil {
		objectMap["policyDefinitionName"] = peVar.PolicyDefinitionName
	}
	if peVar.PolicyDefinitionAction != nil {
		objectMap["policyDefinitionAction"] = peVar.PolicyDefinitionAction
	}
	if peVar.PolicyDefinitionCategory != nil {
		objectMap["policyDefinitionCategory"] = peVar.PolicyDefinitionCategory
	}
	if peVar.PolicySetDefinitionID != nil {
		objectMap["policySetDefinitionId"] = peVar.PolicySetDefinitionID
	}
	if peVar.PolicySetDefinitionName != nil {
		objectMap["policySetDefinitionName"] = peVar.PolicySetDefinitionName
	}
	if peVar.PolicySetDefinitionOwner != nil {
		objectMap["policySetDefinitionOwner"] = peVar.PolicySetDefinitionOwner
	}
	if peVar.PolicySetDefinitionCategory != nil {
		objectMap["policySetDefinitionCategory"] = peVar.PolicySetDefinitionCategory
	}
	if peVar.PolicySetDefinitionParameters != nil {
		objectMap["policySetDefinitionParameters"] = peVar.PolicySetDefinitionParameters
	}
	if peVar.ManagementGroupIds != nil {
		objectMap["managementGroupIds"] = peVar.ManagementGroupIds
	}
	if peVar.PolicyDefinitionReferenceID != nil {
		objectMap["policyDefinitionReferenceId"] = peVar.PolicyDefinitionReferenceID
	}
	if peVar.TenantID != nil {
		objectMap["tenantId"] = peVar.TenantID
	}
	if peVar.PrincipalOid != nil {
		objectMap["principalOid"] = peVar.PrincipalOid
	}
	for k, v := range peVar.AdditionalProperties {
		objectMap[k] = v
	}
	return json.Marshal(objectMap)
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-10-17-preview/policyinsights instead.
// PolicyEventsQueryResults query results.
type PolicyEventsQueryResults struct {
	autorest.Response `json:"-"`
	// OdataContext - OData context string; used by OData clients to resolve type information based on metadata.
	OdataContext *string `json:"@odata.context,omitempty"`
	// OdataCount - OData entity count; represents the number of policy event records returned.
	OdataCount *int32 `json:"@odata.count,omitempty"`
	// Value - Query results.
	Value *[]PolicyEvent `json:"value,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-10-17-preview/policyinsights instead.
// PolicyState policy state record.
type PolicyState struct {
	// AdditionalProperties - Unmatched properties from the message are deserialized this collection
	AdditionalProperties map[string]interface{} `json:""`
	// OdataID - OData entity ID; always set to null since policy state records do not have an entity ID.
	OdataID *string `json:"@odata.id,omitempty"`
	// OdataContext - OData context string; used by OData clients to resolve type information based on metadata.
	OdataContext *string `json:"@odata.context,omitempty"`
	// Timestamp - Timestamp for the policy state record.
	Timestamp *date.Time `json:"timestamp,omitempty"`
	// ResourceID - Resource ID.
	ResourceID *string `json:"resourceId,omitempty"`
	// PolicyAssignmentID - Policy assignment ID.
	PolicyAssignmentID *string `json:"policyAssignmentId,omitempty"`
	// PolicyDefinitionID - Policy definition ID.
	PolicyDefinitionID *string `json:"policyDefinitionId,omitempty"`
	// EffectiveParameters - Effective parameters for the policy assignment.
	EffectiveParameters *string `json:"effectiveParameters,omitempty"`
	// IsCompliant - Flag which states whether the resource is compliant against the policy assignment it was evaluated against.
	IsCompliant *bool `json:"isCompliant,omitempty"`
	// SubscriptionID - Subscription ID.
	SubscriptionID *string `json:"subscriptionId,omitempty"`
	// ResourceType - Resource type.
	ResourceType *string `json:"resourceType,omitempty"`
	// ResourceLocation - Resource location.
	ResourceLocation *string `json:"resourceLocation,omitempty"`
	// ResourceGroup - Resource group name.
	ResourceGroup *string `json:"resourceGroup,omitempty"`
	// ResourceTags - List of resource tags.
	ResourceTags *string `json:"resourceTags,omitempty"`
	// PolicyAssignmentName - Policy assignment name.
	PolicyAssignmentName *string `json:"policyAssignmentName,omitempty"`
	// PolicyAssignmentOwner - Policy assignment owner.
	PolicyAssignmentOwner *string `json:"policyAssignmentOwner,omitempty"`
	// PolicyAssignmentParameters - Policy assignment parameters.
	PolicyAssignmentParameters *string `json:"policyAssignmentParameters,omitempty"`
	// PolicyAssignmentScope - Policy assignment scope.
	PolicyAssignmentScope *string `json:"policyAssignmentScope,omitempty"`
	// PolicyDefinitionName - Policy definition name.
	PolicyDefinitionName *string `json:"policyDefinitionName,omitempty"`
	// PolicyDefinitionAction - Policy definition action, i.e. effect.
	PolicyDefinitionAction *string `json:"policyDefinitionAction,omitempty"`
	// PolicyDefinitionCategory - Policy definition category.
	PolicyDefinitionCategory *string `json:"policyDefinitionCategory,omitempty"`
	// PolicySetDefinitionID - Policy set definition ID, if the policy assignment is for a policy set.
	PolicySetDefinitionID *string `json:"policySetDefinitionId,omitempty"`
	// PolicySetDefinitionName - Policy set definition name, if the policy assignment is for a policy set.
	PolicySetDefinitionName *string `json:"policySetDefinitionName,omitempty"`
	// PolicySetDefinitionOwner - Policy set definition owner, if the policy assignment is for a policy set.
	PolicySetDefinitionOwner *string `json:"policySetDefinitionOwner,omitempty"`
	// PolicySetDefinitionCategory - Policy set definition category, if the policy assignment is for a policy set.
	PolicySetDefinitionCategory *string `json:"policySetDefinitionCategory,omitempty"`
	// PolicySetDefinitionParameters - Policy set definition parameters, if the policy assignment is for a policy set.
	PolicySetDefinitionParameters *string `json:"policySetDefinitionParameters,omitempty"`
	// ManagementGroupIds - Comma seperated list of management group IDs, which represent the hierarchy of the management groups the resource is under.
	ManagementGroupIds *string `json:"managementGroupIds,omitempty"`
	// PolicyDefinitionReferenceID - Reference ID for the policy definition inside the policy set, if the policy assignment is for a policy set.
	PolicyDefinitionReferenceID *string `json:"policyDefinitionReferenceId,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-10-17-preview/policyinsights instead.
// MarshalJSON is the custom marshaler for PolicyState.
func (ps PolicyState) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ps.OdataID != nil {
		objectMap["@odata.id"] = ps.OdataID
	}
	if ps.OdataContext != nil {
		objectMap["@odata.context"] = ps.OdataContext
	}
	if ps.Timestamp != nil {
		objectMap["timestamp"] = ps.Timestamp
	}
	if ps.ResourceID != nil {
		objectMap["resourceId"] = ps.ResourceID
	}
	if ps.PolicyAssignmentID != nil {
		objectMap["policyAssignmentId"] = ps.PolicyAssignmentID
	}
	if ps.PolicyDefinitionID != nil {
		objectMap["policyDefinitionId"] = ps.PolicyDefinitionID
	}
	if ps.EffectiveParameters != nil {
		objectMap["effectiveParameters"] = ps.EffectiveParameters
	}
	if ps.IsCompliant != nil {
		objectMap["isCompliant"] = ps.IsCompliant
	}
	if ps.SubscriptionID != nil {
		objectMap["subscriptionId"] = ps.SubscriptionID
	}
	if ps.ResourceType != nil {
		objectMap["resourceType"] = ps.ResourceType
	}
	if ps.ResourceLocation != nil {
		objectMap["resourceLocation"] = ps.ResourceLocation
	}
	if ps.ResourceGroup != nil {
		objectMap["resourceGroup"] = ps.ResourceGroup
	}
	if ps.ResourceTags != nil {
		objectMap["resourceTags"] = ps.ResourceTags
	}
	if ps.PolicyAssignmentName != nil {
		objectMap["policyAssignmentName"] = ps.PolicyAssignmentName
	}
	if ps.PolicyAssignmentOwner != nil {
		objectMap["policyAssignmentOwner"] = ps.PolicyAssignmentOwner
	}
	if ps.PolicyAssignmentParameters != nil {
		objectMap["policyAssignmentParameters"] = ps.PolicyAssignmentParameters
	}
	if ps.PolicyAssignmentScope != nil {
		objectMap["policyAssignmentScope"] = ps.PolicyAssignmentScope
	}
	if ps.PolicyDefinitionName != nil {
		objectMap["policyDefinitionName"] = ps.PolicyDefinitionName
	}
	if ps.PolicyDefinitionAction != nil {
		objectMap["policyDefinitionAction"] = ps.PolicyDefinitionAction
	}
	if ps.PolicyDefinitionCategory != nil {
		objectMap["policyDefinitionCategory"] = ps.PolicyDefinitionCategory
	}
	if ps.PolicySetDefinitionID != nil {
		objectMap["policySetDefinitionId"] = ps.PolicySetDefinitionID
	}
	if ps.PolicySetDefinitionName != nil {
		objectMap["policySetDefinitionName"] = ps.PolicySetDefinitionName
	}
	if ps.PolicySetDefinitionOwner != nil {
		objectMap["policySetDefinitionOwner"] = ps.PolicySetDefinitionOwner
	}
	if ps.PolicySetDefinitionCategory != nil {
		objectMap["policySetDefinitionCategory"] = ps.PolicySetDefinitionCategory
	}
	if ps.PolicySetDefinitionParameters != nil {
		objectMap["policySetDefinitionParameters"] = ps.PolicySetDefinitionParameters
	}
	if ps.ManagementGroupIds != nil {
		objectMap["managementGroupIds"] = ps.ManagementGroupIds
	}
	if ps.PolicyDefinitionReferenceID != nil {
		objectMap["policyDefinitionReferenceId"] = ps.PolicyDefinitionReferenceID
	}
	for k, v := range ps.AdditionalProperties {
		objectMap[k] = v
	}
	return json.Marshal(objectMap)
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-10-17-preview/policyinsights instead.
// PolicyStatesQueryResults query results.
type PolicyStatesQueryResults struct {
	autorest.Response `json:"-"`
	// OdataContext - OData context string; used by OData clients to resolve type information based on metadata.
	OdataContext *string `json:"@odata.context,omitempty"`
	// OdataCount - OData entity count; represents the number of policy state records returned.
	OdataCount *int32 `json:"@odata.count,omitempty"`
	// Value - Query results.
	Value *[]PolicyState `json:"value,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-10-17-preview/policyinsights instead.
// QueryFailure error response.
type QueryFailure struct {
	// Error - Error definition.
	Error *QueryFailureError `json:"error,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-10-17-preview/policyinsights instead.
// QueryFailureError error definition.
type QueryFailureError struct {
	// Code - Service specific error code which serves as the substatus for the HTTP error code.
	Code *string `json:"code,omitempty"`
	// Message - Description of the error.
	Message *string `json:"message,omitempty"`
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-10-17-preview/policyinsights instead.
// String ...
type String struct {
	autorest.Response `json:"-"`
	Value             *string `json:"value,omitempty"`
}
