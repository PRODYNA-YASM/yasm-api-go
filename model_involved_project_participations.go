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

// checks if the InvolvedProjectParticipations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvolvedProjectParticipations{}

// InvolvedProjectParticipations struct for InvolvedProjectParticipations
type InvolvedProjectParticipations struct {
	ProjectParticipationIds []string `json:"projectParticipationIds,omitempty"`
}

// NewInvolvedProjectParticipations instantiates a new InvolvedProjectParticipations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvolvedProjectParticipations() *InvolvedProjectParticipations {
	this := InvolvedProjectParticipations{}
	return &this
}

// NewInvolvedProjectParticipationsWithDefaults instantiates a new InvolvedProjectParticipations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvolvedProjectParticipationsWithDefaults() *InvolvedProjectParticipations {
	this := InvolvedProjectParticipations{}
	return &this
}

// GetProjectParticipationIds returns the ProjectParticipationIds field value if set, zero value otherwise.
func (o *InvolvedProjectParticipations) GetProjectParticipationIds() []string {
	if o == nil || IsNil(o.ProjectParticipationIds) {
		var ret []string
		return ret
	}
	return o.ProjectParticipationIds
}

// GetProjectParticipationIdsOk returns a tuple with the ProjectParticipationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvolvedProjectParticipations) GetProjectParticipationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProjectParticipationIds) {
		return nil, false
	}
	return o.ProjectParticipationIds, true
}

// HasProjectParticipationIds returns a boolean if a field has been set.
func (o *InvolvedProjectParticipations) HasProjectParticipationIds() bool {
	if o != nil && !IsNil(o.ProjectParticipationIds) {
		return true
	}

	return false
}

// SetProjectParticipationIds gets a reference to the given []string and assigns it to the ProjectParticipationIds field.
func (o *InvolvedProjectParticipations) SetProjectParticipationIds(v []string) {
	o.ProjectParticipationIds = v
}

func (o InvolvedProjectParticipations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvolvedProjectParticipations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectParticipationIds) {
		toSerialize["projectParticipationIds"] = o.ProjectParticipationIds
	}
	return toSerialize, nil
}

type NullableInvolvedProjectParticipations struct {
	value *InvolvedProjectParticipations
	isSet bool
}

func (v NullableInvolvedProjectParticipations) Get() *InvolvedProjectParticipations {
	return v.value
}

func (v *NullableInvolvedProjectParticipations) Set(val *InvolvedProjectParticipations) {
	v.value = val
	v.isSet = true
}

func (v NullableInvolvedProjectParticipations) IsSet() bool {
	return v.isSet
}

func (v *NullableInvolvedProjectParticipations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvolvedProjectParticipations(val *InvolvedProjectParticipations) *NullableInvolvedProjectParticipations {
	return &NullableInvolvedProjectParticipations{value: val, isSet: true}
}

func (v NullableInvolvedProjectParticipations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvolvedProjectParticipations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


