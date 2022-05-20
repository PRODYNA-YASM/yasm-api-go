# PersonSearchProjectsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ParticipationInMonth** | Pointer to [**MinMax**](MinMax.md) |  | [optional] 
**InvolevedOfficeIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPersonSearchProjectsInner

`func NewPersonSearchProjectsInner(id string, ) *PersonSearchProjectsInner`

NewPersonSearchProjectsInner instantiates a new PersonSearchProjectsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonSearchProjectsInnerWithDefaults

`func NewPersonSearchProjectsInnerWithDefaults() *PersonSearchProjectsInner`

NewPersonSearchProjectsInnerWithDefaults instantiates a new PersonSearchProjectsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PersonSearchProjectsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PersonSearchProjectsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PersonSearchProjectsInner) SetId(v string)`

SetId sets Id field to given value.


### GetParticipationInMonth

`func (o *PersonSearchProjectsInner) GetParticipationInMonth() MinMax`

GetParticipationInMonth returns the ParticipationInMonth field if non-nil, zero value otherwise.

### GetParticipationInMonthOk

`func (o *PersonSearchProjectsInner) GetParticipationInMonthOk() (*MinMax, bool)`

GetParticipationInMonthOk returns a tuple with the ParticipationInMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipationInMonth

`func (o *PersonSearchProjectsInner) SetParticipationInMonth(v MinMax)`

SetParticipationInMonth sets ParticipationInMonth field to given value.

### HasParticipationInMonth

`func (o *PersonSearchProjectsInner) HasParticipationInMonth() bool`

HasParticipationInMonth returns a boolean if a field has been set.

### GetInvolevedOfficeIds

`func (o *PersonSearchProjectsInner) GetInvolevedOfficeIds() []string`

GetInvolevedOfficeIds returns the InvolevedOfficeIds field if non-nil, zero value otherwise.

### GetInvolevedOfficeIdsOk

`func (o *PersonSearchProjectsInner) GetInvolevedOfficeIdsOk() (*[]string, bool)`

GetInvolevedOfficeIdsOk returns a tuple with the InvolevedOfficeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvolevedOfficeIds

`func (o *PersonSearchProjectsInner) SetInvolevedOfficeIds(v []string)`

SetInvolevedOfficeIds sets InvolevedOfficeIds field to given value.

### HasInvolevedOfficeIds

`func (o *PersonSearchProjectsInner) HasInvolevedOfficeIds() bool`

HasInvolevedOfficeIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


