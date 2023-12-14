# ProjectParticipationCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeframe** | Pointer to [**Timeframed**](Timeframed.md) |  | [optional] 
**Skills** | Pointer to [**[]SkillLevelUpdate**](SkillLevelUpdate.md) |  | [optional] 
**PersonalDescription** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**ProjectId** | **string** |  | 
**PersonId** | **string** |  | 

## Methods

### NewProjectParticipationCreate

`func NewProjectParticipationCreate(id string, projectId string, personId string, ) *ProjectParticipationCreate`

NewProjectParticipationCreate instantiates a new ProjectParticipationCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectParticipationCreateWithDefaults

`func NewProjectParticipationCreateWithDefaults() *ProjectParticipationCreate`

NewProjectParticipationCreateWithDefaults instantiates a new ProjectParticipationCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeframe

`func (o *ProjectParticipationCreate) GetTimeframe() Timeframed`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *ProjectParticipationCreate) GetTimeframeOk() (*Timeframed, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *ProjectParticipationCreate) SetTimeframe(v Timeframed)`

SetTimeframe sets Timeframe field to given value.

### HasTimeframe

`func (o *ProjectParticipationCreate) HasTimeframe() bool`

HasTimeframe returns a boolean if a field has been set.

### GetSkills

`func (o *ProjectParticipationCreate) GetSkills() []SkillLevelUpdate`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *ProjectParticipationCreate) GetSkillsOk() (*[]SkillLevelUpdate, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *ProjectParticipationCreate) SetSkills(v []SkillLevelUpdate)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *ProjectParticipationCreate) HasSkills() bool`

HasSkills returns a boolean if a field has been set.

### GetPersonalDescription

`func (o *ProjectParticipationCreate) GetPersonalDescription() string`

GetPersonalDescription returns the PersonalDescription field if non-nil, zero value otherwise.

### GetPersonalDescriptionOk

`func (o *ProjectParticipationCreate) GetPersonalDescriptionOk() (*string, bool)`

GetPersonalDescriptionOk returns a tuple with the PersonalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalDescription

`func (o *ProjectParticipationCreate) SetPersonalDescription(v string)`

SetPersonalDescription sets PersonalDescription field to given value.

### HasPersonalDescription

`func (o *ProjectParticipationCreate) HasPersonalDescription() bool`

HasPersonalDescription returns a boolean if a field has been set.

### GetId

`func (o *ProjectParticipationCreate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectParticipationCreate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectParticipationCreate) SetId(v string)`

SetId sets Id field to given value.


### GetProjectId

`func (o *ProjectParticipationCreate) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectParticipationCreate) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectParticipationCreate) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetPersonId

`func (o *ProjectParticipationCreate) GetPersonId() string`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *ProjectParticipationCreate) GetPersonIdOk() (*string, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *ProjectParticipationCreate) SetPersonId(v string)`

SetPersonId sets PersonId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


