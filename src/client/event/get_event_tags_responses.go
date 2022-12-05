// Code generated by go-swagger; DO NOT EDIT.

package event

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetEventTagsReader is a Reader for the GetEventTags structure.
type GetEventTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEventTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEventTagsOK creates a GetEventTagsOK with default headers values
func NewGetEventTagsOK() *GetEventTagsOK {
	return &GetEventTagsOK{}
}

/*
GetEventTagsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetEventTagsOK struct {
	Payload *models.ResponseContainerTagsResponse
}

// IsSuccess returns true when this get event tags o k response has a 2xx status code
func (o *GetEventTagsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get event tags o k response has a 3xx status code
func (o *GetEventTagsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get event tags o k response has a 4xx status code
func (o *GetEventTagsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get event tags o k response has a 5xx status code
func (o *GetEventTagsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get event tags o k response a status code equal to that given
func (o *GetEventTagsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetEventTagsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/event/{id}/tag][%d] getEventTagsOK  %+v", 200, o.Payload)
}

func (o *GetEventTagsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/event/{id}/tag][%d] getEventTagsOK  %+v", 200, o.Payload)
}

func (o *GetEventTagsOK) GetPayload() *models.ResponseContainerTagsResponse {
	return o.Payload
}

func (o *GetEventTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerTagsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
