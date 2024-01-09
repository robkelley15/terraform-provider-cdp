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

// CreateAWSGovCloudCredentialRequest Request object for a create AWS credential request for GovCloud.
//
// swagger:model CreateAWSGovCloudCredentialRequest
type CreateAWSGovCloudCredentialRequest struct {

	// The name of the credential.
	// Required: true
	CredentialName *string `json:"credentialName"`

	// A description for the credential.
	Description string `json:"description,omitempty"`

	// The ARN of the delegated access role.
	// Required: true
	RoleArn *string `json:"roleArn"`
}

// Validate validates this create a w s gov cloud credential request
func (m *CreateAWSGovCloudCredentialRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleArn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAWSGovCloudCredentialRequest) validateCredentialName(formats strfmt.Registry) error {

	if err := validate.Required("credentialName", "body", m.CredentialName); err != nil {
		return err
	}

	return nil
}

func (m *CreateAWSGovCloudCredentialRequest) validateRoleArn(formats strfmt.Registry) error {

	if err := validate.Required("roleArn", "body", m.RoleArn); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create a w s gov cloud credential request based on context it is used
func (m *CreateAWSGovCloudCredentialRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateAWSGovCloudCredentialRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAWSGovCloudCredentialRequest) UnmarshalBinary(b []byte) error {
	var res CreateAWSGovCloudCredentialRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}