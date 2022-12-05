// Code generated by go-swagger; DO NOT EDIT.

package saved_search

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
)

// NewGetSavedSearchParams creates a new GetSavedSearchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSavedSearchParams() *GetSavedSearchParams {
	return &GetSavedSearchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSavedSearchParamsWithTimeout creates a new GetSavedSearchParams object
// with the ability to set a timeout on a request.
func NewGetSavedSearchParamsWithTimeout(timeout time.Duration) *GetSavedSearchParams {
	return &GetSavedSearchParams{
		timeout: timeout,
	}
}

// NewGetSavedSearchParamsWithContext creates a new GetSavedSearchParams object
// with the ability to set a context for a request.
func NewGetSavedSearchParamsWithContext(ctx context.Context) *GetSavedSearchParams {
	return &GetSavedSearchParams{
		Context: ctx,
	}
}

// NewGetSavedSearchParamsWithHTTPClient creates a new GetSavedSearchParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSavedSearchParamsWithHTTPClient(client *http.Client) *GetSavedSearchParams {
	return &GetSavedSearchParams{
		HTTPClient: client,
	}
}

/*
GetSavedSearchParams contains all the parameters to send to the API endpoint

	for the get saved search operation.

	Typically these are written to a http.Request.
*/
type GetSavedSearchParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get saved search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSavedSearchParams) WithDefaults() *GetSavedSearchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get saved search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSavedSearchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get saved search params
func (o *GetSavedSearchParams) WithTimeout(timeout time.Duration) *GetSavedSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get saved search params
func (o *GetSavedSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get saved search params
func (o *GetSavedSearchParams) WithContext(ctx context.Context) *GetSavedSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get saved search params
func (o *GetSavedSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get saved search params
func (o *GetSavedSearchParams) WithHTTPClient(client *http.Client) *GetSavedSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get saved search params
func (o *GetSavedSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get saved search params
func (o *GetSavedSearchParams) WithID(id string) *GetSavedSearchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get saved search params
func (o *GetSavedSearchParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSavedSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
