/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the Layout type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Layout{}

// Layout struct for Layout
type Layout struct {
	BasicDomainModel
	Fields []LayoutField `json:"fields,omitempty"`
}

// NewLayout instantiates a new Layout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLayout(id string) *Layout {
	this := Layout{}
	this.Id = id
	return &this
}

// NewLayoutWithDefaults instantiates a new Layout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLayoutWithDefaults() *Layout {
	this := Layout{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *Layout) GetFields() []LayoutField {
	if o == nil || IsNil(o.Fields) {
		var ret []LayoutField
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Layout) GetFieldsOk() ([]LayoutField, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *Layout) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []LayoutField and assigns it to the Fields field.
func (o *Layout) SetFields(v []LayoutField) {
	o.Fields = v
}

func (o Layout) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Layout) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedBasicDomainModel, errBasicDomainModel := json.Marshal(o.BasicDomainModel)
	if errBasicDomainModel != nil {
		return map[string]interface{}{}, errBasicDomainModel
	}
	errBasicDomainModel = json.Unmarshal([]byte(serializedBasicDomainModel), &toSerialize)
	if errBasicDomainModel != nil {
		return map[string]interface{}{}, errBasicDomainModel
	}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	return toSerialize, nil
}

type NullableLayout struct {
	value *Layout
	isSet bool
}

func (v NullableLayout) Get() *Layout {
	return v.value
}

func (v *NullableLayout) Set(val *Layout) {
	v.value = val
	v.isSet = true
}

func (v NullableLayout) IsSet() bool {
	return v.isSet
}

func (v *NullableLayout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLayout(val *Layout) *NullableLayout {
	return &NullableLayout{value: val, isSet: true}
}

func (v NullableLayout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLayout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


