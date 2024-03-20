# PersonSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**IndustriesIds** | Pointer to **[]string** |  | [optional] 
**CertificationsIds** | Pointer to **[]string** |  | [optional] 

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

### GetIndustriesIds

`func (o *PersonSearch) GetIndustriesIds() []string`

GetIndustriesIds returns the IndustriesIds field if non-nil, zero value otherwise.

### GetIndustriesIdsOk

`func (o *PersonSearch) GetIndustriesIdsOk() (*[]string, bool)`

GetIndustriesIdsOk returns a tuple with the IndustriesIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustriesIds

`func (o *PersonSearch) SetIndustriesIds(v []string)`

SetIndustriesIds sets IndustriesIds field to given value.

### HasIndustriesIds

`func (o *PersonSearch) HasIndustriesIds() bool`

HasIndustriesIds returns a boolean if a field has been set.

### GetCertificationsIds

`func (o *PersonSearch) GetCertificationsIds() []string`

GetCertificationsIds returns the CertificationsIds field if non-nil, zero value otherwise.

### GetCertificationsIdsOk

`func (o *PersonSearch) GetCertificationsIdsOk() (*[]string, bool)`

GetCertificationsIdsOk returns a tuple with the CertificationsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificationsIds

`func (o *PersonSearch) SetCertificationsIds(v []string)`

SetCertificationsIds sets CertificationsIds field to given value.

### HasCertificationsIds

`func (o *PersonSearch) HasCertificationsIds() bool`

HasCertificationsIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


