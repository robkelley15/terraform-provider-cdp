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

// Coprocessor An instance of a coprocessor.
//
// swagger:model Coprocessor
type Coprocessor struct {

	// The last committed command ID of the coprocessor.
	CommandID int64 `json:"commandId,omitempty"`

	// The optional coprocessor arguments in comma-separated list.
	CoprocessorArgs string `json:"coprocessorArgs,omitempty"`

	// The coprocessor canonical name. It is unique per database.
	CoprocessorCanonicalName string `json:"coprocessorCanonicalName,omitempty"`

	// The coprocessor location URL.
	CoprocessorLocationURL string `json:"coprocessorLocationUrl,omitempty"`

	// The current status when listing coprocessors.
	Status CoprocessorStatusType `json:"status,omitempty"`

	// The reason of the current status.
	StatusReason string `json:"statusReason,omitempty"`

	// Fully qualified table name.
	TableName string `json:"tableName,omitempty"`
}

// Validate validates this coprocessor
func (m *Coprocessor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Coprocessor) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// ContextValidate validate this coprocessor based on the context it is used
func (m *Coprocessor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Coprocessor) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Coprocessor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Coprocessor) UnmarshalBinary(b []byte) error {
	var res Coprocessor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}