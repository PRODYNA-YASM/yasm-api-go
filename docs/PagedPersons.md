# PagedPersons

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Persons** | Pointer to [**[]PersonScoreDetail**](PersonScoreDetail.md) |  | [optional] 

## Methods

### NewPagedPersons

`func NewPagedPersons() *PagedPersons`

NewPagedPersons instantiates a new PagedPersons object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagedPersonsWithDefaults

`func NewPagedPersonsWithDefaults() *PagedPersons`

NewPagedPersonsWithDefaults instantiates a new PagedPersons object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkip

`func (o *PagedPersons) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *PagedPersons) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *PagedPersons) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *PagedPersons) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetLimit

`func (o *PagedPersons) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PagedPersons) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PagedPersons) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PagedPersons) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetCount

`func (o *PagedPersons) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PagedPersons) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PagedPersons) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PagedPersons) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetPersons

`func (o *PagedPersons) GetPersons() []PersonScoreDetail`

GetPersons returns the Persons field if non-nil, zero value otherwise.

### GetPersonsOk

`func (o *PagedPersons) GetPersonsOk() (*[]PersonScoreDetail, bool)`

GetPersonsOk returns a tuple with the Persons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersons

`func (o *PagedPersons) SetPersons(v []PersonScoreDetail)`

SetPersons sets Persons field to given value.

### HasPersons

`func (o *PagedPersons) HasPersons() bool`

HasPersons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


