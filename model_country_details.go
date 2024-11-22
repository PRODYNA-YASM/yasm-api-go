/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the CountryDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountryDetails{}

// CountryDetails struct for CountryDetails
type CountryDetails struct {
	Audit *Audit `json:"audit,omitempty"`
	Country *Country `json:"country,omitempty"`
	Languages []Language `json:"languages,omitempty"`
}

// NewCountryDetails instantiates a new CountryDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountryDetails() *CountryDetails {
	this := CountryDetails{}
	return &this
}

// NewCountryDetailsWithDefaults instantiates a new CountryDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountryDetailsWithDefaults() *CountryDetails {
	this := CountryDetails{}
	return &this
}

// GetAudit returns the Audit field value if set, zero value otherwise.
func (o *CountryDetails) GetAudit() Audit {
	if o == nil || IsNil(o.Audit) {
		var ret Audit
		return ret
	}
	return *o.Audit
}

// GetAuditOk returns a tuple with the Audit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryDetails) GetAuditOk() (*Audit, bool) {
	if o == nil || IsNil(o.Audit) {
		return nil, false
	}
	return o.Audit, true
}

// HasAudit returns a boolean if a field has been set.
func (o *CountryDetails) HasAudit() bool {
	if o != nil && !IsNil(o.Audit) {
		return true
	}

	return false
}

// SetAudit gets a reference to the given Audit and assigns it to the Audit field.
func (o *CountryDetails) SetAudit(v Audit) {
	o.Audit = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CountryDetails) GetCountry() Country {
	if o == nil || IsNil(o.Country) {
		var ret Country
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryDetails) GetCountryOk() (*Country, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CountryDetails) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given Country and assigns it to the Country field.
func (o *CountryDetails) SetCountry(v Country) {
	o.Country = &v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *CountryDetails) GetLanguages() []Language {
	if o == nil || IsNil(o.Languages) {
		var ret []Language
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryDetails) GetLanguagesOk() ([]Language, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *CountryDetails) HasLanguages() bool {
	if o != nil && !IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []Language and assigns it to the Languages field.
func (o *CountryDetails) SetLanguages(v []Language) {
	o.Languages = v
}

func (o CountryDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountryDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Audit) {
		toSerialize["audit"] = o.Audit
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Languages) {
		toSerialize["languages"] = o.Languages
	}
	return toSerialize, nil
}

type NullableCountryDetails struct {
	value *CountryDetails
	isSet bool
}

func (v NullableCountryDetails) Get() *CountryDetails {
	return v.value
}

func (v *NullableCountryDetails) Set(val *CountryDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCountryDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCountryDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountryDetails(val *CountryDetails) *NullableCountryDetails {
	return &NullableCountryDetails{value: val, isSet: true}
}

func (v NullableCountryDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountryDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


