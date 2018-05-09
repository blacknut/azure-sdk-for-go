package scheduler

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

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// JobCollectionsClient is the client for the JobCollections methods of the Scheduler service.
type JobCollectionsClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// NewJobCollectionsClient creates an instance of the JobCollectionsClient client.
func NewJobCollectionsClient(subscriptionID string) JobCollectionsClient {
	return NewJobCollectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// NewJobCollectionsClientWithBaseURI creates an instance of the JobCollectionsClient client.
func NewJobCollectionsClientWithBaseURI(baseURI string, subscriptionID string) JobCollectionsClient {
	return JobCollectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// CreateOrUpdate provisions a new job collection or updates an existing job collection.
//
// resourceGroupName is the resource group name. jobCollectionName is the job collection name. jobCollection is the
// job collection definition.
func (client JobCollectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, jobCollectionName string, jobCollection JobCollectionDefinition) (result JobCollectionDefinition, err error) {
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, jobCollectionName, jobCollection)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client JobCollectionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, jobCollectionName string, jobCollection JobCollectionDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobCollectionName": autorest.Encode("path", jobCollectionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}", pathParameters),
		autorest.WithJSON(jobCollection),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client JobCollectionsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client JobCollectionsClient) CreateOrUpdateResponder(resp *http.Response) (result JobCollectionDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// Delete deletes a job collection.
//
// resourceGroupName is the resource group name. jobCollectionName is the job collection name.
func (client JobCollectionsClient) Delete(ctx context.Context, resourceGroupName string, jobCollectionName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, jobCollectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// DeletePreparer prepares the Delete request.
func (client JobCollectionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, jobCollectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobCollectionName": autorest.Encode("path", jobCollectionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client JobCollectionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client JobCollectionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// Disable disables all of the jobs in the job collection.
//
// resourceGroupName is the resource group name. jobCollectionName is the job collection name.
func (client JobCollectionsClient) Disable(ctx context.Context, resourceGroupName string, jobCollectionName string) (result autorest.Response, err error) {
	req, err := client.DisablePreparer(ctx, resourceGroupName, jobCollectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Disable", nil, "Failure preparing request")
		return
	}

	resp, err := client.DisableSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Disable", resp, "Failure sending request")
		return
	}

	result, err = client.DisableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Disable", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// DisablePreparer prepares the Disable request.
func (client JobCollectionsClient) DisablePreparer(ctx context.Context, resourceGroupName string, jobCollectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobCollectionName": autorest.Encode("path", jobCollectionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/disable", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// DisableSender sends the Disable request. The method will close the
// http.Response Body if it receives an error.
func (client JobCollectionsClient) DisableSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// DisableResponder handles the response to the Disable request. The method always
// closes the http.Response Body.
func (client JobCollectionsClient) DisableResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// Enable enables all of the jobs in the job collection.
//
// resourceGroupName is the resource group name. jobCollectionName is the job collection name.
func (client JobCollectionsClient) Enable(ctx context.Context, resourceGroupName string, jobCollectionName string) (result autorest.Response, err error) {
	req, err := client.EnablePreparer(ctx, resourceGroupName, jobCollectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Enable", nil, "Failure preparing request")
		return
	}

	resp, err := client.EnableSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Enable", resp, "Failure sending request")
		return
	}

	result, err = client.EnableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Enable", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// EnablePreparer prepares the Enable request.
func (client JobCollectionsClient) EnablePreparer(ctx context.Context, resourceGroupName string, jobCollectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobCollectionName": autorest.Encode("path", jobCollectionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/enable", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// EnableSender sends the Enable request. The method will close the
// http.Response Body if it receives an error.
func (client JobCollectionsClient) EnableSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// EnableResponder handles the response to the Enable request. The method always
// closes the http.Response Body.
func (client JobCollectionsClient) EnableResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// Get gets a job collection.
//
// resourceGroupName is the resource group name. jobCollectionName is the job collection name.
func (client JobCollectionsClient) Get(ctx context.Context, resourceGroupName string, jobCollectionName string) (result JobCollectionDefinition, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, jobCollectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// GetPreparer prepares the Get request.
func (client JobCollectionsClient) GetPreparer(ctx context.Context, resourceGroupName string, jobCollectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobCollectionName": autorest.Encode("path", jobCollectionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client JobCollectionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client JobCollectionsClient) GetResponder(resp *http.Response) (result JobCollectionDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// ListByResourceGroup gets all job collections under specified resource group.
//
// resourceGroupName is the resource group name.
func (client JobCollectionsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result JobCollectionListResultPage, err error) {
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.jclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.jclr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client JobCollectionsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client JobCollectionsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client JobCollectionsClient) ListByResourceGroupResponder(resp *http.Response) (result JobCollectionListResult, err error) {
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
func (client JobCollectionsClient) listByResourceGroupNextResults(lastResults JobCollectionListResult) (result JobCollectionListResult, err error) {
	req, err := lastResults.jobCollectionListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client JobCollectionsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result JobCollectionListResultIterator, err error) {
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// ListBySubscription gets all job collections under specified subscription.
func (client JobCollectionsClient) ListBySubscription(ctx context.Context) (result JobCollectionListResultPage, err error) {
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.jclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.jclr, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "ListBySubscription", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client JobCollectionsClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Scheduler/jobCollections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client JobCollectionsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client JobCollectionsClient) ListBySubscriptionResponder(resp *http.Response) (result JobCollectionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client JobCollectionsClient) listBySubscriptionNextResults(lastResults JobCollectionListResult) (result JobCollectionListResult, err error) {
	req, err := lastResults.jobCollectionListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client JobCollectionsClient) ListBySubscriptionComplete(ctx context.Context) (result JobCollectionListResultIterator, err error) {
	result.page, err = client.ListBySubscription(ctx)
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// Patch patches an existing job collection.
//
// resourceGroupName is the resource group name. jobCollectionName is the job collection name. jobCollection is the
// job collection definition.
func (client JobCollectionsClient) Patch(ctx context.Context, resourceGroupName string, jobCollectionName string, jobCollection JobCollectionDefinition) (result JobCollectionDefinition, err error) {
	req, err := client.PatchPreparer(ctx, resourceGroupName, jobCollectionName, jobCollection)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Patch", nil, "Failure preparing request")
		return
	}

	resp, err := client.PatchSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Patch", resp, "Failure sending request")
		return
	}

	result, err = client.PatchResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Patch", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// PatchPreparer prepares the Patch request.
func (client JobCollectionsClient) PatchPreparer(ctx context.Context, resourceGroupName string, jobCollectionName string, jobCollection JobCollectionDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobCollectionName": autorest.Encode("path", jobCollectionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}", pathParameters),
		autorest.WithJSON(jobCollection),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// PatchSender sends the Patch request. The method will close the
// http.Response Body if it receives an error.
func (client JobCollectionsClient) PatchSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/scheduler/mgmt/2014-08-01-preview/scheduler instead.
// PatchResponder handles the response to the Patch request. The method always
// closes the http.Response Body.
func (client JobCollectionsClient) PatchResponder(resp *http.Response) (result JobCollectionDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
