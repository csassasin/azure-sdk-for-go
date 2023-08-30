//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
	"net/http"
	"net/url"
	"regexp"
)

// ProximityPlacementGroupsServer is a fake server for instances of the armcompute.ProximityPlacementGroupsClient type.
type ProximityPlacementGroupsServer struct{
	// CreateOrUpdate is the fake for method ProximityPlacementGroupsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters armcompute.ProximityPlacementGroup, options *armcompute.ProximityPlacementGroupsClientCreateOrUpdateOptions) (resp azfake.Responder[armcompute.ProximityPlacementGroupsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ProximityPlacementGroupsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK
	Delete func(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, options *armcompute.ProximityPlacementGroupsClientDeleteOptions) (resp azfake.Responder[armcompute.ProximityPlacementGroupsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ProximityPlacementGroupsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, options *armcompute.ProximityPlacementGroupsClientGetOptions) (resp azfake.Responder[armcompute.ProximityPlacementGroupsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method ProximityPlacementGroupsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armcompute.ProximityPlacementGroupsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armcompute.ProximityPlacementGroupsClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method ProximityPlacementGroupsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armcompute.ProximityPlacementGroupsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armcompute.ProximityPlacementGroupsClientListBySubscriptionResponse])

	// Update is the fake for method ProximityPlacementGroupsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters armcompute.ProximityPlacementGroupUpdate, options *armcompute.ProximityPlacementGroupsClientUpdateOptions) (resp azfake.Responder[armcompute.ProximityPlacementGroupsClientUpdateResponse], errResp azfake.ErrorResponder)

}

// NewProximityPlacementGroupsServerTransport creates a new instance of ProximityPlacementGroupsServerTransport with the provided implementation.
// The returned ProximityPlacementGroupsServerTransport instance is connected to an instance of armcompute.ProximityPlacementGroupsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewProximityPlacementGroupsServerTransport(srv *ProximityPlacementGroupsServer) *ProximityPlacementGroupsServerTransport {
	return &ProximityPlacementGroupsServerTransport{
		srv: srv,
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armcompute.ProximityPlacementGroupsClientListByResourceGroupResponse]](),
		newListBySubscriptionPager: newTracker[azfake.PagerResponder[armcompute.ProximityPlacementGroupsClientListBySubscriptionResponse]](),
	}
}

// ProximityPlacementGroupsServerTransport connects instances of armcompute.ProximityPlacementGroupsClient to instances of ProximityPlacementGroupsServer.
// Don't use this type directly, use NewProximityPlacementGroupsServerTransport instead.
type ProximityPlacementGroupsServerTransport struct {
	srv *ProximityPlacementGroupsServer
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armcompute.ProximityPlacementGroupsClientListByResourceGroupResponse]]
	newListBySubscriptionPager *tracker[azfake.PagerResponder[armcompute.ProximityPlacementGroupsClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for ProximityPlacementGroupsServerTransport.
func (p *ProximityPlacementGroupsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ProximityPlacementGroupsClient.CreateOrUpdate":
		resp, err = p.dispatchCreateOrUpdate(req)
	case "ProximityPlacementGroupsClient.Delete":
		resp, err = p.dispatchDelete(req)
	case "ProximityPlacementGroupsClient.Get":
		resp, err = p.dispatchGet(req)
	case "ProximityPlacementGroupsClient.NewListByResourceGroupPager":
		resp, err = p.dispatchNewListByResourceGroupPager(req)
	case "ProximityPlacementGroupsClient.NewListBySubscriptionPager":
		resp, err = p.dispatchNewListBySubscriptionPager(req)
	case "ProximityPlacementGroupsClient.Update":
		resp, err = p.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *ProximityPlacementGroupsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/proximityPlacementGroups/(?P<proximityPlacementGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcompute.ProximityPlacementGroup](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	proximityPlacementGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("proximityPlacementGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.CreateOrUpdate(req.Context(), resourceGroupNameUnescaped, proximityPlacementGroupNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ProximityPlacementGroup, req)
	if err != nil {		return nil, err
	}
	return resp, nil
}

func (p *ProximityPlacementGroupsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if p.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/proximityPlacementGroups/(?P<proximityPlacementGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	proximityPlacementGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("proximityPlacementGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Delete(req.Context(), resourceGroupNameUnescaped, proximityPlacementGroupNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {		return nil, err
	}
	return resp, nil
}

func (p *ProximityPlacementGroupsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/proximityPlacementGroups/(?P<proximityPlacementGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	proximityPlacementGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("proximityPlacementGroupName")])
	if err != nil {
		return nil, err
	}
	includeColocationStatusUnescaped, err := url.QueryUnescape(qp.Get("includeColocationStatus"))
	if err != nil {
		return nil, err
	}
	includeColocationStatusParam := getOptional(includeColocationStatusUnescaped)
	var options *armcompute.ProximityPlacementGroupsClientGetOptions
	if includeColocationStatusParam != nil {
		options = &armcompute.ProximityPlacementGroupsClientGetOptions{
			IncludeColocationStatus: includeColocationStatusParam,
		}
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameUnescaped, proximityPlacementGroupNameUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ProximityPlacementGroup, req)
	if err != nil {		return nil, err
	}
	return resp, nil
}

func (p *ProximityPlacementGroupsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := p.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/proximityPlacementGroups`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
resp := p.srv.NewListByResourceGroupPager(resourceGroupNameUnescaped, nil)
		newListByResourceGroupPager = &resp
		p.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armcompute.ProximityPlacementGroupsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		p.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (p *ProximityPlacementGroupsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := p.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/proximityPlacementGroups`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
resp := p.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		p.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armcompute.ProximityPlacementGroupsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		p.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (p *ProximityPlacementGroupsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/proximityPlacementGroups/(?P<proximityPlacementGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcompute.ProximityPlacementGroupUpdate](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	proximityPlacementGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("proximityPlacementGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Update(req.Context(), resourceGroupNameUnescaped, proximityPlacementGroupNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ProximityPlacementGroup, req)
	if err != nil {		return nil, err
	}
	return resp, nil
}

