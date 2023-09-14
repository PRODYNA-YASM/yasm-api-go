/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.4.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// AvailabilityDetail struct for AvailabilityDetail
type AvailabilityDetail struct {
	Availability *Availability `json:"availability,omitempty"`
	Percent *float32 `json:"percent,omitempty"`
}

// NewAvailabilityDetail instantiates a new AvailabilityDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailabilityDetail() *AvailabilityDetail {
	this := AvailabilityDetail{}
	return &this
}

// NewAvailabilityDetailWithDefaults instantiates a new AvailabilityDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailabilityDetailWithDefaults() *AvailabilityDetail {
	this := AvailabilityDetail{}
	return &this
}

// GetAvailability returns the Availability field value if set, zero value otherwise.
func (o *AvailabilityDetail) GetAvailability() Availability {
	if o == nil || o.Availability == nil {
		var ret Availability
		return ret
	}
	return *o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityDetail) GetAvailabilityOk() (*Availability, bool) {
	if o == nil || o.Availability == nil {
		return nil, false
	}
	return o.Availability, true
}

// HasAvailability returns a boolean if a field has been set.
func (o *AvailabilityDetail) HasAvailability() bool {
	if o != nil && o.Availability != nil {
		return true
	}

	return false
}

// SetAvailability gets a reference to the given Availability and assigns it to the Availability field.
func (o *AvailabilityDetail) SetAvailability(v Availability) {
	o.Availability = &v
}

// GetPercent returns the Percent field value if set, zero value otherwise.
func (o *AvailabilityDetail) GetPercent() float32 {
	if o == nil || o.Percent == nil {
		var ret float32
		return ret
	}
	return *o.Percent
}

// GetPercentOk returns a tuple with the Percent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityDetail) GetPercentOk() (*float32, bool) {
	if o == nil || o.Percent == nil {
		return nil, false
	}
	return o.Percent, true
}

// HasPercent returns a boolean if a field has been set.
func (o *AvailabilityDetail) HasPercent() bool {
	if o != nil && o.Percent != nil {
		return true
	}

	return false
}

// SetPercent gets a reference to the given float32 and assigns it to the Percent field.
func (o *AvailabilityDetail) SetPercent(v float32) {
	o.Percent = &v
}

func (o AvailabilityDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Availability != nil {
		toSerialize["availability"] = o.Availability
	}
	if o.Percent != nil {
		toSerialize["percent"] = o.Percent
	}
	return json.Marshal(toSerialize)
}

type NullableAvailabilityDetail struct {
	value *AvailabilityDetail
	isSet bool
}

func (v NullableAvailabilityDetail) Get() *AvailabilityDetail {
	return v.value
}

func (v *NullableAvailabilityDetail) Set(val *AvailabilityDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailabilityDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailabilityDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailabilityDetail(val *AvailabilityDetail) *NullableAvailabilityDetail {
	return &NullableAvailabilityDetail{value: val, isSet: true}
}

func (v NullableAvailabilityDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailabilityDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


