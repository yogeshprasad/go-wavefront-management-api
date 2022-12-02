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

// DefaultCustomerTracesSearchReader is a Reader for the DefaultCustomerTracesSearch structure.
type DefaultCustomerTracesSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DefaultCustomerTracesSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDefaultCustomerTracesSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDefaultCustomerTracesSearchOK creates a DefaultCustomerTracesSearchOK with default headers values
func NewDefaultCustomerTracesSearchOK() *DefaultCustomerTracesSearchOK {
	return &DefaultCustomerTracesSearchOK{}
}

/*
DefaultCustomerTracesSearchOK describes a response with status code 200, with default header values.

successful operation
*/
type DefaultCustomerTracesSearchOK struct {
	Payload *models.ResponseContainerString
}

// IsSuccess returns true when this default customer traces search o k response has a 2xx status code
func (o *DefaultCustomerTracesSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this default customer traces search o k response has a 3xx status code
func (o *DefaultCustomerTracesSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this default customer traces search o k response has a 4xx status code
func (o *DefaultCustomerTracesSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this default customer traces search o k response has a 5xx status code
func (o *DefaultCustomerTracesSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this default customer traces search o k response a status code equal to that given
func (o *DefaultCustomerTracesSearchOK) IsCode(code int) bool {
	return code == 200
}

func (o *DefaultCustomerTracesSearchOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/savedtracessearch/defaultCustomerTracesSearch][%d] defaultCustomerTracesSearchOK  %+v", 200, o.Payload)
}

func (o *DefaultCustomerTracesSearchOK) String() string {
	return fmt.Sprintf("[POST /api/v2/savedtracessearch/defaultCustomerTracesSearch][%d] defaultCustomerTracesSearchOK  %+v", 200, o.Payload)
}

func (o *DefaultCustomerTracesSearchOK) GetPayload() *models.ResponseContainerString {
	return o.Payload
}

func (o *DefaultCustomerTracesSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerString)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
