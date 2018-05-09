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
)

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// ServerConnectionPoliciesClient is the the Azure SQL Database management API provides a RESTful set of web services
// that interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve,
// update, and delete databases.
type ServerConnectionPoliciesClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// NewServerConnectionPoliciesClient creates an instance of the ServerConnectionPoliciesClient client.
func NewServerConnectionPoliciesClient(subscriptionID string) ServerConnectionPoliciesClient {
	return NewServerConnectionPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// NewServerConnectionPoliciesClientWithBaseURI creates an instance of the ServerConnectionPoliciesClient client.
func NewServerConnectionPoliciesClientWithBaseURI(baseURI string, subscriptionID string) ServerConnectionPoliciesClient {
	return ServerConnectionPoliciesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// CreateOrUpdate creates or updates the server's connection policy.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal. serverName is the name of the server. parameters is the required
// parameters for updating a secure connection policy.
func (client ServerConnectionPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, parameters ServerConnectionPolicy) (result ServerConnectionPolicy, err error) {
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerConnectionPoliciesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ServerConnectionPoliciesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerConnectionPoliciesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ServerConnectionPoliciesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, parameters ServerConnectionPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionPolicyName": autorest.Encode("path", "default"),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"serverName":           autorest.Encode("path", serverName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/connectionPolicies/{connectionPolicyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ServerConnectionPoliciesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ServerConnectionPoliciesClient) CreateOrUpdateResponder(resp *http.Response) (result ServerConnectionPolicy, err error) {
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
// Get gets the server's secure connection policy.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal. serverName is the name of the server.
func (client ServerConnectionPoliciesClient) Get(ctx context.Context, resourceGroupName string, serverName string) (result ServerConnectionPolicy, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerConnectionPoliciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ServerConnectionPoliciesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ServerConnectionPoliciesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// GetPreparer prepares the Get request.
func (client ServerConnectionPoliciesClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionPolicyName": autorest.Encode("path", "default"),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"serverName":           autorest.Encode("path", serverName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/connectionPolicies/{connectionPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ServerConnectionPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ServerConnectionPoliciesClient) GetResponder(resp *http.Response) (result ServerConnectionPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
