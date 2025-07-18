/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.74.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// Seniority the model 'Seniority'
type Seniority string

// List of Seniority
const (
	UNKNOWN Seniority = "UNKNOWN"
	ENTRY_LEVEL Seniority = "ENTRY_LEVEL"
	PROFESSIONAL Seniority = "PROFESSIONAL"
	SENIOR Seniority = "SENIOR"
	MANAGING Seniority = "MANAGING"
	PRINCIPAL Seniority = "PRINCIPAL"
)

// All allowed values of Seniority enum
var AllowedSeniorityEnumValues = []Seniority{
	"UNKNOWN",
	"ENTRY_LEVEL",
	"PROFESSIONAL",
	"SENIOR",
	"MANAGING",
	"PRINCIPAL",
}

func (v *Seniority) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Seniority(value)
	for _, existing := range AllowedSeniorityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Seniority", value)
}

// NewSeniorityFromValue returns a pointer to a valid Seniority
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSeniorityFromValue(v string) (*Seniority, error) {
	ev := Seniority(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Seniority: valid values are %v", v, AllowedSeniorityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Seniority) IsValid() bool {
	for _, existing := range AllowedSeniorityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Seniority value
func (v Seniority) Ptr() *Seniority {
	return &v
}

type NullableSeniority struct {
	value *Seniority
	isSet bool
}

func (v NullableSeniority) Get() *Seniority {
	return v.value
}

func (v *NullableSeniority) Set(val *Seniority) {
	v.value = val
	v.isSet = true
}

func (v NullableSeniority) IsSet() bool {
	return v.isSet
}

func (v *NullableSeniority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeniority(val *Seniority) *NullableSeniority {
	return &NullableSeniority{value: val, isSet: true}
}

func (v NullableSeniority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeniority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

