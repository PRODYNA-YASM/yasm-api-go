# PersonSearchOrganizationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**AmountOfProjects** | Pointer to [**MinMax**](MinMax.md) |  | [optional] 

## Methods

### NewPersonSearchOrganizationsInner

`func NewPersonSearchOrganizationsInner(id string, ) *PersonSearchOrganizationsInner`

NewPersonSearchOrganizationsInner instantiates a new PersonSearchOrganizationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonSearchOrganizationsInnerWithDefaults

`func NewPersonSearchOrganizationsInnerWithDefaults() *PersonSearchOrganizationsInner`

NewPersonSearchOrganizationsInnerWithDefaults instantiates a new PersonSearchOrganizationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PersonSearchOrganizationsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PersonSearchOrganizationsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PersonSearchOrganizationsInner) SetId(v string)`

SetId sets Id field to given value.


### GetAmountOfProjects

`func (o *PersonSearchOrganizationsInner) GetAmountOfProjects() MinMax`

GetAmountOfProjects returns the AmountOfProjects field if non-nil, zero value otherwise.

### GetAmountOfProjectsOk

`func (o *PersonSearchOrganizationsInner) GetAmountOfProjectsOk() (*MinMax, bool)`

GetAmountOfProjectsOk returns a tuple with the AmountOfProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOfProjects

`func (o *PersonSearchOrganizationsInner) SetAmountOfProjects(v MinMax)`

SetAmountOfProjects sets AmountOfProjects field to given value.

### HasAmountOfProjects

`func (o *PersonSearchOrganizationsInner) HasAmountOfProjects() bool`

HasAmountOfProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


