// Code generated by go-swagger; DO NOT EDIT.

package event

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// UpdateUserEventReader is a Reader for the UpdateUserEvent structure.
type UpdateUserEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserEventOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateUserEventOK creates a UpdateUserEventOK with default headers values
func NewUpdateUserEventOK() *UpdateUserEventOK {
	return &UpdateUserEventOK{}
}

/*
UpdateUserEventOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateUserEventOK struct {
	Payload *models.ResponseContainerEvent
}

// IsSuccess returns true when this update user event o k response has a 2xx status code
func (o *UpdateUserEventOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update user event o k response has a 3xx status code
func (o *UpdateUserEventOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user event o k response has a 4xx status code
func (o *UpdateUserEventOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update user event o k response has a 5xx status code
func (o *UpdateUserEventOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update user event o k response a status code equal to that given
func (o *UpdateUserEventOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateUserEventOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/event/{id}][%d] updateUserEventOK  %+v", 200, o.Payload)
}

func (o *UpdateUserEventOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/event/{id}][%d] updateUserEventOK  %+v", 200, o.Payload)
}

func (o *UpdateUserEventOK) GetPayload() *models.ResponseContainerEvent {
	return o.Payload
}

func (o *UpdateUserEventOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
