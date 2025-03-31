# ProjectSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectIds** | Pointer to **[]string** |  | [optional] 
**MinStartDate** | Pointer to **string** |  | [optional] 
**MaxEndDate** | Pointer to **string** |  | [optional] 
**ProjectStatuses** | Pointer to [**[]ProjectStatus**](ProjectStatus.md) |  | [optional] 
**ParticipationAmountInMonths** | Pointer to [**MinMax**](MinMax.md) |  | [optional] 
**InvolvedCountryIds** | Pointer to **[]string** |  | [optional] 
**OrganizationCountryIds** | Pointer to **[]string** |  | [optional] 
**AmountOfInvolvedPersons** | Pointer to [**MinMax**](MinMax.md) |  | [optional] 
**SkillIds** | Pointer to **[]string** |  | [optional] 
**ProjectType** | Pointer to [**[]ProjectType**](ProjectType.md) |  | [optional] 
**Confidentiality** | Pointer to [**[]Confidentiality**](Confidentiality.md) |  | [optional] 
**ParticipantIds** | Pointer to **[]string** |  | [optional] 
**IndustryIds** | Pointer to **[]string** |  | [optional] 
**OrganizationIds** | Pointer to **[]string** |  | [optional] 
**AwardedParticipant** | Pointer to [**AwardedParticipant**](AwardedParticipant.md) |  | [optional] 
**External** | Pointer to **bool** |  | [optional] 

## Methods

### NewProjectSearch

`func NewProjectSearch() *ProjectSearch`

NewProjectSearch instantiates a new ProjectSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectSearchWithDefaults

`func NewProjectSearchWithDefaults() *ProjectSearch`

NewProjectSearchWithDefaults instantiates a new ProjectSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectIds

