// Code generated by go-swagger; DO NOT EDIT.

package alert

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
	"github.com/go-openapi/swag"
)

// NewUnsnoozeAlertParams creates a new UnsnoozeAlertParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnsnoozeAlertParams() *UnsnoozeAlertParams {
	return &UnsnoozeAlertParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnsnoozeAlertParamsWithTimeout creates a new UnsnoozeAlertParams object
// with the ability to set a timeout on a request.
func NewUnsnoozeAlertParamsWithTimeout(timeout time.Duration) *UnsnoozeAlertParams {
	return &UnsnoozeAlertParams{
		timeout: timeout,
	}
}

// NewUnsnoozeAlertParamsWithContext creates a new UnsnoozeAlertParams object
// with the ability to set a context for a request.
func NewUnsnoozeAlertParamsWithContext(ctx context.Context) *UnsnoozeAlertParams {
	return &UnsnoozeAlertParams{
		Context: ctx,
	}
}

// NewUnsnoozeAlertParamsWithHTTPClient creates a new UnsnoozeAlertParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnsnoozeAlertParamsWithHTTPClient(client *http.Client) *UnsnoozeAlertParams {
	return &UnsnoozeAlertParams{
		HTTPClient: client,
	}
}

/*
UnsnoozeAlertParams contains all the parameters to send to the API endpoint

	for the unsnooze alert operation.

	Typically these are written to a http.Request.
*/
type UnsnoozeAlertParams struct {

	// ID.
	//
	// Format: int64
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the unsnooze alert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnsnoozeAlertParams) WithDefaults() *UnsnoozeAlertParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unsnooze alert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnsnoozeAlertParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the unsnooze alert params
func (o *UnsnoozeAlertParams) WithTimeout(timeout time.Duration) *UnsnoozeAlertParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unsnooze alert params
func (o *UnsnoozeAlertParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unsnooze alert params
func (o *UnsnoozeAlertParams) WithContext(ctx context.Context) *UnsnoozeAlertParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unsnooze alert params
func (o *UnsnoozeAlertParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unsnooze alert params
func (o *UnsnoozeAlertParams) WithHTTPClient(client *http.Client) *UnsnoozeAlertParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unsnooze alert params
func (o *UnsnoozeAlertParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the unsnooze alert params
func (o *UnsnoozeAlertParams) WithID(id int64) *UnsnoozeAlertParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the unsnooze alert params
func (o *UnsnoozeAlertParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UnsnoozeAlertParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
