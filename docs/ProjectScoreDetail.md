# ProjectScoreDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 
**Persons** | Pointer to [**[]Person**](Person.md) |  | [optional] 
**Score** | Pointer to **float32** |  | [optional] 

## Methods

### NewProjectScoreDetail

`func NewProjectScoreDetail() *ProjectScoreDetail`

NewProjectScoreDetail instantiates a new ProjectScoreDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectScoreDetailWithDefaults

`func NewProjectScoreDetailWithDefaults() *ProjectScoreDetail`

NewProjectScoreDetailWithDefaults instantiates a new ProjectScoreDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *ProjectScoreDetail) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ProjectScoreDetail) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ProjectScoreDetail) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *ProjectScoreDetail) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetOrganization

`func (o *ProjectScoreDetail) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ProjectScoreDetail) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ProjectScoreDetail) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ProjectScoreDetail) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPersons

`func (o *ProjectScoreDetail) GetPersons() []Person`

GetPersons returns the Persons field if non-nil, zero value otherwise.

### GetPersonsOk

`func (o *ProjectScoreDetail) GetPersonsOk() (*[]Person, bool)`

GetPersonsOk returns a tuple with the Persons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersons

`func (o *ProjectScoreDetail) SetPersons(v []Person)`

SetPersons sets Persons field to given value.

### HasPersons

`func (o *ProjectScoreDetail) HasPersons() bool`

HasPersons returns a boolean if a field has been set.

### GetScore

`func (o *ProjectScoreDetail) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ProjectScoreDetail) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ProjectScoreDetail) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *ProjectScoreDetail) HasScore() bool`

HasScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


