/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.73.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the AbstractEntityFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AbstractEntityFilter{}

// AbstractEntityFilter struct for AbstractEntityFilter
type AbstractEntityFilter struct {
	Id string `json:"id"`
	ObjectType *string `json:"objectType,omitempty"`
}

// NewAbstractEntityFilter instantiates a new AbstractEntityFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAbstractEntityFilter(id string) *AbstractEntityFilter {
	this := AbstractEntityFilter{}
	this.Id = id
	return &this
}

// NewAbstractEntityFilterWithDefaults instantiates a new AbstractEntityFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAbstractEntityFilterWithDefaults() *AbstractEntityFilter {
	this := AbstractEntityFilter{}
	return &this
}

// GetId returns the Id field value
func (o *AbstractEntityFilter) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AbstractEntityFilter) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AbstractEntityFilter) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *AbstractEntityFilter) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbstractEntityFilter) GetObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *AbstractEntityFilter) HasObjectType() bool {
	if o != nil && !IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *AbstractEntityFilter) SetObjectType(v string) {
	o.ObjectType = &v
}

func (o AbstractEntityFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AbstractEntityFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.ObjectType) {
		toSerialize["objectType"] = o.ObjectType
	}
	return toSerialize, nil
}

type NullableAbstractEntityFilter struct {
	value *AbstractEntityFilter
	isSet bool
}

func (v NullableAbstractEntityFilter) Get() *AbstractEntityFilter {
	return v.value
}

func (v *NullableAbstractEntityFilter) Set(val *AbstractEntityFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableAbstractEntityFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableAbstractEntityFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAbstractEntityFilter(val *AbstractEntityFilter) *NullableAbstractEntityFilter {
	return &NullableAbstractEntityFilter{value: val, isSet: true}
}

func (v NullableAbstractEntityFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAbstractEntityFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


