// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// SearchAlertDeletedEntitiesReader is a Reader for the SearchAlertDeletedEntities structure.
type SearchAlertDeletedEntitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchAlertDeletedEntitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchAlertDeletedEntitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchAlertDeletedEntitiesOK creates a SearchAlertDeletedEntitiesOK with default headers values
func NewSearchAlertDeletedEntitiesOK() *SearchAlertDeletedEntitiesOK {
	return &SearchAlertDeletedEntitiesOK{}
}

/*
SearchAlertDeletedEntitiesOK describes a response with status code 200, with default header values.

successful operation
*/
type SearchAlertDeletedEntitiesOK struct {
	Payload *models.ResponseContainerPagedAlert
}

// IsSuccess returns true when this search alert deleted entities o k response has a 2xx status code
func (o *SearchAlertDeletedEntitiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search alert deleted entities o k response has a 3xx status code
func (o *SearchAlertDeletedEntitiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search alert deleted entities o k response has a 4xx status code
func (o *SearchAlertDeletedEntitiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search alert deleted entities o k response has a 5xx status code
func (o *SearchAlertDeletedEntitiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search alert deleted entities o k response a status code equal to that given
func (o *SearchAlertDeletedEntitiesOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchAlertDeletedEntitiesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/search/alert/deleted][%d] searchAlertDeletedEntitiesOK  %+v", 200, o.Payload)
}

func (o *SearchAlertDeletedEntitiesOK) String() string {
	return fmt.Sprintf("[POST /api/v2/search/alert/deleted][%d] searchAlertDeletedEntitiesOK  %+v", 200, o.Payload)
}

func (o *SearchAlertDeletedEntitiesOK) GetPayload() *models.ResponseContainerPagedAlert {
	return o.Payload
}

func (o *SearchAlertDeletedEntitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerPagedAlert)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
