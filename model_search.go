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

// checks if the Search type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Search{}

// Search struct for Search
type Search struct {
	Term *string `json:"term,omitempty"`
	Skip int32 `json:"skip"`
	Limit int32 `json:"limit"`
	ObjectType *string `json:"objectType,omitempty"`
}

// NewSearch instantiates a new Search object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearch(skip int32, limit int32) *Search {
	this := Search{}
	this.Skip = skip
	this.Limit = limit
	return &this
}

// NewSearchWithDefaults instantiates a new Search object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchWithDefaults() *Search {
	this := Search{}
	var skip int32 = 0
	this.Skip = skip
	var limit int32 = 20
	this.Limit = limit
	return &this
}

// GetTerm returns the Term field value if set, zero value otherwise.
func (o *Search) GetTerm() string {
	if o == nil || IsNil(o.Term) {
		var ret string
		return ret
	}
	return *o.Term
}

// GetTermOk returns a tuple with the Term field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Search) GetTermOk() (*string, bool) {
	if o == nil || IsNil(o.Term) {
		return nil, false
	}
	return o.Term, true
}

// HasTerm returns a boolean if a field has been set.
func (o *Search) HasTerm() bool {
	if o != nil && !IsNil(o.Term) {
		return true
	}

	return false
}

// SetTerm gets a reference to the given string and assigns it to the Term field.
func (o *Search) SetTerm(v string) {
	o.Term = &v
}

// GetSkip returns the Skip field value
func (o *Search) GetSkip() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Skip
}

// GetSkipOk returns a tuple with the Skip field value
// and a boolean to check if the value has been set.
func (o *Search) GetSkipOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Skip, true
}

// SetSkip sets field value
func (o *Search) SetSkip(v int32) {
	o.Skip = v
}

// GetLimit returns the Limit field value
func (o *Search) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *Search) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *Search) SetLimit(v int32) {
	o.Limit = v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *Search) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Search) GetObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *Search) HasObjectType() bool {
	if o != nil && !IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *Search) SetObjectType(v string) {
	o.ObjectType = &v
}

func (o Search) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Search) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Term) {
		toSerialize["term"] = o.Term
	}
	toSerialize["skip"] = o.Skip
	toSerialize["limit"] = o.Limit
	if !IsNil(o.ObjectType) {
		toSerialize["objectType"] = o.ObjectType
	}
	return toSerialize, nil
}

type NullableSearch struct {
	value *Search
	isSet bool
}

func (v NullableSearch) Get() *Search {
	return v.value
}

func (v *NullableSearch) Set(val *Search) {
	v.value = val
	v.isSet = true
}

func (v NullableSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearch(val *Search) *NullableSearch {
	return &NullableSearch{value: val, isSet: true}
}

func (v NullableSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


