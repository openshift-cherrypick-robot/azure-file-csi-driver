//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AvailableServiceAliasesClient contains the methods for the AvailableServiceAliases group.
// Don't use this type directly, use NewAvailableServiceAliasesClient() instead.
type AvailableServiceAliasesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAvailableServiceAliasesClient creates a new instance of AvailableServiceAliasesClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAvailableServiceAliasesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AvailableServiceAliasesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AvailableServiceAliasesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Gets all available service aliases for this subscription in this region.
//
// Generated from API version 2023-05-01
//   - location - The location.
//   - options - AvailableServiceAliasesClientListOptions contains the optional parameters for the AvailableServiceAliasesClient.NewListPager
//     method.
func (client *AvailableServiceAliasesClient) NewListPager(location string, options *AvailableServiceAliasesClientListOptions) *runtime.Pager[AvailableServiceAliasesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AvailableServiceAliasesClientListResponse]{
		More: func(page AvailableServiceAliasesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AvailableServiceAliasesClientListResponse) (AvailableServiceAliasesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AvailableServiceAliasesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, location, options)
			}, nil)
			if err != nil {
				return AvailableServiceAliasesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *AvailableServiceAliasesClient) listCreateRequest(ctx context.Context, location string, options *AvailableServiceAliasesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/availableServiceAliases"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AvailableServiceAliasesClient) listHandleResponse(resp *http.Response) (AvailableServiceAliasesClientListResponse, error) {
	result := AvailableServiceAliasesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailableServiceAliasesResult); err != nil {
		return AvailableServiceAliasesClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets all available service aliases for this resource group in this region.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group.
//   - location - The location.
//   - options - AvailableServiceAliasesClientListByResourceGroupOptions contains the optional parameters for the AvailableServiceAliasesClient.NewListByResourceGroupPager
//     method.
func (client *AvailableServiceAliasesClient) NewListByResourceGroupPager(resourceGroupName string, location string, options *AvailableServiceAliasesClientListByResourceGroupOptions) *runtime.Pager[AvailableServiceAliasesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[AvailableServiceAliasesClientListByResourceGroupResponse]{
		More: func(page AvailableServiceAliasesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AvailableServiceAliasesClientListByResourceGroupResponse) (AvailableServiceAliasesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AvailableServiceAliasesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, location, options)
			}, nil)
			if err != nil {
				return AvailableServiceAliasesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AvailableServiceAliasesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, location string, options *AvailableServiceAliasesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/locations/{location}/availableServiceAliases"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AvailableServiceAliasesClient) listByResourceGroupHandleResponse(resp *http.Response) (AvailableServiceAliasesClientListByResourceGroupResponse, error) {
	result := AvailableServiceAliasesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailableServiceAliasesResult); err != nil {
		return AvailableServiceAliasesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}
