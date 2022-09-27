# PersonScoreDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Person** | Pointer to [**Person**](Person.md) |  | [optional] 
**Projects** | Pointer to [**[]ProjectParticipation**](ProjectParticipation.md) |  | [optional] 
**Industries** | Pointer to [**[]Industry**](Industry.md) |  | [optional] 
**Experiences** | Pointer to [**[]Experience**](Experience.md) |  | [optional] 
**Interests** | Pointer to [**[]Skill**](Skill.md) |  | [optional] 
**Certifications** | Pointer to [**[]CertificationDetails**](CertificationDetails.md) |  | [optional] 
**Languages** | Pointer to [**[]LanguageLevel**](LanguageLevel.md) |  | [optional] 
**Office** | Pointer to [**Office**](Office.md) |  | [optional] 
**Availabilities** | Pointer to [**[]AvailabilityDetail**](AvailabilityDetail.md) |  | [optional] 
**SkillGroups** | Pointer to [**[]SkillGroup**](SkillGroup.md) |  | [optional] 
**Score** | Pointer to **float32** |  | [optional] 

## Methods

### NewPersonScoreDetail

`func NewPersonScoreDetail() *PersonScoreDetail`

NewPersonScoreDetail instantiates a new PersonScoreDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonScoreDetailWithDefaults

`func NewPersonScoreDetailWithDefaults() *PersonScoreDetail`

NewPersonScoreDetailWithDefaults instantiates a new PersonScoreDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPerson

`func (o *PersonScoreDetail) GetPerson() Person`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *PersonScoreDetail) GetPersonOk() (*Person, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *PersonScoreDetail) SetPerson(v Person)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *PersonScoreDetail) HasPerson() bool`

HasPerson returns a boolean if a field has been set.

### GetProjects

`func (o *PersonScoreDetail) GetProjects() []ProjectParticipation`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *PersonScoreDetail) GetProjectsOk() (*[]ProjectParticipation, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *PersonScoreDetail) SetProjects(v []ProjectParticipation)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *PersonScoreDetail) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetIndustries

`func (o *PersonScoreDetail) GetIndustries() []Industry`

GetIndustries returns the Industries field if non-nil, zero value otherwise.

### GetIndustriesOk

`func (o *PersonScoreDetail) GetIndustriesOk() (*[]Industry, bool)`

GetIndustriesOk returns a tuple with the Industries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustries

`func (o *PersonScoreDetail) SetIndustries(v []Industry)`

SetIndustries sets Industries field to given value.

### HasIndustries

`func (o *PersonScoreDetail) HasIndustries() bool`

HasIndustries returns a boolean if a field has been set.

### GetExperiences

`func (o *PersonScoreDetail) GetExperiences() []Experience`

GetExperiences returns the Experiences field if non-nil, zero value otherwise.

### GetExperiencesOk

`func (o *PersonScoreDetail) GetExperiencesOk() (*[]Experience, bool)`

GetExperiencesOk returns a tuple with the Experiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiences

`func (o *PersonScoreDetail) SetExperiences(v []Experience)`

SetExperiences sets Experiences field to given value.

### HasExperiences

`func (o *PersonScoreDetail) HasExperiences() bool`

HasExperiences returns a boolean if a field has been set.

### GetInterests

`func (o *PersonScoreDetail) GetInterests() []Skill`

GetInterests returns the Interests field if non-nil, zero value otherwise.

### GetInterestsOk

`func (o *PersonScoreDetail) GetInterestsOk() (*[]Skill, bool)`

GetInterestsOk returns a tuple with the Interests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterests

`func (o *PersonScoreDetail) SetInterests(v []Skill)`

SetInterests sets Interests field to given value.

### HasInterests

`func (o *PersonScoreDetail) HasInterests() bool`

HasInterests returns a boolean if a field has been set.

### GetCertifications

`func (o *PersonScoreDetail) GetCertifications() []CertificationDetails`

GetCertifications returns the Certifications field if non-nil, zero value otherwise.

### GetCertificationsOk

`func (o *PersonScoreDetail) GetCertificationsOk() (*[]CertificationDetails, bool)`

GetCertificationsOk returns a tuple with the Certifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifications

`func (o *PersonScoreDetail) SetCertifications(v []CertificationDetails)`

SetCertifications sets Certifications field to given value.

### HasCertifications

`func (o *PersonScoreDetail) HasCertifications() bool`

HasCertifications returns a boolean if a field has been set.

### GetLanguages

`func (o *PersonScoreDetail) GetLanguages() []LanguageLevel`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *PersonScoreDetail) GetLanguagesOk() (*[]LanguageLevel, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *PersonScoreDetail) SetLanguages(v []LanguageLevel)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *PersonScoreDetail) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetOffice

`func (o *PersonScoreDetail) GetOffice() Office`

GetOffice returns the Office field if non-nil, zero value otherwise.

### GetOfficeOk

`func (o *PersonScoreDetail) GetOfficeOk() (*Office, bool)`

GetOfficeOk returns a tuple with the Office field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice

`func (o *PersonScoreDetail) SetOffice(v Office)`

SetOffice sets Office field to given value.

### HasOffice

`func (o *PersonScoreDetail) HasOffice() bool`

HasOffice returns a boolean if a field has been set.

### GetAvailabilities

`func (o *PersonScoreDetail) GetAvailabilities() []AvailabilityDetail`

GetAvailabilities returns the Availabilities field if non-nil, zero value otherwise.

### GetAvailabilitiesOk

`func (o *PersonScoreDetail) GetAvailabilitiesOk() (*[]AvailabilityDetail, bool)`

GetAvailabilitiesOk returns a tuple with the Availabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilities

`func (o *PersonScoreDetail) SetAvailabilities(v []AvailabilityDetail)`

SetAvailabilities sets Availabilities field to given value.

### HasAvailabilities

`func (o *PersonScoreDetail) HasAvailabilities() bool`

HasAvailabilities returns a boolean if a field has been set.

### GetSkillGroups

`func (o *PersonScoreDetail) GetSkillGroups() []SkillGroup`

GetSkillGroups returns the SkillGroups field if non-nil, zero value otherwise.

### GetSkillGroupsOk

`func (o *PersonScoreDetail) GetSkillGroupsOk() (*[]SkillGroup, bool)`

GetSkillGroupsOk returns a tuple with the SkillGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillGroups

`func (o *PersonScoreDetail) SetSkillGroups(v []SkillGroup)`

SetSkillGroups sets SkillGroups field to given value.

### HasSkillGroups

`func (o *PersonScoreDetail) HasSkillGroups() bool`

HasSkillGroups returns a boolean if a field has been set.

### GetScore

`func (o *PersonScoreDetail) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *PersonScoreDetail) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *PersonScoreDetail) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *PersonScoreDetail) HasScore() bool`

HasScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


