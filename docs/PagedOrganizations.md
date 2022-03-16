# PagedOrganizations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Organizations** | Pointer to [**[]OrganizationDetails**](OrganizationDetails.md) |  | [optional] 

## Methods

### NewPagedOrganizations

`func NewPagedOrganizations() *PagedOrganizations`

NewPagedOrganizations instantiates a new PagedOrganizations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagedOrganizationsWithDefaults

`func NewPagedOrganizationsWithDefaults() *PagedOrganizations`

NewPagedOrganizationsWithDefaults instantiates a new PagedOrganizations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkip

`func (o *PagedOrganizations) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *PagedOrganizations) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *PagedOrganizations) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *PagedOrganizations) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetLimit

`func (o *PagedOrganizations) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PagedOrganizations) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PagedOrganizations) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PagedOrganizations) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetCount

`func (o *PagedOrganizations) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PagedOrganizations) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PagedOrganizations) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PagedOrganizations) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetOrganizations

`func (o *PagedOrganizations) GetOrganizations() []OrganizationDetails`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *PagedOrganizations) GetOrganizationsOk() (*[]OrganizationDetails, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *PagedOrganizations) SetOrganizations(v []OrganizationDetails)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *PagedOrganizations) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


