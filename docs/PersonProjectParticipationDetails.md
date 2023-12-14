# PersonProjectParticipationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectDetails** | Pointer to [**ProjectDetails**](ProjectDetails.md) |  | [optional] 
**ParticipationItems** | Pointer to [**[]PersonProjectParticipationItem**](PersonProjectParticipationItem.md) |  | [optional] 

## Methods

### NewPersonProjectParticipationDetails

`func NewPersonProjectParticipationDetails() *PersonProjectParticipationDetails`

NewPersonProjectParticipationDetails instantiates a new PersonProjectParticipationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonProjectParticipationDetailsWithDefaults

`func NewPersonProjectParticipationDetailsWithDefaults() *PersonProjectParticipationDetails`

NewPersonProjectParticipationDetailsWithDefaults instantiates a new PersonProjectParticipationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectDetails

`func (o *PersonProjectParticipationDetails) GetProjectDetails() ProjectDetails`

GetProjectDetails returns the ProjectDetails field if non-nil, zero value otherwise.

### GetProjectDetailsOk

`func (o *PersonProjectParticipationDetails) GetProjectDetailsOk() (*ProjectDetails, bool)`

GetProjectDetailsOk returns a tuple with the ProjectDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDetails

`func (o *PersonProjectParticipationDetails) SetProjectDetails(v ProjectDetails)`

SetProjectDetails sets ProjectDetails field to given value.

### HasProjectDetails

`func (o *PersonProjectParticipationDetails) HasProjectDetails() bool`

HasProjectDetails returns a boolean if a field has been set.

### GetParticipationItems

`func (o *PersonProjectParticipationDetails) GetParticipationItems() []PersonProjectParticipationItem`

GetParticipationItems returns the ParticipationItems field if non-nil, zero value otherwise.

### GetParticipationItemsOk

`func (o *PersonProjectParticipationDetails) GetParticipationItemsOk() (*[]PersonProjectParticipationItem, bool)`

GetParticipationItemsOk returns a tuple with the ParticipationItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipationItems

`func (o *PersonProjectParticipationDetails) SetParticipationItems(v []PersonProjectParticipationItem)`

SetParticipationItems sets ParticipationItems field to given value.

### HasParticipationItems

`func (o *PersonProjectParticipationDetails) HasParticipationItems() bool`

HasParticipationItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


