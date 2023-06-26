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

// ProjectParticipationSearch struct for ProjectParticipationSearch
type ProjectParticipationSearch struct {
	ProjectIds []string `json:"projectIds,omitempty"`
	PersonIds []string `json:"personIds,omitempty"`
	ProjectParticipationIds []string `json:"projectParticipationIds,omitempty"`
}

// NewProjectParticipationSearch instantiates a new ProjectParticipationSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectParticipationSearch() *ProjectParticipationSearch {
	this := ProjectParticipationSearch{}
	return &this
}

// NewProjectParticipationSearchWithDefaults instantiates a new ProjectParticipationSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectParticipationSearchWithDefaults() *ProjectParticipationSearch {
	this := ProjectParticipationSearch{}
	return &this
}

// GetProjectIds returns the ProjectIds field value if set, zero value otherwise.
func (o *ProjectParticipationSearch) GetProjectIds() []string {
	if o == nil || o.ProjectIds == nil {
		var ret []string
		return ret
	}
	return o.ProjectIds
}

// GetProjectIdsOk returns a tuple with the ProjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipationSearch) GetProjectIdsOk() ([]string, bool) {
	if o == nil || o.ProjectIds == nil {
		return nil, false
	}
	return o.ProjectIds, true
}

// HasProjectIds returns a boolean if a field has been set.
func (o *ProjectParticipationSearch) HasProjectIds() bool {
	if o != nil && o.ProjectIds != nil {
		return true
	}

	return false
}

// SetProjectIds gets a reference to the given []string and assigns it to the ProjectIds field.
func (o *ProjectParticipationSearch) SetProjectIds(v []string) {
	o.ProjectIds = v
}

// GetPersonIds returns the PersonIds field value if set, zero value otherwise.
func (o *ProjectParticipationSearch) GetPersonIds() []string {
	if o == nil || o.PersonIds == nil {
		var ret []string
		return ret
	}
	return o.PersonIds
}

// GetPersonIdsOk returns a tuple with the PersonIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipationSearch) GetPersonIdsOk() ([]string, bool) {
	if o == nil || o.PersonIds == nil {
		return nil, false
	}
	return o.PersonIds, true
}

// HasPersonIds returns a boolean if a field has been set.
func (o *ProjectParticipationSearch) HasPersonIds() bool {
	if o != nil && o.PersonIds != nil {
		return true
	}

	return false
}

// SetPersonIds gets a reference to the given []string and assigns it to the PersonIds field.
func (o *ProjectParticipationSearch) SetPersonIds(v []string) {
	o.PersonIds = v
}

// GetProjectParticipationIds returns the ProjectParticipationIds field value if set, zero value otherwise.
func (o *ProjectParticipationSearch) GetProjectParticipationIds() []string {
	if o == nil || o.ProjectParticipationIds == nil {
		var ret []string
		return ret
	}
	return o.ProjectParticipationIds
}

// GetProjectParticipationIdsOk returns a tuple with the ProjectParticipationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipationSearch) GetProjectParticipationIdsOk() ([]string, bool) {
	if o == nil || o.ProjectParticipationIds == nil {
		return nil, false
	}
	return o.ProjectParticipationIds, true
}

// HasProjectParticipationIds returns a boolean if a field has been set.
func (o *ProjectParticipationSearch) HasProjectParticipationIds() bool {
	if o != nil && o.ProjectParticipationIds != nil {
		return true
	}

	return false
}

// SetProjectParticipationIds gets a reference to the given []string and assigns it to the ProjectParticipationIds field.
func (o *ProjectParticipationSearch) SetProjectParticipationIds(v []string) {
	o.ProjectParticipationIds = v
}

func (o ProjectParticipationSearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectIds != nil {
		toSerialize["projectIds"] = o.ProjectIds
	}
	if o.PersonIds != nil {
		toSerialize["personIds"] = o.PersonIds
	}
	if o.ProjectParticipationIds != nil {
		toSerialize["projectParticipationIds"] = o.ProjectParticipationIds
	}
	return json.Marshal(toSerialize)
}

type NullableProjectParticipationSearch struct {
	value *ProjectParticipationSearch
	isSet bool
}

func (v NullableProjectParticipationSearch) Get() *ProjectParticipationSearch {
	return v.value
}

func (v *NullableProjectParticipationSearch) Set(val *ProjectParticipationSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectParticipationSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectParticipationSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectParticipationSearch(val *ProjectParticipationSearch) *NullableProjectParticipationSearch {
	return &NullableProjectParticipationSearch{value: val, isSet: true}
}

func (v NullableProjectParticipationSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectParticipationSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

