// Code generated by go-swagger; DO NOT EDIT.

package source

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetAllSourceReader is a Reader for the GetAllSource structure.
type GetAllSourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllSourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllSourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllSourceOK creates a GetAllSourceOK with default headers values
func NewGetAllSourceOK() *GetAllSourceOK {
	return &GetAllSourceOK{}
}

/*
GetAllSourceOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAllSourceOK struct {
	Payload *models.ResponseContainerPagedSource
}

// IsSuccess returns true when this get all source o k response has a 2xx status code
func (o *GetAllSourceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all source o k response has a 3xx status code
func (o *GetAllSourceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all source o k response has a 4xx status code
func (o *GetAllSourceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all source o k response has a 5xx status code
func (o *GetAllSourceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all source o k response a status code equal to that given
func (o *GetAllSourceOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAllSourceOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/source][%d] getAllSourceOK  %+v", 200, o.Payload)
}

func (o *GetAllSourceOK) String() string {
	return fmt.Sprintf("[GET /api/v2/source][%d] getAllSourceOK  %+v", 200, o.Payload)
}

func (o *GetAllSourceOK) GetPayload() *models.ResponseContainerPagedSource {
	return o.Payload
}

func (o *GetAllSourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerPagedSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
