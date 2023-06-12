// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DataVisualizationSummary A Cloudera Data Visualization.
//
// swagger:model DataVisualizationSummary
type DataVisualizationSummary struct {

	// The CRN of the user who created the Cloudera Data Visualization
	CreatorCrn string `json:"creatorCrn,omitempty"`

	// The ID of the Cloudera Data Visualization.
	ID string `json:"id,omitempty"`

	// Current image version of the Cloudera Data Visualization
	ImageVersion string `json:"imageVersion,omitempty"`

	// The name of the Cloudera Data Visualization.
	Name string `json:"name,omitempty"`

	// The template size for the Cloudera Data Visualization
	Size string `json:"size,omitempty"`

	// Status of the Cloudera Data Visualization. Possible values are: Creating, Created, Accepted, Starting, Running, Stopping, Stopped, Updating, PreUpdate, Upgrading, PreUpgrade, Restarting, Deleting, Waiting, Failed, Error.
	Status string `json:"status,omitempty"`
}

// Validate validates this data visualization summary
func (m *DataVisualizationSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this data visualization summary based on context it is used
func (m *DataVisualizationSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataVisualizationSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataVisualizationSummary) UnmarshalBinary(b []byte) error {
	var res DataVisualizationSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}