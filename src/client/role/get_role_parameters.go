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

// NewGetRoleParams creates a new GetRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRoleParams() *GetRoleParams {
	return &GetRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoleParamsWithTimeout creates a new GetRoleParams object
// with the ability to set a timeout on a request.
func NewGetRoleParamsWithTimeout(timeout time.Duration) *GetRoleParams {
	return &GetRoleParams{
		timeout: timeout,
	}
}

// NewGetRoleParamsWithContext creates a new GetRoleParams object
// with the ability to set a context for a request.
func NewGetRoleParamsWithContext(ctx context.Context) *GetRoleParams {
	return &GetRoleParams{
		Context: ctx,
	}
}

// NewGetRoleParamsWithHTTPClient creates a new GetRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRoleParamsWithHTTPClient(client *http.Client) *GetRoleParams {
	return &GetRoleParams{
		HTTPClient: client,
	}
}

/*
GetRoleParams contains all the parameters to send to the API endpoint

	for the get role operation.

	Typically these are written to a http.Request.
*/
type GetRoleParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoleParams) WithDefaults() *GetRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get role params
func (o *GetRoleParams) WithTimeout(timeout time.Duration) *GetRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get role params
func (o *GetRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get role params
func (o *GetRoleParams) WithContext(ctx context.Context) *GetRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get role params
func (o *GetRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get role params
func (o *GetRoleParams) WithHTTPClient(client *http.Client) *GetRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get role params
func (o *GetRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get role params
func (o *GetRoleParams) WithID(id string) *GetRoleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get role params
func (o *GetRoleParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
