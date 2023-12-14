# Experience

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkillLevel** | Pointer to [**SkillLevel**](SkillLevel.md) |  | [optional] 
**ConfirmedBy** | Pointer to [**[]Person**](Person.md) |  | [optional] 

## Methods

### NewExperience

`func NewExperience() *Experience`

NewExperience instantiates a new Experience object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperienceWithDefaults

`func NewExperienceWithDefaults() *Experience`

NewExperienceWithDefaults instantiates a new Experience object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkillLevel

`func (o *Experience) GetSkillLevel() SkillLevel`

GetSkillLevel returns the SkillLevel field if non-nil, zero value otherwise.

### GetSkillLevelOk

`func (o *Experience) GetSkillLevelOk() (*SkillLevel, bool)`

GetSkillLevelOk returns a tuple with the SkillLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillLevel

`func (o *Experience) SetSkillLevel(v SkillLevel)`

SetSkillLevel sets SkillLevel field to given value.

### HasSkillLevel

`func (o *Experience) HasSkillLevel() bool`

HasSkillLevel returns a boolean if a field has been set.

### GetConfirmedBy

`func (o *Experience) GetConfirmedBy() []Person`

GetConfirmedBy returns the ConfirmedBy field if non-nil, zero value otherwise.

### GetConfirmedByOk

`func (o *Experience) GetConfirmedByOk() (*[]Person, bool)`

GetConfirmedByOk returns a tuple with the ConfirmedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedBy

`func (o *Experience) SetConfirmedBy(v []Person)`

SetConfirmedBy sets ConfirmedBy field to given value.

### HasConfirmedBy

`func (o *Experience) HasConfirmedBy() bool`

HasConfirmedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


