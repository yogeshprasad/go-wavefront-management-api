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

// NewRemoveEventTagParams creates a new RemoveEventTagParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveEventTagParams() *RemoveEventTagParams {
	return &RemoveEventTagParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveEventTagParamsWithTimeout creates a new RemoveEventTagParams object
// with the ability to set a timeout on a request.
func NewRemoveEventTagParamsWithTimeout(timeout time.Duration) *RemoveEventTagParams {
	return &RemoveEventTagParams{
		timeout: timeout,
	}
}

// NewRemoveEventTagParamsWithContext creates a new RemoveEventTagParams object
// with the ability to set a context for a request.
func NewRemoveEventTagParamsWithContext(ctx context.Context) *RemoveEventTagParams {
	return &RemoveEventTagParams{
		Context: ctx,
	}
}

// NewRemoveEventTagParamsWithHTTPClient creates a new RemoveEventTagParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveEventTagParamsWithHTTPClient(client *http.Client) *RemoveEventTagParams {
	return &RemoveEventTagParams{
		HTTPClient: client,
	}
}

/*
RemoveEventTagParams contains all the parameters to send to the API endpoint

	for the remove event tag operation.

	Typically these are written to a http.Request.
*/
type RemoveEventTagParams struct {

	// ID.
	ID string

	// TagValue.
	TagValue string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove event tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveEventTagParams) WithDefaults() *RemoveEventTagParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove event tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveEventTagParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove event tag params
func (o *RemoveEventTagParams) WithTimeout(timeout time.Duration) *RemoveEventTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove event tag params
func (o *RemoveEventTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove event tag params
func (o *RemoveEventTagParams) WithContext(ctx context.Context) *RemoveEventTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove event tag params
func (o *RemoveEventTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove event tag params
func (o *RemoveEventTagParams) WithHTTPClient(client *http.Client) *RemoveEventTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove event tag params
func (o *RemoveEventTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the remove event tag params
func (o *RemoveEventTagParams) WithID(id string) *RemoveEventTagParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the remove event tag params
func (o *RemoveEventTagParams) SetID(id string) {
	o.ID = id
}

// WithTagValue adds the tagValue to the remove event tag params
func (o *RemoveEventTagParams) WithTagValue(tagValue string) *RemoveEventTagParams {
	o.SetTagValue(tagValue)
	return o
}

// SetTagValue adds the tagValue to the remove event tag params
func (o *RemoveEventTagParams) SetTagValue(tagValue string) {
	o.TagValue = tagValue
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveEventTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
