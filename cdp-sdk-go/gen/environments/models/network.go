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

// Network The network.
//
// swagger:model Network
type Network struct {

	// AWS network parameters.
	Aws *NetworkAwsParams `json:"aws,omitempty"`

	// Azure network parameters.
	Azure *NetworkAzureParams `json:"azure,omitempty"`

	// The scheme for the endpoint gateway. PUBLIC indicates an external endpoint that can be accessed over the Internet.
	// Enum: [PUBLIC PRIVATE]
	EndpointAccessGatewayScheme string `json:"endpointAccessGatewayScheme,omitempty"`

	// The subnets to use for endpoint access gateway.
	// Unique: true
	EndpointAccessGatewaySubnetIds []string `json:"endpointAccessGatewaySubnetIds"`

	// GCP network parameters.
	Gcp *NetworkGcpParams `json:"gcp,omitempty"`

	// The range of private IPv4 addresses that resources will use under this network.
	NetworkCidr string `json:"networkCidr,omitempty"`

	// Name or id of the network
	// Required: true
	NetworkName *string `json:"networkName"`

	// Subnet names or ids of the network.
	// Required: true
	// Unique: true
	SubnetIds []string `json:"subnetIds"`

	// Additional subnet metadata of the network.
	SubnetMetadata map[string]CloudSubnet `json:"subnetMetadata,omitempty"`
}

// Validate validates this network
func (m *Network) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpointAccessGatewayScheme(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpointAccessGatewaySubnetIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnetIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnetMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Network) validateAws(formats strfmt.Registry) error {
	if swag.IsZero(m.Aws) { // not required
		return nil
	}

	if m.Aws != nil {
		if err := m.Aws.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *Network) validateAzure(formats strfmt.Registry) error {
	if swag.IsZero(m.Azure) { // not required
		return nil
	}

	if m.Azure != nil {
		if err := m.Azure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

var networkTypeEndpointAccessGatewaySchemePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PUBLIC","PRIVATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkTypeEndpointAccessGatewaySchemePropEnum = append(networkTypeEndpointAccessGatewaySchemePropEnum, v)
	}
}

const (

	// NetworkEndpointAccessGatewaySchemePUBLIC captures enum value "PUBLIC"
	NetworkEndpointAccessGatewaySchemePUBLIC string = "PUBLIC"

	// NetworkEndpointAccessGatewaySchemePRIVATE captures enum value "PRIVATE"
	NetworkEndpointAccessGatewaySchemePRIVATE string = "PRIVATE"
)

// prop value enum
func (m *Network) validateEndpointAccessGatewaySchemeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, networkTypeEndpointAccessGatewaySchemePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Network) validateEndpointAccessGatewayScheme(formats strfmt.Registry) error {
	if swag.IsZero(m.EndpointAccessGatewayScheme) { // not required
		return nil
	}

	// value enum
	if err := m.validateEndpointAccessGatewaySchemeEnum("endpointAccessGatewayScheme", "body", m.EndpointAccessGatewayScheme); err != nil {
		return err
	}

	return nil
}

func (m *Network) validateEndpointAccessGatewaySubnetIds(formats strfmt.Registry) error {
	if swag.IsZero(m.EndpointAccessGatewaySubnetIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("endpointAccessGatewaySubnetIds", "body", m.EndpointAccessGatewaySubnetIds); err != nil {
		return err
	}

	return nil
}

func (m *Network) validateGcp(formats strfmt.Registry) error {
	if swag.IsZero(m.Gcp) { // not required
		return nil
	}

	if m.Gcp != nil {
		if err := m.Gcp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcp")
			}
			return err
		}
	}

	return nil
}

func (m *Network) validateNetworkName(formats strfmt.Registry) error {

	if err := validate.Required("networkName", "body", m.NetworkName); err != nil {
		return err
	}

	return nil
}

func (m *Network) validateSubnetIds(formats strfmt.Registry) error {

	if err := validate.Required("subnetIds", "body", m.SubnetIds); err != nil {
		return err
	}

	if err := validate.UniqueItems("subnetIds", "body", m.SubnetIds); err != nil {
		return err
	}

	return nil
}

func (m *Network) validateSubnetMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.SubnetMetadata) { // not required
		return nil
	}

	for k := range m.SubnetMetadata {

		if err := validate.Required("subnetMetadata"+"."+k, "body", m.SubnetMetadata[k]); err != nil {
			return err
		}
		if val, ok := m.SubnetMetadata[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subnetMetadata" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subnetMetadata" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this network based on the context it is used
func (m *Network) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAws(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzure(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGcp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubnetMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Network) contextValidateAws(ctx context.Context, formats strfmt.Registry) error {

	if m.Aws != nil {
		if err := m.Aws.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *Network) contextValidateAzure(ctx context.Context, formats strfmt.Registry) error {

	if m.Azure != nil {
		if err := m.Azure.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

func (m *Network) contextValidateGcp(ctx context.Context, formats strfmt.Registry) error {

	if m.Gcp != nil {
		if err := m.Gcp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcp")
			}
			return err
		}
	}

	return nil
}

func (m *Network) contextValidateSubnetMetadata(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.SubnetMetadata {

		if val, ok := m.SubnetMetadata[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Network) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Network) UnmarshalBinary(b []byte) error {
	var res Network
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}