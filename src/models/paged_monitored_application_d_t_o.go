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

// PagedMonitoredApplicationDTO paged monitored application d t o
//
// swagger:model PagedMonitoredApplicationDTO
type PagedMonitoredApplicationDTO struct {

	// The id at which the current (limited) search can be continued to obtain more matching items
	Cursor string `json:"cursor,omitempty"`

	// List of requested items
	Items []*MonitoredApplicationDTO `json:"items"`

	// limit
	Limit int32 `json:"limit,omitempty"`

	// Whether more items are available for return by increment offset or cursor
	MoreItems bool `json:"moreItems,omitempty"`

	// offset
	// Read Only: true
	Offset int32 `json:"offset,omitempty"`

	// sort
	Sort *Sorting `json:"sort,omitempty"`

	// An estimate (lower-bound) of the total number of items available for return.  May not be a tight estimate for facet queries
	TotalItems int32 `json:"totalItems,omitempty"`
}

// Validate validates this paged monitored application d t o
func (m *PagedMonitoredApplicationDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PagedMonitoredApplicationDTO) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PagedMonitoredApplicationDTO) validateSort(formats strfmt.Registry) error {
	if swag.IsZero(m.Sort) { // not required
		return nil
	}

	if m.Sort != nil {
		if err := m.Sort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sort")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sort")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this paged monitored application d t o based on the context it is used
func (m *PagedMonitoredApplicationDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOffset(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PagedMonitoredApplicationDTO) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items); i++ {

		if m.Items[i] != nil {
			if err := m.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PagedMonitoredApplicationDTO) contextValidateOffset(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "offset", "body", int32(m.Offset)); err != nil {
		return err
	}

	return nil
}

func (m *PagedMonitoredApplicationDTO) contextValidateSort(ctx context.Context, formats strfmt.Registry) error {

	if m.Sort != nil {
		if err := m.Sort.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sort")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sort")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PagedMonitoredApplicationDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PagedMonitoredApplicationDTO) UnmarshalBinary(b []byte) error {
	var res PagedMonitoredApplicationDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
