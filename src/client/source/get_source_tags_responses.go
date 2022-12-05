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

// GetSourceTagsReader is a Reader for the GetSourceTags structure.
type GetSourceTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSourceTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSourceTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSourceTagsOK creates a GetSourceTagsOK with default headers values
func NewGetSourceTagsOK() *GetSourceTagsOK {
	return &GetSourceTagsOK{}
}

/*
GetSourceTagsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetSourceTagsOK struct {
	Payload *models.ResponseContainerTagsResponse
}

// IsSuccess returns true when this get source tags o k response has a 2xx status code
func (o *GetSourceTagsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get source tags o k response has a 3xx status code
func (o *GetSourceTagsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get source tags o k response has a 4xx status code
func (o *GetSourceTagsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get source tags o k response has a 5xx status code
func (o *GetSourceTagsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get source tags o k response a status code equal to that given
func (o *GetSourceTagsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetSourceTagsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/source/{id}/tag][%d] getSourceTagsOK  %+v", 200, o.Payload)
}

func (o *GetSourceTagsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/source/{id}/tag][%d] getSourceTagsOK  %+v", 200, o.Payload)
}

func (o *GetSourceTagsOK) GetPayload() *models.ResponseContainerTagsResponse {
	return o.Payload
}

func (o *GetSourceTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerTagsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
