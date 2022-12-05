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
	"github.com/go-openapi/swag"

	"models"
)

// NewCreateUserParams creates a new CreateUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateUserParams() *CreateUserParams {
	return &CreateUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUserParamsWithTimeout creates a new CreateUserParams object
// with the ability to set a timeout on a request.
func NewCreateUserParamsWithTimeout(timeout time.Duration) *CreateUserParams {
	return &CreateUserParams{
		timeout: timeout,
	}
}

// NewCreateUserParamsWithContext creates a new CreateUserParams object
// with the ability to set a context for a request.
func NewCreateUserParamsWithContext(ctx context.Context) *CreateUserParams {
	return &CreateUserParams{
		Context: ctx,
	}
}

// NewCreateUserParamsWithHTTPClient creates a new CreateUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateUserParamsWithHTTPClient(client *http.Client) *CreateUserParams {
	return &CreateUserParams{
		HTTPClient: client,
	}
}

/*
CreateUserParams contains all the parameters to send to the API endpoint

	for the create user operation.

	Typically these are written to a http.Request.
*/
type CreateUserParams struct {

	/* Body.

	     Example Body:
	<pre>{
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
	}</pre>
	*/
	Body *models.UserToCreate

	/* SendEmail.

	   Whether to send email notification to the user, if created.  Default: false
	*/
	SendEmail *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUserParams) WithDefaults() *CreateUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create user params
func (o *CreateUserParams) WithTimeout(timeout time.Duration) *CreateUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create user params
func (o *CreateUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create user params
func (o *CreateUserParams) WithContext(ctx context.Context) *CreateUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create user params
func (o *CreateUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create user params
func (o *CreateUserParams) WithHTTPClient(client *http.Client) *CreateUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create user params
func (o *CreateUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create user params
func (o *CreateUserParams) WithBody(body *models.UserToCreate) *CreateUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create user params
func (o *CreateUserParams) SetBody(body *models.UserToCreate) {
	o.Body = body
}

// WithSendEmail adds the sendEmail to the create user params
func (o *CreateUserParams) WithSendEmail(sendEmail *bool) *CreateUserParams {
	o.SetSendEmail(sendEmail)
	return o
}

// SetSendEmail adds the sendEmail to the create user params
func (o *CreateUserParams) SetSendEmail(sendEmail *bool) {
	o.SendEmail = sendEmail
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.SendEmail != nil {

		// query param sendEmail
		var qrSendEmail bool

		if o.SendEmail != nil {
			qrSendEmail = *o.SendEmail
		}
		qSendEmail := swag.FormatBool(qrSendEmail)
		if qSendEmail != "" {

			if err := r.SetQueryParam("sendEmail", qSendEmail); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
