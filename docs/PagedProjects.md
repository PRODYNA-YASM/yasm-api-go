# PagedProjects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Projects** | Pointer to [**[]ProjectScoreDetail**](ProjectScoreDetail.md) |  | [optional] 

## Methods

### NewPagedProjects

`func NewPagedProjects() *PagedProjects`

NewPagedProjects instantiates a new PagedProjects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagedProjectsWithDefaults

`func NewPagedProjectsWithDefaults() *PagedProjects`

NewPagedProjectsWithDefaults instantiates a new PagedProjects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkip

`func (o *PagedProjects) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *PagedProjects) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *PagedProjects) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *PagedProjects) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetLimit

`func (o *PagedProjects) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PagedProjects) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PagedProjects) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PagedProjects) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetCount

`func (o *PagedProjects) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PagedProjects) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PagedProjects) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PagedProjects) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetProjects

`func (o *PagedProjects) GetProjects() []ProjectScoreDetail`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *PagedProjects) GetProjectsOk() (*[]ProjectScoreDetail, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *PagedProjects) SetProjects(v []ProjectScoreDetail)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *PagedProjects) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


