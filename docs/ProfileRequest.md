# ProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | Pointer to **string** |  | [optional] 
**Person** | Pointer to [**PersonDetails**](PersonDetails.md) |  | [optional] 
**ProjectParticipation** | Pointer to [**[]PersonProjectParticipationDetails**](PersonProjectParticipationDetails.md) |  | [optional] 

## Methods

### NewProfileRequest

`func NewProfileRequest() *ProfileRequest`

NewProfileRequest instantiates a new ProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileRequestWithDefaults

`func NewProfileRequestWithDefaults() *ProfileRequest`

NewProfileRequestWithDefaults instantiates a new ProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *ProfileRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ProfileRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ProfileRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *ProfileRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetPerson

`func (o *ProfileRequest) GetPerson() PersonDetails`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *ProfileRequest) GetPersonOk() (*PersonDetails, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *ProfileRequest) SetPerson(v PersonDetails)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *ProfileRequest) HasPerson() bool`

HasPerson returns a boolean if a field has been set.

### GetProjectParticipation

`func (o *ProfileRequest) GetProjectParticipation() []PersonProjectParticipationDetails`

GetProjectParticipation returns the ProjectParticipation field if non-nil, zero value otherwise.

### GetProjectParticipationOk

`func (o *ProfileRequest) GetProjectParticipationOk() (*[]PersonProjectParticipationDetails, bool)`

GetProjectParticipationOk returns a tuple with the ProjectParticipation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectParticipation

`func (o *ProfileRequest) SetProjectParticipation(v []PersonProjectParticipationDetails)`

SetProjectParticipation sets ProjectParticipation field to given value.

### HasProjectParticipation

`func (o *ProfileRequest) HasProjectParticipation() bool`

HasProjectParticipation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


