# ProjectParticipation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Timeframe** | Pointer to [**Timeframed**](Timeframed.md) |  | [optional] 
**Project** | [**Project**](Project.md) |  | [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Experiences** | [**[]Experience**](Experience.md) |  | [readonly] 

## Methods

### NewProjectParticipation

`func NewProjectParticipation(id string, project Project, experiences []Experience, ) *ProjectParticipation`

NewProjectParticipation instantiates a new ProjectParticipation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectParticipationWithDefaults

`func NewProjectParticipationWithDefaults() *ProjectParticipation`

NewProjectParticipationWithDefaults instantiates a new ProjectParticipation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectParticipation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectParticipation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectParticipation) SetId(v string)`

SetId sets Id field to given value.


### GetTimeframe

`func (o *ProjectParticipation) GetTimeframe() Timeframed`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *ProjectParticipation) GetTimeframeOk() (*Timeframed, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *ProjectParticipation) SetTimeframe(v Timeframed)`

SetTimeframe sets Timeframe field to given value.

### HasTimeframe

`func (o *ProjectParticipation) HasTimeframe() bool`

HasTimeframe returns a boolean if a field has been set.

### GetProject

`func (o *ProjectParticipation) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ProjectParticipation) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ProjectParticipation) SetProject(v Project)`

SetProject sets Project field to given value.


### GetDescription

`func (o *ProjectParticipation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectParticipation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectParticipation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectParticipation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExperiences

`func (o *ProjectParticipation) GetExperiences() []Experience`

GetExperiences returns the Experiences field if non-nil, zero value otherwise.

### GetExperiencesOk

`func (o *ProjectParticipation) GetExperiencesOk() (*[]Experience, bool)`

GetExperiencesOk returns a tuple with the Experiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiences

`func (o *ProjectParticipation) SetExperiences(v []Experience)`

SetExperiences sets Experiences field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


