/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.4.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ExperienceSkillAllOf struct for ExperienceSkillAllOf
type ExperienceSkillAllOf struct {
	ExperienceInMonth *int32 `json:"experienceInMonth,omitempty"`
}

// NewExperienceSkillAllOf instantiates a new ExperienceSkillAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExperienceSkillAllOf() *ExperienceSkillAllOf {
	this := ExperienceSkillAllOf{}
	var experienceInMonth int32 = 0
	this.ExperienceInMonth = &experienceInMonth
	return &this
}

// NewExperienceSkillAllOfWithDefaults instantiates a new ExperienceSkillAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExperienceSkillAllOfWithDefaults() *ExperienceSkillAllOf {
	this := ExperienceSkillAllOf{}
	var experienceInMonth int32 = 0
	this.ExperienceInMonth = &experienceInMonth
	return &this
}

// GetExperienceInMonth returns the ExperienceInMonth field value if set, zero value otherwise.
func (o *ExperienceSkillAllOf) GetExperienceInMonth() int32 {
	if o == nil || o.ExperienceInMonth == nil {
		var ret int32
		return ret
	}
	return *o.ExperienceInMonth
}

// GetExperienceInMonthOk returns a tuple with the ExperienceInMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperienceSkillAllOf) GetExperienceInMonthOk() (*int32, bool) {
	if o == nil || o.ExperienceInMonth == nil {
		return nil, false
	}
	return o.ExperienceInMonth, true
}

// HasExperienceInMonth returns a boolean if a field has been set.
func (o *ExperienceSkillAllOf) HasExperienceInMonth() bool {
	if o != nil && o.ExperienceInMonth != nil {
		return true
	}

	return false
}

// SetExperienceInMonth gets a reference to the given int32 and assigns it to the ExperienceInMonth field.
func (o *ExperienceSkillAllOf) SetExperienceInMonth(v int32) {
	o.ExperienceInMonth = &v
}

func (o ExperienceSkillAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExperienceInMonth != nil {
		toSerialize["experienceInMonth"] = o.ExperienceInMonth
	}
	return json.Marshal(toSerialize)
}

type NullableExperienceSkillAllOf struct {
	value *ExperienceSkillAllOf
	isSet bool
}

func (v NullableExperienceSkillAllOf) Get() *ExperienceSkillAllOf {
	return v.value
}

func (v *NullableExperienceSkillAllOf) Set(val *ExperienceSkillAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExperienceSkillAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExperienceSkillAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExperienceSkillAllOf(val *ExperienceSkillAllOf) *NullableExperienceSkillAllOf {
	return &NullableExperienceSkillAllOf{value: val, isSet: true}
}

func (v NullableExperienceSkillAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExperienceSkillAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


