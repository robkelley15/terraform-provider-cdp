// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ExternalBucketAccessInfo External bucket definition.
//
// swagger:model ExternalBucketAccessInfo
type ExternalBucketAccessInfo struct {

	// Specifies whether the external bucket will be added in read-only or read-write mode.
	// Enum: [READ READWRITE]
	AccessMode string `json:"accessMode,omitempty"`

	// Enable server side encryption with the specified KMS CMK ARN. If the value is empty or unspecified, default server side encryption will be used when the access mode is read-write. Otherwise this value is ignored.
	KmsCmkArn string `json:"kmsCmkArn,omitempty"`
}

// Validate validates this external bucket access info
func (m *ExternalBucketAccessInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var externalBucketAccessInfoTypeAccessModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["READ","READWRITE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		externalBucketAccessInfoTypeAccessModePropEnum = append(externalBucketAccessInfoTypeAccessModePropEnum, v)
	}
}

const (

	// ExternalBucketAccessInfoAccessModeREAD captures enum value "READ"
	ExternalBucketAccessInfoAccessModeREAD string = "READ"

	// ExternalBucketAccessInfoAccessModeREADWRITE captures enum value "READWRITE"
	ExternalBucketAccessInfoAccessModeREADWRITE string = "READWRITE"
)

// prop value enum
func (m *ExternalBucketAccessInfo) validateAccessModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, externalBucketAccessInfoTypeAccessModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExternalBucketAccessInfo) validateAccessMode(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccessModeEnum("accessMode", "body", m.AccessMode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this external bucket access info based on context it is used
func (m *ExternalBucketAccessInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExternalBucketAccessInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternalBucketAccessInfo) UnmarshalBinary(b []byte) error {
	var res ExternalBucketAccessInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}