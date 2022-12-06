// Code generated by go-swagger; DO NOT EDIT.

package cloud_integration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetCloudIntegrationReader is a Reader for the GetCloudIntegration structure.
type GetCloudIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCloudIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCloudIntegrationOK creates a GetCloudIntegrationOK with default headers values
func NewGetCloudIntegrationOK() *GetCloudIntegrationOK {
	return &GetCloudIntegrationOK{}
}

/*
GetCloudIntegrationOK describes a response with status code 200, with default header values.

successful operation
*/
type GetCloudIntegrationOK struct {
	Payload *models.ResponseContainerCloudIntegration
}

// IsSuccess returns true when this get cloud integration o k response has a 2xx status code
func (o *GetCloudIntegrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cloud integration o k response has a 3xx status code
func (o *GetCloudIntegrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cloud integration o k response has a 4xx status code
func (o *GetCloudIntegrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cloud integration o k response has a 5xx status code
func (o *GetCloudIntegrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cloud integration o k response a status code equal to that given
func (o *GetCloudIntegrationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCloudIntegrationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/cloudintegration/{id}][%d] getCloudIntegrationOK  %+v", 200, o.Payload)
}

func (o *GetCloudIntegrationOK) String() string {
	return fmt.Sprintf("[GET /api/v2/cloudintegration/{id}][%d] getCloudIntegrationOK  %+v", 200, o.Payload)
}

func (o *GetCloudIntegrationOK) GetPayload() *models.ResponseContainerCloudIntegration {
	return o.Payload
}

func (o *GetCloudIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerCloudIntegration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
