/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.73.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PagedProjectImages type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagedProjectImages{}

// PagedProjectImages struct for PagedProjectImages
type PagedProjectImages struct {
	Page
	ProjectImages []ProjectImageDetails `json:"projectImages,omitempty"`
}

// NewPagedProjectImages instantiates a new PagedProjectImages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedProjectImages() *PagedProjectImages {
	this := PagedProjectImages{}
	return &this
}

// NewPagedProjectImagesWithDefaults instantiates a new PagedProjectImages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedProjectImagesWithDefaults() *PagedProjectImages {
	this := PagedProjectImages{}
	return &this
}

// GetProjectImages returns the ProjectImages field value if set, zero value otherwise.
func (o *PagedProjectImages) GetProjectImages() []ProjectImageDetails {
	if o == nil || IsNil(o.ProjectImages) {
		var ret []ProjectImageDetails
		return ret
	}
	return o.ProjectImages
}

// GetProjectImagesOk returns a tuple with the ProjectImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedProjectImages) GetProjectImagesOk() ([]ProjectImageDetails, bool) {
	if o == nil || IsNil(o.ProjectImages) {
		return nil, false
	}
	return o.ProjectImages, true
}

// HasProjectImages returns a boolean if a field has been set.
func (o *PagedProjectImages) HasProjectImages() bool {
	if o != nil && !IsNil(o.ProjectImages) {
		return true
	}

	return false
}

// SetProjectImages gets a reference to the given []ProjectImageDetails and assigns it to the ProjectImages field.
func (o *PagedProjectImages) SetProjectImages(v []ProjectImageDetails) {
	o.ProjectImages = v
}

func (o PagedProjectImages) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagedProjectImages) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPage, errPage := json.Marshal(o.Page)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	errPage = json.Unmarshal([]byte(serializedPage), &toSerialize)
	if errPage != nil {
		return map[string]interface{}{}, errPage
	}
	if !IsNil(o.ProjectImages) {
		toSerialize["projectImages"] = o.ProjectImages
	}
	return toSerialize, nil
}

type NullablePagedProjectImages struct {
	value *PagedProjectImages
	isSet bool
}

func (v NullablePagedProjectImages) Get() *PagedProjectImages {
	return v.value
}

func (v *NullablePagedProjectImages) Set(val *PagedProjectImages) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedProjectImages) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedProjectImages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedProjectImages(val *PagedProjectImages) *NullablePagedProjectImages {
	return &NullablePagedProjectImages{value: val, isSet: true}
}

func (v NullablePagedProjectImages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedProjectImages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


