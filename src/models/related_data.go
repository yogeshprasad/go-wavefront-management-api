// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RelatedData Data object capturing how this event is related to another, requested event.
//
// swagger:model RelatedData
type RelatedData struct {

	// If this event is generated by an alert, the description of that alert.
	// Read Only: true
	AlertDescription string `json:"alertDescription,omitempty"`

	// Chart Link of the anomaly to which this event is related
	// Read Only: true
	AnomalyChartLink string `json:"anomalyChartLink,omitempty"`

	// Set of common dimensions between the 2 events, presented in key=value format
	// Read Only: true
	// Unique: true
	CommonDimensions []string `json:"commonDimensions"`

	// Set of common metrics/labels between the 2 events or anomalies
	// Read Only: true
	// Unique: true
	CommonMetrics []string `json:"commonMetrics"`

	// Set of common sources between the 2 events or anomalies
	// Read Only: true
	// Unique: true
	CommonSources []string `json:"commonSources"`

	// Enhanced score to sort related events and anomalies
	// Read Only: true
	EnhancedScore float64 `json:"enhancedScore,omitempty"`

	// ID of the event to which this event is related
	// Read Only: true
	RelatedID string `json:"relatedId,omitempty"`

	// Text summary of why the two events are related
	// Read Only: true
	Summary string `json:"summary,omitempty"`
}

// Validate validates this related data
func (m *RelatedData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommonDimensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommonMetrics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommonSources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RelatedData) validateCommonDimensions(formats strfmt.Registry) error {
	if swag.IsZero(m.CommonDimensions) { // not required
		return nil
	}

	if err := validate.UniqueItems("commonDimensions", "body", m.CommonDimensions); err != nil {
		return err
	}

	return nil
}

func (m *RelatedData) validateCommonMetrics(formats strfmt.Registry) error {
	if swag.IsZero(m.CommonMetrics) { // not required
		return nil
	}

	if err := validate.UniqueItems("commonMetrics", "body", m.CommonMetrics); err != nil {
		return err
	}

	return nil
}

func (m *RelatedData) validateCommonSources(formats strfmt.Registry) error {
	if swag.IsZero(m.CommonSources) { // not required
		return nil
	}

	if err := validate.UniqueItems("commonSources", "body", m.CommonSources); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this related data based on the context it is used
func (m *RelatedData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAlertDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAnomalyChartLink(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCommonDimensions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCommonMetrics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCommonSources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnhancedScore(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRelatedID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSummary(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RelatedData) contextValidateAlertDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "alertDescription", "body", string(m.AlertDescription)); err != nil {
		return err
	}

	return nil
}

func (m *RelatedData) contextValidateAnomalyChartLink(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "anomalyChartLink", "body", string(m.AnomalyChartLink)); err != nil {
		return err
	}

	return nil
}

func (m *RelatedData) contextValidateCommonDimensions(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "commonDimensions", "body", []string(m.CommonDimensions)); err != nil {
		return err
	}

	return nil
}

func (m *RelatedData) contextValidateCommonMetrics(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "commonMetrics", "body", []string(m.CommonMetrics)); err != nil {
		return err
	}

	return nil
}

func (m *RelatedData) contextValidateCommonSources(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "commonSources", "body", []string(m.CommonSources)); err != nil {
		return err
	}

	return nil
}

func (m *RelatedData) contextValidateEnhancedScore(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "enhancedScore", "body", float64(m.EnhancedScore)); err != nil {
		return err
	}

	return nil
}

func (m *RelatedData) contextValidateRelatedID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "relatedId", "body", string(m.RelatedID)); err != nil {
		return err
	}

	return nil
}

func (m *RelatedData) contextValidateSummary(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "summary", "body", string(m.Summary)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RelatedData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelatedData) UnmarshalBinary(b []byte) error {
	var res RelatedData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
