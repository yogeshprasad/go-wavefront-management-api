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

// NewInviteUserAccountsParams creates a new InviteUserAccountsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInviteUserAccountsParams() *InviteUserAccountsParams {
	return &InviteUserAccountsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInviteUserAccountsParamsWithTimeout creates a new InviteUserAccountsParams object
// with the ability to set a timeout on a request.
func NewInviteUserAccountsParamsWithTimeout(timeout time.Duration) *InviteUserAccountsParams {
	return &InviteUserAccountsParams{
		timeout: timeout,
	}
}

// NewInviteUserAccountsParamsWithContext creates a new InviteUserAccountsParams object
// with the ability to set a context for a request.
func NewInviteUserAccountsParamsWithContext(ctx context.Context) *InviteUserAccountsParams {
	return &InviteUserAccountsParams{
		Context: ctx,
	}
}

// NewInviteUserAccountsParamsWithHTTPClient creates a new InviteUserAccountsParams object
// with the ability to set a custom HTTPClient for a request.
func NewInviteUserAccountsParamsWithHTTPClient(client *http.Client) *InviteUserAccountsParams {
	return &InviteUserAccountsParams{
		HTTPClient: client,
	}
}

/*
InviteUserAccountsParams contains all the parameters to send to the API endpoint

	for the invite user accounts operation.

	Typically these are written to a http.Request.
*/
type InviteUserAccountsParams struct {

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

// WithDefaults hydrates default values in the invite user accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InviteUserAccountsParams) WithDefaults() *InviteUserAccountsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the invite user accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InviteUserAccountsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the invite user accounts params
func (o *InviteUserAccountsParams) WithTimeout(timeout time.Duration) *InviteUserAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the invite user accounts params
func (o *InviteUserAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the invite user accounts params
func (o *InviteUserAccountsParams) WithContext(ctx context.Context) *InviteUserAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the invite user accounts params
func (o *InviteUserAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the invite user accounts params
func (o *InviteUserAccountsParams) WithHTTPClient(client *http.Client) *InviteUserAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the invite user accounts params
func (o *InviteUserAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the invite user accounts params
func (o *InviteUserAccountsParams) WithBody(body []*models.UserToCreate) *InviteUserAccountsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the invite user accounts params
func (o *InviteUserAccountsParams) SetBody(body []*models.UserToCreate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *InviteUserAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
