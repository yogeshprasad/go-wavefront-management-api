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

// SearchDashboardForFacetsReader is a Reader for the SearchDashboardForFacets structure.
type SearchDashboardForFacetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchDashboardForFacetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchDashboardForFacetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchDashboardForFacetsOK creates a SearchDashboardForFacetsOK with default headers values
func NewSearchDashboardForFacetsOK() *SearchDashboardForFacetsOK {
	return &SearchDashboardForFacetsOK{}
}

/*
SearchDashboardForFacetsOK describes a response with status code 200, with default header values.

successful operation
*/
type SearchDashboardForFacetsOK struct {
	Payload *models.ResponseContainerFacetsResponseContainer
}

// IsSuccess returns true when this search dashboard for facets o k response has a 2xx status code
func (o *SearchDashboardForFacetsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search dashboard for facets o k response has a 3xx status code
func (o *SearchDashboardForFacetsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search dashboard for facets o k response has a 4xx status code
func (o *SearchDashboardForFacetsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search dashboard for facets o k response has a 5xx status code
func (o *SearchDashboardForFacetsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search dashboard for facets o k response a status code equal to that given
func (o *SearchDashboardForFacetsOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchDashboardForFacetsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/search/dashboard/facets][%d] searchDashboardForFacetsOK  %+v", 200, o.Payload)
}

func (o *SearchDashboardForFacetsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/search/dashboard/facets][%d] searchDashboardForFacetsOK  %+v", 200, o.Payload)
}

func (o *SearchDashboardForFacetsOK) GetPayload() *models.ResponseContainerFacetsResponseContainer {
	return o.Payload
}

func (o *SearchDashboardForFacetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerFacetsResponseContainer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
