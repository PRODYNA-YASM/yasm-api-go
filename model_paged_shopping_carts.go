/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.69.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PagedShoppingCarts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagedShoppingCarts{}

// PagedShoppingCarts struct for PagedShoppingCarts
type PagedShoppingCarts struct {
	Page
	ShoppingCarts []ShoppingCartDetail `json:"shoppingCarts,omitempty"`
}

// NewPagedShoppingCarts instantiates a new PagedShoppingCarts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedShoppingCarts() *PagedShoppingCarts {
	this := PagedShoppingCarts{}
	return &this
}

// NewPagedShoppingCartsWithDefaults instantiates a new PagedShoppingCarts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedShoppingCartsWithDefaults() *PagedShoppingCarts {
	this := PagedShoppingCarts{}
	return &this
}

// GetShoppingCarts returns the ShoppingCarts field value if set, zero value otherwise.
func (o *PagedShoppingCarts) GetShoppingCarts() []ShoppingCartDetail {
	if o == nil || IsNil(o.ShoppingCarts) {
		var ret []ShoppingCartDetail
		return ret
	}
	return o.ShoppingCarts
}

// GetShoppingCartsOk returns a tuple with the ShoppingCarts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedShoppingCarts) GetShoppingCartsOk() ([]ShoppingCartDetail, bool) {
	if o == nil || IsNil(o.ShoppingCarts) {
		return nil, false
	}
	return o.ShoppingCarts, true
}

// HasShoppingCarts returns a boolean if a field has been set.
func (o *PagedShoppingCarts) HasShoppingCarts() bool {
	if o != nil && !IsNil(o.ShoppingCarts) {
		return true
	}

	return false
}

// SetShoppingCarts gets a reference to the given []ShoppingCartDetail and assigns it to the ShoppingCarts field.
func (o *PagedShoppingCarts) SetShoppingCarts(v []ShoppingCartDetail) {
	o.ShoppingCarts = v
}

func (o PagedShoppingCarts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagedShoppingCarts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPage, errPage := json.Marshal(o.Page)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	errPage = json.Unmarshal([]byte(serializedPage), &toSerialize)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	if !IsNil(o.ShoppingCarts) {
		toSerialize["shoppingCarts"] = o.ShoppingCarts
	}
	return toSerialize, nil
}

type NullablePagedShoppingCarts struct {
	value *PagedShoppingCarts
	isSet bool
}

func (v NullablePagedShoppingCarts) Get() *PagedShoppingCarts {
	return v.value
}

func (v *NullablePagedShoppingCarts) Set(val *PagedShoppingCarts) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedShoppingCarts) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedShoppingCarts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedShoppingCarts(val *PagedShoppingCarts) *NullablePagedShoppingCarts {
	return &NullablePagedShoppingCarts{value: val, isSet: true}
}

func (v NullablePagedShoppingCarts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedShoppingCarts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


