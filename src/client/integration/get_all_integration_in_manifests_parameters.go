// Code generated by go-swagger; DO NOT EDIT.

package integration

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

// NewGetAllIntegrationInManifestsParams creates a new GetAllIntegrationInManifestsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllIntegrationInManifestsParams() *GetAllIntegrationInManifestsParams {
	return &GetAllIntegrationInManifestsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllIntegrationInManifestsParamsWithTimeout creates a new GetAllIntegrationInManifestsParams object
// with the ability to set a timeout on a request.
func NewGetAllIntegrationInManifestsParamsWithTimeout(timeout time.Duration) *GetAllIntegrationInManifestsParams {
	return &GetAllIntegrationInManifestsParams{
		timeout: timeout,
	}
}

// NewGetAllIntegrationInManifestsParamsWithContext creates a new GetAllIntegrationInManifestsParams object
// with the ability to set a context for a request.
func NewGetAllIntegrationInManifestsParamsWithContext(ctx context.Context) *GetAllIntegrationInManifestsParams {
	return &GetAllIntegrationInManifestsParams{
		Context: ctx,
	}
}

// NewGetAllIntegrationInManifestsParamsWithHTTPClient creates a new GetAllIntegrationInManifestsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllIntegrationInManifestsParamsWithHTTPClient(client *http.Client) *GetAllIntegrationInManifestsParams {
	return &GetAllIntegrationInManifestsParams{
		HTTPClient: client,
	}
}

/*
GetAllIntegrationInManifestsParams contains all the parameters to send to the API endpoint

	for the get all integration in manifests operation.

	Typically these are written to a http.Request.
*/
type GetAllIntegrationInManifestsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all integration in manifests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllIntegrationInManifestsParams) WithDefaults() *GetAllIntegrationInManifestsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all integration in manifests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllIntegrationInManifestsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get all integration in manifests params
func (o *GetAllIntegrationInManifestsParams) WithTimeout(timeout time.Duration) *GetAllIntegrationInManifestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all integration in manifests params
func (o *GetAllIntegrationInManifestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all integration in manifests params
func (o *GetAllIntegrationInManifestsParams) WithContext(ctx context.Context) *GetAllIntegrationInManifestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all integration in manifests params
func (o *GetAllIntegrationInManifestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all integration in manifests params
func (o *GetAllIntegrationInManifestsParams) WithHTTPClient(client *http.Client) *GetAllIntegrationInManifestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all integration in manifests params
func (o *GetAllIntegrationInManifestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllIntegrationInManifestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
