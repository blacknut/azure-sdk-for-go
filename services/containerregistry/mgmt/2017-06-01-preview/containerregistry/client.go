// Package containerregistry implements the Azure ARM Containerregistry service API version 2017-06-01-preview.
//
//
package containerregistry

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
	"github.com/Azure/go-autorest/autorest"
)

const (
	// DefaultBaseURI is the default URI used for the service Containerregistry
	DefaultBaseURI = "https://management.azure.com"
)

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/containerregistry/mgmt/2017-06-01-preview/containerregistry instead.
// BaseClient is the base client for Containerregistry.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/containerregistry/mgmt/2017-06-01-preview/containerregistry instead.
// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// Deprecated: Please use package github.com/blacknut/azure-sdk-for-go/services/preview/containerregistry/mgmt/2017-06-01-preview/containerregistry instead.
// NewWithBaseURI creates an instance of the BaseClient client.
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}
