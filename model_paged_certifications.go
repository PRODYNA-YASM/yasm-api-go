/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.16.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PagedCertifications type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagedCertifications{}

// PagedCertifications struct for PagedCertifications
type PagedCertifications struct {
	Page
	Certifications []CertificationDetails `json:"certifications,omitempty"`
}

// NewPagedCertifications instantiates a new PagedCertifications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedCertifications() *PagedCertifications {
	this := PagedCertifications{}
	return &this
}

// NewPagedCertificationsWithDefaults instantiates a new PagedCertifications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedCertificationsWithDefaults() *PagedCertifications {
	this := PagedCertifications{}
	return &this
}

// GetCertifications returns the Certifications field value if set, zero value otherwise.
func (o *PagedCertifications) GetCertifications() []CertificationDetails {
	if o == nil || IsNil(o.Certifications) {
		var ret []CertificationDetails
		return ret
	}
	return o.Certifications
}

// GetCertificationsOk returns a tuple with the Certifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedCertifications) GetCertificationsOk() ([]CertificationDetails, bool) {
	if o == nil || IsNil(o.Certifications) {
		return nil, false
	}
	return o.Certifications, true
}

// HasCertifications returns a boolean if a field has been set.
func (o *PagedCertifications) HasCertifications() bool {
	if o != nil && !IsNil(o.Certifications) {
		return true
	}

	return false
}

// SetCertifications gets a reference to the given []CertificationDetails and assigns it to the Certifications field.
func (o *PagedCertifications) SetCertifications(v []CertificationDetails) {
	o.Certifications = v
}

func (o PagedCertifications) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagedCertifications) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPage, errPage := json.Marshal(o.Page)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	errPage = json.Unmarshal([]byte(serializedPage), &toSerialize)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	if !IsNil(o.Certifications) {
		toSerialize["certifications"] = o.Certifications
	}
	return toSerialize, nil
}

type NullablePagedCertifications struct {
	value *PagedCertifications
	isSet bool
}

func (v NullablePagedCertifications) Get() *PagedCertifications {
	return v.value
}

func (v *NullablePagedCertifications) Set(val *PagedCertifications) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedCertifications) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedCertifications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedCertifications(val *PagedCertifications) *NullablePagedCertifications {
	return &NullablePagedCertifications{value: val, isSet: true}
}

func (v NullablePagedCertifications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedCertifications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


