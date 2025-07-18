/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.74.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PersonSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonSearch{}

// PersonSearch struct for PersonSearch
type PersonSearch struct {
	Search
	CountryIds []string `json:"countryIds,omitempty"`
	PersonIds []string `json:"personIds,omitempty"`
	EmployeeIds []string `json:"employeeIds,omitempty"`
	ProfileIds []string `json:"profileIds,omitempty"`
	OfficeIds []string `json:"officeIds,omitempty"`
	LanguageIds []string `json:"languageIds,omitempty"`
	Availability *AvailabilityFilter `json:"availability,omitempty"`
	OnsiteRatio *MinMaxPercent `json:"onsiteRatio,omitempty"`
	Seniority []Seniority `json:"seniority,omitempty"`
	Skills []PersonSkillFilter `json:"skills,omitempty"`
	ProjectIds []string `json:"projectIds,omitempty"`
	OrganizationIds []string `json:"organizationIds,omitempty"`
	IndustryIds []string `json:"industryIds,omitempty"`
	CertificationIds []string `json:"certificationIds,omitempty"`
	AwardIds []string `json:"awardIds,omitempty"`
	AwardedProject *AwardedProject `json:"awardedProject,omitempty"`
	Inactive *bool `json:"inactive,omitempty"`
}

// NewPersonSearch instantiates a new PersonSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonSearch(skip int32, limit int32) *PersonSearch {
	this := PersonSearch{}
	this.Skip = skip
	this.Limit = limit
	return &this
}

// NewPersonSearchWithDefaults instantiates a new PersonSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonSearchWithDefaults() *PersonSearch {
	this := PersonSearch{}
	return &this
}

// GetCountryIds returns the CountryIds field value if set, zero value otherwise.
func (o *PersonSearch) GetCountryIds() []string {
	if o == nil || IsNil(o.CountryIds) {
		var ret []string
		return ret
	}
	return o.CountryIds
}

// GetCountryIdsOk returns a tuple with the CountryIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetCountryIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CountryIds) {
		return nil, false
	}
	return o.CountryIds, true
}

// HasCountryIds returns a boolean if a field has been set.
func (o *PersonSearch) HasCountryIds() bool {
	if o != nil && !IsNil(o.CountryIds) {
		return true
	}

	return false
}

// SetCountryIds gets a reference to the given []string and assigns it to the CountryIds field.
func (o *PersonSearch) SetCountryIds(v []string) {
	o.CountryIds = v
}

// GetPersonIds returns the PersonIds field value if set, zero value otherwise.
func (o *PersonSearch) GetPersonIds() []string {
	if o == nil || IsNil(o.PersonIds) {
		var ret []string
		return ret
	}
	return o.PersonIds
}

// GetPersonIdsOk returns a tuple with the PersonIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetPersonIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.PersonIds) {
		return nil, false
	}
	return o.PersonIds, true
}

// HasPersonIds returns a boolean if a field has been set.
func (o *PersonSearch) HasPersonIds() bool {
	if o != nil && !IsNil(o.PersonIds) {
		return true
	}

	return false
}

// SetPersonIds gets a reference to the given []string and assigns it to the PersonIds field.
func (o *PersonSearch) SetPersonIds(v []string) {
	o.PersonIds = v
}

// GetEmployeeIds returns the EmployeeIds field value if set, zero value otherwise.
func (o *PersonSearch) GetEmployeeIds() []string {
	if o == nil || IsNil(o.EmployeeIds) {
		var ret []string
		return ret
	}
	return o.EmployeeIds
}

// GetEmployeeIdsOk returns a tuple with the EmployeeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetEmployeeIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.EmployeeIds) {
		return nil, false
	}
	return o.EmployeeIds, true
}

// HasEmployeeIds returns a boolean if a field has been set.
func (o *PersonSearch) HasEmployeeIds() bool {
	if o != nil && !IsNil(o.EmployeeIds) {
		return true
	}

	return false
}

// SetEmployeeIds gets a reference to the given []string and assigns it to the EmployeeIds field.
func (o *PersonSearch) SetEmployeeIds(v []string) {
	o.EmployeeIds = v
}

// GetProfileIds returns the ProfileIds field value if set, zero value otherwise.
func (o *PersonSearch) GetProfileIds() []string {
	if o == nil || IsNil(o.ProfileIds) {
		var ret []string
		return ret
	}
	return o.ProfileIds
}

// GetProfileIdsOk returns a tuple with the ProfileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetProfileIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProfileIds) {
		return nil, false
	}
	return o.ProfileIds, true
}

// HasProfileIds returns a boolean if a field has been set.
func (o *PersonSearch) HasProfileIds() bool {
	if o != nil && !IsNil(o.ProfileIds) {
		return true
	}

	return false
}

