/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.10.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ProjectSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectSearch{}

// ProjectSearch struct for ProjectSearch
type ProjectSearch struct {
	ProjectIds []string `json:"projectIds,omitempty"`
	MinStartDate *string `json:"minStartDate,omitempty"`
	MaxEndDate *string `json:"maxEndDate,omitempty"`
	ProjectStatuses []ProjectStatus `json:"projectStatuses,omitempty"`
	ParticipationAmountInMonths *MinMax `json:"participationAmountInMonths,omitempty"`
	InvolvedOfficeIds []string `json:"involvedOfficeIds,omitempty"`
	AmountOfInvolvedPersons *MinMax `json:"amountOfInvolvedPersons,omitempty"`
	Skills []EntityFilter `json:"skills,omitempty"`
	ProjectType []ProjectType `json:"projectType,omitempty"`
	Confidentiality []Confidentiality `json:"confidentiality,omitempty"`
	Participants []EntityFilter `json:"participants,omitempty"`
	Industries []EntityFilter `json:"industries,omitempty"`
	Organizations []EntityFilter `json:"organizations,omitempty"`
}

// NewProjectSearch instantiates a new ProjectSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectSearch() *ProjectSearch {
	this := ProjectSearch{}
	return &this
}

// NewProjectSearchWithDefaults instantiates a new ProjectSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectSearchWithDefaults() *ProjectSearch {
	this := ProjectSearch{}
	return &this
}

// GetProjectIds returns the ProjectIds field value if set, zero value otherwise.
func (o *ProjectSearch) GetProjectIds() []string {
	if o == nil || IsNil(o.ProjectIds) {
		var ret []string
		return ret
	}
	return o.ProjectIds
}

// GetProjectIdsOk returns a tuple with the ProjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetProjectIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProjectIds) {
		return nil, false
	}
	return o.ProjectIds, true
}

// HasProjectIds returns a boolean if a field has been set.
func (o *ProjectSearch) HasProjectIds() bool {
	if o != nil && !IsNil(o.ProjectIds) {
		return true
	}

	return false
}

// SetProjectIds gets a reference to the given []string and assigns it to the ProjectIds field.
func (o *ProjectSearch) SetProjectIds(v []string) {
	o.ProjectIds = v
}

// GetMinStartDate returns the MinStartDate field value if set, zero value otherwise.
func (o *ProjectSearch) GetMinStartDate() string {
	if o == nil || IsNil(o.MinStartDate) {
		var ret string
		return ret
	}
	return *o.MinStartDate
}

// GetMinStartDateOk returns a tuple with the MinStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetMinStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.MinStartDate) {
		return nil, false
	}
	return o.MinStartDate, true
}

// HasMinStartDate returns a boolean if a field has been set.
func (o *ProjectSearch) HasMinStartDate() bool {
	if o != nil && !IsNil(o.MinStartDate) {
		return true
	}

	return false
}

// SetMinStartDate gets a reference to the given string and assigns it to the MinStartDate field.
func (o *ProjectSearch) SetMinStartDate(v string) {
	o.MinStartDate = &v
}

// GetMaxEndDate returns the MaxEndDate field value if set, zero value otherwise.
func (o *ProjectSearch) GetMaxEndDate() string {
	if o == nil || IsNil(o.MaxEndDate) {
		var ret string
		return ret
	}
	return *o.MaxEndDate
}

// GetMaxEndDateOk returns a tuple with the MaxEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetMaxEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.MaxEndDate) {
		return nil, false
	}
	return o.MaxEndDate, true
}

// HasMaxEndDate returns a boolean if a field has been set.
func (o *ProjectSearch) HasMaxEndDate() bool {
	if o != nil && !IsNil(o.MaxEndDate) {
		return true
	}

	return false
}

// SetMaxEndDate gets a reference to the given string and assigns it to the MaxEndDate field.
func (o *ProjectSearch) SetMaxEndDate(v string) {
	o.MaxEndDate = &v
}

