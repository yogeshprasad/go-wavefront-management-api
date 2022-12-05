// Code generated by go-swagger; DO NOT EDIT.

package proxy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetAllProxyReader is a Reader for the GetAllProxy structure.
type GetAllProxyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllProxyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllProxyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllProxyOK creates a GetAllProxyOK with default headers values
func NewGetAllProxyOK() *GetAllProxyOK {
	return &GetAllProxyOK{}
}

/*
GetAllProxyOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAllProxyOK struct {
	Payload *models.ResponseContainerPagedProxy
}

// IsSuccess returns true when this get all proxy o k response has a 2xx status code
func (o *GetAllProxyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all proxy o k response has a 3xx status code
func (o *GetAllProxyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all proxy o k response has a 4xx status code
func (o *GetAllProxyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all proxy o k response has a 5xx status code
func (o *GetAllProxyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all proxy o k response a status code equal to that given
func (o *GetAllProxyOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAllProxyOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/proxy][%d] getAllProxyOK  %+v", 200, o.Payload)
}

func (o *GetAllProxyOK) String() string {
	return fmt.Sprintf("[GET /api/v2/proxy][%d] getAllProxyOK  %+v", 200, o.Payload)
}

func (o *GetAllProxyOK) GetPayload() *models.ResponseContainerPagedProxy {
	return o.Payload
}

func (o *GetAllProxyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerPagedProxy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
