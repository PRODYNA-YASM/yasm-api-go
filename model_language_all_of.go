/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.4.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// LanguageAllOf struct for LanguageAllOf
type LanguageAllOf struct {
	// Two letter ISO code
	Code *string `json:"code,omitempty"`
	// The native description of a language
	NativeName *string `json:"nativeName,omitempty"`
}

// NewLanguageAllOf instantiates a new LanguageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLanguageAllOf() *LanguageAllOf {
	this := LanguageAllOf{}
	return &this
}

// NewLanguageAllOfWithDefaults instantiates a new LanguageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLanguageAllOfWithDefaults() *LanguageAllOf {
	this := LanguageAllOf{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *LanguageAllOf) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LanguageAllOf) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *LanguageAllOf) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *LanguageAllOf) SetCode(v string) {
	o.Code = &v
}

// GetNativeName returns the NativeName field value if set, zero value otherwise.
func (o *LanguageAllOf) GetNativeName() string {
	if o == nil || o.NativeName == nil {
		var ret string
		return ret
	}
	return *o.NativeName
}

// GetNativeNameOk returns a tuple with the NativeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LanguageAllOf) GetNativeNameOk() (*string, bool) {
	if o == nil || o.NativeName == nil {
		return nil, false
	}
	return o.NativeName, true
}

// HasNativeName returns a boolean if a field has been set.
func (o *LanguageAllOf) HasNativeName() bool {
	if o != nil && o.NativeName != nil {
		return true
	}

	return false
}

// SetNativeName gets a reference to the given string and assigns it to the NativeName field.
func (o *LanguageAllOf) SetNativeName(v string) {
	o.NativeName = &v
}

func (o LanguageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.NativeName != nil {
		toSerialize["nativeName"] = o.NativeName
	}
	return json.Marshal(toSerialize)
}

type NullableLanguageAllOf struct {
	value *LanguageAllOf
	isSet bool
}

func (v NullableLanguageAllOf) Get() *LanguageAllOf {
	return v.value
}

func (v *NullableLanguageAllOf) Set(val *LanguageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLanguageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLanguageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLanguageAllOf(val *LanguageAllOf) *NullableLanguageAllOf {
	return &NullableLanguageAllOf{value: val, isSet: true}
}

func (v NullableLanguageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLanguageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