// GetProjectStatuses returns the ProjectStatuses field value if set, zero value otherwise.
func (o *ProjectSearch) GetProjectStatuses() []ProjectStatus {
	if o == nil || IsNil(o.ProjectStatuses) {
		var ret []ProjectStatus
		return ret
	}
	return o.ProjectStatuses
}

// GetProjectStatusesOk returns a tuple with the ProjectStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetProjectStatusesOk() ([]ProjectStatus, bool) {
	if o == nil || IsNil(o.ProjectStatuses) {
		return nil, false
	}
	return o.ProjectStatuses, true
}

// HasProjectStatuses returns a boolean if a field has been set.
func (o *ProjectSearch) HasProjectStatuses() bool {
	if o != nil && !IsNil(o.ProjectStatuses) {
		return true
	}

	return false
}

// SetProjectStatuses gets a reference to the given []ProjectStatus and assigns it to the ProjectStatuses field.
func (o *ProjectSearch) SetProjectStatuses(v []ProjectStatus) {
	o.ProjectStatuses = v
}

// GetParticipationAmountInMonths returns the ParticipationAmountInMonths field value if set, zero value otherwise.
func (o *ProjectSearch) GetParticipationAmountInMonths() MinMax {
	if o == nil || IsNil(o.ParticipationAmountInMonths) {
		var ret MinMax
		return ret
	}
	return *o.ParticipationAmountInMonths
}

// GetParticipationAmountInMonthsOk returns a tuple with the ParticipationAmountInMonths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetParticipationAmountInMonthsOk() (*MinMax, bool) {
	if o == nil || IsNil(o.ParticipationAmountInMonths) {
		return nil, false
	}
	return o.ParticipationAmountInMonths, true
}

// HasParticipationAmountInMonths returns a boolean if a field has been set.
func (o *ProjectSearch) HasParticipationAmountInMonths() bool {
	if o != nil && !IsNil(o.ParticipationAmountInMonths) {
		return true
	}

	return false
}

// SetParticipationAmountInMonths gets a reference to the given MinMax and assigns it to the ParticipationAmountInMonths field.
func (o *ProjectSearch) SetParticipationAmountInMonths(v MinMax) {
	o.ParticipationAmountInMonths = &v
}

// GetInvolvedOfficeIds returns the InvolvedOfficeIds field value if set, zero value otherwise.
func (o *ProjectSearch) GetInvolvedOfficeIds() []string {
	if o == nil || IsNil(o.InvolvedOfficeIds) {
		var ret []string
		return ret
	}
	return o.InvolvedOfficeIds
}

// GetInvolvedOfficeIdsOk returns a tuple with the InvolvedOfficeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetInvolvedOfficeIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.InvolvedOfficeIds) {
		return nil, false
	}
	return o.InvolvedOfficeIds, true
}

// HasInvolvedOfficeIds returns a boolean if a field has been set.
func (o *ProjectSearch) HasInvolvedOfficeIds() bool {
	if o != nil && !IsNil(o.InvolvedOfficeIds) {
		return true
	}

	return false
}

// SetInvolvedOfficeIds gets a reference to the given []string and assigns it to the InvolvedOfficeIds field.
func (o *ProjectSearch) SetInvolvedOfficeIds(v []string) {
	o.InvolvedOfficeIds = v
}

// GetAmountOfInvolvedPersons returns the AmountOfInvolvedPersons field value if set, zero value otherwise.
func (o *ProjectSearch) GetAmountOfInvolvedPersons() MinMax {
	if o == nil || IsNil(o.AmountOfInvolvedPersons) {
		var ret MinMax
		return ret
	}
	return *o.AmountOfInvolvedPersons
}

