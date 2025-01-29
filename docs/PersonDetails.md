# PersonDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**Person** | Pointer to [**Person**](Person.md) |  | [optional] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 
**Industries** | Pointer to [**[]Industry**](Industry.md) |  | [optional] 
**Experiences** | Pointer to [**[]Experience**](Experience.md) |  | [optional] 
**Interests** | Pointer to [**[]Skill**](Skill.md) |  | [optional] 
**Certifications** | Pointer to [**[]CertificationDetails**](CertificationDetails.md) |  | [optional] 
**Awards** | Pointer to [**[]AwardDetails**](AwardDetails.md) |  | [optional] 
**Languages** | Pointer to [**[]LanguageLevel**](LanguageLevel.md) |  | [optional] 
**Nationalities** | Pointer to [**[]Country**](Country.md) |  | [optional] 
**Office** | Pointer to [**OfficeDetails**](OfficeDetails.md) |  | [optional] 
**Availabilities** | Pointer to [**[]AvailabilityDetail**](AvailabilityDetail.md) |  | [optional] 
**SkillGroups** | Pointer to [**[]ExperienceSkillGroup**](ExperienceSkillGroup.md) |  | [optional] 
**Profiles** | Pointer to [**[]Profile**](Profile.md) |  | [optional] 
**Manager** | Pointer to [**Person**](Person.md) |  | [optional] 
**TeamMembers** | Pointer to [**[]Person**](Person.md) |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 

## Methods

### NewPersonDetails

`func NewPersonDetails() *PersonDetails`

NewPersonDetails instantiates a new PersonDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonDetailsWithDefaults

`func NewPersonDetailsWithDefaults() *PersonDetails`

NewPersonDetailsWithDefaults instantiates a new PersonDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *PersonDetails) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *PersonDetails) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *PersonDetails) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *PersonDetails) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetPerson

`func (o *PersonDetails) GetPerson() Person`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *PersonDetails) GetPersonOk() (*Person, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *PersonDetails) SetPerson(v Person)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *PersonDetails) HasPerson() bool`

HasPerson returns a boolean if a field has been set.

### GetOrganization

`func (o *PersonDetails) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PersonDetails) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PersonDetails) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *PersonDetails) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetIndustries

`func (o *PersonDetails) GetIndustries() []Industry`

GetIndustries returns the Industries field if non-nil, zero value otherwise.

### GetIndustriesOk

`func (o *PersonDetails) GetIndustriesOk() (*[]Industry, bool)`

GetIndustriesOk returns a tuple with the Industries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustries

`func (o *PersonDetails) SetIndustries(v []Industry)`

SetIndustries sets Industries field to given value.

### HasIndustries

`func (o *PersonDetails) HasIndustries() bool`

HasIndustries returns a boolean if a field has been set.

### GetExperiences

`func (o *PersonDetails) GetExperiences() []Experience`

GetExperiences returns the Experiences field if non-nil, zero value otherwise.

### GetExperiencesOk

`func (o *PersonDetails) GetExperiencesOk() (*[]Experience, bool)`

GetExperiencesOk returns a tuple with the Experiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiences

`func (o *PersonDetails) SetExperiences(v []Experience)`

SetExperiences sets Experiences field to given value.

### HasExperiences

`func (o *PersonDetails) HasExperiences() bool`

HasExperiences returns a boolean if a field has been set.

### GetInterests

`func (o *PersonDetails) GetInterests() []Skill`

GetInterests returns the Interests field if non-nil, zero value otherwise.

### GetInterestsOk

`func (o *PersonDetails) GetInterestsOk() (*[]Skill, bool)`

GetInterestsOk returns a tuple with the Interests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterests

`func (o *PersonDetails) SetInterests(v []Skill)`

SetInterests sets Interests field to given value.

### HasInterests

`func (o *PersonDetails) HasInterests() bool`

HasInterests returns a boolean if a field has been set.

### GetCertifications

`func (o *PersonDetails) GetCertifications() []CertificationDetails`

GetCertifications returns the Certifications field if non-nil, zero value otherwise.

### GetCertificationsOk

`func (o *PersonDetails) GetCertificationsOk() (*[]CertificationDetails, bool)`

GetCertificationsOk returns a tuple with the Certifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifications

`func (o *PersonDetails) SetCertifications(v []CertificationDetails)`

SetCertifications sets Certifications field to given value.

### HasCertifications

`func (o *PersonDetails) HasCertifications() bool`

HasCertifications returns a boolean if a field has been set.

### GetAwards

