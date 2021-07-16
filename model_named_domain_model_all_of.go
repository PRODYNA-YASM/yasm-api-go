/*
 * YASM (Yet Another Skill Management) API
 *
 * This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.
 *
 * API version: 0.7.8
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NamedDomainModelAllOf struct for NamedDomainModelAllOf
type NamedDomainModelAllOf struct {
	Name *string `json:"name,omitempty"`
}

// NewNamedDomainModelAllOf instantiates a new NamedDomainModelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamedDomainModelAllOf() *NamedDomainModelAllOf {
	this := NamedDomainModelAllOf{}
	return &this
}

// NewNamedDomainModelAllOfWithDefaults instantiates a new NamedDomainModelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamedDomainModelAllOfWithDefaults() *NamedDomainModelAllOf {
	this := NamedDomainModelAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NamedDomainModelAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamedDomainModelAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NamedDomainModelAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NamedDomainModelAllOf) SetName(v string) {
	o.Name = &v
}

func (o NamedDomainModelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableNamedDomainModelAllOf struct {
	value *NamedDomainModelAllOf
	isSet bool
}

func (v NullableNamedDomainModelAllOf) Get() *NamedDomainModelAllOf {
	return v.value
}

func (v *NullableNamedDomainModelAllOf) Set(val *NamedDomainModelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNamedDomainModelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNamedDomainModelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamedDomainModelAllOf(val *NamedDomainModelAllOf) *NullableNamedDomainModelAllOf {
	return &NullableNamedDomainModelAllOf{value: val, isSet: true}
}

func (v NullableNamedDomainModelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamedDomainModelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


