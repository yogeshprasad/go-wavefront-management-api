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

// RevokePermissionFromUsersReader is a Reader for the RevokePermissionFromUsers structure.
type RevokePermissionFromUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RevokePermissionFromUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRevokePermissionFromUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRevokePermissionFromUsersOK creates a RevokePermissionFromUsersOK with default headers values
func NewRevokePermissionFromUsersOK() *RevokePermissionFromUsersOK {
	return &RevokePermissionFromUsersOK{}
}

/*
RevokePermissionFromUsersOK describes a response with status code 200, with default header values.

successful operation
*/
type RevokePermissionFromUsersOK struct {
	Payload *models.UserModel
}

// IsSuccess returns true when this revoke permission from users o k response has a 2xx status code
func (o *RevokePermissionFromUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this revoke permission from users o k response has a 3xx status code
func (o *RevokePermissionFromUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this revoke permission from users o k response has a 4xx status code
func (o *RevokePermissionFromUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this revoke permission from users o k response has a 5xx status code
func (o *RevokePermissionFromUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this revoke permission from users o k response a status code equal to that given
func (o *RevokePermissionFromUsersOK) IsCode(code int) bool {
	return code == 200
}

func (o *RevokePermissionFromUsersOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/user/revoke/{permission}][%d] revokePermissionFromUsersOK  %+v", 200, o.Payload)
}

func (o *RevokePermissionFromUsersOK) String() string {
	return fmt.Sprintf("[POST /api/v2/user/revoke/{permission}][%d] revokePermissionFromUsersOK  %+v", 200, o.Payload)
}

func (o *RevokePermissionFromUsersOK) GetPayload() *models.UserModel {
	return o.Payload
}

func (o *RevokePermissionFromUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
