/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PagedPersons struct for PagedPersons
type PagedPersons struct {
	Skip *int32 `json:"skip,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	Count *int32 `json:"count,omitempty"`
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

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *PagedPersons) GetSkip() int32 {
	if o == nil || o.Skip == nil {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedPersons) GetSkipOk() (*int32, bool) {
	if o == nil || o.Skip == nil {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *PagedPersons) HasSkip() bool {
	if o != nil && o.Skip != nil {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *PagedPersons) SetSkip(v int32) {
	o.Skip = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *PagedPersons) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedPersons) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *PagedPersons) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *PagedPersons) SetLimit(v int32) {
	o.Limit = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PagedPersons) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedPersons) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PagedPersons) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PagedPersons) SetCount(v int32) {
	o.Count = &v
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
	if o.Skip != nil {
		toSerialize["skip"] = o.Skip
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
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


