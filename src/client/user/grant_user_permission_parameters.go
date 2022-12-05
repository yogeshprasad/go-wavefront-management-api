// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewGrantUserPermissionParams creates a new GrantUserPermissionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGrantUserPermissionParams() *GrantUserPermissionParams {
	return &GrantUserPermissionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGrantUserPermissionParamsWithTimeout creates a new GrantUserPermissionParams object
// with the ability to set a timeout on a request.
func NewGrantUserPermissionParamsWithTimeout(timeout time.Duration) *GrantUserPermissionParams {
	return &GrantUserPermissionParams{
		timeout: timeout,
	}
}

// NewGrantUserPermissionParamsWithContext creates a new GrantUserPermissionParams object
// with the ability to set a context for a request.
func NewGrantUserPermissionParamsWithContext(ctx context.Context) *GrantUserPermissionParams {
	return &GrantUserPermissionParams{
		Context: ctx,
	}
}

// NewGrantUserPermissionParamsWithHTTPClient creates a new GrantUserPermissionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGrantUserPermissionParamsWithHTTPClient(client *http.Client) *GrantUserPermissionParams {
	return &GrantUserPermissionParams{
		HTTPClient: client,
	}
}

/*
GrantUserPermissionParams contains all the parameters to send to the API endpoint

	for the grant user permission operation.

	Typically these are written to a http.Request.
*/
type GrantUserPermissionParams struct {

	/* Group.

	   Permission group to grant to the account. Please note that 'host_tag_management' is the equivalent of the 'Source Tag Management' permission
	*/
	Group *string

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the grant user permission params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GrantUserPermissionParams) WithDefaults() *GrantUserPermissionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the grant user permission params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GrantUserPermissionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the grant user permission params
func (o *GrantUserPermissionParams) WithTimeout(timeout time.Duration) *GrantUserPermissionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the grant user permission params
func (o *GrantUserPermissionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the grant user permission params
func (o *GrantUserPermissionParams) WithContext(ctx context.Context) *GrantUserPermissionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the grant user permission params
func (o *GrantUserPermissionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the grant user permission params
func (o *GrantUserPermissionParams) WithHTTPClient(client *http.Client) *GrantUserPermissionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the grant user permission params
func (o *GrantUserPermissionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroup adds the group to the grant user permission params
func (o *GrantUserPermissionParams) WithGroup(group *string) *GrantUserPermissionParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the grant user permission params
func (o *GrantUserPermissionParams) SetGroup(group *string) {
	o.Group = group
}

// WithID adds the id to the grant user permission params
func (o *GrantUserPermissionParams) WithID(id string) *GrantUserPermissionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the grant user permission params
func (o *GrantUserPermissionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GrantUserPermissionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Group != nil {

		// form param group
		var frGroup string
		if o.Group != nil {
			frGroup = *o.Group
		}
		fGroup := frGroup
		if fGroup != "" {
			if err := r.SetFormParam("group", fGroup); err != nil {
				return err
			}
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
