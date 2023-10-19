/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.4.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// AvailabilityFilter struct for AvailabilityFilter
type AvailabilityFilter struct {
	Timeframed
	Min *int32 `json:"min,omitempty"`
	Max *int32 `json:"max,omitempty"`
}

// NewAvailabilityFilter instantiates a new AvailabilityFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailabilityFilter() *AvailabilityFilter {
	this := AvailabilityFilter{}
	return &this
}

// NewAvailabilityFilterWithDefaults instantiates a new AvailabilityFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailabilityFilterWithDefaults() *AvailabilityFilter {
	this := AvailabilityFilter{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *AvailabilityFilter) GetMin() int32 {
	if o == nil || o.Min == nil {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityFilter) GetMinOk() (*int32, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *AvailabilityFilter) HasMin() bool {
	if o != nil && o.Min != nil {
		return true
	}

	return false
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *AvailabilityFilter) SetMin(v int32) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *AvailabilityFilter) GetMax() int32 {
	if o == nil || o.Max == nil {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityFilter) GetMaxOk() (*int32, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *AvailabilityFilter) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *AvailabilityFilter) SetMax(v int32) {
	o.Max = &v
}

func (o AvailabilityFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedTimeframed, errTimeframed := json.Marshal(o.Timeframed)
	if errTimeframed != nil {
		return []byte{}, errTimeframed
	}
	errTimeframed = json.Unmarshal([]byte(serializedTimeframed), &toSerialize)
	if errTimeframed != nil {
		return []byte{}, errTimeframed
	}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	return json.Marshal(toSerialize)
}

type NullableAvailabilityFilter struct {
	value *AvailabilityFilter
	isSet bool
}

func (v NullableAvailabilityFilter) Get() *AvailabilityFilter {
	return v.value
}

func (v *NullableAvailabilityFilter) Set(val *AvailabilityFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailabilityFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailabilityFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailabilityFilter(val *AvailabilityFilter) *NullableAvailabilityFilter {
	return &NullableAvailabilityFilter{value: val, isSet: true}
}

func (v NullableAvailabilityFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailabilityFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


