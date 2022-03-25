# IndutryDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Industry** | Pointer to [**Industry**](Industry.md) |  | [optional] 
**Organizations** | Pointer to [**[]Organization**](Organization.md) |  | [optional] 

## Methods

### NewIndutryDetails

`func NewIndutryDetails() *IndutryDetails`

NewIndutryDetails instantiates a new IndutryDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndutryDetailsWithDefaults

`func NewIndutryDetailsWithDefaults() *IndutryDetails`

NewIndutryDetailsWithDefaults instantiates a new IndutryDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndustry

`func (o *IndutryDetails) GetIndustry() Industry`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *IndutryDetails) GetIndustryOk() (*Industry, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *IndutryDetails) SetIndustry(v Industry)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *IndutryDetails) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetOrganizations

`func (o *IndutryDetails) GetOrganizations() []Organization`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *IndutryDetails) GetOrganizationsOk() (*[]Organization, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *IndutryDetails) SetOrganizations(v []Organization)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *IndutryDetails) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


