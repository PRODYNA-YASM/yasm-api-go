# ProjectParticipation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ObjectType** | Pointer to **string** |  | [optional] 
**Timeframe** | Pointer to [**Timeframed**](Timeframed.md) |  | [optional] 
**ProjectDetails** | Pointer to [**ProjectDetails**](ProjectDetails.md) |  | [optional] 
**DescriptionOverwrite** | Pointer to **string** |  | [optional] 
**PersonalDescription** | Pointer to **string** |  | [optional] 
**Experiences** | [**[]Experience**](Experience.md) |  | [readonly] 

## Methods

### NewProjectParticipation

`func NewProjectParticipation(id string, experiences []Experience, ) *ProjectParticipation`

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


### GetObjectType

`func (o *ProjectParticipation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ProjectParticipation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ProjectParticipation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *ProjectParticipation) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

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

### GetProjectDetails

`func (o *ProjectParticipation) GetProjectDetails() ProjectDetails`

GetProjectDetails returns the ProjectDetails field if non-nil, zero value otherwise.

### GetProjectDetailsOk

`func (o *ProjectParticipation) GetProjectDetailsOk() (*ProjectDetails, bool)`

GetProjectDetailsOk returns a tuple with the ProjectDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDetails

`func (o *ProjectParticipation) SetProjectDetails(v ProjectDetails)`

SetProjectDetails sets ProjectDetails field to given value.

### HasProjectDetails

`func (o *ProjectParticipation) HasProjectDetails() bool`

HasProjectDetails returns a boolean if a field has been set.

### GetDescriptionOverwrite

`func (o *ProjectParticipation) GetDescriptionOverwrite() string`

GetDescriptionOverwrite returns the DescriptionOverwrite field if non-nil, zero value otherwise.

### GetDescriptionOverwriteOk

`func (o *ProjectParticipation) GetDescriptionOverwriteOk() (*string, bool)`

GetDescriptionOverwriteOk returns a tuple with the DescriptionOverwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionOverwrite

`func (o *ProjectParticipation) SetDescriptionOverwrite(v string)`

SetDescriptionOverwrite sets DescriptionOverwrite field to given value.

### HasDescriptionOverwrite

`func (o *ProjectParticipation) HasDescriptionOverwrite() bool`

HasDescriptionOverwrite returns a boolean if a field has been set.

### GetPersonalDescription

`func (o *ProjectParticipation) GetPersonalDescription() string`

GetPersonalDescription returns the PersonalDescription field if non-nil, zero value otherwise.

### GetPersonalDescriptionOk

`func (o *ProjectParticipation) GetPersonalDescriptionOk() (*string, bool)`

GetPersonalDescriptionOk returns a tuple with the PersonalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalDescription

`func (o *ProjectParticipation) SetPersonalDescription(v string)`

SetPersonalDescription sets PersonalDescription field to given value.

### HasPersonalDescription

`func (o *ProjectParticipation) HasPersonalDescription() bool`

HasPersonalDescription returns a boolean if a field has been set.

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


