# PersonSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryIds** | Pointer to **[]string** |  | [optional] 
**PersonIds** | Pointer to **[]string** |  | [optional] 
**EmployeeIds** | Pointer to **[]string** |  | [optional] 
**ProfileIds** | Pointer to **[]string** |  | [optional] 
**OfficeIds** | Pointer to **[]string** |  | [optional] 
**LanguageIds** | Pointer to **[]string** |  | [optional] 
**Availability** | Pointer to [**AvailabilityFilter**](AvailabilityFilter.md) |  | [optional] 
**OnsiteRatio** | Pointer to [**MinMaxPercent**](MinMaxPercent.md) |  | [optional] 
**Seniority** | Pointer to [**[]Seniority**](Seniority.md) |  | [optional] 
**Skills** | Pointer to [**[]PersonSkillFilter**](PersonSkillFilter.md) |  | [optional] 
**ProjectIds** | Pointer to **[]string** |  | [optional] 
**OrganizationIds** | Pointer to **[]string** |  | [optional] 
**IndustryIds** | Pointer to **[]string** |  | [optional] 
**CertificationIds** | Pointer to **[]string** |  | [optional] 
**AwardIds** | Pointer to **[]string** |  | [optional] 
**AwardedProject** | Pointer to [**AwardedProject**](AwardedProject.md) |  | [optional] 
**Inactive** | Pointer to **bool** |  | [optional] 

## Methods

### NewPersonSearch

`func NewPersonSearch() *PersonSearch`

NewPersonSearch instantiates a new PersonSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonSearchWithDefaults

`func NewPersonSearchWithDefaults() *PersonSearch`

NewPersonSearchWithDefaults instantiates a new PersonSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryIds

`func (o *PersonSearch) GetCountryIds() []string`

GetCountryIds returns the CountryIds field if non-nil, zero value otherwise.

### GetCountryIdsOk

`func (o *PersonSearch) GetCountryIdsOk() (*[]string, bool)`

GetCountryIdsOk returns a tuple with the CountryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIds

`func (o *PersonSearch) SetCountryIds(v []string)`

SetCountryIds sets CountryIds field to given value.

### HasCountryIds

`func (o *PersonSearch) HasCountryIds() bool`

HasCountryIds returns a boolean if a field has been set.

### GetPersonIds

`func (o *PersonSearch) GetPersonIds() []string`

GetPersonIds returns the PersonIds field if non-nil, zero value otherwise.

### GetPersonIdsOk

`func (o *PersonSearch) GetPersonIdsOk() (*[]string, bool)`

GetPersonIdsOk returns a tuple with the PersonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonIds

`func (o *PersonSearch) SetPersonIds(v []string)`

SetPersonIds sets PersonIds field to given value.

### HasPersonIds

`func (o *PersonSearch) HasPersonIds() bool`

HasPersonIds returns a boolean if a field has been set.

### GetEmployeeIds

`func (o *PersonSearch) GetEmployeeIds() []string`

GetEmployeeIds returns the EmployeeIds field if non-nil, zero value otherwise.

### GetEmployeeIdsOk

`func (o *PersonSearch) GetEmployeeIdsOk() (*[]string, bool)`

GetEmployeeIdsOk returns a tuple with the EmployeeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeIds

`func (o *PersonSearch) SetEmployeeIds(v []string)`

SetEmployeeIds sets EmployeeIds field to given value.

### HasEmployeeIds

`func (o *PersonSearch) HasEmployeeIds() bool`

HasEmployeeIds returns a boolean if a field has been set.

### GetProfileIds

`func (o *PersonSearch) GetProfileIds() []string`

GetProfileIds returns the ProfileIds field if non-nil, zero value otherwise.

### GetProfileIdsOk

`func (o *PersonSearch) GetProfileIdsOk() (*[]string, bool)`

GetProfileIdsOk returns a tuple with the ProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIds

`func (o *PersonSearch) SetProfileIds(v []string)`

SetProfileIds sets ProfileIds field to given value.

### HasProfileIds

`func (o *PersonSearch) HasProfileIds() bool`

HasProfileIds returns a boolean if a field has been set.

### GetOfficeIds

`func (o *PersonSearch) GetOfficeIds() []string`

GetOfficeIds returns the OfficeIds field if non-nil, zero value otherwise.

### GetOfficeIdsOk

`func (o *PersonSearch) GetOfficeIdsOk() (*[]string, bool)`

GetOfficeIdsOk returns a tuple with the OfficeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeIds

`func (o *PersonSearch) SetOfficeIds(v []string)`

SetOfficeIds sets OfficeIds field to given value.

### HasOfficeIds

`func (o *PersonSearch) HasOfficeIds() bool`

HasOfficeIds returns a boolean if a field has been set.

### GetLanguageIds

`func (o *PersonSearch) GetLanguageIds() []string`

GetLanguageIds returns the LanguageIds field if non-nil, zero value otherwise.

### GetLanguageIdsOk

`func (o *PersonSearch) GetLanguageIdsOk() (*[]string, bool)`

GetLanguageIdsOk returns a tuple with the LanguageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageIds

`func (o *PersonSearch) SetLanguageIds(v []string)`

SetLanguageIds sets LanguageIds field to given value.

### HasLanguageIds

`func (o *PersonSearch) HasLanguageIds() bool`

HasLanguageIds returns a boolean if a field has been set.

### GetAvailability

`func (o *PersonSearch) GetAvailability() AvailabilityFilter`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *PersonSearch) GetAvailabilityOk() (*AvailabilityFilter, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *PersonSearch) SetAvailability(v AvailabilityFilter)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *PersonSearch) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetOnsiteRatio

