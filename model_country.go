/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.74.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the Country type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Country{}

// Country struct for Country
type Country struct {
	NamedDomainModel
	// base64 encoded image
	Picture *string `json:"picture,omitempty"`
}

// NewCountry instantiates a new Country object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountry(id string, name string) *Country {
	this := Country{}
	this.Id = id
	this.Name = name
	return &this
}

// NewCountryWithDefaults instantiates a new Country object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountryWithDefaults() *Country {
	this := Country{}
	return &this
}

// GetPicture returns the Picture field value if set, zero value otherwise.
func (o *Country) GetPicture() string {
	if o == nil || IsNil(o.Picture) {
		var ret string
		return ret
	}
	return *o.Picture
}

// GetPictureOk returns a tuple with the Picture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Country) GetPictureOk() (*string, bool) {
	if o == nil || IsNil(o.Picture) {
		return nil, false
	}
	return o.Picture, true
}

// HasPicture returns a boolean if a field has been set.
func (o *Country) HasPicture() bool {
	if o != nil && !IsNil(o.Picture) {
		return true
	}

	return false
}

// SetPicture gets a reference to the given string and assigns it to the Picture field.
func (o *Country) SetPicture(v string) {
	o.Picture = &v
}

func (o Country) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Country) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedNamedDomainModel, errNamedDomainModel := json.Marshal(o.NamedDomainModel)
	if errNamedDomainModel != nil {
		return map[string]interface{}{}, errNamedDomainModel
	}
	errNamedDomainModel = json.Unmarshal([]byte(serializedNamedDomainModel), &toSerialize)
	if errNamedDomainModel != nil {
		return map[string]interface{}{}, errNamedDomainModel
	}
	if !IsNil(o.Picture) {
		toSerialize["picture"] = o.Picture
	}
	return toSerialize, nil
}

type NullableCountry struct {
	value *Country
	isSet bool
}

func (v NullableCountry) Get() *Country {
	return v.value
}

func (v *NullableCountry) Set(val *Country) {
	v.value = val
	v.isSet = true
}

func (v NullableCountry) IsSet() bool {
	return v.isSet
}

func (v *NullableCountry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountry(val *Country) *NullableCountry {
	return &NullableCountry{value: val, isSet: true}
}

func (v NullableCountry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


