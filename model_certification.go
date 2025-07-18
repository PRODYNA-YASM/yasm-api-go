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

// checks if the Certification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Certification{}

// Certification struct for Certification
type Certification struct {
	NamedDomainModel
	Synonyms []string `json:"synonyms,omitempty"`
	Description *string `json:"description,omitempty"`
	Validity *string `json:"validity,omitempty"`
}

// NewCertification instantiates a new Certification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertification(id string, name string) *Certification {
	this := Certification{}
	this.Id = id
	this.Name = name
	return &this
}

// NewCertificationWithDefaults instantiates a new Certification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificationWithDefaults() *Certification {
	this := Certification{}
	return &this
}

// GetSynonyms returns the Synonyms field value if set, zero value otherwise.
func (o *Certification) GetSynonyms() []string {
	if o == nil || IsNil(o.Synonyms) {
		var ret []string
		return ret
	}
	return o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certification) GetSynonymsOk() ([]string, bool) {
	if o == nil || IsNil(o.Synonyms) {
		return nil, false
	}
	return o.Synonyms, true
}

// HasSynonyms returns a boolean if a field has been set.
func (o *Certification) HasSynonyms() bool {
	if o != nil && !IsNil(o.Synonyms) {
		return true
	}

	return false
}

// SetSynonyms gets a reference to the given []string and assigns it to the Synonyms field.
func (o *Certification) SetSynonyms(v []string) {
	o.Synonyms = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Certification) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certification) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Certification) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Certification) SetDescription(v string) {
	o.Description = &v
}

// GetValidity returns the Validity field value if set, zero value otherwise.
func (o *Certification) GetValidity() string {
	if o == nil || IsNil(o.Validity) {
		var ret string
		return ret
	}
	return *o.Validity
}

// GetValidityOk returns a tuple with the Validity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certification) GetValidityOk() (*string, bool) {
	if o == nil || IsNil(o.Validity) {
		return nil, false
	}
	return o.Validity, true
}

// HasValidity returns a boolean if a field has been set.
func (o *Certification) HasValidity() bool {
	if o != nil && !IsNil(o.Validity) {
		return true
	}

	return false
}

// SetValidity gets a reference to the given string and assigns it to the Validity field.
func (o *Certification) SetValidity(v string) {
	o.Validity = &v
}

func (o Certification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Certification) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Validity) {
		toSerialize["validity"] = o.Validity
	}
	return toSerialize, nil
}

type NullableCertification struct {
	value *Certification
	isSet bool
}

func (v NullableCertification) Get() *Certification {
	return v.value
}

func (v *NullableCertification) Set(val *Certification) {
	v.value = val
	v.isSet = true
}

func (v NullableCertification) IsSet() bool {
	return v.isSet
}

func (v *NullableCertification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertification(val *Certification) *NullableCertification {
	return &NullableCertification{value: val, isSet: true}
}

func (v NullableCertification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


