/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 0.9.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ProjectAllOf struct for ProjectAllOf
type ProjectAllOf struct {
	Description *string `json:"description,omitempty"`
	Timeframe *Timeframed `json:"timeframe,omitempty"`
}

// NewProjectAllOf instantiates a new ProjectAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectAllOf() *ProjectAllOf {
	this := ProjectAllOf{}
	return &this
}

// NewProjectAllOfWithDefaults instantiates a new ProjectAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectAllOfWithDefaults() *ProjectAllOf {
	this := ProjectAllOf{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProjectAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProjectAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetTimeframe returns the Timeframe field value if set, zero value otherwise.
func (o *ProjectAllOf) GetTimeframe() Timeframed {
	if o == nil || o.Timeframe == nil {
		var ret Timeframed
		return ret
	}
	return *o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectAllOf) GetTimeframeOk() (*Timeframed, bool) {
	if o == nil || o.Timeframe == nil {
		return nil, false
	}
	return o.Timeframe, true
}

// HasTimeframe returns a boolean if a field has been set.
func (o *ProjectAllOf) HasTimeframe() bool {
	if o != nil && o.Timeframe != nil {
		return true
	}

	return false
}

// SetTimeframe gets a reference to the given Timeframed and assigns it to the Timeframe field.
func (o *ProjectAllOf) SetTimeframe(v Timeframed) {
	o.Timeframe = &v
}

func (o ProjectAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Timeframe != nil {
		toSerialize["timeframe"] = o.Timeframe
	}
	return json.Marshal(toSerialize)
}

type NullableProjectAllOf struct {
	value *ProjectAllOf
	isSet bool
}

func (v NullableProjectAllOf) Get() *ProjectAllOf {
	return v.value
}

func (v *NullableProjectAllOf) Set(val *ProjectAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectAllOf(val *ProjectAllOf) *NullableProjectAllOf {
	return &NullableProjectAllOf{value: val, isSet: true}
}

func (v NullableProjectAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


