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

// RevertMetricsPolicyByVersionReader is a Reader for the RevertMetricsPolicyByVersion structure.
type RevertMetricsPolicyByVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RevertMetricsPolicyByVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRevertMetricsPolicyByVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRevertMetricsPolicyByVersionOK creates a RevertMetricsPolicyByVersionOK with default headers values
func NewRevertMetricsPolicyByVersionOK() *RevertMetricsPolicyByVersionOK {
	return &RevertMetricsPolicyByVersionOK{}
}

/*
RevertMetricsPolicyByVersionOK describes a response with status code 200, with default header values.

successful operation
*/
type RevertMetricsPolicyByVersionOK struct {
	Payload *models.ResponseContainerMetricsPolicyReadModel
}

// IsSuccess returns true when this revert metrics policy by version o k response has a 2xx status code
func (o *RevertMetricsPolicyByVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this revert metrics policy by version o k response has a 3xx status code
func (o *RevertMetricsPolicyByVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this revert metrics policy by version o k response has a 4xx status code
func (o *RevertMetricsPolicyByVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this revert metrics policy by version o k response has a 5xx status code
func (o *RevertMetricsPolicyByVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this revert metrics policy by version o k response a status code equal to that given
func (o *RevertMetricsPolicyByVersionOK) IsCode(code int) bool {
	return code == 200
}

func (o *RevertMetricsPolicyByVersionOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/metricspolicy/revert/{version}][%d] revertMetricsPolicyByVersionOK  %+v", 200, o.Payload)
}

func (o *RevertMetricsPolicyByVersionOK) String() string {
	return fmt.Sprintf("[POST /api/v2/metricspolicy/revert/{version}][%d] revertMetricsPolicyByVersionOK  %+v", 200, o.Payload)
}

func (o *RevertMetricsPolicyByVersionOK) GetPayload() *models.ResponseContainerMetricsPolicyReadModel {
	return o.Payload
}

func (o *RevertMetricsPolicyByVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerMetricsPolicyReadModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
