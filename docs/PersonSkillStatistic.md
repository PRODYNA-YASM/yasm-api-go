# PersonSkillStatistic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skill** | Pointer to [**Skill**](Skill.md) |  | [optional] 
**Duration** | Pointer to **string** |  | [optional] 
**Interest** | Pointer to **bool** |  | [optional] 
**UsedInProjects** | Pointer to [**[]SkillStatisticProject**](SkillStatisticProject.md) |  | [optional] 
**TrainedByCertifications** | Pointer to [**[]CertificationView**](CertificationView.md) |  | [optional] 

## Methods

### NewPersonSkillStatistic

`func NewPersonSkillStatistic() *PersonSkillStatistic`

NewPersonSkillStatistic instantiates a new PersonSkillStatistic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonSkillStatisticWithDefaults

`func NewPersonSkillStatisticWithDefaults() *PersonSkillStatistic`

NewPersonSkillStatisticWithDefaults instantiates a new PersonSkillStatistic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkill

`func (o *PersonSkillStatistic) GetSkill() Skill`

GetSkill returns the Skill field if non-nil, zero value otherwise.

### GetSkillOk

`func (o *PersonSkillStatistic) GetSkillOk() (*Skill, bool)`

GetSkillOk returns a tuple with the Skill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkill

`func (o *PersonSkillStatistic) SetSkill(v Skill)`

SetSkill sets Skill field to given value.

### HasSkill

`func (o *PersonSkillStatistic) HasSkill() bool`

HasSkill returns a boolean if a field has been set.

### GetDuration

`func (o *PersonSkillStatistic) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *PersonSkillStatistic) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *PersonSkillStatistic) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *PersonSkillStatistic) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetInterest

`func (o *PersonSkillStatistic) GetInterest() bool`

GetInterest returns the Interest field if non-nil, zero value otherwise.

### GetInterestOk

`func (o *PersonSkillStatistic) GetInterestOk() (*bool, bool)`

GetInterestOk returns a tuple with the Interest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterest

`func (o *PersonSkillStatistic) SetInterest(v bool)`

SetInterest sets Interest field to given value.

### HasInterest

`func (o *PersonSkillStatistic) HasInterest() bool`

HasInterest returns a boolean if a field has been set.

### GetUsedInProjects

`func (o *PersonSkillStatistic) GetUsedInProjects() []SkillStatisticProject`

GetUsedInProjects returns the UsedInProjects field if non-nil, zero value otherwise.

### GetUsedInProjectsOk

`func (o *PersonSkillStatistic) GetUsedInProjectsOk() (*[]SkillStatisticProject, bool)`

GetUsedInProjectsOk returns a tuple with the UsedInProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedInProjects

`func (o *PersonSkillStatistic) SetUsedInProjects(v []SkillStatisticProject)`

SetUsedInProjects sets UsedInProjects field to given value.

### HasUsedInProjects

`func (o *PersonSkillStatistic) HasUsedInProjects() bool`

HasUsedInProjects returns a boolean if a field has been set.

### GetTrainedByCertifications

`func (o *PersonSkillStatistic) GetTrainedByCertifications() []CertificationView`

GetTrainedByCertifications returns the TrainedByCertifications field if non-nil, zero value otherwise.

### GetTrainedByCertificationsOk

`func (o *PersonSkillStatistic) GetTrainedByCertificationsOk() (*[]CertificationView, bool)`

GetTrainedByCertificationsOk returns a tuple with the TrainedByCertifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainedByCertifications

`func (o *PersonSkillStatistic) SetTrainedByCertifications(v []CertificationView)`

SetTrainedByCertifications sets TrainedByCertifications field to given value.

### HasTrainedByCertifications

`func (o *PersonSkillStatistic) HasTrainedByCertifications() bool`

HasTrainedByCertifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


