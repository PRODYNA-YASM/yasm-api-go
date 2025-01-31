/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.58.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PagedProjectParticipation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagedProjectParticipation{}

// PagedProjectParticipation struct for PagedProjectParticipation
type PagedProjectParticipation struct {
	Page
	ProjectParticipation []ProjectParticipationDetails `json:"projectParticipation,omitempty"`
}

// NewPagedProjectParticipation instantiates a new PagedProjectParticipation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedProjectParticipation() *PagedProjectParticipation {
	this := PagedProjectParticipation{}
	return &this
}

// NewPagedProjectParticipationWithDefaults instantiates a new PagedProjectParticipation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedProjectParticipationWithDefaults() *PagedProjectParticipation {
	this := PagedProjectParticipation{}
	return &this
}

// GetProjectParticipation returns the ProjectParticipation field value if set, zero value otherwise.
func (o *PagedProjectParticipation) GetProjectParticipation() []ProjectParticipationDetails {
	if o == nil || IsNil(o.ProjectParticipation) {
		var ret []ProjectParticipationDetails
		return ret
	}
	return o.ProjectParticipation
}

// GetProjectParticipationOk returns a tuple with the ProjectParticipation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedProjectParticipation) GetProjectParticipationOk() ([]ProjectParticipationDetails, bool) {
	if o == nil || IsNil(o.ProjectParticipation) {
		return nil, false
	}
	return o.ProjectParticipation, true
}

// HasProjectParticipation returns a boolean if a field has been set.
func (o *PagedProjectParticipation) HasProjectParticipation() bool {
	if o != nil && !IsNil(o.ProjectParticipation) {
		return true
	}

	return false
}

// SetProjectParticipation gets a reference to the given []ProjectParticipationDetails and assigns it to the ProjectParticipation field.
func (o *PagedProjectParticipation) SetProjectParticipation(v []ProjectParticipationDetails) {
	o.ProjectParticipation = v
}

func (o PagedProjectParticipation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagedProjectParticipation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPage, errPage := json.Marshal(o.Page)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	errPage = json.Unmarshal([]byte(serializedPage), &toSerialize)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	if !IsNil(o.ProjectParticipation) {
		toSerialize["projectParticipation"] = o.ProjectParticipation
	}
	return toSerialize, nil
}

type NullablePagedProjectParticipation struct {
	value *PagedProjectParticipation
	isSet bool
}

func (v NullablePagedProjectParticipation) Get() *PagedProjectParticipation {
	return v.value
}

func (v *NullablePagedProjectParticipation) Set(val *PagedProjectParticipation) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedProjectParticipation) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedProjectParticipation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedProjectParticipation(val *PagedProjectParticipation) *NullablePagedProjectParticipation {
	return &NullablePagedProjectParticipation{value: val, isSet: true}
}

func (v NullablePagedProjectParticipation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedProjectParticipation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


