/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.4.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PagedProjectParticipationsAllOf struct for PagedProjectParticipationsAllOf
type PagedProjectParticipationsAllOf struct {
	ProjectParticipations []ProjectParticipation `json:"projectParticipations,omitempty"`
}

// NewPagedProjectParticipationsAllOf instantiates a new PagedProjectParticipationsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedProjectParticipationsAllOf() *PagedProjectParticipationsAllOf {
	this := PagedProjectParticipationsAllOf{}
	return &this
}

// NewPagedProjectParticipationsAllOfWithDefaults instantiates a new PagedProjectParticipationsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedProjectParticipationsAllOfWithDefaults() *PagedProjectParticipationsAllOf {
	this := PagedProjectParticipationsAllOf{}
	return &this
}

// GetProjectParticipations returns the ProjectParticipations field value if set, zero value otherwise.
func (o *PagedProjectParticipationsAllOf) GetProjectParticipations() []ProjectParticipation {
	if o == nil || o.ProjectParticipations == nil {
		var ret []ProjectParticipation
		return ret
	}
	return o.ProjectParticipations
}

// GetProjectParticipationsOk returns a tuple with the ProjectParticipations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedProjectParticipationsAllOf) GetProjectParticipationsOk() ([]ProjectParticipation, bool) {
	if o == nil || o.ProjectParticipations == nil {
		return nil, false
	}
	return o.ProjectParticipations, true
}

// HasProjectParticipations returns a boolean if a field has been set.
func (o *PagedProjectParticipationsAllOf) HasProjectParticipations() bool {
	if o != nil && o.ProjectParticipations != nil {
		return true
	}

	return false
}

// SetProjectParticipations gets a reference to the given []ProjectParticipation and assigns it to the ProjectParticipations field.
func (o *PagedProjectParticipationsAllOf) SetProjectParticipations(v []ProjectParticipation) {
	o.ProjectParticipations = v
}

func (o PagedProjectParticipationsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectParticipations != nil {
		toSerialize["projectParticipations"] = o.ProjectParticipations
	}
	return json.Marshal(toSerialize)
}

type NullablePagedProjectParticipationsAllOf struct {
	value *PagedProjectParticipationsAllOf
	isSet bool
}

func (v NullablePagedProjectParticipationsAllOf) Get() *PagedProjectParticipationsAllOf {
	return v.value
}

func (v *NullablePagedProjectParticipationsAllOf) Set(val *PagedProjectParticipationsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedProjectParticipationsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedProjectParticipationsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedProjectParticipationsAllOf(val *PagedProjectParticipationsAllOf) *NullablePagedProjectParticipationsAllOf {
	return &NullablePagedProjectParticipationsAllOf{value: val, isSet: true}
}

func (v NullablePagedProjectParticipationsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedProjectParticipationsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


