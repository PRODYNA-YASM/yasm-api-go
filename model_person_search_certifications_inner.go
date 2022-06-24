/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PersonSearchCertificationsInner struct for PersonSearchCertificationsInner
type PersonSearchCertificationsInner struct {
	Id string `json:"id"`
	// Include employees who started the certification
	StartedCertificaiton *bool `json:"startedCertificaiton,omitempty"`
}

// NewPersonSearchCertificationsInner instantiates a new PersonSearchCertificationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonSearchCertificationsInner(id string) *PersonSearchCertificationsInner {
	this := PersonSearchCertificationsInner{}
	this.Id = id
	return &this
}

// NewPersonSearchCertificationsInnerWithDefaults instantiates a new PersonSearchCertificationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonSearchCertificationsInnerWithDefaults() *PersonSearchCertificationsInner {
	this := PersonSearchCertificationsInner{}
	return &this
}

// GetId returns the Id field value
func (o *PersonSearchCertificationsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PersonSearchCertificationsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PersonSearchCertificationsInner) SetId(v string) {
	o.Id = v
}

// GetStartedCertificaiton returns the StartedCertificaiton field value if set, zero value otherwise.
func (o *PersonSearchCertificationsInner) GetStartedCertificaiton() bool {
	if o == nil || o.StartedCertificaiton == nil {
		var ret bool
		return ret
	}
	return *o.StartedCertificaiton
}

// GetStartedCertificaitonOk returns a tuple with the StartedCertificaiton field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearchCertificationsInner) GetStartedCertificaitonOk() (*bool, bool) {
	if o == nil || o.StartedCertificaiton == nil {
		return nil, false
	}
	return o.StartedCertificaiton, true
}

// HasStartedCertificaiton returns a boolean if a field has been set.
func (o *PersonSearchCertificationsInner) HasStartedCertificaiton() bool {
	if o != nil && o.StartedCertificaiton != nil {
		return true
	}

	return false
}

// SetStartedCertificaiton gets a reference to the given bool and assigns it to the StartedCertificaiton field.
func (o *PersonSearchCertificationsInner) SetStartedCertificaiton(v bool) {
	o.StartedCertificaiton = &v
}

func (o PersonSearchCertificationsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.StartedCertificaiton != nil {
		toSerialize["startedCertificaiton"] = o.StartedCertificaiton
	}
	return json.Marshal(toSerialize)
}

type NullablePersonSearchCertificationsInner struct {
	value *PersonSearchCertificationsInner
	isSet bool
}

func (v NullablePersonSearchCertificationsInner) Get() *PersonSearchCertificationsInner {
	return v.value
}

func (v *NullablePersonSearchCertificationsInner) Set(val *PersonSearchCertificationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonSearchCertificationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonSearchCertificationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonSearchCertificationsInner(val *PersonSearchCertificationsInner) *NullablePersonSearchCertificationsInner {
	return &NullablePersonSearchCertificationsInner{value: val, isSet: true}
}

func (v NullablePersonSearchCertificationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonSearchCertificationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


