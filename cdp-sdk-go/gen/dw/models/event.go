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

// Event Represents a Event.
//
// swagger:model Event
type Event struct {

	// The name of the event.
	Event string `json:"event,omitempty"`

	// Detailed message.
	Message string `json:"message,omitempty"`

	// The ID of the operation to which the event belongs.
	OperationID string `json:"operationId,omitempty"`

	// The ID of the service to which the event belongs.
	ServiceID string `json:"serviceId,omitempty"`

	// The timestamp of the event.
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this event
func (m *Event) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Event) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this event based on context it is used
func (m *Event) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Event) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Event) UnmarshalBinary(b []byte) error {
	var res Event
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}