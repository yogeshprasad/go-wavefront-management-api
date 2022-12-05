// Code generated by go-swagger; DO NOT EDIT.

package ingestion_spy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SpyOnDeltaCountersReader is a Reader for the SpyOnDeltaCounters structure.
type SpyOnDeltaCountersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SpyOnDeltaCountersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewSpyOnDeltaCountersDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewSpyOnDeltaCountersDefault creates a SpyOnDeltaCountersDefault with default headers values
func NewSpyOnDeltaCountersDefault(code int) *SpyOnDeltaCountersDefault {
	return &SpyOnDeltaCountersDefault{
		_statusCode: code,
	}
}

/*
SpyOnDeltaCountersDefault describes a response with status code -1, with default header values.

successful operation
*/
type SpyOnDeltaCountersDefault struct {
	_statusCode int
}

// Code gets the status code for the spy on delta counters default response
func (o *SpyOnDeltaCountersDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this spy on delta counters default response has a 2xx status code
func (o *SpyOnDeltaCountersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this spy on delta counters default response has a 3xx status code
func (o *SpyOnDeltaCountersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this spy on delta counters default response has a 4xx status code
func (o *SpyOnDeltaCountersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this spy on delta counters default response has a 5xx status code
func (o *SpyOnDeltaCountersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this spy on delta counters default response a status code equal to that given
func (o *SpyOnDeltaCountersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SpyOnDeltaCountersDefault) Error() string {
	return fmt.Sprintf("[GET /api/spy/deltas][%d] spyOnDeltaCounters default ", o._statusCode)
}

func (o *SpyOnDeltaCountersDefault) String() string {
	return fmt.Sprintf("[GET /api/spy/deltas][%d] spyOnDeltaCounters default ", o._statusCode)
}

func (o *SpyOnDeltaCountersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
