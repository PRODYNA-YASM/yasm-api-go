# PersonOrganizationFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountOfProjects** | Pointer to [**MinMax**](MinMax.md) |  | [optional] 
**Id** | **string** |  | 

## Methods

### NewPersonOrganizationFilter

`func NewPersonOrganizationFilter(id string, ) *PersonOrganizationFilter`

NewPersonOrganizationFilter instantiates a new PersonOrganizationFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonOrganizationFilterWithDefaults

`func NewPersonOrganizationFilterWithDefaults() *PersonOrganizationFilter`

NewPersonOrganizationFilterWithDefaults instantiates a new PersonOrganizationFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountOfProjects

`func (o *PersonOrganizationFilter) GetAmountOfProjects() MinMax`

GetAmountOfProjects returns the AmountOfProjects field if non-nil, zero value otherwise.

### GetAmountOfProjectsOk

`func (o *PersonOrganizationFilter) GetAmountOfProjectsOk() (*MinMax, bool)`

GetAmountOfProjectsOk returns a tuple with the AmountOfProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOfProjects

`func (o *PersonOrganizationFilter) SetAmountOfProjects(v MinMax)`

SetAmountOfProjects sets AmountOfProjects field to given value.

### HasAmountOfProjects

`func (o *PersonOrganizationFilter) HasAmountOfProjects() bool`

HasAmountOfProjects returns a boolean if a field has been set.

### GetId

`func (o *PersonOrganizationFilter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PersonOrganizationFilter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PersonOrganizationFilter) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


