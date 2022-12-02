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

// AccessPolicyRuleDTO access policy rule d t o
//
// swagger:model AccessPolicyRuleDTO
type AccessPolicyRuleDTO struct {

	// action
	// Enum: [ALLOW DENY]
	Action string `json:"action,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// subnet
	Subnet string `json:"subnet,omitempty"`
}

// Validate validates this access policy rule d t o
func (m *AccessPolicyRuleDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var accessPolicyRuleDTOTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ALLOW","DENY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accessPolicyRuleDTOTypeActionPropEnum = append(accessPolicyRuleDTOTypeActionPropEnum, v)
	}
}

const (

	// AccessPolicyRuleDTOActionALLOW captures enum value "ALLOW"
	AccessPolicyRuleDTOActionALLOW string = "ALLOW"

	// AccessPolicyRuleDTOActionDENY captures enum value "DENY"
	AccessPolicyRuleDTOActionDENY string = "DENY"
)

// prop value enum
func (m *AccessPolicyRuleDTO) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accessPolicyRuleDTOTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AccessPolicyRuleDTO) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this access policy rule d t o based on context it is used
func (m *AccessPolicyRuleDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccessPolicyRuleDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccessPolicyRuleDTO) UnmarshalBinary(b []byte) error {
	var res AccessPolicyRuleDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
