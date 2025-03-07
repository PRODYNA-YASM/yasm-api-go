/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.70.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the Score type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Score{}

// Score struct for Score
type Score struct {
	Score *float32 `json:"score,omitempty"`
	// Indicates if the score is a direct hit
	DirectHit *bool `json:"directHit,omitempty"`
	// Indicates if the score is the best match
	BestMatch *bool `json:"bestMatch,omitempty"`
	InputEntity *NamedDomainModel `json:"inputEntity,omitempty"`
	// The matching entities
	MatchingEntity []NamedDomainModel `json:"matchingEntity,omitempty"`
}

// NewScore instantiates a new Score object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScore() *Score {
	this := Score{}
	return &this
}

// NewScoreWithDefaults instantiates a new Score object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScoreWithDefaults() *Score {
	this := Score{}
	return &this
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *Score) GetScore() float32 {
	if o == nil || IsNil(o.Score) {
		var ret float32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Score) GetScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *Score) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given float32 and assigns it to the Score field.
func (o *Score) SetScore(v float32) {
	o.Score = &v
}

// GetDirectHit returns the DirectHit field value if set, zero value otherwise.
func (o *Score) GetDirectHit() bool {
	if o == nil || IsNil(o.DirectHit) {
		var ret bool
		return ret
	}
	return *o.DirectHit
}

// GetDirectHitOk returns a tuple with the DirectHit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Score) GetDirectHitOk() (*bool, bool) {
	if o == nil || IsNil(o.DirectHit) {
		return nil, false
	}
	return o.DirectHit, true
}

// HasDirectHit returns a boolean if a field has been set.
func (o *Score) HasDirectHit() bool {
	if o != nil && !IsNil(o.DirectHit) {
		return true
	}

	return false
}

// SetDirectHit gets a reference to the given bool and assigns it to the DirectHit field.
func (o *Score) SetDirectHit(v bool) {
	o.DirectHit = &v
}

// GetBestMatch returns the BestMatch field value if set, zero value otherwise.
func (o *Score) GetBestMatch() bool {
	if o == nil || IsNil(o.BestMatch) {
		var ret bool
		return ret
	}
	return *o.BestMatch
}

// GetBestMatchOk returns a tuple with the BestMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Score) GetBestMatchOk() (*bool, bool) {
	if o == nil || IsNil(o.BestMatch) {
		return nil, false
	}
	return o.BestMatch, true
}

// HasBestMatch returns a boolean if a field has been set.
func (o *Score) HasBestMatch() bool {
	if o != nil && !IsNil(o.BestMatch) {
		return true
	}

	return false
}

// SetBestMatch gets a reference to the given bool and assigns it to the BestMatch field.
func (o *Score) SetBestMatch(v bool) {
	o.BestMatch = &v
}

// GetInputEntity returns the InputEntity field value if set, zero value otherwise.
func (o *Score) GetInputEntity() NamedDomainModel {
	if o == nil || IsNil(o.InputEntity) {
		var ret NamedDomainModel
		return ret
	}
	return *o.InputEntity
}

// GetInputEntityOk returns a tuple with the InputEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Score) GetInputEntityOk() (*NamedDomainModel, bool) {
	if o == nil || IsNil(o.InputEntity) {
		return nil, false
	}
	return o.InputEntity, true
}

// HasInputEntity returns a boolean if a field has been set.
func (o *Score) HasInputEntity() bool {
	if o != nil && !IsNil(o.InputEntity) {
		return true
	}

	return false
}

// SetInputEntity gets a reference to the given NamedDomainModel and assigns it to the InputEntity field.
func (o *Score) SetInputEntity(v NamedDomainModel) {
	o.InputEntity = &v
}

// GetMatchingEntity returns the MatchingEntity field value if set, zero value otherwise.
func (o *Score) GetMatchingEntity() []NamedDomainModel {
	if o == nil || IsNil(o.MatchingEntity) {
		var ret []NamedDomainModel
		return ret
	}
	return o.MatchingEntity
}

// GetMatchingEntityOk returns a tuple with the MatchingEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Score) GetMatchingEntityOk() ([]NamedDomainModel, bool) {
	if o == nil || IsNil(o.MatchingEntity) {
		return nil, false
	}
	return o.MatchingEntity, true
}

// HasMatchingEntity returns a boolean if a field has been set.
func (o *Score) HasMatchingEntity() bool {
	if o != nil && !IsNil(o.MatchingEntity) {
		return true
	}

	return false
}

// SetMatchingEntity gets a reference to the given []NamedDomainModel and assigns it to the MatchingEntity field.
func (o *Score) SetMatchingEntity(v []NamedDomainModel) {
	o.MatchingEntity = v
}

func (o Score) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Score) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !IsNil(o.DirectHit) {
		toSerialize["directHit"] = o.DirectHit
	}
	if !IsNil(o.BestMatch) {
		toSerialize["bestMatch"] = o.BestMatch
	}
	if !IsNil(o.InputEntity) {
		toSerialize["inputEntity"] = o.InputEntity
	}
	if !IsNil(o.MatchingEntity) {
		toSerialize["matchingEntity"] = o.MatchingEntity
	}
	return toSerialize, nil
}

type NullableScore struct {
	value *Score
	isSet bool
}

func (v NullableScore) Get() *Score {
	return v.value
}

func (v *NullableScore) Set(val *Score) {
	v.value = val
	v.isSet = true
}

func (v NullableScore) IsSet() bool {
	return v.isSet
}

func (v *NullableScore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScore(val *Score) *NullableScore {
	return &NullableScore{value: val, isSet: true}
}

func (v NullableScore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


