// Code generated by go-swagger; DO NOT EDIT.

package role

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

// NewDeleteRoleParams creates a new DeleteRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRoleParams() *DeleteRoleParams {
	return &DeleteRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRoleParamsWithTimeout creates a new DeleteRoleParams object
// with the ability to set a timeout on a request.
func NewDeleteRoleParamsWithTimeout(timeout time.Duration) *DeleteRoleParams {
	return &DeleteRoleParams{
		timeout: timeout,
	}
}

// NewDeleteRoleParamsWithContext creates a new DeleteRoleParams object
// with the ability to set a context for a request.
func NewDeleteRoleParamsWithContext(ctx context.Context) *DeleteRoleParams {
	return &DeleteRoleParams{
		Context: ctx,
	}
}

// NewDeleteRoleParamsWithHTTPClient creates a new DeleteRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRoleParamsWithHTTPClient(client *http.Client) *DeleteRoleParams {
	return &DeleteRoleParams{
		HTTPClient: client,
	}
}

/*
DeleteRoleParams contains all the parameters to send to the API endpoint

	for the delete role operation.

	Typically these are written to a http.Request.
*/
type DeleteRoleParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRoleParams) WithDefaults() *DeleteRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete role params
func (o *DeleteRoleParams) WithTimeout(timeout time.Duration) *DeleteRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete role params
func (o *DeleteRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete role params
func (o *DeleteRoleParams) WithContext(ctx context.Context) *DeleteRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete role params
func (o *DeleteRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete role params
func (o *DeleteRoleParams) WithHTTPClient(client *http.Client) *DeleteRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete role params
func (o *DeleteRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete role params
func (o *DeleteRoleParams) WithID(id string) *DeleteRoleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete role params
func (o *DeleteRoleParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
