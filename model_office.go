/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the Office type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Office{}

// Office struct for Office
type Office struct {
	NamedDomainModel
	Synonyms []string `json:"synonyms,omitempty"`
	Location *string `json:"location,omitempty"`
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

// GetSynonyms returns the Synonyms field value if set, zero value otherwise.
func (o *Office) GetSynonyms() []string {
	if o == nil || IsNil(o.Synonyms) {
		var ret []string
		return ret
	}
	return o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Office) GetSynonymsOk() ([]string, bool) {
	if o == nil || IsNil(o.Synonyms) {
		return nil, false
	}
	return o.Synonyms, true
}

// HasSynonyms returns a boolean if a field has been set.
func (o *Office) HasSynonyms() bool {
	if o != nil && !IsNil(o.Synonyms) {
		return true
	}

	return false
}

// SetSynonyms gets a reference to the given []string and assigns it to the Synonyms field.
func (o *Office) SetSynonyms(v []string) {
	o.Synonyms = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *Office) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Office) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *Office) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *Office) SetLocation(v string) {
	o.Location = &v
}

// GetGeolocation returns the Geolocation field value if set, zero value otherwise.
func (o *Office) GetGeolocation() Geolocation {
	if o == nil || IsNil(o.Geolocation) {
		var ret Geolocation
		return ret
	}
	return *o.Geolocation
}

// GetGeolocationOk returns a tuple with the Geolocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Office) GetGeolocationOk() (*Geolocation, bool) {
	if o == nil || IsNil(o.Geolocation) {
		return nil, false
	}
	return o.Geolocation, true
}

// HasGeolocation returns a boolean if a field has been set.
func (o *Office) HasGeolocation() bool {
	if o != nil && !IsNil(o.Geolocation) {
		return true
	}

	return false
}

// SetGeolocation gets a reference to the given Geolocation and assigns it to the Geolocation field.
func (o *Office) SetGeolocation(v Geolocation) {
	o.Geolocation = &v
}

func (o Office) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Office) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedNamedDomainModel, errNamedDomainModel := json.Marshal(o.NamedDomainModel)
	if errNamedDomainModel != nil {
		return map[string]interface{}{}, errNamedDomainModel
	}
	errNamedDomainModel = json.Unmarshal([]byte(serializedNamedDomainModel), &toSerialize)
	if errNamedDomainModel != nil {
		return map[string]interface{}{}, errNamedDomainModel
	}
	if !IsNil(o.Synonyms) {
		toSerialize["synonyms"] = o.Synonyms
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Geolocation) {
		toSerialize["geolocation"] = o.Geolocation
	}
	return toSerialize, nil
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


