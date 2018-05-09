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

package training

import original "github.com/blacknut/azure-sdk-for-go/services/cognitiveservices/v1.2/customvision/training"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type OrderBy = original.OrderBy

const (
	Newest    OrderBy = original.Newest
	Oldest    OrderBy = original.Oldest
	Suggested OrderBy = original.Suggested
)

type Platform = original.Platform

const (
	CoreML     Platform = original.CoreML
	TensorFlow Platform = original.TensorFlow
)

type Status = original.Status

const (
	ErrorImageFormat    Status = original.ErrorImageFormat
	ErrorImageSize      Status = original.ErrorImageSize
	ErrorLimitExceed    Status = original.ErrorLimitExceed
	ErrorSource         Status = original.ErrorSource
	ErrorStorage        Status = original.ErrorStorage
	ErrorTagLimitExceed Status = original.ErrorTagLimitExceed
	ErrorUnknown        Status = original.ErrorUnknown
	OK                  Status = original.OK
	OKDuplicate         Status = original.OKDuplicate
)

type Status1 = original.Status1

const (
	Done      Status1 = original.Done
	Exporting Status1 = original.Exporting
	Failed    Status1 = original.Failed
)

type Account = original.Account
type AccountQuota = original.AccountQuota
type APIKeys = original.APIKeys
type Domain = original.Domain
type Export = original.Export
type Image = original.Image
type ImageCreateResult = original.ImageCreateResult
type ImageCreateSummary = original.ImageCreateSummary
type ImageFileCreateBatch = original.ImageFileCreateBatch
type ImageFileCreateEntry = original.ImageFileCreateEntry
type ImageIDCreateBatch = original.ImageIDCreateBatch
type ImageIDCreateEntry = original.ImageIDCreateEntry
type ImagePredictionResult = original.ImagePredictionResult
type ImageTag = original.ImageTag
type ImageTagCreateBatch = original.ImageTagCreateBatch
type ImageTagCreateEntry = original.ImageTagCreateEntry
type ImageTagCreateSummary = original.ImageTagCreateSummary
type ImageTagPrediction = original.ImageTagPrediction
type ImageURL = original.ImageURL
type ImageURLCreateBatch = original.ImageURLCreateBatch
type ImageURLCreateEntry = original.ImageURLCreateEntry
type Iteration = original.Iteration
type IterationPerformance = original.IterationPerformance
type KeyPair = original.KeyPair
type ListDomain = original.ListDomain
type ListExport = original.ListExport
type ListImage = original.ListImage
type ListIteration = original.ListIteration
type ListProject = original.ListProject
type PerProjectQuota = original.PerProjectQuota
type Prediction = original.Prediction
type PredictionQuery = original.PredictionQuery
type PredictionQueryTag = original.PredictionQueryTag
type PredictionQueryToken = original.PredictionQueryToken
type PredictionTag = original.PredictionTag
type Project = original.Project
type ProjectSettings = original.ProjectSettings
type Quota = original.Quota
type Tag = original.Tag
type TagList = original.TagList
type TagPerformance = original.TagPerformance

func New(aPIKey string) BaseClient {
	return original.New(aPIKey)
}
func NewWithBaseURI(baseURI string, aPIKey string) BaseClient {
	return original.NewWithBaseURI(baseURI, aPIKey)
}
func PossibleOrderByValues() []OrderBy {
	return original.PossibleOrderByValues()
}
func PossiblePlatformValues() []Platform {
	return original.PossiblePlatformValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func PossibleStatus1Values() []Status1 {
	return original.PossibleStatus1Values()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
