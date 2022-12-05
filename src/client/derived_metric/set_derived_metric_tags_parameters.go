// Code generated by go-swagger; DO NOT EDIT.

package derived_metric

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

// NewSetDerivedMetricTagsParams creates a new SetDerivedMetricTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetDerivedMetricTagsParams() *SetDerivedMetricTagsParams {
	return &SetDerivedMetricTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetDerivedMetricTagsParamsWithTimeout creates a new SetDerivedMetricTagsParams object
// with the ability to set a timeout on a request.
func NewSetDerivedMetricTagsParamsWithTimeout(timeout time.Duration) *SetDerivedMetricTagsParams {
	return &SetDerivedMetricTagsParams{
		timeout: timeout,
	}
}

// NewSetDerivedMetricTagsParamsWithContext creates a new SetDerivedMetricTagsParams object
// with the ability to set a context for a request.
func NewSetDerivedMetricTagsParamsWithContext(ctx context.Context) *SetDerivedMetricTagsParams {
	return &SetDerivedMetricTagsParams{
		Context: ctx,
	}
}

// NewSetDerivedMetricTagsParamsWithHTTPClient creates a new SetDerivedMetricTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetDerivedMetricTagsParamsWithHTTPClient(client *http.Client) *SetDerivedMetricTagsParams {
	return &SetDerivedMetricTagsParams{
		HTTPClient: client,
	}
}

/*
SetDerivedMetricTagsParams contains all the parameters to send to the API endpoint

	for the set derived metric tags operation.

	Typically these are written to a http.Request.
*/
type SetDerivedMetricTagsParams struct {

	// Body.
	Body []string

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set derived metric tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetDerivedMetricTagsParams) WithDefaults() *SetDerivedMetricTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set derived metric tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetDerivedMetricTagsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set derived metric tags params
func (o *SetDerivedMetricTagsParams) WithTimeout(timeout time.Duration) *SetDerivedMetricTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set derived metric tags params
func (o *SetDerivedMetricTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set derived metric tags params
func (o *SetDerivedMetricTagsParams) WithContext(ctx context.Context) *SetDerivedMetricTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set derived metric tags params
func (o *SetDerivedMetricTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set derived metric tags params
func (o *SetDerivedMetricTagsParams) WithHTTPClient(client *http.Client) *SetDerivedMetricTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set derived metric tags params
func (o *SetDerivedMetricTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set derived metric tags params
func (o *SetDerivedMetricTagsParams) WithBody(body []string) *SetDerivedMetricTagsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set derived metric tags params
func (o *SetDerivedMetricTagsParams) SetBody(body []string) {
	o.Body = body
}

// WithID adds the id to the set derived metric tags params
func (o *SetDerivedMetricTagsParams) WithID(id string) *SetDerivedMetricTagsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the set derived metric tags params
func (o *SetDerivedMetricTagsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SetDerivedMetricTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
