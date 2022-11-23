/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.3.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PersonScoreDetail struct for PersonScoreDetail
type PersonScoreDetail struct {
	PersonDetails
	Score *float32 `json:"score,omitempty"`
}

// NewPersonScoreDetail instantiates a new PersonScoreDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonScoreDetail() *PersonScoreDetail {
	this := PersonScoreDetail{}
	return &this
}

// NewPersonScoreDetailWithDefaults instantiates a new PersonScoreDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonScoreDetailWithDefaults() *PersonScoreDetail {
	this := PersonScoreDetail{}
	return &this
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *PersonScoreDetail) GetScore() float32 {
	if o == nil || o.Score == nil {
		var ret float32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonScoreDetail) GetScoreOk() (*float32, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *PersonScoreDetail) HasScore() bool {
	if o != nil && o.Score != nil {
		return true
	}

	return false
}

// SetScore gets a reference to the given float32 and assigns it to the Score field.
func (o *PersonScoreDetail) SetScore(v float32) {
	o.Score = &v
}

func (o PersonScoreDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPersonDetails, errPersonDetails := json.Marshal(o.PersonDetails)
	if errPersonDetails != nil {
		return []byte{}, errPersonDetails
	}
	errPersonDetails = json.Unmarshal([]byte(serializedPersonDetails), &toSerialize)
	if errPersonDetails != nil {
		return []byte{}, errPersonDetails
	}
	if o.Score != nil {
		toSerialize["score"] = o.Score
	}
	return json.Marshal(toSerialize)
}

type NullablePersonScoreDetail struct {
	value *PersonScoreDetail
	isSet bool
}

func (v NullablePersonScoreDetail) Get() *PersonScoreDetail {
	return v.value
}

func (v *NullablePersonScoreDetail) Set(val *PersonScoreDetail) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonScoreDetail) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonScoreDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonScoreDetail(val *PersonScoreDetail) *NullablePersonScoreDetail {
	return &NullablePersonScoreDetail{value: val, isSet: true}
}

func (v NullablePersonScoreDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonScoreDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


