package batchai

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

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// ClustersClient is the the Azure BatchAI Management API.
type ClustersClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// NewClustersClient creates an instance of the ClustersClient client.
func NewClustersClient(subscriptionID string) ClustersClient {
	return NewClustersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// NewClustersClientWithBaseURI creates an instance of the ClustersClient client.
func NewClustersClientWithBaseURI(baseURI string, subscriptionID string) ClustersClient {
	return ClustersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// Create adds a cluster. A cluster is a collection of compute nodes. Multiple jobs can be run on the same cluster.
//
// resourceGroupName is name of the resource group to which the resource belongs. clusterName is the name of the
// cluster within the specified resource group. Cluster names can only contain a combination of alphanumeric
// characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
// parameters is the parameters to provide for cluster creation.
func (client ClustersClient) Create(ctx context.Context, resourceGroupName string, clusterName string, parameters ClusterCreateParameters) (result ClustersCreateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: clusterName,
			Constraints: []validation.Constraint{{Target: "clusterName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "clusterName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "clusterName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Location", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.ClusterBaseProperties", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.VMSize", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.ClusterBaseProperties.ScaleSettings", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.ScaleSettings.Manual", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.ScaleSettings.Manual.TargetNodeCount", Name: validation.Null, Rule: true, Chain: nil}}},
								{Target: "parameters.ClusterBaseProperties.ScaleSettings.AutoScale", Name: validation.Null, Rule: false,
									Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.ScaleSettings.AutoScale.MinimumNodeCount", Name: validation.Null, Rule: true, Chain: nil},
										{Target: "parameters.ClusterBaseProperties.ScaleSettings.AutoScale.MaximumNodeCount", Name: validation.Null, Rule: true, Chain: nil},
									}},
							}},
						{Target: "parameters.ClusterBaseProperties.VirtualMachineConfiguration", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.VirtualMachineConfiguration.ImageReference", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.VirtualMachineConfiguration.ImageReference.Publisher", Name: validation.Null, Rule: true, Chain: nil},
									{Target: "parameters.ClusterBaseProperties.VirtualMachineConfiguration.ImageReference.Offer", Name: validation.Null, Rule: true, Chain: nil},
									{Target: "parameters.ClusterBaseProperties.VirtualMachineConfiguration.ImageReference.Sku", Name: validation.Null, Rule: true, Chain: nil},
								}},
							}},
						{Target: "parameters.ClusterBaseProperties.NodeSetup", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.NodeSetup.SetupTask", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.NodeSetup.SetupTask.CommandLine", Name: validation.Null, Rule: true, Chain: nil},
									{Target: "parameters.ClusterBaseProperties.NodeSetup.SetupTask.StdOutErrPathPrefix", Name: validation.Null, Rule: true, Chain: nil},
								}},
							}},
						{Target: "parameters.ClusterBaseProperties.UserAccountSettings", Name: validation.Null, Rule: true,
							Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.UserAccountSettings.AdminUserName", Name: validation.Null, Rule: true, Chain: nil}}},
						{Target: "parameters.ClusterBaseProperties.Subnet", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.ClusterBaseProperties.Subnet.ID", Name: validation.Null, Rule: true, Chain: nil}}},
					}}}}}); err != nil {
		return result, validation.NewError("batchai.ClustersClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, clusterName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// CreatePreparer prepares the Create request.
func (client ClustersClient) CreatePreparer(ctx context.Context, resourceGroupName string, clusterName string, parameters ClusterCreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/clusters/{clusterName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) CreateSender(req *http.Request) (future ClustersCreateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ClustersClient) CreateResponder(resp *http.Response) (result Cluster, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// Delete deletes a Cluster.
//
// resourceGroupName is name of the resource group to which the resource belongs. clusterName is the name of the
// cluster within the specified resource group. Cluster names can only contain a combination of alphanumeric
// characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
func (client ClustersClient) Delete(ctx context.Context, resourceGroupName string, clusterName string) (result ClustersDeleteFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: clusterName,
			Constraints: []validation.Constraint{{Target: "clusterName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "clusterName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "clusterName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batchai.ClustersClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// DeletePreparer prepares the Delete request.
func (client ClustersClient) DeletePreparer(ctx context.Context, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/clusters/{clusterName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) DeleteSender(req *http.Request) (future ClustersDeleteFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ClustersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// Get gets information about the specified Cluster.
//
// resourceGroupName is name of the resource group to which the resource belongs. clusterName is the name of the
// cluster within the specified resource group. Cluster names can only contain a combination of alphanumeric
// characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
func (client ClustersClient) Get(ctx context.Context, resourceGroupName string, clusterName string) (result Cluster, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: clusterName,
			Constraints: []validation.Constraint{{Target: "clusterName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "clusterName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "clusterName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batchai.ClustersClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Get", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// GetPreparer prepares the Get request.
func (client ClustersClient) GetPreparer(ctx context.Context, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/clusters/{clusterName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ClustersClient) GetResponder(resp *http.Response) (result Cluster, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// List gets information about the Clusters associated with the subscription.
//
// filter is an OData $filter clause.. Used to filter results that are returned in the GET respnose.
// selectParameter is an OData $select clause. Used to select the properties to be returned in the GET respnose.
// maxResults is the maximum number of items to return in the response. A maximum of 1000 files can be returned.
func (client ClustersClient) List(ctx context.Context, filter string, selectParameter string, maxResults *int32) (result ClusterListResultPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: maxResults,
			Constraints: []validation.Constraint{{Target: "maxResults", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "maxResults", Name: validation.InclusiveMaximum, Rule: 1000, Chain: nil},
					{Target: "maxResults", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("batchai.ClustersClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, filter, selectParameter, maxResults)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.clr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "List", resp, "Failure sending request")
		return
	}

	result.clr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "List", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// ListPreparer prepares the List request.
func (client ClustersClient) ListPreparer(ctx context.Context, filter string, selectParameter string, maxResults *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if maxResults != nil {
		queryParameters["maxresults"] = autorest.Encode("query", *maxResults)
	} else {
		queryParameters["maxresults"] = autorest.Encode("query", 1000)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.BatchAI/clusters", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ClustersClient) ListResponder(resp *http.Response) (result ClusterListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ClustersClient) listNextResults(lastResults ClusterListResult) (result ClusterListResult, err error) {
	req, err := lastResults.clusterListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batchai.ClustersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batchai.ClustersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ClustersClient) ListComplete(ctx context.Context, filter string, selectParameter string, maxResults *int32) (result ClusterListResultIterator, err error) {
	result.page, err = client.List(ctx, filter, selectParameter, maxResults)
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// ListByResourceGroup gets information about the Clusters associated within the specified resource group.
//
// resourceGroupName is name of the resource group to which the resource belongs. filter is an OData $filter
// clause.. Used to filter results that are returned in the GET respnose. selectParameter is an OData $select
// clause. Used to select the properties to be returned in the GET respnose. maxResults is the maximum number of
// items to return in the response. A maximum of 1000 files can be returned.
func (client ClustersClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, selectParameter string, maxResults *int32) (result ClusterListResultPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: maxResults,
			Constraints: []validation.Constraint{{Target: "maxResults", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "maxResults", Name: validation.InclusiveMaximum, Rule: 1000, Chain: nil},
					{Target: "maxResults", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("batchai.ClustersClient", "ListByResourceGroup", err.Error())
	}

	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, filter, selectParameter, maxResults)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.clr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.clr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ClustersClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, filter string, selectParameter string, maxResults *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if maxResults != nil {
		queryParameters["maxresults"] = autorest.Encode("query", *maxResults)
	} else {
		queryParameters["maxresults"] = autorest.Encode("query", 1000)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/clusters", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ClustersClient) ListByResourceGroupResponder(resp *http.Response) (result ClusterListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client ClustersClient) listByResourceGroupNextResults(lastResults ClusterListResult) (result ClusterListResult, err error) {
	req, err := lastResults.clusterListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batchai.ClustersClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batchai.ClustersClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client ClustersClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, selectParameter string, maxResults *int32) (result ClusterListResultIterator, err error) {
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, filter, selectParameter, maxResults)
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// ListRemoteLoginInformation get the IP address, port of all the compute nodes in the cluster.
//
// resourceGroupName is name of the resource group to which the resource belongs. clusterName is the name of the
// cluster within the specified resource group. Cluster names can only contain a combination of alphanumeric
// characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
func (client ClustersClient) ListRemoteLoginInformation(ctx context.Context, resourceGroupName string, clusterName string) (result RemoteLoginInformationListResultPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: clusterName,
			Constraints: []validation.Constraint{{Target: "clusterName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "clusterName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "clusterName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batchai.ClustersClient", "ListRemoteLoginInformation", err.Error())
	}

	result.fn = client.listRemoteLoginInformationNextResults
	req, err := client.ListRemoteLoginInformationPreparer(ctx, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "ListRemoteLoginInformation", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListRemoteLoginInformationSender(req)
	if err != nil {
		result.rlilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "ListRemoteLoginInformation", resp, "Failure sending request")
		return
	}

	result.rlilr, err = client.ListRemoteLoginInformationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "ListRemoteLoginInformation", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// ListRemoteLoginInformationPreparer prepares the ListRemoteLoginInformation request.
func (client ClustersClient) ListRemoteLoginInformationPreparer(ctx context.Context, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/clusters/{clusterName}/listRemoteLoginInformation", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// ListRemoteLoginInformationSender sends the ListRemoteLoginInformation request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) ListRemoteLoginInformationSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// ListRemoteLoginInformationResponder handles the response to the ListRemoteLoginInformation request. The method always
// closes the http.Response Body.
func (client ClustersClient) ListRemoteLoginInformationResponder(resp *http.Response) (result RemoteLoginInformationListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listRemoteLoginInformationNextResults retrieves the next set of results, if any.
func (client ClustersClient) listRemoteLoginInformationNextResults(lastResults RemoteLoginInformationListResult) (result RemoteLoginInformationListResult, err error) {
	req, err := lastResults.remoteLoginInformationListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batchai.ClustersClient", "listRemoteLoginInformationNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListRemoteLoginInformationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batchai.ClustersClient", "listRemoteLoginInformationNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListRemoteLoginInformationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "listRemoteLoginInformationNextResults", resp, "Failure responding to next results request")
	}
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// ListRemoteLoginInformationComplete enumerates all values, automatically crossing page boundaries as required.
func (client ClustersClient) ListRemoteLoginInformationComplete(ctx context.Context, resourceGroupName string, clusterName string) (result RemoteLoginInformationListResultIterator, err error) {
	result.page, err = client.ListRemoteLoginInformation(ctx, resourceGroupName, clusterName)
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// Update update the properties of a given cluster.
//
// resourceGroupName is name of the resource group to which the resource belongs. clusterName is the name of the
// cluster within the specified resource group. Cluster names can only contain a combination of alphanumeric
// characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
// parameters is additional parameters for cluster update.
func (client ClustersClient) Update(ctx context.Context, resourceGroupName string, clusterName string, parameters ClusterUpdateParameters) (result Cluster, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: clusterName,
			Constraints: []validation.Constraint{{Target: "clusterName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "clusterName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "clusterName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batchai.ClustersClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, clusterName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchai.ClustersClient", "Update", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// UpdatePreparer prepares the Update request.
func (client ClustersClient) UpdatePreparer(ctx context.Context, resourceGroupName string, clusterName string, parameters ClusterUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/clusters/{clusterName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ClustersClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/batchai/mgmt/2017-09-preview/batchai instead.
// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ClustersClient) UpdateResponder(resp *http.Response) (result Cluster, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
