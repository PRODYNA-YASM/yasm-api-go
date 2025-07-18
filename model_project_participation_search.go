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

// checks if the ProjectParticipationSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectParticipationSearch{}

// ProjectParticipationSearch struct for ProjectParticipationSearch
type ProjectParticipationSearch struct {
	Search
	LinkToErp *bool `json:"linkToErp,omitempty"`
	ProjectIds []string `json:"projectIds,omitempty"`
	PersonIds []string `json:"personIds,omitempty"`
	ProjectParticipationIds []string `json:"projectParticipationIds,omitempty"`
	AwardIds []string `json:"awardIds,omitempty"`
}

// NewProjectParticipationSearch instantiates a new ProjectParticipationSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectParticipationSearch(skip int32, limit int32) *ProjectParticipationSearch {
	this := ProjectParticipationSearch{}
	this.Skip = skip
	this.Limit = limit
	return &this
}

// NewProjectParticipationSearchWithDefaults instantiates a new ProjectParticipationSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectParticipationSearchWithDefaults() *ProjectParticipationSearch {
	this := ProjectParticipationSearch{}
	return &this
}

// GetLinkToErp returns the LinkToErp field value if set, zero value otherwise.
func (o *ProjectParticipationSearch) GetLinkToErp() bool {
	if o == nil || IsNil(o.LinkToErp) {
		var ret bool
		return ret
	}
	return *o.LinkToErp
}

// GetLinkToErpOk returns a tuple with the LinkToErp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipationSearch) GetLinkToErpOk() (*bool, bool) {
	if o == nil || IsNil(o.LinkToErp) {
		return nil, false
	}
	return o.LinkToErp, true
}

// HasLinkToErp returns a boolean if a field has been set.
func (o *ProjectParticipationSearch) HasLinkToErp() bool {
	if o != nil && !IsNil(o.LinkToErp) {
		return true
	}

	return false
}

// SetLinkToErp gets a reference to the given bool and assigns it to the LinkToErp field.
func (o *ProjectParticipationSearch) SetLinkToErp(v bool) {
	o.LinkToErp = &v
}

// GetProjectIds returns the ProjectIds field value if set, zero value otherwise.
func (o *ProjectParticipationSearch) GetProjectIds() []string {
	if o == nil || IsNil(o.ProjectIds) {
		var ret []string
		return ret
	}
	return o.ProjectIds
}

// GetProjectIdsOk returns a tuple with the ProjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipationSearch) GetProjectIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProjectIds) {
		return nil, false
	}
	return o.ProjectIds, true
}

// HasProjectIds returns a boolean if a field has been set.
func (o *ProjectParticipationSearch) HasProjectIds() bool {
	if o != nil && !IsNil(o.ProjectIds) {
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
	if o == nil || IsNil(o.PersonIds) {
		var ret []string
		return ret
	}
	return o.PersonIds
}

// GetPersonIdsOk returns a tuple with the PersonIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipationSearch) GetPersonIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.PersonIds) {
		return nil, false
	}
	return o.PersonIds, true
}

// HasPersonIds returns a boolean if a field has been set.
func (o *ProjectParticipationSearch) HasPersonIds() bool {
	if o != nil && !IsNil(o.PersonIds) {
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
	if o == nil || IsNil(o.ProjectParticipationIds) {
		var ret []string
		return ret
	}
	return o.ProjectParticipationIds
}

// GetProjectParticipationIdsOk returns a tuple with the ProjectParticipationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipationSearch) GetProjectParticipationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProjectParticipationIds) {
		return nil, false
	}
	return o.ProjectParticipationIds, true
}

// HasProjectParticipationIds returns a boolean if a field has been set.
func (o *ProjectParticipationSearch) HasProjectParticipationIds() bool {
	if o != nil && !IsNil(o.ProjectParticipationIds) {
		return true
	}

	return false
}

// SetProjectParticipationIds gets a reference to the given []string and assigns it to the ProjectParticipationIds field.
func (o *ProjectParticipationSearch) SetProjectParticipationIds(v []string) {
	o.ProjectParticipationIds = v
}

// GetAwardIds returns the AwardIds field value if set, zero value otherwise.
func (o *ProjectParticipationSearch) GetAwardIds() []string {
	if o == nil || IsNil(o.AwardIds) {
		var ret []string
		return ret
	}
	return o.AwardIds
}

// GetAwardIdsOk returns a tuple with the AwardIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipationSearch) GetAwardIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AwardIds) {
		return nil, false
	}
	return o.AwardIds, true
}

// HasAwardIds returns a boolean if a field has been set.
func (o *ProjectParticipationSearch) HasAwardIds() bool {
	if o != nil && !IsNil(o.AwardIds) {
		return true
	}

	return false
}

// SetAwardIds gets a reference to the given []string and assigns it to the AwardIds field.
func (o *ProjectParticipationSearch) SetAwardIds(v []string) {
	o.AwardIds = v
}

func (o ProjectParticipationSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectParticipationSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSearch, errSearch := json.Marshal(o.Search)
	if errSearch != nil {
		return map[string]interface{}{}, errSearch
	}
	errSearch = json.Unmarshal([]byte(serializedSearch), &toSerialize)
	if errSearch != nil {
		return map[string]interface{}{}, errSearch
	}
	if !IsNil(o.LinkToErp) {
		toSerialize["linkToErp"] = o.LinkToErp
	}
	if !IsNil(o.ProjectIds) {
		toSerialize["projectIds"] = o.ProjectIds
	}
	if !IsNil(o.PersonIds) {
		toSerialize["personIds"] = o.PersonIds
	}
	if !IsNil(o.ProjectParticipationIds) {
		toSerialize["projectParticipationIds"] = o.ProjectParticipationIds
	}
	if !IsNil(o.AwardIds) {
		toSerialize["awardIds"] = o.AwardIds
	}
	return toSerialize, nil
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


