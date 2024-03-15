/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.16.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PagedLanguages type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagedLanguages{}

// PagedLanguages struct for PagedLanguages
type PagedLanguages struct {
	Page
	Languages []Language `json:"languages,omitempty"`
}

// NewPagedLanguages instantiates a new PagedLanguages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedLanguages() *PagedLanguages {
	this := PagedLanguages{}
	return &this
}

// NewPagedLanguagesWithDefaults instantiates a new PagedLanguages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedLanguagesWithDefaults() *PagedLanguages {
	this := PagedLanguages{}
	return &this
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *PagedLanguages) GetLanguages() []Language {
	if o == nil || IsNil(o.Languages) {
		var ret []Language
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedLanguages) GetLanguagesOk() ([]Language, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *PagedLanguages) HasLanguages() bool {
	if o != nil && !IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []Language and assigns it to the Languages field.
func (o *PagedLanguages) SetLanguages(v []Language) {
	o.Languages = v
}

func (o PagedLanguages) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagedLanguages) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPage, errPage := json.Marshal(o.Page)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	errPage = json.Unmarshal([]byte(serializedPage), &toSerialize)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	if !IsNil(o.Languages) {
		toSerialize["languages"] = o.Languages
	}
	return toSerialize, nil
}

type NullablePagedLanguages struct {
	value *PagedLanguages
	isSet bool
}

func (v NullablePagedLanguages) Get() *PagedLanguages {
	return v.value
}

func (v *NullablePagedLanguages) Set(val *PagedLanguages) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedLanguages) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedLanguages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedLanguages(val *PagedLanguages) *NullablePagedLanguages {
	return &NullablePagedLanguages{value: val, isSet: true}
}

func (v NullablePagedLanguages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedLanguages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


