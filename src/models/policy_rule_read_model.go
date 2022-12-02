// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PolicyRuleReadModel Policy rule entity for read requests.
//
// swagger:model PolicyRuleReadModel
type PolicyRuleReadModel struct {

	// The access type of the policy rule
	// Enum: [ALLOW BLOCK]
	AccessType string `json:"accessType,omitempty"`

	// The list of accounts of the policy rule
	Accounts []*AccessControlElement `json:"accounts"`

	// The description of the policy rule
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// The name of the policy rule
	// Required: true
	Name *string `json:"name"`

	// The prefixes of the policy rule
	Prefixes []string `json:"prefixes"`

	// The list of roles of the policy rule
	Roles []*AccessControlElement `json:"roles"`

	// The tag/value pairs of the policy rule
	Tags []Annotation `json:"tags"`

	// Whether tags should be AND-ed or OR-ed.If true, a metric must contain all tags in order for the policy rule to apply. If false, the tags are OR-ed, and a metric must contain one of the tags. Default: false
	TagsAnded bool `json:"tagsAnded,omitempty"`

	// The list of user groups of the policy rule
	UserGroups []*AccessControlElement `json:"userGroups"`
}

// Validate validates this policy rule read model
func (m *PolicyRuleReadModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

var policyRuleReadModelTypeAccessTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ALLOW","BLOCK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		policyRuleReadModelTypeAccessTypePropEnum = append(policyRuleReadModelTypeAccessTypePropEnum, v)
	}
}

const (

	// PolicyRuleReadModelAccessTypeALLOW captures enum value "ALLOW"
	PolicyRuleReadModelAccessTypeALLOW string = "ALLOW"

	// PolicyRuleReadModelAccessTypeBLOCK captures enum value "BLOCK"
	PolicyRuleReadModelAccessTypeBLOCK string = "BLOCK"
)

// prop value enum
func (m *PolicyRuleReadModel) validateAccessTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, policyRuleReadModelTypeAccessTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PolicyRuleReadModel) validateAccessType(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccessTypeEnum("accessType", "body", m.AccessType); err != nil {
		return err
	}

	return nil
}

func (m *PolicyRuleReadModel) validateAccounts(formats strfmt.Registry) error {
	if swag.IsZero(m.Accounts) { // not required
		return nil
	}

	for i := 0; i < len(m.Accounts); i++ {
		if swag.IsZero(m.Accounts[i]) { // not required
			continue
		}

		if m.Accounts[i] != nil {
			if err := m.Accounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("accounts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("accounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PolicyRuleReadModel) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PolicyRuleReadModel) validateRoles(formats strfmt.Registry) error {
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

func (m *PolicyRuleReadModel) validateUserGroups(formats strfmt.Registry) error {
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

// ContextValidate validate this policy rule read model based on the context it is used
func (m *PolicyRuleReadModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccounts(ctx, formats); err != nil {
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

func (m *PolicyRuleReadModel) contextValidateAccounts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Accounts); i++ {

		if m.Accounts[i] != nil {
			if err := m.Accounts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("accounts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("accounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PolicyRuleReadModel) contextValidateRoles(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PolicyRuleReadModel) contextValidateUserGroups(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PolicyRuleReadModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyRuleReadModel) UnmarshalBinary(b []byte) error {
	var res PolicyRuleReadModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
