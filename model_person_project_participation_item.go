/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PersonProjectParticipationItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonProjectParticipationItem{}

// PersonProjectParticipationItem struct for PersonProjectParticipationItem
type PersonProjectParticipationItem struct {
	Participation *ProjectParticipation `json:"participation,omitempty"`
	SkillGroups []ExperienceSkillGroup `json:"skillGroups,omitempty"`
}

// NewPersonProjectParticipationItem instantiates a new PersonProjectParticipationItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonProjectParticipationItem() *PersonProjectParticipationItem {
	this := PersonProjectParticipationItem{}
	return &this
}

// NewPersonProjectParticipationItemWithDefaults instantiates a new PersonProjectParticipationItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonProjectParticipationItemWithDefaults() *PersonProjectParticipationItem {
	this := PersonProjectParticipationItem{}
	return &this
}

// GetParticipation returns the Participation field value if set, zero value otherwise.
func (o *PersonProjectParticipationItem) GetParticipation() ProjectParticipation {
	if o == nil || IsNil(o.Participation) {
		var ret ProjectParticipation
		return ret
	}
	return *o.Participation
}

// GetParticipationOk returns a tuple with the Participation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonProjectParticipationItem) GetParticipationOk() (*ProjectParticipation, bool) {
	if o == nil || IsNil(o.Participation) {
		return nil, false
	}
	return o.Participation, true
}

// HasParticipation returns a boolean if a field has been set.
func (o *PersonProjectParticipationItem) HasParticipation() bool {
	if o != nil && !IsNil(o.Participation) {
		return true
	}

	return false
}

// SetParticipation gets a reference to the given ProjectParticipation and assigns it to the Participation field.
func (o *PersonProjectParticipationItem) SetParticipation(v ProjectParticipation) {
	o.Participation = &v
}

// GetSkillGroups returns the SkillGroups field value if set, zero value otherwise.
func (o *PersonProjectParticipationItem) GetSkillGroups() []ExperienceSkillGroup {
	if o == nil || IsNil(o.SkillGroups) {
		var ret []ExperienceSkillGroup
		return ret
	}
	return o.SkillGroups
}

// GetSkillGroupsOk returns a tuple with the SkillGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonProjectParticipationItem) GetSkillGroupsOk() ([]ExperienceSkillGroup, bool) {
	if o == nil || IsNil(o.SkillGroups) {
		return nil, false
	}
	return o.SkillGroups, true
}

// HasSkillGroups returns a boolean if a field has been set.
func (o *PersonProjectParticipationItem) HasSkillGroups() bool {
	if o != nil && !IsNil(o.SkillGroups) {
		return true
	}

	return false
}

// SetSkillGroups gets a reference to the given []ExperienceSkillGroup and assigns it to the SkillGroups field.
func (o *PersonProjectParticipationItem) SetSkillGroups(v []ExperienceSkillGroup) {
	o.SkillGroups = v
}

func (o PersonProjectParticipationItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonProjectParticipationItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Participation) {
		toSerialize["participation"] = o.Participation
	}
	if !IsNil(o.SkillGroups) {
		toSerialize["skillGroups"] = o.SkillGroups
	}
	return toSerialize, nil
}

type NullablePersonProjectParticipationItem struct {
	value *PersonProjectParticipationItem
	isSet bool
}

func (v NullablePersonProjectParticipationItem) Get() *PersonProjectParticipationItem {
	return v.value
}

func (v *NullablePersonProjectParticipationItem) Set(val *PersonProjectParticipationItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonProjectParticipationItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonProjectParticipationItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonProjectParticipationItem(val *PersonProjectParticipationItem) *NullablePersonProjectParticipationItem {
	return &NullablePersonProjectParticipationItem{value: val, isSet: true}
}

func (v NullablePersonProjectParticipationItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonProjectParticipationItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


