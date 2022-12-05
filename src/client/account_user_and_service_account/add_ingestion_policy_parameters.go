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

	"models"
)

// NewAddIngestionPolicyParams creates a new AddIngestionPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddIngestionPolicyParams() *AddIngestionPolicyParams {
	return &AddIngestionPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddIngestionPolicyParamsWithTimeout creates a new AddIngestionPolicyParams object
// with the ability to set a timeout on a request.
func NewAddIngestionPolicyParamsWithTimeout(timeout time.Duration) *AddIngestionPolicyParams {
	return &AddIngestionPolicyParams{
		timeout: timeout,
	}
}

// NewAddIngestionPolicyParamsWithContext creates a new AddIngestionPolicyParams object
// with the ability to set a context for a request.
func NewAddIngestionPolicyParamsWithContext(ctx context.Context) *AddIngestionPolicyParams {
	return &AddIngestionPolicyParams{
		Context: ctx,
	}
}

// NewAddIngestionPolicyParamsWithHTTPClient creates a new AddIngestionPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddIngestionPolicyParamsWithHTTPClient(client *http.Client) *AddIngestionPolicyParams {
	return &AddIngestionPolicyParams{
		HTTPClient: client,
	}
}

/*
AddIngestionPolicyParams contains all the parameters to send to the API endpoint

	for the add ingestion policy operation.

	Typically these are written to a http.Request.
*/
type AddIngestionPolicyParams struct {

	/* Body.

	     Example Body:
	<pre>{
	  "ingestionPolicyId": "Ingestion policy identifier",
	  "accounts": [
	  "account1",
	  "account2",
	  "account3"
	  ],
	  "groups": [
	  "group1",
	  "group2"
	  ]
	}</pre>
	*/
	Body *models.IngestionPolicyMapping

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add ingestion policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddIngestionPolicyParams) WithDefaults() *AddIngestionPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add ingestion policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddIngestionPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add ingestion policy params
func (o *AddIngestionPolicyParams) WithTimeout(timeout time.Duration) *AddIngestionPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add ingestion policy params
func (o *AddIngestionPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add ingestion policy params
func (o *AddIngestionPolicyParams) WithContext(ctx context.Context) *AddIngestionPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add ingestion policy params
func (o *AddIngestionPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add ingestion policy params
func (o *AddIngestionPolicyParams) WithHTTPClient(client *http.Client) *AddIngestionPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add ingestion policy params
func (o *AddIngestionPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add ingestion policy params
func (o *AddIngestionPolicyParams) WithBody(body *models.IngestionPolicyMapping) *AddIngestionPolicyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add ingestion policy params
func (o *AddIngestionPolicyParams) SetBody(body *models.IngestionPolicyMapping) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddIngestionPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
