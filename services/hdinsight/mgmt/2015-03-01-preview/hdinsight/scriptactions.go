package hdinsight

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

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hdinsight/mgmt/2015-03-01-preview/hdinsight instead.
// ScriptActionsClient is the hDInsight Management Client
type ScriptActionsClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hdinsight/mgmt/2015-03-01-preview/hdinsight instead.
// NewScriptActionsClient creates an instance of the ScriptActionsClient client.
func NewScriptActionsClient(subscriptionID string) ScriptActionsClient {
	return NewScriptActionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hdinsight/mgmt/2015-03-01-preview/hdinsight instead.
// NewScriptActionsClientWithBaseURI creates an instance of the ScriptActionsClient client.
func NewScriptActionsClientWithBaseURI(baseURI string, subscriptionID string) ScriptActionsClient {
	return ScriptActionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hdinsight/mgmt/2015-03-01-preview/hdinsight instead.
// Delete deletes a specified persisted script action of the cluster.
//
// resourceGroupName is the name of the resource group. clusterName is the name of the cluster. scriptName is the
// name of the script.
func (client ScriptActionsClient) Delete(ctx context.Context, resourceGroupName string, clusterName string, scriptName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, clusterName, scriptName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ScriptActionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "hdinsight.ScriptActionsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ScriptActionsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hdinsight/mgmt/2015-03-01-preview/hdinsight instead.
// DeletePreparer prepares the Delete request.
func (client ScriptActionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, clusterName string, scriptName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"scriptName":        autorest.Encode("path", scriptName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/scriptActions/{scriptName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hdinsight/mgmt/2015-03-01-preview/hdinsight instead.
// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ScriptActionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hdinsight/mgmt/2015-03-01-preview/hdinsight instead.
// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ScriptActionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hdinsight/mgmt/2015-03-01-preview/hdinsight instead.
// GetExecutionDetail gets the script execution detail for the given script execution ID.
//
// resourceGroupName is the name of the resource group. clusterName is the name of the cluster. scriptExecutionID
// is the script execution Id
func (client ScriptActionsClient) GetExecutionDetail(ctx context.Context, resourceGroupName string, clusterName string, scriptExecutionID string) (result RuntimeScriptActionDetail, err error) {
	req, err := client.GetExecutionDetailPreparer(ctx, resourceGroupName, clusterName, scriptExecutionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ScriptActionsClient", "GetExecutionDetail", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetExecutionDetailSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hdinsight.ScriptActionsClient", "GetExecutionDetail", resp, "Failure sending request")
		return
	}

	result, err = client.GetExecutionDetailResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ScriptActionsClient", "GetExecutionDetail", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hdinsight/mgmt/2015-03-01-preview/hdinsight instead.
// GetExecutionDetailPreparer prepares the GetExecutionDetail request.
func (client ScriptActionsClient) GetExecutionDetailPreparer(ctx context.Context, resourceGroupName string, clusterName string, scriptExecutionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"scriptExecutionId": autorest.Encode("path", scriptExecutionID),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/scriptExecutionHistory/{scriptExecutionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hdinsight/mgmt/2015-03-01-preview/hdinsight instead.
// GetExecutionDetailSender sends the GetExecutionDetail request. The method will close the
// http.Response Body if it receives an error.
func (client ScriptActionsClient) GetExecutionDetailSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hdinsight/mgmt/2015-03-01-preview/hdinsight instead.
// GetExecutionDetailResponder handles the response to the GetExecutionDetail request. The method always
// closes the http.Response Body.
func (client ScriptActionsClient) GetExecutionDetailResponder(resp *http.Response) (result RuntimeScriptActionDetail, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hdinsight/mgmt/2015-03-01-preview/hdinsight instead.
// ListPersistedScripts lists all the persisted script actions for the specified cluster.
//
// resourceGroupName is the name of the resource group. clusterName is the name of the cluster.
func (client ScriptActionsClient) ListPersistedScripts(ctx context.Context, resourceGroupName string, clusterName string) (result ScriptActionsListPage, err error) {
	result.fn = client.listPersistedScriptsNextResults
	req, err := client.ListPersistedScriptsPreparer(ctx, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ScriptActionsClient", "ListPersistedScripts", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListPersistedScriptsSender(req)
	if err != nil {
		result.sal.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hdinsight.ScriptActionsClient", "ListPersistedScripts", resp, "Failure sending request")
		return
	}

	result.sal, err = client.ListPersistedScriptsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ScriptActionsClient", "ListPersistedScripts", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hdinsight/mgmt/2015-03-01-preview/hdinsight instead.
// ListPersistedScriptsPreparer prepares the ListPersistedScripts request.
func (client ScriptActionsClient) ListPersistedScriptsPreparer(ctx context.Context, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/scriptActions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hdinsight/mgmt/2015-03-01-preview/hdinsight instead.
// ListPersistedScriptsSender sends the ListPersistedScripts request. The method will close the
// http.Response Body if it receives an error.
func (client ScriptActionsClient) ListPersistedScriptsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hdinsight/mgmt/2015-03-01-preview/hdinsight instead.
// ListPersistedScriptsResponder handles the response to the ListPersistedScripts request. The method always
// closes the http.Response Body.
func (client ScriptActionsClient) ListPersistedScriptsResponder(resp *http.Response) (result ScriptActionsList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listPersistedScriptsNextResults retrieves the next set of results, if any.
func (client ScriptActionsClient) listPersistedScriptsNextResults(lastResults ScriptActionsList) (result ScriptActionsList, err error) {
	req, err := lastResults.scriptActionsListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "hdinsight.ScriptActionsClient", "listPersistedScriptsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListPersistedScriptsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "hdinsight.ScriptActionsClient", "listPersistedScriptsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListPersistedScriptsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hdinsight.ScriptActionsClient", "listPersistedScriptsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/hdinsight/mgmt/2015-03-01-preview/hdinsight instead.
// ListPersistedScriptsComplete enumerates all values, automatically crossing page boundaries as required.
func (client ScriptActionsClient) ListPersistedScriptsComplete(ctx context.Context, resourceGroupName string, clusterName string) (result ScriptActionsListIterator, err error) {
	result.page, err = client.ListPersistedScripts(ctx, resourceGroupName, clusterName)
	return
}
