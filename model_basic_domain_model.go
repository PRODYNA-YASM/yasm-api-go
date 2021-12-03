/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 0.9.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// BasicDomainModel struct for BasicDomainModel
type BasicDomainModel struct {
	Id string `json:"id"`
}

// NewBasicDomainModel instantiates a new BasicDomainModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicDomainModel(id string) *BasicDomainModel {
	this := BasicDomainModel{}
	this.Id = id
	return &this
}

// NewBasicDomainModelWithDefaults instantiates a new BasicDomainModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicDomainModelWithDefaults() *BasicDomainModel {
	this := BasicDomainModel{}
	return &this
}

// GetId returns the Id field value
func (o *BasicDomainModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BasicDomainModel) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BasicDomainModel) SetId(v string) {
	o.Id = v
}

func (o BasicDomainModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableBasicDomainModel struct {
	value *BasicDomainModel
	isSet bool
}

func (v NullableBasicDomainModel) Get() *BasicDomainModel {
	return v.value
}

func (v *NullableBasicDomainModel) Set(val *BasicDomainModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicDomainModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicDomainModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicDomainModel(val *BasicDomainModel) *NullableBasicDomainModel {
	return &NullableBasicDomainModel{value: val, isSet: true}
}

func (v NullableBasicDomainModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicDomainModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


