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

// MaintenanceWindow Wavefront maintenance window entity
//
// swagger:model MaintenanceWindow
type MaintenanceWindow struct {

	// created epoch millis
	// Read Only: true
	CreatedEpochMillis int64 `json:"createdEpochMillis,omitempty"`

	// creator Id
	// Read Only: true
	CreatorID string `json:"creatorId,omitempty"`

	// customer Id
	// Read Only: true
	CustomerID string `json:"customerId,omitempty"`

	// The time in epoch seconds when this maintenance window will end
	// Required: true
	EndTimeInSeconds *int64 `json:"endTimeInSeconds"`

	// The name of an event associated with the creation/update of this maintenance window
	// Read Only: true
	EventName string `json:"eventName,omitempty"`

	// If true, a source/host must be in 'relevantHostNames' and have tags matching the specification formed by 'relevantHostTags' and 'relevantHostTagsAnded' in order for this maintenance window to apply. If false, a source/host must either be in 'relevantHostNames' or match 'relevantHostTags' and 'relevantHostTagsAnded'. Default: false
	HostTagGroupHostNamesGroupAnded bool `json:"hostTagGroupHostNamesGroupAnded,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// Query that filters on point tags of timeseries scanned by alert.
	PointTagFilter string `json:"pointTagFilter,omitempty"`

	// Whether to AND point tags filter listed in pointTagFilter. If true, a timeseries must contain the point tags along with other filters in order for the maintenance window to apply.If false, the tags are OR'ed, the customer must contain one of the tags. Default: false
	PointTagFilterAnded bool `json:"pointTagFilterAnded,omitempty"`

	// The purpose of this maintenance window
	// Required: true
	Reason *string `json:"reason"`

	// List of alert tags whose matching alerts will be put into maintenance because of this maintenance window
	// Required: true
	RelevantCustomerTags []string `json:"relevantCustomerTags"`

	// Whether to AND customer tags listed in relevantCustomerTags. If true, a customer must contain all tags in order for the maintenance window to apply.  If false, the tags are OR'ed, and a customer must contain one of the tags. Default: false
	RelevantCustomerTagsAnded bool `json:"relevantCustomerTagsAnded,omitempty"`

	// List of source/host names that will be put into maintenance because of this maintenance window
	RelevantHostNames []string `json:"relevantHostNames"`

	// List of source/host tags whose matching sources/hosts will be put into maintenance because of this maintenance window
	RelevantHostTags []string `json:"relevantHostTags"`

	// Whether to AND source/host tags listed in relevantHostTags. If true, a source/host must contain all tags in order for the maintenance window to apply.  If false, the tags are OR'ed, and a source/host must contain one of the tags. Default: false
	RelevantHostTagsAnded bool `json:"relevantHostTagsAnded,omitempty"`

	// running state
	// Read Only: true
	// Enum: [ONGOING PENDING ENDED]
	RunningState string `json:"runningState,omitempty"`

	// Numeric value used in default sorting
	// Read Only: true
	SortAttr int32 `json:"sortAttr,omitempty"`

	// The time in epoch seconds when this maintenance window will start
	// Required: true
	StartTimeInSeconds *int64 `json:"startTimeInSeconds"`

	// List of targets to notify, overriding the alert's targets.
	Targets []string `json:"targets"`

	// Title of this maintenance window
	// Required: true
	Title *string `json:"title"`

	// updated epoch millis
	// Read Only: true
	UpdatedEpochMillis int64 `json:"updatedEpochMillis,omitempty"`

	// updater Id
	// Read Only: true
	UpdaterID string `json:"updaterId,omitempty"`
}

// Validate validates this maintenance window
func (m *MaintenanceWindow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndTimeInSeconds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelevantCustomerTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunningState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTimeInSeconds(formats); err != nil {
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

func (m *MaintenanceWindow) validateEndTimeInSeconds(formats strfmt.Registry) error {

	if err := validate.Required("endTimeInSeconds", "body", m.EndTimeInSeconds); err != nil {
		return err
	}

	return nil
}

func (m *MaintenanceWindow) validateReason(formats strfmt.Registry) error {

	if err := validate.Required("reason", "body", m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *MaintenanceWindow) validateRelevantCustomerTags(formats strfmt.Registry) error {

	if err := validate.Required("relevantCustomerTags", "body", m.RelevantCustomerTags); err != nil {
		return err
	}

	return nil
}

var maintenanceWindowTypeRunningStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ONGOING","PENDING","ENDED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		maintenanceWindowTypeRunningStatePropEnum = append(maintenanceWindowTypeRunningStatePropEnum, v)
	}
}

const (

	// MaintenanceWindowRunningStateONGOING captures enum value "ONGOING"
	MaintenanceWindowRunningStateONGOING string = "ONGOING"

	// MaintenanceWindowRunningStatePENDING captures enum value "PENDING"
	MaintenanceWindowRunningStatePENDING string = "PENDING"

	// MaintenanceWindowRunningStateENDED captures enum value "ENDED"
	MaintenanceWindowRunningStateENDED string = "ENDED"
)

// prop value enum
func (m *MaintenanceWindow) validateRunningStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, maintenanceWindowTypeRunningStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MaintenanceWindow) validateRunningState(formats strfmt.Registry) error {
	if swag.IsZero(m.RunningState) { // not required
		return nil
	}

	// value enum
	if err := m.validateRunningStateEnum("runningState", "body", m.RunningState); err != nil {
		return err
	}

	return nil
}

func (m *MaintenanceWindow) validateStartTimeInSeconds(formats strfmt.Registry) error {

	if err := validate.Required("startTimeInSeconds", "body", m.StartTimeInSeconds); err != nil {
		return err
	}

	return nil
}

func (m *MaintenanceWindow) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this maintenance window based on the context it is used
func (m *MaintenanceWindow) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedEpochMillis(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatorID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomerID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEventName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRunningState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSortAttr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedEpochMillis(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdaterID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MaintenanceWindow) contextValidateCreatedEpochMillis(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdEpochMillis", "body", int64(m.CreatedEpochMillis)); err != nil {
		return err
	}

	return nil
}

func (m *MaintenanceWindow) contextValidateCreatorID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "creatorId", "body", string(m.CreatorID)); err != nil {
		return err
	}

	return nil
}

func (m *MaintenanceWindow) contextValidateCustomerID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "customerId", "body", string(m.CustomerID)); err != nil {
		return err
	}

	return nil
}

func (m *MaintenanceWindow) contextValidateEventName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "eventName", "body", string(m.EventName)); err != nil {
		return err
	}

	return nil
}

func (m *MaintenanceWindow) contextValidateRunningState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "runningState", "body", string(m.RunningState)); err != nil {
		return err
	}

	return nil
}

func (m *MaintenanceWindow) contextValidateSortAttr(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "sortAttr", "body", int32(m.SortAttr)); err != nil {
		return err
	}

	return nil
}

func (m *MaintenanceWindow) contextValidateUpdatedEpochMillis(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updatedEpochMillis", "body", int64(m.UpdatedEpochMillis)); err != nil {
		return err
	}

	return nil
}

func (m *MaintenanceWindow) contextValidateUpdaterID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updaterId", "body", string(m.UpdaterID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MaintenanceWindow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MaintenanceWindow) UnmarshalBinary(b []byte) error {
	var res MaintenanceWindow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
