/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.24.0
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
	// Optionally filter skills based on suggestion
	Suggestions *string `json:"suggestions,omitempty"`
	// Optionally filter skills based on linkable
	Linkable *bool `json:"linkable,omitempty"`
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
	var suggestions string = "all"
	this.Suggestions = &suggestions
	return &this
}

// NewSkillSearchWithDefaults instantiates a new SkillSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkillSearchWithDefaults() *SkillSearch {
	this := SkillSearch{}
	var types string = "all"
	this.Types = &types
	var suggestions string = "all"
	this.Suggestions = &suggestions
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

// GetSuggestions returns the Suggestions field value if set, zero value otherwise.
func (o *SkillSearch) GetSuggestions() string {
	if o == nil || IsNil(o.Suggestions) {
		var ret string
		return ret
	}
	return *o.Suggestions
}

// GetSuggestionsOk returns a tuple with the Suggestions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillSearch) GetSuggestionsOk() (*string, bool) {
	if o == nil || IsNil(o.Suggestions) {
		return nil, false
	}
	return o.Suggestions, true
}

// HasSuggestions returns a boolean if a field has been set.
func (o *SkillSearch) HasSuggestions() bool {
	if o != nil && !IsNil(o.Suggestions) {
		return true
	}

	return false
}

// SetSuggestions gets a reference to the given string and assigns it to the Suggestions field.
func (o *SkillSearch) SetSuggestions(v string) {
	o.Suggestions = &v
}

// GetLinkable returns the Linkable field value if set, zero value otherwise.
func (o *SkillSearch) GetLinkable() bool {
	if o == nil || IsNil(o.Linkable) {
		var ret bool
		return ret
	}
	return *o.Linkable
}

// GetLinkableOk returns a tuple with the Linkable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillSearch) GetLinkableOk() (*bool, bool) {
	if o == nil || IsNil(o.Linkable) {
		return nil, false
	}
	return o.Linkable, true
}

// HasLinkable returns a boolean if a field has been set.
func (o *SkillSearch) HasLinkable() bool {
	if o != nil && !IsNil(o.Linkable) {
		return true
	}

	return false
}

// SetLinkable gets a reference to the given bool and assigns it to the Linkable field.
func (o *SkillSearch) SetLinkable(v bool) {
	o.Linkable = &v
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
	if !IsNil(o.Suggestions) {
		toSerialize["suggestions"] = o.Suggestions
	}
	if !IsNil(o.Linkable) {
		toSerialize["linkable"] = o.Linkable
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


