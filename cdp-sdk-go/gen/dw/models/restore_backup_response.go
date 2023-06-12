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

// RestoreBackupResponse Response object for the restore backup request.
//
// swagger:model RestoreBackupResponse
type RestoreBackupResponse struct {

	// The CRN of the restore.
	// Required: true
	RestoreCrn *string `json:"restoreCrn"`
}

// Validate validates this restore backup response
func (m *RestoreBackupResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRestoreCrn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreBackupResponse) validateRestoreCrn(formats strfmt.Registry) error {

	if err := validate.Required("restoreCrn", "body", m.RestoreCrn); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this restore backup response based on context it is used
func (m *RestoreBackupResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestoreBackupResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestoreBackupResponse) UnmarshalBinary(b []byte) error {
	var res RestoreBackupResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}