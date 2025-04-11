/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.73.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the Award type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Award{}

// Award struct for Award
type Award struct {
	NamedDomainModel
	Synonyms []string `json:"synonyms,omitempty"`
	Description *string `json:"description,omitempty"`
	// The year in which the award was received
	YearOfAward NullableInt32 `json:"yearOfAward,omitempty"`
}

// NewAward instantiates a new Award object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAward(id string, name string) *Award {
	this := Award{}
	this.Id = id
	this.Name = name
	return &this
}

// NewAwardWithDefaults instantiates a new Award object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwardWithDefaults() *Award {
	this := Award{}
	return &this
}

// GetSynonyms returns the Synonyms field value if set, zero value otherwise.
func (o *Award) GetSynonyms() []string {
	if o == nil || IsNil(o.Synonyms) {
		var ret []string
		return ret
	}
	return o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Award) GetSynonymsOk() ([]string, bool) {
	if o == nil || IsNil(o.Synonyms) {
		return nil, false
	}
	return o.Synonyms, true
}

// HasSynonyms returns a boolean if a field has been set.
func (o *Award) HasSynonyms() bool {
	if o != nil && !IsNil(o.Synonyms) {
		return true
	}

	return false
}

// SetSynonyms gets a reference to the given []string and assigns it to the Synonyms field.
func (o *Award) SetSynonyms(v []string) {
	o.Synonyms = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Award) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Award) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Award) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Award) SetDescription(v string) {
	o.Description = &v
}

// GetYearOfAward returns the YearOfAward field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Award) GetYearOfAward() int32 {
	if o == nil || IsNil(o.YearOfAward.Get()) {
		var ret int32
		return ret
	}
	return *o.YearOfAward.Get()
}

// GetYearOfAwardOk returns a tuple with the YearOfAward field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Award) GetYearOfAwardOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.YearOfAward.Get(), o.YearOfAward.IsSet()
}

// HasYearOfAward returns a boolean if a field has been set.
func (o *Award) HasYearOfAward() bool {
	if o != nil && o.YearOfAward.IsSet() {
		return true
	}

	return false
}

// SetYearOfAward gets a reference to the given NullableInt32 and assigns it to the YearOfAward field.
func (o *Award) SetYearOfAward(v int32) {
	o.YearOfAward.Set(&v)
}
// SetYearOfAwardNil sets the value for YearOfAward to be an explicit nil
func (o *Award) SetYearOfAwardNil() {
	o.YearOfAward.Set(nil)
}

// UnsetYearOfAward ensures that no value is present for YearOfAward, not even an explicit nil
func (o *Award) UnsetYearOfAward() {
	o.YearOfAward.Unset()
}

func (o Award) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Award) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedNamedDomainModel, errNamedDomainModel := json.Marshal(o.NamedDomainModel)
	if errNamedDomainModel != nil {
		return map[string]interface{}{}, errNamedDomainModel
	}
	errNamedDomainModel = json.Unmarshal([]byte(serializedNamedDomainModel), &toSerialize)
	if errNamedDomainModel != nil {
		return map[string]interface{}{}, errNamedDomainModel
	}
	if !IsNil(o.Synonyms) {
		toSerialize["synonyms"] = o.Synonyms
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.YearOfAward.IsSet() {
		toSerialize["yearOfAward"] = o.YearOfAward.Get()
	}
	return toSerialize, nil
}

type NullableAward struct {
	value *Award
	isSet bool
}

func (v NullableAward) Get() *Award {
	return v.value
}

func (v *NullableAward) Set(val *Award) {
	v.value = val
	v.isSet = true
}

func (v NullableAward) IsSet() bool {
	return v.isSet
}

func (v *NullableAward) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAward(val *Award) *NullableAward {
	return &NullableAward{value: val, isSet: true}
}

func (v NullableAward) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAward) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


