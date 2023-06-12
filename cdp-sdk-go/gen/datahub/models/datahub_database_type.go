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

// DatahubDatabaseType Database type for datahub. Currently supported values: NONE, NON_HA, HA
//
// swagger:model DatahubDatabaseType
type DatahubDatabaseType string

func NewDatahubDatabaseType(value DatahubDatabaseType) *DatahubDatabaseType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated DatahubDatabaseType.
func (m DatahubDatabaseType) Pointer() *DatahubDatabaseType {
	return &m
}

const (

	// DatahubDatabaseTypeHA captures enum value "HA"
	DatahubDatabaseTypeHA DatahubDatabaseType = "HA"

	// DatahubDatabaseTypeNONHA captures enum value "NON_HA"
	DatahubDatabaseTypeNONHA DatahubDatabaseType = "NON_HA"

	// DatahubDatabaseTypeNONE captures enum value "NONE"
	DatahubDatabaseTypeNONE DatahubDatabaseType = "NONE"
)

// for schema
var datahubDatabaseTypeEnum []interface{}

func init() {
	var res []DatahubDatabaseType
	if err := json.Unmarshal([]byte(`["HA","NON_HA","NONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		datahubDatabaseTypeEnum = append(datahubDatabaseTypeEnum, v)
	}
}

func (m DatahubDatabaseType) validateDatahubDatabaseTypeEnum(path, location string, value DatahubDatabaseType) error {
	if err := validate.EnumCase(path, location, value, datahubDatabaseTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this datahub database type
func (m DatahubDatabaseType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDatahubDatabaseTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this datahub database type based on context it is used
func (m DatahubDatabaseType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}