# ProjectParticipationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeframe** | Pointer to [**Timeframed**](Timeframed.md) |  | [optional] 
**Project** | [**Project**](Project.md) |  | [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Experiences** | [**[]Experience**](Experience.md) |  | [readonly] 

## Methods

### NewProjectParticipationAllOf

`func NewProjectParticipationAllOf(project Project, experiences []Experience, ) *ProjectParticipationAllOf`

NewProjectParticipationAllOf instantiates a new ProjectParticipationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectParticipationAllOfWithDefaults

`func NewProjectParticipationAllOfWithDefaults() *ProjectParticipationAllOf`

NewProjectParticipationAllOfWithDefaults instantiates a new ProjectParticipationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeframe

`func (o *ProjectParticipationAllOf) GetTimeframe() Timeframed`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *ProjectParticipationAllOf) GetTimeframeOk() (*Timeframed, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *ProjectParticipationAllOf) SetTimeframe(v Timeframed)`

SetTimeframe sets Timeframe field to given value.

### HasTimeframe

`func (o *ProjectParticipationAllOf) HasTimeframe() bool`

HasTimeframe returns a boolean if a field has been set.

### GetProject

`func (o *ProjectParticipationAllOf) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ProjectParticipationAllOf) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ProjectParticipationAllOf) SetProject(v Project)`

SetProject sets Project field to given value.


### GetDescription

`func (o *ProjectParticipationAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectParticipationAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectParticipationAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectParticipationAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExperiences

`func (o *ProjectParticipationAllOf) GetExperiences() []Experience`

GetExperiences returns the Experiences field if non-nil, zero value otherwise.

### GetExperiencesOk

`func (o *ProjectParticipationAllOf) GetExperiencesOk() (*[]Experience, bool)`

GetExperiencesOk returns a tuple with the Experiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiences

`func (o *ProjectParticipationAllOf) SetExperiences(v []Experience)`

SetExperiences sets Experiences field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


