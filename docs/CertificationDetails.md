# CertificationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**Certification** | Pointer to [**Certification**](Certification.md) |  | [optional] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 
**SkillLevels** | Pointer to [**[]SkillLevel**](SkillLevel.md) |  | [optional] 

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

### GetAudit

`func (o *CertificationDetails) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *CertificationDetails) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *CertificationDetails) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *CertificationDetails) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

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

### GetSkillLevels

`func (o *CertificationDetails) GetSkillLevels() []SkillLevel`

GetSkillLevels returns the SkillLevels field if non-nil, zero value otherwise.

### GetSkillLevelsOk

`func (o *CertificationDetails) GetSkillLevelsOk() (*[]SkillLevel, bool)`

GetSkillLevelsOk returns a tuple with the SkillLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillLevels

`func (o *CertificationDetails) SetSkillLevels(v []SkillLevel)`

SetSkillLevels sets SkillLevels field to given value.

### HasSkillLevels

`func (o *CertificationDetails) HasSkillLevels() bool`

HasSkillLevels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