`func (o *PersonDetails) GetAwards() []AwardDetails`

GetAwards returns the Awards field if non-nil, zero value otherwise.

### GetAwardsOk

`func (o *PersonDetails) GetAwardsOk() (*[]AwardDetails, bool)`

GetAwardsOk returns a tuple with the Awards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwards

`func (o *PersonDetails) SetAwards(v []AwardDetails)`

SetAwards sets Awards field to given value.

### HasAwards

`func (o *PersonDetails) HasAwards() bool`

HasAwards returns a boolean if a field has been set.

### GetLanguages

`func (o *PersonDetails) GetLanguages() []LanguageLevel`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *PersonDetails) GetLanguagesOk() (*[]LanguageLevel, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *PersonDetails) SetLanguages(v []LanguageLevel)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *PersonDetails) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetNationalities

`func (o *PersonDetails) GetNationalities() []Country`

GetNationalities returns the Nationalities field if non-nil, zero value otherwise.

### GetNationalitiesOk

`func (o *PersonDetails) GetNationalitiesOk() (*[]Country, bool)`

GetNationalitiesOk returns a tuple with the Nationalities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalities

`func (o *PersonDetails) SetNationalities(v []Country)`

SetNationalities sets Nationalities field to given value.

### HasNationalities

`func (o *PersonDetails) HasNationalities() bool`

HasNationalities returns a boolean if a field has been set.

### GetOffice

`func (o *PersonDetails) GetOffice() OfficeDetails`

GetOffice returns the Office field if non-nil, zero value otherwise.

### GetOfficeOk

`func (o *PersonDetails) GetOfficeOk() (*OfficeDetails, bool)`

GetOfficeOk returns a tuple with the Office field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice

`func (o *PersonDetails) SetOffice(v OfficeDetails)`

SetOffice sets Office field to given value.

### HasOffice

`func (o *PersonDetails) HasOffice() bool`

HasOffice returns a boolean if a field has been set.

### GetAvailabilities

`func (o *PersonDetails) GetAvailabilities() []AvailabilityDetail`

GetAvailabilities returns the Availabilities field if non-nil, zero value otherwise.

### GetAvailabilitiesOk

`func (o *PersonDetails) GetAvailabilitiesOk() (*[]AvailabilityDetail, bool)`

GetAvailabilitiesOk returns a tuple with the Availabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilities

`func (o *PersonDetails) SetAvailabilities(v []AvailabilityDetail)`

SetAvailabilities sets Availabilities field to given value.

### HasAvailabilities

`func (o *PersonDetails) HasAvailabilities() bool`

HasAvailabilities returns a boolean if a field has been set.

### GetSkillGroups

`func (o *PersonDetails) GetSkillGroups() []ExperienceSkillGroup`

GetSkillGroups returns the SkillGroups field if non-nil, zero value otherwise.

### GetSkillGroupsOk

`func (o *PersonDetails) GetSkillGroupsOk() (*[]ExperienceSkillGroup, bool)`

GetSkillGroupsOk returns a tuple with the SkillGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillGroups

`func (o *PersonDetails) SetSkillGroups(v []ExperienceSkillGroup)`

SetSkillGroups sets SkillGroups field to given value.

### HasSkillGroups

`func (o *PersonDetails) HasSkillGroups() bool`

HasSkillGroups returns a boolean if a field has been set.

### GetProfiles

`func (o *PersonDetails) GetProfiles() []Profile`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *PersonDetails) GetProfilesOk() (*[]Profile, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *PersonDetails) SetProfiles(v []Profile)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *PersonDetails) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetManager

`func (o *PersonDetails) GetManager() Person`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *PersonDetails) GetManagerOk() (*Person, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *PersonDetails) SetManager(v Person)`

SetManager sets Manager field to given value.

### HasManager

`func (o *PersonDetails) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetTeamMembers

`func (o *PersonDetails) GetTeamMembers() []Person`

GetTeamMembers returns the TeamMembers field if non-nil, zero value otherwise.

### GetTeamMembersOk

`func (o *PersonDetails) GetTeamMembersOk() (*[]Person, bool)`

GetTeamMembersOk returns a tuple with the TeamMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamMembers

`func (o *PersonDetails) SetTeamMembers(v []Person)`

SetTeamMembers sets TeamMembers field to given value.

### HasTeamMembers

`func (o *PersonDetails) HasTeamMembers() bool`

HasTeamMembers returns a boolean if a field has been set.

### GetObjectType

`func (o *PersonDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PersonDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PersonDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *PersonDetails) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


