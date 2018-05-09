package operationsmanagement

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
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
)

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// ManagementAssociationsClient is the operations Management Client
type ManagementAssociationsClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// NewManagementAssociationsClient creates an instance of the ManagementAssociationsClient client.
func NewManagementAssociationsClient(subscriptionID string, providerName string, resourceType string, resourceName string) ManagementAssociationsClient {
	return NewManagementAssociationsClientWithBaseURI(DefaultBaseURI, subscriptionID, providerName, resourceType, resourceName)
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// NewManagementAssociationsClientWithBaseURI creates an instance of the ManagementAssociationsClient client.
func NewManagementAssociationsClientWithBaseURI(baseURI string, subscriptionID string, providerName string, resourceType string, resourceName string) ManagementAssociationsClient {
	return ManagementAssociationsClient{NewWithBaseURI(baseURI, subscriptionID, providerName, resourceType, resourceName)}
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// CreateOrUpdate creates or updates the ManagementAssociation.
//
// resourceGroupName is the name of the resource group to get. The name is case insensitive.
// managementAssociationName is user ManagementAssociation Name. parameters is the parameters required to create
// ManagementAssociation extension.
func (client ManagementAssociationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, managementAssociationName string, parameters ManagementAssociation) (result ManagementAssociation, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Properties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.Properties.ApplicationID", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("operationsmanagement.ManagementAssociationsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, managementAssociationName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationsmanagement.ManagementAssociationsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "operationsmanagement.ManagementAssociationsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationsmanagement.ManagementAssociationsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ManagementAssociationsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, managementAssociationName string, parameters ManagementAssociation) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementAssociationName": autorest.Encode("path", managementAssociationName),
		"providerName":              autorest.Encode("path", client.ProviderName),
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"resourceName":              autorest.Encode("path", client.ResourceName),
		"resourceType":              autorest.Encode("path", client.ResourceType),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.OperationsManagement/ManagementAssociations/{managementAssociationName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementAssociationsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ManagementAssociationsClient) CreateOrUpdateResponder(resp *http.Response) (result ManagementAssociation, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// Delete deletes the ManagementAssociation in the subscription.
//
// resourceGroupName is the name of the resource group to get. The name is case insensitive.
// managementAssociationName is user ManagementAssociation Name.
func (client ManagementAssociationsClient) Delete(ctx context.Context, resourceGroupName string, managementAssociationName string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("operationsmanagement.ManagementAssociationsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, managementAssociationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationsmanagement.ManagementAssociationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "operationsmanagement.ManagementAssociationsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationsmanagement.ManagementAssociationsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// DeletePreparer prepares the Delete request.
func (client ManagementAssociationsClient) DeletePreparer(ctx context.Context, resourceGroupName string, managementAssociationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementAssociationName": autorest.Encode("path", managementAssociationName),
		"providerName":              autorest.Encode("path", client.ProviderName),
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"resourceName":              autorest.Encode("path", client.ResourceName),
		"resourceType":              autorest.Encode("path", client.ResourceType),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.OperationsManagement/ManagementAssociations/{managementAssociationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementAssociationsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ManagementAssociationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// Get retrieves the user ManagementAssociation.
//
// resourceGroupName is the name of the resource group to get. The name is case insensitive.
// managementAssociationName is user ManagementAssociation Name.
func (client ManagementAssociationsClient) Get(ctx context.Context, resourceGroupName string, managementAssociationName string) (result ManagementAssociation, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("operationsmanagement.ManagementAssociationsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, managementAssociationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationsmanagement.ManagementAssociationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "operationsmanagement.ManagementAssociationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationsmanagement.ManagementAssociationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// GetPreparer prepares the Get request.
func (client ManagementAssociationsClient) GetPreparer(ctx context.Context, resourceGroupName string, managementAssociationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementAssociationName": autorest.Encode("path", managementAssociationName),
		"providerName":              autorest.Encode("path", client.ProviderName),
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"resourceName":              autorest.Encode("path", client.ResourceName),
		"resourceType":              autorest.Encode("path", client.ResourceType),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.OperationsManagement/ManagementAssociations/{managementAssociationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementAssociationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ManagementAssociationsClient) GetResponder(resp *http.Response) (result ManagementAssociation, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// ListBySubscription retrieves the ManagementAssociatons list.
func (client ManagementAssociationsClient) ListBySubscription(ctx context.Context) (result ManagementAssociationPropertiesList, err error) {
	req, err := client.ListBySubscriptionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationsmanagement.ManagementAssociationsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "operationsmanagement.ManagementAssociationsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationsmanagement.ManagementAssociationsClient", "ListBySubscription", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client ManagementAssociationsClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.OperationsManagement/ManagementAssociations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementAssociationsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/operationsmanagement/mgmt/2015-11-01-preview/operationsmanagement instead.
// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client ManagementAssociationsClient) ListBySubscriptionResponder(resp *http.Response) (result ManagementAssociationPropertiesList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
