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
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ReplicationsClient is the client for the Replications methods of the Containerregistry service.
type ReplicationsClient struct {
	BaseClient
}

// NewReplicationsClient creates an instance of the ReplicationsClient client.
func NewReplicationsClient(subscriptionID string) ReplicationsClient {
	return NewReplicationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewReplicationsClientWithBaseURI creates an instance of the ReplicationsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewReplicationsClientWithBaseURI(baseURI string, subscriptionID string) ReplicationsClient {
	return ReplicationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a replication for a container registry with the specified parameters.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// replicationName - the name of the replication.
// replication - the parameters for creating a replication.
func (client ReplicationsClient) Create(ctx context.Context, resourceGroupName string, registryName string, replicationName string, replication Replication) (result ReplicationsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationsClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: replicationName,
			Constraints: []validation.Constraint{{Target: "replicationName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "replicationName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "replicationName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.ReplicationsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, registryName, replicationName, replication)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ReplicationsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ReplicationsClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ReplicationsClient) CreatePreparer(ctx context.Context, resourceGroupName string, registryName string, replicationName string, replication Replication) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"replicationName":   autorest.Encode("path", replicationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-12-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}", pathParameters),
		autorest.WithJSON(replication),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationsClient) CreateSender(req *http.Request) (future ReplicationsCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ReplicationsClient) CreateResponder(resp *http.Response) (result Replication, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a replication from a container registry.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// replicationName - the name of the replication.
func (client ReplicationsClient) Delete(ctx context.Context, resourceGroupName string, registryName string, replicationName string) (result ReplicationsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: replicationName,
			Constraints: []validation.Constraint{{Target: "replicationName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "replicationName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "replicationName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.ReplicationsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, registryName, replicationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ReplicationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ReplicationsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ReplicationsClient) DeletePreparer(ctx context.Context, resourceGroupName string, registryName string, replicationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"replicationName":   autorest.Encode("path", replicationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-12-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationsClient) DeleteSender(req *http.Request) (future ReplicationsDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ReplicationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the properties of the specified replication.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// replicationName - the name of the replication.
func (client ReplicationsClient) Get(ctx context.Context, resourceGroupName string, registryName string, replicationName string) (result Replication, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: replicationName,
			Constraints: []validation.Constraint{{Target: "replicationName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "replicationName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "replicationName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.ReplicationsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, registryName, replicationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ReplicationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.ReplicationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ReplicationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ReplicationsClient) GetPreparer(ctx context.Context, resourceGroupName string, registryName string, replicationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"replicationName":   autorest.Encode("path", replicationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-12-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ReplicationsClient) GetResponder(resp *http.Response) (result Replication, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all the replications for the specified container registry.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
func (client ReplicationsClient) List(ctx context.Context, resourceGroupName string, registryName string) (result ReplicationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationsClient.List")
		defer func() {
			sc := -1
			if result.rlr.Response.Response != nil {
				sc = result.rlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.ReplicationsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, registryName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ReplicationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.ReplicationsClient", "List", resp, "Failure sending request")
		return
	}

	result.rlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ReplicationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ReplicationsClient) ListPreparer(ctx context.Context, resourceGroupName string, registryName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-12-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ReplicationsClient) ListResponder(resp *http.Response) (result ReplicationListResult, err error) {
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
func (client ReplicationsClient) listNextResults(ctx context.Context, lastResults ReplicationListResult) (result ReplicationListResult, err error) {
	req, err := lastResults.replicationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerregistry.ReplicationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerregistry.ReplicationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ReplicationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReplicationsClient) ListComplete(ctx context.Context, resourceGroupName string, registryName string) (result ReplicationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, registryName)
	return
}

// Update updates a replication for a container registry with the specified parameters.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// replicationName - the name of the replication.
// replicationUpdateParameters - the parameters for updating a replication.
func (client ReplicationsClient) Update(ctx context.Context, resourceGroupName string, registryName string, replicationName string, replicationUpdateParameters ReplicationUpdateParameters) (result ReplicationsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationsClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: replicationName,
			Constraints: []validation.Constraint{{Target: "replicationName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "replicationName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "replicationName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.ReplicationsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, registryName, replicationName, replicationUpdateParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ReplicationsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.ReplicationsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ReplicationsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, registryName string, replicationName string, replicationUpdateParameters ReplicationUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"replicationName":   autorest.Encode("path", replicationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-12-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}", pathParameters),
		autorest.WithJSON(replicationUpdateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationsClient) UpdateSender(req *http.Request) (future ReplicationsUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ReplicationsClient) UpdateResponder(resp *http.Response) (result Replication, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}