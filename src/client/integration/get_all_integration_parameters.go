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
	"github.com/go-openapi/swag"
)

// NewGetAllIntegrationParams creates a new GetAllIntegrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllIntegrationParams() *GetAllIntegrationParams {
	return &GetAllIntegrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllIntegrationParamsWithTimeout creates a new GetAllIntegrationParams object
// with the ability to set a timeout on a request.
func NewGetAllIntegrationParamsWithTimeout(timeout time.Duration) *GetAllIntegrationParams {
	return &GetAllIntegrationParams{
		timeout: timeout,
	}
}

// NewGetAllIntegrationParamsWithContext creates a new GetAllIntegrationParams object
// with the ability to set a context for a request.
func NewGetAllIntegrationParamsWithContext(ctx context.Context) *GetAllIntegrationParams {
	return &GetAllIntegrationParams{
		Context: ctx,
	}
}

// NewGetAllIntegrationParamsWithHTTPClient creates a new GetAllIntegrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllIntegrationParamsWithHTTPClient(client *http.Client) *GetAllIntegrationParams {
	return &GetAllIntegrationParams{
		HTTPClient: client,
	}
}

/*
GetAllIntegrationParams contains all the parameters to send to the API endpoint

	for the get all integration operation.

	Typically these are written to a http.Request.
*/
type GetAllIntegrationParams struct {

	// Limit.
	//
	// Format: int32
	// Default: 100
	Limit *int32

	// Offset.
	//
	// Format: int32
	Offset *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllIntegrationParams) WithDefaults() *GetAllIntegrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllIntegrationParams) SetDefaults() {
	var (
		limitDefault = int32(100)

		offsetDefault = int32(0)
	)

	val := GetAllIntegrationParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get all integration params
func (o *GetAllIntegrationParams) WithTimeout(timeout time.Duration) *GetAllIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all integration params
func (o *GetAllIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all integration params
func (o *GetAllIntegrationParams) WithContext(ctx context.Context) *GetAllIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all integration params
func (o *GetAllIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all integration params
func (o *GetAllIntegrationParams) WithHTTPClient(client *http.Client) *GetAllIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all integration params
func (o *GetAllIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get all integration params
func (o *GetAllIntegrationParams) WithLimit(limit *int32) *GetAllIntegrationParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get all integration params
func (o *GetAllIntegrationParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get all integration params
func (o *GetAllIntegrationParams) WithOffset(offset *int32) *GetAllIntegrationParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get all integration params
func (o *GetAllIntegrationParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int32

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
