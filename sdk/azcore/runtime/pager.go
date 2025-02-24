//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
)

// PagingHandler contains the required data for constructing a Pager.
type PagingHandler[T any] struct {
	// More returns a boolean indicating if there are more pages to fetch.
	// It uses the provided page to make the determination.
	More func(T) bool

	// Fetcher fetches the first and subsequent pages.
	Fetcher func(context.Context, *T) (T, error)

	// Tracer contains the Tracer from the client that's creating the Pager.
	Tracer tracing.Tracer
}

// Pager provides operations for iterating over paged responses.
type Pager[T any] struct {
	current   *T
	handler   PagingHandler[T]
	tracer    tracing.Tracer
	firstPage bool
}

// NewPager creates an instance of Pager using the specified PagingHandler.
// Pass a non-nil T for firstPage if the first page has already been retrieved.
func NewPager[T any](handler PagingHandler[T]) *Pager[T] {
	return &Pager[T]{
		handler:   handler,
		tracer:    handler.Tracer,
		firstPage: true,
	}
}

// More returns true if there are more pages to retrieve.
func (p *Pager[T]) More() bool {
	if p.current != nil {
		return p.handler.More(*p.current)
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *Pager[T]) NextPage(ctx context.Context) (T, error) {
	var resp T
	var err error
	if p.current != nil {
		if p.firstPage {
			// we get here if it's an LRO-pager, we already have the first page
			p.firstPage = false
			return *p.current, nil
		} else if !p.handler.More(*p.current) {
			return *new(T), errors.New("no more pages")
		}
		ctx, endSpan := StartSpan(ctx, fmt.Sprintf("%s.NextPage", shortenTypeName(reflect.TypeOf(*p).Name())), p.tracer, nil)
		defer endSpan(err)
		resp, err = p.handler.Fetcher(ctx, p.current)
	} else {
		// non-LRO case, first page
		p.firstPage = false
		ctx, endSpan := StartSpan(ctx, fmt.Sprintf("%s.NextPage", shortenTypeName(reflect.TypeOf(*p).Name())), p.tracer, nil)
		defer endSpan(err)
		resp, err = p.handler.Fetcher(ctx, nil)
	}
	if err != nil {
		return *new(T), err
	}
	p.current = &resp
	return *p.current, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for Pager[T].
func (p *Pager[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &p.current)
}

// FetcherForNextLink is a helper containing boilerplate code to simplify creating a PagingHandler[T].Fetcher from a next link URL.
func FetcherForNextLink(ctx context.Context, pl Pipeline, nextLink string, createReq func(context.Context) (*policy.Request, error)) (*http.Response, error) {
	var req *policy.Request
	var err error
	if nextLink == "" {
		req, err = createReq(ctx)
	} else if nextLink, err = EncodeQueryParams(nextLink); err == nil {
		req, err = NewRequest(ctx, http.MethodGet, nextLink)
	}
	if err != nil {
		return nil, err
	}
	resp, err := pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !HasStatusCode(resp, http.StatusOK) {
		return nil, NewResponseError(resp)
	}
	return resp, nil
}
