# PersonSkillFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ObjectType** | Pointer to **string** |  | [optional] 
**ExperienceInMonth** | Pointer to [**MinMax**](MinMax.md) |  | [optional] 
**LastAssignment** | Pointer to **string** | filters the last time the employee used the skill in a project | [optional] 

## Methods

### NewPersonSkillFilter

`func NewPersonSkillFilter(id string, ) *PersonSkillFilter`

NewPersonSkillFilter instantiates a new PersonSkillFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonSkillFilterWithDefaults

`func NewPersonSkillFilterWithDefaults() *PersonSkillFilter`

NewPersonSkillFilterWithDefaults instantiates a new PersonSkillFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PersonSkillFilter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PersonSkillFilter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PersonSkillFilter) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *PersonSkillFilter) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PersonSkillFilter) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PersonSkillFilter) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *PersonSkillFilter) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetExperienceInMonth

`func (o *PersonSkillFilter) GetExperienceInMonth() MinMax`

GetExperienceInMonth returns the ExperienceInMonth field if non-nil, zero value otherwise.

### GetExperienceInMonthOk

`func (o *PersonSkillFilter) GetExperienceInMonthOk() (*MinMax, bool)`

GetExperienceInMonthOk returns a tuple with the ExperienceInMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperienceInMonth

`func (o *PersonSkillFilter) SetExperienceInMonth(v MinMax)`

SetExperienceInMonth sets ExperienceInMonth field to given value.

### HasExperienceInMonth

`func (o *PersonSkillFilter) HasExperienceInMonth() bool`

HasExperienceInMonth returns a boolean if a field has been set.

### GetLastAssignment

`func (o *PersonSkillFilter) GetLastAssignment() string`

GetLastAssignment returns the LastAssignment field if non-nil, zero value otherwise.

### GetLastAssignmentOk

`func (o *PersonSkillFilter) GetLastAssignmentOk() (*string, bool)`

GetLastAssignmentOk returns a tuple with the LastAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAssignment

`func (o *PersonSkillFilter) SetLastAssignment(v string)`

SetLastAssignment sets LastAssignment field to given value.

### HasLastAssignment

`func (o *PersonSkillFilter) HasLastAssignment() bool`

HasLastAssignment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


