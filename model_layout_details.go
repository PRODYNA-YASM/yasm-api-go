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

// checks if the LayoutDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LayoutDetails{}

// LayoutDetails struct for LayoutDetails
type LayoutDetails struct {
	Audit *Audit `json:"audit,omitempty"`
	Layout *Layout `json:"layout,omitempty"`
}

// NewLayoutDetails instantiates a new LayoutDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLayoutDetails() *LayoutDetails {
	this := LayoutDetails{}
	return &this
}

// NewLayoutDetailsWithDefaults instantiates a new LayoutDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLayoutDetailsWithDefaults() *LayoutDetails {
	this := LayoutDetails{}
	return &this
}

// GetAudit returns the Audit field value if set, zero value otherwise.
func (o *LayoutDetails) GetAudit() Audit {
	if o == nil || IsNil(o.Audit) {
		var ret Audit
		return ret
	}
	return *o.Audit
}

// GetAuditOk returns a tuple with the Audit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutDetails) GetAuditOk() (*Audit, bool) {
	if o == nil || IsNil(o.Audit) {
		return nil, false
	}
	return o.Audit, true
}

// HasAudit returns a boolean if a field has been set.
func (o *LayoutDetails) HasAudit() bool {
	if o != nil && !IsNil(o.Audit) {
		return true
	}

	return false
}

// SetAudit gets a reference to the given Audit and assigns it to the Audit field.
func (o *LayoutDetails) SetAudit(v Audit) {
	o.Audit = &v
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *LayoutDetails) GetLayout() Layout {
	if o == nil || IsNil(o.Layout) {
		var ret Layout
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LayoutDetails) GetLayoutOk() (*Layout, bool) {
	if o == nil || IsNil(o.Layout) {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *LayoutDetails) HasLayout() bool {
	if o != nil && !IsNil(o.Layout) {
		return true
	}

	return false
}

// SetLayout gets a reference to the given Layout and assigns it to the Layout field.
func (o *LayoutDetails) SetLayout(v Layout) {
	o.Layout = &v
}

func (o LayoutDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LayoutDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Audit) {
		toSerialize["audit"] = o.Audit
	}
	if !IsNil(o.Layout) {
		toSerialize["layout"] = o.Layout
	}
	return toSerialize, nil
}

type NullableLayoutDetails struct {
	value *LayoutDetails
	isSet bool
}

func (v NullableLayoutDetails) Get() *LayoutDetails {
	return v.value
}

func (v *NullableLayoutDetails) Set(val *LayoutDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableLayoutDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableLayoutDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLayoutDetails(val *LayoutDetails) *NullableLayoutDetails {
	return &NullableLayoutDetails{value: val, isSet: true}
}

func (v NullableLayoutDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLayoutDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


