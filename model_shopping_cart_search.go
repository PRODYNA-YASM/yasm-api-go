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

// checks if the ShoppingCartSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShoppingCartSearch{}

// ShoppingCartSearch struct for ShoppingCartSearch
type ShoppingCartSearch struct {
	Search
	OwnerIds []string `json:"ownerIds,omitempty"`
	SharedWithIds []string `json:"sharedWithIds,omitempty"`
	PersonIds []string `json:"personIds,omitempty"`
}

// NewShoppingCartSearch instantiates a new ShoppingCartSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShoppingCartSearch(skip int32, limit int32) *ShoppingCartSearch {
	this := ShoppingCartSearch{}
	this.Skip = skip
	this.Limit = limit
	return &this
}

// NewShoppingCartSearchWithDefaults instantiates a new ShoppingCartSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShoppingCartSearchWithDefaults() *ShoppingCartSearch {
	this := ShoppingCartSearch{}
	return &this
}

// GetOwnerIds returns the OwnerIds field value if set, zero value otherwise.
func (o *ShoppingCartSearch) GetOwnerIds() []string {
	if o == nil || IsNil(o.OwnerIds) {
		var ret []string
		return ret
	}
	return o.OwnerIds
}

// GetOwnerIdsOk returns a tuple with the OwnerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShoppingCartSearch) GetOwnerIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OwnerIds) {
		return nil, false
	}
	return o.OwnerIds, true
}

// HasOwnerIds returns a boolean if a field has been set.
func (o *ShoppingCartSearch) HasOwnerIds() bool {
	if o != nil && !IsNil(o.OwnerIds) {
		return true
	}

	return false
}

// SetOwnerIds gets a reference to the given []string and assigns it to the OwnerIds field.
func (o *ShoppingCartSearch) SetOwnerIds(v []string) {
	o.OwnerIds = v
}

// GetSharedWithIds returns the SharedWithIds field value if set, zero value otherwise.
func (o *ShoppingCartSearch) GetSharedWithIds() []string {
	if o == nil || IsNil(o.SharedWithIds) {
		var ret []string
		return ret
	}
	return o.SharedWithIds
}

// GetSharedWithIdsOk returns a tuple with the SharedWithIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShoppingCartSearch) GetSharedWithIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SharedWithIds) {
		return nil, false
	}
	return o.SharedWithIds, true
}

// HasSharedWithIds returns a boolean if a field has been set.
func (o *ShoppingCartSearch) HasSharedWithIds() bool {
	if o != nil && !IsNil(o.SharedWithIds) {
		return true
	}

	return false
}

// SetSharedWithIds gets a reference to the given []string and assigns it to the SharedWithIds field.
func (o *ShoppingCartSearch) SetSharedWithIds(v []string) {
	o.SharedWithIds = v
}

// GetPersonIds returns the PersonIds field value if set, zero value otherwise.
func (o *ShoppingCartSearch) GetPersonIds() []string {
	if o == nil || IsNil(o.PersonIds) {
		var ret []string
		return ret
	}
	return o.PersonIds
}

// GetPersonIdsOk returns a tuple with the PersonIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShoppingCartSearch) GetPersonIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.PersonIds) {
		return nil, false
	}
	return o.PersonIds, true
}

// HasPersonIds returns a boolean if a field has been set.
func (o *ShoppingCartSearch) HasPersonIds() bool {
	if o != nil && !IsNil(o.PersonIds) {
		return true
	}

	return false
}

// SetPersonIds gets a reference to the given []string and assigns it to the PersonIds field.
func (o *ShoppingCartSearch) SetPersonIds(v []string) {
	o.PersonIds = v
}

func (o ShoppingCartSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShoppingCartSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSearch, errSearch := json.Marshal(o.Search)
	if errSearch != nil {
		return map[string]interface{}{}, errSearch
	}
	errSearch = json.Unmarshal([]byte(serializedSearch), &toSerialize)
	if errSearch != nil {
		return map[string]interface{}{}, errSearch
	}
	if !IsNil(o.OwnerIds) {
		toSerialize["ownerIds"] = o.OwnerIds
	}
	if !IsNil(o.SharedWithIds) {
		toSerialize["sharedWithIds"] = o.SharedWithIds
	}
	if !IsNil(o.PersonIds) {
		toSerialize["personIds"] = o.PersonIds
	}
	return toSerialize, nil
}

type NullableShoppingCartSearch struct {
	value *ShoppingCartSearch
	isSet bool
}

func (v NullableShoppingCartSearch) Get() *ShoppingCartSearch {
	return v.value
}

func (v *NullableShoppingCartSearch) Set(val *ShoppingCartSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableShoppingCartSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableShoppingCartSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShoppingCartSearch(val *ShoppingCartSearch) *NullableShoppingCartSearch {
	return &NullableShoppingCartSearch{value: val, isSet: true}
}

func (v NullableShoppingCartSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShoppingCartSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


