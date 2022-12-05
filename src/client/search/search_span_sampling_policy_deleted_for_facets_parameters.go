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

// NewSearchSpanSamplingPolicyDeletedForFacetsParams creates a new SearchSpanSamplingPolicyDeletedForFacetsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchSpanSamplingPolicyDeletedForFacetsParams() *SearchSpanSamplingPolicyDeletedForFacetsParams {
	return &SearchSpanSamplingPolicyDeletedForFacetsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchSpanSamplingPolicyDeletedForFacetsParamsWithTimeout creates a new SearchSpanSamplingPolicyDeletedForFacetsParams object
// with the ability to set a timeout on a request.
func NewSearchSpanSamplingPolicyDeletedForFacetsParamsWithTimeout(timeout time.Duration) *SearchSpanSamplingPolicyDeletedForFacetsParams {
	return &SearchSpanSamplingPolicyDeletedForFacetsParams{
		timeout: timeout,
	}
}

// NewSearchSpanSamplingPolicyDeletedForFacetsParamsWithContext creates a new SearchSpanSamplingPolicyDeletedForFacetsParams object
// with the ability to set a context for a request.
func NewSearchSpanSamplingPolicyDeletedForFacetsParamsWithContext(ctx context.Context) *SearchSpanSamplingPolicyDeletedForFacetsParams {
	return &SearchSpanSamplingPolicyDeletedForFacetsParams{
		Context: ctx,
	}
}

// NewSearchSpanSamplingPolicyDeletedForFacetsParamsWithHTTPClient creates a new SearchSpanSamplingPolicyDeletedForFacetsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchSpanSamplingPolicyDeletedForFacetsParamsWithHTTPClient(client *http.Client) *SearchSpanSamplingPolicyDeletedForFacetsParams {
	return &SearchSpanSamplingPolicyDeletedForFacetsParams{
		HTTPClient: client,
	}
}

/*
SearchSpanSamplingPolicyDeletedForFacetsParams contains all the parameters to send to the API endpoint

	for the search span sampling policy deleted for facets operation.

	Typically these are written to a http.Request.
*/
type SearchSpanSamplingPolicyDeletedForFacetsParams struct {

	// Body.
	Body *models.FacetsSearchRequestContainer

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search span sampling policy deleted for facets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchSpanSamplingPolicyDeletedForFacetsParams) WithDefaults() *SearchSpanSamplingPolicyDeletedForFacetsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search span sampling policy deleted for facets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchSpanSamplingPolicyDeletedForFacetsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search span sampling policy deleted for facets params
func (o *SearchSpanSamplingPolicyDeletedForFacetsParams) WithTimeout(timeout time.Duration) *SearchSpanSamplingPolicyDeletedForFacetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search span sampling policy deleted for facets params
func (o *SearchSpanSamplingPolicyDeletedForFacetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search span sampling policy deleted for facets params
func (o *SearchSpanSamplingPolicyDeletedForFacetsParams) WithContext(ctx context.Context) *SearchSpanSamplingPolicyDeletedForFacetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search span sampling policy deleted for facets params
func (o *SearchSpanSamplingPolicyDeletedForFacetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search span sampling policy deleted for facets params
func (o *SearchSpanSamplingPolicyDeletedForFacetsParams) WithHTTPClient(client *http.Client) *SearchSpanSamplingPolicyDeletedForFacetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search span sampling policy deleted for facets params
func (o *SearchSpanSamplingPolicyDeletedForFacetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search span sampling policy deleted for facets params
func (o *SearchSpanSamplingPolicyDeletedForFacetsParams) WithBody(body *models.FacetsSearchRequestContainer) *SearchSpanSamplingPolicyDeletedForFacetsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search span sampling policy deleted for facets params
func (o *SearchSpanSamplingPolicyDeletedForFacetsParams) SetBody(body *models.FacetsSearchRequestContainer) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SearchSpanSamplingPolicyDeletedForFacetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
