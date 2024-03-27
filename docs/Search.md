# Search

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Term** | Pointer to **string** |  | [optional] 
**Skip** | **int32** |  | [default to 0]
**Limit** | **int32** |  | [default to 20]
**ObjectType** | Pointer to **string** |  | [optional] 

## Methods

### NewSearch

`func NewSearch(skip int32, limit int32, ) *Search`

NewSearch instantiates a new Search object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchWithDefaults

`func NewSearchWithDefaults() *Search`

NewSearchWithDefaults instantiates a new Search object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTerm

`func (o *Search) GetTerm() string`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *Search) GetTermOk() (*string, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *Search) SetTerm(v string)`

SetTerm sets Term field to given value.

### HasTerm

`func (o *Search) HasTerm() bool`

HasTerm returns a boolean if a field has been set.

### GetSkip

`func (o *Search) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *Search) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *Search) SetSkip(v int32)`

SetSkip sets Skip field to given value.


### GetLimit

`func (o *Search) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Search) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Search) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetObjectType

`func (o *Search) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *Search) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *Search) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *Search) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


