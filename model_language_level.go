/*
 * YASM (Yet Another Skill Management) API
 *
 * This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.
 *
 * API version: 0.6.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// LanguageLevel struct for LanguageLevel
type LanguageLevel struct {
	Language *Language `json:"language,omitempty"`
	Level *Level `json:"level,omitempty"`
}

// NewLanguageLevel instantiates a new LanguageLevel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLanguageLevel() *LanguageLevel {
	this := LanguageLevel{}
	return &this
}

// NewLanguageLevelWithDefaults instantiates a new LanguageLevel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLanguageLevelWithDefaults() *LanguageLevel {
	this := LanguageLevel{}
	return &this
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *LanguageLevel) GetLanguage() Language {
	if o == nil || o.Language == nil {
		var ret Language
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LanguageLevel) GetLanguageOk() (*Language, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *LanguageLevel) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given Language and assigns it to the Language field.
func (o *LanguageLevel) SetLanguage(v Language) {
	o.Language = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *LanguageLevel) GetLevel() Level {
	if o == nil || o.Level == nil {
		var ret Level
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LanguageLevel) GetLevelOk() (*Level, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *LanguageLevel) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given Level and assigns it to the Level field.
func (o *LanguageLevel) SetLevel(v Level) {
	o.Level = &v
}

func (o LanguageLevel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	return json.Marshal(toSerialize)
}

type NullableLanguageLevel struct {
	value *LanguageLevel
	isSet bool
}

func (v NullableLanguageLevel) Get() *LanguageLevel {
	return v.value
}

func (v *NullableLanguageLevel) Set(val *LanguageLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableLanguageLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableLanguageLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLanguageLevel(val *LanguageLevel) *NullableLanguageLevel {
	return &NullableLanguageLevel{value: val, isSet: true}
}

func (v NullableLanguageLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLanguageLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


