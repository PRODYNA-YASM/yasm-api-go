# SkillGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**Skill**](Skill.md) |  | [optional] 
**Skills** | Pointer to [**[]Skill**](Skill.md) |  | [optional] 

## Methods

### NewSkillGroup

`func NewSkillGroup() *SkillGroup`

NewSkillGroup instantiates a new SkillGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillGroupWithDefaults

`func NewSkillGroupWithDefaults() *SkillGroup`

NewSkillGroupWithDefaults instantiates a new SkillGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *SkillGroup) GetGroup() Skill`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SkillGroup) GetGroupOk() (*Skill, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SkillGroup) SetGroup(v Skill)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *SkillGroup) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetSkills

`func (o *SkillGroup) GetSkills() []Skill`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *SkillGroup) GetSkillsOk() (*[]Skill, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *SkillGroup) SetSkills(v []Skill)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *SkillGroup) HasSkills() bool`

HasSkills returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


