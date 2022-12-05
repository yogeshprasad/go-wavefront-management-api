// Code generated by go-swagger; DO NOT EDIT.

package api_token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetAllTokensReader is a Reader for the GetAllTokens structure.
type GetAllTokensReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllTokensReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllTokensOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllTokensOK creates a GetAllTokensOK with default headers values
func NewGetAllTokensOK() *GetAllTokensOK {
	return &GetAllTokensOK{}
}

/*
GetAllTokensOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAllTokensOK struct {
	Payload *models.ResponseContainerListUserAPIToken
}

// IsSuccess returns true when this get all tokens o k response has a 2xx status code
func (o *GetAllTokensOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all tokens o k response has a 3xx status code
func (o *GetAllTokensOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all tokens o k response has a 4xx status code
func (o *GetAllTokensOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all tokens o k response has a 5xx status code
func (o *GetAllTokensOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all tokens o k response a status code equal to that given
func (o *GetAllTokensOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAllTokensOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/apitoken][%d] getAllTokensOK  %+v", 200, o.Payload)
}

func (o *GetAllTokensOK) String() string {
	return fmt.Sprintf("[GET /api/v2/apitoken][%d] getAllTokensOK  %+v", 200, o.Payload)
}

func (o *GetAllTokensOK) GetPayload() *models.ResponseContainerListUserAPIToken {
	return o.Payload
}

func (o *GetAllTokensOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerListUserAPIToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