// SetProfileIds gets a reference to the given []string and assigns it to the ProfileIds field.
func (o *PersonSearch) SetProfileIds(v []string) {
	o.ProfileIds = v
}

// GetOfficeIds returns the OfficeIds field value if set, zero value otherwise.
func (o *PersonSearch) GetOfficeIds() []string {
	if o == nil || IsNil(o.OfficeIds) {
		var ret []string
		return ret
	}
	return o.OfficeIds
}

// GetOfficeIdsOk returns a tuple with the OfficeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetOfficeIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OfficeIds) {
		return nil, false
	}
	return o.OfficeIds, true
}

// HasOfficeIds returns a boolean if a field has been set.
func (o *PersonSearch) HasOfficeIds() bool {
	if o != nil && !IsNil(o.OfficeIds) {
		return true
	}

	return false
}

// SetOfficeIds gets a reference to the given []string and assigns it to the OfficeIds field.
func (o *PersonSearch) SetOfficeIds(v []string) {
	o.OfficeIds = v
}

// GetLanguageIds returns the LanguageIds field value if set, zero value otherwise.
func (o *PersonSearch) GetLanguageIds() []string {
	if o == nil || IsNil(o.LanguageIds) {
		var ret []string
		return ret
	}
	return o.LanguageIds
}

// GetLanguageIdsOk returns a tuple with the LanguageIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetLanguageIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.LanguageIds) {
		return nil, false
	}
	return o.LanguageIds, true
}

// HasLanguageIds returns a boolean if a field has been set.
func (o *PersonSearch) HasLanguageIds() bool {
	if o != nil && !IsNil(o.LanguageIds) {
		return true
	}

	return false
}

// SetLanguageIds gets a reference to the given []string and assigns it to the LanguageIds field.
func (o *PersonSearch) SetLanguageIds(v []string) {
	o.LanguageIds = v
}

// GetAvailability returns the Availability field value if set, zero value otherwise.
func (o *PersonSearch) GetAvailability() AvailabilityFilter {
	if o == nil || IsNil(o.Availability) {
		var ret AvailabilityFilter
		return ret
	}
	return *o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetAvailabilityOk() (*AvailabilityFilter, bool) {
	if o == nil || IsNil(o.Availability) {
		return nil, false
	}
	return o.Availability, true
}

// HasAvailability returns a boolean if a field has been set.
func (o *PersonSearch) HasAvailability() bool {
	if o != nil && !IsNil(o.Availability) {
		return true
	}

	return false
}

// SetAvailability gets a reference to the given AvailabilityFilter and assigns it to the Availability field.
func (o *PersonSearch) SetAvailability(v AvailabilityFilter) {
	o.Availability = &v
}

// GetOnsiteRatio returns the OnsiteRatio field value if set, zero value otherwise.
func (o *PersonSearch) GetOnsiteRatio() MinMaxPercent {
	if o == nil || IsNil(o.OnsiteRatio) {
		var ret MinMaxPercent
		return ret
	}
	return *o.OnsiteRatio
}

// GetOnsiteRatioOk returns a tuple with the OnsiteRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetOnsiteRatioOk() (*MinMaxPercent, bool) {
	if o == nil || IsNil(o.OnsiteRatio) {
		return nil, false
	}
	return o.OnsiteRatio, true
}

// HasOnsiteRatio returns a boolean if a field has been set.
func (o *PersonSearch) HasOnsiteRatio() bool {
	if o != nil && !IsNil(o.OnsiteRatio) {
		return true
	}

	return false
}

// SetOnsiteRatio gets a reference to the given MinMaxPercent and assigns it to the OnsiteRatio field.
func (o *PersonSearch) SetOnsiteRatio(v MinMaxPercent) {
	o.OnsiteRatio = &v
}

// GetSeniority returns the Seniority field value if set, zero value otherwise.
func (o *PersonSearch) GetSeniority() []Seniority {
	if o == nil || IsNil(o.Seniority) {
		var ret []Seniority
		return ret
	}
	return o.Seniority
}

// GetSeniorityOk returns a tuple with the Seniority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetSeniorityOk() ([]Seniority, bool) {
	if o == nil || IsNil(o.Seniority) {
		return nil, false
	}
	return o.Seniority, true
}

// HasSeniority returns a boolean if a field has been set.
func (o *PersonSearch) HasSeniority() bool {
	if o != nil && !IsNil(o.Seniority) {
		return true
	}

	return false
}

// SetSeniority gets a reference to the given []Seniority and assigns it to the Seniority field.
func (o *PersonSearch) SetSeniority(v []Seniority) {
	o.Seniority = v
}

