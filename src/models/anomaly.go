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

// Anomaly DTO for wavefront anomaly
//
// swagger:model Anomaly
type Anomaly struct {

	// chart hash(as unique identifier) for this anomaly
	// Required: true
	ChartHash *string `json:"chartHash"`

	// chart link for this anomaly
	ChartLink string `json:"chartLink,omitempty"`

	// column number for this anomaly
	// Required: true
	Col *int32 `json:"col"`

	// created epoch millis
	CreatedEpochMillis int64 `json:"createdEpochMillis,omitempty"`

	// creator Id
	CreatorID string `json:"creatorId,omitempty"`

	// id of the customer to which this anomaly belongs
	// Read Only: true
	Customer string `json:"customer,omitempty"`

	// dashboard id for this anomaly
	// Required: true
	DashboardID *string `json:"dashboardId"`

	// deleted
	Deleted bool `json:"deleted,omitempty"`

	// endMs for this anomaly
	// Required: true
	EndMs *int64 `json:"endMs"`

	// list of hosts affected of this anomaly
	HostsUsed []string `json:"hostsUsed"`

	// unique id that defines this anomaly
	// Read Only: true
	ID string `json:"id,omitempty"`

	// image link for this anomaly
	ImageLink string `json:"imageLink,omitempty"`

	// list of metrics used of this anomaly
	MetricsUsed []string `json:"metricsUsed"`

	// model for this anomaly
	Model string `json:"model,omitempty"`

	// list of originalStripe belongs to this anomaly
	OriginalStripes []*Stripe `json:"originalStripes"`

	// param hash for this anomaly
	// Required: true
	ParamHash *string `json:"paramHash"`

	// query hash for this anomaly
	// Required: true
	QueryHash *string `json:"queryHash"`

	// map of query params belongs to this anomaly
	// Required: true
	QueryParams map[string]string `json:"queryParams"`

	// row number for this anomaly
	// Required: true
	Row *int32 `json:"row"`

	// section number for this anomaly
	// Required: true
	Section *int32 `json:"section"`

	// startMs for this anomaly
	// Required: true
	StartMs *int64 `json:"startMs"`

	// updated epoch millis
	UpdatedEpochMillis int64 `json:"updatedEpochMillis,omitempty"`

	// updateMs for this anomaly
	// Required: true
	UpdatedMs *int64 `json:"updatedMs"`

	// updater Id
	UpdaterID string `json:"updaterId,omitempty"`
}

// Validate validates this anomaly
func (m *Anomaly) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChartHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDashboardID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndMs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginalStripes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParamHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartMs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedMs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Anomaly) validateChartHash(formats strfmt.Registry) error {

	if err := validate.Required("chartHash", "body", m.ChartHash); err != nil {
		return err
	}

	return nil
}

func (m *Anomaly) validateCol(formats strfmt.Registry) error {

	if err := validate.Required("col", "body", m.Col); err != nil {
		return err
	}

	return nil
}

func (m *Anomaly) validateDashboardID(formats strfmt.Registry) error {

	if err := validate.Required("dashboardId", "body", m.DashboardID); err != nil {
		return err
	}

	return nil
}

func (m *Anomaly) validateEndMs(formats strfmt.Registry) error {

	if err := validate.Required("endMs", "body", m.EndMs); err != nil {
		return err
	}

	return nil
}

func (m *Anomaly) validateOriginalStripes(formats strfmt.Registry) error {
	if swag.IsZero(m.OriginalStripes) { // not required
		return nil
	}

	for i := 0; i < len(m.OriginalStripes); i++ {
		if swag.IsZero(m.OriginalStripes[i]) { // not required
			continue
		}

		if m.OriginalStripes[i] != nil {
			if err := m.OriginalStripes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("originalStripes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("originalStripes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Anomaly) validateParamHash(formats strfmt.Registry) error {

	if err := validate.Required("paramHash", "body", m.ParamHash); err != nil {
		return err
	}

	return nil
}

func (m *Anomaly) validateQueryHash(formats strfmt.Registry) error {

	if err := validate.Required("queryHash", "body", m.QueryHash); err != nil {
		return err
	}

	return nil
}

func (m *Anomaly) validateQueryParams(formats strfmt.Registry) error {

	if err := validate.Required("queryParams", "body", m.QueryParams); err != nil {
		return err
	}

	return nil
}

func (m *Anomaly) validateRow(formats strfmt.Registry) error {

	if err := validate.Required("row", "body", m.Row); err != nil {
		return err
	}

	return nil
}

func (m *Anomaly) validateSection(formats strfmt.Registry) error {

	if err := validate.Required("section", "body", m.Section); err != nil {
		return err
	}

	return nil
}

func (m *Anomaly) validateStartMs(formats strfmt.Registry) error {

	if err := validate.Required("startMs", "body", m.StartMs); err != nil {
		return err
	}

	return nil
}

func (m *Anomaly) validateUpdatedMs(formats strfmt.Registry) error {

	if err := validate.Required("updatedMs", "body", m.UpdatedMs); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this anomaly based on the context it is used
func (m *Anomaly) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOriginalStripes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Anomaly) contextValidateCustomer(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "customer", "body", string(m.Customer)); err != nil {
		return err
	}

	return nil
}

func (m *Anomaly) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Anomaly) contextValidateOriginalStripes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OriginalStripes); i++ {

		if m.OriginalStripes[i] != nil {
			if err := m.OriginalStripes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("originalStripes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("originalStripes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Anomaly) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Anomaly) UnmarshalBinary(b []byte) error {
	var res Anomaly
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
