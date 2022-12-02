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

// IntegrationAlert A alert definition belonging to a particular integration
//
// swagger:model IntegrationAlert
type IntegrationAlert struct {

	// alert min obj
	AlertMinObj *AlertMin `json:"alertMinObj,omitempty"`

	// alert obj
	AlertObj *Alert `json:"alertObj,omitempty"`

	// Alert description
	// Required: true
	Description *string `json:"description"`

	// Alert name
	// Required: true
	Name *string `json:"name"`

	// URL path to the JSON definition of this alert
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this integration alert
func (m *IntegrationAlert) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlertMinObj(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAlertObj(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntegrationAlert) validateAlertMinObj(formats strfmt.Registry) error {
	if swag.IsZero(m.AlertMinObj) { // not required
		return nil
	}

	if m.AlertMinObj != nil {
		if err := m.AlertMinObj.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alertMinObj")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("alertMinObj")
			}
			return err
		}
	}

	return nil
}

func (m *IntegrationAlert) validateAlertObj(formats strfmt.Registry) error {
	if swag.IsZero(m.AlertObj) { // not required
		return nil
	}

	if m.AlertObj != nil {
		if err := m.AlertObj.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alertObj")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("alertObj")
			}
			return err
		}
	}

	return nil
}

func (m *IntegrationAlert) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *IntegrationAlert) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IntegrationAlert) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this integration alert based on the context it is used
func (m *IntegrationAlert) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAlertMinObj(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAlertObj(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntegrationAlert) contextValidateAlertMinObj(ctx context.Context, formats strfmt.Registry) error {

	if m.AlertMinObj != nil {
		if err := m.AlertMinObj.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alertMinObj")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("alertMinObj")
			}
			return err
		}
	}

	return nil
}

func (m *IntegrationAlert) contextValidateAlertObj(ctx context.Context, formats strfmt.Registry) error {

	if m.AlertObj != nil {
		if err := m.AlertObj.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alertObj")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("alertObj")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IntegrationAlert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegrationAlert) UnmarshalBinary(b []byte) error {
	var res IntegrationAlert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
