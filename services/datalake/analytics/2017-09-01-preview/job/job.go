package job

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

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// Client is the creates an Azure Data Lake Analytics job client.
type Client struct {
	BaseClient
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// NewClient creates an instance of the Client client.
func NewClient() Client {
	return Client{New()}
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// Build builds (compiles) the specified job in the specified Data Lake Analytics account for job correctness and
// validation.
//
// accountName is the Azure Data Lake Analytics account to execute job operations on. parameters is the parameters
// to build a job.
func (client Client) Build(ctx context.Context, accountName string, parameters BuildJobParameters) (result Information, err error) {
	req, err := client.BuildPreparer(ctx, accountName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "Build", nil, "Failure preparing request")
		return
	}

	resp, err := client.BuildSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "job.Client", "Build", resp, "Failure sending request")
		return
	}

	result, err = client.BuildResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "Build", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// BuildPreparer prepares the Build request.
func (client Client) BuildPreparer(ctx context.Context, accountName string, parameters BuildJobParameters) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPath("/buildJob"),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// BuildSender sends the Build request. The method will close the
// http.Response Body if it receives an error.
func (client Client) BuildSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// BuildResponder handles the response to the Build request. The method always
// closes the http.Response Body.
func (client Client) BuildResponder(resp *http.Response) (result Information, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// Cancel cancels the running job specified by the job ID.
//
// accountName is the Azure Data Lake Analytics account to execute job operations on. jobIdentity is job
// identifier. Uniquely identifies the job across all jobs submitted to the service.
func (client Client) Cancel(ctx context.Context, accountName string, jobIdentity uuid.UUID) (result CancelFuture, err error) {
	req, err := client.CancelPreparer(ctx, accountName, jobIdentity)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "Cancel", nil, "Failure preparing request")
		return
	}

	result, err = client.CancelSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "Cancel", result.Response(), "Failure sending request")
		return
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// CancelPreparer prepares the Cancel request.
func (client Client) CancelPreparer(ctx context.Context, accountName string, jobIdentity uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	pathParameters := map[string]interface{}{
		"jobIdentity": autorest.Encode("path", jobIdentity),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/jobs/{jobIdentity}/CancelJob", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// CancelSender sends the Cancel request. The method will close the
// http.Response Body if it receives an error.
func (client Client) CancelSender(req *http.Request) (future CancelFuture, err error) {
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
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

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// CancelResponder handles the response to the Cancel request. The method always
// closes the http.Response Body.
func (client Client) CancelResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// Create submits a job to the specified Data Lake Analytics account.
//
// accountName is the Azure Data Lake Analytics account to execute job operations on. jobIdentity is job
// identifier. Uniquely identifies the job across all jobs submitted to the service. parameters is the parameters
// to submit a job.
func (client Client) Create(ctx context.Context, accountName string, jobIdentity uuid.UUID, parameters CreateJobParameters) (result Information, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.Related", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.Related.PipelineName", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.Related.PipelineName", Name: validation.MaxLength, Rule: 260, Chain: nil}}},
						{Target: "parameters.Related.RecurrenceID", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.Related.RecurrenceName", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.Related.RecurrenceName", Name: validation.MaxLength, Rule: 260, Chain: nil}}},
					}}}}}); err != nil {
		return result, validation.NewError("job.Client", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, accountName, jobIdentity, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "job.Client", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "Create", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// CreatePreparer prepares the Create request.
func (client Client) CreatePreparer(ctx context.Context, accountName string, jobIdentity uuid.UUID, parameters CreateJobParameters) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	pathParameters := map[string]interface{}{
		"jobIdentity": autorest.Encode("path", jobIdentity),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/jobs/{jobIdentity}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client Client) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client Client) CreateResponder(resp *http.Response) (result Information, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// Get gets the job information for the specified job ID.
//
// accountName is the Azure Data Lake Analytics account to execute job operations on. jobIdentity is jobInfo ID.
func (client Client) Get(ctx context.Context, accountName string, jobIdentity uuid.UUID) (result Information, err error) {
	req, err := client.GetPreparer(ctx, accountName, jobIdentity)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "job.Client", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "Get", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// GetPreparer prepares the Get request.
func (client Client) GetPreparer(ctx context.Context, accountName string, jobIdentity uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	pathParameters := map[string]interface{}{
		"jobIdentity": autorest.Encode("path", jobIdentity),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/jobs/{jobIdentity}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client Client) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client Client) GetResponder(resp *http.Response) (result Information, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// GetDebugDataPath gets the job debug data information specified by the job ID.
//
// accountName is the Azure Data Lake Analytics account to execute job operations on. jobIdentity is job
// identifier. Uniquely identifies the job across all jobs submitted to the service.
func (client Client) GetDebugDataPath(ctx context.Context, accountName string, jobIdentity uuid.UUID) (result DataPath, err error) {
	req, err := client.GetDebugDataPathPreparer(ctx, accountName, jobIdentity)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "GetDebugDataPath", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDebugDataPathSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "job.Client", "GetDebugDataPath", resp, "Failure sending request")
		return
	}

	result, err = client.GetDebugDataPathResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "GetDebugDataPath", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// GetDebugDataPathPreparer prepares the GetDebugDataPath request.
func (client Client) GetDebugDataPathPreparer(ctx context.Context, accountName string, jobIdentity uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	pathParameters := map[string]interface{}{
		"jobIdentity": autorest.Encode("path", jobIdentity),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/jobs/{jobIdentity}/GetDebugDataPath", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// GetDebugDataPathSender sends the GetDebugDataPath request. The method will close the
// http.Response Body if it receives an error.
func (client Client) GetDebugDataPathSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// GetDebugDataPathResponder handles the response to the GetDebugDataPath request. The method always
// closes the http.Response Body.
func (client Client) GetDebugDataPathResponder(resp *http.Response) (result DataPath, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// GetStatistics gets statistics of the specified job.
//
// accountName is the Azure Data Lake Analytics account to execute job operations on. jobIdentity is job
// Information ID.
func (client Client) GetStatistics(ctx context.Context, accountName string, jobIdentity uuid.UUID) (result Statistics, err error) {
	req, err := client.GetStatisticsPreparer(ctx, accountName, jobIdentity)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "GetStatistics", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetStatisticsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "job.Client", "GetStatistics", resp, "Failure sending request")
		return
	}

	result, err = client.GetStatisticsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "GetStatistics", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// GetStatisticsPreparer prepares the GetStatistics request.
func (client Client) GetStatisticsPreparer(ctx context.Context, accountName string, jobIdentity uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	pathParameters := map[string]interface{}{
		"jobIdentity": autorest.Encode("path", jobIdentity),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/jobs/{jobIdentity}/GetStatistics", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// GetStatisticsSender sends the GetStatistics request. The method will close the
// http.Response Body if it receives an error.
func (client Client) GetStatisticsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// GetStatisticsResponder handles the response to the GetStatistics request. The method always
// closes the http.Response Body.
func (client Client) GetStatisticsResponder(resp *http.Response) (result Statistics, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// List lists the jobs, if any, associated with the specified Data Lake Analytics account. The response includes a link
// to the next page of results, if any.
//
// accountName is the Azure Data Lake Analytics account to execute job operations on. filter is oData filter.
// Optional. top is the number of items to return. Optional. skip is the number of items to skip over before
// returning elements. Optional. selectParameter is oData Select statement. Limits the properties on each entry to
// just those requested, e.g. Categories?$select=CategoryName,Description. Optional. orderby is orderBy clause. One
// or more comma-separated expressions with an optional "asc" (the default) or "desc" depending on the order you'd
// like the values sorted, e.g. Categories?$orderby=CategoryName desc. Optional. count is the Boolean value of true
// or false to request a count of the matching resources included with the resources in the response, e.g.
// Categories?$count=true. Optional.
func (client Client) List(ctx context.Context, accountName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result InfoListResultPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("job.Client", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, accountName, filter, top, skip, selectParameter, orderby, count)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.ilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "job.Client", "List", resp, "Failure sending request")
		return
	}

	result.ilr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "List", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// ListPreparer prepares the List request.
func (client Client) ListPreparer(ctx context.Context, accountName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}
	if count != nil {
		queryParameters["$count"] = autorest.Encode("query", *count)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPath("/jobs"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client Client) ListResponder(resp *http.Response) (result InfoListResult, err error) {
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
func (client Client) listNextResults(lastResults InfoListResult) (result InfoListResult, err error) {
	req, err := lastResults.infoListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "job.Client", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "job.Client", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client Client) ListComplete(ctx context.Context, accountName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result InfoListResultIterator, err error) {
	result.page, err = client.List(ctx, accountName, filter, top, skip, selectParameter, orderby, count)
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// Update updates the job information for the specified job ID. (Only for use internally with Scope job type.)
//
// accountName is the Azure Data Lake Analytics account to execute job operations on. jobIdentity is job
// identifier. Uniquely identifies the job across all jobs submitted to the service. parameters is the parameters
// to update a job.
func (client Client) Update(ctx context.Context, accountName string, jobIdentity uuid.UUID, parameters *UpdateJobParameters) (result UpdateFuture, err error) {
	req, err := client.UpdatePreparer(ctx, accountName, jobIdentity, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// UpdatePreparer prepares the Update request.
func (client Client) UpdatePreparer(ctx context.Context, accountName string, jobIdentity uuid.UUID, parameters *UpdateJobParameters) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	pathParameters := map[string]interface{}{
		"jobIdentity": autorest.Encode("path", jobIdentity),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/jobs/{jobIdentity}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if parameters != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(parameters))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client Client) UpdateSender(req *http.Request) (future UpdateFuture, err error) {
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted))
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client Client) UpdateResponder(resp *http.Response) (result Information, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// Yield pauses the specified job and places it back in the job queue, behind other jobs of equal or higher importance,
// based on priority. (Only for use internally with Scope job type.)
//
// accountName is the Azure Data Lake Analytics account to execute job operations on. jobIdentity is job
// identifier. Uniquely identifies the job across all jobs submitted to the service.
func (client Client) Yield(ctx context.Context, accountName string, jobIdentity uuid.UUID) (result YieldFuture, err error) {
	req, err := client.YieldPreparer(ctx, accountName, jobIdentity)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "Yield", nil, "Failure preparing request")
		return
	}

	result, err = client.YieldSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.Client", "Yield", result.Response(), "Failure sending request")
		return
	}

	return
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// YieldPreparer prepares the Yield request.
func (client Client) YieldPreparer(ctx context.Context, accountName string, jobIdentity uuid.UUID) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName":      accountName,
		"adlaJobDnsSuffix": client.AdlaJobDNSSuffix,
	}

	pathParameters := map[string]interface{}{
		"jobIdentity": autorest.Encode("path", jobIdentity),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{accountName}.{adlaJobDnsSuffix}", urlParameters),
		autorest.WithPathParameters("/jobs/{jobIdentity}/YieldJob", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// YieldSender sends the Yield request. The method will close the
// http.Response Body if it receives an error.
func (client Client) YieldSender(req *http.Request) (future YieldFuture, err error) {
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
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

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/datalake/analytics/2017-09-01-preview/job instead.
// YieldResponder handles the response to the Yield request. The method always
// closes the http.Response Body.
func (client Client) YieldResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
