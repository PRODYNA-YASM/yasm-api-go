/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 0.9.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PersonProjectFilter struct for PersonProjectFilter
type PersonProjectFilter struct {
	Id string `json:"id"`
	ParticipationInMonth *MinMax `json:"participationInMonth,omitempty"`
	InvolevedOfficeIds *[]string `json:"involevedOfficeIds,omitempty"`
}

// NewPersonProjectFilter instantiates a new PersonProjectFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonProjectFilter(id string) *PersonProjectFilter {
	this := PersonProjectFilter{}
	this.Id = id
	return &this
}

// NewPersonProjectFilterWithDefaults instantiates a new PersonProjectFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonProjectFilterWithDefaults() *PersonProjectFilter {
	this := PersonProjectFilter{}
	return &this
}

// GetId returns the Id field value
func (o *PersonProjectFilter) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PersonProjectFilter) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PersonProjectFilter) SetId(v string) {
	o.Id = v
}

// GetParticipationInMonth returns the ParticipationInMonth field value if set, zero value otherwise.
func (o *PersonProjectFilter) GetParticipationInMonth() MinMax {
	if o == nil || o.ParticipationInMonth == nil {
		var ret MinMax
		return ret
	}
	return *o.ParticipationInMonth
}

// GetParticipationInMonthOk returns a tuple with the ParticipationInMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonProjectFilter) GetParticipationInMonthOk() (*MinMax, bool) {
	if o == nil || o.ParticipationInMonth == nil {
		return nil, false
	}
	return o.ParticipationInMonth, true
}

// HasParticipationInMonth returns a boolean if a field has been set.
func (o *PersonProjectFilter) HasParticipationInMonth() bool {
	if o != nil && o.ParticipationInMonth != nil {
		return true
	}

	return false
}

// SetParticipationInMonth gets a reference to the given MinMax and assigns it to the ParticipationInMonth field.
func (o *PersonProjectFilter) SetParticipationInMonth(v MinMax) {
	o.ParticipationInMonth = &v
}

// GetInvolevedOfficeIds returns the InvolevedOfficeIds field value if set, zero value otherwise.
func (o *PersonProjectFilter) GetInvolevedOfficeIds() []string {
	if o == nil || o.InvolevedOfficeIds == nil {
		var ret []string
		return ret
	}
	return *o.InvolevedOfficeIds
}

// GetInvolevedOfficeIdsOk returns a tuple with the InvolevedOfficeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonProjectFilter) GetInvolevedOfficeIdsOk() (*[]string, bool) {
	if o == nil || o.InvolevedOfficeIds == nil {
		return nil, false
	}
	return o.InvolevedOfficeIds, true
}

// HasInvolevedOfficeIds returns a boolean if a field has been set.
func (o *PersonProjectFilter) HasInvolevedOfficeIds() bool {
	if o != nil && o.InvolevedOfficeIds != nil {
		return true
	}

	return false
}

// SetInvolevedOfficeIds gets a reference to the given []string and assigns it to the InvolevedOfficeIds field.
func (o *PersonProjectFilter) SetInvolevedOfficeIds(v []string) {
	o.InvolevedOfficeIds = &v
}

func (o PersonProjectFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.ParticipationInMonth != nil {
		toSerialize["participationInMonth"] = o.ParticipationInMonth
	}
	if o.InvolevedOfficeIds != nil {
		toSerialize["involevedOfficeIds"] = o.InvolevedOfficeIds
	}
	return json.Marshal(toSerialize)
}

type NullablePersonProjectFilter struct {
	value *PersonProjectFilter
	isSet bool
}

func (v NullablePersonProjectFilter) Get() *PersonProjectFilter {
	return v.value
}

func (v *NullablePersonProjectFilter) Set(val *PersonProjectFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonProjectFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonProjectFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonProjectFilter(val *PersonProjectFilter) *NullablePersonProjectFilter {
	return &NullablePersonProjectFilter{value: val, isSet: true}
}

func (v NullablePersonProjectFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonProjectFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


