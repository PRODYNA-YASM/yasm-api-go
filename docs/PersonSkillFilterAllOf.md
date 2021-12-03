# PersonSkillFilterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExperienceInMonth** | Pointer to [**MinMax**](MinMax.md) |  | [optional] 
**LastAssignment** | Pointer to **string** | filters the last time the employee used the skill in a project | [optional] 

## Methods

### NewPersonSkillFilterAllOf

`func NewPersonSkillFilterAllOf() *PersonSkillFilterAllOf`

NewPersonSkillFilterAllOf instantiates a new PersonSkillFilterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonSkillFilterAllOfWithDefaults

`func NewPersonSkillFilterAllOfWithDefaults() *PersonSkillFilterAllOf`

NewPersonSkillFilterAllOfWithDefaults instantiates a new PersonSkillFilterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExperienceInMonth

`func (o *PersonSkillFilterAllOf) GetExperienceInMonth() MinMax`

GetExperienceInMonth returns the ExperienceInMonth field if non-nil, zero value otherwise.

### GetExperienceInMonthOk

`func (o *PersonSkillFilterAllOf) GetExperienceInMonthOk() (*MinMax, bool)`

GetExperienceInMonthOk returns a tuple with the ExperienceInMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperienceInMonth

`func (o *PersonSkillFilterAllOf) SetExperienceInMonth(v MinMax)`

SetExperienceInMonth sets ExperienceInMonth field to given value.

### HasExperienceInMonth

`func (o *PersonSkillFilterAllOf) HasExperienceInMonth() bool`

HasExperienceInMonth returns a boolean if a field has been set.

### GetLastAssignment

`func (o *PersonSkillFilterAllOf) GetLastAssignment() string`

GetLastAssignment returns the LastAssignment field if non-nil, zero value otherwise.

### GetLastAssignmentOk

`func (o *PersonSkillFilterAllOf) GetLastAssignmentOk() (*string, bool)`

GetLastAssignmentOk returns a tuple with the LastAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAssignment

`func (o *PersonSkillFilterAllOf) SetLastAssignment(v string)`

SetLastAssignment sets LastAssignment field to given value.

### HasLastAssignment

`func (o *PersonSkillFilterAllOf) HasLastAssignment() bool`

HasLastAssignment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


