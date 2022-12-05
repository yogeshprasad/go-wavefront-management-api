// Code generated by go-swagger; DO NOT EDIT.

package recent_traces_search

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

// NewGetAllRecentTracesSearchesParams creates a new GetAllRecentTracesSearchesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllRecentTracesSearchesParams() *GetAllRecentTracesSearchesParams {
	return &GetAllRecentTracesSearchesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllRecentTracesSearchesParamsWithTimeout creates a new GetAllRecentTracesSearchesParams object
// with the ability to set a timeout on a request.
func NewGetAllRecentTracesSearchesParamsWithTimeout(timeout time.Duration) *GetAllRecentTracesSearchesParams {
	return &GetAllRecentTracesSearchesParams{
		timeout: timeout,
	}
}

// NewGetAllRecentTracesSearchesParamsWithContext creates a new GetAllRecentTracesSearchesParams object
// with the ability to set a context for a request.
func NewGetAllRecentTracesSearchesParamsWithContext(ctx context.Context) *GetAllRecentTracesSearchesParams {
	return &GetAllRecentTracesSearchesParams{
		Context: ctx,
	}
}

// NewGetAllRecentTracesSearchesParamsWithHTTPClient creates a new GetAllRecentTracesSearchesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllRecentTracesSearchesParamsWithHTTPClient(client *http.Client) *GetAllRecentTracesSearchesParams {
	return &GetAllRecentTracesSearchesParams{
		HTTPClient: client,
	}
}

/*
GetAllRecentTracesSearchesParams contains all the parameters to send to the API endpoint

	for the get all recent traces searches operation.

	Typically these are written to a http.Request.
*/
type GetAllRecentTracesSearchesParams struct {

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

// WithDefaults hydrates default values in the get all recent traces searches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllRecentTracesSearchesParams) WithDefaults() *GetAllRecentTracesSearchesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all recent traces searches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllRecentTracesSearchesParams) SetDefaults() {
	var (
		limitDefault = int32(100)

		offsetDefault = int32(0)
	)

	val := GetAllRecentTracesSearchesParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get all recent traces searches params
func (o *GetAllRecentTracesSearchesParams) WithTimeout(timeout time.Duration) *GetAllRecentTracesSearchesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all recent traces searches params
func (o *GetAllRecentTracesSearchesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all recent traces searches params
func (o *GetAllRecentTracesSearchesParams) WithContext(ctx context.Context) *GetAllRecentTracesSearchesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all recent traces searches params
func (o *GetAllRecentTracesSearchesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all recent traces searches params
func (o *GetAllRecentTracesSearchesParams) WithHTTPClient(client *http.Client) *GetAllRecentTracesSearchesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all recent traces searches params
func (o *GetAllRecentTracesSearchesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get all recent traces searches params
func (o *GetAllRecentTracesSearchesParams) WithLimit(limit *int32) *GetAllRecentTracesSearchesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get all recent traces searches params
func (o *GetAllRecentTracesSearchesParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get all recent traces searches params
func (o *GetAllRecentTracesSearchesParams) WithOffset(offset *int32) *GetAllRecentTracesSearchesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get all recent traces searches params
func (o *GetAllRecentTracesSearchesParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllRecentTracesSearchesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
