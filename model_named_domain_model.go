/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the NamedDomainModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NamedDomainModel{}

// NamedDomainModel struct for NamedDomainModel
type NamedDomainModel struct {
	BasicDomainModel
	Name string `json:"name"`
}

// NewNamedDomainModel instantiates a new NamedDomainModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamedDomainModel(name string, id string) *NamedDomainModel {
	this := NamedDomainModel{}
	this.Id = id
	this.Name = name
	return &this
}

// NewNamedDomainModelWithDefaults instantiates a new NamedDomainModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamedDomainModelWithDefaults() *NamedDomainModel {
	this := NamedDomainModel{}
	return &this
}

// GetName returns the Name field value
func (o *NamedDomainModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NamedDomainModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NamedDomainModel) SetName(v string) {
	o.Name = v
}

func (o NamedDomainModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NamedDomainModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedBasicDomainModel, errBasicDomainModel := json.Marshal(o.BasicDomainModel)
	if errBasicDomainModel != nil {
		return map[string]interface{}{}, errBasicDomainModel
	}
	errBasicDomainModel = json.Unmarshal([]byte(serializedBasicDomainModel), &toSerialize)
	if errBasicDomainModel != nil {
		return map[string]interface{}{}, errBasicDomainModel
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableNamedDomainModel struct {
	value *NamedDomainModel
	isSet bool
}

func (v NullableNamedDomainModel) Get() *NamedDomainModel {
	return v.value
}

func (v *NullableNamedDomainModel) Set(val *NamedDomainModel) {
	v.value = val
	v.isSet = true
}

func (v NullableNamedDomainModel) IsSet() bool {
	return v.isSet
}

func (v *NullableNamedDomainModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamedDomainModel(val *NamedDomainModel) *NullableNamedDomainModel {
	return &NullableNamedDomainModel{value: val, isSet: true}
}

func (v NullableNamedDomainModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamedDomainModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


