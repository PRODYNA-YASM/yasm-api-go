/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// Confidentiality the model 'Confidentiality'
type Confidentiality string

// List of Confidentiality
const (
	UNKNOWN_CONFIDENTIALITY Confidentiality = "UNKNOWN_CONFIDENTIALITY"
	PUBLIC Confidentiality = "PUBLIC"
	CONFIDENTIAL Confidentiality = "CONFIDENTIAL"
)

// All allowed values of Confidentiality enum
var AllowedConfidentialityEnumValues = []Confidentiality{
	"UNKNOWN_CONFIDENTIALITY",
	"PUBLIC",
	"CONFIDENTIAL",
}

func (v *Confidentiality) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Confidentiality(value)
	for _, existing := range AllowedConfidentialityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Confidentiality", value)
}

// NewConfidentialityFromValue returns a pointer to a valid Confidentiality
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConfidentialityFromValue(v string) (*Confidentiality, error) {
	ev := Confidentiality(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Confidentiality: valid values are %v", v, AllowedConfidentialityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Confidentiality) IsValid() bool {
	for _, existing := range AllowedConfidentialityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Confidentiality value
func (v Confidentiality) Ptr() *Confidentiality {
	return &v
}

type NullableConfidentiality struct {
	value *Confidentiality
	isSet bool
}

func (v NullableConfidentiality) Get() *Confidentiality {
	return v.value
}

func (v *NullableConfidentiality) Set(val *Confidentiality) {
	v.value = val
	v.isSet = true
}

func (v NullableConfidentiality) IsSet() bool {
	return v.isSet
}

func (v *NullableConfidentiality) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfidentiality(val *Confidentiality) *NullableConfidentiality {
	return &NullableConfidentiality{value: val, isSet: true}
}

func (v NullableConfidentiality) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfidentiality) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

