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

// NewSearchSpanSamplingPolicyForFacetsParams creates a new SearchSpanSamplingPolicyForFacetsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchSpanSamplingPolicyForFacetsParams() *SearchSpanSamplingPolicyForFacetsParams {
	return &SearchSpanSamplingPolicyForFacetsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchSpanSamplingPolicyForFacetsParamsWithTimeout creates a new SearchSpanSamplingPolicyForFacetsParams object
// with the ability to set a timeout on a request.
func NewSearchSpanSamplingPolicyForFacetsParamsWithTimeout(timeout time.Duration) *SearchSpanSamplingPolicyForFacetsParams {
	return &SearchSpanSamplingPolicyForFacetsParams{
		timeout: timeout,
	}
}

// NewSearchSpanSamplingPolicyForFacetsParamsWithContext creates a new SearchSpanSamplingPolicyForFacetsParams object
// with the ability to set a context for a request.
func NewSearchSpanSamplingPolicyForFacetsParamsWithContext(ctx context.Context) *SearchSpanSamplingPolicyForFacetsParams {
	return &SearchSpanSamplingPolicyForFacetsParams{
		Context: ctx,
	}
}

// NewSearchSpanSamplingPolicyForFacetsParamsWithHTTPClient creates a new SearchSpanSamplingPolicyForFacetsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchSpanSamplingPolicyForFacetsParamsWithHTTPClient(client *http.Client) *SearchSpanSamplingPolicyForFacetsParams {
	return &SearchSpanSamplingPolicyForFacetsParams{
		HTTPClient: client,
	}
}

/*
SearchSpanSamplingPolicyForFacetsParams contains all the parameters to send to the API endpoint

	for the search span sampling policy for facets operation.

	Typically these are written to a http.Request.
*/
type SearchSpanSamplingPolicyForFacetsParams struct {

	// Body.
	Body *models.FacetsSearchRequestContainer

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search span sampling policy for facets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchSpanSamplingPolicyForFacetsParams) WithDefaults() *SearchSpanSamplingPolicyForFacetsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search span sampling policy for facets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchSpanSamplingPolicyForFacetsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search span sampling policy for facets params
func (o *SearchSpanSamplingPolicyForFacetsParams) WithTimeout(timeout time.Duration) *SearchSpanSamplingPolicyForFacetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search span sampling policy for facets params
func (o *SearchSpanSamplingPolicyForFacetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search span sampling policy for facets params
func (o *SearchSpanSamplingPolicyForFacetsParams) WithContext(ctx context.Context) *SearchSpanSamplingPolicyForFacetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search span sampling policy for facets params
func (o *SearchSpanSamplingPolicyForFacetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search span sampling policy for facets params
func (o *SearchSpanSamplingPolicyForFacetsParams) WithHTTPClient(client *http.Client) *SearchSpanSamplingPolicyForFacetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search span sampling policy for facets params
func (o *SearchSpanSamplingPolicyForFacetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search span sampling policy for facets params
func (o *SearchSpanSamplingPolicyForFacetsParams) WithBody(body *models.FacetsSearchRequestContainer) *SearchSpanSamplingPolicyForFacetsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search span sampling policy for facets params
func (o *SearchSpanSamplingPolicyForFacetsParams) SetBody(body *models.FacetsSearchRequestContainer) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SearchSpanSamplingPolicyForFacetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
