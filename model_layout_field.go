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

// checks if the LayoutField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LayoutField{}

// LayoutField struct for LayoutField
type LayoutField struct {
	NamedDomainModel
	// with inside the svg tag
	Width *string `json:"width,omitempty"`
	// height of the svg tag
	Height *string `json:"height,omitempty"`
}

// NewLayoutField instantiates a new LayoutField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLayoutField(id string, name string) *LayoutField {
	this := LayoutField{}
	this.Id = id
	this.Name = name
	return &this
}

// NewLayoutFieldWithDefaults instantiates a new LayoutField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLayoutFieldWithDefaults() *LayoutField {
	this := LayoutField{}
	return &this
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *LayoutField) GetWidth() string {
	if o == nil || IsNil(o.Width) {
		var ret string
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutField) GetWidthOk() (*string, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *LayoutField) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given string and assigns it to the Width field.
func (o *LayoutField) SetWidth(v string) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *LayoutField) GetHeight() string {
	if o == nil || IsNil(o.Height) {
		var ret string
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutField) GetHeightOk() (*string, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *LayoutField) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given string and assigns it to the Height field.
func (o *LayoutField) SetHeight(v string) {
	o.Height = &v
}

func (o LayoutField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LayoutField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedNamedDomainModel, errNamedDomainModel := json.Marshal(o.NamedDomainModel)
	if errNamedDomainModel != nil {
		return map[string]interface{}{}, errNamedDomainModel
	}
	errNamedDomainModel = json.Unmarshal([]byte(serializedNamedDomainModel), &toSerialize)
	if errNamedDomainModel != nil {
		return map[string]interface{}{}, errNamedDomainModel
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	return toSerialize, nil
}

type NullableLayoutField struct {
	value *LayoutField
	isSet bool
}

func (v NullableLayoutField) Get() *LayoutField {
	return v.value
}

func (v *NullableLayoutField) Set(val *LayoutField) {
	v.value = val
	v.isSet = true
}

func (v NullableLayoutField) IsSet() bool {
	return v.isSet
}

func (v *NullableLayoutField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLayoutField(val *LayoutField) *NullableLayoutField {
	return &NullableLayoutField{value: val, isSet: true}
}

func (v NullableLayoutField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLayoutField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


