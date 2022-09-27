/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PersonSearchProjectsInner struct for PersonSearchProjectsInner
type PersonSearchProjectsInner struct {
	Id string `json:"id"`
	ParticipationInMonth *MinMax `json:"participationInMonth,omitempty"`
	InvolevedOfficeIds []string `json:"involevedOfficeIds,omitempty"`
}

// NewPersonSearchProjectsInner instantiates a new PersonSearchProjectsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonSearchProjectsInner(id string) *PersonSearchProjectsInner {
	this := PersonSearchProjectsInner{}
	this.Id = id
	return &this
}

// NewPersonSearchProjectsInnerWithDefaults instantiates a new PersonSearchProjectsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonSearchProjectsInnerWithDefaults() *PersonSearchProjectsInner {
	this := PersonSearchProjectsInner{}
	return &this
}

// GetId returns the Id field value
func (o *PersonSearchProjectsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PersonSearchProjectsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PersonSearchProjectsInner) SetId(v string) {
	o.Id = v
}

// GetParticipationInMonth returns the ParticipationInMonth field value if set, zero value otherwise.
func (o *PersonSearchProjectsInner) GetParticipationInMonth() MinMax {
	if o == nil || o.ParticipationInMonth == nil {
		var ret MinMax
		return ret
	}
	return *o.ParticipationInMonth
}

// GetParticipationInMonthOk returns a tuple with the ParticipationInMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearchProjectsInner) GetParticipationInMonthOk() (*MinMax, bool) {
	if o == nil || o.ParticipationInMonth == nil {
		return nil, false
	}
	return o.ParticipationInMonth, true
}

// HasParticipationInMonth returns a boolean if a field has been set.
func (o *PersonSearchProjectsInner) HasParticipationInMonth() bool {
	if o != nil && o.ParticipationInMonth != nil {
		return true
	}

	return false
}

// SetParticipationInMonth gets a reference to the given MinMax and assigns it to the ParticipationInMonth field.
func (o *PersonSearchProjectsInner) SetParticipationInMonth(v MinMax) {
	o.ParticipationInMonth = &v
}

// GetInvolevedOfficeIds returns the InvolevedOfficeIds field value if set, zero value otherwise.
func (o *PersonSearchProjectsInner) GetInvolevedOfficeIds() []string {
	if o == nil || o.InvolevedOfficeIds == nil {
		var ret []string
		return ret
	}
	return o.InvolevedOfficeIds
}

// GetInvolevedOfficeIdsOk returns a tuple with the InvolevedOfficeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearchProjectsInner) GetInvolevedOfficeIdsOk() ([]string, bool) {
	if o == nil || o.InvolevedOfficeIds == nil {
		return nil, false
	}
	return o.InvolevedOfficeIds, true
}

// HasInvolevedOfficeIds returns a boolean if a field has been set.
func (o *PersonSearchProjectsInner) HasInvolevedOfficeIds() bool {
	if o != nil && o.InvolevedOfficeIds != nil {
		return true
	}

	return false
}

// SetInvolevedOfficeIds gets a reference to the given []string and assigns it to the InvolevedOfficeIds field.
func (o *PersonSearchProjectsInner) SetInvolevedOfficeIds(v []string) {
	o.InvolevedOfficeIds = v
}

func (o PersonSearchProjectsInner) MarshalJSON() ([]byte, error) {
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

type NullablePersonSearchProjectsInner struct {
	value *PersonSearchProjectsInner
	isSet bool
}

func (v NullablePersonSearchProjectsInner) Get() *PersonSearchProjectsInner {
	return v.value
}

func (v *NullablePersonSearchProjectsInner) Set(val *PersonSearchProjectsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonSearchProjectsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonSearchProjectsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonSearchProjectsInner(val *PersonSearchProjectsInner) *NullablePersonSearchProjectsInner {
	return &NullablePersonSearchProjectsInner{value: val, isSet: true}
}

func (v NullablePersonSearchProjectsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonSearchProjectsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


