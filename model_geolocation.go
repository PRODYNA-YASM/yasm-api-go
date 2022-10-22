/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.3.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Geolocation struct for Geolocation
type Geolocation struct {
	Longitude float32 `json:"longitude"`
	Latitude float32 `json:"latitude"`
}

// NewGeolocation instantiates a new Geolocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeolocation(longitude float32, latitude float32) *Geolocation {
	this := Geolocation{}
	this.Longitude = longitude
	this.Latitude = latitude
	return &this
}

// NewGeolocationWithDefaults instantiates a new Geolocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeolocationWithDefaults() *Geolocation {
	this := Geolocation{}
	return &this
}

// GetLongitude returns the Longitude field value
func (o *Geolocation) GetLongitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *Geolocation) GetLongitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *Geolocation) SetLongitude(v float32) {
	o.Longitude = v
}

// GetLatitude returns the Latitude field value
func (o *Geolocation) GetLatitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *Geolocation) GetLatitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *Geolocation) SetLatitude(v float32) {
	o.Latitude = v
}

func (o Geolocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["longitude"] = o.Longitude
	}
	if true {
		toSerialize["latitude"] = o.Latitude
	}
	return json.Marshal(toSerialize)
}

type NullableGeolocation struct {
	value *Geolocation
	isSet bool
}

func (v NullableGeolocation) Get() *Geolocation {
	return v.value
}

func (v *NullableGeolocation) Set(val *Geolocation) {
	v.value = val
	v.isSet = true
}

func (v NullableGeolocation) IsSet() bool {
	return v.isSet
}

func (v *NullableGeolocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeolocation(val *Geolocation) *NullableGeolocation {
	return &NullableGeolocation{value: val, isSet: true}
}

func (v NullableGeolocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeolocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


