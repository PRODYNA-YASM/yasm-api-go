/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PersonSearchSkillsInner struct for PersonSearchSkillsInner
type PersonSearchSkillsInner struct {
	Id string `json:"id"`
	ExperienceInMonth *MinMax `json:"experienceInMonth,omitempty"`
	// filters the last time the employee used the skill in a project
	LastAssignment *string `json:"lastAssignment,omitempty"`
}

// NewPersonSearchSkillsInner instantiates a new PersonSearchSkillsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonSearchSkillsInner(id string) *PersonSearchSkillsInner {
	this := PersonSearchSkillsInner{}
	this.Id = id
	return &this
}

// NewPersonSearchSkillsInnerWithDefaults instantiates a new PersonSearchSkillsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonSearchSkillsInnerWithDefaults() *PersonSearchSkillsInner {
	this := PersonSearchSkillsInner{}
	return &this
}

// GetId returns the Id field value
func (o *PersonSearchSkillsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PersonSearchSkillsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PersonSearchSkillsInner) SetId(v string) {
	o.Id = v
}

// GetExperienceInMonth returns the ExperienceInMonth field value if set, zero value otherwise.
func (o *PersonSearchSkillsInner) GetExperienceInMonth() MinMax {
	if o == nil || o.ExperienceInMonth == nil {
		var ret MinMax
		return ret
	}
	return *o.ExperienceInMonth
}

// GetExperienceInMonthOk returns a tuple with the ExperienceInMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearchSkillsInner) GetExperienceInMonthOk() (*MinMax, bool) {
	if o == nil || o.ExperienceInMonth == nil {
		return nil, false
	}
	return o.ExperienceInMonth, true
}

// HasExperienceInMonth returns a boolean if a field has been set.
func (o *PersonSearchSkillsInner) HasExperienceInMonth() bool {
	if o != nil && o.ExperienceInMonth != nil {
		return true
	}

	return false
}

// SetExperienceInMonth gets a reference to the given MinMax and assigns it to the ExperienceInMonth field.
func (o *PersonSearchSkillsInner) SetExperienceInMonth(v MinMax) {
	o.ExperienceInMonth = &v
}

// GetLastAssignment returns the LastAssignment field value if set, zero value otherwise.
func (o *PersonSearchSkillsInner) GetLastAssignment() string {
	if o == nil || o.LastAssignment == nil {
		var ret string
		return ret
	}
	return *o.LastAssignment
}

// GetLastAssignmentOk returns a tuple with the LastAssignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearchSkillsInner) GetLastAssignmentOk() (*string, bool) {
	if o == nil || o.LastAssignment == nil {
		return nil, false
	}
	return o.LastAssignment, true
}

// HasLastAssignment returns a boolean if a field has been set.
func (o *PersonSearchSkillsInner) HasLastAssignment() bool {
	if o != nil && o.LastAssignment != nil {
		return true
	}

	return false
}

// SetLastAssignment gets a reference to the given string and assigns it to the LastAssignment field.
func (o *PersonSearchSkillsInner) SetLastAssignment(v string) {
	o.LastAssignment = &v
}

func (o PersonSearchSkillsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.ExperienceInMonth != nil {
		toSerialize["experienceInMonth"] = o.ExperienceInMonth
	}
	if o.LastAssignment != nil {
		toSerialize["lastAssignment"] = o.LastAssignment
	}
	return json.Marshal(toSerialize)
}

type NullablePersonSearchSkillsInner struct {
	value *PersonSearchSkillsInner
	isSet bool
}

func (v NullablePersonSearchSkillsInner) Get() *PersonSearchSkillsInner {
	return v.value
}

func (v *NullablePersonSearchSkillsInner) Set(val *PersonSearchSkillsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonSearchSkillsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonSearchSkillsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonSearchSkillsInner(val *PersonSearchSkillsInner) *NullablePersonSearchSkillsInner {
	return &NullablePersonSearchSkillsInner{value: val, isSet: true}
}

func (v NullablePersonSearchSkillsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonSearchSkillsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


