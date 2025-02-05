/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.60.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the SkillStatisticProject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkillStatisticProject{}

// SkillStatisticProject struct for SkillStatisticProject
type SkillStatisticProject struct {
	ProjectView
	Duration *string `json:"duration,omitempty"`
}

// NewSkillStatisticProject instantiates a new SkillStatisticProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkillStatisticProject() *SkillStatisticProject {
	this := SkillStatisticProject{}
	return &this
}

// NewSkillStatisticProjectWithDefaults instantiates a new SkillStatisticProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkillStatisticProjectWithDefaults() *SkillStatisticProject {
	this := SkillStatisticProject{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *SkillStatisticProject) GetDuration() string {
	if o == nil || IsNil(o.Duration) {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillStatisticProject) GetDurationOk() (*string, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *SkillStatisticProject) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *SkillStatisticProject) SetDuration(v string) {
	o.Duration = &v
}

func (o SkillStatisticProject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkillStatisticProject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedProjectView, errProjectView := json.Marshal(o.ProjectView)
	if errProjectView != nil {
		return map[string]interface{}{}, errProjectView
	}
	errProjectView = json.Unmarshal([]byte(serializedProjectView), &toSerialize)
	if errProjectView != nil {
		return map[string]interface{}{}, errProjectView
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	return toSerialize, nil
}

type NullableSkillStatisticProject struct {
	value *SkillStatisticProject
	isSet bool
}

func (v NullableSkillStatisticProject) Get() *SkillStatisticProject {
	return v.value
}

func (v *NullableSkillStatisticProject) Set(val *SkillStatisticProject) {
	v.value = val
	v.isSet = true
}

func (v NullableSkillStatisticProject) IsSet() bool {
	return v.isSet
}

func (v *NullableSkillStatisticProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkillStatisticProject(val *SkillStatisticProject) *NullableSkillStatisticProject {
	return &NullableSkillStatisticProject{value: val, isSet: true}
}

func (v NullableSkillStatisticProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkillStatisticProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


