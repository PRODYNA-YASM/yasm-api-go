/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PersonDetails struct for PersonDetails
type PersonDetails struct {
	Person *Person `json:"person,omitempty"`
	Industries []Industry `json:"industries,omitempty"`
	Experiences []Experience `json:"experiences,omitempty"`
	Interests []Skill `json:"interests,omitempty"`
	Certifications []CertificationDetails `json:"certifications,omitempty"`
	Languages []LanguageLevel `json:"languages,omitempty"`
	Office *Office `json:"office,omitempty"`
	Availabilities []AvailabilityDetail `json:"availabilities,omitempty"`
	SkillGroups []ExperienceSkillGroup `json:"skillGroups,omitempty"`
	Profiles []Profile `json:"profiles,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
}

// NewPersonDetails instantiates a new PersonDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonDetails() *PersonDetails {
	this := PersonDetails{}
	return &this
}

// NewPersonDetailsWithDefaults instantiates a new PersonDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonDetailsWithDefaults() *PersonDetails {
	this := PersonDetails{}
	return &this
}

// GetPerson returns the Person field value if set, zero value otherwise.
func (o *PersonDetails) GetPerson() Person {
	if o == nil || o.Person == nil {
		var ret Person
		return ret
	}
	return *o.Person
}

// GetPersonOk returns a tuple with the Person field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetPersonOk() (*Person, bool) {
	if o == nil || o.Person == nil {
		return nil, false
	}
	return o.Person, true
}

// HasPerson returns a boolean if a field has been set.
func (o *PersonDetails) HasPerson() bool {
	if o != nil && o.Person != nil {
		return true
	}

	return false
}

// SetPerson gets a reference to the given Person and assigns it to the Person field.
func (o *PersonDetails) SetPerson(v Person) {
	o.Person = &v
}

// GetIndustries returns the Industries field value if set, zero value otherwise.
func (o *PersonDetails) GetIndustries() []Industry {
	if o == nil || o.Industries == nil {
		var ret []Industry
		return ret
	}
	return o.Industries
}

// GetIndustriesOk returns a tuple with the Industries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetIndustriesOk() ([]Industry, bool) {
	if o == nil || o.Industries == nil {
		return nil, false
	}
	return o.Industries, true
}

// HasIndustries returns a boolean if a field has been set.
func (o *PersonDetails) HasIndustries() bool {
	if o != nil && o.Industries != nil {
		return true
	}

	return false
}

// SetIndustries gets a reference to the given []Industry and assigns it to the Industries field.
func (o *PersonDetails) SetIndustries(v []Industry) {
	o.Industries = v
}

// GetExperiences returns the Experiences field value if set, zero value otherwise.
func (o *PersonDetails) GetExperiences() []Experience {
	if o == nil || o.Experiences == nil {
		var ret []Experience
		return ret
	}
	return o.Experiences
}

// GetExperiencesOk returns a tuple with the Experiences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetExperiencesOk() ([]Experience, bool) {
	if o == nil || o.Experiences == nil {
		return nil, false
	}
	return o.Experiences, true
}

// HasExperiences returns a boolean if a field has been set.
func (o *PersonDetails) HasExperiences() bool {
	if o != nil && o.Experiences != nil {
		return true
	}

	return false
}

// SetExperiences gets a reference to the given []Experience and assigns it to the Experiences field.
func (o *PersonDetails) SetExperiences(v []Experience) {
	o.Experiences = v
}

// GetInterests returns the Interests field value if set, zero value otherwise.
func (o *PersonDetails) GetInterests() []Skill {
	if o == nil || o.Interests == nil {
		var ret []Skill
		return ret
	}
	return o.Interests
}

// GetInterestsOk returns a tuple with the Interests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetInterestsOk() ([]Skill, bool) {
	if o == nil || o.Interests == nil {
		return nil, false
	}
	return o.Interests, true
}

// HasInterests returns a boolean if a field has been set.
func (o *PersonDetails) HasInterests() bool {
	if o != nil && o.Interests != nil {
		return true
	}

	return false
}

// SetInterests gets a reference to the given []Skill and assigns it to the Interests field.
func (o *PersonDetails) SetInterests(v []Skill) {
	o.Interests = v
}

// GetCertifications returns the Certifications field value if set, zero value otherwise.
func (o *PersonDetails) GetCertifications() []CertificationDetails {
	if o == nil || o.Certifications == nil {
		var ret []CertificationDetails
		return ret
	}
	return o.Certifications
}

// GetCertificationsOk returns a tuple with the Certifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetCertificationsOk() ([]CertificationDetails, bool) {
	if o == nil || o.Certifications == nil {
		return nil, false
	}
	return o.Certifications, true
}

// HasCertifications returns a boolean if a field has been set.
func (o *PersonDetails) HasCertifications() bool {
	if o != nil && o.Certifications != nil {
		return true
	}

	return false
}

// SetCertifications gets a reference to the given []CertificationDetails and assigns it to the Certifications field.
func (o *PersonDetails) SetCertifications(v []CertificationDetails) {
	o.Certifications = v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *PersonDetails) GetLanguages() []LanguageLevel {
	if o == nil || o.Languages == nil {
		var ret []LanguageLevel
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetLanguagesOk() ([]LanguageLevel, bool) {
	if o == nil || o.Languages == nil {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *PersonDetails) HasLanguages() bool {
	if o != nil && o.Languages != nil {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []LanguageLevel and assigns it to the Languages field.
func (o *PersonDetails) SetLanguages(v []LanguageLevel) {
	o.Languages = v
}

// GetOffice returns the Office field value if set, zero value otherwise.
func (o *PersonDetails) GetOffice() Office {
	if o == nil || o.Office == nil {
		var ret Office
		return ret
	}
	return *o.Office
}

// GetOfficeOk returns a tuple with the Office field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetOfficeOk() (*Office, bool) {
	if o == nil || o.Office == nil {
		return nil, false
	}
	return o.Office, true
}

// HasOffice returns a boolean if a field has been set.
func (o *PersonDetails) HasOffice() bool {
	if o != nil && o.Office != nil {
		return true
	}

	return false
}

// SetOffice gets a reference to the given Office and assigns it to the Office field.
func (o *PersonDetails) SetOffice(v Office) {
	o.Office = &v
}

// GetAvailabilities returns the Availabilities field value if set, zero value otherwise.
func (o *PersonDetails) GetAvailabilities() []AvailabilityDetail {
	if o == nil || o.Availabilities == nil {
		var ret []AvailabilityDetail
		return ret
	}
	return o.Availabilities
}

// GetAvailabilitiesOk returns a tuple with the Availabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetAvailabilitiesOk() ([]AvailabilityDetail, bool) {
	if o == nil || o.Availabilities == nil {
		return nil, false
	}
	return o.Availabilities, true
}

// HasAvailabilities returns a boolean if a field has been set.
func (o *PersonDetails) HasAvailabilities() bool {
	if o != nil && o.Availabilities != nil {
		return true
	}

	return false
}

// SetAvailabilities gets a reference to the given []AvailabilityDetail and assigns it to the Availabilities field.
func (o *PersonDetails) SetAvailabilities(v []AvailabilityDetail) {
	o.Availabilities = v
}

// GetSkillGroups returns the SkillGroups field value if set, zero value otherwise.
func (o *PersonDetails) GetSkillGroups() []ExperienceSkillGroup {
	if o == nil || o.SkillGroups == nil {
		var ret []ExperienceSkillGroup
		return ret
	}
	return o.SkillGroups
}

// GetSkillGroupsOk returns a tuple with the SkillGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetSkillGroupsOk() ([]ExperienceSkillGroup, bool) {
	if o == nil || o.SkillGroups == nil {
		return nil, false
	}
	return o.SkillGroups, true
}

// HasSkillGroups returns a boolean if a field has been set.
func (o *PersonDetails) HasSkillGroups() bool {
	if o != nil && o.SkillGroups != nil {
		return true
	}

	return false
}

// SetSkillGroups gets a reference to the given []ExperienceSkillGroup and assigns it to the SkillGroups field.
func (o *PersonDetails) SetSkillGroups(v []ExperienceSkillGroup) {
	o.SkillGroups = v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *PersonDetails) GetProfiles() []Profile {
	if o == nil || o.Profiles == nil {
		var ret []Profile
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetProfilesOk() ([]Profile, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *PersonDetails) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []Profile and assigns it to the Profiles field.
func (o *PersonDetails) SetProfiles(v []Profile) {
	o.Profiles = v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *PersonDetails) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *PersonDetails) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *PersonDetails) SetObjectType(v string) {
	o.ObjectType = &v
}

func (o PersonDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Person != nil {
		toSerialize["person"] = o.Person
	}
	if o.Industries != nil {
		toSerialize["industries"] = o.Industries
	}
	if o.Experiences != nil {
		toSerialize["experiences"] = o.Experiences
	}
	if o.Interests != nil {
		toSerialize["interests"] = o.Interests
	}
	if o.Certifications != nil {
		toSerialize["certifications"] = o.Certifications
	}
	if o.Languages != nil {
		toSerialize["languages"] = o.Languages
	}
	if o.Office != nil {
		toSerialize["office"] = o.Office
	}
	if o.Availabilities != nil {
		toSerialize["availabilities"] = o.Availabilities
	}
	if o.SkillGroups != nil {
		toSerialize["skillGroups"] = o.SkillGroups
	}
	if o.Profiles != nil {
		toSerialize["profiles"] = o.Profiles
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	return json.Marshal(toSerialize)
}

type NullablePersonDetails struct {
	value *PersonDetails
	isSet bool
}

func (v NullablePersonDetails) Get() *PersonDetails {
	return v.value
}

func (v *NullablePersonDetails) Set(val *PersonDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonDetails(val *PersonDetails) *NullablePersonDetails {
	return &NullablePersonDetails{value: val, isSet: true}
}

func (v NullablePersonDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