// GetSkills returns the Skills field value if set, zero value otherwise.
func (o *PersonSearch) GetSkills() []PersonSkillFilter {
	if o == nil || IsNil(o.Skills) {
		var ret []PersonSkillFilter
		return ret
	}
	return o.Skills
}

// GetSkillsOk returns a tuple with the Skills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetSkillsOk() ([]PersonSkillFilter, bool) {
	if o == nil || IsNil(o.Skills) {
		return nil, false
	}
	return o.Skills, true
}

// HasSkills returns a boolean if a field has been set.
func (o *PersonSearch) HasSkills() bool {
	if o != nil && !IsNil(o.Skills) {
		return true
	}

	return false
}

// SetSkills gets a reference to the given []PersonSkillFilter and assigns it to the Skills field.
func (o *PersonSearch) SetSkills(v []PersonSkillFilter) {
	o.Skills = v
}

// GetProjectIds returns the ProjectIds field value if set, zero value otherwise.
func (o *PersonSearch) GetProjectIds() []string {
	if o == nil || IsNil(o.ProjectIds) {
		var ret []string
		return ret
	}
	return o.ProjectIds
}

// GetProjectIdsOk returns a tuple with the ProjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetProjectIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProjectIds) {
		return nil, false
	}
	return o.ProjectIds, true
}

// HasProjectIds returns a boolean if a field has been set.
func (o *PersonSearch) HasProjectIds() bool {
	if o != nil && !IsNil(o.ProjectIds) {
		return true
	}

	return false
}

// SetProjectIds gets a reference to the given []string and assigns it to the ProjectIds field.
func (o *PersonSearch) SetProjectIds(v []string) {
	o.ProjectIds = v
}

// GetOrganizationIds returns the OrganizationIds field value if set, zero value otherwise.
func (o *PersonSearch) GetOrganizationIds() []string {
	if o == nil || IsNil(o.OrganizationIds) {
		var ret []string
		return ret
	}
	return o.OrganizationIds
}

// GetOrganizationIdsOk returns a tuple with the OrganizationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetOrganizationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OrganizationIds) {
		return nil, false
	}
	return o.OrganizationIds, true
}

// HasOrganizationIds returns a boolean if a field has been set.
func (o *PersonSearch) HasOrganizationIds() bool {
	if o != nil && !IsNil(o.OrganizationIds) {
		return true
	}

	return false
}

// SetOrganizationIds gets a reference to the given []string and assigns it to the OrganizationIds field.
func (o *PersonSearch) SetOrganizationIds(v []string) {
	o.OrganizationIds = v
}

// GetIndustryIds returns the IndustryIds field value if set, zero value otherwise.
func (o *PersonSearch) GetIndustryIds() []string {
	if o == nil || IsNil(o.IndustryIds) {
		var ret []string
		return ret
	}
	return o.IndustryIds
}

// GetIndustryIdsOk returns a tuple with the IndustryIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetIndustryIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.IndustryIds) {
		return nil, false
	}
	return o.IndustryIds, true
}

// HasIndustryIds returns a boolean if a field has been set.
func (o *PersonSearch) HasIndustryIds() bool {
	if o != nil && !IsNil(o.IndustryIds) {
		return true
	}

	return false
}

// SetIndustryIds gets a reference to the given []string and assigns it to the IndustryIds field.
func (o *PersonSearch) SetIndustryIds(v []string) {
	o.IndustryIds = v
}

// GetCertificationIds returns the CertificationIds field value if set, zero value otherwise.
func (o *PersonSearch) GetCertificationIds() []string {
	if o == nil || IsNil(o.CertificationIds) {
		var ret []string
		return ret
	}
	return o.CertificationIds
}

// GetCertificationIdsOk returns a tuple with the CertificationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetCertificationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CertificationIds) {
		return nil, false
	}
	return o.CertificationIds, true
}

// HasCertificationIds returns a boolean if a field has been set.
func (o *PersonSearch) HasCertificationIds() bool {
	if o != nil && !IsNil(o.CertificationIds) {
		return true
	}

	return false
}

// SetCertificationIds gets a reference to the given []string and assigns it to the CertificationIds field.
func (o *PersonSearch) SetCertificationIds(v []string) {
	o.CertificationIds = v
}

// GetAwardIds returns the AwardIds field value if set, zero value otherwise.
func (o *PersonSearch) GetAwardIds() []string {
	if o == nil || IsNil(o.AwardIds) {
		var ret []string
		return ret
	}
	return o.AwardIds
}

// GetAwardIdsOk returns a tuple with the AwardIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetAwardIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AwardIds) {
		return nil, false
	}
	return o.AwardIds, true
}

