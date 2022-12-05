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

// SearchServiceAccountEntitiesReader is a Reader for the SearchServiceAccountEntities structure.
type SearchServiceAccountEntitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchServiceAccountEntitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchServiceAccountEntitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchServiceAccountEntitiesOK creates a SearchServiceAccountEntitiesOK with default headers values
func NewSearchServiceAccountEntitiesOK() *SearchServiceAccountEntitiesOK {
	return &SearchServiceAccountEntitiesOK{}
}

/*
SearchServiceAccountEntitiesOK describes a response with status code 200, with default header values.

successful operation
*/
type SearchServiceAccountEntitiesOK struct {
	Payload *models.ResponseContainerPagedServiceAccount
}

// IsSuccess returns true when this search service account entities o k response has a 2xx status code
func (o *SearchServiceAccountEntitiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search service account entities o k response has a 3xx status code
func (o *SearchServiceAccountEntitiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search service account entities o k response has a 4xx status code
func (o *SearchServiceAccountEntitiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search service account entities o k response has a 5xx status code
func (o *SearchServiceAccountEntitiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search service account entities o k response a status code equal to that given
func (o *SearchServiceAccountEntitiesOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchServiceAccountEntitiesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/search/serviceaccount][%d] searchServiceAccountEntitiesOK  %+v", 200, o.Payload)
}

func (o *SearchServiceAccountEntitiesOK) String() string {
	return fmt.Sprintf("[POST /api/v2/search/serviceaccount][%d] searchServiceAccountEntitiesOK  %+v", 200, o.Payload)
}

func (o *SearchServiceAccountEntitiesOK) GetPayload() *models.ResponseContainerPagedServiceAccount {
	return o.Payload
}

func (o *SearchServiceAccountEntitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerPagedServiceAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