`func (o *ProjectSearch) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *ProjectSearch) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *ProjectSearch) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *ProjectSearch) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### GetMinStartDate

`func (o *ProjectSearch) GetMinStartDate() string`

GetMinStartDate returns the MinStartDate field if non-nil, zero value otherwise.

### GetMinStartDateOk

`func (o *ProjectSearch) GetMinStartDateOk() (*string, bool)`

GetMinStartDateOk returns a tuple with the MinStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinStartDate

`func (o *ProjectSearch) SetMinStartDate(v string)`

SetMinStartDate sets MinStartDate field to given value.

### HasMinStartDate

`func (o *ProjectSearch) HasMinStartDate() bool`

HasMinStartDate returns a boolean if a field has been set.

### GetMaxEndDate

`func (o *ProjectSearch) GetMaxEndDate() string`

GetMaxEndDate returns the MaxEndDate field if non-nil, zero value otherwise.

### GetMaxEndDateOk

`func (o *ProjectSearch) GetMaxEndDateOk() (*string, bool)`

GetMaxEndDateOk returns a tuple with the MaxEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEndDate

`func (o *ProjectSearch) SetMaxEndDate(v string)`

SetMaxEndDate sets MaxEndDate field to given value.

### HasMaxEndDate

`func (o *ProjectSearch) HasMaxEndDate() bool`

HasMaxEndDate returns a boolean if a field has been set.

### GetProjectStatuses

`func (o *ProjectSearch) GetProjectStatuses() []ProjectStatus`

GetProjectStatuses returns the ProjectStatuses field if non-nil, zero value otherwise.

### GetProjectStatusesOk

`func (o *ProjectSearch) GetProjectStatusesOk() (*[]ProjectStatus, bool)`

GetProjectStatusesOk returns a tuple with the ProjectStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectStatuses

`func (o *ProjectSearch) SetProjectStatuses(v []ProjectStatus)`

SetProjectStatuses sets ProjectStatuses field to given value.

### HasProjectStatuses

`func (o *ProjectSearch) HasProjectStatuses() bool`

HasProjectStatuses returns a boolean if a field has been set.

### GetParticipationAmountInMonths

`func (o *ProjectSearch) GetParticipationAmountInMonths() MinMax`

GetParticipationAmountInMonths returns the ParticipationAmountInMonths field if non-nil, zero value otherwise.

### GetParticipationAmountInMonthsOk

`func (o *ProjectSearch) GetParticipationAmountInMonthsOk() (*MinMax, bool)`

GetParticipationAmountInMonthsOk returns a tuple with the ParticipationAmountInMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipationAmountInMonths

`func (o *ProjectSearch) SetParticipationAmountInMonths(v MinMax)`

SetParticipationAmountInMonths sets ParticipationAmountInMonths field to given value.

### HasParticipationAmountInMonths

`func (o *ProjectSearch) HasParticipationAmountInMonths() bool`

HasParticipationAmountInMonths returns a boolean if a field has been set.

### GetInvolvedCountryIds

`func (o *ProjectSearch) GetInvolvedCountryIds() []string`

GetInvolvedCountryIds returns the InvolvedCountryIds field if non-nil, zero value otherwise.

### GetInvolvedCountryIdsOk

`func (o *ProjectSearch) GetInvolvedCountryIdsOk() (*[]string, bool)`

GetInvolvedCountryIdsOk returns a tuple with the InvolvedCountryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvolvedCountryIds

`func (o *ProjectSearch) SetInvolvedCountryIds(v []string)`

SetInvolvedCountryIds sets InvolvedCountryIds field to given value.

### HasInvolvedCountryIds

`func (o *ProjectSearch) HasInvolvedCountryIds() bool`

HasInvolvedCountryIds returns a boolean if a field has been set.

### GetOrganizationCountryIds

`func (o *ProjectSearch) GetOrganizationCountryIds() []string`

GetOrganizationCountryIds returns the OrganizationCountryIds field if non-nil, zero value otherwise.

### GetOrganizationCountryIdsOk

`func (o *ProjectSearch) GetOrganizationCountryIdsOk() (*[]string, bool)`

GetOrganizationCountryIdsOk returns a tuple with the OrganizationCountryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationCountryIds

`func (o *ProjectSearch) SetOrganizationCountryIds(v []string)`

SetOrganizationCountryIds sets OrganizationCountryIds field to given value.

### HasOrganizationCountryIds

`func (o *ProjectSearch) HasOrganizationCountryIds() bool`

HasOrganizationCountryIds returns a boolean if a field has been set.

### GetAmountOfInvolvedPersons

`func (o *ProjectSearch) GetAmountOfInvolvedPersons() MinMax`

GetAmountOfInvolvedPersons returns the AmountOfInvolvedPersons field if non-nil, zero value otherwise.

### GetAmountOfInvolvedPersonsOk

`func (o *ProjectSearch) GetAmountOfInvolvedPersonsOk() (*MinMax, bool)`

GetAmountOfInvolvedPersonsOk returns a tuple with the AmountOfInvolvedPersons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOfInvolvedPersons

`func (o *ProjectSearch) SetAmountOfInvolvedPersons(v MinMax)`

SetAmountOfInvolvedPersons sets AmountOfInvolvedPersons field to given value.

### HasAmountOfInvolvedPersons

`func (o *ProjectSearch) HasAmountOfInvolvedPersons() bool`

HasAmountOfInvolvedPersons returns a boolean if a field has been set.

### GetSkillIds

`func (o *ProjectSearch) GetSkillIds() []string`

GetSkillIds returns the SkillIds field if non-nil, zero value otherwise.

### GetSkillIdsOk

`func (o *ProjectSearch) GetSkillIdsOk() (*[]string, bool)`

GetSkillIdsOk returns a tuple with the SkillIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillIds

`func (o *ProjectSearch) SetSkillIds(v []string)`

SetSkillIds sets SkillIds field to given value.

### HasSkillIds

`func (o *ProjectSearch) HasSkillIds() bool`

HasSkillIds returns a boolean if a field has been set.

### GetProjectType

`func (o *ProjectSearch) GetProjectType() []ProjectType`

GetProjectType returns the ProjectType field if non-nil, zero value otherwise.

### GetProjectTypeOk

`func (o *ProjectSearch) GetProjectTypeOk() (*[]ProjectType, bool)`

GetProjectTypeOk returns a tuple with the ProjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectType

`func (o *ProjectSearch) SetProjectType(v []ProjectType)`

SetProjectType sets ProjectType field to given value.

### HasProjectType

`func (o *ProjectSearch) HasProjectType() bool`

HasProjectType returns a boolean if a field has been set.

### GetConfidentiality

`func (o *ProjectSearch) GetConfidentiality() []Confidentiality`

GetConfidentiality returns the Confidentiality field if non-nil, zero value otherwise.

### GetConfidentialityOk

`func (o *ProjectSearch) GetConfidentialityOk() (*[]Confidentiality, bool)`

GetConfidentialityOk returns a tuple with the Confidentiality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentiality

`func (o *ProjectSearch) SetConfidentiality(v []Confidentiality)`

SetConfidentiality sets Confidentiality field to given value.

### HasConfidentiality

`func (o *ProjectSearch) HasConfidentiality() bool`

HasConfidentiality returns a boolean if a field has been set.

### GetParticipantIds

`func (o *ProjectSearch) GetParticipantIds() []string`

GetParticipantIds returns the ParticipantIds field if non-nil, zero value otherwise.

### GetParticipantIdsOk

`func (o *ProjectSearch) GetParticipantIdsOk() (*[]string, bool)`

GetParticipantIdsOk returns a tuple with the ParticipantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantIds

`func (o *ProjectSearch) SetParticipantIds(v []string)`

SetParticipantIds sets ParticipantIds field to given value.

### HasParticipantIds

`func (o *ProjectSearch) HasParticipantIds() bool`

HasParticipantIds returns a boolean if a field has been set.

### GetIndustryIds

`func (o *ProjectSearch) GetIndustryIds() []string`

GetIndustryIds returns the IndustryIds field if non-nil, zero value otherwise.

### GetIndustryIdsOk

`func (o *ProjectSearch) GetIndustryIdsOk() (*[]string, bool)`

GetIndustryIdsOk returns a tuple with the IndustryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryIds

`func (o *ProjectSearch) SetIndustryIds(v []string)`

SetIndustryIds sets IndustryIds field to given value.

### HasIndustryIds

`func (o *ProjectSearch) HasIndustryIds() bool`

HasIndustryIds returns a boolean if a field has been set.

### GetOrganizationIds

`func (o *ProjectSearch) GetOrganizationIds() []string`

GetOrganizationIds returns the OrganizationIds field if non-nil, zero value otherwise.

### GetOrganizationIdsOk

`func (o *ProjectSearch) GetOrganizationIdsOk() (*[]string, bool)`

GetOrganizationIdsOk returns a tuple with the OrganizationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationIds

`func (o *ProjectSearch) SetOrganizationIds(v []string)`

SetOrganizationIds sets OrganizationIds field to given value.

### HasOrganizationIds

`func (o *ProjectSearch) HasOrganizationIds() bool`

HasOrganizationIds returns a boolean if a field has been set.

### GetAwardedParticipant

`func (o *ProjectSearch) GetAwardedParticipant() AwardedParticipant`

GetAwardedParticipant returns the AwardedParticipant field if non-nil, zero value otherwise.

### GetAwardedParticipantOk

`func (o *ProjectSearch) GetAwardedParticipantOk() (*AwardedParticipant, bool)`

GetAwardedParticipantOk returns a tuple with the AwardedParticipant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwardedParticipant

`func (o *ProjectSearch) SetAwardedParticipant(v AwardedParticipant)`

SetAwardedParticipant sets AwardedParticipant field to given value.

### HasAwardedParticipant

`func (o *ProjectSearch) HasAwardedParticipant() bool`

HasAwardedParticipant returns a boolean if a field has been set.

### GetExternal

`func (o *ProjectSearch) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *ProjectSearch) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *ProjectSearch) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *ProjectSearch) HasExternal() bool`

HasExternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


