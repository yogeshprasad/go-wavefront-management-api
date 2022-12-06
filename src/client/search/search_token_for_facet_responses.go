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

// SearchTokenForFacetReader is a Reader for the SearchTokenForFacet structure.
type SearchTokenForFacetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchTokenForFacetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchTokenForFacetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchTokenForFacetOK creates a SearchTokenForFacetOK with default headers values
func NewSearchTokenForFacetOK() *SearchTokenForFacetOK {
	return &SearchTokenForFacetOK{}
}

/*
SearchTokenForFacetOK describes a response with status code 200, with default header values.

successful operation
*/
type SearchTokenForFacetOK struct {
	Payload *models.ResponseContainerFacetResponse
}

// IsSuccess returns true when this search token for facet o k response has a 2xx status code
func (o *SearchTokenForFacetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search token for facet o k response has a 3xx status code
func (o *SearchTokenForFacetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search token for facet o k response has a 4xx status code
func (o *SearchTokenForFacetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search token for facet o k response has a 5xx status code
func (o *SearchTokenForFacetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search token for facet o k response a status code equal to that given
func (o *SearchTokenForFacetOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchTokenForFacetOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/search/token/{facet}][%d] searchTokenForFacetOK  %+v", 200, o.Payload)
}

func (o *SearchTokenForFacetOK) String() string {
	return fmt.Sprintf("[POST /api/v2/search/token/{facet}][%d] searchTokenForFacetOK  %+v", 200, o.Payload)
}

func (o *SearchTokenForFacetOK) GetPayload() *models.ResponseContainerFacetResponse {
	return o.Payload
}

func (o *SearchTokenForFacetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerFacetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
