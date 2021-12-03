/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 0.9.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// MinMax struct for MinMax
type MinMax struct {
	Min *float32 `json:"min,omitempty"`
	Max *float32 `json:"max,omitempty"`
}

// NewMinMax instantiates a new MinMax object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinMax() *MinMax {
	this := MinMax{}
	return &this
}

// NewMinMaxWithDefaults instantiates a new MinMax object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinMaxWithDefaults() *MinMax {
	this := MinMax{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *MinMax) GetMin() float32 {
	if o == nil || o.Min == nil {
		var ret float32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinMax) GetMinOk() (*float32, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *MinMax) HasMin() bool {
	if o != nil && o.Min != nil {
		return true
	}

	return false
}

// SetMin gets a reference to the given float32 and assigns it to the Min field.
func (o *MinMax) SetMin(v float32) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *MinMax) GetMax() float32 {
	if o == nil || o.Max == nil {
		var ret float32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinMax) GetMaxOk() (*float32, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *MinMax) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given float32 and assigns it to the Max field.
func (o *MinMax) SetMax(v float32) {
	o.Max = &v
}

func (o MinMax) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	return json.Marshal(toSerialize)
}

type NullableMinMax struct {
	value *MinMax
	isSet bool
}

func (v NullableMinMax) Get() *MinMax {
	return v.value
}

func (v *NullableMinMax) Set(val *MinMax) {
	v.value = val
	v.isSet = true
}

func (v NullableMinMax) IsSet() bool {
	return v.isSet
}

func (v *NullableMinMax) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinMax(val *MinMax) *NullableMinMax {
	return &NullableMinMax{value: val, isSet: true}
}

func (v NullableMinMax) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinMax) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


