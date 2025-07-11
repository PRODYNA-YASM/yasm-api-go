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

// checks if the Profile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Profile{}

// Profile struct for Profile
type Profile struct {
	NamedDomainModel
	Synonyms []string `json:"synonyms,omitempty"`
	Confidentiality *Confidentiality `json:"confidentiality,omitempty"`
}

// NewProfile instantiates a new Profile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfile(id string, name string) *Profile {
	this := Profile{}
	this.Id = id
	this.Name = name
	return &this
}

// NewProfileWithDefaults instantiates a new Profile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileWithDefaults() *Profile {
	this := Profile{}
	return &this
}

// GetSynonyms returns the Synonyms field value if set, zero value otherwise.
func (o *Profile) GetSynonyms() []string {
	if o == nil || IsNil(o.Synonyms) {
		var ret []string
		return ret
	}
	return o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetSynonymsOk() ([]string, bool) {
	if o == nil || IsNil(o.Synonyms) {
		return nil, false
	}
	return o.Synonyms, true
}

// HasSynonyms returns a boolean if a field has been set.
func (o *Profile) HasSynonyms() bool {
	if o != nil && !IsNil(o.Synonyms) {
		return true
	}

	return false
}

// SetSynonyms gets a reference to the given []string and assigns it to the Synonyms field.
func (o *Profile) SetSynonyms(v []string) {
	o.Synonyms = v
}

// GetConfidentiality returns the Confidentiality field value if set, zero value otherwise.
func (o *Profile) GetConfidentiality() Confidentiality {
	if o == nil || IsNil(o.Confidentiality) {
		var ret Confidentiality
		return ret
	}
	return *o.Confidentiality
}

// GetConfidentialityOk returns a tuple with the Confidentiality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetConfidentialityOk() (*Confidentiality, bool) {
	if o == nil || IsNil(o.Confidentiality) {
		return nil, false
	}
	return o.Confidentiality, true
}

// HasConfidentiality returns a boolean if a field has been set.
func (o *Profile) HasConfidentiality() bool {
	if o != nil && !IsNil(o.Confidentiality) {
		return true
	}

	return false
}

// SetConfidentiality gets a reference to the given Confidentiality and assigns it to the Confidentiality field.
func (o *Profile) SetConfidentiality(v Confidentiality) {
	o.Confidentiality = &v
}

func (o Profile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Profile) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Confidentiality) {
		toSerialize["confidentiality"] = o.Confidentiality
	}
	return toSerialize, nil
}

type NullableProfile struct {
	value *Profile
	isSet bool
}

func (v NullableProfile) Get() *Profile {
	return v.value
}

func (v *NullableProfile) Set(val *Profile) {
	v.value = val
	v.isSet = true
}

func (v NullableProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfile(val *Profile) *NullableProfile {
	return &NullableProfile{value: val, isSet: true}
}

func (v NullableProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


