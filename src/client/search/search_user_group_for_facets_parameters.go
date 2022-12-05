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

// NewSearchUserGroupForFacetsParams creates a new SearchUserGroupForFacetsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchUserGroupForFacetsParams() *SearchUserGroupForFacetsParams {
	return &SearchUserGroupForFacetsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchUserGroupForFacetsParamsWithTimeout creates a new SearchUserGroupForFacetsParams object
// with the ability to set a timeout on a request.
func NewSearchUserGroupForFacetsParamsWithTimeout(timeout time.Duration) *SearchUserGroupForFacetsParams {
	return &SearchUserGroupForFacetsParams{
		timeout: timeout,
	}
}

// NewSearchUserGroupForFacetsParamsWithContext creates a new SearchUserGroupForFacetsParams object
// with the ability to set a context for a request.
func NewSearchUserGroupForFacetsParamsWithContext(ctx context.Context) *SearchUserGroupForFacetsParams {
	return &SearchUserGroupForFacetsParams{
		Context: ctx,
	}
}

// NewSearchUserGroupForFacetsParamsWithHTTPClient creates a new SearchUserGroupForFacetsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchUserGroupForFacetsParamsWithHTTPClient(client *http.Client) *SearchUserGroupForFacetsParams {
	return &SearchUserGroupForFacetsParams{
		HTTPClient: client,
	}
}

/*
SearchUserGroupForFacetsParams contains all the parameters to send to the API endpoint

	for the search user group for facets operation.

	Typically these are written to a http.Request.
*/
type SearchUserGroupForFacetsParams struct {

	// Body.
	Body *models.FacetsSearchRequestContainer

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search user group for facets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchUserGroupForFacetsParams) WithDefaults() *SearchUserGroupForFacetsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search user group for facets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchUserGroupForFacetsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search user group for facets params
func (o *SearchUserGroupForFacetsParams) WithTimeout(timeout time.Duration) *SearchUserGroupForFacetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search user group for facets params
func (o *SearchUserGroupForFacetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search user group for facets params
func (o *SearchUserGroupForFacetsParams) WithContext(ctx context.Context) *SearchUserGroupForFacetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search user group for facets params
func (o *SearchUserGroupForFacetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search user group for facets params
func (o *SearchUserGroupForFacetsParams) WithHTTPClient(client *http.Client) *SearchUserGroupForFacetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search user group for facets params
func (o *SearchUserGroupForFacetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search user group for facets params
func (o *SearchUserGroupForFacetsParams) WithBody(body *models.FacetsSearchRequestContainer) *SearchUserGroupForFacetsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search user group for facets params
func (o *SearchUserGroupForFacetsParams) SetBody(body *models.FacetsSearchRequestContainer) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SearchUserGroupForFacetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
