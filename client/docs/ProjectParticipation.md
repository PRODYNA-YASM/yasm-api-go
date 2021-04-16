# ProjectParticipation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Startdate** | [**Date**](Date.md) |  | 
**Enddate** | Pointer to [**Date**](Date.md) |  | [optional] 
**Project** | [**Project**](Project.md) |  | [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Experiences** | [**[]Experience**](Experience.md) |  | [readonly] 

## Methods

### NewProjectParticipation

`func NewProjectParticipation(id string, startdate Date, project Project, experiences []Experience, ) *ProjectParticipation`

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


### GetStartdate

`func (o *ProjectParticipation) GetStartdate() Date`

GetStartdate returns the Startdate field if non-nil, zero value otherwise.

### GetStartdateOk

`func (o *ProjectParticipation) GetStartdateOk() (*Date, bool)`

GetStartdateOk returns a tuple with the Startdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartdate

`func (o *ProjectParticipation) SetStartdate(v Date)`

SetStartdate sets Startdate field to given value.


### GetEnddate

`func (o *ProjectParticipation) GetEnddate() Date`

GetEnddate returns the Enddate field if non-nil, zero value otherwise.

### GetEnddateOk

`func (o *ProjectParticipation) GetEnddateOk() (*Date, bool)`

GetEnddateOk returns a tuple with the Enddate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnddate

`func (o *ProjectParticipation) SetEnddate(v Date)`

SetEnddate sets Enddate field to given value.

### HasEnddate

`func (o *ProjectParticipation) HasEnddate() bool`

HasEnddate returns a boolean if a field has been set.

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


