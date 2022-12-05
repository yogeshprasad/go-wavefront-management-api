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

// CloudWatchConfiguration Configuration specific to the CloudWatch AWS integration.  Only applicable when the containing Credential has service=CLOUDWATCH
//
// swagger:model CloudWatchConfiguration
type CloudWatchConfiguration struct {

	// base credentials
	BaseCredentials *AWSBaseCredentials `json:"baseCredentials,omitempty"`

	// A string->string map of allow list of AWS instance tag-value pairs (in AWS).  If the instance's AWS tags match this allow list, CloudWatch data about this instance is ingested.  Multiple entries are OR'ed
	InstanceSelectionTags map[string]string `json:"instanceSelectionTags,omitempty"`

	// A string expressing the allow list of AWS instance tag-value pairs.  If the instance's AWS tags match this allow list, CloudWatch data about this instance is ingested.  Multiple entries are OR'ed and also OR'ed with entries from instanceSelectionTags.  Key-value pairs in the string are separated by commas and in the form k=v.  Example: "k1=v1, k1=v2, k3=v3".
	InstanceSelectionTagsExpr string `json:"instanceSelectionTagsExpr,omitempty"`

	// A regular expression that a CloudWatch metric name must match (case-insensitively) in order to be ingested
	// Example: ^aws.(billing|instance|sqs|sns|reservedInstance|ebs|route53.health|ec2.status|elb).*$
	MetricFilterRegex string `json:"metricFilterRegex,omitempty"`

	// A list of namespace that limit what we query from CloudWatch.
	Namespaces []string `json:"namespaces"`

	// A regular expression that AWS tag key name must match (case-insensitively) in order to be ingested
	// Example: (region|name)
	PointTagFilterRegex string `json:"pointTagFilterRegex,omitempty"`

	// A regular expression that a AWS S3 Bucket name must match (case-insensitively) in order to be ingested
	// Example: ^(test|demo).*$
	S3BucketNameFilterRegex string `json:"s3BucketNameFilterRegex,omitempty"`

	// ThreadDistributionInMins
	ThreadDistributionInMins int32 `json:"threadDistributionInMins,omitempty"`

	// A string->string map of allow list of AWS volume tag-value pairs (in AWS).  If the volume's AWS tags match this allow list, CloudWatch data about this volume is ingested.  Multiple entries are OR'ed
	VolumeSelectionTags map[string]string `json:"volumeSelectionTags,omitempty"`

	// A string expressing the allow list of AWS volume tag-value pairs.  If the volume's AWS tags match this allow list, CloudWatch data about this volume is ingested.  Multiple entries are OR'ed and also OR'ed with entries from volumeSelectionTags.  Key-value pairs in the string are separated by commas and in the form k=v.  Example: "k1=v1, k1=v2, k3=v3".
	VolumeSelectionTagsExpr string `json:"volumeSelectionTagsExpr,omitempty"`
}

// Validate validates this cloud watch configuration
func (m *CloudWatchConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudWatchConfiguration) validateBaseCredentials(formats strfmt.Registry) error {
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

// ContextValidate validate this cloud watch configuration based on the context it is used
func (m *CloudWatchConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBaseCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudWatchConfiguration) contextValidateBaseCredentials(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CloudWatchConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudWatchConfiguration) UnmarshalBinary(b []byte) error {
	var res CloudWatchConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
