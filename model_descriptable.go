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

// checks if the Descriptable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Descriptable{}

// Descriptable struct for Descriptable
type Descriptable struct {
	Description *string `json:"description,omitempty"`
}

// NewDescriptable instantiates a new Descriptable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescriptable() *Descriptable {
	this := Descriptable{}
	return &this
}

// NewDescriptableWithDefaults instantiates a new Descriptable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescriptableWithDefaults() *Descriptable {
	this := Descriptable{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Descriptable) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Descriptable) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Descriptable) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Descriptable) SetDescription(v string) {
	o.Description = &v
}

func (o Descriptable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Descriptable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableDescriptable struct {
	value *Descriptable
	isSet bool
}

func (v NullableDescriptable) Get() *Descriptable {
	return v.value
}

func (v *NullableDescriptable) Set(val *Descriptable) {
	v.value = val
	v.isSet = true
}

func (v NullableDescriptable) IsSet() bool {
	return v.isSet
}

func (v *NullableDescriptable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescriptable(val *Descriptable) *NullableDescriptable {
	return &NullableDescriptable{value: val, isSet: true}
}

func (v NullableDescriptable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescriptable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


