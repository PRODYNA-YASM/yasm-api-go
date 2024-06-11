/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the SkillSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkillSearch{}

// SkillSearch struct for SkillSearch
type SkillSearch struct {
	Search
	// Gives you either all skills, only the root kills
	Types *string `json:"types,omitempty"`
	SkillIds []string `json:"skillIds,omitempty"`
}

// NewSkillSearch instantiates a new SkillSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkillSearch(skip int32, limit int32) *SkillSearch {
	this := SkillSearch{}
	this.Skip = skip
	this.Limit = limit
	var types string = "all"
	this.Types = &types
	return &this
}

// NewSkillSearchWithDefaults instantiates a new SkillSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkillSearchWithDefaults() *SkillSearch {
	this := SkillSearch{}
	var types string = "all"
	this.Types = &types
	return &this
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *SkillSearch) GetTypes() string {
	if o == nil || IsNil(o.Types) {
		var ret string
		return ret
	}
	return *o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillSearch) GetTypesOk() (*string, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *SkillSearch) HasTypes() bool {
	if o != nil && !IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given string and assigns it to the Types field.
func (o *SkillSearch) SetTypes(v string) {
	o.Types = &v
}

// GetSkillIds returns the SkillIds field value if set, zero value otherwise.
func (o *SkillSearch) GetSkillIds() []string {
	if o == nil || IsNil(o.SkillIds) {
		var ret []string
		return ret
	}
	return o.SkillIds
}

// GetSkillIdsOk returns a tuple with the SkillIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillSearch) GetSkillIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SkillIds) {
		return nil, false
	}
	return o.SkillIds, true
}

// HasSkillIds returns a boolean if a field has been set.
func (o *SkillSearch) HasSkillIds() bool {
	if o != nil && !IsNil(o.SkillIds) {
		return true
	}

	return false
}

// SetSkillIds gets a reference to the given []string and assigns it to the SkillIds field.
func (o *SkillSearch) SetSkillIds(v []string) {
	o.SkillIds = v
}

func (o SkillSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkillSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSearch, errSearch := json.Marshal(o.Search)
	if errSearch != nil {
		return map[string]interface{}{}, errSearch
	}
	errSearch = json.Unmarshal([]byte(serializedSearch), &toSerialize)
	if errSearch != nil {
		return map[string]interface{}{}, errSearch
	}
	if !IsNil(o.Types) {
		toSerialize["types"] = o.Types
	}
	if !IsNil(o.SkillIds) {
		toSerialize["skillIds"] = o.SkillIds
	}
	return toSerialize, nil
}

type NullableSkillSearch struct {
	value *SkillSearch
	isSet bool
}

func (v NullableSkillSearch) Get() *SkillSearch {
	return v.value
}

func (v *NullableSkillSearch) Set(val *SkillSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableSkillSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableSkillSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkillSearch(val *SkillSearch) *NullableSkillSearch {
	return &NullableSkillSearch{value: val, isSet: true}
}

func (v NullableSkillSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkillSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

