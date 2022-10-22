/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.3.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PagedPersons struct for PagedPersons
type PagedPersons struct {
	Page
	Persons []PersonScoreDetail `json:"persons,omitempty"`
}

// NewPagedPersons instantiates a new PagedPersons object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedPersons() *PagedPersons {
	this := PagedPersons{}
	return &this
}

// NewPagedPersonsWithDefaults instantiates a new PagedPersons object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedPersonsWithDefaults() *PagedPersons {
	this := PagedPersons{}
	return &this
}

// GetPersons returns the Persons field value if set, zero value otherwise.
func (o *PagedPersons) GetPersons() []PersonScoreDetail {
	if o == nil || o.Persons == nil {
		var ret []PersonScoreDetail
		return ret
	}
	return o.Persons
}

// GetPersonsOk returns a tuple with the Persons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedPersons) GetPersonsOk() ([]PersonScoreDetail, bool) {
	if o == nil || o.Persons == nil {
		return nil, false
	}
	return o.Persons, true
}

// HasPersons returns a boolean if a field has been set.
func (o *PagedPersons) HasPersons() bool {
	if o != nil && o.Persons != nil {
		return true
	}

	return false
}

// SetPersons gets a reference to the given []PersonScoreDetail and assigns it to the Persons field.
func (o *PagedPersons) SetPersons(v []PersonScoreDetail) {
	o.Persons = v
}

func (o PagedPersons) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPage, errPage := json.Marshal(o.Page)
	if errPage != nil {
		return []byte{}, errPage
	}
	errPage = json.Unmarshal([]byte(serializedPage), &toSerialize)
	if errPage != nil {
		return []byte{}, errPage
	}
	if o.Persons != nil {
		toSerialize["persons"] = o.Persons
	}
	return json.Marshal(toSerialize)
}

type NullablePagedPersons struct {
	value *PagedPersons
	isSet bool
}

func (v NullablePagedPersons) Get() *PagedPersons {
	return v.value
}

func (v *NullablePagedPersons) Set(val *PagedPersons) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedPersons) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedPersons) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedPersons(val *PagedPersons) *NullablePagedPersons {
	return &NullablePagedPersons{value: val, isSet: true}
}

func (v NullablePagedPersons) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedPersons) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


