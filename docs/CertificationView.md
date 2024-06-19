# CertificationView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certification** | Pointer to [**Certification**](Certification.md) |  | [optional] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 

## Methods

### NewCertificationView

`func NewCertificationView() *CertificationView`

NewCertificationView instantiates a new CertificationView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificationViewWithDefaults

`func NewCertificationViewWithDefaults() *CertificationView`

NewCertificationViewWithDefaults instantiates a new CertificationView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertification

`func (o *CertificationView) GetCertification() Certification`

GetCertification returns the Certification field if non-nil, zero value otherwise.

### GetCertificationOk

`func (o *CertificationView) GetCertificationOk() (*Certification, bool)`

GetCertificationOk returns a tuple with the Certification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertification

`func (o *CertificationView) SetCertification(v Certification)`

SetCertification sets Certification field to given value.

### HasCertification

`func (o *CertificationView) HasCertification() bool`

HasCertification returns a boolean if a field has been set.

### GetOrganization

`func (o *CertificationView) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CertificationView) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CertificationView) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *CertificationView) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


