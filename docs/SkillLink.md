# SkillLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skill** | Pointer to [**Skill**](Skill.md) |  | [optional] 
**KindGiver** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSkillLink

`func NewSkillLink() *SkillLink`

NewSkillLink instantiates a new SkillLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillLinkWithDefaults

`func NewSkillLinkWithDefaults() *SkillLink`

NewSkillLinkWithDefaults instantiates a new SkillLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkill

`func (o *SkillLink) GetSkill() Skill`

GetSkill returns the Skill field if non-nil, zero value otherwise.

### GetSkillOk

`func (o *SkillLink) GetSkillOk() (*Skill, bool)`

GetSkillOk returns a tuple with the Skill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkill

`func (o *SkillLink) SetSkill(v Skill)`

SetSkill sets Skill field to given value.

### HasSkill

`func (o *SkillLink) HasSkill() bool`

HasSkill returns a boolean if a field has been set.

### GetKindGiver

`func (o *SkillLink) GetKindGiver() bool`

GetKindGiver returns the KindGiver field if non-nil, zero value otherwise.

### GetKindGiverOk

`func (o *SkillLink) GetKindGiverOk() (*bool, bool)`

GetKindGiverOk returns a tuple with the KindGiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKindGiver

`func (o *SkillLink) SetKindGiver(v bool)`

SetKindGiver sets KindGiver field to given value.

### HasKindGiver

`func (o *SkillLink) HasKindGiver() bool`

HasKindGiver returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


