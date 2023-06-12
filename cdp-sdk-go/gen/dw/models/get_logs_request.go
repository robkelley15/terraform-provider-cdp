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

// GetLogsRequest Request object for the get logs request.
//
// swagger:model GetLogsRequest
type GetLogsRequest struct {

	// The CRN of the backup, restore or backup deletion.
	// Required: true
	Crn *string `json:"crn"`
}

// Validate validates this get logs request
func (m *GetLogsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetLogsRequest) validateCrn(formats strfmt.Registry) error {

	if err := validate.Required("crn", "body", m.Crn); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get logs request based on context it is used
func (m *GetLogsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetLogsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetLogsRequest) UnmarshalBinary(b []byte) error {
	var res GetLogsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}