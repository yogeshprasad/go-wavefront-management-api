// Code generated by go-swagger; DO NOT EDIT.

package cloud_integration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new cloud integration API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cloud integration API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateAWSExternalID(params *CreateAWSExternalIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAWSExternalIDOK, error)

	CreateCloudIntegration(params *CreateCloudIntegrationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateCloudIntegrationOK, error)

	DeleteAWSExternalID(params *DeleteAWSExternalIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAWSExternalIDOK, error)

	DeleteCloudIntegration(params *DeleteCloudIntegrationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteCloudIntegrationOK, error)

	DisableCloudIntegration(params *DisableCloudIntegrationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DisableCloudIntegrationOK, error)

	EnableCloudIntegration(params *EnableCloudIntegrationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EnableCloudIntegrationOK, error)

	GetAWSExternalID(params *GetAWSExternalIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAWSExternalIDOK, error)

	GetAllCloudIntegration(params *GetAllCloudIntegrationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllCloudIntegrationOK, error)

	GetCloudIntegration(params *GetCloudIntegrationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCloudIntegrationOK, error)

	UndeleteCloudIntegration(params *UndeleteCloudIntegrationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UndeleteCloudIntegrationOK, error)

	UpdateCloudIntegration(params *UpdateCloudIntegrationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateCloudIntegrationOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateAWSExternalID creates an external id
*/
func (a *Client) CreateAWSExternalID(params *CreateAWSExternalIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAWSExternalIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAWSExternalIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createAWSExternalId",
		Method:             "POST",
		PathPattern:        "/api/v2/cloudintegration/awsExternalId",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateAWSExternalIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateAWSExternalIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createAWSExternalId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateCloudIntegration creates a cloud integration
*/
func (a *Client) CreateCloudIntegration(params *CreateCloudIntegrationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateCloudIntegrationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCloudIntegrationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createCloudIntegration",
		Method:             "POST",
		PathPattern:        "/api/v2/cloudintegration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateCloudIntegrationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateCloudIntegrationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createCloudIntegration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteAWSExternalID ds e l e t es an external id that was created by wavefront
*/
func (a *Client) DeleteAWSExternalID(params *DeleteAWSExternalIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAWSExternalIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAWSExternalIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteAWSExternalId",
		Method:             "DELETE",
		PathPattern:        "/api/v2/cloudintegration/awsExternalId/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteAWSExternalIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteAWSExternalIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteAWSExternalId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteCloudIntegration deletes a specific cloud integration
*/
func (a *Client) DeleteCloudIntegration(params *DeleteCloudIntegrationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteCloudIntegrationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCloudIntegrationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteCloudIntegration",
		Method:             "DELETE",
		PathPattern:        "/api/v2/cloudintegration/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteCloudIntegrationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteCloudIntegrationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteCloudIntegration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DisableCloudIntegration disables a specific cloud integration
*/
func (a *Client) DisableCloudIntegration(params *DisableCloudIntegrationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DisableCloudIntegrationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisableCloudIntegrationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "disableCloudIntegration",
		Method:             "POST",
		PathPattern:        "/api/v2/cloudintegration/{id}/disable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DisableCloudIntegrationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DisableCloudIntegrationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for disableCloudIntegration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EnableCloudIntegration enables a specific cloud integration
*/
func (a *Client) EnableCloudIntegration(params *EnableCloudIntegrationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EnableCloudIntegrationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnableCloudIntegrationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "enableCloudIntegration",
		Method:             "POST",
		PathPattern:        "/api/v2/cloudintegration/{id}/enable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EnableCloudIntegrationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EnableCloudIntegrationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for enableCloudIntegration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAWSExternalID gs e ts confirms a valid external id that was created by wavefront
*/
func (a *Client) GetAWSExternalID(params *GetAWSExternalIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAWSExternalIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAWSExternalIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAWSExternalId",
		Method:             "GET",
		PathPattern:        "/api/v2/cloudintegration/awsExternalId/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAWSExternalIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAWSExternalIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAWSExternalId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAllCloudIntegration gets all cloud integrations for a customer
*/
func (a *Client) GetAllCloudIntegration(params *GetAllCloudIntegrationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAllCloudIntegrationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllCloudIntegrationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllCloudIntegration",
		Method:             "GET",
		PathPattern:        "/api/v2/cloudintegration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllCloudIntegrationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAllCloudIntegrationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllCloudIntegration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCloudIntegration gets a specific cloud integration
*/
func (a *Client) GetCloudIntegration(params *GetCloudIntegrationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCloudIntegrationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCloudIntegrationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCloudIntegration",
		Method:             "GET",
		PathPattern:        "/api/v2/cloudintegration/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCloudIntegrationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCloudIntegrationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCloudIntegration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UndeleteCloudIntegration undeletes a specific cloud integration
*/
func (a *Client) UndeleteCloudIntegration(params *UndeleteCloudIntegrationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UndeleteCloudIntegrationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUndeleteCloudIntegrationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "undeleteCloudIntegration",
		Method:             "POST",
		PathPattern:        "/api/v2/cloudintegration/{id}/undelete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UndeleteCloudIntegrationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UndeleteCloudIntegrationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for undeleteCloudIntegration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateCloudIntegration updates a specific cloud integration
*/
func (a *Client) UpdateCloudIntegration(params *UpdateCloudIntegrationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateCloudIntegrationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCloudIntegrationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateCloudIntegration",
		Method:             "PUT",
		PathPattern:        "/api/v2/cloudintegration/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateCloudIntegrationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateCloudIntegrationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateCloudIntegration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
