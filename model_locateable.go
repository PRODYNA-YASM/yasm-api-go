/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the Locateable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Locateable{}

// Locateable struct for Locateable
type Locateable struct {
	Location *string `json:"location,omitempty"`
	Geolocation *Geolocation `json:"geolocation,omitempty"`
}

// NewLocateable instantiates a new Locateable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocateable() *Locateable {
	this := Locateable{}
	return &this
}

// NewLocateableWithDefaults instantiates a new Locateable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocateableWithDefaults() *Locateable {
	this := Locateable{}
	return &this
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *Locateable) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Locateable) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *Locateable) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *Locateable) SetLocation(v string) {
	o.Location = &v
}

// GetGeolocation returns the Geolocation field value if set, zero value otherwise.
func (o *Locateable) GetGeolocation() Geolocation {
	if o == nil || IsNil(o.Geolocation) {
		var ret Geolocation
		return ret
	}
	return *o.Geolocation
}

// GetGeolocationOk returns a tuple with the Geolocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Locateable) GetGeolocationOk() (*Geolocation, bool) {
	if o == nil || IsNil(o.Geolocation) {
		return nil, false
	}
	return o.Geolocation, true
}

// HasGeolocation returns a boolean if a field has been set.
func (o *Locateable) HasGeolocation() bool {
	if o != nil && !IsNil(o.Geolocation) {
		return true
	}

	return false
}

// SetGeolocation gets a reference to the given Geolocation and assigns it to the Geolocation field.
func (o *Locateable) SetGeolocation(v Geolocation) {
	o.Geolocation = &v
}

func (o Locateable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Locateable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Geolocation) {
		toSerialize["geolocation"] = o.Geolocation
	}
	return toSerialize, nil
}

type NullableLocateable struct {
	value *Locateable
	isSet bool
}

func (v NullableLocateable) Get() *Locateable {
	return v.value
}

func (v *NullableLocateable) Set(val *Locateable) {
	v.value = val
	v.isSet = true
}

func (v NullableLocateable) IsSet() bool {
	return v.isSet
}

func (v *NullableLocateable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocateable(val *Locateable) *NullableLocateable {
	return &NullableLocateable{value: val, isSet: true}
}

func (v NullableLocateable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocateable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


