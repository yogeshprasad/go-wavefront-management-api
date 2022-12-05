// Code generated by go-swagger; DO NOT EDIT.

package saved_traces_search_group

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

// NewGetSavedTracesSearchGroupParams creates a new GetSavedTracesSearchGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSavedTracesSearchGroupParams() *GetSavedTracesSearchGroupParams {
	return &GetSavedTracesSearchGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSavedTracesSearchGroupParamsWithTimeout creates a new GetSavedTracesSearchGroupParams object
// with the ability to set a timeout on a request.
func NewGetSavedTracesSearchGroupParamsWithTimeout(timeout time.Duration) *GetSavedTracesSearchGroupParams {
	return &GetSavedTracesSearchGroupParams{
		timeout: timeout,
	}
}

// NewGetSavedTracesSearchGroupParamsWithContext creates a new GetSavedTracesSearchGroupParams object
// with the ability to set a context for a request.
func NewGetSavedTracesSearchGroupParamsWithContext(ctx context.Context) *GetSavedTracesSearchGroupParams {
	return &GetSavedTracesSearchGroupParams{
		Context: ctx,
	}
}

// NewGetSavedTracesSearchGroupParamsWithHTTPClient creates a new GetSavedTracesSearchGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSavedTracesSearchGroupParamsWithHTTPClient(client *http.Client) *GetSavedTracesSearchGroupParams {
	return &GetSavedTracesSearchGroupParams{
		HTTPClient: client,
	}
}

/*
GetSavedTracesSearchGroupParams contains all the parameters to send to the API endpoint

	for the get saved traces search group operation.

	Typically these are written to a http.Request.
*/
type GetSavedTracesSearchGroupParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get saved traces search group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSavedTracesSearchGroupParams) WithDefaults() *GetSavedTracesSearchGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get saved traces search group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSavedTracesSearchGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get saved traces search group params
func (o *GetSavedTracesSearchGroupParams) WithTimeout(timeout time.Duration) *GetSavedTracesSearchGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get saved traces search group params
func (o *GetSavedTracesSearchGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get saved traces search group params
func (o *GetSavedTracesSearchGroupParams) WithContext(ctx context.Context) *GetSavedTracesSearchGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get saved traces search group params
func (o *GetSavedTracesSearchGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get saved traces search group params
func (o *GetSavedTracesSearchGroupParams) WithHTTPClient(client *http.Client) *GetSavedTracesSearchGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get saved traces search group params
func (o *GetSavedTracesSearchGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get saved traces search group params
func (o *GetSavedTracesSearchGroupParams) WithID(id string) *GetSavedTracesSearchGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get saved traces search group params
func (o *GetSavedTracesSearchGroupParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSavedTracesSearchGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
