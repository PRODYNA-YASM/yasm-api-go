/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.0.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Linkable struct for Linkable
type Linkable struct {
	// The entity can be linked
	Linkable *bool `json:"linkable,omitempty"`
}

// NewLinkable instantiates a new Linkable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkable() *Linkable {
	this := Linkable{}
	var linkable bool = false
	this.Linkable = &linkable
	return &this
}

// NewLinkableWithDefaults instantiates a new Linkable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkableWithDefaults() *Linkable {
	this := Linkable{}
	var linkable bool = false
	this.Linkable = &linkable
	return &this
}

// GetLinkable returns the Linkable field value if set, zero value otherwise.
func (o *Linkable) GetLinkable() bool {
	if o == nil || o.Linkable == nil {
		var ret bool
		return ret
	}
	return *o.Linkable
}

// GetLinkableOk returns a tuple with the Linkable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linkable) GetLinkableOk() (*bool, bool) {
	if o == nil || o.Linkable == nil {
		return nil, false
	}
	return o.Linkable, true
}

// HasLinkable returns a boolean if a field has been set.
func (o *Linkable) HasLinkable() bool {
	if o != nil && o.Linkable != nil {
		return true
	}

	return false
}

// SetLinkable gets a reference to the given bool and assigns it to the Linkable field.
func (o *Linkable) SetLinkable(v bool) {
	o.Linkable = &v
}

func (o Linkable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Linkable != nil {
		toSerialize["linkable"] = o.Linkable
	}
	return json.Marshal(toSerialize)
}

type NullableLinkable struct {
	value *Linkable
	isSet bool
}

func (v NullableLinkable) Get() *Linkable {
	return v.value
}

func (v *NullableLinkable) Set(val *Linkable) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkable) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkable(val *Linkable) *NullableLinkable {
	return &NullableLinkable{value: val, isSet: true}
}

func (v NullableLinkable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

