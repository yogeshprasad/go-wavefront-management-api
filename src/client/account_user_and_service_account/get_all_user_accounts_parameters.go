// Code generated by go-swagger; DO NOT EDIT.

package account_user_and_service_account

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

// NewGetAllUserAccountsParams creates a new GetAllUserAccountsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllUserAccountsParams() *GetAllUserAccountsParams {
	return &GetAllUserAccountsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllUserAccountsParamsWithTimeout creates a new GetAllUserAccountsParams object
// with the ability to set a timeout on a request.
func NewGetAllUserAccountsParamsWithTimeout(timeout time.Duration) *GetAllUserAccountsParams {
	return &GetAllUserAccountsParams{
		timeout: timeout,
	}
}

// NewGetAllUserAccountsParamsWithContext creates a new GetAllUserAccountsParams object
// with the ability to set a context for a request.
func NewGetAllUserAccountsParamsWithContext(ctx context.Context) *GetAllUserAccountsParams {
	return &GetAllUserAccountsParams{
		Context: ctx,
	}
}

// NewGetAllUserAccountsParamsWithHTTPClient creates a new GetAllUserAccountsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllUserAccountsParamsWithHTTPClient(client *http.Client) *GetAllUserAccountsParams {
	return &GetAllUserAccountsParams{
		HTTPClient: client,
	}
}

/*
GetAllUserAccountsParams contains all the parameters to send to the API endpoint

	for the get all user accounts operation.

	Typically these are written to a http.Request.
*/
type GetAllUserAccountsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all user accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllUserAccountsParams) WithDefaults() *GetAllUserAccountsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all user accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllUserAccountsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all user accounts params
func (o *GetAllUserAccountsParams) WithTimeout(timeout time.Duration) *GetAllUserAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all user accounts params
func (o *GetAllUserAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all user accounts params
func (o *GetAllUserAccountsParams) WithContext(ctx context.Context) *GetAllUserAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all user accounts params
func (o *GetAllUserAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all user accounts params
func (o *GetAllUserAccountsParams) WithHTTPClient(client *http.Client) *GetAllUserAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all user accounts params
func (o *GetAllUserAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllUserAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
