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

// BatchUpdateReader is a Reader for the BatchUpdate structure.
type BatchUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BatchUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBatchUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewBatchUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewBatchUpdateMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBatchUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBatchUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBatchUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBatchUpdateOK creates a BatchUpdateOK with default headers values
func NewBatchUpdateOK() *BatchUpdateOK {
	return &BatchUpdateOK{}
}

/*
BatchUpdateOK describes a response with status code 200, with default header values.

successful operation
*/
type BatchUpdateOK struct {
	Payload *models.ResponseContainer
}

// IsSuccess returns true when this batch update o k response has a 2xx status code
func (o *BatchUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this batch update o k response has a 3xx status code
func (o *BatchUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this batch update o k response has a 4xx status code
func (o *BatchUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this batch update o k response has a 5xx status code
func (o *BatchUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this batch update o k response a status code equal to that given
func (o *BatchUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *BatchUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/monitoredservice/services][%d] batchUpdateOK  %+v", 200, o.Payload)
}

func (o *BatchUpdateOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/monitoredservice/services][%d] batchUpdateOK  %+v", 200, o.Payload)
}

func (o *BatchUpdateOK) GetPayload() *models.ResponseContainer {
	return o.Payload
}

func (o *BatchUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBatchUpdateNoContent creates a BatchUpdateNoContent with default headers values
func NewBatchUpdateNoContent() *BatchUpdateNoContent {
	return &BatchUpdateNoContent{}
}

/*
BatchUpdateNoContent describes a response with status code 204, with default header values.

Successful Update. No content.
*/
type BatchUpdateNoContent struct {
}

// IsSuccess returns true when this batch update no content response has a 2xx status code
func (o *BatchUpdateNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this batch update no content response has a 3xx status code
func (o *BatchUpdateNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this batch update no content response has a 4xx status code
func (o *BatchUpdateNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this batch update no content response has a 5xx status code
func (o *BatchUpdateNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this batch update no content response a status code equal to that given
func (o *BatchUpdateNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *BatchUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v2/monitoredservice/services][%d] batchUpdateNoContent ", 204)
}

func (o *BatchUpdateNoContent) String() string {
	return fmt.Sprintf("[PUT /api/v2/monitoredservice/services][%d] batchUpdateNoContent ", 204)
}

func (o *BatchUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBatchUpdateMultiStatus creates a BatchUpdateMultiStatus with default headers values
func NewBatchUpdateMultiStatus() *BatchUpdateMultiStatus {
	return &BatchUpdateMultiStatus{}
}

/*
BatchUpdateMultiStatus describes a response with status code 207, with default header values.

Bad request payload.  Partial payload updated. One or more errors
*/
type BatchUpdateMultiStatus struct {
}

// IsSuccess returns true when this batch update multi status response has a 2xx status code
func (o *BatchUpdateMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this batch update multi status response has a 3xx status code
func (o *BatchUpdateMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this batch update multi status response has a 4xx status code
func (o *BatchUpdateMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this batch update multi status response has a 5xx status code
func (o *BatchUpdateMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this batch update multi status response a status code equal to that given
func (o *BatchUpdateMultiStatus) IsCode(code int) bool {
	return code == 207
}

func (o *BatchUpdateMultiStatus) Error() string {
	return fmt.Sprintf("[PUT /api/v2/monitoredservice/services][%d] batchUpdateMultiStatus ", 207)
}

func (o *BatchUpdateMultiStatus) String() string {
	return fmt.Sprintf("[PUT /api/v2/monitoredservice/services][%d] batchUpdateMultiStatus ", 207)
}

func (o *BatchUpdateMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBatchUpdateBadRequest creates a BatchUpdateBadRequest with default headers values
func NewBatchUpdateBadRequest() *BatchUpdateBadRequest {
	return &BatchUpdateBadRequest{}
}

/*
BatchUpdateBadRequest describes a response with status code 400, with default header values.

Bad request payload. One or more entities were not found for update
*/
type BatchUpdateBadRequest struct {
}

// IsSuccess returns true when this batch update bad request response has a 2xx status code
func (o *BatchUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this batch update bad request response has a 3xx status code
func (o *BatchUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this batch update bad request response has a 4xx status code
func (o *BatchUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this batch update bad request response has a 5xx status code
func (o *BatchUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this batch update bad request response a status code equal to that given
func (o *BatchUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *BatchUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/monitoredservice/services][%d] batchUpdateBadRequest ", 400)
}

func (o *BatchUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/monitoredservice/services][%d] batchUpdateBadRequest ", 400)
}

func (o *BatchUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBatchUpdateForbidden creates a BatchUpdateForbidden with default headers values
func NewBatchUpdateForbidden() *BatchUpdateForbidden {
	return &BatchUpdateForbidden{}
}

/*
BatchUpdateForbidden describes a response with status code 403, with default header values.

Don't have permission for this operation. Check token.
*/
type BatchUpdateForbidden struct {
}

// IsSuccess returns true when this batch update forbidden response has a 2xx status code
func (o *BatchUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this batch update forbidden response has a 3xx status code
func (o *BatchUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this batch update forbidden response has a 4xx status code
func (o *BatchUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this batch update forbidden response has a 5xx status code
func (o *BatchUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this batch update forbidden response a status code equal to that given
func (o *BatchUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *BatchUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/monitoredservice/services][%d] batchUpdateForbidden ", 403)
}

func (o *BatchUpdateForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/monitoredservice/services][%d] batchUpdateForbidden ", 403)
}

func (o *BatchUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBatchUpdateInternalServerError creates a BatchUpdateInternalServerError with default headers values
func NewBatchUpdateInternalServerError() *BatchUpdateInternalServerError {
	return &BatchUpdateInternalServerError{}
}

/*
BatchUpdateInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type BatchUpdateInternalServerError struct {
}

// IsSuccess returns true when this batch update internal server error response has a 2xx status code
func (o *BatchUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this batch update internal server error response has a 3xx status code
func (o *BatchUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this batch update internal server error response has a 4xx status code
func (o *BatchUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this batch update internal server error response has a 5xx status code
func (o *BatchUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this batch update internal server error response a status code equal to that given
func (o *BatchUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *BatchUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/monitoredservice/services][%d] batchUpdateInternalServerError ", 500)
}

func (o *BatchUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/monitoredservice/services][%d] batchUpdateInternalServerError ", 500)
}

func (o *BatchUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
