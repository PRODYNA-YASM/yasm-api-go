/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.74.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the Synonymable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Synonymable{}

// Synonymable struct for Synonymable
type Synonymable struct {
	Synonyms []string `json:"synonyms,omitempty"`
}

// NewSynonymable instantiates a new Synonymable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSynonymable() *Synonymable {
	this := Synonymable{}
	return &this
}

// NewSynonymableWithDefaults instantiates a new Synonymable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSynonymableWithDefaults() *Synonymable {
	this := Synonymable{}
	return &this
}

// GetSynonyms returns the Synonyms field value if set, zero value otherwise.
func (o *Synonymable) GetSynonyms() []string {
	if o == nil || IsNil(o.Synonyms) {
		var ret []string
		return ret
	}
	return o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Synonymable) GetSynonymsOk() ([]string, bool) {
	if o == nil || IsNil(o.Synonyms) {
		return nil, false
	}
	return o.Synonyms, true
}

// HasSynonyms returns a boolean if a field has been set.
func (o *Synonymable) HasSynonyms() bool {
	if o != nil && !IsNil(o.Synonyms) {
		return true
	}

	return false
}

// SetSynonyms gets a reference to the given []string and assigns it to the Synonyms field.
func (o *Synonymable) SetSynonyms(v []string) {
	o.Synonyms = v
}

func (o Synonymable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Synonymable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Synonyms) {
		toSerialize["synonyms"] = o.Synonyms
	}
	return toSerialize, nil
}

type NullableSynonymable struct {
	value *Synonymable
	isSet bool
}

func (v NullableSynonymable) Get() *Synonymable {
	return v.value
}

func (v *NullableSynonymable) Set(val *Synonymable) {
	v.value = val
	v.isSet = true
}

func (v NullableSynonymable) IsSet() bool {
	return v.isSet
}

func (v *NullableSynonymable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSynonymable(val *Synonymable) *NullableSynonymable {
	return &NullableSynonymable{value: val, isSet: true}
}

func (v NullableSynonymable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSynonymable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


