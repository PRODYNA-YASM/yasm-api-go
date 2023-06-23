/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.4.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Language struct for Language
type Language struct {
	NamedDomainModel
	// Two letter ISO code
	Code *string `json:"code,omitempty"`
	// The native description of a language
	NativeName *string `json:"nativeName,omitempty"`
}

// NewLanguage instantiates a new Language object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLanguage(id string, name string) *Language {
	this := Language{}
	this.Id = id
	this.Name = name
	return &this
}

// NewLanguageWithDefaults instantiates a new Language object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLanguageWithDefaults() *Language {
	this := Language{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Language) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Language) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Language) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Language) SetCode(v string) {
	o.Code = &v
}

// GetNativeName returns the NativeName field value if set, zero value otherwise.
func (o *Language) GetNativeName() string {
	if o == nil || o.NativeName == nil {
		var ret string
		return ret
	}
	return *o.NativeName
}

// GetNativeNameOk returns a tuple with the NativeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Language) GetNativeNameOk() (*string, bool) {
	if o == nil || o.NativeName == nil {
		return nil, false
	}
	return o.NativeName, true
}

// HasNativeName returns a boolean if a field has been set.
func (o *Language) HasNativeName() bool {
	if o != nil && o.NativeName != nil {
		return true
	}

	return false
}

// SetNativeName gets a reference to the given string and assigns it to the NativeName field.
func (o *Language) SetNativeName(v string) {
	o.NativeName = &v
}

func (o Language) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedNamedDomainModel, errNamedDomainModel := json.Marshal(o.NamedDomainModel)
	if errNamedDomainModel != nil {
		return []byte{}, errNamedDomainModel
	}
	errNamedDomainModel = json.Unmarshal([]byte(serializedNamedDomainModel), &toSerialize)
	if errNamedDomainModel != nil {
		return []byte{}, errNamedDomainModel
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.NativeName != nil {
		toSerialize["nativeName"] = o.NativeName
	}
	return json.Marshal(toSerialize)
}

type NullableLanguage struct {
	value *Language
	isSet bool
}

func (v NullableLanguage) Get() *Language {
	return v.value
}

func (v *NullableLanguage) Set(val *Language) {
	v.value = val
	v.isSet = true
}

func (v NullableLanguage) IsSet() bool {
	return v.isSet
}

func (v *NullableLanguage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLanguage(val *Language) *NullableLanguage {
	return &NullableLanguage{value: val, isSet: true}
}

func (v NullableLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLanguage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


