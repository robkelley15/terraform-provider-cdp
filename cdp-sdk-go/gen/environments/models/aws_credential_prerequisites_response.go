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

// AwsCredentialPrerequisitesResponse Response object for getting AWS credential prerequisites.
//
// swagger:model AwsCredentialPrerequisitesResponse
type AwsCredentialPrerequisitesResponse struct {

	// The cross-account external ID.
	// Required: true
	ExternalID *string `json:"externalId"`

	// The fine-grained policies related to each service.
	Policies []*CredentialGranularPolicyResponse `json:"policies"`

	// The related policy json encoded in base64
	// Required: true
	PolicyJSON *string `json:"policyJson"`
}

// Validate validates this aws credential prerequisites response
func (m *AwsCredentialPrerequisitesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExternalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyJSON(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsCredentialPrerequisitesResponse) validateExternalID(formats strfmt.Registry) error {

	if err := validate.Required("externalId", "body", m.ExternalID); err != nil {
		return err
	}

	return nil
}

func (m *AwsCredentialPrerequisitesResponse) validatePolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.Policies) { // not required
		return nil
	}

	for i := 0; i < len(m.Policies); i++ {
		if swag.IsZero(m.Policies[i]) { // not required
			continue
		}

		if m.Policies[i] != nil {
			if err := m.Policies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AwsCredentialPrerequisitesResponse) validatePolicyJSON(formats strfmt.Registry) error {

	if err := validate.Required("policyJson", "body", m.PolicyJSON); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this aws credential prerequisites response based on the context it is used
func (m *AwsCredentialPrerequisitesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsCredentialPrerequisitesResponse) contextValidatePolicies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Policies); i++ {

		if m.Policies[i] != nil {
			if err := m.Policies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AwsCredentialPrerequisitesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsCredentialPrerequisitesResponse) UnmarshalBinary(b []byte) error {
	var res AwsCredentialPrerequisitesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}