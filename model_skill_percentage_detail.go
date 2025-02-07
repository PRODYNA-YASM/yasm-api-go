/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.62.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the SkillPercentageDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkillPercentageDetail{}

// SkillPercentageDetail struct for SkillPercentageDetail
type SkillPercentageDetail struct {
	Skill *Skill `json:"skill,omitempty"`
	Percentage *float64 `json:"percentage,omitempty"`
	DaysOfUsage *float32 `json:"daysOfUsage,omitempty"`
	ParticipationCount *float32 `json:"participationCount,omitempty"`
}

// NewSkillPercentageDetail instantiates a new SkillPercentageDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkillPercentageDetail() *SkillPercentageDetail {
	this := SkillPercentageDetail{}
	return &this
}

// NewSkillPercentageDetailWithDefaults instantiates a new SkillPercentageDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkillPercentageDetailWithDefaults() *SkillPercentageDetail {
	this := SkillPercentageDetail{}
	return &this
}

// GetSkill returns the Skill field value if set, zero value otherwise.
func (o *SkillPercentageDetail) GetSkill() Skill {
	if o == nil || IsNil(o.Skill) {
		var ret Skill
		return ret
	}
	return *o.Skill
}

// GetSkillOk returns a tuple with the Skill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillPercentageDetail) GetSkillOk() (*Skill, bool) {
	if o == nil || IsNil(o.Skill) {
		return nil, false
	}
	return o.Skill, true
}

// HasSkill returns a boolean if a field has been set.
func (o *SkillPercentageDetail) HasSkill() bool {
	if o != nil && !IsNil(o.Skill) {
		return true
	}

	return false
}

// SetSkill gets a reference to the given Skill and assigns it to the Skill field.
func (o *SkillPercentageDetail) SetSkill(v Skill) {
	o.Skill = &v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *SkillPercentageDetail) GetPercentage() float64 {
	if o == nil || IsNil(o.Percentage) {
		var ret float64
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillPercentageDetail) GetPercentageOk() (*float64, bool) {
	if o == nil || IsNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *SkillPercentageDetail) HasPercentage() bool {
	if o != nil && !IsNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given float64 and assigns it to the Percentage field.
func (o *SkillPercentageDetail) SetPercentage(v float64) {
	o.Percentage = &v
}

// GetDaysOfUsage returns the DaysOfUsage field value if set, zero value otherwise.
func (o *SkillPercentageDetail) GetDaysOfUsage() float32 {
	if o == nil || IsNil(o.DaysOfUsage) {
		var ret float32
		return ret
	}
	return *o.DaysOfUsage
}

// GetDaysOfUsageOk returns a tuple with the DaysOfUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillPercentageDetail) GetDaysOfUsageOk() (*float32, bool) {
	if o == nil || IsNil(o.DaysOfUsage) {
		return nil, false
	}
	return o.DaysOfUsage, true
}

// HasDaysOfUsage returns a boolean if a field has been set.
func (o *SkillPercentageDetail) HasDaysOfUsage() bool {
	if o != nil && !IsNil(o.DaysOfUsage) {
		return true
	}

	return false
}

// SetDaysOfUsage gets a reference to the given float32 and assigns it to the DaysOfUsage field.
func (o *SkillPercentageDetail) SetDaysOfUsage(v float32) {
	o.DaysOfUsage = &v
}

// GetParticipationCount returns the ParticipationCount field value if set, zero value otherwise.
func (o *SkillPercentageDetail) GetParticipationCount() float32 {
	if o == nil || IsNil(o.ParticipationCount) {
		var ret float32
		return ret
	}
	return *o.ParticipationCount
}

// GetParticipationCountOk returns a tuple with the ParticipationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillPercentageDetail) GetParticipationCountOk() (*float32, bool) {
	if o == nil || IsNil(o.ParticipationCount) {
		return nil, false
	}
	return o.ParticipationCount, true
}

// HasParticipationCount returns a boolean if a field has been set.
func (o *SkillPercentageDetail) HasParticipationCount() bool {
	if o != nil && !IsNil(o.ParticipationCount) {
		return true
	}

	return false
}

// SetParticipationCount gets a reference to the given float32 and assigns it to the ParticipationCount field.
func (o *SkillPercentageDetail) SetParticipationCount(v float32) {
	o.ParticipationCount = &v
}

func (o SkillPercentageDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkillPercentageDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Skill) {
		toSerialize["skill"] = o.Skill
	}
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	if !IsNil(o.DaysOfUsage) {
		toSerialize["daysOfUsage"] = o.DaysOfUsage
	}
	if !IsNil(o.ParticipationCount) {
		toSerialize["participationCount"] = o.ParticipationCount
	}
	return toSerialize, nil
}

type NullableSkillPercentageDetail struct {
	value *SkillPercentageDetail
	isSet bool
}

func (v NullableSkillPercentageDetail) Get() *SkillPercentageDetail {
	return v.value
}

func (v *NullableSkillPercentageDetail) Set(val *SkillPercentageDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableSkillPercentageDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableSkillPercentageDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkillPercentageDetail(val *SkillPercentageDetail) *NullableSkillPercentageDetail {
	return &NullableSkillPercentageDetail{value: val, isSet: true}
}

func (v NullableSkillPercentageDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkillPercentageDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


