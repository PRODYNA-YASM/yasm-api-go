/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PageWithDirectHitCount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageWithDirectHitCount{}

// PageWithDirectHitCount struct for PageWithDirectHitCount
type PageWithDirectHitCount struct {
	Page
	DirectHits *int32 `json:"directHits,omitempty"`
}

// NewPageWithDirectHitCount instantiates a new PageWithDirectHitCount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageWithDirectHitCount() *PageWithDirectHitCount {
	this := PageWithDirectHitCount{}
	return &this
}

// NewPageWithDirectHitCountWithDefaults instantiates a new PageWithDirectHitCount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageWithDirectHitCountWithDefaults() *PageWithDirectHitCount {
	this := PageWithDirectHitCount{}
	return &this
}

// GetDirectHits returns the DirectHits field value if set, zero value otherwise.
func (o *PageWithDirectHitCount) GetDirectHits() int32 {
	if o == nil || IsNil(o.DirectHits) {
		var ret int32
		return ret
	}
	return *o.DirectHits
}

// GetDirectHitsOk returns a tuple with the DirectHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageWithDirectHitCount) GetDirectHitsOk() (*int32, bool) {
	if o == nil || IsNil(o.DirectHits) {
		return nil, false
	}
	return o.DirectHits, true
}

// HasDirectHits returns a boolean if a field has been set.
func (o *PageWithDirectHitCount) HasDirectHits() bool {
	if o != nil && !IsNil(o.DirectHits) {
		return true
	}

	return false
}

// SetDirectHits gets a reference to the given int32 and assigns it to the DirectHits field.
func (o *PageWithDirectHitCount) SetDirectHits(v int32) {
	o.DirectHits = &v
}

func (o PageWithDirectHitCount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageWithDirectHitCount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPage, errPage := json.Marshal(o.Page)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	errPage = json.Unmarshal([]byte(serializedPage), &toSerialize)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	if !IsNil(o.DirectHits) {
		toSerialize["directHits"] = o.DirectHits
	}
	return toSerialize, nil
}

type NullablePageWithDirectHitCount struct {
	value *PageWithDirectHitCount
	isSet bool
}

func (v NullablePageWithDirectHitCount) Get() *PageWithDirectHitCount {
	return v.value
}

func (v *NullablePageWithDirectHitCount) Set(val *PageWithDirectHitCount) {
	v.value = val
	v.isSet = true
}

func (v NullablePageWithDirectHitCount) IsSet() bool {
	return v.isSet
}

func (v *NullablePageWithDirectHitCount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageWithDirectHitCount(val *PageWithDirectHitCount) *NullablePageWithDirectHitCount {
	return &NullablePageWithDirectHitCount{value: val, isSet: true}
}

func (v NullablePageWithDirectHitCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageWithDirectHitCount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


