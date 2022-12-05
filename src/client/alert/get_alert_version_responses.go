// Code generated by go-swagger; DO NOT EDIT.

package alert

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetAlertVersionReader is a Reader for the GetAlertVersion structure.
type GetAlertVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAlertVersionOK creates a GetAlertVersionOK with default headers values
func NewGetAlertVersionOK() *GetAlertVersionOK {
	return &GetAlertVersionOK{}
}

/*
GetAlertVersionOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAlertVersionOK struct {
	Payload *models.ResponseContainerAlert
}

// IsSuccess returns true when this get alert version o k response has a 2xx status code
func (o *GetAlertVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get alert version o k response has a 3xx status code
func (o *GetAlertVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert version o k response has a 4xx status code
func (o *GetAlertVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alert version o k response has a 5xx status code
func (o *GetAlertVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert version o k response a status code equal to that given
func (o *GetAlertVersionOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAlertVersionOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/alert/{id}/history/{version}][%d] getAlertVersionOK  %+v", 200, o.Payload)
}

func (o *GetAlertVersionOK) String() string {
	return fmt.Sprintf("[GET /api/v2/alert/{id}/history/{version}][%d] getAlertVersionOK  %+v", 200, o.Payload)
}

func (o *GetAlertVersionOK) GetPayload() *models.ResponseContainerAlert {
	return o.Payload
}

func (o *GetAlertVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerAlert)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
