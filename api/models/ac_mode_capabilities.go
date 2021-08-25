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
)

// AcModeCapabilities ac mode capabilities
//
// swagger:model AcModeCapabilities
type AcModeCapabilities struct {

	// Cooling system fan speed.
	FanSpeeds []AcFanSpeed `json:"fanSpeeds"`

	// Cooling system swing mode.
	Swings []Power `json:"swings"`

	// temperatures
	Temperatures *TemperatureRange `json:"temperatures,omitempty"`
}

// Validate validates this ac mode capabilities
func (m *AcModeCapabilities) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFanSpeeds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemperatures(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AcModeCapabilities) validateFanSpeeds(formats strfmt.Registry) error {
	if swag.IsZero(m.FanSpeeds) { // not required
		return nil
	}

	for i := 0; i < len(m.FanSpeeds); i++ {

		if err := m.FanSpeeds[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fanSpeeds" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AcModeCapabilities) validateSwings(formats strfmt.Registry) error {
	if swag.IsZero(m.Swings) { // not required
		return nil
	}

	for i := 0; i < len(m.Swings); i++ {

		if err := m.Swings[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("swings" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AcModeCapabilities) validateTemperatures(formats strfmt.Registry) error {
	if swag.IsZero(m.Temperatures) { // not required
		return nil
	}

	if m.Temperatures != nil {
		if err := m.Temperatures.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("temperatures")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ac mode capabilities based on the context it is used
func (m *AcModeCapabilities) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFanSpeeds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSwings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTemperatures(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AcModeCapabilities) contextValidateFanSpeeds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FanSpeeds); i++ {

		if err := m.FanSpeeds[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fanSpeeds" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AcModeCapabilities) contextValidateSwings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Swings); i++ {

		if err := m.Swings[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("swings" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *AcModeCapabilities) contextValidateTemperatures(ctx context.Context, formats strfmt.Registry) error {

	if m.Temperatures != nil {
		if err := m.Temperatures.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("temperatures")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AcModeCapabilities) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AcModeCapabilities) UnmarshalBinary(b []byte) error {
	var res AcModeCapabilities
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
