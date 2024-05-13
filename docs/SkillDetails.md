# SkillDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skill** | Pointer to [**Skill**](Skill.md) |  | [optional] 
**ParentPath** | Pointer to [**[]Skill**](Skill.md) |  | [optional] 
**Children** | Pointer to [**[]SkillLink**](SkillLink.md) |  | [optional] 
**AdditionalParents** | Pointer to [**[]Skill**](Skill.md) |  | [optional] 

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

### GetParentPath

`func (o *SkillDetails) GetParentPath() []Skill`

GetParentPath returns the ParentPath field if non-nil, zero value otherwise.

### GetParentPathOk

`func (o *SkillDetails) GetParentPathOk() (*[]Skill, bool)`

GetParentPathOk returns a tuple with the ParentPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPath

`func (o *SkillDetails) SetParentPath(v []Skill)`

SetParentPath sets ParentPath field to given value.

### HasParentPath

`func (o *SkillDetails) HasParentPath() bool`

HasParentPath returns a boolean if a field has been set.

### GetChildren

`func (o *SkillDetails) GetChildren() []SkillLink`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *SkillDetails) GetChildrenOk() (*[]SkillLink, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *SkillDetails) SetChildren(v []SkillLink)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *SkillDetails) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetAdditionalParents

`func (o *SkillDetails) GetAdditionalParents() []Skill`

GetAdditionalParents returns the AdditionalParents field if non-nil, zero value otherwise.

### GetAdditionalParentsOk

`func (o *SkillDetails) GetAdditionalParentsOk() (*[]Skill, bool)`

GetAdditionalParentsOk returns a tuple with the AdditionalParents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalParents

`func (o *SkillDetails) SetAdditionalParents(v []Skill)`

SetAdditionalParents sets AdditionalParents field to given value.

### HasAdditionalParents

`func (o *SkillDetails) HasAdditionalParents() bool`

HasAdditionalParents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


