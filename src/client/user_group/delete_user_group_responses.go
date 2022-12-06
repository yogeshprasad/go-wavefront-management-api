// Code generated by go-swagger; DO NOT EDIT.

package user_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// DeleteUserGroupReader is a Reader for the DeleteUserGroup structure.
type DeleteUserGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUserGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserGroupOK creates a DeleteUserGroupOK with default headers values
func NewDeleteUserGroupOK() *DeleteUserGroupOK {
	return &DeleteUserGroupOK{}
}

/*
DeleteUserGroupOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteUserGroupOK struct {
	Payload *models.ResponseContainerUserGroupModel
}

// IsSuccess returns true when this delete user group o k response has a 2xx status code
func (o *DeleteUserGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete user group o k response has a 3xx status code
func (o *DeleteUserGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user group o k response has a 4xx status code
func (o *DeleteUserGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete user group o k response has a 5xx status code
func (o *DeleteUserGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user group o k response a status code equal to that given
func (o *DeleteUserGroupOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteUserGroupOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/usergroup/{id}][%d] deleteUserGroupOK  %+v", 200, o.Payload)
}

func (o *DeleteUserGroupOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/usergroup/{id}][%d] deleteUserGroupOK  %+v", 200, o.Payload)
}

func (o *DeleteUserGroupOK) GetPayload() *models.ResponseContainerUserGroupModel {
	return o.Payload
}

func (o *DeleteUserGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerUserGroupModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