// HasAwardIds returns a boolean if a field has been set.
func (o *PersonSearch) HasAwardIds() bool {
	if o != nil && !IsNil(o.AwardIds) {
		return true
	}

	return false
}

// SetAwardIds gets a reference to the given []string and assigns it to the AwardIds field.
func (o *PersonSearch) SetAwardIds(v []string) {
	o.AwardIds = v
}

// GetAwardedProject returns the AwardedProject field value if set, zero value otherwise.
func (o *PersonSearch) GetAwardedProject() AwardedProject {
	if o == nil || IsNil(o.AwardedProject) {
		var ret AwardedProject
		return ret
	}
	return *o.AwardedProject
}

// GetAwardedProjectOk returns a tuple with the AwardedProject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetAwardedProjectOk() (*AwardedProject, bool) {
	if o == nil || IsNil(o.AwardedProject) {
		return nil, false
	}
	return o.AwardedProject, true
}

// HasAwardedProject returns a boolean if a field has been set.
func (o *PersonSearch) HasAwardedProject() bool {
	if o != nil && !IsNil(o.AwardedProject) {
		return true
	}

	return false
}

// SetAwardedProject gets a reference to the given AwardedProject and assigns it to the AwardedProject field.
func (o *PersonSearch) SetAwardedProject(v AwardedProject) {
	o.AwardedProject = &v
}

// GetInactive returns the Inactive field value if set, zero value otherwise.
func (o *PersonSearch) GetInactive() bool {
	if o == nil || IsNil(o.Inactive) {
		var ret bool
		return ret
	}
	return *o.Inactive
}

// GetInactiveOk returns a tuple with the Inactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetInactiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Inactive) {
		return nil, false
	}
	return o.Inactive, true
}

// HasInactive returns a boolean if a field has been set.
func (o *PersonSearch) HasInactive() bool {
	if o != nil && !IsNil(o.Inactive) {
		return true
	}

	return false
}

// SetInactive gets a reference to the given bool and assigns it to the Inactive field.
func (o *PersonSearch) SetInactive(v bool) {
	o.Inactive = &v
}

func (o PersonSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSearch, errSearch := json.Marshal(o.Search)
	if errSearch != nil {
		return map[string]interface{}{}, errSearch
	}
	errSearch = json.Unmarshal([]byte(serializedSearch), &toSerialize)
	if errSearch != nil {
		return map[string]interface{}{}, errSearch
	}
	if !IsNil(o.CountryIds) {
		toSerialize["countryIds"] = o.CountryIds
	}
	if !IsNil(o.PersonIds) {
		toSerialize["personIds"] = o.PersonIds
	}
	if !IsNil(o.EmployeeIds) {
		toSerialize["employeeIds"] = o.EmployeeIds
	}
	if !IsNil(o.ProfileIds) {
		toSerialize["profileIds"] = o.ProfileIds
	}
	if !IsNil(o.OfficeIds) {
		toSerialize["officeIds"] = o.OfficeIds
	}
	if !IsNil(o.LanguageIds) {
		toSerialize["languageIds"] = o.LanguageIds
	}
	if !IsNil(o.Availability) {
		toSerialize["availability"] = o.Availability
	}
	if !IsNil(o.OnsiteRatio) {
		toSerialize["onsiteRatio"] = o.OnsiteRatio
	}
	if !IsNil(o.Seniority) {
		toSerialize["seniority"] = o.Seniority
	}
	if !IsNil(o.Skills) {
		toSerialize["skills"] = o.Skills
	}
	if !IsNil(o.ProjectIds) {
		toSerialize["projectIds"] = o.ProjectIds
	}
	if !IsNil(o.OrganizationIds) {
		toSerialize["organizationIds"] = o.OrganizationIds
	}
	if !IsNil(o.IndustryIds) {
		toSerialize["industryIds"] = o.IndustryIds
	}
	if !IsNil(o.CertificationIds) {
		toSerialize["certificationIds"] = o.CertificationIds
	}
	if !IsNil(o.AwardIds) {
		toSerialize["awardIds"] = o.AwardIds
	}
	if !IsNil(o.AwardedProject) {
		toSerialize["awardedProject"] = o.AwardedProject
	}
	if !IsNil(o.Inactive) {
		toSerialize["inactive"] = o.Inactive
	}
	return toSerialize, nil
}

type NullablePersonSearch struct {
	value *PersonSearch
	isSet bool
}

func (v NullablePersonSearch) Get() *PersonSearch {
	return v.value
}

func (v *NullablePersonSearch) Set(val *PersonSearch) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonSearch) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonSearch(val *PersonSearch) *NullablePersonSearch {
	return &NullablePersonSearch{value: val, isSet: true}
}

func (v NullablePersonSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


