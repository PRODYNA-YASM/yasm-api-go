/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ProjectScoreDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectScoreDetail{}

// ProjectScoreDetail struct for ProjectScoreDetail
type ProjectScoreDetail struct {
	ProjectDetails
	DirectHit *bool `json:"directHit,omitempty"`
	Scores []Score `json:"scores,omitempty"`
}

// NewProjectScoreDetail instantiates a new ProjectScoreDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectScoreDetail() *ProjectScoreDetail {
	this := ProjectScoreDetail{}
	return &this
}

// NewProjectScoreDetailWithDefaults instantiates a new ProjectScoreDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectScoreDetailWithDefaults() *ProjectScoreDetail {
	this := ProjectScoreDetail{}
	return &this
}

// GetDirectHit returns the DirectHit field value if set, zero value otherwise.
func (o *ProjectScoreDetail) GetDirectHit() bool {
	if o == nil || IsNil(o.DirectHit) {
		var ret bool
		return ret
	}
	return *o.DirectHit
}

// GetDirectHitOk returns a tuple with the DirectHit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectScoreDetail) GetDirectHitOk() (*bool, bool) {
	if o == nil || IsNil(o.DirectHit) {
		return nil, false
	}
	return o.DirectHit, true
}

// HasDirectHit returns a boolean if a field has been set.
func (o *ProjectScoreDetail) HasDirectHit() bool {
	if o != nil && !IsNil(o.DirectHit) {
		return true
	}

	return false
}

// SetDirectHit gets a reference to the given bool and assigns it to the DirectHit field.
func (o *ProjectScoreDetail) SetDirectHit(v bool) {
	o.DirectHit = &v
}

// GetScores returns the Scores field value if set, zero value otherwise.
func (o *ProjectScoreDetail) GetScores() []Score {
	if o == nil || IsNil(o.Scores) {
		var ret []Score
		return ret
	}
	return o.Scores
}

// GetScoresOk returns a tuple with the Scores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectScoreDetail) GetScoresOk() ([]Score, bool) {
	if o == nil || IsNil(o.Scores) {
		return nil, false
	}
	return o.Scores, true
}

// HasScores returns a boolean if a field has been set.
func (o *ProjectScoreDetail) HasScores() bool {
	if o != nil && !IsNil(o.Scores) {
		return true
	}

	return false
}

// SetScores gets a reference to the given []Score and assigns it to the Scores field.
func (o *ProjectScoreDetail) SetScores(v []Score) {
	o.Scores = v
}

func (o ProjectScoreDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectScoreDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedProjectDetails, errProjectDetails := json.Marshal(o.ProjectDetails)
	if errProjectDetails != nil {
		return map[string]interface{}{}, errProjectDetails
	}
	errProjectDetails = json.Unmarshal([]byte(serializedProjectDetails), &toSerialize)
	if errProjectDetails != nil {
		return map[string]interface{}{}, errProjectDetails
	}
	if !IsNil(o.DirectHit) {
		toSerialize["directHit"] = o.DirectHit
	}
	if !IsNil(o.Scores) {
		toSerialize["scores"] = o.Scores
	}
	return toSerialize, nil
}

type NullableProjectScoreDetail struct {
	value *ProjectScoreDetail
	isSet bool
}

func (v NullableProjectScoreDetail) Get() *ProjectScoreDetail {
	return v.value
}

func (v *NullableProjectScoreDetail) Set(val *ProjectScoreDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectScoreDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectScoreDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectScoreDetail(val *ProjectScoreDetail) *NullableProjectScoreDetail {
	return &NullableProjectScoreDetail{value: val, isSet: true}
}

func (v NullableProjectScoreDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectScoreDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


