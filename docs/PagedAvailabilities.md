# PagedAvailabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Availabilities** | Pointer to [**[]AvailabilityDetail**](AvailabilityDetail.md) |  | [optional] 

## Methods

### NewPagedAvailabilities

`func NewPagedAvailabilities() *PagedAvailabilities`

NewPagedAvailabilities instantiates a new PagedAvailabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagedAvailabilitiesWithDefaults

`func NewPagedAvailabilitiesWithDefaults() *PagedAvailabilities`

NewPagedAvailabilitiesWithDefaults instantiates a new PagedAvailabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkip

`func (o *PagedAvailabilities) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *PagedAvailabilities) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *PagedAvailabilities) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *PagedAvailabilities) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetLimit

`func (o *PagedAvailabilities) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PagedAvailabilities) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PagedAvailabilities) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PagedAvailabilities) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetCount

`func (o *PagedAvailabilities) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PagedAvailabilities) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PagedAvailabilities) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PagedAvailabilities) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetAvailabilities

`func (o *PagedAvailabilities) GetAvailabilities() []AvailabilityDetail`

GetAvailabilities returns the Availabilities field if non-nil, zero value otherwise.

### GetAvailabilitiesOk

`func (o *PagedAvailabilities) GetAvailabilitiesOk() (*[]AvailabilityDetail, bool)`

GetAvailabilitiesOk returns a tuple with the Availabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilities

`func (o *PagedAvailabilities) SetAvailabilities(v []AvailabilityDetail)`

SetAvailabilities sets Availabilities field to given value.

### HasAvailabilities

`func (o *PagedAvailabilities) HasAvailabilities() bool`

HasAvailabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


