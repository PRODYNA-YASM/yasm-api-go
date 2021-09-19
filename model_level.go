/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 0.8.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Level struct for Level
type Level struct {
	Level *int32 `json:"level,omitempty"`
}

// NewLevel instantiates a new Level object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLevel() *Level {
	this := Level{}
	return &this
}

// NewLevelWithDefaults instantiates a new Level object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLevelWithDefaults() *Level {
	this := Level{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *Level) GetLevel() int32 {
	if o == nil || o.Level == nil {
		var ret int32
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Level) GetLevelOk() (*int32, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *Level) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given int32 and assigns it to the Level field.
func (o *Level) SetLevel(v int32) {
	o.Level = &v
}

func (o Level) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	return json.Marshal(toSerialize)
}

type NullableLevel struct {
	value *Level
	isSet bool
}

func (v NullableLevel) Get() *Level {
	return v.value
}

func (v *NullableLevel) Set(val *Level) {
	v.value = val
	v.isSet = true
}

func (v NullableLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLevel(val *Level) *NullableLevel {
	return &NullableLevel{value: val, isSet: true}
}

func (v NullableLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


