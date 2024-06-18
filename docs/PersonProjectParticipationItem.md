# PersonProjectParticipationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Participation** | Pointer to [**ProjectParticipation**](ProjectParticipation.md) |  | [optional] 
**Experiences** | Pointer to [**[]Experience**](Experience.md) |  | [optional] 

## Methods

### NewPersonProjectParticipationItem

`func NewPersonProjectParticipationItem() *PersonProjectParticipationItem`

NewPersonProjectParticipationItem instantiates a new PersonProjectParticipationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonProjectParticipationItemWithDefaults

`func NewPersonProjectParticipationItemWithDefaults() *PersonProjectParticipationItem`

NewPersonProjectParticipationItemWithDefaults instantiates a new PersonProjectParticipationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParticipation

`func (o *PersonProjectParticipationItem) GetParticipation() ProjectParticipation`

GetParticipation returns the Participation field if non-nil, zero value otherwise.

### GetParticipationOk

`func (o *PersonProjectParticipationItem) GetParticipationOk() (*ProjectParticipation, bool)`

GetParticipationOk returns a tuple with the Participation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipation

`func (o *PersonProjectParticipationItem) SetParticipation(v ProjectParticipation)`

SetParticipation sets Participation field to given value.

### HasParticipation

`func (o *PersonProjectParticipationItem) HasParticipation() bool`

HasParticipation returns a boolean if a field has been set.

### GetExperiences

`func (o *PersonProjectParticipationItem) GetExperiences() []Experience`

GetExperiences returns the Experiences field if non-nil, zero value otherwise.

### GetExperiencesOk

`func (o *PersonProjectParticipationItem) GetExperiencesOk() (*[]Experience, bool)`

GetExperiencesOk returns a tuple with the Experiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiences

`func (o *PersonProjectParticipationItem) SetExperiences(v []Experience)`

SetExperiences sets Experiences field to given value.

### HasExperiences

`func (o *PersonProjectParticipationItem) HasExperiences() bool`

HasExperiences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


