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

// Setup A setup definition belonging to a particular integration
//
// swagger:model setup
type Setup struct {

	// Setup description
	// Required: true
	Description *string `json:"description"`

	// Relative file path to the setup.md file
	// Required: true
	FilePath *string `json:"filePath"`

	// Setup title
	Title string `json:"title,omitempty"`

	// Setup Type
	// Required: true
	// Enum: [METRICS LOGS]
	Type *string `json:"type"`
}

// Validate validates this setup
func (m *Setup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Setup) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *Setup) validateFilePath(formats strfmt.Registry) error {

	if err := validate.Required("filePath", "body", m.FilePath); err != nil {
		return err
	}

	return nil
}

var setupTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["METRICS","LOGS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		setupTypeTypePropEnum = append(setupTypeTypePropEnum, v)
	}
}

const (

	// SetupTypeMETRICS captures enum value "METRICS"
	SetupTypeMETRICS string = "METRICS"

	// SetupTypeLOGS captures enum value "LOGS"
	SetupTypeLOGS string = "LOGS"
)

// prop value enum
func (m *Setup) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, setupTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Setup) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this setup based on context it is used
func (m *Setup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Setup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Setup) UnmarshalBinary(b []byte) error {
	var res Setup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
