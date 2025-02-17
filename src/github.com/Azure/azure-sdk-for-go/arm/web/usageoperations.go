package web

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.14.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
	"net/url"
)

// UsageOperationsClient is the use these APIs to manage Azure Websites
// resources through the Azure Resource Manager. All task operations conform
// to the HTTP/1.1 protocol specification and each operation returns an
// x-ms-request-id header that can be used to obtain information about the
// request. You must make sure that requests made to these resources are
// secure. For more information, see <a
// href="https://msdn.microsoft.com/en-us/library/azure/dn790557.aspx">Authenticating
// Azure Resource Manager requests.</a>
type UsageOperationsClient struct {
	ManagementClient
}

// NewUsageOperationsClient creates an instance of the UsageOperationsClient
// client.
func NewUsageOperationsClient(subscriptionID string) UsageOperationsClient {
	return NewUsageOperationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUsageOperationsClientWithBaseURI creates an instance of the
// UsageOperationsClient client.
func NewUsageOperationsClientWithBaseURI(baseURI string, subscriptionID string) UsageOperationsClient {
	return UsageOperationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetUsage sends the get usage request.
//
// resourceGroupName is name of resource group environmentName is environment
// name lastID is last marker that was returned from the batch batchSize is
// size of the batch to be returned.
func (client UsageOperationsClient) GetUsage(resourceGroupName string, environmentName string, lastID string, batchSize int32) (result ObjectSet, err error) {
	req, err := client.GetUsagePreparer(resourceGroupName, environmentName, lastID, batchSize)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web/UsageOperationsClient", "GetUsage", nil, "Failure preparing request")
	}

	resp, err := client.GetUsageSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web/UsageOperationsClient", "GetUsage", resp, "Failure sending request")
	}

	result, err = client.GetUsageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web/UsageOperationsClient", "GetUsage", resp, "Failure responding to request")
	}

	return
}

// GetUsagePreparer prepares the GetUsage request.
func (client UsageOperationsClient) GetUsagePreparer(resourceGroupName string, environmentName string, lastID string, batchSize int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentName":   url.QueryEscape(environmentName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
		"batchSize":   batchSize,
		"lastId":      lastID,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web.Admin/environments/{environmentName}/usage"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetUsageSender sends the GetUsage request. The method will close the
// http.Response Body if it receives an error.
func (client UsageOperationsClient) GetUsageSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetUsageResponder handles the response to the GetUsage request. The method always
// closes the http.Response Body.
func (client UsageOperationsClient) GetUsageResponder(resp *http.Response) (result ObjectSet, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
