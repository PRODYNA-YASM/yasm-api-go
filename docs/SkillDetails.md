# SkillDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skill** | Pointer to [**Skill**](Skill.md) |  | [optional] 
**Children** | Pointer to [**[]Skill**](Skill.md) |  | [optional] 
**Kinds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSkillDetails

`func NewSkillDetails() *SkillDetails`

NewSkillDetails instantiates a new SkillDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillDetailsWithDefaults

`func NewSkillDetailsWithDefaults() *SkillDetails`

NewSkillDetailsWithDefaults instantiates a new SkillDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkill

`func (o *SkillDetails) GetSkill() Skill`

GetSkill returns the Skill field if non-nil, zero value otherwise.

### GetSkillOk

`func (o *SkillDetails) GetSkillOk() (*Skill, bool)`

GetSkillOk returns a tuple with the Skill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkill

`func (o *SkillDetails) SetSkill(v Skill)`

SetSkill sets Skill field to given value.

### HasSkill

`func (o *SkillDetails) HasSkill() bool`

HasSkill returns a boolean if a field has been set.

### GetChildren

`func (o *SkillDetails) GetChildren() []Skill`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *SkillDetails) GetChildrenOk() (*[]Skill, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *SkillDetails) SetChildren(v []Skill)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *SkillDetails) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetKinds

`func (o *SkillDetails) GetKinds() []string`

GetKinds returns the Kinds field if non-nil, zero value otherwise.

### GetKindsOk

`func (o *SkillDetails) GetKindsOk() (*[]string, bool)`

GetKindsOk returns a tuple with the Kinds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKinds

`func (o *SkillDetails) SetKinds(v []string)`

SetKinds sets Kinds field to given value.

### HasKinds

`func (o *SkillDetails) HasKinds() bool`

HasKinds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


