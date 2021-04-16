# ExperienceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skill** | Pointer to [**SkillLevel**](SkillLevel.md) |  | [optional] 
**ConfirmedBy** | Pointer to [**[]Person**](Person.md) |  | [optional] 

## Methods

### NewExperienceAllOf

`func NewExperienceAllOf() *ExperienceAllOf`

NewExperienceAllOf instantiates a new ExperienceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperienceAllOfWithDefaults

`func NewExperienceAllOfWithDefaults() *ExperienceAllOf`

NewExperienceAllOfWithDefaults instantiates a new ExperienceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkill

`func (o *ExperienceAllOf) GetSkill() SkillLevel`

GetSkill returns the Skill field if non-nil, zero value otherwise.

### GetSkillOk

`func (o *ExperienceAllOf) GetSkillOk() (*SkillLevel, bool)`

GetSkillOk returns a tuple with the Skill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkill

`func (o *ExperienceAllOf) SetSkill(v SkillLevel)`

SetSkill sets Skill field to given value.

### HasSkill

`func (o *ExperienceAllOf) HasSkill() bool`

HasSkill returns a boolean if a field has been set.

### GetConfirmedBy

`func (o *ExperienceAllOf) GetConfirmedBy() []Person`

GetConfirmedBy returns the ConfirmedBy field if non-nil, zero value otherwise.

### GetConfirmedByOk

`func (o *ExperienceAllOf) GetConfirmedByOk() (*[]Person, bool)`

GetConfirmedByOk returns a tuple with the ConfirmedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedBy

`func (o *ExperienceAllOf) SetConfirmedBy(v []Person)`

SetConfirmedBy sets ConfirmedBy field to given value.

### HasConfirmedBy

`func (o *ExperienceAllOf) HasConfirmedBy() bool`

HasConfirmedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


