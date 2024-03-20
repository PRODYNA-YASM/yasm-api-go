# ProjectParticipationUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeframe** | [**Timeframed**](Timeframed.md) |  | 
**Skills** | Pointer to [**[]SkillLevelUpdate**](SkillLevelUpdate.md) |  | [optional] 
**PersonalDescription** | Pointer to **string** |  | [optional] 

## Methods

### NewProjectParticipationUpdate

`func NewProjectParticipationUpdate(timeframe Timeframed, ) *ProjectParticipationUpdate`

NewProjectParticipationUpdate instantiates a new ProjectParticipationUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectParticipationUpdateWithDefaults

`func NewProjectParticipationUpdateWithDefaults() *ProjectParticipationUpdate`

NewProjectParticipationUpdateWithDefaults instantiates a new ProjectParticipationUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeframe

`func (o *ProjectParticipationUpdate) GetTimeframe() Timeframed`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *ProjectParticipationUpdate) GetTimeframeOk() (*Timeframed, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *ProjectParticipationUpdate) SetTimeframe(v Timeframed)`

SetTimeframe sets Timeframe field to given value.


### GetSkills

`func (o *ProjectParticipationUpdate) GetSkills() []SkillLevelUpdate`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *ProjectParticipationUpdate) GetSkillsOk() (*[]SkillLevelUpdate, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *ProjectParticipationUpdate) SetSkills(v []SkillLevelUpdate)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *ProjectParticipationUpdate) HasSkills() bool`

HasSkills returns a boolean if a field has been set.

### GetPersonalDescription

`func (o *ProjectParticipationUpdate) GetPersonalDescription() string`

GetPersonalDescription returns the PersonalDescription field if non-nil, zero value otherwise.

### GetPersonalDescriptionOk

`func (o *ProjectParticipationUpdate) GetPersonalDescriptionOk() (*string, bool)`

GetPersonalDescriptionOk returns a tuple with the PersonalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalDescription

`func (o *ProjectParticipationUpdate) SetPersonalDescription(v string)`

SetPersonalDescription sets PersonalDescription field to given value.

### HasPersonalDescription

`func (o *ProjectParticipationUpdate) HasPersonalDescription() bool`

HasPersonalDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


