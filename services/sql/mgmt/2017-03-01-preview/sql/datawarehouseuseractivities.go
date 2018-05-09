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

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2017-03-01-preview/sql instead.
// DataWarehouseUserActivitiesClient is the the Azure SQL Database management API provides a RESTful set of web
// services that interact with Azure SQL Database services to manage your databases. The API enables you to create,
// retrieve, update, and delete databases.
type DataWarehouseUserActivitiesClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2017-03-01-preview/sql instead.
// NewDataWarehouseUserActivitiesClient creates an instance of the DataWarehouseUserActivitiesClient client.
func NewDataWarehouseUserActivitiesClient(subscriptionID string) DataWarehouseUserActivitiesClient {
	return NewDataWarehouseUserActivitiesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2017-03-01-preview/sql instead.
// NewDataWarehouseUserActivitiesClientWithBaseURI creates an instance of the DataWarehouseUserActivitiesClient client.
func NewDataWarehouseUserActivitiesClientWithBaseURI(baseURI string, subscriptionID string) DataWarehouseUserActivitiesClient {
	return DataWarehouseUserActivitiesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2017-03-01-preview/sql instead.
// Get gets the user activities of a data warehouse which includes running and suspended queries
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal. serverName is the name of the server. databaseName is the name of
// the database.
func (client DataWarehouseUserActivitiesClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result DataWarehouseUserActivities, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, databaseName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DataWarehouseUserActivitiesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.DataWarehouseUserActivitiesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DataWarehouseUserActivitiesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2017-03-01-preview/sql instead.
// GetPreparer prepares the Get request.
func (client DataWarehouseUserActivitiesClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":                  autorest.Encode("path", databaseName),
		"dataWarehouseUserActivityName": autorest.Encode("path", "current"),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"serverName":                    autorest.Encode("path", serverName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/dataWarehouseUserActivities/{dataWarehouseUserActivityName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2017-03-01-preview/sql instead.
// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DataWarehouseUserActivitiesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/sql/mgmt/2017-03-01-preview/sql instead.
// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DataWarehouseUserActivitiesClient) GetResponder(resp *http.Response) (result DataWarehouseUserActivities, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
