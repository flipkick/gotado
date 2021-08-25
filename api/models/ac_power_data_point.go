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

// AcPowerDataPoint ac power data point
//
// swagger:model AcPowerDataPoint
type AcPowerDataPoint struct {

	// Measurement timestamp
	// Required: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp"`

	// AC Power state (OFF, ON)
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this ac power data point
func (m *AcPowerDataPoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AcPowerDataPoint) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AcPowerDataPoint) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this ac power data point based on context it is used
func (m *AcPowerDataPoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AcPowerDataPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AcPowerDataPoint) UnmarshalBinary(b []byte) error {
	var res AcPowerDataPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
