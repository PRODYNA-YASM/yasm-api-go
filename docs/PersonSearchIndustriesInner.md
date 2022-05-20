# PersonSearchIndustriesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ActiveProjects** | Pointer to **bool** |  | [optional] 
**AmountOfProjects** | Pointer to [**MinMax**](MinMax.md) |  | [optional] 
**ExperienceInMonth** | Pointer to [**MinMax**](MinMax.md) |  | [optional] 

## Methods

### NewPersonSearchIndustriesInner

`func NewPersonSearchIndustriesInner(id string, ) *PersonSearchIndustriesInner`

NewPersonSearchIndustriesInner instantiates a new PersonSearchIndustriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonSearchIndustriesInnerWithDefaults

`func NewPersonSearchIndustriesInnerWithDefaults() *PersonSearchIndustriesInner`

NewPersonSearchIndustriesInnerWithDefaults instantiates a new PersonSearchIndustriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PersonSearchIndustriesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PersonSearchIndustriesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PersonSearchIndustriesInner) SetId(v string)`

SetId sets Id field to given value.


### GetActiveProjects

`func (o *PersonSearchIndustriesInner) GetActiveProjects() bool`

GetActiveProjects returns the ActiveProjects field if non-nil, zero value otherwise.

### GetActiveProjectsOk

`func (o *PersonSearchIndustriesInner) GetActiveProjectsOk() (*bool, bool)`

GetActiveProjectsOk returns a tuple with the ActiveProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveProjects

`func (o *PersonSearchIndustriesInner) SetActiveProjects(v bool)`

SetActiveProjects sets ActiveProjects field to given value.

### HasActiveProjects

`func (o *PersonSearchIndustriesInner) HasActiveProjects() bool`

HasActiveProjects returns a boolean if a field has been set.

### GetAmountOfProjects

`func (o *PersonSearchIndustriesInner) GetAmountOfProjects() MinMax`

GetAmountOfProjects returns the AmountOfProjects field if non-nil, zero value otherwise.

### GetAmountOfProjectsOk

`func (o *PersonSearchIndustriesInner) GetAmountOfProjectsOk() (*MinMax, bool)`

GetAmountOfProjectsOk returns a tuple with the AmountOfProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOfProjects

`func (o *PersonSearchIndustriesInner) SetAmountOfProjects(v MinMax)`

SetAmountOfProjects sets AmountOfProjects field to given value.

### HasAmountOfProjects

`func (o *PersonSearchIndustriesInner) HasAmountOfProjects() bool`

HasAmountOfProjects returns a boolean if a field has been set.

### GetExperienceInMonth

`func (o *PersonSearchIndustriesInner) GetExperienceInMonth() MinMax`

GetExperienceInMonth returns the ExperienceInMonth field if non-nil, zero value otherwise.

### GetExperienceInMonthOk

`func (o *PersonSearchIndustriesInner) GetExperienceInMonthOk() (*MinMax, bool)`

GetExperienceInMonthOk returns a tuple with the ExperienceInMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperienceInMonth

`func (o *PersonSearchIndustriesInner) SetExperienceInMonth(v MinMax)`

SetExperienceInMonth sets ExperienceInMonth field to given value.

### HasExperienceInMonth

`func (o *PersonSearchIndustriesInner) HasExperienceInMonth() bool`

HasExperienceInMonth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


