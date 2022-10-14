# PagedPersons

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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


