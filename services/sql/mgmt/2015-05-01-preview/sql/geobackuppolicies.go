package sql

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

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// GeoBackupPoliciesClient is the the Azure SQL Database management API provides a RESTful set of web services that
// interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update,
// and delete databases.
type GeoBackupPoliciesClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// NewGeoBackupPoliciesClient creates an instance of the GeoBackupPoliciesClient client.
func NewGeoBackupPoliciesClient(subscriptionID string) GeoBackupPoliciesClient {
	return NewGeoBackupPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// NewGeoBackupPoliciesClientWithBaseURI creates an instance of the GeoBackupPoliciesClient client.
func NewGeoBackupPoliciesClientWithBaseURI(baseURI string, subscriptionID string) GeoBackupPoliciesClient {
	return GeoBackupPoliciesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// CreateOrUpdate updates a database geo backup policy.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal. serverName is the name of the server. databaseName is the name of
// the database. parameters is the required parameters for creating or updating the geo backup policy.
func (client GeoBackupPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters GeoBackupPolicy) (result GeoBackupPolicy, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.GeoBackupPolicyProperties", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("sql.GeoBackupPoliciesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, databaseName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.GeoBackupPoliciesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.GeoBackupPoliciesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.GeoBackupPoliciesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client GeoBackupPoliciesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters GeoBackupPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":        autorest.Encode("path", databaseName),
		"geoBackupPolicyName": autorest.Encode("path", "Default"),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"serverName":          autorest.Encode("path", serverName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/geoBackupPolicies/{geoBackupPolicyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client GeoBackupPoliciesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client GeoBackupPoliciesClient) CreateOrUpdateResponder(resp *http.Response) (result GeoBackupPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// Get gets a geo backup policy.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal. serverName is the name of the server. databaseName is the name of
// the database.
func (client GeoBackupPoliciesClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result GeoBackupPolicy, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, databaseName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.GeoBackupPoliciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.GeoBackupPoliciesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.GeoBackupPoliciesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// GetPreparer prepares the Get request.
func (client GeoBackupPoliciesClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":        autorest.Encode("path", databaseName),
		"geoBackupPolicyName": autorest.Encode("path", "Default"),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"serverName":          autorest.Encode("path", serverName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/geoBackupPolicies/{geoBackupPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client GeoBackupPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client GeoBackupPoliciesClient) GetResponder(resp *http.Response) (result GeoBackupPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// ListByDatabase returns a list of geo backup policies.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal. serverName is the name of the server. databaseName is the name of
// the database.
func (client GeoBackupPoliciesClient) ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result GeoBackupPolicyListResult, err error) {
	req, err := client.ListByDatabasePreparer(ctx, resourceGroupName, serverName, databaseName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.GeoBackupPoliciesClient", "ListByDatabase", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDatabaseSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.GeoBackupPoliciesClient", "ListByDatabase", resp, "Failure sending request")
		return
	}

	result, err = client.ListByDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.GeoBackupPoliciesClient", "ListByDatabase", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// ListByDatabasePreparer prepares the ListByDatabase request.
func (client GeoBackupPoliciesClient) ListByDatabasePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/geoBackupPolicies", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// ListByDatabaseSender sends the ListByDatabase request. The method will close the
// http.Response Body if it receives an error.
func (client GeoBackupPoliciesClient) ListByDatabaseSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// ListByDatabaseResponder handles the response to the ListByDatabase request. The method always
// closes the http.Response Body.
func (client GeoBackupPoliciesClient) ListByDatabaseResponder(resp *http.Response) (result GeoBackupPolicyListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
