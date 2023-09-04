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

// PagedIndustries struct for PagedIndustries
type PagedIndustries struct {
	Page
	Industries []IndustryDetails `json:"industries,omitempty"`
}

// NewPagedIndustries instantiates a new PagedIndustries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedIndustries() *PagedIndustries {
	this := PagedIndustries{}
	return &this
}

// NewPagedIndustriesWithDefaults instantiates a new PagedIndustries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedIndustriesWithDefaults() *PagedIndustries {
	this := PagedIndustries{}
	return &this
}

// GetIndustries returns the Industries field value if set, zero value otherwise.
func (o *PagedIndustries) GetIndustries() []IndustryDetails {
	if o == nil || o.Industries == nil {
		var ret []IndustryDetails
		return ret
	}
	return o.Industries
}

// GetIndustriesOk returns a tuple with the Industries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedIndustries) GetIndustriesOk() ([]IndustryDetails, bool) {
	if o == nil || o.Industries == nil {
		return nil, false
	}
	return o.Industries, true
}

// HasIndustries returns a boolean if a field has been set.
func (o *PagedIndustries) HasIndustries() bool {
	if o != nil && o.Industries != nil {
		return true
	}

	return false
}

// SetIndustries gets a reference to the given []IndustryDetails and assigns it to the Industries field.
func (o *PagedIndustries) SetIndustries(v []IndustryDetails) {
	o.Industries = v
}

func (o PagedIndustries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPage, errPage := json.Marshal(o.Page)
	if errPage != nil {
		return []byte{}, errPage
	}
	errPage = json.Unmarshal([]byte(serializedPage), &toSerialize)
	if errPage != nil {
		return []byte{}, errPage
	}
	if o.Industries != nil {
		toSerialize["industries"] = o.Industries
	}
	return json.Marshal(toSerialize)
}

type NullablePagedIndustries struct {
	value *PagedIndustries
	isSet bool
}

func (v NullablePagedIndustries) Get() *PagedIndustries {
	return v.value
}

func (v *NullablePagedIndustries) Set(val *PagedIndustries) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedIndustries) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedIndustries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedIndustries(val *PagedIndustries) *NullablePagedIndustries {
	return &NullablePagedIndustries{value: val, isSet: true}
}

func (v NullablePagedIndustries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedIndustries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


