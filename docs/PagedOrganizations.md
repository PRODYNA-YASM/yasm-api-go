# PagedOrganizations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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


