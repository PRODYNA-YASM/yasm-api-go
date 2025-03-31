/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.72.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PagedAwards type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagedAwards{}

// PagedAwards struct for PagedAwards
type PagedAwards struct {
	Page
	Awards []AwardDetails `json:"awards,omitempty"`
}

// NewPagedAwards instantiates a new PagedAwards object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedAwards() *PagedAwards {
	this := PagedAwards{}
	return &this
}

// NewPagedAwardsWithDefaults instantiates a new PagedAwards object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedAwardsWithDefaults() *PagedAwards {
	this := PagedAwards{}
	return &this
}

// GetAwards returns the Awards field value if set, zero value otherwise.
func (o *PagedAwards) GetAwards() []AwardDetails {
	if o == nil || IsNil(o.Awards) {
		var ret []AwardDetails
		return ret
	}
	return o.Awards
}

// GetAwardsOk returns a tuple with the Awards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedAwards) GetAwardsOk() ([]AwardDetails, bool) {
	if o == nil || IsNil(o.Awards) {
		return nil, false
	}
	return o.Awards, true
}

// HasAwards returns a boolean if a field has been set.
func (o *PagedAwards) HasAwards() bool {
	if o != nil && !IsNil(o.Awards) {
		return true
	}

	return false
}

// SetAwards gets a reference to the given []AwardDetails and assigns it to the Awards field.
func (o *PagedAwards) SetAwards(v []AwardDetails) {
	o.Awards = v
}

func (o PagedAwards) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagedAwards) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPage, errPage := json.Marshal(o.Page)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	errPage = json.Unmarshal([]byte(serializedPage), &toSerialize)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	if !IsNil(o.Awards) {
		toSerialize["awards"] = o.Awards
	}
	return toSerialize, nil
}

type NullablePagedAwards struct {
	value *PagedAwards
	isSet bool
}

func (v NullablePagedAwards) Get() *PagedAwards {
	return v.value
}

func (v *NullablePagedAwards) Set(val *PagedAwards) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedAwards) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedAwards) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedAwards(val *PagedAwards) *NullablePagedAwards {
	return &NullablePagedAwards{value: val, isSet: true}
}

func (v NullablePagedAwards) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedAwards) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


