# CertificationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certification** | Pointer to [**Certification**](Certification.md) |  | [optional] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 
**Skills** | Pointer to [**[]SkillLevel**](SkillLevel.md) |  | [optional] 

## Methods

### NewCertificationDetails

`func NewCertificationDetails() *CertificationDetails`

NewCertificationDetails instantiates a new CertificationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificationDetailsWithDefaults

`func NewCertificationDetailsWithDefaults() *CertificationDetails`

NewCertificationDetailsWithDefaults instantiates a new CertificationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertification

`func (o *CertificationDetails) GetCertification() Certification`

GetCertification returns the Certification field if non-nil, zero value otherwise.

### GetCertificationOk

`func (o *CertificationDetails) GetCertificationOk() (*Certification, bool)`

GetCertificationOk returns a tuple with the Certification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertification

`func (o *CertificationDetails) SetCertification(v Certification)`

SetCertification sets Certification field to given value.

### HasCertification

`func (o *CertificationDetails) HasCertification() bool`

HasCertification returns a boolean if a field has been set.

### GetOrganization

`func (o *CertificationDetails) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CertificationDetails) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CertificationDetails) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *CertificationDetails) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSkills

`func (o *CertificationDetails) GetSkills() []SkillLevel`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *CertificationDetails) GetSkillsOk() (*[]SkillLevel, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *CertificationDetails) SetSkills(v []SkillLevel)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *CertificationDetails) HasSkills() bool`

HasSkills returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


