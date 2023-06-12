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

// SetWorkloadPasswordPolicyRequest Request object for a set workload password policy request.
//
// swagger:model SetWorkloadPasswordPolicyRequest
type SetWorkloadPasswordPolicyRequest struct {

	// The global password policy object. If set, maxPasswordLifetimeDays is ignored, and if not set the default values for the different password policies are used. See PasswordPolicy for more details on the different default values.
	GlobalPasswordPolicy *PasswordPolicy `json:"globalPasswordPolicy,omitempty"`

	// The password policy object for machine users. If set, this will be used for enforcing password complexity for machine users instead of the global password policy.
	MachineUsersPasswordPolicy *PasswordPolicy `json:"machineUsersPasswordPolicy,omitempty"`
}

// Validate validates this set workload password policy request
func (m *SetWorkloadPasswordPolicyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGlobalPasswordPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachineUsersPasswordPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetWorkloadPasswordPolicyRequest) validateGlobalPasswordPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.GlobalPasswordPolicy) { // not required
		return nil
	}

	if m.GlobalPasswordPolicy != nil {
		if err := m.GlobalPasswordPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("globalPasswordPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("globalPasswordPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *SetWorkloadPasswordPolicyRequest) validateMachineUsersPasswordPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.MachineUsersPasswordPolicy) { // not required
		return nil
	}

	if m.MachineUsersPasswordPolicy != nil {
		if err := m.MachineUsersPasswordPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("machineUsersPasswordPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("machineUsersPasswordPolicy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this set workload password policy request based on the context it is used
func (m *SetWorkloadPasswordPolicyRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGlobalPasswordPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMachineUsersPasswordPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetWorkloadPasswordPolicyRequest) contextValidateGlobalPasswordPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.GlobalPasswordPolicy != nil {
		if err := m.GlobalPasswordPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("globalPasswordPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("globalPasswordPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *SetWorkloadPasswordPolicyRequest) contextValidateMachineUsersPasswordPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.MachineUsersPasswordPolicy != nil {
		if err := m.MachineUsersPasswordPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("machineUsersPasswordPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("machineUsersPasswordPolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SetWorkloadPasswordPolicyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetWorkloadPasswordPolicyRequest) UnmarshalBinary(b []byte) error {
	var res SetWorkloadPasswordPolicyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}