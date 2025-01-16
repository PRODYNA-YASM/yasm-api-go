/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.51.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the FieldImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FieldImage{}

// FieldImage struct for FieldImage
type FieldImage struct {
	Image *ImageDetails `json:"image,omitempty"`
	LayoutFieldId *string `json:"layoutFieldId,omitempty"`
}

// NewFieldImage instantiates a new FieldImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldImage() *FieldImage {
	this := FieldImage{}
	return &this
}

// NewFieldImageWithDefaults instantiates a new FieldImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldImageWithDefaults() *FieldImage {
	this := FieldImage{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *FieldImage) GetImage() ImageDetails {
	if o == nil || IsNil(o.Image) {
		var ret ImageDetails
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldImage) GetImageOk() (*ImageDetails, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *FieldImage) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given ImageDetails and assigns it to the Image field.
func (o *FieldImage) SetImage(v ImageDetails) {
	o.Image = &v
}

// GetLayoutFieldId returns the LayoutFieldId field value if set, zero value otherwise.
func (o *FieldImage) GetLayoutFieldId() string {
	if o == nil || IsNil(o.LayoutFieldId) {
		var ret string
		return ret
	}
	return *o.LayoutFieldId
}

// GetLayoutFieldIdOk returns a tuple with the LayoutFieldId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldImage) GetLayoutFieldIdOk() (*string, bool) {
	if o == nil || IsNil(o.LayoutFieldId) {
		return nil, false
	}
	return o.LayoutFieldId, true
}

// HasLayoutFieldId returns a boolean if a field has been set.
func (o *FieldImage) HasLayoutFieldId() bool {
	if o != nil && !IsNil(o.LayoutFieldId) {
		return true
	}

	return false
}

// SetLayoutFieldId gets a reference to the given string and assigns it to the LayoutFieldId field.
func (o *FieldImage) SetLayoutFieldId(v string) {
	o.LayoutFieldId = &v
}

func (o FieldImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FieldImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.LayoutFieldId) {
		toSerialize["layoutFieldId"] = o.LayoutFieldId
	}
	return toSerialize, nil
}

type NullableFieldImage struct {
	value *FieldImage
	isSet bool
}

func (v NullableFieldImage) Get() *FieldImage {
	return v.value
}

func (v *NullableFieldImage) Set(val *FieldImage) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldImage) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldImage(val *FieldImage) *NullableFieldImage {
	return &NullableFieldImage{value: val, isSet: true}
}

func (v NullableFieldImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


