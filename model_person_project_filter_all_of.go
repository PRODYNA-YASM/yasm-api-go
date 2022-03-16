/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PersonProjectFilterAllOf struct for PersonProjectFilterAllOf
type PersonProjectFilterAllOf struct {
	ParticipationInMonth *MinMax `json:"participationInMonth,omitempty"`
	InvolevedOfficeIds []string `json:"involevedOfficeIds,omitempty"`
}

// NewPersonProjectFilterAllOf instantiates a new PersonProjectFilterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonProjectFilterAllOf() *PersonProjectFilterAllOf {
	this := PersonProjectFilterAllOf{}
	return &this
}

// NewPersonProjectFilterAllOfWithDefaults instantiates a new PersonProjectFilterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonProjectFilterAllOfWithDefaults() *PersonProjectFilterAllOf {
	this := PersonProjectFilterAllOf{}
	return &this
}

// GetParticipationInMonth returns the ParticipationInMonth field value if set, zero value otherwise.
func (o *PersonProjectFilterAllOf) GetParticipationInMonth() MinMax {
	if o == nil || o.ParticipationInMonth == nil {
		var ret MinMax
		return ret
	}
	return *o.ParticipationInMonth
}

// GetParticipationInMonthOk returns a tuple with the ParticipationInMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonProjectFilterAllOf) GetParticipationInMonthOk() (*MinMax, bool) {
	if o == nil || o.ParticipationInMonth == nil {
		return nil, false
	}
	return o.ParticipationInMonth, true
}

// HasParticipationInMonth returns a boolean if a field has been set.
func (o *PersonProjectFilterAllOf) HasParticipationInMonth() bool {
	if o != nil && o.ParticipationInMonth != nil {
		return true
	}

	return false
}

// SetParticipationInMonth gets a reference to the given MinMax and assigns it to the ParticipationInMonth field.
func (o *PersonProjectFilterAllOf) SetParticipationInMonth(v MinMax) {
	o.ParticipationInMonth = &v
}

// GetInvolevedOfficeIds returns the InvolevedOfficeIds field value if set, zero value otherwise.
func (o *PersonProjectFilterAllOf) GetInvolevedOfficeIds() []string {
	if o == nil || o.InvolevedOfficeIds == nil {
		var ret []string
		return ret
	}
	return o.InvolevedOfficeIds
}

// GetInvolevedOfficeIdsOk returns a tuple with the InvolevedOfficeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonProjectFilterAllOf) GetInvolevedOfficeIdsOk() ([]string, bool) {
	if o == nil || o.InvolevedOfficeIds == nil {
		return nil, false
	}
	return o.InvolevedOfficeIds, true
}

// HasInvolevedOfficeIds returns a boolean if a field has been set.
func (o *PersonProjectFilterAllOf) HasInvolevedOfficeIds() bool {
	if o != nil && o.InvolevedOfficeIds != nil {
		return true
	}

	return false
}

// SetInvolevedOfficeIds gets a reference to the given []string and assigns it to the InvolevedOfficeIds field.
func (o *PersonProjectFilterAllOf) SetInvolevedOfficeIds(v []string) {
	o.InvolevedOfficeIds = v
}

func (o PersonProjectFilterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ParticipationInMonth != nil {
		toSerialize["participationInMonth"] = o.ParticipationInMonth
	}
	if o.InvolevedOfficeIds != nil {
		toSerialize["involevedOfficeIds"] = o.InvolevedOfficeIds
	}
	return json.Marshal(toSerialize)
}

type NullablePersonProjectFilterAllOf struct {
	value *PersonProjectFilterAllOf
	isSet bool
}

func (v NullablePersonProjectFilterAllOf) Get() *PersonProjectFilterAllOf {
	return v.value
}

func (v *NullablePersonProjectFilterAllOf) Set(val *PersonProjectFilterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonProjectFilterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonProjectFilterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonProjectFilterAllOf(val *PersonProjectFilterAllOf) *NullablePersonProjectFilterAllOf {
	return &NullablePersonProjectFilterAllOf{value: val, isSet: true}
}

func (v NullablePersonProjectFilterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonProjectFilterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

