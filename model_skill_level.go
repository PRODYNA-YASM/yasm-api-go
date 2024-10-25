/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the SkillLevel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkillLevel{}

// SkillLevel struct for SkillLevel
type SkillLevel struct {
	Skill *Skill `json:"skill,omitempty"`
	Level *Level `json:"level,omitempty"`
}

// NewSkillLevel instantiates a new SkillLevel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkillLevel() *SkillLevel {
	this := SkillLevel{}
	return &this
}

// NewSkillLevelWithDefaults instantiates a new SkillLevel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkillLevelWithDefaults() *SkillLevel {
	this := SkillLevel{}
	return &this
}

// GetSkill returns the Skill field value if set, zero value otherwise.
func (o *SkillLevel) GetSkill() Skill {
	if o == nil || IsNil(o.Skill) {
		var ret Skill
		return ret
	}
	return *o.Skill
}

// GetSkillOk returns a tuple with the Skill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillLevel) GetSkillOk() (*Skill, bool) {
	if o == nil || IsNil(o.Skill) {
		return nil, false
	}
	return o.Skill, true
}

// HasSkill returns a boolean if a field has been set.
func (o *SkillLevel) HasSkill() bool {
	if o != nil && !IsNil(o.Skill) {
		return true
	}

	return false
}

// SetSkill gets a reference to the given Skill and assigns it to the Skill field.
func (o *SkillLevel) SetSkill(v Skill) {
	o.Skill = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *SkillLevel) GetLevel() Level {
	if o == nil || IsNil(o.Level) {
		var ret Level
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillLevel) GetLevelOk() (*Level, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *SkillLevel) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given Level and assigns it to the Level field.
func (o *SkillLevel) SetLevel(v Level) {
	o.Level = &v
}

func (o SkillLevel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkillLevel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Skill) {
		toSerialize["skill"] = o.Skill
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	return toSerialize, nil
}

type NullableSkillLevel struct {
	value *SkillLevel
	isSet bool
}

func (v NullableSkillLevel) Get() *SkillLevel {
	return v.value
}

func (v *NullableSkillLevel) Set(val *SkillLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableSkillLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableSkillLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkillLevel(val *SkillLevel) *NullableSkillLevel {
	return &NullableSkillLevel{value: val, isSet: true}
}

func (v NullableSkillLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkillLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


