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

// IntegrationManifestGroup A functional group of integrations defined together in a manifest
//
// swagger:model IntegrationManifestGroup
type IntegrationManifestGroup struct {

	// Materialized JSONs for integrations belonging to this group, as referenced by `integrations`
	IntegrationObjs []*Integration `json:"integrationObjs"`

	// A list of paths to JSON definitions for integrations in this group
	// Required: true
	Integrations []string `json:"integrations"`

	// Subtitle of this integration group
	// Required: true
	Subtitle *string `json:"subtitle"`

	// Title of this integration group
	// Required: true
	Title *string `json:"title"`
}

// Validate validates this integration manifest group
func (m *IntegrationManifestGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIntegrationObjs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntegrations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubtitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntegrationManifestGroup) validateIntegrationObjs(formats strfmt.Registry) error {
	if swag.IsZero(m.IntegrationObjs) { // not required
		return nil
	}

	for i := 0; i < len(m.IntegrationObjs); i++ {
		if swag.IsZero(m.IntegrationObjs[i]) { // not required
			continue
		}

		if m.IntegrationObjs[i] != nil {
			if err := m.IntegrationObjs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("integrationObjs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("integrationObjs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IntegrationManifestGroup) validateIntegrations(formats strfmt.Registry) error {

	if err := validate.Required("integrations", "body", m.Integrations); err != nil {
		return err
	}

	return nil
}

func (m *IntegrationManifestGroup) validateSubtitle(formats strfmt.Registry) error {

	if err := validate.Required("subtitle", "body", m.Subtitle); err != nil {
		return err
	}

	return nil
}

func (m *IntegrationManifestGroup) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this integration manifest group based on the context it is used
func (m *IntegrationManifestGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIntegrationObjs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntegrationManifestGroup) contextValidateIntegrationObjs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IntegrationObjs); i++ {

		if m.IntegrationObjs[i] != nil {
			if err := m.IntegrationObjs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("integrationObjs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("integrationObjs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IntegrationManifestGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegrationManifestGroup) UnmarshalBinary(b []byte) error {
	var res IntegrationManifestGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
