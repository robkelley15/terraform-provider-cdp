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

// ListDatalakeDiagnosticsResponse Response object for listing recent Datalake diagnostics collections.
//
// swagger:model ListDatalakeDiagnosticsResponse
type ListDatalakeDiagnosticsResponse struct {

	// description.
	Collections []*DatalakeDiagnosticsCollectionResponse `json:"collections"`
}

// Validate validates this list datalake diagnostics response
func (m *ListDatalakeDiagnosticsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCollections(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListDatalakeDiagnosticsResponse) validateCollections(formats strfmt.Registry) error {
	if swag.IsZero(m.Collections) { // not required
		return nil
	}

	for i := 0; i < len(m.Collections); i++ {
		if swag.IsZero(m.Collections[i]) { // not required
			continue
		}

		if m.Collections[i] != nil {
			if err := m.Collections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("collections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("collections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list datalake diagnostics response based on the context it is used
func (m *ListDatalakeDiagnosticsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCollections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListDatalakeDiagnosticsResponse) contextValidateCollections(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Collections); i++ {

		if m.Collections[i] != nil {
			if err := m.Collections[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("collections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("collections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListDatalakeDiagnosticsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListDatalakeDiagnosticsResponse) UnmarshalBinary(b []byte) error {
	var res ListDatalakeDiagnosticsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}