# OrganizationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 
**ServiceManager** | Pointer to [**Person**](Person.md) |  | [optional] 
**Children** | Pointer to [**[]Organization**](Organization.md) |  | [optional] 
**Parents** | Pointer to [**[]Organization**](Organization.md) |  | [optional] 
**Projects** | Pointer to [**[]Project**](Project.md) |  | [optional] [readonly] 
**Industries** | Pointer to [**[]Industry**](Industry.md) |  | [optional] 
**Certifications** | Pointer to [**[]Certification**](Certification.md) |  | [optional] 
**Offices** | Pointer to [**[]Office**](Office.md) |  | [optional] 

## Methods

### NewOrganizationDetails

`func NewOrganizationDetails() *OrganizationDetails`

NewOrganizationDetails instantiates a new OrganizationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationDetailsWithDefaults

`func NewOrganizationDetailsWithDefaults() *OrganizationDetails`

NewOrganizationDetailsWithDefaults instantiates a new OrganizationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *OrganizationDetails) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *OrganizationDetails) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *OrganizationDetails) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *OrganizationDetails) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetOrganization

`func (o *OrganizationDetails) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OrganizationDetails) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OrganizationDetails) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *OrganizationDetails) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetServiceManager

`func (o *OrganizationDetails) GetServiceManager() Person`

GetServiceManager returns the ServiceManager field if non-nil, zero value otherwise.

### GetServiceManagerOk

`func (o *OrganizationDetails) GetServiceManagerOk() (*Person, bool)`

GetServiceManagerOk returns a tuple with the ServiceManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceManager

`func (o *OrganizationDetails) SetServiceManager(v Person)`

SetServiceManager sets ServiceManager field to given value.

### HasServiceManager

`func (o *OrganizationDetails) HasServiceManager() bool`

HasServiceManager returns a boolean if a field has been set.

### GetChildren

`func (o *OrganizationDetails) GetChildren() []Organization`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *OrganizationDetails) GetChildrenOk() (*[]Organization, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *OrganizationDetails) SetChildren(v []Organization)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *OrganizationDetails) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetParents

`func (o *OrganizationDetails) GetParents() []Organization`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *OrganizationDetails) GetParentsOk() (*[]Organization, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *OrganizationDetails) SetParents(v []Organization)`

SetParents sets Parents field to given value.

### HasParents

`func (o *OrganizationDetails) HasParents() bool`

HasParents returns a boolean if a field has been set.

### GetProjects

`func (o *OrganizationDetails) GetProjects() []Project`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *OrganizationDetails) GetProjectsOk() (*[]Project, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *OrganizationDetails) SetProjects(v []Project)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *OrganizationDetails) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetIndustries

`func (o *OrganizationDetails) GetIndustries() []Industry`

GetIndustries returns the Industries field if non-nil, zero value otherwise.

### GetIndustriesOk

`func (o *OrganizationDetails) GetIndustriesOk() (*[]Industry, bool)`

GetIndustriesOk returns a tuple with the Industries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustries

`func (o *OrganizationDetails) SetIndustries(v []Industry)`

SetIndustries sets Industries field to given value.

### HasIndustries

`func (o *OrganizationDetails) HasIndustries() bool`

HasIndustries returns a boolean if a field has been set.

### GetCertifications

`func (o *OrganizationDetails) GetCertifications() []Certification`

GetCertifications returns the Certifications field if non-nil, zero value otherwise.

### GetCertificationsOk

`func (o *OrganizationDetails) GetCertificationsOk() (*[]Certification, bool)`

GetCertificationsOk returns a tuple with the Certifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifications

`func (o *OrganizationDetails) SetCertifications(v []Certification)`

SetCertifications sets Certifications field to given value.

### HasCertifications

`func (o *OrganizationDetails) HasCertifications() bool`

HasCertifications returns a boolean if a field has been set.

### GetOffices

`func (o *OrganizationDetails) GetOffices() []Office`

GetOffices returns the Offices field if non-nil, zero value otherwise.

### GetOfficesOk

`func (o *OrganizationDetails) GetOfficesOk() (*[]Office, bool)`

GetOfficesOk returns a tuple with the Offices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffices

`func (o *OrganizationDetails) SetOffices(v []Office)`

SetOffices sets Offices field to given value.

### HasOffices

`func (o *OrganizationDetails) HasOffices() bool`

HasOffices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


