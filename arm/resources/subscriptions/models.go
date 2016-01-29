package subscriptions

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.12.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// DeploymentExtendedFilter is deployment filter.
type DeploymentExtendedFilter struct {
	ProvisioningState *string `json:"provisioningState,omitempty"`
}

// GenericResourceFilter is resource filter.
type GenericResourceFilter struct {
	ResourceType *string `json:"resourceType,omitempty"`
	Tagname      *string `json:"tagname,omitempty"`
	Tagvalue     *string `json:"tagvalue,omitempty"`
}

// Location is location information.
type Location struct {
	ID             *string `json:"id,omitempty"`
	SubscriptionID *string `json:"subscriptionId,omitempty"`
	Name           *string `json:"name,omitempty"`
	DisplayName    *string `json:"displayName,omitempty"`
	Latitude       *string `json:"latitude,omitempty"`
	Longitude      *string `json:"longitude,omitempty"`
}

// LocationListResult is location list operation response.
type LocationListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Location `json:"value,omitempty"`
	NextLink          *string     `json:",omitempty"`
}

// LocationListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client LocationListResult) LocationListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// Resource is
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// ResourceGroupFilter is resource group filter.
type ResourceGroupFilter struct {
	TagName  *string `json:"tagName,omitempty"`
	TagValue *string `json:"tagValue,omitempty"`
}

// SubResource is
type SubResource struct {
	ID *string `json:"id,omitempty"`
}

// Subscription is subscription information.
type Subscription struct {
	autorest.Response `json:"-"`
	ID                *string `json:"id,omitempty"`
	SubscriptionID    *string `json:"subscriptionId,omitempty"`
	DisplayName       *string `json:"displayName,omitempty"`
	State             *string `json:"state,omitempty"`
}

// SubscriptionListResult is subscription list operation response.
type SubscriptionListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Subscription `json:"value,omitempty"`
	NextLink          *string         `json:"nextLink,omitempty"`
}

// SubscriptionListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client SubscriptionListResult) SubscriptionListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// TenantIDDescription is tenant Id information
type TenantIDDescription struct {
	ID       *string `json:"id,omitempty"`
	TenantID *string `json:"tenantId,omitempty"`
}

// TenantListResult is tenant Ids information.
type TenantListResult struct {
	autorest.Response `json:"-"`
	Value             *[]TenantIDDescription `json:"value,omitempty"`
	NextLink          *string                `json:"nextLink,omitempty"`
}

// TenantListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client TenantListResult) TenantListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}