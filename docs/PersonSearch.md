# PersonSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PersonIds** | Pointer to **[]string** |  | [optional] 
**ProfileIds** | Pointer to **[]string** |  | [optional] 
**OfficeIds** | Pointer to **[]string** |  | [optional] 
**LanguageIds** | Pointer to **[]string** |  | [optional] 
**Availability** | Pointer to [**AvailabilityFilter**](AvailabilityFilter.md) |  | [optional] 
**OnsiteRatio** | Pointer to [**MinMaxPercent**](MinMaxPercent.md) |  | [optional] 
**Seniority** | Pointer to [**[]Seniority**](Seniority.md) |  | [optional] 
**Skills** | Pointer to [**[]PersonSkillFilter**](PersonSkillFilter.md) |  | [optional] 
**Projects** | Pointer to [**[]PersonProjectFilter**](PersonProjectFilter.md) |  | [optional] 
**Organizations** | Pointer to [**[]PersonOrganizationFilter**](PersonOrganizationFilter.md) |  | [optional] 
**Industries** | Pointer to [**[]PersonIndustryFilter**](PersonIndustryFilter.md) |  | [optional] 
**Certifications** | Pointer to [**[]PersonCertificationFilter**](PersonCertificationFilter.md) |  | [optional] 

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

### GetProjects

`func (o *PersonSearch) GetProjects() []PersonProjectFilter`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *PersonSearch) GetProjectsOk() (*[]PersonProjectFilter, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *PersonSearch) SetProjects(v []PersonProjectFilter)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *PersonSearch) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetOrganizations

`func (o *PersonSearch) GetOrganizations() []PersonOrganizationFilter`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *PersonSearch) GetOrganizationsOk() (*[]PersonOrganizationFilter, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *PersonSearch) SetOrganizations(v []PersonOrganizationFilter)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *PersonSearch) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetIndustries

`func (o *PersonSearch) GetIndustries() []PersonIndustryFilter`

GetIndustries returns the Industries field if non-nil, zero value otherwise.

### GetIndustriesOk

`func (o *PersonSearch) GetIndustriesOk() (*[]PersonIndustryFilter, bool)`

GetIndustriesOk returns a tuple with the Industries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustries

`func (o *PersonSearch) SetIndustries(v []PersonIndustryFilter)`

SetIndustries sets Industries field to given value.

### HasIndustries

`func (o *PersonSearch) HasIndustries() bool`

HasIndustries returns a boolean if a field has been set.

### GetCertifications

`func (o *PersonSearch) GetCertifications() []PersonCertificationFilter`

GetCertifications returns the Certifications field if non-nil, zero value otherwise.

### GetCertificationsOk

`func (o *PersonSearch) GetCertificationsOk() (*[]PersonCertificationFilter, bool)`

GetCertificationsOk returns a tuple with the Certifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifications

`func (o *PersonSearch) SetCertifications(v []PersonCertificationFilter)`

SetCertifications sets Certifications field to given value.

### HasCertifications

`func (o *PersonSearch) HasCertifications() bool`

HasCertifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


