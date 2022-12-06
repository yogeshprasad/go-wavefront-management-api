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

// NewSearchRegisteredQueryDeletedForFacetsParams creates a new SearchRegisteredQueryDeletedForFacetsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchRegisteredQueryDeletedForFacetsParams() *SearchRegisteredQueryDeletedForFacetsParams {
	return &SearchRegisteredQueryDeletedForFacetsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchRegisteredQueryDeletedForFacetsParamsWithTimeout creates a new SearchRegisteredQueryDeletedForFacetsParams object
// with the ability to set a timeout on a request.
func NewSearchRegisteredQueryDeletedForFacetsParamsWithTimeout(timeout time.Duration) *SearchRegisteredQueryDeletedForFacetsParams {
	return &SearchRegisteredQueryDeletedForFacetsParams{
		timeout: timeout,
	}
}

// NewSearchRegisteredQueryDeletedForFacetsParamsWithContext creates a new SearchRegisteredQueryDeletedForFacetsParams object
// with the ability to set a context for a request.
func NewSearchRegisteredQueryDeletedForFacetsParamsWithContext(ctx context.Context) *SearchRegisteredQueryDeletedForFacetsParams {
	return &SearchRegisteredQueryDeletedForFacetsParams{
		Context: ctx,
	}
}

// NewSearchRegisteredQueryDeletedForFacetsParamsWithHTTPClient creates a new SearchRegisteredQueryDeletedForFacetsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchRegisteredQueryDeletedForFacetsParamsWithHTTPClient(client *http.Client) *SearchRegisteredQueryDeletedForFacetsParams {
	return &SearchRegisteredQueryDeletedForFacetsParams{
		HTTPClient: client,
	}
}

/*
SearchRegisteredQueryDeletedForFacetsParams contains all the parameters to send to the API endpoint

	for the search registered query deleted for facets operation.

	Typically these are written to a http.Request.
*/
type SearchRegisteredQueryDeletedForFacetsParams struct {

	// Body.
	Body *models.FacetsSearchRequestContainer

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search registered query deleted for facets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchRegisteredQueryDeletedForFacetsParams) WithDefaults() *SearchRegisteredQueryDeletedForFacetsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search registered query deleted for facets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchRegisteredQueryDeletedForFacetsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search registered query deleted for facets params
func (o *SearchRegisteredQueryDeletedForFacetsParams) WithTimeout(timeout time.Duration) *SearchRegisteredQueryDeletedForFacetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search registered query deleted for facets params
func (o *SearchRegisteredQueryDeletedForFacetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search registered query deleted for facets params
func (o *SearchRegisteredQueryDeletedForFacetsParams) WithContext(ctx context.Context) *SearchRegisteredQueryDeletedForFacetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search registered query deleted for facets params
func (o *SearchRegisteredQueryDeletedForFacetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search registered query deleted for facets params
func (o *SearchRegisteredQueryDeletedForFacetsParams) WithHTTPClient(client *http.Client) *SearchRegisteredQueryDeletedForFacetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search registered query deleted for facets params
func (o *SearchRegisteredQueryDeletedForFacetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search registered query deleted for facets params
func (o *SearchRegisteredQueryDeletedForFacetsParams) WithBody(body *models.FacetsSearchRequestContainer) *SearchRegisteredQueryDeletedForFacetsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search registered query deleted for facets params
func (o *SearchRegisteredQueryDeletedForFacetsParams) SetBody(body *models.FacetsSearchRequestContainer) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SearchRegisteredQueryDeletedForFacetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
