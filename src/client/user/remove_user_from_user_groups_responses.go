// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// RemoveUserFromUserGroupsReader is a Reader for the RemoveUserFromUserGroups structure.
type RemoveUserFromUserGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveUserFromUserGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveUserFromUserGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoveUserFromUserGroupsOK creates a RemoveUserFromUserGroupsOK with default headers values
func NewRemoveUserFromUserGroupsOK() *RemoveUserFromUserGroupsOK {
	return &RemoveUserFromUserGroupsOK{}
}

/*
RemoveUserFromUserGroupsOK describes a response with status code 200, with default header values.

successful operation
*/
type RemoveUserFromUserGroupsOK struct {
	Payload *models.UserModel
}

// IsSuccess returns true when this remove user from user groups o k response has a 2xx status code
func (o *RemoveUserFromUserGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove user from user groups o k response has a 3xx status code
func (o *RemoveUserFromUserGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove user from user groups o k response has a 4xx status code
func (o *RemoveUserFromUserGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove user from user groups o k response has a 5xx status code
func (o *RemoveUserFromUserGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this remove user from user groups o k response a status code equal to that given
func (o *RemoveUserFromUserGroupsOK) IsCode(code int) bool {
	return code == 200
}

func (o *RemoveUserFromUserGroupsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/user/{id}/removeUserGroups][%d] removeUserFromUserGroupsOK  %+v", 200, o.Payload)
}

func (o *RemoveUserFromUserGroupsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/user/{id}/removeUserGroups][%d] removeUserFromUserGroupsOK  %+v", 200, o.Payload)
}

func (o *RemoveUserFromUserGroupsOK) GetPayload() *models.UserModel {
	return o.Payload
}

func (o *RemoveUserFromUserGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
