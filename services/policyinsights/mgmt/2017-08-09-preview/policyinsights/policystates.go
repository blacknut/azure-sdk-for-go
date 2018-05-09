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
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/validation"
)

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// PolicyStatesClient is the client for the PolicyStates methods of the Policyinsights service.
type PolicyStatesClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// NewPolicyStatesClient creates an instance of the PolicyStatesClient client.
func NewPolicyStatesClient() PolicyStatesClient {
	return NewPolicyStatesClientWithBaseURI(DefaultBaseURI)
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// NewPolicyStatesClientWithBaseURI creates an instance of the PolicyStatesClient client.
func NewPolicyStatesClientWithBaseURI(baseURI string) PolicyStatesClient {
	return PolicyStatesClient{NewWithBaseURI(baseURI)}
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// ListQueryResultsForManagementGroup queries policy states for the resources under the management group.
//
// policyStatesResource is the virtual resource under PolicyStates resource type. In a given time range, 'latest'
// represents the latest policy state(s), whereas 'default' represents all policy state(s). managementGroupName is
// management group name. top is maximum number of records to return. orderBy is ordering expression using OData
// notation. One or more comma-separated column names with an optional "desc" (the default) or "asc", e.g.
// "$orderby=PolicyAssignmentId, ResourceId asc". selectParameter is select expression using OData notation. Limits
// the columns on each record to just those requested, e.g. "$select=PolicyAssignmentId, ResourceId". from is ISO
// 8601 formatted timestamp specifying the start time of the interval to query. When not specified, the service
// uses ($to - 1-day). toParameter is ISO 8601 formatted timestamp specifying the end time of the interval to
// query. When not specified, the service uses request time. filter is oData filter expression. apply is oData
// apply expression for aggregations.
func (client PolicyStatesClient) ListQueryResultsForManagementGroup(ctx context.Context, policyStatesResource PolicyStatesResource, managementGroupName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result PolicyStatesQueryResults, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("policyinsights.PolicyStatesClient", "ListQueryResultsForManagementGroup", err.Error())
	}

	req, err := client.ListQueryResultsForManagementGroupPreparer(ctx, policyStatesResource, managementGroupName, top, orderBy, selectParameter, from, toParameter, filter, apply)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyStatesClient", "ListQueryResultsForManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListQueryResultsForManagementGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyStatesClient", "ListQueryResultsForManagementGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListQueryResultsForManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyStatesClient", "ListQueryResultsForManagementGroup", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// ListQueryResultsForManagementGroupPreparer prepares the ListQueryResultsForManagementGroup request.
func (client PolicyStatesClient) ListQueryResultsForManagementGroupPreparer(ctx context.Context, policyStatesResource PolicyStatesResource, managementGroupName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupName":       autorest.Encode("path", managementGroupName),
		"managementGroupsNamespace": autorest.Encode("path", "Microsoft.Management"),
		"policyStatesResource":      autorest.Encode("path", policyStatesResource),
	}

	const APIVersion = "2017-08-09-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderBy) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderBy)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if from != nil {
		queryParameters["$from"] = autorest.Encode("query", *from)
	}
	if toParameter != nil {
		queryParameters["$to"] = autorest.Encode("query", *toParameter)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(apply) > 0 {
		queryParameters["$apply"] = autorest.Encode("query", apply)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/{managementGroupsNamespace}/managementGroups/{managementGroupName}/providers/Microsoft.PolicyInsights/policyStates/{policyStatesResource}/queryResults", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// ListQueryResultsForManagementGroupSender sends the ListQueryResultsForManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client PolicyStatesClient) ListQueryResultsForManagementGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// ListQueryResultsForManagementGroupResponder handles the response to the ListQueryResultsForManagementGroup request. The method always
// closes the http.Response Body.
func (client PolicyStatesClient) ListQueryResultsForManagementGroupResponder(resp *http.Response) (result PolicyStatesQueryResults, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// ListQueryResultsForResource queries policy states for the resource.
//
// policyStatesResource is the virtual resource under PolicyStates resource type. In a given time range, 'latest'
// represents the latest policy state(s), whereas 'default' represents all policy state(s). resourceID is resource
// ID. top is maximum number of records to return. orderBy is ordering expression using OData notation. One or more
// comma-separated column names with an optional "desc" (the default) or "asc", e.g. "$orderby=PolicyAssignmentId,
// ResourceId asc". selectParameter is select expression using OData notation. Limits the columns on each record to
// just those requested, e.g. "$select=PolicyAssignmentId, ResourceId". from is ISO 8601 formatted timestamp
// specifying the start time of the interval to query. When not specified, the service uses ($to - 1-day).
// toParameter is ISO 8601 formatted timestamp specifying the end time of the interval to query. When not
// specified, the service uses request time. filter is oData filter expression. apply is oData apply expression for
// aggregations.
func (client PolicyStatesClient) ListQueryResultsForResource(ctx context.Context, policyStatesResource PolicyStatesResource, resourceID string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result PolicyStatesQueryResults, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("policyinsights.PolicyStatesClient", "ListQueryResultsForResource", err.Error())
	}

	req, err := client.ListQueryResultsForResourcePreparer(ctx, policyStatesResource, resourceID, top, orderBy, selectParameter, from, toParameter, filter, apply)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyStatesClient", "ListQueryResultsForResource", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListQueryResultsForResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyStatesClient", "ListQueryResultsForResource", resp, "Failure sending request")
		return
	}

	result, err = client.ListQueryResultsForResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyStatesClient", "ListQueryResultsForResource", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// ListQueryResultsForResourcePreparer prepares the ListQueryResultsForResource request.
func (client PolicyStatesClient) ListQueryResultsForResourcePreparer(ctx context.Context, policyStatesResource PolicyStatesResource, resourceID string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyStatesResource": autorest.Encode("path", policyStatesResource),
		"resourceId":           resourceID,
	}

	const APIVersion = "2017-08-09-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderBy) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderBy)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if from != nil {
		queryParameters["$from"] = autorest.Encode("query", *from)
	}
	if toParameter != nil {
		queryParameters["$to"] = autorest.Encode("query", *toParameter)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(apply) > 0 {
		queryParameters["$apply"] = autorest.Encode("query", apply)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceId}/providers/Microsoft.PolicyInsights/policyStates/{policyStatesResource}/queryResults", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// ListQueryResultsForResourceSender sends the ListQueryResultsForResource request. The method will close the
// http.Response Body if it receives an error.
func (client PolicyStatesClient) ListQueryResultsForResourceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// ListQueryResultsForResourceResponder handles the response to the ListQueryResultsForResource request. The method always
// closes the http.Response Body.
func (client PolicyStatesClient) ListQueryResultsForResourceResponder(resp *http.Response) (result PolicyStatesQueryResults, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// ListQueryResultsForResourceGroup queries policy states for the resources under the resource group.
//
// policyStatesResource is the virtual resource under PolicyStates resource type. In a given time range, 'latest'
// represents the latest policy state(s), whereas 'default' represents all policy state(s). subscriptionID is
// microsoft Azure subscription ID. resourceGroupName is resource group name. top is maximum number of records to
// return. orderBy is ordering expression using OData notation. One or more comma-separated column names with an
// optional "desc" (the default) or "asc", e.g. "$orderby=PolicyAssignmentId, ResourceId asc". selectParameter is
// select expression using OData notation. Limits the columns on each record to just those requested, e.g.
// "$select=PolicyAssignmentId, ResourceId". from is ISO 8601 formatted timestamp specifying the start time of the
// interval to query. When not specified, the service uses ($to - 1-day). toParameter is ISO 8601 formatted
// timestamp specifying the end time of the interval to query. When not specified, the service uses request time.
// filter is oData filter expression. apply is oData apply expression for aggregations.
func (client PolicyStatesClient) ListQueryResultsForResourceGroup(ctx context.Context, policyStatesResource PolicyStatesResource, subscriptionID string, resourceGroupName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result PolicyStatesQueryResults, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("policyinsights.PolicyStatesClient", "ListQueryResultsForResourceGroup", err.Error())
	}

	req, err := client.ListQueryResultsForResourceGroupPreparer(ctx, policyStatesResource, subscriptionID, resourceGroupName, top, orderBy, selectParameter, from, toParameter, filter, apply)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyStatesClient", "ListQueryResultsForResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListQueryResultsForResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyStatesClient", "ListQueryResultsForResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListQueryResultsForResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyStatesClient", "ListQueryResultsForResourceGroup", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// ListQueryResultsForResourceGroupPreparer prepares the ListQueryResultsForResourceGroup request.
func (client PolicyStatesClient) ListQueryResultsForResourceGroupPreparer(ctx context.Context, policyStatesResource PolicyStatesResource, subscriptionID string, resourceGroupName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyStatesResource": autorest.Encode("path", policyStatesResource),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2017-08-09-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderBy) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderBy)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if from != nil {
		queryParameters["$from"] = autorest.Encode("query", *from)
	}
	if toParameter != nil {
		queryParameters["$to"] = autorest.Encode("query", *toParameter)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(apply) > 0 {
		queryParameters["$apply"] = autorest.Encode("query", apply)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PolicyInsights/policyStates/{policyStatesResource}/queryResults", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// ListQueryResultsForResourceGroupSender sends the ListQueryResultsForResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client PolicyStatesClient) ListQueryResultsForResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// ListQueryResultsForResourceGroupResponder handles the response to the ListQueryResultsForResourceGroup request. The method always
// closes the http.Response Body.
func (client PolicyStatesClient) ListQueryResultsForResourceGroupResponder(resp *http.Response) (result PolicyStatesQueryResults, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// ListQueryResultsForSubscription queries policy states for the resources under the subscription.
//
// policyStatesResource is the virtual resource under PolicyStates resource type. In a given time range, 'latest'
// represents the latest policy state(s), whereas 'default' represents all policy state(s). subscriptionID is
// microsoft Azure subscription ID. top is maximum number of records to return. orderBy is ordering expression
// using OData notation. One or more comma-separated column names with an optional "desc" (the default) or "asc",
// e.g. "$orderby=PolicyAssignmentId, ResourceId asc". selectParameter is select expression using OData notation.
// Limits the columns on each record to just those requested, e.g. "$select=PolicyAssignmentId, ResourceId". from
// is ISO 8601 formatted timestamp specifying the start time of the interval to query. When not specified, the
// service uses ($to - 1-day). toParameter is ISO 8601 formatted timestamp specifying the end time of the interval
// to query. When not specified, the service uses request time. filter is oData filter expression. apply is oData
// apply expression for aggregations.
func (client PolicyStatesClient) ListQueryResultsForSubscription(ctx context.Context, policyStatesResource PolicyStatesResource, subscriptionID string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result PolicyStatesQueryResults, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("policyinsights.PolicyStatesClient", "ListQueryResultsForSubscription", err.Error())
	}

	req, err := client.ListQueryResultsForSubscriptionPreparer(ctx, policyStatesResource, subscriptionID, top, orderBy, selectParameter, from, toParameter, filter, apply)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyStatesClient", "ListQueryResultsForSubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListQueryResultsForSubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyStatesClient", "ListQueryResultsForSubscription", resp, "Failure sending request")
		return
	}

	result, err = client.ListQueryResultsForSubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policyinsights.PolicyStatesClient", "ListQueryResultsForSubscription", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// ListQueryResultsForSubscriptionPreparer prepares the ListQueryResultsForSubscription request.
func (client PolicyStatesClient) ListQueryResultsForSubscriptionPreparer(ctx context.Context, policyStatesResource PolicyStatesResource, subscriptionID string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyStatesResource": autorest.Encode("path", policyStatesResource),
		"subscriptionId":       autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2017-08-09-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderBy) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderBy)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if from != nil {
		queryParameters["$from"] = autorest.Encode("query", *from)
	}
	if toParameter != nil {
		queryParameters["$to"] = autorest.Encode("query", *toParameter)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(apply) > 0 {
		queryParameters["$apply"] = autorest.Encode("query", apply)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.PolicyInsights/policyStates/{policyStatesResource}/queryResults", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// ListQueryResultsForSubscriptionSender sends the ListQueryResultsForSubscription request. The method will close the
// http.Response Body if it receives an error.
func (client PolicyStatesClient) ListQueryResultsForSubscriptionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/policyinsights/mgmt/2017-08-09-preview/policyinsights instead.
// ListQueryResultsForSubscriptionResponder handles the response to the ListQueryResultsForSubscription request. The method always
// closes the http.Response Body.
func (client PolicyStatesClient) ListQueryResultsForSubscriptionResponder(resp *http.Response) (result PolicyStatesQueryResults, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
