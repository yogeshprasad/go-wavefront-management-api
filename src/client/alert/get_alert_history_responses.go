// Code generated by go-swagger; DO NOT EDIT.

package alert

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetAlertHistoryReader is a Reader for the GetAlertHistory structure.
type GetAlertHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAlertHistoryOK creates a GetAlertHistoryOK with default headers values
func NewGetAlertHistoryOK() *GetAlertHistoryOK {
	return &GetAlertHistoryOK{}
}

/*
GetAlertHistoryOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAlertHistoryOK struct {
	Payload *models.ResponseContainerHistoryResponse
}

// IsSuccess returns true when this get alert history o k response has a 2xx status code
func (o *GetAlertHistoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get alert history o k response has a 3xx status code
func (o *GetAlertHistoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert history o k response has a 4xx status code
func (o *GetAlertHistoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alert history o k response has a 5xx status code
func (o *GetAlertHistoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert history o k response a status code equal to that given
func (o *GetAlertHistoryOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAlertHistoryOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/alert/{id}/history][%d] getAlertHistoryOK  %+v", 200, o.Payload)
}

func (o *GetAlertHistoryOK) String() string {
	return fmt.Sprintf("[GET /api/v2/alert/{id}/history][%d] getAlertHistoryOK  %+v", 200, o.Payload)
}

func (o *GetAlertHistoryOK) GetPayload() *models.ResponseContainerHistoryResponse {
	return o.Payload
}

func (o *GetAlertHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerHistoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
