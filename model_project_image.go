/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.70.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ProjectImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectImage{}

// ProjectImage struct for ProjectImage
type ProjectImage struct {
	BasicDomainModel
	ProjectId *string `json:"projectId,omitempty"`
	FieldImages []FieldImage `json:"fieldImages,omitempty"`
}

// NewProjectImage instantiates a new ProjectImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectImage(id string) *ProjectImage {
	this := ProjectImage{}
	this.Id = id
	return &this
}

// NewProjectImageWithDefaults instantiates a new ProjectImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectImageWithDefaults() *ProjectImage {
	this := ProjectImage{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ProjectImage) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectImage) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ProjectImage) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *ProjectImage) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetFieldImages returns the FieldImages field value if set, zero value otherwise.
func (o *ProjectImage) GetFieldImages() []FieldImage {
	if o == nil || IsNil(o.FieldImages) {
		var ret []FieldImage
		return ret
	}
	return o.FieldImages
}

// GetFieldImagesOk returns a tuple with the FieldImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectImage) GetFieldImagesOk() ([]FieldImage, bool) {
	if o == nil || IsNil(o.FieldImages) {
		return nil, false
	}
	return o.FieldImages, true
}

// HasFieldImages returns a boolean if a field has been set.
func (o *ProjectImage) HasFieldImages() bool {
	if o != nil && !IsNil(o.FieldImages) {
		return true
	}

	return false
}

// SetFieldImages gets a reference to the given []FieldImage and assigns it to the FieldImages field.
func (o *ProjectImage) SetFieldImages(v []FieldImage) {
	o.FieldImages = v
}

func (o ProjectImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedBasicDomainModel, errBasicDomainModel := json.Marshal(o.BasicDomainModel)
	if errBasicDomainModel != nil {
		return map[string]interface{}{}, errBasicDomainModel
	}
	errBasicDomainModel = json.Unmarshal([]byte(serializedBasicDomainModel), &toSerialize)
	if errBasicDomainModel != nil {
		return map[string]interface{}{}, errBasicDomainModel
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.FieldImages) {
		toSerialize["fieldImages"] = o.FieldImages
	}
	return toSerialize, nil
}

type NullableProjectImage struct {
	value *ProjectImage
	isSet bool
}

func (v NullableProjectImage) Get() *ProjectImage {
	return v.value
}

func (v *NullableProjectImage) Set(val *ProjectImage) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectImage) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectImage(val *ProjectImage) *NullableProjectImage {
	return &NullableProjectImage{value: val, isSet: true}
}

func (v NullableProjectImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


