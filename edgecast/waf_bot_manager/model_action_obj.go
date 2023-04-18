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

// ActionObj struct for ActionObj
type ActionObj struct {
	ALERT                *AlertAction            `json:"ALERT,omitempty"`
	CUSTOM_RESPONSE      *CustomResponseAction   `json:"CUSTOM_RESPONSE,omitempty"`
	BLOCK_REQUEST        *BlockRequestAction     `json:"BLOCK_REQUEST,omitempty"`
	REDIRECT302          *RedirectAction         `json:"REDIRECT_302,omitempty"`
	BROWSER_CHALLENGE    *BrowserChallengeAction `json:"BROWSER_CHALLENGE,omitempty"`
	RECAPTCHA    		 *RecaptchaAction 		 `json:"RECAPTCHA,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ActionObj ActionObj

// NewActionObj instantiates a new ActionObj object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionObj() *ActionObj {
	this := ActionObj{}
	return &this
}

// NewActionObjWithDefaults instantiates a new ActionObj object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionObjWithDefaults() *ActionObj {
	this := ActionObj{}
	return &this
}

// GetALERT returns the ALERT field value if set, zero value otherwise.
func (o *ActionObj) GetALERT() AlertAction {
	if o == nil || o.ALERT == nil {
		var ret AlertAction
		return ret
	}
	return *o.ALERT
}

// GetALERTOk returns a tuple with the ALERT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionObj) GetALERTOk() (*AlertAction, bool) {
	if o == nil || o.ALERT == nil {
		return nil, false
	}
	return o.ALERT, true
}

// HasALERT returns a boolean if a field has been set.
func (o *ActionObj) HasALERT() bool {
	if o != nil && o.ALERT != nil {
		return true
	}

	return false
}

// SetALERT gets a reference to the given AlertAction and assigns it to the ALERT field.
func (o *ActionObj) SetALERT(v AlertAction) {
	o.ALERT = &v
}

// GetCUSTOM_RESPONSE returns the CUSTOM_RESPONSE field value if set, zero value otherwise.
func (o *ActionObj) GetCUSTOM_RESPONSE() CustomResponseAction {
	if o == nil || o.CUSTOM_RESPONSE == nil {
		var ret CustomResponseAction
		return ret
	}
	return *o.CUSTOM_RESPONSE
}

// GetCUSTOM_RESPONSEOk returns a tuple with the CUSTOM_RESPONSE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionObj) GetCUSTOM_RESPONSEOk() (*CustomResponseAction, bool) {
	if o == nil || o.CUSTOM_RESPONSE == nil {
		return nil, false
	}
	return o.CUSTOM_RESPONSE, true
}

// HasCUSTOM_RESPONSE returns a boolean if a field has been set.
func (o *ActionObj) HasCUSTOM_RESPONSE() bool {
	if o != nil && o.CUSTOM_RESPONSE != nil {
		return true
	}

	return false
}

// SetCUSTOM_RESPONSE gets a reference to the given CustomResponseAction and assigns it to the CUSTOM_RESPONSE field.
func (o *ActionObj) SetCUSTOM_RESPONSE(v CustomResponseAction) {
	o.CUSTOM_RESPONSE = &v
}

// GetBLOCK_REQUEST returns the BLOCK_REQUEST field value if set, zero value otherwise.
func (o *ActionObj) GetBLOCK_REQUEST() BlockRequestAction {
	if o == nil || o.BLOCK_REQUEST == nil {
		var ret BlockRequestAction
		return ret
	}
	return *o.BLOCK_REQUEST
}

// GetBLOCK_REQUESTOk returns a tuple with the BLOCK_REQUEST field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionObj) GetBLOCK_REQUESTOk() (*BlockRequestAction, bool) {
	if o == nil || o.BLOCK_REQUEST == nil {
		return nil, false
	}
	return o.BLOCK_REQUEST, true
}

// HasBLOCK_REQUEST returns a boolean if a field has been set.
func (o *ActionObj) HasBLOCK_REQUEST() bool {
	if o != nil && o.BLOCK_REQUEST != nil {
		return true
	}

	return false
}

// SetBLOCK_REQUEST gets a reference to the given BlockRequestAction and assigns it to the BLOCK_REQUEST field.
func (o *ActionObj) SetBLOCK_REQUEST(v BlockRequestAction) {
	o.BLOCK_REQUEST = &v
}

// GetREDIRECT302 returns the REDIRECT302 field value if set, zero value otherwise.
func (o *ActionObj) GetREDIRECT302() RedirectAction {
	if o == nil || o.REDIRECT302 == nil {
		var ret RedirectAction
		return ret
	}
	return *o.REDIRECT302
}

// GetREDIRECT302Ok returns a tuple with the REDIRECT302 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionObj) GetREDIRECT302Ok() (*RedirectAction, bool) {
	if o == nil || o.REDIRECT302 == nil {
		return nil, false
	}
	return o.REDIRECT302, true
}

// HasREDIRECT302 returns a boolean if a field has been set.
func (o *ActionObj) HasREDIRECT302() bool {
	if o != nil && o.REDIRECT302 != nil {
		return true
	}

	return false
}

// SetREDIRECT302 gets a reference to the given RedirectAction and assigns it to the REDIRECT302 field.
func (o *ActionObj) SetREDIRECT302(v RedirectAction) {
	o.REDIRECT302 = &v
}

// GetBROWSER_CHALLENGE returns the BROWSER_CHALLENGE field value if set, zero value otherwise.
func (o *ActionObj) GetBROWSER_CHALLENGE() BrowserChallengeAction {
	if o == nil || o.BROWSER_CHALLENGE == nil {
		var ret BrowserChallengeAction
		return ret
	}
	return *o.BROWSER_CHALLENGE
}

// GetBROWSER_CHALLENGEOk returns a tuple with the BROWSER_CHALLENGE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionObj) GetBROWSER_CHALLENGEOk() (*BrowserChallengeAction, bool) {
	if o == nil || o.BROWSER_CHALLENGE == nil {
		return nil, false
	}
	return o.BROWSER_CHALLENGE, true
}

// HasBROWSER_CHALLENGE returns a boolean if a field has been set.
func (o *ActionObj) HasBROWSER_CHALLENGE() bool {
	if o != nil && o.BROWSER_CHALLENGE != nil {
		return true
	}

	return false
}

// SetBROWSER_CHALLENGE gets a reference to the given BrowserChallengeAction and assigns it to the BROWSER_CHALLENGE field.
func (o *ActionObj) SetBROWSER_CHALLENGE(v BrowserChallengeAction) {
	o.BROWSER_CHALLENGE = &v
}

// GetRECAPTCHAOk returns a tuple with the RECAPTCHA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionObj) GetRECAPTCHAOk() (*RecaptchaAction, bool) {
	if o == nil || o.RECAPTCHA == nil {
		return nil, false
	}
	return o.RECAPTCHA, true
}

// HasRECAPTCHA returns a boolean if a field has been set.
func (o *ActionObj) HasRECAPTCHA() bool {
	if o != nil && o.RECAPTCHA != nil {
		return true
	}

	return false
}

// SetRECAPTCHA gets a reference to the given RecaptchaAction and assigns it to the RECAPTCHA field.
func (o *ActionObj) SetRECAPTCHA(v RecaptchaAction) {
	o.RECAPTCHA = &v
}

func (o ActionObj) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ALERT != nil {
		toSerialize["ALERT"] = o.ALERT
	}
	if o.CUSTOM_RESPONSE != nil {
		toSerialize["CUSTOM_RESPONSE"] = o.CUSTOM_RESPONSE
	}
	if o.BLOCK_REQUEST != nil {
		toSerialize["BLOCK_REQUEST"] = o.BLOCK_REQUEST
	}
	if o.REDIRECT302 != nil {
		toSerialize["REDIRECT_302"] = o.REDIRECT302
	}
	if o.BROWSER_CHALLENGE != nil {
		toSerialize["BROWSER_CHALLENGE"] = o.BROWSER_CHALLENGE
	}
	if o.RECAPTCHA != nil{
		toSerialize["RECAPTCHA"] = o.RECAPTCHA
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ActionObj) UnmarshalJSON(bytes []byte) (err error) {
	varActionObj := _ActionObj{}

	if err = json.Unmarshal(bytes, &varActionObj); err == nil {
		*o = ActionObj(varActionObj)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ALERT")
		delete(additionalProperties, "CUSTOM_RESPONSE")
		delete(additionalProperties, "BLOCK_REQUEST")
		delete(additionalProperties, "REDIRECT_302")
		delete(additionalProperties, "BROWSER_CHALLENGE")
		delete(additionalProperties, "RECAPTCHA")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableActionObj struct {
	value *ActionObj
	isSet bool
}

func (v NullableActionObj) Get() *ActionObj {
	return v.value
}

func (v *NullableActionObj) Set(val *ActionObj) {
	v.value = val
	v.isSet = true
}

func (v NullableActionObj) IsSet() bool {
	return v.isSet
}

func (v *NullableActionObj) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionObj(val ActionObj) NullableActionObj {
	return NullableActionObj{value: &val, isSet: true}
}

func (v NullableActionObj) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionObj) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
