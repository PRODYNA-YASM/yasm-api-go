# ProjectDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 
**Industries** | Pointer to [**[]Industry**](Industry.md) |  | [optional] 
**Persons** | Pointer to [**[]Person**](Person.md) |  | [optional] 

## Methods

### NewProjectDetails

`func NewProjectDetails() *ProjectDetails`

NewProjectDetails instantiates a new ProjectDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDetailsWithDefaults

`func NewProjectDetailsWithDefaults() *ProjectDetails`

NewProjectDetailsWithDefaults instantiates a new ProjectDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *ProjectDetails) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ProjectDetails) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ProjectDetails) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *ProjectDetails) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetOrganization

`func (o *ProjectDetails) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ProjectDetails) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ProjectDetails) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ProjectDetails) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetIndustries

`func (o *ProjectDetails) GetIndustries() []Industry`

GetIndustries returns the Industries field if non-nil, zero value otherwise.

### GetIndustriesOk

`func (o *ProjectDetails) GetIndustriesOk() (*[]Industry, bool)`

GetIndustriesOk returns a tuple with the Industries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustries

`func (o *ProjectDetails) SetIndustries(v []Industry)`

SetIndustries sets Industries field to given value.

### HasIndustries

`func (o *ProjectDetails) HasIndustries() bool`

HasIndustries returns a boolean if a field has been set.

### GetPersons

`func (o *ProjectDetails) GetPersons() []Person`

GetPersons returns the Persons field if non-nil, zero value otherwise.

### GetPersonsOk

`func (o *ProjectDetails) GetPersonsOk() (*[]Person, bool)`

GetPersonsOk returns a tuple with the Persons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersons

`func (o *ProjectDetails) SetPersons(v []Person)`

SetPersons sets Persons field to given value.

### HasPersons

`func (o *ProjectDetails) HasPersons() bool`

HasPersons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


