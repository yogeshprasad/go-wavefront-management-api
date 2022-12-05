// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ResponseContainerPagedMonitoredServiceDTO JSON container for the HTTP response along with status
//
// swagger:model ResponseContainerPagedMonitoredServiceDTO
type ResponseContainerPagedMonitoredServiceDTO struct {

	// response
	Response *PagedMonitoredServiceDTO `json:"response,omitempty"`

	// status
	// Required: true
	Status *ResponseStatus `json:"status"`
}

// Validate validates this response container paged monitored service d t o
func (m *ResponseContainerPagedMonitoredServiceDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponseContainerPagedMonitoredServiceDTO) validateResponse(formats strfmt.Registry) error {
	if swag.IsZero(m.Response) { // not required
		return nil
	}

	if m.Response != nil {
		if err := m.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("response")
			}
			return err
		}
	}

	return nil
}

func (m *ResponseContainerPagedMonitoredServiceDTO) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this response container paged monitored service d t o based on the context it is used
func (m *ResponseContainerPagedMonitoredServiceDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResponse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponseContainerPagedMonitoredServiceDTO) contextValidateResponse(ctx context.Context, formats strfmt.Registry) error {

	if m.Response != nil {
		if err := m.Response.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("response")
			}
			return err
		}
	}

	return nil
}

func (m *ResponseContainerPagedMonitoredServiceDTO) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResponseContainerPagedMonitoredServiceDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseContainerPagedMonitoredServiceDTO) UnmarshalBinary(b []byte) error {
	var res ResponseContainerPagedMonitoredServiceDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
