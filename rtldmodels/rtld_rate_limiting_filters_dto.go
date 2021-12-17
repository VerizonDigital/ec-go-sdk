// Code generated by go-swagger; DO NOT EDIT.

package rtldmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RtldRateLimitingFiltersDto rtld rate limiting filters dto
//
// swagger:model RtldRateLimitingFiltersDto
type RtldRateLimitingFiltersDto struct {
	BaseRtldFiltersDto

	// action limit id
	ActionLimitID []string `json:"action_limit_id"`

	// action type
	ActionType []string `json:"action_type"`

	// action type condition
	ActionTypeCondition FilterConditions `json:"action_type_condition,omitempty"`

	// client ip
	ClientIP []string `json:"client_ip"`

	// client ip condition
	ClientIPCondition FilterConditions `json:"client_ip_condition,omitempty"`

	// country code
	CountryCode []string `json:"country_code"`

	// country code condition
	CountryCodeCondition FilterConditions `json:"country_code_condition,omitempty"`

	// request method
	RequestMethod []string `json:"request_method"`

	// request method condition
	RequestMethodCondition FilterConditions `json:"request_method_condition,omitempty"`

	// scope name
	ScopeName []string `json:"scope_name"`

	// url regexp
	URLRegexp string `json:"url_regexp,omitempty"`

	// user agent regexp
	UserAgentRegexp string `json:"user_agent_regexp,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *RtldRateLimitingFiltersDto) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseRtldFiltersDto
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseRtldFiltersDto = aO0

	// AO1
	var dataAO1 struct {
		ActionLimitID []string `json:"action_limit_id"`

		ActionType []string `json:"action_type"`

		ActionTypeCondition FilterConditions `json:"action_type_condition,omitempty"`

		ClientIP []string `json:"client_ip"`

		ClientIPCondition FilterConditions `json:"client_ip_condition,omitempty"`

		CountryCode []string `json:"country_code"`

		CountryCodeCondition FilterConditions `json:"country_code_condition,omitempty"`

		RequestMethod []string `json:"request_method"`

		RequestMethodCondition FilterConditions `json:"request_method_condition,omitempty"`

		ScopeName []string `json:"scope_name"`

		URLRegexp string `json:"url_regexp,omitempty"`

		UserAgentRegexp string `json:"user_agent_regexp,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ActionLimitID = dataAO1.ActionLimitID

	m.ActionType = dataAO1.ActionType

	m.ActionTypeCondition = dataAO1.ActionTypeCondition

	m.ClientIP = dataAO1.ClientIP

	m.ClientIPCondition = dataAO1.ClientIPCondition

	m.CountryCode = dataAO1.CountryCode

	m.CountryCodeCondition = dataAO1.CountryCodeCondition

	m.RequestMethod = dataAO1.RequestMethod

	m.RequestMethodCondition = dataAO1.RequestMethodCondition

	m.ScopeName = dataAO1.ScopeName

	m.URLRegexp = dataAO1.URLRegexp

	m.UserAgentRegexp = dataAO1.UserAgentRegexp

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m RtldRateLimitingFiltersDto) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BaseRtldFiltersDto)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		ActionLimitID []string `json:"action_limit_id"`

		ActionType []string `json:"action_type"`

		ActionTypeCondition FilterConditions `json:"action_type_condition,omitempty"`

		ClientIP []string `json:"client_ip"`

		ClientIPCondition FilterConditions `json:"client_ip_condition,omitempty"`

		CountryCode []string `json:"country_code"`

		CountryCodeCondition FilterConditions `json:"country_code_condition,omitempty"`

		RequestMethod []string `json:"request_method"`

		RequestMethodCondition FilterConditions `json:"request_method_condition,omitempty"`

		ScopeName []string `json:"scope_name"`

		URLRegexp string `json:"url_regexp,omitempty"`

		UserAgentRegexp string `json:"user_agent_regexp,omitempty"`
	}

	dataAO1.ActionLimitID = m.ActionLimitID

	dataAO1.ActionType = m.ActionType

	dataAO1.ActionTypeCondition = m.ActionTypeCondition

	dataAO1.ClientIP = m.ClientIP

	dataAO1.ClientIPCondition = m.ClientIPCondition

	dataAO1.CountryCode = m.CountryCode

	dataAO1.CountryCodeCondition = m.CountryCodeCondition

	dataAO1.RequestMethod = m.RequestMethod

	dataAO1.RequestMethodCondition = m.RequestMethodCondition

	dataAO1.ScopeName = m.ScopeName

	dataAO1.URLRegexp = m.URLRegexp

	dataAO1.UserAgentRegexp = m.UserAgentRegexp

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this rtld rate limiting filters dto
func (m *RtldRateLimitingFiltersDto) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseRtldFiltersDto
	if err := m.BaseRtldFiltersDto.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActionTypeCondition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientIPCondition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountryCodeCondition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestMethodCondition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RtldRateLimitingFiltersDto) validateActionTypeCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.ActionTypeCondition) { // not required
		return nil
	}

	if err := m.ActionTypeCondition.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("action_type_condition")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("action_type_condition")
		}
		return err
	}

	return nil
}

func (m *RtldRateLimitingFiltersDto) validateClientIPCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.ClientIPCondition) { // not required
		return nil
	}

	if err := m.ClientIPCondition.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("client_ip_condition")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("client_ip_condition")
		}
		return err
	}

	return nil
}

func (m *RtldRateLimitingFiltersDto) validateCountryCodeCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.CountryCodeCondition) { // not required
		return nil
	}

	if err := m.CountryCodeCondition.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("country_code_condition")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("country_code_condition")
		}
		return err
	}

	return nil
}

func (m *RtldRateLimitingFiltersDto) validateRequestMethodCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestMethodCondition) { // not required
		return nil
	}

	if err := m.RequestMethodCondition.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("request_method_condition")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("request_method_condition")
		}
		return err
	}

	return nil
}

// ContextValidate validate this rtld rate limiting filters dto based on the context it is used
func (m *RtldRateLimitingFiltersDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseRtldFiltersDto
	if err := m.BaseRtldFiltersDto.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateActionTypeCondition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClientIPCondition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCountryCodeCondition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequestMethodCondition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RtldRateLimitingFiltersDto) contextValidateActionTypeCondition(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ActionTypeCondition.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("action_type_condition")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("action_type_condition")
		}
		return err
	}

	return nil
}

func (m *RtldRateLimitingFiltersDto) contextValidateClientIPCondition(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ClientIPCondition.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("client_ip_condition")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("client_ip_condition")
		}
		return err
	}

	return nil
}

func (m *RtldRateLimitingFiltersDto) contextValidateCountryCodeCondition(ctx context.Context, formats strfmt.Registry) error {

	if err := m.CountryCodeCondition.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("country_code_condition")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("country_code_condition")
		}
		return err
	}

	return nil
}

func (m *RtldRateLimitingFiltersDto) contextValidateRequestMethodCondition(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RequestMethodCondition.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("request_method_condition")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("request_method_condition")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RtldRateLimitingFiltersDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RtldRateLimitingFiltersDto) UnmarshalBinary(b []byte) error {
	var res RtldRateLimitingFiltersDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
