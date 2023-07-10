package network

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
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

// StaticMembersClient is the network Client
type StaticMembersClient struct {
	BaseClient
}

// NewStaticMembersClient creates an instance of the StaticMembersClient client.
func NewStaticMembersClient(subscriptionID string) StaticMembersClient {
	return NewStaticMembersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewStaticMembersClientWithBaseURI creates an instance of the StaticMembersClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewStaticMembersClientWithBaseURI(baseURI string, subscriptionID string) StaticMembersClient {
	return StaticMembersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a static member.
// Parameters:
// parameters - parameters supplied to the specify the static member to create
// resourceGroupName - the name of the resource group.
// networkManagerName - the name of the network manager.
// networkGroupName - the name of the network group.
// staticMemberName - the name of the static member.
func (client StaticMembersClient) CreateOrUpdate(ctx context.Context, parameters StaticMember, resourceGroupName string, networkManagerName string, networkGroupName string, staticMemberName string) (result StaticMember, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StaticMembersClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, parameters, resourceGroupName, networkManagerName, networkGroupName, staticMemberName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.StaticMembersClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.StaticMembersClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.StaticMembersClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client StaticMembersClient) CreateOrUpdatePreparer(ctx context.Context, parameters StaticMember, resourceGroupName string, networkManagerName string, networkGroupName string, staticMemberName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkGroupName":   autorest.Encode("path", networkGroupName),
		"networkManagerName": autorest.Encode("path", networkManagerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"staticMemberName":   autorest.Encode("path", staticMemberName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/networkGroups/{networkGroupName}/staticMembers/{staticMemberName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client StaticMembersClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client StaticMembersClient) CreateOrUpdateResponder(resp *http.Response) (result StaticMember, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a static member.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkManagerName - the name of the network manager.
// networkGroupName - the name of the network group.
// staticMemberName - the name of the static member.
func (client StaticMembersClient) Delete(ctx context.Context, resourceGroupName string, networkManagerName string, networkGroupName string, staticMemberName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StaticMembersClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, networkManagerName, networkGroupName, staticMemberName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.StaticMembersClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "network.StaticMembersClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.StaticMembersClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client StaticMembersClient) DeletePreparer(ctx context.Context, resourceGroupName string, networkManagerName string, networkGroupName string, staticMemberName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkGroupName":   autorest.Encode("path", networkGroupName),
		"networkManagerName": autorest.Encode("path", networkManagerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"staticMemberName":   autorest.Encode("path", staticMemberName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/networkGroups/{networkGroupName}/staticMembers/{staticMemberName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client StaticMembersClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client StaticMembersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified static member.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkManagerName - the name of the network manager.
// networkGroupName - the name of the network group.
// staticMemberName - the name of the static member.
func (client StaticMembersClient) Get(ctx context.Context, resourceGroupName string, networkManagerName string, networkGroupName string, staticMemberName string) (result StaticMember, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StaticMembersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, networkManagerName, networkGroupName, staticMemberName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.StaticMembersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.StaticMembersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.StaticMembersClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client StaticMembersClient) GetPreparer(ctx context.Context, resourceGroupName string, networkManagerName string, networkGroupName string, staticMemberName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkGroupName":   autorest.Encode("path", networkGroupName),
		"networkManagerName": autorest.Encode("path", networkManagerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"staticMemberName":   autorest.Encode("path", staticMemberName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/networkGroups/{networkGroupName}/staticMembers/{staticMemberName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client StaticMembersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client StaticMembersClient) GetResponder(resp *http.Response) (result StaticMember, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the specified static member.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkManagerName - the name of the network manager.
// networkGroupName - the name of the network group.
// top - an optional query parameter which specifies the maximum number of records to be returned by the
// server.
// skipToken - skipToken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skipToken parameter that
// specifies a starting point to use for subsequent calls.
func (client StaticMembersClient) List(ctx context.Context, resourceGroupName string, networkManagerName string, networkGroupName string, top *int32, skipToken string) (result StaticMemberListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StaticMembersClient.List")
		defer func() {
			sc := -1
			if result.smlr.Response.Response != nil {
				sc = result.smlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: int64(20), Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("network.StaticMembersClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, networkManagerName, networkGroupName, top, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.StaticMembersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.smlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.StaticMembersClient", "List", resp, "Failure sending request")
		return
	}

	result.smlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.StaticMembersClient", "List", resp, "Failure responding to request")
		return
	}
	if result.smlr.hasNextLink() && result.smlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client StaticMembersClient) ListPreparer(ctx context.Context, resourceGroupName string, networkManagerName string, networkGroupName string, top *int32, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkGroupName":   autorest.Encode("path", networkGroupName),
		"networkManagerName": autorest.Encode("path", networkManagerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/networkGroups/{networkGroupName}/staticMembers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client StaticMembersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client StaticMembersClient) ListResponder(resp *http.Response) (result StaticMemberListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client StaticMembersClient) listNextResults(ctx context.Context, lastResults StaticMemberListResult) (result StaticMemberListResult, err error) {
	req, err := lastResults.staticMemberListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.StaticMembersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.StaticMembersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.StaticMembersClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client StaticMembersClient) ListComplete(ctx context.Context, resourceGroupName string, networkManagerName string, networkGroupName string, top *int32, skipToken string) (result StaticMemberListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StaticMembersClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, networkManagerName, networkGroupName, top, skipToken)
	return
}
