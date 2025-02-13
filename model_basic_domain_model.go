/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.66.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the BasicDomainModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BasicDomainModel{}

// BasicDomainModel struct for BasicDomainModel
type BasicDomainModel struct {
	Id string `json:"id"`
	ObjectType *string `json:"objectType,omitempty"`
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
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BasicDomainModel) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *BasicDomainModel) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicDomainModel) GetObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *BasicDomainModel) HasObjectType() bool {
	if o != nil && !IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *BasicDomainModel) SetObjectType(v string) {
	o.ObjectType = &v
}

func (o BasicDomainModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BasicDomainModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.ObjectType) {
		toSerialize["objectType"] = o.ObjectType
	}
	return toSerialize, nil
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


