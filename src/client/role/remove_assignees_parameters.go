// Code generated by go-swagger; DO NOT EDIT.

package role

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

// NewRemoveAssigneesParams creates a new RemoveAssigneesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveAssigneesParams() *RemoveAssigneesParams {
	return &RemoveAssigneesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveAssigneesParamsWithTimeout creates a new RemoveAssigneesParams object
// with the ability to set a timeout on a request.
func NewRemoveAssigneesParamsWithTimeout(timeout time.Duration) *RemoveAssigneesParams {
	return &RemoveAssigneesParams{
		timeout: timeout,
	}
}

// NewRemoveAssigneesParamsWithContext creates a new RemoveAssigneesParams object
// with the ability to set a context for a request.
func NewRemoveAssigneesParamsWithContext(ctx context.Context) *RemoveAssigneesParams {
	return &RemoveAssigneesParams{
		Context: ctx,
	}
}

// NewRemoveAssigneesParamsWithHTTPClient creates a new RemoveAssigneesParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveAssigneesParamsWithHTTPClient(client *http.Client) *RemoveAssigneesParams {
	return &RemoveAssigneesParams{
		HTTPClient: client,
	}
}

/*
RemoveAssigneesParams contains all the parameters to send to the API endpoint

	for the remove assignees operation.

	Typically these are written to a http.Request.
*/
type RemoveAssigneesParams struct {

	/* Body.

	   List of users and user groups thatshould be removed from role
	*/
	Body []string

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove assignees params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveAssigneesParams) WithDefaults() *RemoveAssigneesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove assignees params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveAssigneesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove assignees params
func (o *RemoveAssigneesParams) WithTimeout(timeout time.Duration) *RemoveAssigneesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove assignees params
func (o *RemoveAssigneesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove assignees params
func (o *RemoveAssigneesParams) WithContext(ctx context.Context) *RemoveAssigneesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove assignees params
func (o *RemoveAssigneesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove assignees params
func (o *RemoveAssigneesParams) WithHTTPClient(client *http.Client) *RemoveAssigneesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove assignees params
func (o *RemoveAssigneesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the remove assignees params
func (o *RemoveAssigneesParams) WithBody(body []string) *RemoveAssigneesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the remove assignees params
func (o *RemoveAssigneesParams) SetBody(body []string) {
	o.Body = body
}

// WithID adds the id to the remove assignees params
func (o *RemoveAssigneesParams) WithID(id string) *RemoveAssigneesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the remove assignees params
func (o *RemoveAssigneesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveAssigneesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
