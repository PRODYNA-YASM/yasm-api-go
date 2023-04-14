# ExperienceSkillGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**Skill**](Skill.md) |  | [optional] 
**Skills** | Pointer to [**[]ExperienceSkill**](ExperienceSkill.md) |  | [optional] 

## Methods

### NewExperienceSkillGroup

`func NewExperienceSkillGroup() *ExperienceSkillGroup`

NewExperienceSkillGroup instantiates a new ExperienceSkillGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperienceSkillGroupWithDefaults

`func NewExperienceSkillGroupWithDefaults() *ExperienceSkillGroup`

NewExperienceSkillGroupWithDefaults instantiates a new ExperienceSkillGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *ExperienceSkillGroup) GetGroup() Skill`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ExperienceSkillGroup) GetGroupOk() (*Skill, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ExperienceSkillGroup) SetGroup(v Skill)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ExperienceSkillGroup) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetSkills

`func (o *ExperienceSkillGroup) GetSkills() []ExperienceSkill`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *ExperienceSkillGroup) GetSkillsOk() (*[]ExperienceSkill, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *ExperienceSkillGroup) SetSkills(v []ExperienceSkill)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *ExperienceSkillGroup) HasSkills() bool`

HasSkills returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


