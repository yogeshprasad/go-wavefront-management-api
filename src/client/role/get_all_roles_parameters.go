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
	"github.com/go-openapi/swag"
)

// NewGetAllRolesParams creates a new GetAllRolesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllRolesParams() *GetAllRolesParams {
	return &GetAllRolesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllRolesParamsWithTimeout creates a new GetAllRolesParams object
// with the ability to set a timeout on a request.
func NewGetAllRolesParamsWithTimeout(timeout time.Duration) *GetAllRolesParams {
	return &GetAllRolesParams{
		timeout: timeout,
	}
}

// NewGetAllRolesParamsWithContext creates a new GetAllRolesParams object
// with the ability to set a context for a request.
func NewGetAllRolesParamsWithContext(ctx context.Context) *GetAllRolesParams {
	return &GetAllRolesParams{
		Context: ctx,
	}
}

// NewGetAllRolesParamsWithHTTPClient creates a new GetAllRolesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllRolesParamsWithHTTPClient(client *http.Client) *GetAllRolesParams {
	return &GetAllRolesParams{
		HTTPClient: client,
	}
}

/*
GetAllRolesParams contains all the parameters to send to the API endpoint

	for the get all roles operation.

	Typically these are written to a http.Request.
*/
type GetAllRolesParams struct {

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

// WithDefaults hydrates default values in the get all roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllRolesParams) WithDefaults() *GetAllRolesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllRolesParams) SetDefaults() {
	var (
		limitDefault = int32(100)

		offsetDefault = int32(0)
	)

	val := GetAllRolesParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get all roles params
func (o *GetAllRolesParams) WithTimeout(timeout time.Duration) *GetAllRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all roles params
func (o *GetAllRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all roles params
func (o *GetAllRolesParams) WithContext(ctx context.Context) *GetAllRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all roles params
func (o *GetAllRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all roles params
func (o *GetAllRolesParams) WithHTTPClient(client *http.Client) *GetAllRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all roles params
func (o *GetAllRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get all roles params
func (o *GetAllRolesParams) WithLimit(limit *int32) *GetAllRolesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get all roles params
func (o *GetAllRolesParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get all roles params
func (o *GetAllRolesParams) WithOffset(offset *int32) *GetAllRolesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get all roles params
func (o *GetAllRolesParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
