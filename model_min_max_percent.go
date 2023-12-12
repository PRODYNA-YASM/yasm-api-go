/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the MinMaxPercent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MinMaxPercent{}

// MinMaxPercent struct for MinMaxPercent
type MinMaxPercent struct {
	Min *int32 `json:"min,omitempty"`
	Max *int32 `json:"max,omitempty"`
}

// NewMinMaxPercent instantiates a new MinMaxPercent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinMaxPercent() *MinMaxPercent {
	this := MinMaxPercent{}
	return &this
}

// NewMinMaxPercentWithDefaults instantiates a new MinMaxPercent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinMaxPercentWithDefaults() *MinMaxPercent {
	this := MinMaxPercent{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *MinMaxPercent) GetMin() int32 {
	if o == nil || IsNil(o.Min) {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinMaxPercent) GetMinOk() (*int32, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *MinMaxPercent) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *MinMaxPercent) SetMin(v int32) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *MinMaxPercent) GetMax() int32 {
	if o == nil || IsNil(o.Max) {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinMaxPercent) GetMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *MinMaxPercent) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *MinMaxPercent) SetMax(v int32) {
	o.Max = &v
}

func (o MinMaxPercent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MinMaxPercent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	return toSerialize, nil
}

type NullableMinMaxPercent struct {
	value *MinMaxPercent
	isSet bool
}

func (v NullableMinMaxPercent) Get() *MinMaxPercent {
	return v.value
}

func (v *NullableMinMaxPercent) Set(val *MinMaxPercent) {
	v.value = val
	v.isSet = true
}

func (v NullableMinMaxPercent) IsSet() bool {
	return v.isSet
}

func (v *NullableMinMaxPercent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinMaxPercent(val *MinMaxPercent) *NullableMinMaxPercent {
	return &NullableMinMaxPercent{value: val, isSet: true}
}

func (v NullableMinMaxPercent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinMaxPercent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


