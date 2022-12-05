// Code generated by go-swagger; DO NOT EDIT.

package saved_search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// CreateSavedSearchReader is a Reader for the CreateSavedSearch structure.
type CreateSavedSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSavedSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSavedSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateSavedSearchOK creates a CreateSavedSearchOK with default headers values
func NewCreateSavedSearchOK() *CreateSavedSearchOK {
	return &CreateSavedSearchOK{}
}

/*
CreateSavedSearchOK describes a response with status code 200, with default header values.

successful operation
*/
type CreateSavedSearchOK struct {
	Payload *models.ResponseContainerSavedSearch
}

// IsSuccess returns true when this create saved search o k response has a 2xx status code
func (o *CreateSavedSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create saved search o k response has a 3xx status code
func (o *CreateSavedSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create saved search o k response has a 4xx status code
func (o *CreateSavedSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create saved search o k response has a 5xx status code
func (o *CreateSavedSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create saved search o k response a status code equal to that given
func (o *CreateSavedSearchOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateSavedSearchOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/savedsearch][%d] createSavedSearchOK  %+v", 200, o.Payload)
}

func (o *CreateSavedSearchOK) String() string {
	return fmt.Sprintf("[POST /api/v2/savedsearch][%d] createSavedSearchOK  %+v", 200, o.Payload)
}

func (o *CreateSavedSearchOK) GetPayload() *models.ResponseContainerSavedSearch {
	return o.Payload
}

func (o *CreateSavedSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerSavedSearch)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
