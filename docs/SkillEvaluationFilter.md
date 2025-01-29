# SkillEvaluationFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | **string** |  | 
**EndDate** | **string** |  | 
**EmployeeIds** | Pointer to **[]string** |  | [optional] 
**ErpIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSkillEvaluationFilter

`func NewSkillEvaluationFilter(startDate string, endDate string, ) *SkillEvaluationFilter`

NewSkillEvaluationFilter instantiates a new SkillEvaluationFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillEvaluationFilterWithDefaults

`func NewSkillEvaluationFilterWithDefaults() *SkillEvaluationFilter`

NewSkillEvaluationFilterWithDefaults instantiates a new SkillEvaluationFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *SkillEvaluationFilter) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SkillEvaluationFilter) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SkillEvaluationFilter) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *SkillEvaluationFilter) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *SkillEvaluationFilter) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *SkillEvaluationFilter) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetEmployeeIds

`func (o *SkillEvaluationFilter) GetEmployeeIds() []string`

GetEmployeeIds returns the EmployeeIds field if non-nil, zero value otherwise.

### GetEmployeeIdsOk

`func (o *SkillEvaluationFilter) GetEmployeeIdsOk() (*[]string, bool)`

GetEmployeeIdsOk returns a tuple with the EmployeeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeIds

`func (o *SkillEvaluationFilter) SetEmployeeIds(v []string)`

SetEmployeeIds sets EmployeeIds field to given value.

### HasEmployeeIds

`func (o *SkillEvaluationFilter) HasEmployeeIds() bool`

HasEmployeeIds returns a boolean if a field has been set.

### GetErpIds

`func (o *SkillEvaluationFilter) GetErpIds() []string`

GetErpIds returns the ErpIds field if non-nil, zero value otherwise.

### GetErpIdsOk

`func (o *SkillEvaluationFilter) GetErpIdsOk() (*[]string, bool)`

GetErpIdsOk returns a tuple with the ErpIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErpIds

`func (o *SkillEvaluationFilter) SetErpIds(v []string)`

SetErpIds sets ErpIds field to given value.

### HasErpIds

`func (o *SkillEvaluationFilter) HasErpIds() bool`

HasErpIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


