// Code generated by go-swagger; DO NOT EDIT.

package dashboard

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

// NewAddDashboardAccessParams creates a new AddDashboardAccessParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddDashboardAccessParams() *AddDashboardAccessParams {
	return &AddDashboardAccessParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddDashboardAccessParamsWithTimeout creates a new AddDashboardAccessParams object
// with the ability to set a timeout on a request.
func NewAddDashboardAccessParamsWithTimeout(timeout time.Duration) *AddDashboardAccessParams {
	return &AddDashboardAccessParams{
		timeout: timeout,
	}
}

// NewAddDashboardAccessParamsWithContext creates a new AddDashboardAccessParams object
// with the ability to set a context for a request.
func NewAddDashboardAccessParamsWithContext(ctx context.Context) *AddDashboardAccessParams {
	return &AddDashboardAccessParams{
		Context: ctx,
	}
}

// NewAddDashboardAccessParamsWithHTTPClient creates a new AddDashboardAccessParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddDashboardAccessParamsWithHTTPClient(client *http.Client) *AddDashboardAccessParams {
	return &AddDashboardAccessParams{
		HTTPClient: client,
	}
}

/*
AddDashboardAccessParams contains all the parameters to send to the API endpoint

	for the add dashboard access operation.

	Typically these are written to a http.Request.
*/
type AddDashboardAccessParams struct {

	// Body.
	Body []*models.AccessControlListWriteDTO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add dashboard access params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddDashboardAccessParams) WithDefaults() *AddDashboardAccessParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add dashboard access params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddDashboardAccessParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add dashboard access params
func (o *AddDashboardAccessParams) WithTimeout(timeout time.Duration) *AddDashboardAccessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add dashboard access params
func (o *AddDashboardAccessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add dashboard access params
func (o *AddDashboardAccessParams) WithContext(ctx context.Context) *AddDashboardAccessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add dashboard access params
func (o *AddDashboardAccessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add dashboard access params
func (o *AddDashboardAccessParams) WithHTTPClient(client *http.Client) *AddDashboardAccessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add dashboard access params
func (o *AddDashboardAccessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add dashboard access params
func (o *AddDashboardAccessParams) WithBody(body []*models.AccessControlListWriteDTO) *AddDashboardAccessParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add dashboard access params
func (o *AddDashboardAccessParams) SetBody(body []*models.AccessControlListWriteDTO) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddDashboardAccessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
