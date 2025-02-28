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

// checks if the CertificationView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificationView{}

// CertificationView struct for CertificationView
type CertificationView struct {
	Certification *Certification `json:"certification,omitempty"`
	Organization *Organization `json:"organization,omitempty"`
}

// NewCertificationView instantiates a new CertificationView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificationView() *CertificationView {
	this := CertificationView{}
	return &this
}

// NewCertificationViewWithDefaults instantiates a new CertificationView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificationViewWithDefaults() *CertificationView {
	this := CertificationView{}
	return &this
}

// GetCertification returns the Certification field value if set, zero value otherwise.
func (o *CertificationView) GetCertification() Certification {
	if o == nil || IsNil(o.Certification) {
		var ret Certification
		return ret
	}
	return *o.Certification
}

// GetCertificationOk returns a tuple with the Certification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificationView) GetCertificationOk() (*Certification, bool) {
	if o == nil || IsNil(o.Certification) {
		return nil, false
	}
	return o.Certification, true
}

// HasCertification returns a boolean if a field has been set.
func (o *CertificationView) HasCertification() bool {
	if o != nil && !IsNil(o.Certification) {
		return true
	}

	return false
}

// SetCertification gets a reference to the given Certification and assigns it to the Certification field.
func (o *CertificationView) SetCertification(v Certification) {
	o.Certification = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *CertificationView) GetOrganization() Organization {
	if o == nil || IsNil(o.Organization) {
		var ret Organization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificationView) GetOrganizationOk() (*Organization, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *CertificationView) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Organization and assigns it to the Organization field.
func (o *CertificationView) SetOrganization(v Organization) {
	o.Organization = &v
}

func (o CertificationView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificationView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Certification) {
		toSerialize["certification"] = o.Certification
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	return toSerialize, nil
}

type NullableCertificationView struct {
	value *CertificationView
	isSet bool
}

func (v NullableCertificationView) Get() *CertificationView {
	return v.value
}

func (v *NullableCertificationView) Set(val *CertificationView) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificationView) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificationView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificationView(val *CertificationView) *NullableCertificationView {
	return &NullableCertificationView{value: val, isSet: true}
}

func (v NullableCertificationView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificationView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


