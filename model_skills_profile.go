/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the SkillsProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkillsProfile{}

// SkillsProfile struct for SkillsProfile
type SkillsProfile struct {
	NamedDomainModel
}

// NewSkillsProfile instantiates a new SkillsProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkillsProfile(id string, name string) *SkillsProfile {
	this := SkillsProfile{}
	this.Id = id
	this.Name = name
	return &this
}

// NewSkillsProfileWithDefaults instantiates a new SkillsProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkillsProfileWithDefaults() *SkillsProfile {
	this := SkillsProfile{}
	return &this
}

func (o SkillsProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkillsProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedNamedDomainModel, errNamedDomainModel := json.Marshal(o.NamedDomainModel)
	if errNamedDomainModel != nil {
		return map[string]interface{}{}, errNamedDomainModel
	}
	errNamedDomainModel = json.Unmarshal([]byte(serializedNamedDomainModel), &toSerialize)
	if errNamedDomainModel != nil {
		return map[string]interface{}{}, errNamedDomainModel
	}
	return toSerialize, nil
}

type NullableSkillsProfile struct {
	value *SkillsProfile
	isSet bool
}

func (v NullableSkillsProfile) Get() *SkillsProfile {
	return v.value
}

func (v *NullableSkillsProfile) Set(val *SkillsProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableSkillsProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableSkillsProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkillsProfile(val *SkillsProfile) *NullableSkillsProfile {
	return &NullableSkillsProfile{value: val, isSet: true}
}

func (v NullableSkillsProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkillsProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

