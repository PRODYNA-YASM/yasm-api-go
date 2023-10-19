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

// PersonProjectFilter struct for PersonProjectFilter
type PersonProjectFilter struct {
	AbstractEntityFilter
	ParticipationInMonth *MinMax `json:"participationInMonth,omitempty"`
	InvolvedOfficeIds []string `json:"involvedOfficeIds,omitempty"`
}

// NewPersonProjectFilter instantiates a new PersonProjectFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonProjectFilter(id string) *PersonProjectFilter {
	this := PersonProjectFilter{}
	this.Id = id
	return &this
}

// NewPersonProjectFilterWithDefaults instantiates a new PersonProjectFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonProjectFilterWithDefaults() *PersonProjectFilter {
	this := PersonProjectFilter{}
	return &this
}

// GetParticipationInMonth returns the ParticipationInMonth field value if set, zero value otherwise.
func (o *PersonProjectFilter) GetParticipationInMonth() MinMax {
	if o == nil || o.ParticipationInMonth == nil {
		var ret MinMax
		return ret
	}
	return *o.ParticipationInMonth
}

// GetParticipationInMonthOk returns a tuple with the ParticipationInMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonProjectFilter) GetParticipationInMonthOk() (*MinMax, bool) {
	if o == nil || o.ParticipationInMonth == nil {
		return nil, false
	}
	return o.ParticipationInMonth, true
}

// HasParticipationInMonth returns a boolean if a field has been set.
func (o *PersonProjectFilter) HasParticipationInMonth() bool {
	if o != nil && o.ParticipationInMonth != nil {
		return true
	}

	return false
}

// SetParticipationInMonth gets a reference to the given MinMax and assigns it to the ParticipationInMonth field.
func (o *PersonProjectFilter) SetParticipationInMonth(v MinMax) {
	o.ParticipationInMonth = &v
}

// GetInvolvedOfficeIds returns the InvolvedOfficeIds field value if set, zero value otherwise.
func (o *PersonProjectFilter) GetInvolvedOfficeIds() []string {
	if o == nil || o.InvolvedOfficeIds == nil {
		var ret []string
		return ret
	}
	return o.InvolvedOfficeIds
}

// GetInvolvedOfficeIdsOk returns a tuple with the InvolvedOfficeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonProjectFilter) GetInvolvedOfficeIdsOk() ([]string, bool) {
	if o == nil || o.InvolvedOfficeIds == nil {
		return nil, false
	}
	return o.InvolvedOfficeIds, true
}

// HasInvolvedOfficeIds returns a boolean if a field has been set.
func (o *PersonProjectFilter) HasInvolvedOfficeIds() bool {
	if o != nil && o.InvolvedOfficeIds != nil {
		return true
	}

	return false
}

// SetInvolvedOfficeIds gets a reference to the given []string and assigns it to the InvolvedOfficeIds field.
func (o *PersonProjectFilter) SetInvolvedOfficeIds(v []string) {
	o.InvolvedOfficeIds = v
}

func (o PersonProjectFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAbstractEntityFilter, errAbstractEntityFilter := json.Marshal(o.AbstractEntityFilter)
	if errAbstractEntityFilter != nil {
		return []byte{}, errAbstractEntityFilter
	}
	errAbstractEntityFilter = json.Unmarshal([]byte(serializedAbstractEntityFilter), &toSerialize)
	if errAbstractEntityFilter != nil {
		return []byte{}, errAbstractEntityFilter
	}
	if o.ParticipationInMonth != nil {
		toSerialize["participationInMonth"] = o.ParticipationInMonth
	}
	if o.InvolvedOfficeIds != nil {
		toSerialize["involvedOfficeIds"] = o.InvolvedOfficeIds
	}
	return json.Marshal(toSerialize)
}

type NullablePersonProjectFilter struct {
	value *PersonProjectFilter
	isSet bool
}

func (v NullablePersonProjectFilter) Get() *PersonProjectFilter {
	return v.value
}

func (v *NullablePersonProjectFilter) Set(val *PersonProjectFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonProjectFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonProjectFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonProjectFilter(val *PersonProjectFilter) *NullablePersonProjectFilter {
	return &NullablePersonProjectFilter{value: val, isSet: true}
}

func (v NullablePersonProjectFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonProjectFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


