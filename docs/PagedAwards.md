# PagedAwards

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Awards** | Pointer to [**[]AwardDetails**](AwardDetails.md) |  | [optional] 

## Methods

### NewPagedAwards

`func NewPagedAwards() *PagedAwards`

NewPagedAwards instantiates a new PagedAwards object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagedAwardsWithDefaults

`func NewPagedAwardsWithDefaults() *PagedAwards`

NewPagedAwardsWithDefaults instantiates a new PagedAwards object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwards

`func (o *PagedAwards) GetAwards() []AwardDetails`

GetAwards returns the Awards field if non-nil, zero value otherwise.

### GetAwardsOk

`func (o *PagedAwards) GetAwardsOk() (*[]AwardDetails, bool)`

GetAwardsOk returns a tuple with the Awards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwards

`func (o *PagedAwards) SetAwards(v []AwardDetails)`

SetAwards sets Awards field to given value.

### HasAwards

`func (o *PagedAwards) HasAwards() bool`

HasAwards returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


