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

// RobotCreate The request for robot account creation.
//
// swagger:model RobotCreate
type RobotCreate struct {

	// The description of the robot
	Description string `json:"description,omitempty"`

	// The disable status of the robot
	Disable bool `json:"disable,omitempty"`

	// The duration of the robot in days
	Duration int64 `json:"duration,omitempty"`

	// The level of the robot, project or system
	Level string `json:"level,omitempty"`

	// The name of the robot
	Name string `json:"name,omitempty"`

	// permissions
	Permissions []*RobotPermission `json:"permissions"`

	// The secret of the robot
	Secret string `json:"secret,omitempty"`
}

// Validate validates this robot create
func (m *RobotCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RobotCreate) validatePermissions(formats strfmt.Registry) error {
	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	for i := 0; i < len(m.Permissions); i++ {
		if swag.IsZero(m.Permissions[i]) { // not required
			continue
		}

		if m.Permissions[i] != nil {
			if err := m.Permissions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("permissions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("permissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this robot create based on the context it is used
func (m *RobotCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePermissions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RobotCreate) contextValidatePermissions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Permissions); i++ {

		if m.Permissions[i] != nil {

			if swag.IsZero(m.Permissions[i]) { // not required
				return nil
			}

			if err := m.Permissions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("permissions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("permissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RobotCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RobotCreate) UnmarshalBinary(b []byte) error {
	var res RobotCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}