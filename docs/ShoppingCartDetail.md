# ShoppingCartDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShoppingCart** | Pointer to [**ShoppingCart**](ShoppingCart.md) |  | [optional] 
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**Owner** | Pointer to [**Person**](Person.md) |  | [optional] 
**SharedWith** | Pointer to [**[]Person**](Person.md) |  | [optional] 
**Persons** | Pointer to [**[]Person**](Person.md) |  | [optional] 

## Methods

### NewShoppingCartDetail

`func NewShoppingCartDetail() *ShoppingCartDetail`

NewShoppingCartDetail instantiates a new ShoppingCartDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShoppingCartDetailWithDefaults

`func NewShoppingCartDetailWithDefaults() *ShoppingCartDetail`

NewShoppingCartDetailWithDefaults instantiates a new ShoppingCartDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShoppingCart

`func (o *ShoppingCartDetail) GetShoppingCart() ShoppingCart`

GetShoppingCart returns the ShoppingCart field if non-nil, zero value otherwise.

### GetShoppingCartOk

`func (o *ShoppingCartDetail) GetShoppingCartOk() (*ShoppingCart, bool)`

GetShoppingCartOk returns a tuple with the ShoppingCart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShoppingCart

`func (o *ShoppingCartDetail) SetShoppingCart(v ShoppingCart)`

SetShoppingCart sets ShoppingCart field to given value.

### HasShoppingCart

`func (o *ShoppingCartDetail) HasShoppingCart() bool`

HasShoppingCart returns a boolean if a field has been set.

### GetAudit

`func (o *ShoppingCartDetail) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *ShoppingCartDetail) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *ShoppingCartDetail) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *ShoppingCartDetail) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetOwner

`func (o *ShoppingCartDetail) GetOwner() Person`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ShoppingCartDetail) GetOwnerOk() (*Person, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ShoppingCartDetail) SetOwner(v Person)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ShoppingCartDetail) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSharedWith

`func (o *ShoppingCartDetail) GetSharedWith() []Person`

GetSharedWith returns the SharedWith field if non-nil, zero value otherwise.

### GetSharedWithOk

`func (o *ShoppingCartDetail) GetSharedWithOk() (*[]Person, bool)`

GetSharedWithOk returns a tuple with the SharedWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedWith

`func (o *ShoppingCartDetail) SetSharedWith(v []Person)`

SetSharedWith sets SharedWith field to given value.

### HasSharedWith

`func (o *ShoppingCartDetail) HasSharedWith() bool`

HasSharedWith returns a boolean if a field has been set.

### GetPersons

`func (o *ShoppingCartDetail) GetPersons() []Person`

GetPersons returns the Persons field if non-nil, zero value otherwise.

### GetPersonsOk

`func (o *ShoppingCartDetail) GetPersonsOk() (*[]Person, bool)`

GetPersonsOk returns a tuple with the Persons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersons

`func (o *ShoppingCartDetail) SetPersons(v []Person)`

SetPersons sets Persons field to given value.

### HasPersons

`func (o *ShoppingCartDetail) HasPersons() bool`

HasPersons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


