// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DashboardSection dashboard section
//
// swagger:model DashboardSection
type DashboardSection struct {

	// Id of this section
	ID string `json:"id,omitempty"`

	// Name of this section
	// Required: true
	Name *string `json:"name"`

	// Rows of this section
	// Required: true
	Rows []*DashboardSectionRow `json:"rows"`

	// Display filter for conditional dashboard section
	SectionFilter JSONNode `json:"sectionFilter,omitempty"`
}

// Validate validates this dashboard section
func (m *DashboardSection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRows(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DashboardSection) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DashboardSection) validateRows(formats strfmt.Registry) error {

	if err := validate.Required("rows", "body", m.Rows); err != nil {
		return err
	}

	for i := 0; i < len(m.Rows); i++ {
		if swag.IsZero(m.Rows[i]) { // not required
			continue
		}

		if m.Rows[i] != nil {
			if err := m.Rows[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rows" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dashboard section based on the context it is used
func (m *DashboardSection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRows(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DashboardSection) contextValidateRows(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Rows); i++ {

		if m.Rows[i] != nil {
			if err := m.Rows[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rows" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DashboardSection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DashboardSection) UnmarshalBinary(b []byte) error {
	var res DashboardSection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
