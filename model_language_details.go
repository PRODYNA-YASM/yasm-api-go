/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.69.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the LanguageDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LanguageDetails{}

// LanguageDetails struct for LanguageDetails
type LanguageDetails struct {
	Audit *Audit `json:"audit,omitempty"`
	Language *Language `json:"language,omitempty"`
	Countries []Country `json:"countries,omitempty"`
}

// NewLanguageDetails instantiates a new LanguageDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLanguageDetails() *LanguageDetails {
	this := LanguageDetails{}
	return &this
}

// NewLanguageDetailsWithDefaults instantiates a new LanguageDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLanguageDetailsWithDefaults() *LanguageDetails {
	this := LanguageDetails{}
	return &this
}

// GetAudit returns the Audit field value if set, zero value otherwise.
func (o *LanguageDetails) GetAudit() Audit {
	if o == nil || IsNil(o.Audit) {
		var ret Audit
		return ret
	}
	return *o.Audit
}

// GetAuditOk returns a tuple with the Audit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LanguageDetails) GetAuditOk() (*Audit, bool) {
	if o == nil || IsNil(o.Audit) {
		return nil, false
	}
	return o.Audit, true
}

// HasAudit returns a boolean if a field has been set.
func (o *LanguageDetails) HasAudit() bool {
	if o != nil && !IsNil(o.Audit) {
		return true
	}

	return false
}

// SetAudit gets a reference to the given Audit and assigns it to the Audit field.
func (o *LanguageDetails) SetAudit(v Audit) {
	o.Audit = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *LanguageDetails) GetLanguage() Language {
	if o == nil || IsNil(o.Language) {
		var ret Language
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LanguageDetails) GetLanguageOk() (*Language, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *LanguageDetails) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given Language and assigns it to the Language field.
func (o *LanguageDetails) SetLanguage(v Language) {
	o.Language = &v
}

// GetCountries returns the Countries field value if set, zero value otherwise.
func (o *LanguageDetails) GetCountries() []Country {
	if o == nil || IsNil(o.Countries) {
		var ret []Country
		return ret
	}
	return o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LanguageDetails) GetCountriesOk() ([]Country, bool) {
	if o == nil || IsNil(o.Countries) {
		return nil, false
	}
	return o.Countries, true
}

// HasCountries returns a boolean if a field has been set.
func (o *LanguageDetails) HasCountries() bool {
	if o != nil && !IsNil(o.Countries) {
		return true
	}

	return false
}

// SetCountries gets a reference to the given []Country and assigns it to the Countries field.
func (o *LanguageDetails) SetCountries(v []Country) {
	o.Countries = v
}

func (o LanguageDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LanguageDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Audit) {
		toSerialize["audit"] = o.Audit
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Countries) {
		toSerialize["countries"] = o.Countries
	}
	return toSerialize, nil
}

type NullableLanguageDetails struct {
	value *LanguageDetails
	isSet bool
}

func (v NullableLanguageDetails) Get() *LanguageDetails {
	return v.value
}

func (v *NullableLanguageDetails) Set(val *LanguageDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableLanguageDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableLanguageDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLanguageDetails(val *LanguageDetails) *NullableLanguageDetails {
	return &NullableLanguageDetails{value: val, isSet: true}
}

func (v NullableLanguageDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLanguageDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


