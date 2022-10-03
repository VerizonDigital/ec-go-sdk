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

// DomainDcvToken domain dcv token
//
// swagger:model DomainDcvToken
type DomainDcvToken struct {

	// dcv method
	// Enum: [Email DnsCnameToken DnsTxtToken]
	DcvMethod string `json:"dcv_method,omitempty"`

	// dcv token
	DcvToken *DcvToken `json:"dcv_token,omitempty"`

	// domain id
	DomainID int64 `json:"domain_id,omitempty"`
}

// Validate validates this domain dcv token
func (m *DomainDcvToken) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDcvMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDcvToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var domainDcvTokenTypeDcvMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Email","DnsCnameToken","DnsTxtToken"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		domainDcvTokenTypeDcvMethodPropEnum = append(domainDcvTokenTypeDcvMethodPropEnum, v)
	}
}

const (

	// DomainDcvTokenDcvMethodEmail captures enum value "Email"
	DomainDcvTokenDcvMethodEmail string = "Email"

	// DomainDcvTokenDcvMethodDNSCnameToken captures enum value "DnsCnameToken"
	DomainDcvTokenDcvMethodDNSCnameToken string = "DnsCnameToken"

	// DomainDcvTokenDcvMethodDNSTxtToken captures enum value "DnsTxtToken"
	DomainDcvTokenDcvMethodDNSTxtToken string = "DnsTxtToken"
)

// prop value enum
func (m *DomainDcvToken) validateDcvMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, domainDcvTokenTypeDcvMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DomainDcvToken) validateDcvMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.DcvMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateDcvMethodEnum("dcv_method", "body", m.DcvMethod); err != nil {
		return err
	}

	return nil
}

func (m *DomainDcvToken) validateDcvToken(formats strfmt.Registry) error {
	if swag.IsZero(m.DcvToken) { // not required
		return nil
	}

	if m.DcvToken != nil {
		if err := m.DcvToken.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dcv_token")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dcv_token")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this domain dcv token based on the context it is used
func (m *DomainDcvToken) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDcvToken(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainDcvToken) contextValidateDcvToken(ctx context.Context, formats strfmt.Registry) error {

	if m.DcvToken != nil {
		if err := m.DcvToken.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dcv_token")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dcv_token")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainDcvToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainDcvToken) UnmarshalBinary(b []byte) error {
	var res DomainDcvToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
