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

// checks if the LayoutField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LayoutField{}

// LayoutField struct for LayoutField
type LayoutField struct {
	BasicDomainModel
	Description *string `json:"description,omitempty"`
	Width *float32 `json:"width,omitempty"`
	Height *float32 `json:"height,omitempty"`
	DeviceType *string `json:"deviceType,omitempty"`
}

// NewLayoutField instantiates a new LayoutField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLayoutField(id string) *LayoutField {
	this := LayoutField{}
	this.Id = id
	return &this
}

// NewLayoutFieldWithDefaults instantiates a new LayoutField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLayoutFieldWithDefaults() *LayoutField {
	this := LayoutField{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LayoutField) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutField) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LayoutField) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LayoutField) SetDescription(v string) {
	o.Description = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *LayoutField) GetWidth() float32 {
	if o == nil || IsNil(o.Width) {
		var ret float32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutField) GetWidthOk() (*float32, bool) {
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

// SetWidth gets a reference to the given float32 and assigns it to the Width field.
func (o *LayoutField) SetWidth(v float32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *LayoutField) GetHeight() float32 {
	if o == nil || IsNil(o.Height) {
		var ret float32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutField) GetHeightOk() (*float32, bool) {
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

// SetHeight gets a reference to the given float32 and assigns it to the Height field.
func (o *LayoutField) SetHeight(v float32) {
	o.Height = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *LayoutField) GetDeviceType() string {
	if o == nil || IsNil(o.DeviceType) {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutField) GetDeviceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceType) {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *LayoutField) HasDeviceType() bool {
	if o != nil && !IsNil(o.DeviceType) {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *LayoutField) SetDeviceType(v string) {
	o.DeviceType = &v
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
	serializedBasicDomainModel, errBasicDomainModel := json.Marshal(o.BasicDomainModel)
	if errBasicDomainModel != nil {
		return map[string]interface{}{}, errBasicDomainModel
	}
	errBasicDomainModel = json.Unmarshal([]byte(serializedBasicDomainModel), &toSerialize)
	if errBasicDomainModel != nil {
		return map[string]interface{}{}, errBasicDomainModel
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.DeviceType) {
		toSerialize["deviceType"] = o.DeviceType
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


