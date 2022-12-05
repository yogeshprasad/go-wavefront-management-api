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
)

// NewAddDashboardTagParams creates a new AddDashboardTagParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddDashboardTagParams() *AddDashboardTagParams {
	return &AddDashboardTagParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddDashboardTagParamsWithTimeout creates a new AddDashboardTagParams object
// with the ability to set a timeout on a request.
func NewAddDashboardTagParamsWithTimeout(timeout time.Duration) *AddDashboardTagParams {
	return &AddDashboardTagParams{
		timeout: timeout,
	}
}

// NewAddDashboardTagParamsWithContext creates a new AddDashboardTagParams object
// with the ability to set a context for a request.
func NewAddDashboardTagParamsWithContext(ctx context.Context) *AddDashboardTagParams {
	return &AddDashboardTagParams{
		Context: ctx,
	}
}

// NewAddDashboardTagParamsWithHTTPClient creates a new AddDashboardTagParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddDashboardTagParamsWithHTTPClient(client *http.Client) *AddDashboardTagParams {
	return &AddDashboardTagParams{
		HTTPClient: client,
	}
}

/*
AddDashboardTagParams contains all the parameters to send to the API endpoint

	for the add dashboard tag operation.

	Typically these are written to a http.Request.
*/
type AddDashboardTagParams struct {

	// ID.
	ID string

	// TagValue.
	TagValue string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add dashboard tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddDashboardTagParams) WithDefaults() *AddDashboardTagParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add dashboard tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddDashboardTagParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add dashboard tag params
func (o *AddDashboardTagParams) WithTimeout(timeout time.Duration) *AddDashboardTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add dashboard tag params
func (o *AddDashboardTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add dashboard tag params
func (o *AddDashboardTagParams) WithContext(ctx context.Context) *AddDashboardTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add dashboard tag params
func (o *AddDashboardTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add dashboard tag params
func (o *AddDashboardTagParams) WithHTTPClient(client *http.Client) *AddDashboardTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add dashboard tag params
func (o *AddDashboardTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the add dashboard tag params
func (o *AddDashboardTagParams) WithID(id string) *AddDashboardTagParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the add dashboard tag params
func (o *AddDashboardTagParams) SetID(id string) {
	o.ID = id
}

// WithTagValue adds the tagValue to the add dashboard tag params
func (o *AddDashboardTagParams) WithTagValue(tagValue string) *AddDashboardTagParams {
	o.SetTagValue(tagValue)
	return o
}

// SetTagValue adds the tagValue to the add dashboard tag params
func (o *AddDashboardTagParams) SetTagValue(tagValue string) {
	o.TagValue = tagValue
}

// WriteToRequest writes these params to a swagger request
func (o *AddDashboardTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param tagValue
	if err := r.SetPathParam("tagValue", o.TagValue); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
