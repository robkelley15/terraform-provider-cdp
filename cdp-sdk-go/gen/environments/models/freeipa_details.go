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

// FreeipaDetails Details of a FreeIPA cluster.
//
// swagger:model FreeipaDetails
type FreeipaDetails struct {

	// CRN of the FreeIPA cluster.
	Crn string `json:"crn,omitempty"`

	// The domain name of the FreeIPA cluster.
	Domain string `json:"domain,omitempty"`

	// The hostname of the FreeIPA cluster.
	Hostname string `json:"hostname,omitempty"`

	// The recipes for the FreeIPA cluster.
	Recipes []string `json:"recipes"`

	// The IP addresses of the FreeIPA cluster.
	// Unique: true
	ServerIP []string `json:"serverIP"`
}

// Validate validates this freeipa details
func (m *FreeipaDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServerIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FreeipaDetails) validateServerIP(formats strfmt.Registry) error {
	if swag.IsZero(m.ServerIP) { // not required
		return nil
	}

	if err := validate.UniqueItems("serverIP", "body", m.ServerIP); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this freeipa details based on context it is used
func (m *FreeipaDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FreeipaDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FreeipaDetails) UnmarshalBinary(b []byte) error {
	var res FreeipaDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}