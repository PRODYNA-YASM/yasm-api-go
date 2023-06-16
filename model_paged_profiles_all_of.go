/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.4.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PagedProfilesAllOf struct for PagedProfilesAllOf
type PagedProfilesAllOf struct {
	Profiles []ProfileDetails `json:"profiles,omitempty"`
}

// NewPagedProfilesAllOf instantiates a new PagedProfilesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedProfilesAllOf() *PagedProfilesAllOf {
	this := PagedProfilesAllOf{}
	return &this
}

// NewPagedProfilesAllOfWithDefaults instantiates a new PagedProfilesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedProfilesAllOfWithDefaults() *PagedProfilesAllOf {
	this := PagedProfilesAllOf{}
	return &this
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *PagedProfilesAllOf) GetProfiles() []ProfileDetails {
	if o == nil || o.Profiles == nil {
		var ret []ProfileDetails
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedProfilesAllOf) GetProfilesOk() ([]ProfileDetails, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *PagedProfilesAllOf) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []ProfileDetails and assigns it to the Profiles field.
func (o *PagedProfilesAllOf) SetProfiles(v []ProfileDetails) {
	o.Profiles = v
}

func (o PagedProfilesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Profiles != nil {
		toSerialize["profiles"] = o.Profiles
	}
	return json.Marshal(toSerialize)
}

type NullablePagedProfilesAllOf struct {
	value *PagedProfilesAllOf
	isSet bool
}

func (v NullablePagedProfilesAllOf) Get() *PagedProfilesAllOf {
	return v.value
}

func (v *NullablePagedProfilesAllOf) Set(val *PagedProfilesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedProfilesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedProfilesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedProfilesAllOf(val *PagedProfilesAllOf) *NullablePagedProfilesAllOf {
	return &NullablePagedProfilesAllOf{value: val, isSet: true}
}

func (v NullablePagedProfilesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedProfilesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


