// Code generated by go-swagger; DO NOT EDIT.

package notificant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetAllNotificantsReader is a Reader for the GetAllNotificants structure.
type GetAllNotificantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllNotificantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllNotificantsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllNotificantsOK creates a GetAllNotificantsOK with default headers values
func NewGetAllNotificantsOK() *GetAllNotificantsOK {
	return &GetAllNotificantsOK{}
}

/*
GetAllNotificantsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAllNotificantsOK struct {
	Payload *models.ResponseContainerPagedNotificant
}

// IsSuccess returns true when this get all notificants o k response has a 2xx status code
func (o *GetAllNotificantsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all notificants o k response has a 3xx status code
func (o *GetAllNotificantsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all notificants o k response has a 4xx status code
func (o *GetAllNotificantsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all notificants o k response has a 5xx status code
func (o *GetAllNotificantsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all notificants o k response a status code equal to that given
func (o *GetAllNotificantsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAllNotificantsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/notificant][%d] getAllNotificantsOK  %+v", 200, o.Payload)
}

func (o *GetAllNotificantsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/notificant][%d] getAllNotificantsOK  %+v", 200, o.Payload)
}

func (o *GetAllNotificantsOK) GetPayload() *models.ResponseContainerPagedNotificant {
	return o.Payload
}

func (o *GetAllNotificantsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerPagedNotificant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
