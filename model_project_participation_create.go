/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ProjectParticipationCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectParticipationCreate{}

// ProjectParticipationCreate struct for ProjectParticipationCreate
type ProjectParticipationCreate struct {
	Timeframe *Timeframed `json:"timeframe,omitempty"`
	Skills []SkillLevelUpdate `json:"skills,omitempty"`
	PersonalDescription *string `json:"personalDescription,omitempty"`
	Id string `json:"id"`
	ProjectId string `json:"projectId"`
	PersonId string `json:"personId"`
}

// NewProjectParticipationCreate instantiates a new ProjectParticipationCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectParticipationCreate(id string, projectId string, personId string) *ProjectParticipationCreate {
	this := ProjectParticipationCreate{}
	this.Id = id
	this.ProjectId = projectId
	this.PersonId = personId
	return &this
}

// NewProjectParticipationCreateWithDefaults instantiates a new ProjectParticipationCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectParticipationCreateWithDefaults() *ProjectParticipationCreate {
	this := ProjectParticipationCreate{}
	return &this
}

// GetTimeframe returns the Timeframe field value if set, zero value otherwise.
func (o *ProjectParticipationCreate) GetTimeframe() Timeframed {
	if o == nil || IsNil(o.Timeframe) {
		var ret Timeframed
		return ret
	}
	return *o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipationCreate) GetTimeframeOk() (*Timeframed, bool) {
	if o == nil || IsNil(o.Timeframe) {
		return nil, false
	}
	return o.Timeframe, true
}

// HasTimeframe returns a boolean if a field has been set.
func (o *ProjectParticipationCreate) HasTimeframe() bool {
	if o != nil && !IsNil(o.Timeframe) {
		return true
	}

	return false
}

// SetTimeframe gets a reference to the given Timeframed and assigns it to the Timeframe field.
func (o *ProjectParticipationCreate) SetTimeframe(v Timeframed) {
	o.Timeframe = &v
}

// GetSkills returns the Skills field value if set, zero value otherwise.
func (o *ProjectParticipationCreate) GetSkills() []SkillLevelUpdate {
	if o == nil || IsNil(o.Skills) {
		var ret []SkillLevelUpdate
		return ret
	}
	return o.Skills
}

// GetSkillsOk returns a tuple with the Skills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipationCreate) GetSkillsOk() ([]SkillLevelUpdate, bool) {
	if o == nil || IsNil(o.Skills) {
		return nil, false
	}
	return o.Skills, true
}

// HasSkills returns a boolean if a field has been set.
func (o *ProjectParticipationCreate) HasSkills() bool {
	if o != nil && !IsNil(o.Skills) {
		return true
	}

	return false
}

// SetSkills gets a reference to the given []SkillLevelUpdate and assigns it to the Skills field.
func (o *ProjectParticipationCreate) SetSkills(v []SkillLevelUpdate) {
	o.Skills = v
}

// GetPersonalDescription returns the PersonalDescription field value if set, zero value otherwise.
func (o *ProjectParticipationCreate) GetPersonalDescription() string {
	if o == nil || IsNil(o.PersonalDescription) {
		var ret string
		return ret
	}
	return *o.PersonalDescription
}

// GetPersonalDescriptionOk returns a tuple with the PersonalDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipationCreate) GetPersonalDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PersonalDescription) {
		return nil, false
	}
	return o.PersonalDescription, true
}

// HasPersonalDescription returns a boolean if a field has been set.
func (o *ProjectParticipationCreate) HasPersonalDescription() bool {
	if o != nil && !IsNil(o.PersonalDescription) {
		return true
	}

	return false
}

// SetPersonalDescription gets a reference to the given string and assigns it to the PersonalDescription field.
func (o *ProjectParticipationCreate) SetPersonalDescription(v string) {
	o.PersonalDescription = &v
}

// GetId returns the Id field value
func (o *ProjectParticipationCreate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectParticipationCreate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectParticipationCreate) SetId(v string) {
	o.Id = v
}

// GetProjectId returns the ProjectId field value
func (o *ProjectParticipationCreate) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *ProjectParticipationCreate) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *ProjectParticipationCreate) SetProjectId(v string) {
	o.ProjectId = v
}

// GetPersonId returns the PersonId field value
func (o *ProjectParticipationCreate) GetPersonId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PersonId
}

// GetPersonIdOk returns a tuple with the PersonId field value
// and a boolean to check if the value has been set.
func (o *ProjectParticipationCreate) GetPersonIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PersonId, true
}

// SetPersonId sets field value
func (o *ProjectParticipationCreate) SetPersonId(v string) {
	o.PersonId = v
}

func (o ProjectParticipationCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectParticipationCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timeframe) {
		toSerialize["timeframe"] = o.Timeframe
	}
	if !IsNil(o.Skills) {
		toSerialize["skills"] = o.Skills
	}
	if !IsNil(o.PersonalDescription) {
		toSerialize["personalDescription"] = o.PersonalDescription
	}
	toSerialize["id"] = o.Id
	toSerialize["projectId"] = o.ProjectId
	toSerialize["personId"] = o.PersonId
	return toSerialize, nil
}

type NullableProjectParticipationCreate struct {
	value *ProjectParticipationCreate
	isSet bool
}

func (v NullableProjectParticipationCreate) Get() *ProjectParticipationCreate {
	return v.value
}

func (v *NullableProjectParticipationCreate) Set(val *ProjectParticipationCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectParticipationCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectParticipationCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectParticipationCreate(val *ProjectParticipationCreate) *NullableProjectParticipationCreate {
	return &NullableProjectParticipationCreate{value: val, isSet: true}
}

func (v NullableProjectParticipationCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectParticipationCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


