/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the Skill type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Skill{}

// Skill struct for Skill
type Skill struct {
	NamedDomainModel
	Synonyms []string `json:"synonyms,omitempty"`
	Description *string `json:"description,omitempty"`
	GroupPriority *int32 `json:"groupPriority,omitempty"`
	Important *bool `json:"important,omitempty"`
}

// NewSkill instantiates a new Skill object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkill(id string, name string) *Skill {
	this := Skill{}
	this.Id = id
	this.Name = name
	var important bool = false
	this.Important = &important
	return &this
}

// NewSkillWithDefaults instantiates a new Skill object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkillWithDefaults() *Skill {
	this := Skill{}
	var important bool = false
	this.Important = &important
	return &this
}

// GetSynonyms returns the Synonyms field value if set, zero value otherwise.
func (o *Skill) GetSynonyms() []string {
	if o == nil || IsNil(o.Synonyms) {
		var ret []string
		return ret
	}
	return o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Skill) GetSynonymsOk() ([]string, bool) {
	if o == nil || IsNil(o.Synonyms) {
		return nil, false
	}
	return o.Synonyms, true
}

// HasSynonyms returns a boolean if a field has been set.
func (o *Skill) HasSynonyms() bool {
	if o != nil && !IsNil(o.Synonyms) {
		return true
	}

	return false
}

// SetSynonyms gets a reference to the given []string and assigns it to the Synonyms field.
func (o *Skill) SetSynonyms(v []string) {
	o.Synonyms = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Skill) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Skill) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Skill) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Skill) SetDescription(v string) {
	o.Description = &v
}

// GetGroupPriority returns the GroupPriority field value if set, zero value otherwise.
func (o *Skill) GetGroupPriority() int32 {
	if o == nil || IsNil(o.GroupPriority) {
		var ret int32
		return ret
	}
	return *o.GroupPriority
}

// GetGroupPriorityOk returns a tuple with the GroupPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Skill) GetGroupPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.GroupPriority) {
		return nil, false
	}
	return o.GroupPriority, true
}

// HasGroupPriority returns a boolean if a field has been set.
func (o *Skill) HasGroupPriority() bool {
	if o != nil && !IsNil(o.GroupPriority) {
		return true
	}

	return false
}

// SetGroupPriority gets a reference to the given int32 and assigns it to the GroupPriority field.
func (o *Skill) SetGroupPriority(v int32) {
	o.GroupPriority = &v
}

// GetImportant returns the Important field value if set, zero value otherwise.
func (o *Skill) GetImportant() bool {
	if o == nil || IsNil(o.Important) {
		var ret bool
		return ret
	}
	return *o.Important
}

// GetImportantOk returns a tuple with the Important field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Skill) GetImportantOk() (*bool, bool) {
	if o == nil || IsNil(o.Important) {
		return nil, false
	}
	return o.Important, true
}

// HasImportant returns a boolean if a field has been set.
func (o *Skill) HasImportant() bool {
	if o != nil && !IsNil(o.Important) {
		return true
	}

	return false
}

// SetImportant gets a reference to the given bool and assigns it to the Important field.
func (o *Skill) SetImportant(v bool) {
	o.Important = &v
}

func (o Skill) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Skill) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.GroupPriority) {
		toSerialize["groupPriority"] = o.GroupPriority
	}
	if !IsNil(o.Important) {
		toSerialize["important"] = o.Important
	}
	return toSerialize, nil
}

type NullableSkill struct {
	value *Skill
	isSet bool
}

func (v NullableSkill) Get() *Skill {
	return v.value
}

func (v *NullableSkill) Set(val *Skill) {
	v.value = val
	v.isSet = true
}

func (v NullableSkill) IsSet() bool {
	return v.isSet
}

func (v *NullableSkill) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkill(val *Skill) *NullableSkill {
	return &NullableSkill{value: val, isSet: true}
}

func (v NullableSkill) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkill) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


