// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetWebhookReader is a Reader for the GetWebhook structure.
type GetWebhookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebhookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWebhookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWebhookOK creates a GetWebhookOK with default headers values
func NewGetWebhookOK() *GetWebhookOK {
	return &GetWebhookOK{}
}

/*
GetWebhookOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWebhookOK struct {
	Payload *models.ResponseContainerNotificant
}

// IsSuccess returns true when this get webhook o k response has a 2xx status code
func (o *GetWebhookOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get webhook o k response has a 3xx status code
func (o *GetWebhookOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webhook o k response has a 4xx status code
func (o *GetWebhookOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get webhook o k response has a 5xx status code
func (o *GetWebhookOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get webhook o k response a status code equal to that given
func (o *GetWebhookOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetWebhookOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/webhook/{id}][%d] getWebhookOK  %+v", 200, o.Payload)
}

func (o *GetWebhookOK) String() string {
	return fmt.Sprintf("[GET /api/v2/webhook/{id}][%d] getWebhookOK  %+v", 200, o.Payload)
}

func (o *GetWebhookOK) GetPayload() *models.ResponseContainerNotificant {
	return o.Payload
}

func (o *GetWebhookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerNotificant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