// GetAmountOfInvolvedPersonsOk returns a tuple with the AmountOfInvolvedPersons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetAmountOfInvolvedPersonsOk() (*MinMax, bool) {
	if o == nil || IsNil(o.AmountOfInvolvedPersons) {
		return nil, false
	}
	return o.AmountOfInvolvedPersons, true
}

// HasAmountOfInvolvedPersons returns a boolean if a field has been set.
func (o *ProjectSearch) HasAmountOfInvolvedPersons() bool {
	if o != nil && !IsNil(o.AmountOfInvolvedPersons) {
		return true
	}

	return false
}

// SetAmountOfInvolvedPersons gets a reference to the given MinMax and assigns it to the AmountOfInvolvedPersons field.
func (o *ProjectSearch) SetAmountOfInvolvedPersons(v MinMax) {
	o.AmountOfInvolvedPersons = &v
}

// GetSkills returns the Skills field value if set, zero value otherwise.
func (o *ProjectSearch) GetSkills() []EntityFilter {
	if o == nil || IsNil(o.Skills) {
		var ret []EntityFilter
		return ret
	}
	return o.Skills
}

// GetSkillsOk returns a tuple with the Skills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetSkillsOk() ([]EntityFilter, bool) {
	if o == nil || IsNil(o.Skills) {
		return nil, false
	}
	return o.Skills, true
}

// HasSkills returns a boolean if a field has been set.
func (o *ProjectSearch) HasSkills() bool {
	if o != nil && !IsNil(o.Skills) {
		return true
	}

	return false
}

// SetSkills gets a reference to the given []EntityFilter and assigns it to the Skills field.
func (o *ProjectSearch) SetSkills(v []EntityFilter) {
	o.Skills = v
}

// GetProjectType returns the ProjectType field value if set, zero value otherwise.
func (o *ProjectSearch) GetProjectType() []ProjectType {
	if o == nil || IsNil(o.ProjectType) {
		var ret []ProjectType
		return ret
	}
	return o.ProjectType
}

// GetProjectTypeOk returns a tuple with the ProjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetProjectTypeOk() ([]ProjectType, bool) {
	if o == nil || IsNil(o.ProjectType) {
		return nil, false
	}
	return o.ProjectType, true
}

// HasProjectType returns a boolean if a field has been set.
func (o *ProjectSearch) HasProjectType() bool {
	if o != nil && !IsNil(o.ProjectType) {
		return true
	}

	return false
}

// SetProjectType gets a reference to the given []ProjectType and assigns it to the ProjectType field.
func (o *ProjectSearch) SetProjectType(v []ProjectType) {
	o.ProjectType = v
}

// GetConfidentiality returns the Confidentiality field value if set, zero value otherwise.
func (o *ProjectSearch) GetConfidentiality() []Confidentiality {
	if o == nil || IsNil(o.Confidentiality) {
		var ret []Confidentiality
		return ret
	}
	return o.Confidentiality
}

// GetConfidentialityOk returns a tuple with the Confidentiality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetConfidentialityOk() ([]Confidentiality, bool) {
	if o == nil || IsNil(o.Confidentiality) {
		return nil, false
	}
	return o.Confidentiality, true
}

// HasConfidentiality returns a boolean if a field has been set.
func (o *ProjectSearch) HasConfidentiality() bool {
	if o != nil && !IsNil(o.Confidentiality) {
		return true
	}

	return false
}

// SetConfidentiality gets a reference to the given []Confidentiality and assigns it to the Confidentiality field.
func (o *ProjectSearch) SetConfidentiality(v []Confidentiality) {
	o.Confidentiality = v
}

// GetParticipants returns the Participants field value if set, zero value otherwise.
func (o *ProjectSearch) GetParticipants() []EntityFilter {
	if o == nil || IsNil(o.Participants) {
		var ret []EntityFilter
		return ret
	}
	return o.Participants
}

// GetParticipantsOk returns a tuple with the Participants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetParticipantsOk() ([]EntityFilter, bool) {
	if o == nil || IsNil(o.Participants) {
		return nil, false
	}
	return o.Participants, true
}

