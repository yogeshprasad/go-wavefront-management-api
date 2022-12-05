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
)

// UserDTO user d t o
//
// swagger:model UserDTO
type UserDTO struct {

	// customer
	Customer string `json:"customer,omitempty"`

	// groups
	Groups []string `json:"groups"`

	// identifier
	Identifier string `json:"identifier,omitempty"`

	// ingestion policies
	IngestionPolicies []*IngestionPolicy `json:"ingestionPolicies"`

	// ingestion policy
	IngestionPolicy *IngestionPolicy `json:"ingestionPolicy,omitempty"`

	// last successful login
	LastSuccessfulLogin int64 `json:"lastSuccessfulLogin,omitempty"`

	// roles
	Roles []*RoleDTO `json:"roles"`

	// sso Id
	SsoID string `json:"ssoId,omitempty"`

	// user groups
	UserGroups []*UserGroup `json:"userGroups"`
}

// Validate validates this user d t o
func (m *UserDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIngestionPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngestionPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserDTO) validateIngestionPolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.IngestionPolicies) { // not required
		return nil
	}

	for i := 0; i < len(m.IngestionPolicies); i++ {
		if swag.IsZero(m.IngestionPolicies[i]) { // not required
			continue
		}

		if m.IngestionPolicies[i] != nil {
			if err := m.IngestionPolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ingestionPolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ingestionPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserDTO) validateIngestionPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.IngestionPolicy) { // not required
		return nil
	}

	if m.IngestionPolicy != nil {
		if err := m.IngestionPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ingestionPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ingestionPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *UserDTO) validateRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.Roles) { // not required
		return nil
	}

	for i := 0; i < len(m.Roles); i++ {
		if swag.IsZero(m.Roles[i]) { // not required
			continue
		}

		if m.Roles[i] != nil {
			if err := m.Roles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserDTO) validateUserGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.UserGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.UserGroups); i++ {
		if swag.IsZero(m.UserGroups[i]) { // not required
			continue
		}

		if m.UserGroups[i] != nil {
			if err := m.UserGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("userGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this user d t o based on the context it is used
func (m *UserDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIngestionPolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIngestionPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserDTO) contextValidateIngestionPolicies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IngestionPolicies); i++ {

		if m.IngestionPolicies[i] != nil {
			if err := m.IngestionPolicies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ingestionPolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ingestionPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserDTO) contextValidateIngestionPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.IngestionPolicy != nil {
		if err := m.IngestionPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ingestionPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ingestionPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *UserDTO) contextValidateRoles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Roles); i++ {

		if m.Roles[i] != nil {
			if err := m.Roles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserDTO) contextValidateUserGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UserGroups); i++ {

		if m.UserGroups[i] != nil {
			if err := m.UserGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("userGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserDTO) UnmarshalBinary(b []byte) error {
	var res UserDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
