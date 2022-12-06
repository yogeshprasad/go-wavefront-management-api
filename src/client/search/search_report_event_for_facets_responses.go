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

// SearchReportEventForFacetsReader is a Reader for the SearchReportEventForFacets structure.
type SearchReportEventForFacetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchReportEventForFacetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchReportEventForFacetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchReportEventForFacetsOK creates a SearchReportEventForFacetsOK with default headers values
func NewSearchReportEventForFacetsOK() *SearchReportEventForFacetsOK {
	return &SearchReportEventForFacetsOK{}
}

/*
SearchReportEventForFacetsOK describes a response with status code 200, with default header values.

successful operation
*/
type SearchReportEventForFacetsOK struct {
	Payload *models.ResponseContainerFacetsResponseContainer
}

// IsSuccess returns true when this search report event for facets o k response has a 2xx status code
func (o *SearchReportEventForFacetsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search report event for facets o k response has a 3xx status code
func (o *SearchReportEventForFacetsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search report event for facets o k response has a 4xx status code
func (o *SearchReportEventForFacetsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search report event for facets o k response has a 5xx status code
func (o *SearchReportEventForFacetsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search report event for facets o k response a status code equal to that given
func (o *SearchReportEventForFacetsOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchReportEventForFacetsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/search/event/facets][%d] searchReportEventForFacetsOK  %+v", 200, o.Payload)
}

func (o *SearchReportEventForFacetsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/search/event/facets][%d] searchReportEventForFacetsOK  %+v", 200, o.Payload)
}

func (o *SearchReportEventForFacetsOK) GetPayload() *models.ResponseContainerFacetsResponseContainer {
	return o.Payload
}

func (o *SearchReportEventForFacetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerFacetsResponseContainer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
