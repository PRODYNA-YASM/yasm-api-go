# PagedSkills

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Skills** | Pointer to [**[]Skill**](Skill.md) |  | [optional] 

## Methods

### NewPagedSkills

`func NewPagedSkills() *PagedSkills`

NewPagedSkills instantiates a new PagedSkills object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagedSkillsWithDefaults

`func NewPagedSkillsWithDefaults() *PagedSkills`

NewPagedSkillsWithDefaults instantiates a new PagedSkills object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkip

`func (o *PagedSkills) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *PagedSkills) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *PagedSkills) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *PagedSkills) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetLimit

`func (o *PagedSkills) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PagedSkills) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PagedSkills) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PagedSkills) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetCount

`func (o *PagedSkills) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PagedSkills) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PagedSkills) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PagedSkills) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSkills

`func (o *PagedSkills) GetSkills() []Skill`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *PagedSkills) GetSkillsOk() (*[]Skill, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *PagedSkills) SetSkills(v []Skill)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *PagedSkills) HasSkills() bool`

HasSkills returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


