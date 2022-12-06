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

// NewSearchTaggedSourceEntitiesParams creates a new SearchTaggedSourceEntitiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchTaggedSourceEntitiesParams() *SearchTaggedSourceEntitiesParams {
	return &SearchTaggedSourceEntitiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchTaggedSourceEntitiesParamsWithTimeout creates a new SearchTaggedSourceEntitiesParams object
// with the ability to set a timeout on a request.
func NewSearchTaggedSourceEntitiesParamsWithTimeout(timeout time.Duration) *SearchTaggedSourceEntitiesParams {
	return &SearchTaggedSourceEntitiesParams{
		timeout: timeout,
	}
}

// NewSearchTaggedSourceEntitiesParamsWithContext creates a new SearchTaggedSourceEntitiesParams object
// with the ability to set a context for a request.
func NewSearchTaggedSourceEntitiesParamsWithContext(ctx context.Context) *SearchTaggedSourceEntitiesParams {
	return &SearchTaggedSourceEntitiesParams{
		Context: ctx,
	}
}

// NewSearchTaggedSourceEntitiesParamsWithHTTPClient creates a new SearchTaggedSourceEntitiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchTaggedSourceEntitiesParamsWithHTTPClient(client *http.Client) *SearchTaggedSourceEntitiesParams {
	return &SearchTaggedSourceEntitiesParams{
		HTTPClient: client,
	}
}

/*
SearchTaggedSourceEntitiesParams contains all the parameters to send to the API endpoint

	for the search tagged source entities operation.

	Typically these are written to a http.Request.
*/
type SearchTaggedSourceEntitiesParams struct {

	// Body.
	Body *models.SourceSearchRequestContainer

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search tagged source entities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchTaggedSourceEntitiesParams) WithDefaults() *SearchTaggedSourceEntitiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search tagged source entities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchTaggedSourceEntitiesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search tagged source entities params
func (o *SearchTaggedSourceEntitiesParams) WithTimeout(timeout time.Duration) *SearchTaggedSourceEntitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search tagged source entities params
func (o *SearchTaggedSourceEntitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search tagged source entities params
func (o *SearchTaggedSourceEntitiesParams) WithContext(ctx context.Context) *SearchTaggedSourceEntitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search tagged source entities params
func (o *SearchTaggedSourceEntitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search tagged source entities params
func (o *SearchTaggedSourceEntitiesParams) WithHTTPClient(client *http.Client) *SearchTaggedSourceEntitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search tagged source entities params
func (o *SearchTaggedSourceEntitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search tagged source entities params
func (o *SearchTaggedSourceEntitiesParams) WithBody(body *models.SourceSearchRequestContainer) *SearchTaggedSourceEntitiesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search tagged source entities params
func (o *SearchTaggedSourceEntitiesParams) SetBody(body *models.SourceSearchRequestContainer) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SearchTaggedSourceEntitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
