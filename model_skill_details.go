/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.6.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the SkillDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkillDetails{}

// SkillDetails struct for SkillDetails
type SkillDetails struct {
	Skill *Skill `json:"skill,omitempty"`
	Children []Skill `json:"children,omitempty"`
	Parents []Skill `json:"parents,omitempty"`
	Kinds []Skill `json:"kinds,omitempty"`
	// The display name of the skill, e.g. \"Communication (Computing)\" or \"Communication (Design)\" or \"Protcool (Communication,IT)\"
	DisplayName *string `json:"displayName,omitempty"`
}

// NewSkillDetails instantiates a new SkillDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkillDetails() *SkillDetails {
	this := SkillDetails{}
	return &this
}

// NewSkillDetailsWithDefaults instantiates a new SkillDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkillDetailsWithDefaults() *SkillDetails {
	this := SkillDetails{}
	return &this
}

// GetSkill returns the Skill field value if set, zero value otherwise.
func (o *SkillDetails) GetSkill() Skill {
	if o == nil || IsNil(o.Skill) {
		var ret Skill
		return ret
	}
	return *o.Skill
}

// GetSkillOk returns a tuple with the Skill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillDetails) GetSkillOk() (*Skill, bool) {
	if o == nil || IsNil(o.Skill) {
		return nil, false
	}
	return o.Skill, true
}

// HasSkill returns a boolean if a field has been set.
func (o *SkillDetails) HasSkill() bool {
	if o != nil && !IsNil(o.Skill) {
		return true
	}

	return false
}

// SetSkill gets a reference to the given Skill and assigns it to the Skill field.
func (o *SkillDetails) SetSkill(v Skill) {
	o.Skill = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *SkillDetails) GetChildren() []Skill {
	if o == nil || IsNil(o.Children) {
		var ret []Skill
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillDetails) GetChildrenOk() ([]Skill, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *SkillDetails) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []Skill and assigns it to the Children field.
func (o *SkillDetails) SetChildren(v []Skill) {
	o.Children = v
}

// GetParents returns the Parents field value if set, zero value otherwise.
func (o *SkillDetails) GetParents() []Skill {
	if o == nil || IsNil(o.Parents) {
		var ret []Skill
		return ret
	}
	return o.Parents
}

// GetParentsOk returns a tuple with the Parents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillDetails) GetParentsOk() ([]Skill, bool) {
	if o == nil || IsNil(o.Parents) {
		return nil, false
	}
	return o.Parents, true
}

// HasParents returns a boolean if a field has been set.
func (o *SkillDetails) HasParents() bool {
	if o != nil && !IsNil(o.Parents) {
		return true
	}

	return false
}

// SetParents gets a reference to the given []Skill and assigns it to the Parents field.
func (o *SkillDetails) SetParents(v []Skill) {
	o.Parents = v
}

// GetKinds returns the Kinds field value if set, zero value otherwise.
func (o *SkillDetails) GetKinds() []Skill {
	if o == nil || IsNil(o.Kinds) {
		var ret []Skill
		return ret
	}
	return o.Kinds
}

// GetKindsOk returns a tuple with the Kinds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillDetails) GetKindsOk() ([]Skill, bool) {
	if o == nil || IsNil(o.Kinds) {
		return nil, false
	}
	return o.Kinds, true
}

// HasKinds returns a boolean if a field has been set.
func (o *SkillDetails) HasKinds() bool {
	if o != nil && !IsNil(o.Kinds) {
		return true
	}

	return false
}

// SetKinds gets a reference to the given []Skill and assigns it to the Kinds field.
func (o *SkillDetails) SetKinds(v []Skill) {
	o.Kinds = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SkillDetails) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillDetails) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SkillDetails) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SkillDetails) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o SkillDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkillDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Skill) {
		toSerialize["skill"] = o.Skill
	}
	if !IsNil(o.Children) {
		toSerialize["children"] = o.Children
	}
	if !IsNil(o.Parents) {
		toSerialize["parents"] = o.Parents
	}
	if !IsNil(o.Kinds) {
		toSerialize["kinds"] = o.Kinds
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	return toSerialize, nil
}

type NullableSkillDetails struct {
	value *SkillDetails
	isSet bool
}

func (v NullableSkillDetails) Get() *SkillDetails {
	return v.value
}

func (v *NullableSkillDetails) Set(val *SkillDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSkillDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSkillDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkillDetails(val *SkillDetails) *NullableSkillDetails {
	return &NullableSkillDetails{value: val, isSet: true}
}

func (v NullableSkillDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkillDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


