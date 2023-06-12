// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// VwType Type of Virtual Warehouse: 'hive' or 'impala'.
//
// swagger:model VwType
type VwType string

func NewVwType(value VwType) *VwType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VwType.
func (m VwType) Pointer() *VwType {
	return &m
}

const (

	// VwTypeHive captures enum value "hive"
	VwTypeHive VwType = "hive"

	// VwTypeImpala captures enum value "impala"
	VwTypeImpala VwType = "impala"
)

// for schema
var vwTypeEnum []interface{}

func init() {
	var res []VwType
	if err := json.Unmarshal([]byte(`["hive","impala"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vwTypeEnum = append(vwTypeEnum, v)
	}
}

func (m VwType) validateVwTypeEnum(path, location string, value VwType) error {
	if err := validate.EnumCase(path, location, value, vwTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this vw type
func (m VwType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVwTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this vw type based on context it is used
func (m VwType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}