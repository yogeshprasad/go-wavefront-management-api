// Code generated by go-swagger; DO NOT EDIT.

package query

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// QueryRawReader is a Reader for the QueryRaw structure.
type QueryRawReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryRawReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryRawOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryRawBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewQueryRawUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQueryRawNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryRawInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryRawOK creates a QueryRawOK with default headers values
func NewQueryRawOK() *QueryRawOK {
	return &QueryRawOK{}
}

/*
QueryRawOK describes a response with status code 200, with default header values.

successful operation
*/
type QueryRawOK struct {
	Payload []*models.RawTimeseries
}

// IsSuccess returns true when this query raw o k response has a 2xx status code
func (o *QueryRawOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query raw o k response has a 3xx status code
func (o *QueryRawOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query raw o k response has a 4xx status code
func (o *QueryRawOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query raw o k response has a 5xx status code
func (o *QueryRawOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query raw o k response a status code equal to that given
func (o *QueryRawOK) IsCode(code int) bool {
	return code == 200
}

func (o *QueryRawOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart/raw][%d] queryRawOK  %+v", 200, o.Payload)
}

func (o *QueryRawOK) String() string {
	return fmt.Sprintf("[GET /api/v2/chart/raw][%d] queryRawOK  %+v", 200, o.Payload)
}

func (o *QueryRawOK) GetPayload() []*models.RawTimeseries {
	return o.Payload
}

func (o *QueryRawOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryRawBadRequest creates a QueryRawBadRequest with default headers values
func NewQueryRawBadRequest() *QueryRawBadRequest {
	return &QueryRawBadRequest{}
}

/*
QueryRawBadRequest describes a response with status code 400, with default header values.

Invalid request parameters
*/
type QueryRawBadRequest struct {
}

// IsSuccess returns true when this query raw bad request response has a 2xx status code
func (o *QueryRawBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query raw bad request response has a 3xx status code
func (o *QueryRawBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query raw bad request response has a 4xx status code
func (o *QueryRawBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query raw bad request response has a 5xx status code
func (o *QueryRawBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query raw bad request response a status code equal to that given
func (o *QueryRawBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *QueryRawBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart/raw][%d] queryRawBadRequest ", 400)
}

func (o *QueryRawBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/chart/raw][%d] queryRawBadRequest ", 400)
}

func (o *QueryRawBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewQueryRawUnauthorized creates a QueryRawUnauthorized with default headers values
func NewQueryRawUnauthorized() *QueryRawUnauthorized {
	return &QueryRawUnauthorized{}
}

/*
QueryRawUnauthorized describes a response with status code 401, with default header values.

Authorization Error
*/
type QueryRawUnauthorized struct {
}

// IsSuccess returns true when this query raw unauthorized response has a 2xx status code
func (o *QueryRawUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query raw unauthorized response has a 3xx status code
func (o *QueryRawUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query raw unauthorized response has a 4xx status code
func (o *QueryRawUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this query raw unauthorized response has a 5xx status code
func (o *QueryRawUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this query raw unauthorized response a status code equal to that given
func (o *QueryRawUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *QueryRawUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart/raw][%d] queryRawUnauthorized ", 401)
}

func (o *QueryRawUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/chart/raw][%d] queryRawUnauthorized ", 401)
}

func (o *QueryRawUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewQueryRawNotFound creates a QueryRawNotFound with default headers values
func NewQueryRawNotFound() *QueryRawNotFound {
	return &QueryRawNotFound{}
}

/*
QueryRawNotFound describes a response with status code 404, with default header values.

Metric not found for specified source/host
*/
type QueryRawNotFound struct {
}

// IsSuccess returns true when this query raw not found response has a 2xx status code
func (o *QueryRawNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query raw not found response has a 3xx status code
func (o *QueryRawNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query raw not found response has a 4xx status code
func (o *QueryRawNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this query raw not found response has a 5xx status code
func (o *QueryRawNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this query raw not found response a status code equal to that given
func (o *QueryRawNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *QueryRawNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart/raw][%d] queryRawNotFound ", 404)
}

func (o *QueryRawNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/chart/raw][%d] queryRawNotFound ", 404)
}

func (o *QueryRawNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewQueryRawInternalServerError creates a QueryRawInternalServerError with default headers values
func NewQueryRawInternalServerError() *QueryRawInternalServerError {
	return &QueryRawInternalServerError{}
}

/*
QueryRawInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type QueryRawInternalServerError struct {
}

// IsSuccess returns true when this query raw internal server error response has a 2xx status code
func (o *QueryRawInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query raw internal server error response has a 3xx status code
func (o *QueryRawInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query raw internal server error response has a 4xx status code
func (o *QueryRawInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query raw internal server error response has a 5xx status code
func (o *QueryRawInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query raw internal server error response a status code equal to that given
func (o *QueryRawInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *QueryRawInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/chart/raw][%d] queryRawInternalServerError ", 500)
}

func (o *QueryRawInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/chart/raw][%d] queryRawInternalServerError ", 500)
}

func (o *QueryRawInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
