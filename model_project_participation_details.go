/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ProjectParticipationDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectParticipationDetails{}

// ProjectParticipationDetails struct for ProjectParticipationDetails
type ProjectParticipationDetails struct {
	Audit *Audit `json:"audit,omitempty"`
	Participation *ProjectParticipation `json:"participation,omitempty"`
	SkillGroups []ExperienceSkillGroup `json:"skillGroups,omitempty"`
	Person *Person `json:"person,omitempty"`
	Project *Project `json:"project,omitempty"`
}

// NewProjectParticipationDetails instantiates a new ProjectParticipationDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectParticipationDetails() *ProjectParticipationDetails {
	this := ProjectParticipationDetails{}
	return &this
}

// NewProjectParticipationDetailsWithDefaults instantiates a new ProjectParticipationDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectParticipationDetailsWithDefaults() *ProjectParticipationDetails {
	this := ProjectParticipationDetails{}
	return &this
}

// GetAudit returns the Audit field value if set, zero value otherwise.
func (o *ProjectParticipationDetails) GetAudit() Audit {
	if o == nil || IsNil(o.Audit) {
		var ret Audit
		return ret
	}
	return *o.Audit
}

// GetAuditOk returns a tuple with the Audit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipationDetails) GetAuditOk() (*Audit, bool) {
	if o == nil || IsNil(o.Audit) {
		return nil, false
	}
	return o.Audit, true
}

// HasAudit returns a boolean if a field has been set.
func (o *ProjectParticipationDetails) HasAudit() bool {
	if o != nil && !IsNil(o.Audit) {
		return true
	}

	return false
}

// SetAudit gets a reference to the given Audit and assigns it to the Audit field.
func (o *ProjectParticipationDetails) SetAudit(v Audit) {
	o.Audit = &v
}

// GetParticipation returns the Participation field value if set, zero value otherwise.
func (o *ProjectParticipationDetails) GetParticipation() ProjectParticipation {
	if o == nil || IsNil(o.Participation) {
		var ret ProjectParticipation
		return ret
	}
	return *o.Participation
}

// GetParticipationOk returns a tuple with the Participation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipationDetails) GetParticipationOk() (*ProjectParticipation, bool) {
	if o == nil || IsNil(o.Participation) {
		return nil, false
	}
	return o.Participation, true
}

// HasParticipation returns a boolean if a field has been set.
func (o *ProjectParticipationDetails) HasParticipation() bool {
	if o != nil && !IsNil(o.Participation) {
		return true
	}

	return false
}

// SetParticipation gets a reference to the given ProjectParticipation and assigns it to the Participation field.
func (o *ProjectParticipationDetails) SetParticipation(v ProjectParticipation) {
	o.Participation = &v
}

// GetSkillGroups returns the SkillGroups field value if set, zero value otherwise.
func (o *ProjectParticipationDetails) GetSkillGroups() []ExperienceSkillGroup {
	if o == nil || IsNil(o.SkillGroups) {
		var ret []ExperienceSkillGroup
		return ret
	}
	return o.SkillGroups
}

// GetSkillGroupsOk returns a tuple with the SkillGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipationDetails) GetSkillGroupsOk() ([]ExperienceSkillGroup, bool) {
	if o == nil || IsNil(o.SkillGroups) {
		return nil, false
	}
	return o.SkillGroups, true
}

// HasSkillGroups returns a boolean if a field has been set.
func (o *ProjectParticipationDetails) HasSkillGroups() bool {
	if o != nil && !IsNil(o.SkillGroups) {
		return true
	}

	return false
}

// SetSkillGroups gets a reference to the given []ExperienceSkillGroup and assigns it to the SkillGroups field.
func (o *ProjectParticipationDetails) SetSkillGroups(v []ExperienceSkillGroup) {
	o.SkillGroups = v
}

// GetPerson returns the Person field value if set, zero value otherwise.
func (o *ProjectParticipationDetails) GetPerson() Person {
	if o == nil || IsNil(o.Person) {
		var ret Person
		return ret
	}
	return *o.Person
}

// GetPersonOk returns a tuple with the Person field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipationDetails) GetPersonOk() (*Person, bool) {
	if o == nil || IsNil(o.Person) {
		return nil, false
	}
	return o.Person, true
}

// HasPerson returns a boolean if a field has been set.
func (o *ProjectParticipationDetails) HasPerson() bool {
	if o != nil && !IsNil(o.Person) {
		return true
	}

	return false
}

// SetPerson gets a reference to the given Person and assigns it to the Person field.
func (o *ProjectParticipationDetails) SetPerson(v Person) {
	o.Person = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ProjectParticipationDetails) GetProject() Project {
	if o == nil || IsNil(o.Project) {
		var ret Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipationDetails) GetProjectOk() (*Project, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ProjectParticipationDetails) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Project and assigns it to the Project field.
func (o *ProjectParticipationDetails) SetProject(v Project) {
	o.Project = &v
}

func (o ProjectParticipationDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectParticipationDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Audit) {
		toSerialize["audit"] = o.Audit
	}
	if !IsNil(o.Participation) {
		toSerialize["participation"] = o.Participation
	}
	if !IsNil(o.SkillGroups) {
		toSerialize["skillGroups"] = o.SkillGroups
	}
	if !IsNil(o.Person) {
		toSerialize["person"] = o.Person
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	return toSerialize, nil
}

type NullableProjectParticipationDetails struct {
	value *ProjectParticipationDetails
	isSet bool
}

func (v NullableProjectParticipationDetails) Get() *ProjectParticipationDetails {
	return v.value
}

func (v *NullableProjectParticipationDetails) Set(val *ProjectParticipationDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectParticipationDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectParticipationDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectParticipationDetails(val *ProjectParticipationDetails) *NullableProjectParticipationDetails {
	return &NullableProjectParticipationDetails{value: val, isSet: true}
}

func (v NullableProjectParticipationDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectParticipationDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


