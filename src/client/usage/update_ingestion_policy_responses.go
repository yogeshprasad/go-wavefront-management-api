// Code generated by go-swagger; DO NOT EDIT.

package usage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// UpdateIngestionPolicyReader is a Reader for the UpdateIngestionPolicy structure.
type UpdateIngestionPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateIngestionPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateIngestionPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateIngestionPolicyOK creates a UpdateIngestionPolicyOK with default headers values
func NewUpdateIngestionPolicyOK() *UpdateIngestionPolicyOK {
	return &UpdateIngestionPolicyOK{}
}

/*
UpdateIngestionPolicyOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateIngestionPolicyOK struct {
	Payload *models.ResponseContainerIngestionPolicy
}

// IsSuccess returns true when this update ingestion policy o k response has a 2xx status code
func (o *UpdateIngestionPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update ingestion policy o k response has a 3xx status code
func (o *UpdateIngestionPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update ingestion policy o k response has a 4xx status code
func (o *UpdateIngestionPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update ingestion policy o k response has a 5xx status code
func (o *UpdateIngestionPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update ingestion policy o k response a status code equal to that given
func (o *UpdateIngestionPolicyOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateIngestionPolicyOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/usage/ingestionpolicy/{id}][%d] updateIngestionPolicyOK  %+v", 200, o.Payload)
}

func (o *UpdateIngestionPolicyOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/usage/ingestionpolicy/{id}][%d] updateIngestionPolicyOK  %+v", 200, o.Payload)
}

func (o *UpdateIngestionPolicyOK) GetPayload() *models.ResponseContainerIngestionPolicy {
	return o.Payload
}

func (o *UpdateIngestionPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerIngestionPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
