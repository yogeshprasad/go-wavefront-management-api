// Code generated by go-swagger; DO NOT EDIT.

package metrics_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetMetricsPolicyByVersionReader is a Reader for the GetMetricsPolicyByVersion structure.
type GetMetricsPolicyByVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMetricsPolicyByVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMetricsPolicyByVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMetricsPolicyByVersionOK creates a GetMetricsPolicyByVersionOK with default headers values
func NewGetMetricsPolicyByVersionOK() *GetMetricsPolicyByVersionOK {
	return &GetMetricsPolicyByVersionOK{}
}

/*
GetMetricsPolicyByVersionOK describes a response with status code 200, with default header values.

successful operation
*/
type GetMetricsPolicyByVersionOK struct {
	Payload *models.ResponseContainerMetricsPolicyReadModel
}

// IsSuccess returns true when this get metrics policy by version o k response has a 2xx status code
func (o *GetMetricsPolicyByVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get metrics policy by version o k response has a 3xx status code
func (o *GetMetricsPolicyByVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get metrics policy by version o k response has a 4xx status code
func (o *GetMetricsPolicyByVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get metrics policy by version o k response has a 5xx status code
func (o *GetMetricsPolicyByVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get metrics policy by version o k response a status code equal to that given
func (o *GetMetricsPolicyByVersionOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetMetricsPolicyByVersionOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/metricspolicy/history/{version}][%d] getMetricsPolicyByVersionOK  %+v", 200, o.Payload)
}

func (o *GetMetricsPolicyByVersionOK) String() string {
	return fmt.Sprintf("[GET /api/v2/metricspolicy/history/{version}][%d] getMetricsPolicyByVersionOK  %+v", 200, o.Payload)
}

func (o *GetMetricsPolicyByVersionOK) GetPayload() *models.ResponseContainerMetricsPolicyReadModel {
	return o.Payload
}

func (o *GetMetricsPolicyByVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerMetricsPolicyReadModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
