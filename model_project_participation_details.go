/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.27.1
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
	Participation *ProjectParticipation `json:"participation,omitempty"`
	Experiences []Experience `json:"experiences,omitempty"`
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

// GetExperiences returns the Experiences field value if set, zero value otherwise.
func (o *ProjectParticipationDetails) GetExperiences() []Experience {
	if o == nil || IsNil(o.Experiences) {
		var ret []Experience
		return ret
	}
	return o.Experiences
}

// GetExperiencesOk returns a tuple with the Experiences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipationDetails) GetExperiencesOk() ([]Experience, bool) {
	if o == nil || IsNil(o.Experiences) {
		return nil, false
	}
	return o.Experiences, true
}

// HasExperiences returns a boolean if a field has been set.
func (o *ProjectParticipationDetails) HasExperiences() bool {
	if o != nil && !IsNil(o.Experiences) {
		return true
	}

	return false
}

// SetExperiences gets a reference to the given []Experience and assigns it to the Experiences field.
func (o *ProjectParticipationDetails) SetExperiences(v []Experience) {
	o.Experiences = v
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
	if !IsNil(o.Participation) {
		toSerialize["participation"] = o.Participation
	}
	if !IsNil(o.Experiences) {
		toSerialize["experiences"] = o.Experiences
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


