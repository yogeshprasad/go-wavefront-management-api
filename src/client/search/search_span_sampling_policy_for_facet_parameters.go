// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"models"
)

// NewSearchSpanSamplingPolicyForFacetParams creates a new SearchSpanSamplingPolicyForFacetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchSpanSamplingPolicyForFacetParams() *SearchSpanSamplingPolicyForFacetParams {
	return &SearchSpanSamplingPolicyForFacetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchSpanSamplingPolicyForFacetParamsWithTimeout creates a new SearchSpanSamplingPolicyForFacetParams object
// with the ability to set a timeout on a request.
func NewSearchSpanSamplingPolicyForFacetParamsWithTimeout(timeout time.Duration) *SearchSpanSamplingPolicyForFacetParams {
	return &SearchSpanSamplingPolicyForFacetParams{
		timeout: timeout,
	}
}

// NewSearchSpanSamplingPolicyForFacetParamsWithContext creates a new SearchSpanSamplingPolicyForFacetParams object
// with the ability to set a context for a request.
func NewSearchSpanSamplingPolicyForFacetParamsWithContext(ctx context.Context) *SearchSpanSamplingPolicyForFacetParams {
	return &SearchSpanSamplingPolicyForFacetParams{
		Context: ctx,
	}
}

// NewSearchSpanSamplingPolicyForFacetParamsWithHTTPClient creates a new SearchSpanSamplingPolicyForFacetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchSpanSamplingPolicyForFacetParamsWithHTTPClient(client *http.Client) *SearchSpanSamplingPolicyForFacetParams {
	return &SearchSpanSamplingPolicyForFacetParams{
		HTTPClient: client,
	}
}

/*
SearchSpanSamplingPolicyForFacetParams contains all the parameters to send to the API endpoint

	for the search span sampling policy for facet operation.

	Typically these are written to a http.Request.
*/
type SearchSpanSamplingPolicyForFacetParams struct {

	// Body.
	Body *models.FacetSearchRequestContainer

	// Facet.
	Facet string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search span sampling policy for facet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchSpanSamplingPolicyForFacetParams) WithDefaults() *SearchSpanSamplingPolicyForFacetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search span sampling policy for facet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchSpanSamplingPolicyForFacetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search span sampling policy for facet params
func (o *SearchSpanSamplingPolicyForFacetParams) WithTimeout(timeout time.Duration) *SearchSpanSamplingPolicyForFacetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search span sampling policy for facet params
func (o *SearchSpanSamplingPolicyForFacetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search span sampling policy for facet params
func (o *SearchSpanSamplingPolicyForFacetParams) WithContext(ctx context.Context) *SearchSpanSamplingPolicyForFacetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search span sampling policy for facet params
func (o *SearchSpanSamplingPolicyForFacetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search span sampling policy for facet params
func (o *SearchSpanSamplingPolicyForFacetParams) WithHTTPClient(client *http.Client) *SearchSpanSamplingPolicyForFacetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search span sampling policy for facet params
func (o *SearchSpanSamplingPolicyForFacetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search span sampling policy for facet params
func (o *SearchSpanSamplingPolicyForFacetParams) WithBody(body *models.FacetSearchRequestContainer) *SearchSpanSamplingPolicyForFacetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search span sampling policy for facet params
func (o *SearchSpanSamplingPolicyForFacetParams) SetBody(body *models.FacetSearchRequestContainer) {
	o.Body = body
}

// WithFacet adds the facet to the search span sampling policy for facet params
func (o *SearchSpanSamplingPolicyForFacetParams) WithFacet(facet string) *SearchSpanSamplingPolicyForFacetParams {
	o.SetFacet(facet)
	return o
}

// SetFacet adds the facet to the search span sampling policy for facet params
func (o *SearchSpanSamplingPolicyForFacetParams) SetFacet(facet string) {
	o.Facet = facet
}

// WriteToRequest writes these params to a swagger request
func (o *SearchSpanSamplingPolicyForFacetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param facet
	if err := r.SetPathParam("facet", o.Facet); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
