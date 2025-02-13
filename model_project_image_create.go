/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.66.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ProjectImageCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectImageCreate{}

// ProjectImageCreate struct for ProjectImageCreate
type ProjectImageCreate struct {
	ProjectImage
	LayoutId *string `json:"layoutId,omitempty"`
}

// NewProjectImageCreate instantiates a new ProjectImageCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectImageCreate(id string) *ProjectImageCreate {
	this := ProjectImageCreate{}
	this.Id = id
	return &this
}

// NewProjectImageCreateWithDefaults instantiates a new ProjectImageCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectImageCreateWithDefaults() *ProjectImageCreate {
	this := ProjectImageCreate{}
	return &this
}

// GetLayoutId returns the LayoutId field value if set, zero value otherwise.
func (o *ProjectImageCreate) GetLayoutId() string {
	if o == nil || IsNil(o.LayoutId) {
		var ret string
		return ret
	}
	return *o.LayoutId
}

// GetLayoutIdOk returns a tuple with the LayoutId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectImageCreate) GetLayoutIdOk() (*string, bool) {
	if o == nil || IsNil(o.LayoutId) {
		return nil, false
	}
	return o.LayoutId, true
}

// HasLayoutId returns a boolean if a field has been set.
func (o *ProjectImageCreate) HasLayoutId() bool {
	if o != nil && !IsNil(o.LayoutId) {
		return true
	}

	return false
}

// SetLayoutId gets a reference to the given string and assigns it to the LayoutId field.
func (o *ProjectImageCreate) SetLayoutId(v string) {
	o.LayoutId = &v
}

func (o ProjectImageCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectImageCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedProjectImage, errProjectImage := json.Marshal(o.ProjectImage)
	if errProjectImage != nil {
		return map[string]interface{}{}, errProjectImage
	}
	errProjectImage = json.Unmarshal([]byte(serializedProjectImage), &toSerialize)
	if errProjectImage != nil {
		return map[string]interface{}{}, errProjectImage
	}
	if !IsNil(o.LayoutId) {
		toSerialize["layoutId"] = o.LayoutId
	}
	return toSerialize, nil
}

type NullableProjectImageCreate struct {
	value *ProjectImageCreate
	isSet bool
}

func (v NullableProjectImageCreate) Get() *ProjectImageCreate {
	return v.value
}

func (v *NullableProjectImageCreate) Set(val *ProjectImageCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectImageCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectImageCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectImageCreate(val *ProjectImageCreate) *NullableProjectImageCreate {
	return &NullableProjectImageCreate{value: val, isSet: true}
}

func (v NullableProjectImageCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectImageCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


