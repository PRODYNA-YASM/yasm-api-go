/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.74.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PagedSkills type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagedSkills{}

// PagedSkills struct for PagedSkills
type PagedSkills struct {
	Page
	Skills []SkillDetails `json:"skills,omitempty"`
}

// NewPagedSkills instantiates a new PagedSkills object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedSkills() *PagedSkills {
	this := PagedSkills{}
	return &this
}

// NewPagedSkillsWithDefaults instantiates a new PagedSkills object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedSkillsWithDefaults() *PagedSkills {
	this := PagedSkills{}
	return &this
}

// GetSkills returns the Skills field value if set, zero value otherwise.
func (o *PagedSkills) GetSkills() []SkillDetails {
	if o == nil || IsNil(o.Skills) {
		var ret []SkillDetails
		return ret
	}
	return o.Skills
}

// GetSkillsOk returns a tuple with the Skills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedSkills) GetSkillsOk() ([]SkillDetails, bool) {
	if o == nil || IsNil(o.Skills) {
		return nil, false
	}
	return o.Skills, true
}

// HasSkills returns a boolean if a field has been set.
func (o *PagedSkills) HasSkills() bool {
	if o != nil && !IsNil(o.Skills) {
		return true
	}

	return false
}

// SetSkills gets a reference to the given []SkillDetails and assigns it to the Skills field.
func (o *PagedSkills) SetSkills(v []SkillDetails) {
	o.Skills = v
}

func (o PagedSkills) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagedSkills) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPage, errPage := json.Marshal(o.Page)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	errPage = json.Unmarshal([]byte(serializedPage), &toSerialize)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	if !IsNil(o.Skills) {
		toSerialize["skills"] = o.Skills
	}
	return toSerialize, nil
}

type NullablePagedSkills struct {
	value *PagedSkills
	isSet bool
}

func (v NullablePagedSkills) Get() *PagedSkills {
	return v.value
}

func (v *NullablePagedSkills) Set(val *PagedSkills) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedSkills) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedSkills) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedSkills(val *PagedSkills) *NullablePagedSkills {
	return &NullablePagedSkills{value: val, isSet: true}
}

func (v NullablePagedSkills) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedSkills) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


