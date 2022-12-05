// Code generated by go-swagger; DO NOT EDIT.

package metric

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetMetricDetailsReader is a Reader for the GetMetricDetails structure.
type GetMetricDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMetricDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMetricDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMetricDetailsOK creates a GetMetricDetailsOK with default headers values
func NewGetMetricDetailsOK() *GetMetricDetailsOK {
	return &GetMetricDetailsOK{}
}

/*
GetMetricDetailsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetMetricDetailsOK struct {
	Payload *models.MetricDetailsResponse
}

// IsSuccess returns true when this get metric details o k response has a 2xx status code
func (o *GetMetricDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get metric details o k response has a 3xx status code
func (o *GetMetricDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get metric details o k response has a 4xx status code
func (o *GetMetricDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get metric details o k response has a 5xx status code
func (o *GetMetricDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get metric details o k response a status code equal to that given
func (o *GetMetricDetailsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetMetricDetailsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart/metric/detail][%d] getMetricDetailsOK  %+v", 200, o.Payload)
}

func (o *GetMetricDetailsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/chart/metric/detail][%d] getMetricDetailsOK  %+v", 200, o.Payload)
}

func (o *GetMetricDetailsOK) GetPayload() *models.MetricDetailsResponse {
	return o.Payload
}

func (o *GetMetricDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetricDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
