// Code generated by go-swagger; DO NOT EDIT.

package saved_traces_search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// DefaultTracesSearchReader is a Reader for the DefaultTracesSearch structure.
type DefaultTracesSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DefaultTracesSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDefaultTracesSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDefaultTracesSearchOK creates a DefaultTracesSearchOK with default headers values
func NewDefaultTracesSearchOK() *DefaultTracesSearchOK {
	return &DefaultTracesSearchOK{}
}

/*
DefaultTracesSearchOK describes a response with status code 200, with default header values.

successful operation
*/
type DefaultTracesSearchOK struct {
	Payload *models.ResponseContainerDefaultSavedTracesSearch
}

// IsSuccess returns true when this default traces search o k response has a 2xx status code
func (o *DefaultTracesSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this default traces search o k response has a 3xx status code
func (o *DefaultTracesSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this default traces search o k response has a 4xx status code
func (o *DefaultTracesSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this default traces search o k response has a 5xx status code
func (o *DefaultTracesSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this default traces search o k response a status code equal to that given
func (o *DefaultTracesSearchOK) IsCode(code int) bool {
	return code == 200
}

func (o *DefaultTracesSearchOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/savedtracessearch/defaultTracesSearch][%d] defaultTracesSearchOK  %+v", 200, o.Payload)
}

func (o *DefaultTracesSearchOK) String() string {
	return fmt.Sprintf("[GET /api/v2/savedtracessearch/defaultTracesSearch][%d] defaultTracesSearchOK  %+v", 200, o.Payload)
}

func (o *DefaultTracesSearchOK) GetPayload() *models.ResponseContainerDefaultSavedTracesSearch {
	return o.Payload
}

func (o *DefaultTracesSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerDefaultSavedTracesSearch)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
