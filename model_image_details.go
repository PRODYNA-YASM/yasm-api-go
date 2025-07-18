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

// checks if the ImageDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageDetails{}

// ImageDetails struct for ImageDetails
type ImageDetails struct {
	BasicDomainModel
	Audit *Audit `json:"audit,omitempty"`
}

// NewImageDetails instantiates a new ImageDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageDetails(id string) *ImageDetails {
	this := ImageDetails{}
	this.Id = id
	return &this
}

// NewImageDetailsWithDefaults instantiates a new ImageDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageDetailsWithDefaults() *ImageDetails {
	this := ImageDetails{}
	return &this
}

// GetAudit returns the Audit field value if set, zero value otherwise.
func (o *ImageDetails) GetAudit() Audit {
	if o == nil || IsNil(o.Audit) {
		var ret Audit
		return ret
	}
	return *o.Audit
}

// GetAuditOk returns a tuple with the Audit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageDetails) GetAuditOk() (*Audit, bool) {
	if o == nil || IsNil(o.Audit) {
		return nil, false
	}
	return o.Audit, true
}

// HasAudit returns a boolean if a field has been set.
func (o *ImageDetails) HasAudit() bool {
	if o != nil && !IsNil(o.Audit) {
		return true
	}

	return false
}

// SetAudit gets a reference to the given Audit and assigns it to the Audit field.
func (o *ImageDetails) SetAudit(v Audit) {
	o.Audit = &v
}

func (o ImageDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedBasicDomainModel, errBasicDomainModel := json.Marshal(o.BasicDomainModel)
	if errBasicDomainModel != nil {
		return map[string]interface{}{}, errBasicDomainModel
	}
	errBasicDomainModel = json.Unmarshal([]byte(serializedBasicDomainModel), &toSerialize)
	if errBasicDomainModel != nil {
		return map[string]interface{}{}, errBasicDomainModel
	}
	if !IsNil(o.Audit) {
		toSerialize["audit"] = o.Audit
	}
	return toSerialize, nil
}

type NullableImageDetails struct {
	value *ImageDetails
	isSet bool
}

func (v NullableImageDetails) Get() *ImageDetails {
	return v.value
}

func (v *NullableImageDetails) Set(val *ImageDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableImageDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableImageDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageDetails(val *ImageDetails) *NullableImageDetails {
	return &NullableImageDetails{value: val, isSet: true}
}

func (v NullableImageDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


