/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 0.8.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ProjectParticipationAllOf struct for ProjectParticipationAllOf
type ProjectParticipationAllOf struct {
	Timeframe *Timeframed `json:"timeframe,omitempty"`
	Project Project `json:"project"`
	DescriptionOverwrite *string `json:"descriptionOverwrite,omitempty"`
	PersonalDescription *string `json:"personalDescription,omitempty"`
	Experiences []Experience `json:"experiences"`
}

// NewProjectParticipationAllOf instantiates a new ProjectParticipationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectParticipationAllOf(project Project, experiences []Experience) *ProjectParticipationAllOf {
	this := ProjectParticipationAllOf{}
	this.Project = project
	this.Experiences = experiences
	return &this
}

// NewProjectParticipationAllOfWithDefaults instantiates a new ProjectParticipationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectParticipationAllOfWithDefaults() *ProjectParticipationAllOf {
	this := ProjectParticipationAllOf{}
	return &this
}

// GetTimeframe returns the Timeframe field value if set, zero value otherwise.
func (o *ProjectParticipationAllOf) GetTimeframe() Timeframed {
	if o == nil || o.Timeframe == nil {
		var ret Timeframed
		return ret
	}
	return *o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipationAllOf) GetTimeframeOk() (*Timeframed, bool) {
	if o == nil || o.Timeframe == nil {
		return nil, false
	}
	return o.Timeframe, true
}

// HasTimeframe returns a boolean if a field has been set.
func (o *ProjectParticipationAllOf) HasTimeframe() bool {
	if o != nil && o.Timeframe != nil {
		return true
	}

	return false
}

// SetTimeframe gets a reference to the given Timeframed and assigns it to the Timeframe field.
func (o *ProjectParticipationAllOf) SetTimeframe(v Timeframed) {
	o.Timeframe = &v
}

// GetProject returns the Project field value
func (o *ProjectParticipationAllOf) GetProject() Project {
	if o == nil {
		var ret Project
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *ProjectParticipationAllOf) GetProjectOk() (*Project, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *ProjectParticipationAllOf) SetProject(v Project) {
	o.Project = v
}

// GetDescriptionOverwrite returns the DescriptionOverwrite field value if set, zero value otherwise.
func (o *ProjectParticipationAllOf) GetDescriptionOverwrite() string {
	if o == nil || o.DescriptionOverwrite == nil {
		var ret string
		return ret
	}
	return *o.DescriptionOverwrite
}

// GetDescriptionOverwriteOk returns a tuple with the DescriptionOverwrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipationAllOf) GetDescriptionOverwriteOk() (*string, bool) {
	if o == nil || o.DescriptionOverwrite == nil {
		return nil, false
	}
	return o.DescriptionOverwrite, true
}

// HasDescriptionOverwrite returns a boolean if a field has been set.
func (o *ProjectParticipationAllOf) HasDescriptionOverwrite() bool {
	if o != nil && o.DescriptionOverwrite != nil {
		return true
	}

	return false
}

// SetDescriptionOverwrite gets a reference to the given string and assigns it to the DescriptionOverwrite field.
func (o *ProjectParticipationAllOf) SetDescriptionOverwrite(v string) {
	o.DescriptionOverwrite = &v
}

// GetPersonalDescription returns the PersonalDescription field value if set, zero value otherwise.
func (o *ProjectParticipationAllOf) GetPersonalDescription() string {
	if o == nil || o.PersonalDescription == nil {
		var ret string
		return ret
	}
	return *o.PersonalDescription
}

// GetPersonalDescriptionOk returns a tuple with the PersonalDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipationAllOf) GetPersonalDescriptionOk() (*string, bool) {
	if o == nil || o.PersonalDescription == nil {
		return nil, false
	}
	return o.PersonalDescription, true
}

// HasPersonalDescription returns a boolean if a field has been set.
func (o *ProjectParticipationAllOf) HasPersonalDescription() bool {
	if o != nil && o.PersonalDescription != nil {
		return true
	}

	return false
}

// SetPersonalDescription gets a reference to the given string and assigns it to the PersonalDescription field.
func (o *ProjectParticipationAllOf) SetPersonalDescription(v string) {
	o.PersonalDescription = &v
}

// GetExperiences returns the Experiences field value
func (o *ProjectParticipationAllOf) GetExperiences() []Experience {
	if o == nil {
		var ret []Experience
		return ret
	}

	return o.Experiences
}

// GetExperiencesOk returns a tuple with the Experiences field value
// and a boolean to check if the value has been set.
func (o *ProjectParticipationAllOf) GetExperiencesOk() (*[]Experience, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Experiences, true
}

// SetExperiences sets field value
func (o *ProjectParticipationAllOf) SetExperiences(v []Experience) {
	o.Experiences = v
}

func (o ProjectParticipationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Timeframe != nil {
		toSerialize["timeframe"] = o.Timeframe
	}
	if true {
		toSerialize["project"] = o.Project
	}
	if o.DescriptionOverwrite != nil {
		toSerialize["descriptionOverwrite"] = o.DescriptionOverwrite
	}
	if o.PersonalDescription != nil {
		toSerialize["personalDescription"] = o.PersonalDescription
	}
	if true {
		toSerialize["experiences"] = o.Experiences
	}
	return json.Marshal(toSerialize)
}

type NullableProjectParticipationAllOf struct {
	value *ProjectParticipationAllOf
	isSet bool
}

func (v NullableProjectParticipationAllOf) Get() *ProjectParticipationAllOf {
	return v.value
}

func (v *NullableProjectParticipationAllOf) Set(val *ProjectParticipationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectParticipationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectParticipationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectParticipationAllOf(val *ProjectParticipationAllOf) *NullableProjectParticipationAllOf {
	return &NullableProjectParticipationAllOf{value: val, isSet: true}
}

func (v NullableProjectParticipationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectParticipationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


