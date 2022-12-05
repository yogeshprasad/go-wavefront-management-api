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

// NewSearchRegisteredQueryDeletedForFacetParams creates a new SearchRegisteredQueryDeletedForFacetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchRegisteredQueryDeletedForFacetParams() *SearchRegisteredQueryDeletedForFacetParams {
	return &SearchRegisteredQueryDeletedForFacetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchRegisteredQueryDeletedForFacetParamsWithTimeout creates a new SearchRegisteredQueryDeletedForFacetParams object
// with the ability to set a timeout on a request.
func NewSearchRegisteredQueryDeletedForFacetParamsWithTimeout(timeout time.Duration) *SearchRegisteredQueryDeletedForFacetParams {
	return &SearchRegisteredQueryDeletedForFacetParams{
		timeout: timeout,
	}
}

// NewSearchRegisteredQueryDeletedForFacetParamsWithContext creates a new SearchRegisteredQueryDeletedForFacetParams object
// with the ability to set a context for a request.
func NewSearchRegisteredQueryDeletedForFacetParamsWithContext(ctx context.Context) *SearchRegisteredQueryDeletedForFacetParams {
	return &SearchRegisteredQueryDeletedForFacetParams{
		Context: ctx,
	}
}

// NewSearchRegisteredQueryDeletedForFacetParamsWithHTTPClient creates a new SearchRegisteredQueryDeletedForFacetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchRegisteredQueryDeletedForFacetParamsWithHTTPClient(client *http.Client) *SearchRegisteredQueryDeletedForFacetParams {
	return &SearchRegisteredQueryDeletedForFacetParams{
		HTTPClient: client,
	}
}

/*
SearchRegisteredQueryDeletedForFacetParams contains all the parameters to send to the API endpoint

	for the search registered query deleted for facet operation.

	Typically these are written to a http.Request.
*/
type SearchRegisteredQueryDeletedForFacetParams struct {

	// Body.
	Body *models.FacetSearchRequestContainer

	// Facet.
	Facet string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search registered query deleted for facet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchRegisteredQueryDeletedForFacetParams) WithDefaults() *SearchRegisteredQueryDeletedForFacetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search registered query deleted for facet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchRegisteredQueryDeletedForFacetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search registered query deleted for facet params
func (o *SearchRegisteredQueryDeletedForFacetParams) WithTimeout(timeout time.Duration) *SearchRegisteredQueryDeletedForFacetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search registered query deleted for facet params
func (o *SearchRegisteredQueryDeletedForFacetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search registered query deleted for facet params
func (o *SearchRegisteredQueryDeletedForFacetParams) WithContext(ctx context.Context) *SearchRegisteredQueryDeletedForFacetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search registered query deleted for facet params
func (o *SearchRegisteredQueryDeletedForFacetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search registered query deleted for facet params
func (o *SearchRegisteredQueryDeletedForFacetParams) WithHTTPClient(client *http.Client) *SearchRegisteredQueryDeletedForFacetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search registered query deleted for facet params
func (o *SearchRegisteredQueryDeletedForFacetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search registered query deleted for facet params
func (o *SearchRegisteredQueryDeletedForFacetParams) WithBody(body *models.FacetSearchRequestContainer) *SearchRegisteredQueryDeletedForFacetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search registered query deleted for facet params
func (o *SearchRegisteredQueryDeletedForFacetParams) SetBody(body *models.FacetSearchRequestContainer) {
	o.Body = body
}

// WithFacet adds the facet to the search registered query deleted for facet params
func (o *SearchRegisteredQueryDeletedForFacetParams) WithFacet(facet string) *SearchRegisteredQueryDeletedForFacetParams {
	o.SetFacet(facet)
	return o
}

// SetFacet adds the facet to the search registered query deleted for facet params
func (o *SearchRegisteredQueryDeletedForFacetParams) SetFacet(facet string) {
	o.Facet = facet
}

// WriteToRequest writes these params to a swagger request
func (o *SearchRegisteredQueryDeletedForFacetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
