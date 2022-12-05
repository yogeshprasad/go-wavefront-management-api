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

// NewSearchUserForFacetParams creates a new SearchUserForFacetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchUserForFacetParams() *SearchUserForFacetParams {
	return &SearchUserForFacetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchUserForFacetParamsWithTimeout creates a new SearchUserForFacetParams object
// with the ability to set a timeout on a request.
func NewSearchUserForFacetParamsWithTimeout(timeout time.Duration) *SearchUserForFacetParams {
	return &SearchUserForFacetParams{
		timeout: timeout,
	}
}

// NewSearchUserForFacetParamsWithContext creates a new SearchUserForFacetParams object
// with the ability to set a context for a request.
func NewSearchUserForFacetParamsWithContext(ctx context.Context) *SearchUserForFacetParams {
	return &SearchUserForFacetParams{
		Context: ctx,
	}
}

// NewSearchUserForFacetParamsWithHTTPClient creates a new SearchUserForFacetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchUserForFacetParamsWithHTTPClient(client *http.Client) *SearchUserForFacetParams {
	return &SearchUserForFacetParams{
		HTTPClient: client,
	}
}

/*
SearchUserForFacetParams contains all the parameters to send to the API endpoint

	for the search user for facet operation.

	Typically these are written to a http.Request.
*/
type SearchUserForFacetParams struct {

	// Body.
	Body *models.FacetSearchRequestContainer

	// Facet.
	Facet string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search user for facet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchUserForFacetParams) WithDefaults() *SearchUserForFacetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search user for facet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchUserForFacetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search user for facet params
func (o *SearchUserForFacetParams) WithTimeout(timeout time.Duration) *SearchUserForFacetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search user for facet params
func (o *SearchUserForFacetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search user for facet params
func (o *SearchUserForFacetParams) WithContext(ctx context.Context) *SearchUserForFacetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search user for facet params
func (o *SearchUserForFacetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search user for facet params
func (o *SearchUserForFacetParams) WithHTTPClient(client *http.Client) *SearchUserForFacetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search user for facet params
func (o *SearchUserForFacetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search user for facet params
func (o *SearchUserForFacetParams) WithBody(body *models.FacetSearchRequestContainer) *SearchUserForFacetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search user for facet params
func (o *SearchUserForFacetParams) SetBody(body *models.FacetSearchRequestContainer) {
	o.Body = body
}

// WithFacet adds the facet to the search user for facet params
func (o *SearchUserForFacetParams) WithFacet(facet string) *SearchUserForFacetParams {
	o.SetFacet(facet)
	return o
}

// SetFacet adds the facet to the search user for facet params
func (o *SearchUserForFacetParams) SetFacet(facet string) {
	o.Facet = facet
}

// WriteToRequest writes these params to a swagger request
func (o *SearchUserForFacetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
