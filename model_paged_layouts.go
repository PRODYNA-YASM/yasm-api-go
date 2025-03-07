/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.70.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PagedLayouts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagedLayouts{}

// PagedLayouts struct for PagedLayouts
type PagedLayouts struct {
	Page
	Layouts []LayoutDetails `json:"layouts,omitempty"`
}

// NewPagedLayouts instantiates a new PagedLayouts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedLayouts() *PagedLayouts {
	this := PagedLayouts{}
	return &this
}

// NewPagedLayoutsWithDefaults instantiates a new PagedLayouts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedLayoutsWithDefaults() *PagedLayouts {
	this := PagedLayouts{}
	return &this
}

// GetLayouts returns the Layouts field value if set, zero value otherwise.
func (o *PagedLayouts) GetLayouts() []LayoutDetails {
	if o == nil || IsNil(o.Layouts) {
		var ret []LayoutDetails
		return ret
	}
	return o.Layouts
}

// GetLayoutsOk returns a tuple with the Layouts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedLayouts) GetLayoutsOk() ([]LayoutDetails, bool) {
	if o == nil || IsNil(o.Layouts) {
		return nil, false
	}
	return o.Layouts, true
}

// HasLayouts returns a boolean if a field has been set.
func (o *PagedLayouts) HasLayouts() bool {
	if o != nil && !IsNil(o.Layouts) {
		return true
	}

	return false
}

// SetLayouts gets a reference to the given []LayoutDetails and assigns it to the Layouts field.
func (o *PagedLayouts) SetLayouts(v []LayoutDetails) {
	o.Layouts = v
}

func (o PagedLayouts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagedLayouts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPage, errPage := json.Marshal(o.Page)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	errPage = json.Unmarshal([]byte(serializedPage), &toSerialize)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	if !IsNil(o.Layouts) {
		toSerialize["layouts"] = o.Layouts
	}
	return toSerialize, nil
}

type NullablePagedLayouts struct {
	value *PagedLayouts
	isSet bool
}

func (v NullablePagedLayouts) Get() *PagedLayouts {
	return v.value
}

func (v *NullablePagedLayouts) Set(val *PagedLayouts) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedLayouts) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedLayouts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedLayouts(val *PagedLayouts) *NullablePagedLayouts {
	return &NullablePagedLayouts{value: val, isSet: true}
}

func (v NullablePagedLayouts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedLayouts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


