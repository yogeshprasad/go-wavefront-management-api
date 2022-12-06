// Code generated by go-swagger; DO NOT EDIT.

package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetAllRolesReader is a Reader for the GetAllRoles structure.
type GetAllRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllRolesOK creates a GetAllRolesOK with default headers values
func NewGetAllRolesOK() *GetAllRolesOK {
	return &GetAllRolesOK{}
}

/*
GetAllRolesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAllRolesOK struct {
	Payload *models.ResponseContainerPagedRoleDTO
}

// IsSuccess returns true when this get all roles o k response has a 2xx status code
func (o *GetAllRolesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all roles o k response has a 3xx status code
func (o *GetAllRolesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all roles o k response has a 4xx status code
func (o *GetAllRolesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all roles o k response has a 5xx status code
func (o *GetAllRolesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all roles o k response a status code equal to that given
func (o *GetAllRolesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAllRolesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/role][%d] getAllRolesOK  %+v", 200, o.Payload)
}

func (o *GetAllRolesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/role][%d] getAllRolesOK  %+v", 200, o.Payload)
}

func (o *GetAllRolesOK) GetPayload() *models.ResponseContainerPagedRoleDTO {
	return o.Payload
}

func (o *GetAllRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseContainerPagedRoleDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
