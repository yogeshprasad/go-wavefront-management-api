// Code generated by go-swagger; DO NOT EDIT.

package monitored_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetServicesOfApplicationReader is a Reader for the GetServicesOfApplication structure.
type GetServicesOfApplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServicesOfApplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServicesOfApplicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetServicesOfApplicationOK creates a GetServicesOfApplicationOK with default headers values
func NewGetServicesOfApplicationOK() *GetServicesOfApplicationOK {
	return &GetServicesOfApplicationOK{}
}

/*
GetServicesOfApplicationOK describes a response with status code 200, with default header values.

successful operation
*/
type GetServicesOfApplicationOK struct {
	Payload *models.ResponseContainerPagedMonitoredServiceDTO
}

// IsSuccess returns true when this get services of application o k response has a 2xx status code
func (o *GetServicesOfApplicationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get services of application o k response has a 3xx status code
func (o *GetServicesOfApplicationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get services of application o k response has a 4xx status code
func (o *GetServicesOfApplicationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get services of application o k response has a 5xx status code
func (o *GetServicesOfApplicationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get services of application o k response a status code equal to that given
func (o *GetServicesOfApplicationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetServicesOfApplicationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/monitoredservice/{application}][%d] getServicesOfApplicationOK  %+v", 200, o.Payload)
}

func (o *GetServicesOfApplicationOK) String() string {
	return fmt.Sprintf("[GET /api/v2/monitoredservice/{application}][%d] getServicesOfApplicationOK  %+v", 200, o.Payload)
}

func (o *GetServicesOfApplicationOK) GetPayload() *models.ResponseContainerPagedMonitoredServiceDTO {
	return o.Payload
}

func (o *GetServicesOfApplicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerPagedMonitoredServiceDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
