# ProjectView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 
**Timeframe** | Pointer to [**Timeframed**](Timeframed.md) |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 

## Methods

### NewProjectView

`func NewProjectView() *ProjectView`

NewProjectView instantiates a new ProjectView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectViewWithDefaults

`func NewProjectViewWithDefaults() *ProjectView`

NewProjectViewWithDefaults instantiates a new ProjectView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *ProjectView) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ProjectView) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ProjectView) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *ProjectView) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetOrganization

`func (o *ProjectView) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ProjectView) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ProjectView) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ProjectView) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetTimeframe

`func (o *ProjectView) GetTimeframe() Timeframed`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *ProjectView) GetTimeframeOk() (*Timeframed, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *ProjectView) SetTimeframe(v Timeframed)`

SetTimeframe sets Timeframe field to given value.

### HasTimeframe

`func (o *ProjectView) HasTimeframe() bool`

HasTimeframe returns a boolean if a field has been set.

### GetObjectType

`func (o *ProjectView) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ProjectView) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ProjectView) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *ProjectView) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


