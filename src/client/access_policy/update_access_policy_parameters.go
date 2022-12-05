// Code generated by go-swagger; DO NOT EDIT.

package access_policy

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

// NewUpdateAccessPolicyParams creates a new UpdateAccessPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAccessPolicyParams() *UpdateAccessPolicyParams {
	return &UpdateAccessPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAccessPolicyParamsWithTimeout creates a new UpdateAccessPolicyParams object
// with the ability to set a timeout on a request.
func NewUpdateAccessPolicyParamsWithTimeout(timeout time.Duration) *UpdateAccessPolicyParams {
	return &UpdateAccessPolicyParams{
		timeout: timeout,
	}
}

// NewUpdateAccessPolicyParamsWithContext creates a new UpdateAccessPolicyParams object
// with the ability to set a context for a request.
func NewUpdateAccessPolicyParamsWithContext(ctx context.Context) *UpdateAccessPolicyParams {
	return &UpdateAccessPolicyParams{
		Context: ctx,
	}
}

// NewUpdateAccessPolicyParamsWithHTTPClient creates a new UpdateAccessPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAccessPolicyParamsWithHTTPClient(client *http.Client) *UpdateAccessPolicyParams {
	return &UpdateAccessPolicyParams{
		HTTPClient: client,
	}
}

/*
UpdateAccessPolicyParams contains all the parameters to send to the API endpoint

	for the update access policy operation.

	Typically these are written to a http.Request.
*/
type UpdateAccessPolicyParams struct {

	/* Body.

	     Example Body:
	<pre>{
	 "policyRules": [{
	   "name": "rule name",
	   "description": "desc",
	   "action": "ALLOW",
	   "subnet": "12.148.72.0/23"
	 }]
	}</pre>
	*/
	Body *models.AccessPolicy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update access policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAccessPolicyParams) WithDefaults() *UpdateAccessPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update access policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAccessPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update access policy params
func (o *UpdateAccessPolicyParams) WithTimeout(timeout time.Duration) *UpdateAccessPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update access policy params
func (o *UpdateAccessPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update access policy params
func (o *UpdateAccessPolicyParams) WithContext(ctx context.Context) *UpdateAccessPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update access policy params
func (o *UpdateAccessPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update access policy params
func (o *UpdateAccessPolicyParams) WithHTTPClient(client *http.Client) *UpdateAccessPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update access policy params
func (o *UpdateAccessPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update access policy params
func (o *UpdateAccessPolicyParams) WithBody(body *models.AccessPolicy) *UpdateAccessPolicyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update access policy params
func (o *UpdateAccessPolicyParams) SetBody(body *models.AccessPolicy) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAccessPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
