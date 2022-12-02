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

	"models"
)

// NewInviteUsersParams creates a new InviteUsersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInviteUsersParams() *InviteUsersParams {
	return &InviteUsersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInviteUsersParamsWithTimeout creates a new InviteUsersParams object
// with the ability to set a timeout on a request.
func NewInviteUsersParamsWithTimeout(timeout time.Duration) *InviteUsersParams {
	return &InviteUsersParams{
		timeout: timeout,
	}
}

// NewInviteUsersParamsWithContext creates a new InviteUsersParams object
// with the ability to set a context for a request.
func NewInviteUsersParamsWithContext(ctx context.Context) *InviteUsersParams {
	return &InviteUsersParams{
		Context: ctx,
	}
}

// NewInviteUsersParamsWithHTTPClient creates a new InviteUsersParams object
// with the ability to set a custom HTTPClient for a request.
func NewInviteUsersParamsWithHTTPClient(client *http.Client) *InviteUsersParams {
	return &InviteUsersParams{
		HTTPClient: client,
	}
}

/*
InviteUsersParams contains all the parameters to send to the API endpoint

	for the invite users operation.

	Typically these are written to a http.Request.
*/
type InviteUsersParams struct {

	/* Body.

	     Example Body:
	<pre>[
	{
	  "emailAddress": "user@example.com",
	  "groups": [
	    "user_management"
	  ],
	  "userGroups": [
	    "8b23136b-ecd2-4cb5-8c92-62477dcc4090"
	  ],
	  "roles": [
	    "Role"
	  ],
	  "ingestionPolicies": [
	    "policyId1",
	    "policyId2"
	  ]
	}
	]</pre>
	*/
	Body []*models.UserToCreate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the invite users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InviteUsersParams) WithDefaults() *InviteUsersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the invite users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InviteUsersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the invite users params
func (o *InviteUsersParams) WithTimeout(timeout time.Duration) *InviteUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the invite users params
func (o *InviteUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the invite users params
func (o *InviteUsersParams) WithContext(ctx context.Context) *InviteUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the invite users params
func (o *InviteUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the invite users params
func (o *InviteUsersParams) WithHTTPClient(client *http.Client) *InviteUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the invite users params
func (o *InviteUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the invite users params
func (o *InviteUsersParams) WithBody(body []*models.UserToCreate) *InviteUsersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the invite users params
func (o *InviteUsersParams) SetBody(body []*models.UserToCreate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *InviteUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