// HasParticipants returns a boolean if a field has been set.
func (o *ProjectSearch) HasParticipants() bool {
	if o != nil && !IsNil(o.Participants) {
		return true
	}

	return false
}

// SetParticipants gets a reference to the given []EntityFilter and assigns it to the Participants field.
func (o *ProjectSearch) SetParticipants(v []EntityFilter) {
	o.Participants = v
}

// GetIndustries returns the Industries field value if set, zero value otherwise.
func (o *ProjectSearch) GetIndustries() []EntityFilter {
	if o == nil || IsNil(o.Industries) {
		var ret []EntityFilter
		return ret
	}
	return o.Industries
}

// GetIndustriesOk returns a tuple with the Industries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetIndustriesOk() ([]EntityFilter, bool) {
	if o == nil || IsNil(o.Industries) {
		return nil, false
	}
	return o.Industries, true
}

// HasIndustries returns a boolean if a field has been set.
func (o *ProjectSearch) HasIndustries() bool {
	if o != nil && !IsNil(o.Industries) {
		return true
	}

	return false
}

// SetIndustries gets a reference to the given []EntityFilter and assigns it to the Industries field.
func (o *ProjectSearch) SetIndustries(v []EntityFilter) {
	o.Industries = v
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *ProjectSearch) GetOrganizations() []EntityFilter {
	if o == nil || IsNil(o.Organizations) {
		var ret []EntityFilter
		return ret
	}
	return o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSearch) GetOrganizationsOk() ([]EntityFilter, bool) {
	if o == nil || IsNil(o.Organizations) {
		return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *ProjectSearch) HasOrganizations() bool {
	if o != nil && !IsNil(o.Organizations) {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given []EntityFilter and assigns it to the Organizations field.
func (o *ProjectSearch) SetOrganizations(v []EntityFilter) {
	o.Organizations = v
}

func (o ProjectSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectIds) {
		toSerialize["projectIds"] = o.ProjectIds
	}
	if !IsNil(o.MinStartDate) {
		toSerialize["minStartDate"] = o.MinStartDate
	}
	if !IsNil(o.MaxEndDate) {
		toSerialize["maxEndDate"] = o.MaxEndDate
	}
	if !IsNil(o.ProjectStatuses) {
		toSerialize["projectStatuses"] = o.ProjectStatuses
	}
	if !IsNil(o.ParticipationAmountInMonths) {
		toSerialize["participationAmountInMonths"] = o.ParticipationAmountInMonths
	}
	if !IsNil(o.InvolvedOfficeIds) {
		toSerialize["involvedOfficeIds"] = o.InvolvedOfficeIds
	}
	if !IsNil(o.AmountOfInvolvedPersons) {
		toSerialize["amountOfInvolvedPersons"] = o.AmountOfInvolvedPersons
	}
	if !IsNil(o.Skills) {
		toSerialize["skills"] = o.Skills
	}
	if !IsNil(o.ProjectType) {
		toSerialize["projectType"] = o.ProjectType
	}
	if !IsNil(o.Confidentiality) {
		toSerialize["confidentiality"] = o.Confidentiality
	}
	if !IsNil(o.Participants) {
		toSerialize["participants"] = o.Participants
	}
	if !IsNil(o.Industries) {
		toSerialize["industries"] = o.Industries
	}
	if !IsNil(o.Organizations) {
		toSerialize["organizations"] = o.Organizations
	}
	return toSerialize, nil
}

type NullableProjectSearch struct {
	value *ProjectSearch
	isSet bool
}

func (v NullableProjectSearch) Get() *ProjectSearch {
	return v.value
}

func (v *NullableProjectSearch) Set(val *ProjectSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectSearch(val *ProjectSearch) *NullableProjectSearch {
	return &NullableProjectSearch{value: val, isSet: true}
}

func (v NullableProjectSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


