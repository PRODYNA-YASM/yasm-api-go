/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 0.7.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Experience struct for Experience
type Experience struct {
	Skill *SkillLevel `json:"skill,omitempty"`
	ConfirmedBy *[]Person `json:"confirmedBy,omitempty"`
}

// NewExperience instantiates a new Experience object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExperience() *Experience {
	this := Experience{}
	return &this
}

// NewExperienceWithDefaults instantiates a new Experience object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExperienceWithDefaults() *Experience {
	this := Experience{}
	return &this
}

// GetSkill returns the Skill field value if set, zero value otherwise.
func (o *Experience) GetSkill() SkillLevel {
	if o == nil || o.Skill == nil {
		var ret SkillLevel
		return ret
	}
	return *o.Skill
}

// GetSkillOk returns a tuple with the Skill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Experience) GetSkillOk() (*SkillLevel, bool) {
	if o == nil || o.Skill == nil {
		return nil, false
	}
	return o.Skill, true
}

// HasSkill returns a boolean if a field has been set.
func (o *Experience) HasSkill() bool {
	if o != nil && o.Skill != nil {
		return true
	}

	return false
}

// SetSkill gets a reference to the given SkillLevel and assigns it to the Skill field.
func (o *Experience) SetSkill(v SkillLevel) {
	o.Skill = &v
}

// GetConfirmedBy returns the ConfirmedBy field value if set, zero value otherwise.
func (o *Experience) GetConfirmedBy() []Person {
	if o == nil || o.ConfirmedBy == nil {
		var ret []Person
		return ret
	}
	return *o.ConfirmedBy
}

// GetConfirmedByOk returns a tuple with the ConfirmedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Experience) GetConfirmedByOk() (*[]Person, bool) {
	if o == nil || o.ConfirmedBy == nil {
		return nil, false
	}
	return o.ConfirmedBy, true
}

// HasConfirmedBy returns a boolean if a field has been set.
func (o *Experience) HasConfirmedBy() bool {
	if o != nil && o.ConfirmedBy != nil {
		return true
	}

	return false
}

// SetConfirmedBy gets a reference to the given []Person and assigns it to the ConfirmedBy field.
func (o *Experience) SetConfirmedBy(v []Person) {
	o.ConfirmedBy = &v
}

func (o Experience) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Skill != nil {
		toSerialize["skill"] = o.Skill
	}
	if o.ConfirmedBy != nil {
		toSerialize["confirmedBy"] = o.ConfirmedBy
	}
	return json.Marshal(toSerialize)
}

type NullableExperience struct {
	value *Experience
	isSet bool
}

func (v NullableExperience) Get() *Experience {
	return v.value
}

func (v *NullableExperience) Set(val *Experience) {
	v.value = val
	v.isSet = true
}

func (v NullableExperience) IsSet() bool {
	return v.isSet
}

func (v *NullableExperience) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExperience(val *Experience) *NullableExperience {
	return &NullableExperience{value: val, isSet: true}
}

func (v NullableExperience) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExperience) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


