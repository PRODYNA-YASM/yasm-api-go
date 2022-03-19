# PagedIndustries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Industries** | Pointer to [**[]IndutryDetails**](IndutryDetails.md) |  | [optional] 

## Methods

### NewPagedIndustries

`func NewPagedIndustries() *PagedIndustries`

NewPagedIndustries instantiates a new PagedIndustries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagedIndustriesWithDefaults

`func NewPagedIndustriesWithDefaults() *PagedIndustries`

NewPagedIndustriesWithDefaults instantiates a new PagedIndustries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkip

`func (o *PagedIndustries) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *PagedIndustries) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *PagedIndustries) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *PagedIndustries) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetLimit

`func (o *PagedIndustries) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PagedIndustries) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PagedIndustries) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PagedIndustries) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetCount

`func (o *PagedIndustries) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PagedIndustries) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PagedIndustries) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PagedIndustries) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetIndustries

`func (o *PagedIndustries) GetIndustries() []IndutryDetails`

GetIndustries returns the Industries field if non-nil, zero value otherwise.

### GetIndustriesOk

`func (o *PagedIndustries) GetIndustriesOk() (*[]IndutryDetails, bool)`

GetIndustriesOk returns a tuple with the Industries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustries

`func (o *PagedIndustries) SetIndustries(v []IndutryDetails)`

SetIndustries sets Industries field to given value.

### HasIndustries

`func (o *PagedIndustries) HasIndustries() bool`

HasIndustries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


