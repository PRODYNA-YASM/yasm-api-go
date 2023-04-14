/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.3.17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PersonCertificationFilterAllOf struct for PersonCertificationFilterAllOf
type PersonCertificationFilterAllOf struct {
	// Include employees who started the certification
	StartedCertificaiton *bool `json:"startedCertificaiton,omitempty"`
}

// NewPersonCertificationFilterAllOf instantiates a new PersonCertificationFilterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonCertificationFilterAllOf() *PersonCertificationFilterAllOf {
	this := PersonCertificationFilterAllOf{}
	return &this
}

// NewPersonCertificationFilterAllOfWithDefaults instantiates a new PersonCertificationFilterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonCertificationFilterAllOfWithDefaults() *PersonCertificationFilterAllOf {
	this := PersonCertificationFilterAllOf{}
	return &this
}

// GetStartedCertificaiton returns the StartedCertificaiton field value if set, zero value otherwise.
func (o *PersonCertificationFilterAllOf) GetStartedCertificaiton() bool {
	if o == nil || o.StartedCertificaiton == nil {
		var ret bool
		return ret
	}
	return *o.StartedCertificaiton
}

// GetStartedCertificaitonOk returns a tuple with the StartedCertificaiton field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonCertificationFilterAllOf) GetStartedCertificaitonOk() (*bool, bool) {
	if o == nil || o.StartedCertificaiton == nil {
		return nil, false
	}
	return o.StartedCertificaiton, true
}

// HasStartedCertificaiton returns a boolean if a field has been set.
func (o *PersonCertificationFilterAllOf) HasStartedCertificaiton() bool {
	if o != nil && o.StartedCertificaiton != nil {
		return true
	}

	return false
}

// SetStartedCertificaiton gets a reference to the given bool and assigns it to the StartedCertificaiton field.
func (o *PersonCertificationFilterAllOf) SetStartedCertificaiton(v bool) {
	o.StartedCertificaiton = &v
}

func (o PersonCertificationFilterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StartedCertificaiton != nil {
		toSerialize["startedCertificaiton"] = o.StartedCertificaiton
	}
	return json.Marshal(toSerialize)
}

type NullablePersonCertificationFilterAllOf struct {
	value *PersonCertificationFilterAllOf
	isSet bool
}

func (v NullablePersonCertificationFilterAllOf) Get() *PersonCertificationFilterAllOf {
	return v.value
}

func (v *NullablePersonCertificationFilterAllOf) Set(val *PersonCertificationFilterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonCertificationFilterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonCertificationFilterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonCertificationFilterAllOf(val *PersonCertificationFilterAllOf) *NullablePersonCertificationFilterAllOf {
	return &NullablePersonCertificationFilterAllOf{value: val, isSet: true}
}

func (v NullablePersonCertificationFilterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonCertificationFilterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


