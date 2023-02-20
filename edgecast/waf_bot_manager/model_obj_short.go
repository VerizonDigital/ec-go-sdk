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

// ObjShort struct for ObjShort
type ObjShort struct {
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	LastModifiedDate     string `json:"last_modified_date"`
	AdditionalProperties map[string]interface{}
}

type _ObjShort ObjShort

// NewObjShort instantiates a new ObjShort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjShort(id string, name string, lastModifiedDate string) *ObjShort {
	this := ObjShort{}
	this.Id = id
	this.Name = name
	this.LastModifiedDate = lastModifiedDate
	return &this
}

// NewObjShortWithDefaults instantiates a new ObjShort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjShortWithDefaults() *ObjShort {
	this := ObjShort{}
	return &this
}

// GetId returns the Id field value
func (o *ObjShort) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObjShort) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ObjShort) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ObjShort) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ObjShort) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ObjShort) SetName(v string) {
	o.Name = v
}

// GetLastModifiedDate returns the LastModifiedDate field value
func (o *ObjShort) GetLastModifiedDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastModifiedDate
}

// GetLastModifiedDateOk returns a tuple with the LastModifiedDate field value
// and a boolean to check if the value has been set.
func (o *ObjShort) GetLastModifiedDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedDate, true
}

// SetLastModifiedDate sets field value
func (o *ObjShort) SetLastModifiedDate(v string) {
	o.LastModifiedDate = v
}

func (o ObjShort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["last_modified_date"] = o.LastModifiedDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ObjShort) UnmarshalJSON(bytes []byte) (err error) {
	varObjShort := _ObjShort{}

	if err = json.Unmarshal(bytes, &varObjShort); err == nil {
		*o = ObjShort(varObjShort)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "last_modified_date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableObjShort struct {
	value *ObjShort
	isSet bool
}

func (v NullableObjShort) Get() *ObjShort {
	return v.value
}

func (v *NullableObjShort) Set(val *ObjShort) {
	v.value = val
	v.isSet = true
}

func (v NullableObjShort) IsSet() bool {
	return v.isSet
}

func (v *NullableObjShort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjShort(val ObjShort) NullableObjShort {
	return NullableObjShort{value: &val, isSet: true}
}

func (v NullableObjShort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjShort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
