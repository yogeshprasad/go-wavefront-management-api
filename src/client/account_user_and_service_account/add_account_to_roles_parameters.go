// Code generated by go-swagger; DO NOT EDIT.

package account_user_and_service_account

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

// NewAddAccountToRolesParams creates a new AddAccountToRolesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddAccountToRolesParams() *AddAccountToRolesParams {
	return &AddAccountToRolesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddAccountToRolesParamsWithTimeout creates a new AddAccountToRolesParams object
// with the ability to set a timeout on a request.
func NewAddAccountToRolesParamsWithTimeout(timeout time.Duration) *AddAccountToRolesParams {
	return &AddAccountToRolesParams{
		timeout: timeout,
	}
}

// NewAddAccountToRolesParamsWithContext creates a new AddAccountToRolesParams object
// with the ability to set a context for a request.
func NewAddAccountToRolesParamsWithContext(ctx context.Context) *AddAccountToRolesParams {
	return &AddAccountToRolesParams{
		Context: ctx,
	}
}

// NewAddAccountToRolesParamsWithHTTPClient creates a new AddAccountToRolesParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddAccountToRolesParamsWithHTTPClient(client *http.Client) *AddAccountToRolesParams {
	return &AddAccountToRolesParams{
		HTTPClient: client,
	}
}

/*
AddAccountToRolesParams contains all the parameters to send to the API endpoint

	for the add account to roles operation.

	Typically these are written to a http.Request.
*/
type AddAccountToRolesParams struct {

	/* Body.

	   The list of roles that should be added to the account
	*/
	Body []string

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add account to roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAccountToRolesParams) WithDefaults() *AddAccountToRolesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add account to roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAccountToRolesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add account to roles params
func (o *AddAccountToRolesParams) WithTimeout(timeout time.Duration) *AddAccountToRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add account to roles params
func (o *AddAccountToRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add account to roles params
func (o *AddAccountToRolesParams) WithContext(ctx context.Context) *AddAccountToRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add account to roles params
func (o *AddAccountToRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add account to roles params
func (o *AddAccountToRolesParams) WithHTTPClient(client *http.Client) *AddAccountToRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add account to roles params
func (o *AddAccountToRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add account to roles params
func (o *AddAccountToRolesParams) WithBody(body []string) *AddAccountToRolesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add account to roles params
func (o *AddAccountToRolesParams) SetBody(body []string) {
	o.Body = body
}

// WithID adds the id to the add account to roles params
func (o *AddAccountToRolesParams) WithID(id string) *AddAccountToRolesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the add account to roles params
func (o *AddAccountToRolesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AddAccountToRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
