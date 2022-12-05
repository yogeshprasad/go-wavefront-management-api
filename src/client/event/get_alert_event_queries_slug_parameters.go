// Code generated by go-swagger; DO NOT EDIT.

package event

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

// NewGetAlertEventQueriesSlugParams creates a new GetAlertEventQueriesSlugParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAlertEventQueriesSlugParams() *GetAlertEventQueriesSlugParams {
	return &GetAlertEventQueriesSlugParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAlertEventQueriesSlugParamsWithTimeout creates a new GetAlertEventQueriesSlugParams object
// with the ability to set a timeout on a request.
func NewGetAlertEventQueriesSlugParamsWithTimeout(timeout time.Duration) *GetAlertEventQueriesSlugParams {
	return &GetAlertEventQueriesSlugParams{
		timeout: timeout,
	}
}

// NewGetAlertEventQueriesSlugParamsWithContext creates a new GetAlertEventQueriesSlugParams object
// with the ability to set a context for a request.
func NewGetAlertEventQueriesSlugParamsWithContext(ctx context.Context) *GetAlertEventQueriesSlugParams {
	return &GetAlertEventQueriesSlugParams{
		Context: ctx,
	}
}

// NewGetAlertEventQueriesSlugParamsWithHTTPClient creates a new GetAlertEventQueriesSlugParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAlertEventQueriesSlugParamsWithHTTPClient(client *http.Client) *GetAlertEventQueriesSlugParams {
	return &GetAlertEventQueriesSlugParams{
		HTTPClient: client,
	}
}

/*
GetAlertEventQueriesSlugParams contains all the parameters to send to the API endpoint

	for the get alert event queries slug operation.

	Typically these are written to a http.Request.
*/
type GetAlertEventQueriesSlugParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get alert event queries slug params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertEventQueriesSlugParams) WithDefaults() *GetAlertEventQueriesSlugParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get alert event queries slug params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertEventQueriesSlugParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get alert event queries slug params
func (o *GetAlertEventQueriesSlugParams) WithTimeout(timeout time.Duration) *GetAlertEventQueriesSlugParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get alert event queries slug params
func (o *GetAlertEventQueriesSlugParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get alert event queries slug params
func (o *GetAlertEventQueriesSlugParams) WithContext(ctx context.Context) *GetAlertEventQueriesSlugParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get alert event queries slug params
func (o *GetAlertEventQueriesSlugParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get alert event queries slug params
func (o *GetAlertEventQueriesSlugParams) WithHTTPClient(client *http.Client) *GetAlertEventQueriesSlugParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get alert event queries slug params
func (o *GetAlertEventQueriesSlugParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get alert event queries slug params
func (o *GetAlertEventQueriesSlugParams) WithID(id string) *GetAlertEventQueriesSlugParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get alert event queries slug params
func (o *GetAlertEventQueriesSlugParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAlertEventQueriesSlugParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
