/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.4.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PagedProjectParticipations struct for PagedProjectParticipations
type PagedProjectParticipations struct {
	Page
	ProjectParticipations []ProjectParticipation `json:"projectParticipations,omitempty"`
}

// NewPagedProjectParticipations instantiates a new PagedProjectParticipations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedProjectParticipations() *PagedProjectParticipations {
	this := PagedProjectParticipations{}
	return &this
}

// NewPagedProjectParticipationsWithDefaults instantiates a new PagedProjectParticipations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedProjectParticipationsWithDefaults() *PagedProjectParticipations {
	this := PagedProjectParticipations{}
	return &this
}

// GetProjectParticipations returns the ProjectParticipations field value if set, zero value otherwise.
func (o *PagedProjectParticipations) GetProjectParticipations() []ProjectParticipation {
	if o == nil || o.ProjectParticipations == nil {
		var ret []ProjectParticipation
		return ret
	}
	return o.ProjectParticipations
}

// GetProjectParticipationsOk returns a tuple with the ProjectParticipations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedProjectParticipations) GetProjectParticipationsOk() ([]ProjectParticipation, bool) {
	if o == nil || o.ProjectParticipations == nil {
		return nil, false
	}
	return o.ProjectParticipations, true
}

// HasProjectParticipations returns a boolean if a field has been set.
func (o *PagedProjectParticipations) HasProjectParticipations() bool {
	if o != nil && o.ProjectParticipations != nil {
		return true
	}

	return false
}

// SetProjectParticipations gets a reference to the given []ProjectParticipation and assigns it to the ProjectParticipations field.
func (o *PagedProjectParticipations) SetProjectParticipations(v []ProjectParticipation) {
	o.ProjectParticipations = v
}

func (o PagedProjectParticipations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPage, errPage := json.Marshal(o.Page)
	if errPage != nil {
		return []byte{}, errPage
	}
	errPage = json.Unmarshal([]byte(serializedPage), &toSerialize)
	if errPage != nil {
		return []byte{}, errPage
	}
	if o.ProjectParticipations != nil {
		toSerialize["projectParticipations"] = o.ProjectParticipations
	}
	return json.Marshal(toSerialize)
}

type NullablePagedProjectParticipations struct {
	value *PagedProjectParticipations
	isSet bool
}

func (v NullablePagedProjectParticipations) Get() *PagedProjectParticipations {
	return v.value
}

func (v *NullablePagedProjectParticipations) Set(val *PagedProjectParticipations) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedProjectParticipations) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedProjectParticipations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedProjectParticipations(val *PagedProjectParticipations) *NullablePagedProjectParticipations {
	return &NullablePagedProjectParticipations{value: val, isSet: true}
}

func (v NullablePagedProjectParticipations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedProjectParticipations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

