// Code generated by go-swagger; DO NOT EDIT.

package account_user_and_service_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// InviteUserAccountsReader is a Reader for the InviteUserAccounts structure.
type InviteUserAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InviteUserAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInviteUserAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInviteUserAccountsOK creates a InviteUserAccountsOK with default headers values
func NewInviteUserAccountsOK() *InviteUserAccountsOK {
	return &InviteUserAccountsOK{}
}

/*
InviteUserAccountsOK describes a response with status code 200, with default header values.

successful operation
*/
type InviteUserAccountsOK struct {
	Payload *models.UserModel
}

// IsSuccess returns true when this invite user accounts o k response has a 2xx status code
func (o *InviteUserAccountsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this invite user accounts o k response has a 3xx status code
func (o *InviteUserAccountsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invite user accounts o k response has a 4xx status code
func (o *InviteUserAccountsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this invite user accounts o k response has a 5xx status code
func (o *InviteUserAccountsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this invite user accounts o k response a status code equal to that given
func (o *InviteUserAccountsOK) IsCode(code int) bool {
	return code == 200
}

func (o *InviteUserAccountsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/account/user/invite][%d] inviteUserAccountsOK  %+v", 200, o.Payload)
}

func (o *InviteUserAccountsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/account/user/invite][%d] inviteUserAccountsOK  %+v", 200, o.Payload)
}

func (o *InviteUserAccountsOK) GetPayload() *models.UserModel {
	return o.Payload
}

func (o *InviteUserAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
