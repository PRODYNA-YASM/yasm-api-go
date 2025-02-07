/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.62.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the AvailabilityDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailabilityDetail{}

// AvailabilityDetail struct for AvailabilityDetail
type AvailabilityDetail struct {
	Audit *Audit `json:"audit,omitempty"`
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

// GetAudit returns the Audit field value if set, zero value otherwise.
func (o *AvailabilityDetail) GetAudit() Audit {
	if o == nil || IsNil(o.Audit) {
		var ret Audit
		return ret
	}
	return *o.Audit
}

// GetAuditOk returns a tuple with the Audit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityDetail) GetAuditOk() (*Audit, bool) {
	if o == nil || IsNil(o.Audit) {
		return nil, false
	}
	return o.Audit, true
}

// HasAudit returns a boolean if a field has been set.
func (o *AvailabilityDetail) HasAudit() bool {
	if o != nil && !IsNil(o.Audit) {
		return true
	}

	return false
}

// SetAudit gets a reference to the given Audit and assigns it to the Audit field.
func (o *AvailabilityDetail) SetAudit(v Audit) {
	o.Audit = &v
}

// GetAvailability returns the Availability field value if set, zero value otherwise.
func (o *AvailabilityDetail) GetAvailability() Availability {
	if o == nil || IsNil(o.Availability) {
		var ret Availability
		return ret
	}
	return *o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityDetail) GetAvailabilityOk() (*Availability, bool) {
	if o == nil || IsNil(o.Availability) {
		return nil, false
	}
	return o.Availability, true
}

// HasAvailability returns a boolean if a field has been set.
func (o *AvailabilityDetail) HasAvailability() bool {
	if o != nil && !IsNil(o.Availability) {
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
	if o == nil || IsNil(o.Percent) {
		var ret float32
		return ret
	}
	return *o.Percent
}

// GetPercentOk returns a tuple with the Percent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityDetail) GetPercentOk() (*float32, bool) {
	if o == nil || IsNil(o.Percent) {
		return nil, false
	}
	return o.Percent, true
}

// HasPercent returns a boolean if a field has been set.
func (o *AvailabilityDetail) HasPercent() bool {
	if o != nil && !IsNil(o.Percent) {
		return true
	}

	return false
}

// SetPercent gets a reference to the given float32 and assigns it to the Percent field.
func (o *AvailabilityDetail) SetPercent(v float32) {
	o.Percent = &v
}

func (o AvailabilityDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvailabilityDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Audit) {
		toSerialize["audit"] = o.Audit
	}
	if !IsNil(o.Availability) {
		toSerialize["availability"] = o.Availability
	}
	if !IsNil(o.Percent) {
		toSerialize["percent"] = o.Percent
	}
	return toSerialize, nil
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


