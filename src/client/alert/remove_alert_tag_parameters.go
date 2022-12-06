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
)

// NewRemoveAlertTagParams creates a new RemoveAlertTagParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveAlertTagParams() *RemoveAlertTagParams {
	return &RemoveAlertTagParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveAlertTagParamsWithTimeout creates a new RemoveAlertTagParams object
// with the ability to set a timeout on a request.
func NewRemoveAlertTagParamsWithTimeout(timeout time.Duration) *RemoveAlertTagParams {
	return &RemoveAlertTagParams{
		timeout: timeout,
	}
}

// NewRemoveAlertTagParamsWithContext creates a new RemoveAlertTagParams object
// with the ability to set a context for a request.
func NewRemoveAlertTagParamsWithContext(ctx context.Context) *RemoveAlertTagParams {
	return &RemoveAlertTagParams{
		Context: ctx,
	}
}

// NewRemoveAlertTagParamsWithHTTPClient creates a new RemoveAlertTagParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveAlertTagParamsWithHTTPClient(client *http.Client) *RemoveAlertTagParams {
	return &RemoveAlertTagParams{
		HTTPClient: client,
	}
}

/*
RemoveAlertTagParams contains all the parameters to send to the API endpoint

	for the remove alert tag operation.

	Typically these are written to a http.Request.
*/
type RemoveAlertTagParams struct {

	// ID.
	ID string

	/* TagValue.

	     Supported Characters of Tags:
	<pre>Tag names can contain alphanumeric (a-z, A-Z, 0-9),
	dash (-), underscore (_), and colon (:) characters.
	The space character is not supported.</pre>

	*/
	TagValue string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove alert tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveAlertTagParams) WithDefaults() *RemoveAlertTagParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove alert tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveAlertTagParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove alert tag params
func (o *RemoveAlertTagParams) WithTimeout(timeout time.Duration) *RemoveAlertTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove alert tag params
func (o *RemoveAlertTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove alert tag params
func (o *RemoveAlertTagParams) WithContext(ctx context.Context) *RemoveAlertTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove alert tag params
func (o *RemoveAlertTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove alert tag params
func (o *RemoveAlertTagParams) WithHTTPClient(client *http.Client) *RemoveAlertTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove alert tag params
func (o *RemoveAlertTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the remove alert tag params
func (o *RemoveAlertTagParams) WithID(id string) *RemoveAlertTagParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the remove alert tag params
func (o *RemoveAlertTagParams) SetID(id string) {
	o.ID = id
}

// WithTagValue adds the tagValue to the remove alert tag params
func (o *RemoveAlertTagParams) WithTagValue(tagValue string) *RemoveAlertTagParams {
	o.SetTagValue(tagValue)
	return o
}

// SetTagValue adds the tagValue to the remove alert tag params
func (o *RemoveAlertTagParams) SetTagValue(tagValue string) {
	o.TagValue = tagValue
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveAlertTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
