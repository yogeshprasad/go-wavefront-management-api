// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TargetInfo Alert target display information that includes type, id, and the name of the alert target.
//
// swagger:model TargetInfo
type TargetInfo struct {

	// ID of the alert target
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Notification method of the alert target
	// Read Only: true
	// Enum: [EMAIL PAGERDUTY WEBHOOK]
	Method string `json:"method,omitempty"`

	// Name of the alert target
	// Read Only: true
	Name string `json:"name,omitempty"`
}

// Validate validates this target info
func (m *TargetInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var targetInfoTypeMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EMAIL","PAGERDUTY","WEBHOOK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		targetInfoTypeMethodPropEnum = append(targetInfoTypeMethodPropEnum, v)
	}
}

const (

	// TargetInfoMethodEMAIL captures enum value "EMAIL"
	TargetInfoMethodEMAIL string = "EMAIL"

	// TargetInfoMethodPAGERDUTY captures enum value "PAGERDUTY"
	TargetInfoMethodPAGERDUTY string = "PAGERDUTY"

	// TargetInfoMethodWEBHOOK captures enum value "WEBHOOK"
	TargetInfoMethodWEBHOOK string = "WEBHOOK"
)

// prop value enum
func (m *TargetInfo) validateMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, targetInfoTypeMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TargetInfo) validateMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.Method) { // not required
		return nil
	}

	// value enum
	if err := m.validateMethodEnum("method", "body", m.Method); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this target info based on the context it is used
func (m *TargetInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMethod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TargetInfo) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *TargetInfo) contextValidateMethod(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "method", "body", string(m.Method)); err != nil {
		return err
	}

	return nil
}

func (m *TargetInfo) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TargetInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TargetInfo) UnmarshalBinary(b []byte) error {
	var res TargetInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
