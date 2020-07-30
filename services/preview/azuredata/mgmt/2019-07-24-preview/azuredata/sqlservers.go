package azuredata

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// SQLServersClient is the the AzureData management API provides a RESTful set of web APIs to manage Azure Data
// Resources.
type SQLServersClient struct {
	BaseClient
}

// NewSQLServersClient creates an instance of the SQLServersClient client.
func NewSQLServersClient(subscriptionID string, subscriptionID1 string) SQLServersClient {
	return NewSQLServersClientWithBaseURI(DefaultBaseURI, subscriptionID, subscriptionID1)
}

// NewSQLServersClientWithBaseURI creates an instance of the SQLServersClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewSQLServersClientWithBaseURI(baseURI string, subscriptionID string, subscriptionID1 string) SQLServersClient {
	return SQLServersClient{NewWithBaseURI(baseURI, subscriptionID, subscriptionID1)}
}

// CreateOrUpdate creates or updates a SQL Server.
// Parameters:
// resourceGroupName - name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal.
// SQLServerRegistrationName - name of the SQL Server registration.
// SQLServerName - name of the SQL Server.
// parameters - the SQL Server to be created or updated.
func (client SQLServersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, SQLServerName string, parameters SQLServer) (result SQLServer, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLServersClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, SQLServerRegistrationName, SQLServerName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServersClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azuredata.SQLServersClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServersClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SQLServersClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, SQLServerName string, parameters SQLServer) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"sqlServerName":             autorest.Encode("path", SQLServerName),
		"sqlServerRegistrationName": autorest.Encode("path", SQLServerRegistrationName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerRegistrations/{sqlServerRegistrationName}/sqlServers/{sqlServerName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SQLServersClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SQLServersClient) CreateOrUpdateResponder(resp *http.Response) (result SQLServer, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a SQL Server.
// Parameters:
// resourceGroupName - name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal.
// SQLServerRegistrationName - name of the SQL Server registration.
// SQLServerName - name of the SQL Server.
func (client SQLServersClient) Delete(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, SQLServerName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLServersClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, SQLServerRegistrationName, SQLServerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServersClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "azuredata.SQLServersClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServersClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SQLServersClient) DeletePreparer(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, SQLServerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"sqlServerName":             autorest.Encode("path", SQLServerName),
		"sqlServerRegistrationName": autorest.Encode("path", SQLServerRegistrationName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerRegistrations/{sqlServerRegistrationName}/sqlServers/{sqlServerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SQLServersClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SQLServersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a SQL Server.
// Parameters:
// resourceGroupName - name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal.
// SQLServerRegistrationName - name of the SQL Server registration.
// SQLServerName - name of the SQL Server.
// expand - the child resources to include in the response.
func (client SQLServersClient) Get(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, SQLServerName string, expand string) (result SQLServer, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLServersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, SQLServerRegistrationName, SQLServerName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azuredata.SQLServersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServersClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SQLServersClient) GetPreparer(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, SQLServerName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"sqlServerName":             autorest.Encode("path", SQLServerName),
		"sqlServerRegistrationName": autorest.Encode("path", SQLServerRegistrationName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerRegistrations/{sqlServerRegistrationName}/sqlServers/{sqlServerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SQLServersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SQLServersClient) GetResponder(resp *http.Response) (result SQLServer, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup gets all SQL Servers in a SQL Server Registration.
// Parameters:
// resourceGroupName - name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal.
// SQLServerRegistrationName - name of the SQL Server registration.
// expand - the child resources to include in the response.
func (client SQLServersClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, expand string) (result SQLServerListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLServersClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.sslr.Response.Response != nil {
				sc = result.sslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, SQLServerRegistrationName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServersClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.sslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "azuredata.SQLServersClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.sslr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServersClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client SQLServersClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"sqlServerRegistrationName": autorest.Encode("path", SQLServerRegistrationName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerRegistrations/{sqlServerRegistrationName}/sqlServers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client SQLServersClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client SQLServersClient) ListByResourceGroupResponder(resp *http.Response) (result SQLServerListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client SQLServersClient) listByResourceGroupNextResults(ctx context.Context, lastResults SQLServerListResult) (result SQLServerListResult, err error) {
	req, err := lastResults.sQLServerListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "azuredata.SQLServersClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "azuredata.SQLServersClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azuredata.SQLServersClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client SQLServersClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, SQLServerRegistrationName string, expand string) (result SQLServerListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLServersClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, SQLServerRegistrationName, expand)
	return
}