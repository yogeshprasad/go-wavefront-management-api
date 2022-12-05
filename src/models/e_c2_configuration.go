// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EC2Configuration Configurations specific to the EC2 AWS integration.  Only applicable when the containing Credential has service=EC2
//
// swagger:model EC2Configuration
type EC2Configuration struct {

	// base credentials
	BaseCredentials *AWSBaseCredentials `json:"baseCredentials,omitempty"`

	// A list of AWS instance tags that, when found, will be used as the "source" name in a series.  Default: ["hostname", "host", "name"].  If no tag in this list is found, the series source is set to the instance id.
	HostNameTags []string `json:"hostNameTags"`

	// A string expressing the allow list of AWS instance tag-value pairs.  If the instance's AWS tags match this allow list, data about this instance is ingested from EC2 APIs  Multiple entries are OR'ed.  Key-value pairs in the string are separated by commas and in the form k=v.  Example: "k1=v1, k1=v2, k3=v3".
	InstanceSelectionTagsExpr string `json:"instanceSelectionTagsExpr,omitempty"`

	// A regular expression that a metric name must match (case-insensitively) in order to be ingested
	// Example: ^aws.(sqs|ebs).*$
	MetricFilterRegex string `json:"metricFilterRegex,omitempty"`

	// A regular expression that AWS tag key name must match (case-insensitively) in order to be ingested
	// Example: (region|name)
	PointTagFilterRegex string `json:"pointTagFilterRegex,omitempty"`

	// A string expressing the allow list of AWS volume tag-value pairs.  If the volume's AWS tags match this allow list, Data about this volume is ingested from EBS APIs.  Multiple entries are OR'ed.  Key-value pairs in the string are separated by commas and in the form k=v.  Example: "k1=v1, k1=v2, k3=v3".
	VolumeSelectionTagsExpr string `json:"volumeSelectionTagsExpr,omitempty"`
}

// Validate validates this e c2 configuration
func (m *EC2Configuration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EC2Configuration) validateBaseCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.BaseCredentials) { // not required
		return nil
	}

	if m.BaseCredentials != nil {
		if err := m.BaseCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baseCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("baseCredentials")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this e c2 configuration based on the context it is used
func (m *EC2Configuration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBaseCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EC2Configuration) contextValidateBaseCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.BaseCredentials != nil {
		if err := m.BaseCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("baseCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("baseCredentials")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EC2Configuration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EC2Configuration) UnmarshalBinary(b []byte) error {
	var res EC2Configuration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
