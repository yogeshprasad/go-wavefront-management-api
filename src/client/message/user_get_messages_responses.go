// Code generated by go-swagger; DO NOT EDIT.

package message

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// UserGetMessagesReader is a Reader for the UserGetMessages structure.
type UserGetMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserGetMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserGetMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserGetMessagesOK creates a UserGetMessagesOK with default headers values
func NewUserGetMessagesOK() *UserGetMessagesOK {
	return &UserGetMessagesOK{}
}

/*
UserGetMessagesOK describes a response with status code 200, with default header values.

successful operation
*/
type UserGetMessagesOK struct {
	Payload *models.ResponseContainerPagedMessage
}

// IsSuccess returns true when this user get messages o k response has a 2xx status code
func (o *UserGetMessagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user get messages o k response has a 3xx status code
func (o *UserGetMessagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user get messages o k response has a 4xx status code
func (o *UserGetMessagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user get messages o k response has a 5xx status code
func (o *UserGetMessagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user get messages o k response a status code equal to that given
func (o *UserGetMessagesOK) IsCode(code int) bool {
	return code == 200
}

func (o *UserGetMessagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/message][%d] userGetMessagesOK  %+v", 200, o.Payload)
}

func (o *UserGetMessagesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/message][%d] userGetMessagesOK  %+v", 200, o.Payload)
}

func (o *UserGetMessagesOK) GetPayload() *models.ResponseContainerPagedMessage {
	return o.Payload
}

func (o *UserGetMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerPagedMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
