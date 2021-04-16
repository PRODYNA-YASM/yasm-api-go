# PagedCountries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Countries** | Pointer to [**[]Country**](Country.md) |  | [optional] 

## Methods

### NewPagedCountries

`func NewPagedCountries() *PagedCountries`

NewPagedCountries instantiates a new PagedCountries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagedCountriesWithDefaults

`func NewPagedCountriesWithDefaults() *PagedCountries`

NewPagedCountriesWithDefaults instantiates a new PagedCountries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkip

`func (o *PagedCountries) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *PagedCountries) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *PagedCountries) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *PagedCountries) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetLimit

`func (o *PagedCountries) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PagedCountries) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PagedCountries) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PagedCountries) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetCount

`func (o *PagedCountries) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PagedCountries) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PagedCountries) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PagedCountries) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCountries

`func (o *PagedCountries) GetCountries() []Country`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *PagedCountries) GetCountriesOk() (*[]Country, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *PagedCountries) SetCountries(v []Country)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *PagedCountries) HasCountries() bool`

HasCountries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


