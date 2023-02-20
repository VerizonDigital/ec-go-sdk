// Code generated by the Code Generation Kit (CGK) using OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
//
// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

/*
WAF API

The WAF API is a RESTful server application for managing customer configuration settings.

API version: 1.0
*/

package waf_bot_manager

import (
	"encoding/json"
)

// BrowserChallengeAction struct for BrowserChallengeAction
type BrowserChallengeAction struct {
	Id                   *string `json:"id,omitempty"`
	Name                 *string `json:"name,omitempty"`
	EnfType              *string `json:"enf_type,omitempty"`
	IsCustomChallenge    *bool   `json:"is_custom_challenge,omitempty"`
	ResponseBodyBase64   *string `json:"response_body_base64,omitempty"`
	ValidForSec          *int32  `json:"valid_for_sec,omitempty"`
	Status               *int32  `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BrowserChallengeAction BrowserChallengeAction

// NewBrowserChallengeAction instantiates a new BrowserChallengeAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrowserChallengeAction() *BrowserChallengeAction {
	this := BrowserChallengeAction{}
	return &this
}

// NewBrowserChallengeActionWithDefaults instantiates a new BrowserChallengeAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrowserChallengeActionWithDefaults() *BrowserChallengeAction {
	this := BrowserChallengeAction{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BrowserChallengeAction) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserChallengeAction) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BrowserChallengeAction) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BrowserChallengeAction) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BrowserChallengeAction) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserChallengeAction) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BrowserChallengeAction) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BrowserChallengeAction) SetName(v string) {
	o.Name = &v
}

// GetEnfType returns the EnfType field value if set, zero value otherwise.
func (o *BrowserChallengeAction) GetEnfType() string {
	if o == nil || o.EnfType == nil {
		var ret string
		return ret
	}
	return *o.EnfType
}

// GetEnfTypeOk returns a tuple with the EnfType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserChallengeAction) GetEnfTypeOk() (*string, bool) {
	if o == nil || o.EnfType == nil {
		return nil, false
	}
	return o.EnfType, true
}

// HasEnfType returns a boolean if a field has been set.
func (o *BrowserChallengeAction) HasEnfType() bool {
	if o != nil && o.EnfType != nil {
		return true
	}

	return false
}

// SetEnfType gets a reference to the given string and assigns it to the EnfType field.
func (o *BrowserChallengeAction) SetEnfType(v string) {
	o.EnfType = &v
}

// GetIsCustomChallenge returns the IsCustomChallenge field value if set, zero value otherwise.
func (o *BrowserChallengeAction) GetIsCustomChallenge() bool {
	if o == nil || o.IsCustomChallenge == nil {
		var ret bool
		return ret
	}
	return *o.IsCustomChallenge
}

// GetIsCustomChallengeOk returns a tuple with the IsCustomChallenge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserChallengeAction) GetIsCustomChallengeOk() (*bool, bool) {
	if o == nil || o.IsCustomChallenge == nil {
		return nil, false
	}
	return o.IsCustomChallenge, true
}

// HasIsCustomChallenge returns a boolean if a field has been set.
func (o *BrowserChallengeAction) HasIsCustomChallenge() bool {
	if o != nil && o.IsCustomChallenge != nil {
		return true
	}

	return false
}

// SetIsCustomChallenge gets a reference to the given bool and assigns it to the IsCustomChallenge field.
func (o *BrowserChallengeAction) SetIsCustomChallenge(v bool) {
	o.IsCustomChallenge = &v
}

// GetResponseBodyBase64 returns the ResponseBodyBase64 field value if set, zero value otherwise.
func (o *BrowserChallengeAction) GetResponseBodyBase64() string {
	if o == nil || o.ResponseBodyBase64 == nil {
		var ret string
		return ret
	}
	return *o.ResponseBodyBase64
}

// GetResponseBodyBase64Ok returns a tuple with the ResponseBodyBase64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserChallengeAction) GetResponseBodyBase64Ok() (*string, bool) {
	if o == nil || o.ResponseBodyBase64 == nil {
		return nil, false
	}
	return o.ResponseBodyBase64, true
}

// HasResponseBodyBase64 returns a boolean if a field has been set.
func (o *BrowserChallengeAction) HasResponseBodyBase64() bool {
	if o != nil && o.ResponseBodyBase64 != nil {
		return true
	}

	return false
}

// SetResponseBodyBase64 gets a reference to the given string and assigns it to the ResponseBodyBase64 field.
func (o *BrowserChallengeAction) SetResponseBodyBase64(v string) {
	o.ResponseBodyBase64 = &v
}

// GetValidForSec returns the ValidForSec field value if set, zero value otherwise.
func (o *BrowserChallengeAction) GetValidForSec() int32 {
	if o == nil || o.ValidForSec == nil {
		var ret int32
		return ret
	}
	return *o.ValidForSec
}

// GetValidForSecOk returns a tuple with the ValidForSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserChallengeAction) GetValidForSecOk() (*int32, bool) {
	if o == nil || o.ValidForSec == nil {
		return nil, false
	}
	return o.ValidForSec, true
}

// HasValidForSec returns a boolean if a field has been set.
func (o *BrowserChallengeAction) HasValidForSec() bool {
	if o != nil && o.ValidForSec != nil {
		return true
	}

	return false
}

// SetValidForSec gets a reference to the given int32 and assigns it to the ValidForSec field.
func (o *BrowserChallengeAction) SetValidForSec(v int32) {
	o.ValidForSec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BrowserChallengeAction) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserChallengeAction) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BrowserChallengeAction) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *BrowserChallengeAction) SetStatus(v int32) {
	o.Status = &v
}

func (o BrowserChallengeAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.EnfType != nil {
		toSerialize["enf_type"] = o.EnfType
	}
	if o.IsCustomChallenge != nil {
		toSerialize["is_custom_challenge"] = o.IsCustomChallenge
	}
	if o.ResponseBodyBase64 != nil {
		toSerialize["response_body_base64"] = o.ResponseBodyBase64
	}
	if o.ValidForSec != nil {
		toSerialize["valid_for_sec"] = o.ValidForSec
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BrowserChallengeAction) UnmarshalJSON(bytes []byte) (err error) {
	varBrowserChallengeAction := _BrowserChallengeAction{}

	if err = json.Unmarshal(bytes, &varBrowserChallengeAction); err == nil {
		*o = BrowserChallengeAction(varBrowserChallengeAction)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "enf_type")
		delete(additionalProperties, "is_custom_challenge")
		delete(additionalProperties, "response_body_base64")
		delete(additionalProperties, "valid_for_sec")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBrowserChallengeAction struct {
	value *BrowserChallengeAction
	isSet bool
}

func (v NullableBrowserChallengeAction) Get() *BrowserChallengeAction {
	return v.value
}

func (v *NullableBrowserChallengeAction) Set(val *BrowserChallengeAction) {
	v.value = val
	v.isSet = true
}

func (v NullableBrowserChallengeAction) IsSet() bool {
	return v.isSet
}

func (v *NullableBrowserChallengeAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrowserChallengeAction(val BrowserChallengeAction) NullableBrowserChallengeAction {
	return NullableBrowserChallengeAction{value: &val, isSet: true}
}

func (v NullableBrowserChallengeAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrowserChallengeAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
