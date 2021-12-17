// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0 license. See LICENSE file in project root for terms.

package profiles_cdn

// This file was generated by codegen-sdk-go.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"

	"github.com/EdgeCast/ec-sdk-go/rtldmodels"
)

// NewProfilesAddCustomerSettingParams creates a new ProfilesAddCustomerSettingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProfilesAddCustomerSettingParams() *ProfilesAddCustomerSettingParams {
	return &ProfilesAddCustomerSettingParams{}
}

/* ProfilesAddCustomerSettingParams contains all the parameters to send to the API endpoint
   for the profiles add customer setting operation.

   Typically these are written to a http.Request.
*/
type ProfilesAddCustomerSettingParams struct {

	// SettingDto.
	SettingDto *rtldmodels.CdnProfileDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the profiles add customer setting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProfilesAddCustomerSettingParams) WithDefaults() *ProfilesAddCustomerSettingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the profiles add customer setting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProfilesAddCustomerSettingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WriteToRequest extracts parameters and sets for the request to be consumed
func WriteToRequestProfilesAddCustomerSettingParams(o *ProfilesAddCustomerSettingParams) (RequestParameters, error) {

	var res []error

	params := NewRequestParameters()
	if o.SettingDto != nil {
		params.Body = o.SettingDto
	}

	if len(res) > 0 {
		return *params, errors.CompositeValidationError(res...)
	}
	return *params, nil
}
