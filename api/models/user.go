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

// User user
//
// swagger:model User
type User struct {

	// Email of the user.
	// Required: true
	Email *string `json:"email"`

	// List of Homes of the user, with their IDs.
	Homes []*UserHomesItems0 `json:"homes"`

	// Unique indentifier of the user.
	ID string `json:"id,omitempty"`

	// Locale string
	Locale string `json:"locale,omitempty"`

	// Full name of the user.
	// Required: true
	Name *string `json:"name"`

	// Username of the user.
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHomes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *User) validateHomes(formats strfmt.Registry) error {
	if swag.IsZero(m.Homes) { // not required
		return nil
	}

	for i := 0; i < len(m.Homes); i++ {
		if swag.IsZero(m.Homes[i]) { // not required
			continue
		}

		if m.Homes[i] != nil {
			if err := m.Homes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("homes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *User) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this user based on the context it is used
func (m *User) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHomes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) contextValidateHomes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Homes); i++ {

		if m.Homes[i] != nil {
			if err := m.Homes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("homes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UserHomesItems0 user homes items0
//
// swagger:model UserHomesItems0
type UserHomesItems0 struct {

	// The ID of the home
	ID int32 `json:"id,omitempty"`
}

// Validate validates this user homes items0
func (m *UserHomesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user homes items0 based on context it is used
func (m *UserHomesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserHomesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserHomesItems0) UnmarshalBinary(b []byte) error {
	var res UserHomesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
