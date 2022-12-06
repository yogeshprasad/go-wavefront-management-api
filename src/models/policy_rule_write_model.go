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

// PolicyRuleWriteModel Policy rule entity for write requests.
//
// swagger:model PolicyRuleWriteModel
type PolicyRuleWriteModel struct {

	// The access type of the policy rule
	// Enum: [ALLOW BLOCK]
	AccessType string `json:"accessType,omitempty"`

	// The list of account identifiers of the policy rule
	Accounts []string `json:"accounts"`

	// The description of the policy rule
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// The name of the policy rule
	// Required: true
	Name *string `json:"name"`

	// The prefixes of the policy rule
	Prefixes []string `json:"prefixes"`

	// The list of role identifiers of the policy rule
	Roles []string `json:"roles"`

	// The tag/value pairs of the policy rule
	Tags []*Annotation `json:"tags"`

	// Whether tags should be AND-ed or OR-ed.If true, a metric must contain all tags in order for the policy rule to apply. If false, the tags are OR-ed, and a metric must contain one of the tags. Default: false
	TagsAnded bool `json:"tagsAnded,omitempty"`

	// The list of user group identifiers of the policy rule
	UserGroups []string `json:"userGroups"`
}

// Validate validates this policy rule write model
func (m *PolicyRuleWriteModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var policyRuleWriteModelTypeAccessTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ALLOW","BLOCK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		policyRuleWriteModelTypeAccessTypePropEnum = append(policyRuleWriteModelTypeAccessTypePropEnum, v)
	}
}

const (

	// PolicyRuleWriteModelAccessTypeALLOW captures enum value "ALLOW"
	PolicyRuleWriteModelAccessTypeALLOW string = "ALLOW"

	// PolicyRuleWriteModelAccessTypeBLOCK captures enum value "BLOCK"
	PolicyRuleWriteModelAccessTypeBLOCK string = "BLOCK"
)

// prop value enum
func (m *PolicyRuleWriteModel) validateAccessTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, policyRuleWriteModelTypeAccessTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PolicyRuleWriteModel) validateAccessType(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccessTypeEnum("accessType", "body", m.AccessType); err != nil {
		return err
	}

	return nil
}

func (m *PolicyRuleWriteModel) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PolicyRuleWriteModel) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this policy rule write model based on the context it is used
func (m *PolicyRuleWriteModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyRuleWriteModel) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PolicyRuleWriteModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyRuleWriteModel) UnmarshalBinary(b []byte) error {
	var res PolicyRuleWriteModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
