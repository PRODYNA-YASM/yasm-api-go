# ShoppingCartGroupDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShoppingCartGroup** | Pointer to [**ShoppingCartGroup**](ShoppingCartGroup.md) |  | [optional] 
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**Persons** | Pointer to [**[]Person**](Person.md) |  | [optional] 

## Methods

### NewShoppingCartGroupDetail

`func NewShoppingCartGroupDetail() *ShoppingCartGroupDetail`

NewShoppingCartGroupDetail instantiates a new ShoppingCartGroupDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShoppingCartGroupDetailWithDefaults

`func NewShoppingCartGroupDetailWithDefaults() *ShoppingCartGroupDetail`

NewShoppingCartGroupDetailWithDefaults instantiates a new ShoppingCartGroupDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShoppingCartGroup

`func (o *ShoppingCartGroupDetail) GetShoppingCartGroup() ShoppingCartGroup`

GetShoppingCartGroup returns the ShoppingCartGroup field if non-nil, zero value otherwise.

### GetShoppingCartGroupOk

`func (o *ShoppingCartGroupDetail) GetShoppingCartGroupOk() (*ShoppingCartGroup, bool)`

GetShoppingCartGroupOk returns a tuple with the ShoppingCartGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShoppingCartGroup

`func (o *ShoppingCartGroupDetail) SetShoppingCartGroup(v ShoppingCartGroup)`

SetShoppingCartGroup sets ShoppingCartGroup field to given value.

### HasShoppingCartGroup

`func (o *ShoppingCartGroupDetail) HasShoppingCartGroup() bool`

HasShoppingCartGroup returns a boolean if a field has been set.

### GetAudit

`func (o *ShoppingCartGroupDetail) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *ShoppingCartGroupDetail) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *ShoppingCartGroupDetail) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *ShoppingCartGroupDetail) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetPersons

`func (o *ShoppingCartGroupDetail) GetPersons() []Person`

GetPersons returns the Persons field if non-nil, zero value otherwise.

### GetPersonsOk

`func (o *ShoppingCartGroupDetail) GetPersonsOk() (*[]Person, bool)`

GetPersonsOk returns a tuple with the Persons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersons

`func (o *ShoppingCartGroupDetail) SetPersons(v []Person)`

SetPersons sets Persons field to given value.

### HasPersons

`func (o *ShoppingCartGroupDetail) HasPersons() bool`

HasPersons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


