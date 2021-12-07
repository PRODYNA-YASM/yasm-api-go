# ProjectSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinStartDate** | Pointer to **string** |  | [optional] 
**MaxEndDate** | Pointer to **string** |  | [optional] 
**ProjectStatuses** | Pointer to [**[]ProjectStatus**](ProjectStatus.md) |  | [optional] 
**ParticipationAmountInMonths** | Pointer to [**MinMax**](MinMax.md) |  | [optional] 
**InvolvedOfficeIds** | Pointer to **[]string** |  | [optional] 
**AmountOfInvolvedPersons** | Pointer to [**MinMax**](MinMax.md) |  | [optional] 
**Skills** | Pointer to [**[]EntityFilter**](EntityFilter.md) |  | [optional] 

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

### GetInvolvedOfficeIds

`func (o *ProjectSearch) GetInvolvedOfficeIds() []string`

GetInvolvedOfficeIds returns the InvolvedOfficeIds field if non-nil, zero value otherwise.

### GetInvolvedOfficeIdsOk

`func (o *ProjectSearch) GetInvolvedOfficeIdsOk() (*[]string, bool)`

GetInvolvedOfficeIdsOk returns a tuple with the InvolvedOfficeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvolvedOfficeIds

`func (o *ProjectSearch) SetInvolvedOfficeIds(v []string)`

SetInvolvedOfficeIds sets InvolvedOfficeIds field to given value.

### HasInvolvedOfficeIds

`func (o *ProjectSearch) HasInvolvedOfficeIds() bool`

HasInvolvedOfficeIds returns a boolean if a field has been set.

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

### GetSkills

`func (o *ProjectSearch) GetSkills() []EntityFilter`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *ProjectSearch) GetSkillsOk() (*[]EntityFilter, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *ProjectSearch) SetSkills(v []EntityFilter)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *ProjectSearch) HasSkills() bool`

HasSkills returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