`func (o *PersonSearch) GetOnsiteRatio() MinMaxPercent`

GetOnsiteRatio returns the OnsiteRatio field if non-nil, zero value otherwise.

### GetOnsiteRatioOk

`func (o *PersonSearch) GetOnsiteRatioOk() (*MinMaxPercent, bool)`

GetOnsiteRatioOk returns a tuple with the OnsiteRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnsiteRatio

`func (o *PersonSearch) SetOnsiteRatio(v MinMaxPercent)`

SetOnsiteRatio sets OnsiteRatio field to given value.

### HasOnsiteRatio

`func (o *PersonSearch) HasOnsiteRatio() bool`

HasOnsiteRatio returns a boolean if a field has been set.

### GetSeniority

`func (o *PersonSearch) GetSeniority() []Seniority`

GetSeniority returns the Seniority field if non-nil, zero value otherwise.

### GetSeniorityOk

`func (o *PersonSearch) GetSeniorityOk() (*[]Seniority, bool)`

GetSeniorityOk returns a tuple with the Seniority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeniority

`func (o *PersonSearch) SetSeniority(v []Seniority)`

SetSeniority sets Seniority field to given value.

### HasSeniority

`func (o *PersonSearch) HasSeniority() bool`

HasSeniority returns a boolean if a field has been set.

### GetSkills

`func (o *PersonSearch) GetSkills() []PersonSkillFilter`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *PersonSearch) GetSkillsOk() (*[]PersonSkillFilter, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *PersonSearch) SetSkills(v []PersonSkillFilter)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *PersonSearch) HasSkills() bool`

HasSkills returns a boolean if a field has been set.

### GetProjectIds

`func (o *PersonSearch) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *PersonSearch) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *PersonSearch) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *PersonSearch) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### GetOrganizationIds

`func (o *PersonSearch) GetOrganizationIds() []string`

GetOrganizationIds returns the OrganizationIds field if non-nil, zero value otherwise.

### GetOrganizationIdsOk

`func (o *PersonSearch) GetOrganizationIdsOk() (*[]string, bool)`

GetOrganizationIdsOk returns a tuple with the OrganizationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationIds

`func (o *PersonSearch) SetOrganizationIds(v []string)`

SetOrganizationIds sets OrganizationIds field to given value.

### HasOrganizationIds

`func (o *PersonSearch) HasOrganizationIds() bool`

HasOrganizationIds returns a boolean if a field has been set.

### GetIndustryIds

`func (o *PersonSearch) GetIndustryIds() []string`

GetIndustryIds returns the IndustryIds field if non-nil, zero value otherwise.

### GetIndustryIdsOk

`func (o *PersonSearch) GetIndustryIdsOk() (*[]string, bool)`

GetIndustryIdsOk returns a tuple with the IndustryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryIds

`func (o *PersonSearch) SetIndustryIds(v []string)`

SetIndustryIds sets IndustryIds field to given value.

### HasIndustryIds

`func (o *PersonSearch) HasIndustryIds() bool`

HasIndustryIds returns a boolean if a field has been set.

### GetCertificationIds

`func (o *PersonSearch) GetCertificationIds() []string`

GetCertificationIds returns the CertificationIds field if non-nil, zero value otherwise.

### GetCertificationIdsOk

`func (o *PersonSearch) GetCertificationIdsOk() (*[]string, bool)`

GetCertificationIdsOk returns a tuple with the CertificationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificationIds

`func (o *PersonSearch) SetCertificationIds(v []string)`

SetCertificationIds sets CertificationIds field to given value.

### HasCertificationIds

`func (o *PersonSearch) HasCertificationIds() bool`

HasCertificationIds returns a boolean if a field has been set.

### GetAwardIds

`func (o *PersonSearch) GetAwardIds() []string`

GetAwardIds returns the AwardIds field if non-nil, zero value otherwise.

### GetAwardIdsOk

`func (o *PersonSearch) GetAwardIdsOk() (*[]string, bool)`

GetAwardIdsOk returns a tuple with the AwardIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwardIds

`func (o *PersonSearch) SetAwardIds(v []string)`

SetAwardIds sets AwardIds field to given value.

### HasAwardIds

`func (o *PersonSearch) HasAwardIds() bool`

HasAwardIds returns a boolean if a field has been set.

### GetAwardedProject

`func (o *PersonSearch) GetAwardedProject() AwardedProject`

GetAwardedProject returns the AwardedProject field if non-nil, zero value otherwise.

### GetAwardedProjectOk

`func (o *PersonSearch) GetAwardedProjectOk() (*AwardedProject, bool)`

GetAwardedProjectOk returns a tuple with the AwardedProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwardedProject

`func (o *PersonSearch) SetAwardedProject(v AwardedProject)`

SetAwardedProject sets AwardedProject field to given value.

### HasAwardedProject

`func (o *PersonSearch) HasAwardedProject() bool`

HasAwardedProject returns a boolean if a field has been set.

### GetInactive

`func (o *PersonSearch) GetInactive() bool`

GetInactive returns the Inactive field if non-nil, zero value otherwise.

### GetInactiveOk

`func (o *PersonSearch) GetInactiveOk() (*bool, bool)`

GetInactiveOk returns a tuple with the Inactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactive

`func (o *PersonSearch) SetInactive(v bool)`

SetInactive sets Inactive field to given value.

### HasInactive

`func (o *PersonSearch) HasInactive() bool`

HasInactive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


