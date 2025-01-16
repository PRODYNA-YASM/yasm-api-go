/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.51.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ExperienceSkill type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExperienceSkill{}

// ExperienceSkill struct for ExperienceSkill
type ExperienceSkill struct {
	Skill
	ExperienceInMonth *int32 `json:"experienceInMonth,omitempty"`
}

// NewExperienceSkill instantiates a new ExperienceSkill object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExperienceSkill(id string, name string) *ExperienceSkill {
	this := ExperienceSkill{}
	this.Id = id
	this.Name = name
	var important bool = false
	this.Important = &important
	var experienceInMonth int32 = 0
	this.ExperienceInMonth = &experienceInMonth
	return &this
}

// NewExperienceSkillWithDefaults instantiates a new ExperienceSkill object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExperienceSkillWithDefaults() *ExperienceSkill {
	this := ExperienceSkill{}
	var experienceInMonth int32 = 0
	this.ExperienceInMonth = &experienceInMonth
	return &this
}

// GetExperienceInMonth returns the ExperienceInMonth field value if set, zero value otherwise.
func (o *ExperienceSkill) GetExperienceInMonth() int32 {
	if o == nil || IsNil(o.ExperienceInMonth) {
		var ret int32
		return ret
	}
	return *o.ExperienceInMonth
}

// GetExperienceInMonthOk returns a tuple with the ExperienceInMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperienceSkill) GetExperienceInMonthOk() (*int32, bool) {
	if o == nil || IsNil(o.ExperienceInMonth) {
		return nil, false
	}
	return o.ExperienceInMonth, true
}

// HasExperienceInMonth returns a boolean if a field has been set.
func (o *ExperienceSkill) HasExperienceInMonth() bool {
	if o != nil && !IsNil(o.ExperienceInMonth) {
		return true
	}

	return false
}

// SetExperienceInMonth gets a reference to the given int32 and assigns it to the ExperienceInMonth field.
func (o *ExperienceSkill) SetExperienceInMonth(v int32) {
	o.ExperienceInMonth = &v
}

func (o ExperienceSkill) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExperienceSkill) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSkill, errSkill := json.Marshal(o.Skill)
	if errSkill != nil {
		return map[string]interface{}{}, errSkill
	}
	errSkill = json.Unmarshal([]byte(serializedSkill), &toSerialize)
	if errSkill != nil {
		return map[string]interface{}{}, errSkill
	}
	if !IsNil(o.ExperienceInMonth) {
		toSerialize["experienceInMonth"] = o.ExperienceInMonth
	}
	return toSerialize, nil
}

type NullableExperienceSkill struct {
	value *ExperienceSkill
	isSet bool
}

func (v NullableExperienceSkill) Get() *ExperienceSkill {
	return v.value
}

func (v *NullableExperienceSkill) Set(val *ExperienceSkill) {
	v.value = val
	v.isSet = true
}

func (v NullableExperienceSkill) IsSet() bool {
	return v.isSet
}

func (v *NullableExperienceSkill) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExperienceSkill(val *ExperienceSkill) *NullableExperienceSkill {
	return &NullableExperienceSkill{value: val, isSet: true}
}

func (v NullableExperienceSkill) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExperienceSkill) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


