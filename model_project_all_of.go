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

// ProjectAllOf struct for ProjectAllOf
type ProjectAllOf struct {
	// true if project was done outside of the organization
	External *bool `json:"external,omitempty"`
}

// NewProjectAllOf instantiates a new ProjectAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectAllOf() *ProjectAllOf {
	this := ProjectAllOf{}
	var external bool = false
	this.External = &external
	return &this
}

// NewProjectAllOfWithDefaults instantiates a new ProjectAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectAllOfWithDefaults() *ProjectAllOf {
	this := ProjectAllOf{}
	var external bool = false
	this.External = &external
	return &this
}

// GetExternal returns the External field value if set, zero value otherwise.
func (o *ProjectAllOf) GetExternal() bool {
	if o == nil || o.External == nil {
		var ret bool
		return ret
	}
	return *o.External
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAllOf) GetExternalOk() (*bool, bool) {
	if o == nil || o.External == nil {
		return nil, false
	}
	return o.External, true
}

// HasExternal returns a boolean if a field has been set.
func (o *ProjectAllOf) HasExternal() bool {
	if o != nil && o.External != nil {
		return true
	}

	return false
}

// SetExternal gets a reference to the given bool and assigns it to the External field.
func (o *ProjectAllOf) SetExternal(v bool) {
	o.External = &v
}

func (o ProjectAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.External != nil {
		toSerialize["external"] = o.External
	}
	return json.Marshal(toSerialize)
}

type NullableProjectAllOf struct {
	value *ProjectAllOf
	isSet bool
}

func (v NullableProjectAllOf) Get() *ProjectAllOf {
	return v.value
}

func (v *NullableProjectAllOf) Set(val *ProjectAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectAllOf(val *ProjectAllOf) *NullableProjectAllOf {
	return &NullableProjectAllOf{value: val, isSet: true}
}

func (v NullableProjectAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


