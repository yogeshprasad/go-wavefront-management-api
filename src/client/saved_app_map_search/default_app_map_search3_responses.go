// Code generated by go-swagger; DO NOT EDIT.

package saved_app_map_search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// DefaultAppMapSearch3Reader is a Reader for the DefaultAppMapSearch3 structure.
type DefaultAppMapSearch3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DefaultAppMapSearch3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDefaultAppMapSearch3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDefaultAppMapSearch3OK creates a DefaultAppMapSearch3OK with default headers values
func NewDefaultAppMapSearch3OK() *DefaultAppMapSearch3OK {
	return &DefaultAppMapSearch3OK{}
}

/*
DefaultAppMapSearch3OK describes a response with status code 200, with default header values.

successful operation
*/
type DefaultAppMapSearch3OK struct {
	Payload *models.ResponseContainerDefaultSavedAppMapSearch
}

// IsSuccess returns true when this default app map search3 o k response has a 2xx status code
func (o *DefaultAppMapSearch3OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this default app map search3 o k response has a 3xx status code
func (o *DefaultAppMapSearch3OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this default app map search3 o k response has a 4xx status code
func (o *DefaultAppMapSearch3OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this default app map search3 o k response has a 5xx status code
func (o *DefaultAppMapSearch3OK) IsServerError() bool {
	return false
}

// IsCode returns true when this default app map search3 o k response a status code equal to that given
func (o *DefaultAppMapSearch3OK) IsCode(code int) bool {
	return code == 200
}

func (o *DefaultAppMapSearch3OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/savedappmapsearch/defaultAppMapSearch][%d] defaultAppMapSearch3OK  %+v", 200, o.Payload)
}

func (o *DefaultAppMapSearch3OK) String() string {
	return fmt.Sprintf("[GET /api/v2/savedappmapsearch/defaultAppMapSearch][%d] defaultAppMapSearch3OK  %+v", 200, o.Payload)
}

func (o *DefaultAppMapSearch3OK) GetPayload() *models.ResponseContainerDefaultSavedAppMapSearch {
	return o.Payload
}

func (o *DefaultAppMapSearch3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerDefaultSavedAppMapSearch)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
