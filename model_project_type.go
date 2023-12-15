/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// ProjectType the model 'ProjectType'
type ProjectType string

// List of ProjectType
const (
	UNKNOWN_TYPE ProjectType = "UNKNOWN_TYPE"
	INTERNAL ProjectType = "INTERNAL"
	EXTERNAL ProjectType = "EXTERNAL"
)

// All allowed values of ProjectType enum
var AllowedProjectTypeEnumValues = []ProjectType{
	"UNKNOWN_TYPE",
	"INTERNAL",
	"EXTERNAL",
}

func (v *ProjectType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProjectType(value)
	for _, existing := range AllowedProjectTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProjectType", value)
}

// NewProjectTypeFromValue returns a pointer to a valid ProjectType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectTypeFromValue(v string) (*ProjectType, error) {
	ev := ProjectType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectType: valid values are %v", v, AllowedProjectTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectType) IsValid() bool {
	for _, existing := range AllowedProjectTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProjectType value
func (v ProjectType) Ptr() *ProjectType {
	return &v
}

type NullableProjectType struct {
	value *ProjectType
	isSet bool
}

func (v NullableProjectType) Get() *ProjectType {
	return v.value
}

func (v *NullableProjectType) Set(val *ProjectType) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectType) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectType(val *ProjectType) *NullableProjectType {
	return &NullableProjectType{value: val, isSet: true}
}

func (v NullableProjectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

