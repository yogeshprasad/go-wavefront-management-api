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

// NewDeleteSavedSearchParams creates a new DeleteSavedSearchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSavedSearchParams() *DeleteSavedSearchParams {
	return &DeleteSavedSearchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSavedSearchParamsWithTimeout creates a new DeleteSavedSearchParams object
// with the ability to set a timeout on a request.
func NewDeleteSavedSearchParamsWithTimeout(timeout time.Duration) *DeleteSavedSearchParams {
	return &DeleteSavedSearchParams{
		timeout: timeout,
	}
}

// NewDeleteSavedSearchParamsWithContext creates a new DeleteSavedSearchParams object
// with the ability to set a context for a request.
func NewDeleteSavedSearchParamsWithContext(ctx context.Context) *DeleteSavedSearchParams {
	return &DeleteSavedSearchParams{
		Context: ctx,
	}
}

// NewDeleteSavedSearchParamsWithHTTPClient creates a new DeleteSavedSearchParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSavedSearchParamsWithHTTPClient(client *http.Client) *DeleteSavedSearchParams {
	return &DeleteSavedSearchParams{
		HTTPClient: client,
	}
}

/*
DeleteSavedSearchParams contains all the parameters to send to the API endpoint

	for the delete saved search operation.

	Typically these are written to a http.Request.
*/
type DeleteSavedSearchParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete saved search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSavedSearchParams) WithDefaults() *DeleteSavedSearchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete saved search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSavedSearchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete saved search params
func (o *DeleteSavedSearchParams) WithTimeout(timeout time.Duration) *DeleteSavedSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete saved search params
func (o *DeleteSavedSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete saved search params
func (o *DeleteSavedSearchParams) WithContext(ctx context.Context) *DeleteSavedSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete saved search params
func (o *DeleteSavedSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete saved search params
func (o *DeleteSavedSearchParams) WithHTTPClient(client *http.Client) *DeleteSavedSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete saved search params
func (o *DeleteSavedSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete saved search params
func (o *DeleteSavedSearchParams) WithID(id string) *DeleteSavedSearchParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete saved search params
func (o *DeleteSavedSearchParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSavedSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
