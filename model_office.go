/*
 * YASM (Yet Another Skill Management) API
 *
 * This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.
 *
 * API version: 0.6.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Office struct for Office
type Office struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Geolocation *Geolocation `json:"geolocation,omitempty"`
}

// NewOffice instantiates a new Office object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOffice(id string, name string) *Office {
	this := Office{}
	this.Id = id
	this.Name = name
	return &this
}

// NewOfficeWithDefaults instantiates a new Office object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfficeWithDefaults() *Office {
	this := Office{}
	return &this
}

// GetId returns the Id field value
func (o *Office) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Office) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Office) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Office) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Office) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Office) SetName(v string) {
	o.Name = v
}

// GetGeolocation returns the Geolocation field value if set, zero value otherwise.
func (o *Office) GetGeolocation() Geolocation {
	if o == nil || o.Geolocation == nil {
		var ret Geolocation
		return ret
	}
	return *o.Geolocation
}

// GetGeolocationOk returns a tuple with the Geolocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Office) GetGeolocationOk() (*Geolocation, bool) {
	if o == nil || o.Geolocation == nil {
		return nil, false
	}
	return o.Geolocation, true
}

// HasGeolocation returns a boolean if a field has been set.
func (o *Office) HasGeolocation() bool {
	if o != nil && o.Geolocation != nil {
		return true
	}

	return false
}

// SetGeolocation gets a reference to the given Geolocation and assigns it to the Geolocation field.
func (o *Office) SetGeolocation(v Geolocation) {
	o.Geolocation = &v
}

func (o Office) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Geolocation != nil {
		toSerialize["geolocation"] = o.Geolocation
	}
	return json.Marshal(toSerialize)
}

type NullableOffice struct {
	value *Office
	isSet bool
}

func (v NullableOffice) Get() *Office {
	return v.value
}

func (v *NullableOffice) Set(val *Office) {
	v.value = val
	v.isSet = true
}

func (v NullableOffice) IsSet() bool {
	return v.isSet
}

func (v *NullableOffice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOffice(val *Office) *NullableOffice {
	return &NullableOffice{value: val, isSet: true}
}

func (v NullableOffice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOffice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


