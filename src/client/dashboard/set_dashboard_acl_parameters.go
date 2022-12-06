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

// NewSetDashboardACLParams creates a new SetDashboardACLParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetDashboardACLParams() *SetDashboardACLParams {
	return &SetDashboardACLParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetDashboardACLParamsWithTimeout creates a new SetDashboardACLParams object
// with the ability to set a timeout on a request.
func NewSetDashboardACLParamsWithTimeout(timeout time.Duration) *SetDashboardACLParams {
	return &SetDashboardACLParams{
		timeout: timeout,
	}
}

// NewSetDashboardACLParamsWithContext creates a new SetDashboardACLParams object
// with the ability to set a context for a request.
func NewSetDashboardACLParamsWithContext(ctx context.Context) *SetDashboardACLParams {
	return &SetDashboardACLParams{
		Context: ctx,
	}
}

// NewSetDashboardACLParamsWithHTTPClient creates a new SetDashboardACLParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetDashboardACLParamsWithHTTPClient(client *http.Client) *SetDashboardACLParams {
	return &SetDashboardACLParams{
		HTTPClient: client,
	}
}

/*
SetDashboardACLParams contains all the parameters to send to the API endpoint

	for the set dashboard Acl operation.

	Typically these are written to a http.Request.
*/
type SetDashboardACLParams struct {

	// Body.
	Body []*models.AccessControlListWriteDTO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set dashboard Acl params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetDashboardACLParams) WithDefaults() *SetDashboardACLParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set dashboard Acl params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetDashboardACLParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set dashboard Acl params
func (o *SetDashboardACLParams) WithTimeout(timeout time.Duration) *SetDashboardACLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set dashboard Acl params
func (o *SetDashboardACLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set dashboard Acl params
func (o *SetDashboardACLParams) WithContext(ctx context.Context) *SetDashboardACLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set dashboard Acl params
func (o *SetDashboardACLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set dashboard Acl params
func (o *SetDashboardACLParams) WithHTTPClient(client *http.Client) *SetDashboardACLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set dashboard Acl params
func (o *SetDashboardACLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set dashboard Acl params
func (o *SetDashboardACLParams) WithBody(body []*models.AccessControlListWriteDTO) *SetDashboardACLParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set dashboard Acl params
func (o *SetDashboardACLParams) SetBody(body []*models.AccessControlListWriteDTO) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SetDashboardACLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
