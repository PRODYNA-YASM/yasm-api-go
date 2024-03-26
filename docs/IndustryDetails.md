# IndustryDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Industry** | Pointer to [**Industry**](Industry.md) |  | [optional] 
**OrganizationsCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewIndustryDetails

`func NewIndustryDetails() *IndustryDetails`

NewIndustryDetails instantiates a new IndustryDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndustryDetailsWithDefaults

`func NewIndustryDetailsWithDefaults() *IndustryDetails`

NewIndustryDetailsWithDefaults instantiates a new IndustryDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndustry

`func (o *IndustryDetails) GetIndustry() Industry`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *IndustryDetails) GetIndustryOk() (*Industry, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *IndustryDetails) SetIndustry(v Industry)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *IndustryDetails) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetOrganizationsCount

`func (o *IndustryDetails) GetOrganizationsCount() int32`

GetOrganizationsCount returns the OrganizationsCount field if non-nil, zero value otherwise.

### GetOrganizationsCountOk

`func (o *IndustryDetails) GetOrganizationsCountOk() (*int32, bool)`

GetOrganizationsCountOk returns a tuple with the OrganizationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationsCount

`func (o *IndustryDetails) SetOrganizationsCount(v int32)`

SetOrganizationsCount sets OrganizationsCount field to given value.

### HasOrganizationsCount

`func (o *IndustryDetails) HasOrganizationsCount() bool`

HasOrganizationsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


