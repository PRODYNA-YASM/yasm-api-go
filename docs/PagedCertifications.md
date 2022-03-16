# PagedCertifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Certifications** | Pointer to [**[]CertificationDetails**](CertificationDetails.md) |  | [optional] 

## Methods

### NewPagedCertifications

`func NewPagedCertifications() *PagedCertifications`

NewPagedCertifications instantiates a new PagedCertifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagedCertificationsWithDefaults

`func NewPagedCertificationsWithDefaults() *PagedCertifications`

NewPagedCertificationsWithDefaults instantiates a new PagedCertifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkip

`func (o *PagedCertifications) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *PagedCertifications) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *PagedCertifications) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *PagedCertifications) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetLimit

`func (o *PagedCertifications) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PagedCertifications) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PagedCertifications) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PagedCertifications) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetCount

`func (o *PagedCertifications) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PagedCertifications) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PagedCertifications) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PagedCertifications) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCertifications

`func (o *PagedCertifications) GetCertifications() []CertificationDetails`

GetCertifications returns the Certifications field if non-nil, zero value otherwise.

### GetCertificationsOk

`func (o *PagedCertifications) GetCertificationsOk() (*[]CertificationDetails, bool)`

GetCertificationsOk returns a tuple with the Certifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifications

`func (o *PagedCertifications) SetCertifications(v []CertificationDetails)`

SetCertifications sets Certifications field to given value.

### HasCertifications

`func (o *PagedCertifications) HasCertifications() bool`

HasCertifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


