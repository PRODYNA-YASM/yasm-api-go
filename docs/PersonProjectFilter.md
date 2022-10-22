# PersonProjectFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParticipationInMonth** | Pointer to [**MinMax**](MinMax.md) |  | [optional] 
**InvolvedOfficeIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPersonProjectFilter

`func NewPersonProjectFilter() *PersonProjectFilter`

NewPersonProjectFilter instantiates a new PersonProjectFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonProjectFilterWithDefaults

`func NewPersonProjectFilterWithDefaults() *PersonProjectFilter`

NewPersonProjectFilterWithDefaults instantiates a new PersonProjectFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParticipationInMonth

`func (o *PersonProjectFilter) GetParticipationInMonth() MinMax`

GetParticipationInMonth returns the ParticipationInMonth field if non-nil, zero value otherwise.

### GetParticipationInMonthOk

`func (o *PersonProjectFilter) GetParticipationInMonthOk() (*MinMax, bool)`

GetParticipationInMonthOk returns a tuple with the ParticipationInMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipationInMonth

`func (o *PersonProjectFilter) SetParticipationInMonth(v MinMax)`

SetParticipationInMonth sets ParticipationInMonth field to given value.

### HasParticipationInMonth

`func (o *PersonProjectFilter) HasParticipationInMonth() bool`

HasParticipationInMonth returns a boolean if a field has been set.

### GetInvolvedOfficeIds

`func (o *PersonProjectFilter) GetInvolvedOfficeIds() []string`

GetInvolvedOfficeIds returns the InvolvedOfficeIds field if non-nil, zero value otherwise.

### GetInvolvedOfficeIdsOk

`func (o *PersonProjectFilter) GetInvolvedOfficeIdsOk() (*[]string, bool)`

GetInvolvedOfficeIdsOk returns a tuple with the InvolvedOfficeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvolvedOfficeIds

`func (o *PersonProjectFilter) SetInvolvedOfficeIds(v []string)`

SetInvolvedOfficeIds sets InvolvedOfficeIds field to given value.

### HasInvolvedOfficeIds

`func (o *PersonProjectFilter) HasInvolvedOfficeIds() bool`

HasInvolvedOfficeIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


