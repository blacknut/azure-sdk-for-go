// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/blacknut/azure-sdk-for-go/tools/profileBuilder

package consumption

import original "github.com/blacknut/azure-sdk-for-go/services/consumption/mgmt/2018-03-31/consumption"

type BudgetsClient = original.BudgetsClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type CostTagsClient = original.CostTagsClient
type MarketplacesClient = original.MarketplacesClient
type CategoryType = original.CategoryType

const (
	Cost  CategoryType = original.Cost
	Usage CategoryType = original.Usage
)

type Datagrain = original.Datagrain

const (
	DailyGrain   Datagrain = original.DailyGrain
	MonthlyGrain Datagrain = original.MonthlyGrain
)

type OperatorType = original.OperatorType

const (
	EqualTo              OperatorType = original.EqualTo
	GreaterThan          OperatorType = original.GreaterThan
	GreaterThanOrEqualTo OperatorType = original.GreaterThanOrEqualTo
)

type TimeGrainType = original.TimeGrainType

const (
	Annually  TimeGrainType = original.Annually
	Monthly   TimeGrainType = original.Monthly
	Quarterly TimeGrainType = original.Quarterly
)

type Budget = original.Budget
type BudgetProperties = original.BudgetProperties
type BudgetsListResult = original.BudgetsListResult
type BudgetsListResultIterator = original.BudgetsListResultIterator
type BudgetsListResultPage = original.BudgetsListResultPage
type BudgetTimePeriod = original.BudgetTimePeriod
type CostTag = original.CostTag
type CostTagProperties = original.CostTagProperties
type CostTags = original.CostTags
type CurrentSpend = original.CurrentSpend
type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type Filters = original.Filters
type Marketplace = original.Marketplace
type MarketplaceProperties = original.MarketplaceProperties
type MarketplacesListResult = original.MarketplacesListResult
type MarketplacesListResultIterator = original.MarketplacesListResultIterator
type MarketplacesListResultPage = original.MarketplacesListResultPage
type MeterDetails = original.MeterDetails
type Notification = original.Notification
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type PriceSheetModel = original.PriceSheetModel
type PriceSheetProperties = original.PriceSheetProperties
type PriceSheetResult = original.PriceSheetResult
type ProxyResource = original.ProxyResource
type ReservationDetails = original.ReservationDetails
type ReservationDetailsListResult = original.ReservationDetailsListResult
type ReservationDetailsListResultIterator = original.ReservationDetailsListResultIterator
type ReservationDetailsListResultPage = original.ReservationDetailsListResultPage
type ReservationDetailsProperties = original.ReservationDetailsProperties
type ReservationRecommendations = original.ReservationRecommendations
type ReservationRecommendationsListResult = original.ReservationRecommendationsListResult
type ReservationRecommendationsListResultIterator = original.ReservationRecommendationsListResultIterator
type ReservationRecommendationsListResultPage = original.ReservationRecommendationsListResultPage
type ReservationRecommendationsProperties = original.ReservationRecommendationsProperties
type ReservationSummaries = original.ReservationSummaries
type ReservationSummariesListResult = original.ReservationSummariesListResult
type ReservationSummariesListResultIterator = original.ReservationSummariesListResultIterator
type ReservationSummariesListResultPage = original.ReservationSummariesListResultPage
type ReservationSummariesProperties = original.ReservationSummariesProperties
type Resource = original.Resource
type ResourceAttributes = original.ResourceAttributes
type Tag = original.Tag
type TagProperties = original.TagProperties
type Tags = original.Tags
type UsageDetail = original.UsageDetail
type UsageDetailProperties = original.UsageDetailProperties
type UsageDetailsListResult = original.UsageDetailsListResult
type UsageDetailsListResultIterator = original.UsageDetailsListResultIterator
type UsageDetailsListResultPage = original.UsageDetailsListResultPage
type OperationsClient = original.OperationsClient
type PriceSheetClient = original.PriceSheetClient
type ReservationRecommendationsClient = original.ReservationRecommendationsClient
type ReservationsDetailsClient = original.ReservationsDetailsClient
type ReservationsSummariesClient = original.ReservationsSummariesClient
type TagsClient = original.TagsClient
type UsageDetailsClient = original.UsageDetailsClient

func NewBudgetsClient(subscriptionID string) BudgetsClient {
	return original.NewBudgetsClient(subscriptionID)
}
func NewBudgetsClientWithBaseURI(baseURI string, subscriptionID string) BudgetsClient {
	return original.NewBudgetsClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewCostTagsClient(subscriptionID string) CostTagsClient {
	return original.NewCostTagsClient(subscriptionID)
}
func NewCostTagsClientWithBaseURI(baseURI string, subscriptionID string) CostTagsClient {
	return original.NewCostTagsClientWithBaseURI(baseURI, subscriptionID)
}
func NewMarketplacesClient(subscriptionID string) MarketplacesClient {
	return original.NewMarketplacesClient(subscriptionID)
}
func NewMarketplacesClientWithBaseURI(baseURI string, subscriptionID string) MarketplacesClient {
	return original.NewMarketplacesClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleCategoryTypeValues() []CategoryType {
	return original.PossibleCategoryTypeValues()
}
func PossibleDatagrainValues() []Datagrain {
	return original.PossibleDatagrainValues()
}
func PossibleOperatorTypeValues() []OperatorType {
	return original.PossibleOperatorTypeValues()
}
func PossibleTimeGrainTypeValues() []TimeGrainType {
	return original.PossibleTimeGrainTypeValues()
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPriceSheetClient(subscriptionID string) PriceSheetClient {
	return original.NewPriceSheetClient(subscriptionID)
}
func NewPriceSheetClientWithBaseURI(baseURI string, subscriptionID string) PriceSheetClient {
	return original.NewPriceSheetClientWithBaseURI(baseURI, subscriptionID)
}
func NewReservationRecommendationsClient(subscriptionID string) ReservationRecommendationsClient {
	return original.NewReservationRecommendationsClient(subscriptionID)
}
func NewReservationRecommendationsClientWithBaseURI(baseURI string, subscriptionID string) ReservationRecommendationsClient {
	return original.NewReservationRecommendationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewReservationsDetailsClient(subscriptionID string) ReservationsDetailsClient {
	return original.NewReservationsDetailsClient(subscriptionID)
}
func NewReservationsDetailsClientWithBaseURI(baseURI string, subscriptionID string) ReservationsDetailsClient {
	return original.NewReservationsDetailsClientWithBaseURI(baseURI, subscriptionID)
}
func NewReservationsSummariesClient(subscriptionID string) ReservationsSummariesClient {
	return original.NewReservationsSummariesClient(subscriptionID)
}
func NewReservationsSummariesClientWithBaseURI(baseURI string, subscriptionID string) ReservationsSummariesClient {
	return original.NewReservationsSummariesClientWithBaseURI(baseURI, subscriptionID)
}
func NewTagsClient(subscriptionID string) TagsClient {
	return original.NewTagsClient(subscriptionID)
}
func NewTagsClientWithBaseURI(baseURI string, subscriptionID string) TagsClient {
	return original.NewTagsClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsageDetailsClient(subscriptionID string) UsageDetailsClient {
	return original.NewUsageDetailsClient(subscriptionID)
}
func NewUsageDetailsClientWithBaseURI(baseURI string, subscriptionID string) UsageDetailsClient {
	return original.NewUsageDetailsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
