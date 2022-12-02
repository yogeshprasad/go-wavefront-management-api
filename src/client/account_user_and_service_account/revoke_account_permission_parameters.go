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

// NewRevokeAccountPermissionParams creates a new RevokeAccountPermissionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRevokeAccountPermissionParams() *RevokeAccountPermissionParams {
	return &RevokeAccountPermissionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRevokeAccountPermissionParamsWithTimeout creates a new RevokeAccountPermissionParams object
// with the ability to set a timeout on a request.
func NewRevokeAccountPermissionParamsWithTimeout(timeout time.Duration) *RevokeAccountPermissionParams {
	return &RevokeAccountPermissionParams{
		timeout: timeout,
	}
}

// NewRevokeAccountPermissionParamsWithContext creates a new RevokeAccountPermissionParams object
// with the ability to set a context for a request.
func NewRevokeAccountPermissionParamsWithContext(ctx context.Context) *RevokeAccountPermissionParams {
	return &RevokeAccountPermissionParams{
		Context: ctx,
	}
}

// NewRevokeAccountPermissionParamsWithHTTPClient creates a new RevokeAccountPermissionParams object
// with the ability to set a custom HTTPClient for a request.
func NewRevokeAccountPermissionParamsWithHTTPClient(client *http.Client) *RevokeAccountPermissionParams {
	return &RevokeAccountPermissionParams{
		HTTPClient: client,
	}
}

/*
RevokeAccountPermissionParams contains all the parameters to send to the API endpoint

	for the revoke account permission operation.

	Typically these are written to a http.Request.
*/
type RevokeAccountPermissionParams struct {

	// ID.
	ID string

	// Permission.
	Permission string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the revoke account permission params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevokeAccountPermissionParams) WithDefaults() *RevokeAccountPermissionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the revoke account permission params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevokeAccountPermissionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the revoke account permission params
func (o *RevokeAccountPermissionParams) WithTimeout(timeout time.Duration) *RevokeAccountPermissionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the revoke account permission params
func (o *RevokeAccountPermissionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the revoke account permission params
func (o *RevokeAccountPermissionParams) WithContext(ctx context.Context) *RevokeAccountPermissionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the revoke account permission params
func (o *RevokeAccountPermissionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the revoke account permission params
func (o *RevokeAccountPermissionParams) WithHTTPClient(client *http.Client) *RevokeAccountPermissionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the revoke account permission params
func (o *RevokeAccountPermissionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the revoke account permission params
func (o *RevokeAccountPermissionParams) WithID(id string) *RevokeAccountPermissionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the revoke account permission params
func (o *RevokeAccountPermissionParams) SetID(id string) {
	o.ID = id
}

// WithPermission adds the permission to the revoke account permission params
func (o *RevokeAccountPermissionParams) WithPermission(permission string) *RevokeAccountPermissionParams {
	o.SetPermission(permission)
	return o
}

// SetPermission adds the permission to the revoke account permission params
func (o *RevokeAccountPermissionParams) SetPermission(permission string) {
	o.Permission = permission
}

// WriteToRequest writes these params to a swagger request
func (o *RevokeAccountPermissionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param permission
	if err := r.SetPathParam("permission", o.Permission); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
