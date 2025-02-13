/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.66.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ScoreResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScoreResult{}

// ScoreResult struct for ScoreResult
type ScoreResult struct {
	Score *float32 `json:"score,omitempty"`
	// Indicates if the score is a direct hit
	DirectHit *bool `json:"directHit,omitempty"`
	Scores []Score `json:"scores,omitempty"`
}

// NewScoreResult instantiates a new ScoreResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScoreResult() *ScoreResult {
	this := ScoreResult{}
	return &this
}

// NewScoreResultWithDefaults instantiates a new ScoreResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScoreResultWithDefaults() *ScoreResult {
	this := ScoreResult{}
	return &this
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *ScoreResult) GetScore() float32 {
	if o == nil || IsNil(o.Score) {
		var ret float32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScoreResult) GetScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *ScoreResult) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given float32 and assigns it to the Score field.
func (o *ScoreResult) SetScore(v float32) {
	o.Score = &v
}

// GetDirectHit returns the DirectHit field value if set, zero value otherwise.
func (o *ScoreResult) GetDirectHit() bool {
	if o == nil || IsNil(o.DirectHit) {
		var ret bool
		return ret
	}
	return *o.DirectHit
}

// GetDirectHitOk returns a tuple with the DirectHit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScoreResult) GetDirectHitOk() (*bool, bool) {
	if o == nil || IsNil(o.DirectHit) {
		return nil, false
	}
	return o.DirectHit, true
}

// HasDirectHit returns a boolean if a field has been set.
func (o *ScoreResult) HasDirectHit() bool {
	if o != nil && !IsNil(o.DirectHit) {
		return true
	}

	return false
}

// SetDirectHit gets a reference to the given bool and assigns it to the DirectHit field.
func (o *ScoreResult) SetDirectHit(v bool) {
	o.DirectHit = &v
}

// GetScores returns the Scores field value if set, zero value otherwise.
func (o *ScoreResult) GetScores() []Score {
	if o == nil || IsNil(o.Scores) {
		var ret []Score
		return ret
	}
	return o.Scores
}

// GetScoresOk returns a tuple with the Scores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScoreResult) GetScoresOk() ([]Score, bool) {
	if o == nil || IsNil(o.Scores) {
		return nil, false
	}
	return o.Scores, true
}

// HasScores returns a boolean if a field has been set.
func (o *ScoreResult) HasScores() bool {
	if o != nil && !IsNil(o.Scores) {
		return true
	}

	return false
}

// SetScores gets a reference to the given []Score and assigns it to the Scores field.
func (o *ScoreResult) SetScores(v []Score) {
	o.Scores = v
}

func (o ScoreResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScoreResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !IsNil(o.DirectHit) {
		toSerialize["directHit"] = o.DirectHit
	}
	if !IsNil(o.Scores) {
		toSerialize["scores"] = o.Scores
	}
	return toSerialize, nil
}

type NullableScoreResult struct {
	value *ScoreResult
	isSet bool
}

func (v NullableScoreResult) Get() *ScoreResult {
	return v.value
}

func (v *NullableScoreResult) Set(val *ScoreResult) {
	v.value = val
	v.isSet = true
}

func (v NullableScoreResult) IsSet() bool {
	return v.isSet
}

func (v *NullableScoreResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScoreResult(val *ScoreResult) *NullableScoreResult {
	return &NullableScoreResult{value: val, isSet: true}
}

func (v NullableScoreResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScoreResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


