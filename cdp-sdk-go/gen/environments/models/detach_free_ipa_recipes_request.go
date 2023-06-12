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

// DetachFreeIpaRecipesRequest Request object for detaching recipes from FreeIPA.
//
// swagger:model DetachFreeIpaRecipesRequest
type DetachFreeIpaRecipesRequest struct {

	// The environment name or CRN of the FreeIPA.
	// Required: true
	Environment *string `json:"environment"`

	// The list of recipes to detach.
	// Required: true
	Recipes []string `json:"recipes"`
}

// Validate validates this detach free ipa recipes request
func (m *DetachFreeIpaRecipesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DetachFreeIpaRecipesRequest) validateEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("environment", "body", m.Environment); err != nil {
		return err
	}

	return nil
}

func (m *DetachFreeIpaRecipesRequest) validateRecipes(formats strfmt.Registry) error {

	if err := validate.Required("recipes", "body", m.Recipes); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this detach free ipa recipes request based on context it is used
func (m *DetachFreeIpaRecipesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DetachFreeIpaRecipesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DetachFreeIpaRecipesRequest) UnmarshalBinary(b []byte) error {
	var res DetachFreeIpaRecipesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}