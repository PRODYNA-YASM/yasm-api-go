/*
 * YASM (Yet Another Skill Management) API
 *
 * This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.
 *
 * API version: 0.5.9
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PagedSkillsAllOf struct for PagedSkillsAllOf
type PagedSkillsAllOf struct {
	Skills *[]Skill `json:"skills,omitempty"`
}

// NewPagedSkillsAllOf instantiates a new PagedSkillsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedSkillsAllOf() *PagedSkillsAllOf {
	this := PagedSkillsAllOf{}
	return &this
}

// NewPagedSkillsAllOfWithDefaults instantiates a new PagedSkillsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedSkillsAllOfWithDefaults() *PagedSkillsAllOf {
	this := PagedSkillsAllOf{}
	return &this
}

// GetSkills returns the Skills field value if set, zero value otherwise.
func (o *PagedSkillsAllOf) GetSkills() []Skill {
	if o == nil || o.Skills == nil {
		var ret []Skill
		return ret
	}
	return *o.Skills
}

// GetSkillsOk returns a tuple with the Skills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedSkillsAllOf) GetSkillsOk() (*[]Skill, bool) {
	if o == nil || o.Skills == nil {
		return nil, false
	}
	return o.Skills, true
}

// HasSkills returns a boolean if a field has been set.
func (o *PagedSkillsAllOf) HasSkills() bool {
	if o != nil && o.Skills != nil {
		return true
	}

	return false
}

// SetSkills gets a reference to the given []Skill and assigns it to the Skills field.
func (o *PagedSkillsAllOf) SetSkills(v []Skill) {
	o.Skills = &v
}

func (o PagedSkillsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Skills != nil {
		toSerialize["skills"] = o.Skills
	}
	return json.Marshal(toSerialize)
}

type NullablePagedSkillsAllOf struct {
	value *PagedSkillsAllOf
	isSet bool
}

func (v NullablePagedSkillsAllOf) Get() *PagedSkillsAllOf {
	return v.value
}

func (v *NullablePagedSkillsAllOf) Set(val *PagedSkillsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedSkillsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedSkillsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedSkillsAllOf(val *PagedSkillsAllOf) *NullablePagedSkillsAllOf {
	return &NullablePagedSkillsAllOf{value: val, isSet: true}
}

func (v NullablePagedSkillsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedSkillsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


