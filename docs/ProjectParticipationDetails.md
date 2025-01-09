# ProjectParticipationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**Participation** | Pointer to [**ProjectParticipation**](ProjectParticipation.md) |  | [optional] 
**SkillGroups** | Pointer to [**[]ExperienceSkillGroup**](ExperienceSkillGroup.md) |  | [optional] 
**Person** | Pointer to [**Person**](Person.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Awards** | Pointer to [**[]Award**](Award.md) |  | [optional] 

## Methods

### NewProjectParticipationDetails

`func NewProjectParticipationDetails() *ProjectParticipationDetails`

NewProjectParticipationDetails instantiates a new ProjectParticipationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectParticipationDetailsWithDefaults

`func NewProjectParticipationDetailsWithDefaults() *ProjectParticipationDetails`

NewProjectParticipationDetailsWithDefaults instantiates a new ProjectParticipationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *ProjectParticipationDetails) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *ProjectParticipationDetails) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *ProjectParticipationDetails) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *ProjectParticipationDetails) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetParticipation

`func (o *ProjectParticipationDetails) GetParticipation() ProjectParticipation`

GetParticipation returns the Participation field if non-nil, zero value otherwise.

### GetParticipationOk

`func (o *ProjectParticipationDetails) GetParticipationOk() (*ProjectParticipation, bool)`

GetParticipationOk returns a tuple with the Participation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipation

`func (o *ProjectParticipationDetails) SetParticipation(v ProjectParticipation)`

SetParticipation sets Participation field to given value.

### HasParticipation

`func (o *ProjectParticipationDetails) HasParticipation() bool`

HasParticipation returns a boolean if a field has been set.

### GetSkillGroups

`func (o *ProjectParticipationDetails) GetSkillGroups() []ExperienceSkillGroup`

GetSkillGroups returns the SkillGroups field if non-nil, zero value otherwise.

### GetSkillGroupsOk

`func (o *ProjectParticipationDetails) GetSkillGroupsOk() (*[]ExperienceSkillGroup, bool)`

GetSkillGroupsOk returns a tuple with the SkillGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillGroups

`func (o *ProjectParticipationDetails) SetSkillGroups(v []ExperienceSkillGroup)`

SetSkillGroups sets SkillGroups field to given value.

### HasSkillGroups

`func (o *ProjectParticipationDetails) HasSkillGroups() bool`

HasSkillGroups returns a boolean if a field has been set.

### GetPerson

`func (o *ProjectParticipationDetails) GetPerson() Person`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *ProjectParticipationDetails) GetPersonOk() (*Person, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *ProjectParticipationDetails) SetPerson(v Person)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *ProjectParticipationDetails) HasPerson() bool`

HasPerson returns a boolean if a field has been set.

### GetProject

`func (o *ProjectParticipationDetails) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ProjectParticipationDetails) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ProjectParticipationDetails) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *ProjectParticipationDetails) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetAwards

`func (o *ProjectParticipationDetails) GetAwards() []Award`

GetAwards returns the Awards field if non-nil, zero value otherwise.

### GetAwardsOk

`func (o *ProjectParticipationDetails) GetAwardsOk() (*[]Award, bool)`

GetAwardsOk returns a tuple with the Awards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwards

`func (o *ProjectParticipationDetails) SetAwards(v []Award)`

SetAwards sets Awards field to given value.

### HasAwards

`func (o *ProjectParticipationDetails) HasAwards() bool`

HasAwards returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


