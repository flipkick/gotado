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

// Link Details about the current link of a system.
//
// swagger:model Link
type Link struct {

	// reason
	Reason *LinkReason `json:"reason,omitempty"`

	// The state of the link.
	// Required: true
	// Enum: [ONLINE OFFLINE]
	State *string `json:"state"`
}

// Validate validates this link
func (m *Link) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Link) validateReason(formats strfmt.Registry) error {
	if swag.IsZero(m.Reason) { // not required
		return nil
	}

	if m.Reason != nil {
		if err := m.Reason.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reason")
			}
			return err
		}
	}

	return nil
}

var linkTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ONLINE","OFFLINE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		linkTypeStatePropEnum = append(linkTypeStatePropEnum, v)
	}
}

const (

	// LinkStateONLINE captures enum value "ONLINE"
	LinkStateONLINE string = "ONLINE"

	// LinkStateOFFLINE captures enum value "OFFLINE"
	LinkStateOFFLINE string = "OFFLINE"
)

// prop value enum
func (m *Link) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, linkTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Link) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this link based on the context it is used
func (m *Link) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Link) contextValidateReason(ctx context.Context, formats strfmt.Registry) error {

	if m.Reason != nil {
		if err := m.Reason.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reason")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Link) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Link) UnmarshalBinary(b []byte) error {
	var res Link
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LinkReason Details about the current state.
//
// swagger:model LinkReason
type LinkReason struct {

	// Message key intended for i18n.
	Code string `json:"code,omitempty"`

	// Short english error description, should not be presented to the user.
	Title string `json:"title,omitempty"`
}

// Validate validates this link reason
func (m *LinkReason) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this link reason based on context it is used
func (m *LinkReason) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LinkReason) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinkReason) UnmarshalBinary(b []byte) error {
	var res LinkReason
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
