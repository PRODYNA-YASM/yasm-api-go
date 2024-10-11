/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PagedOffices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagedOffices{}

// PagedOffices struct for PagedOffices
type PagedOffices struct {
	Page
	Offices []OfficeDetails `json:"offices,omitempty"`
}

// NewPagedOffices instantiates a new PagedOffices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedOffices() *PagedOffices {
	this := PagedOffices{}
	return &this
}

// NewPagedOfficesWithDefaults instantiates a new PagedOffices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedOfficesWithDefaults() *PagedOffices {
	this := PagedOffices{}
	return &this
}

// GetOffices returns the Offices field value if set, zero value otherwise.
func (o *PagedOffices) GetOffices() []OfficeDetails {
	if o == nil || IsNil(o.Offices) {
		var ret []OfficeDetails
		return ret
	}
	return o.Offices
}

// GetOfficesOk returns a tuple with the Offices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedOffices) GetOfficesOk() ([]OfficeDetails, bool) {
	if o == nil || IsNil(o.Offices) {
		return nil, false
	}
	return o.Offices, true
}

// HasOffices returns a boolean if a field has been set.
func (o *PagedOffices) HasOffices() bool {
	if o != nil && !IsNil(o.Offices) {
		return true
	}

	return false
}

// SetOffices gets a reference to the given []OfficeDetails and assigns it to the Offices field.
func (o *PagedOffices) SetOffices(v []OfficeDetails) {
	o.Offices = v
}

func (o PagedOffices) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagedOffices) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPage, errPage := json.Marshal(o.Page)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	errPage = json.Unmarshal([]byte(serializedPage), &toSerialize)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	if !IsNil(o.Offices) {
		toSerialize["offices"] = o.Offices
	}
	return toSerialize, nil
}

type NullablePagedOffices struct {
	value *PagedOffices
	isSet bool
}

func (v NullablePagedOffices) Get() *PagedOffices {
	return v.value
}

func (v *NullablePagedOffices) Set(val *PagedOffices) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedOffices) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedOffices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedOffices(val *PagedOffices) *NullablePagedOffices {
	return &NullablePagedOffices{value: val, isSet: true}
}

func (v NullablePagedOffices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedOffices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


