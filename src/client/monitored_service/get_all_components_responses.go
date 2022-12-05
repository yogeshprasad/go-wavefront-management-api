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

// GetAllComponentsReader is a Reader for the GetAllComponents structure.
type GetAllComponentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllComponentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllComponentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllComponentsOK creates a GetAllComponentsOK with default headers values
func NewGetAllComponentsOK() *GetAllComponentsOK {
	return &GetAllComponentsOK{}
}

/*
GetAllComponentsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAllComponentsOK struct {
	Payload *models.ResponseContainerPagedMonitoredServiceDTO
}

// IsSuccess returns true when this get all components o k response has a 2xx status code
func (o *GetAllComponentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all components o k response has a 3xx status code
func (o *GetAllComponentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all components o k response has a 4xx status code
func (o *GetAllComponentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all components o k response has a 5xx status code
func (o *GetAllComponentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all components o k response a status code equal to that given
func (o *GetAllComponentsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAllComponentsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/monitoredservice/components][%d] getAllComponentsOK  %+v", 200, o.Payload)
}

func (o *GetAllComponentsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/monitoredservice/components][%d] getAllComponentsOK  %+v", 200, o.Payload)
}

func (o *GetAllComponentsOK) GetPayload() *models.ResponseContainerPagedMonitoredServiceDTO {
	return o.Payload
}

func (o *GetAllComponentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerPagedMonitoredServiceDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
