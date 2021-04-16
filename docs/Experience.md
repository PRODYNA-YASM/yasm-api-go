# Experience

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Skill** | Pointer to [**SkillLevel**](SkillLevel.md) |  | [optional] 
**ConfirmedBy** | Pointer to [**[]Person**](Person.md) |  | [optional] 

## Methods

### NewExperience

`func NewExperience(id string, name string, ) *Experience`

NewExperience instantiates a new Experience object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperienceWithDefaults

`func NewExperienceWithDefaults() *Experience`

NewExperienceWithDefaults instantiates a new Experience object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Experience) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Experience) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Experience) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Experience) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Experience) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Experience) SetName(v string)`

SetName sets Name field to given value.


### GetSkill

`func (o *Experience) GetSkill() SkillLevel`

GetSkill returns the Skill field if non-nil, zero value otherwise.

### GetSkillOk

`func (o *Experience) GetSkillOk() (*SkillLevel, bool)`

GetSkillOk returns a tuple with the Skill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkill

`func (o *Experience) SetSkill(v SkillLevel)`

SetSkill sets Skill field to given value.

### HasSkill

`func (o *Experience) HasSkill() bool`

HasSkill returns a boolean if a field has been set.

### GetConfirmedBy

`func (o *Experience) GetConfirmedBy() []Person`

GetConfirmedBy returns the ConfirmedBy field if non-nil, zero value otherwise.

### GetConfirmedByOk

`func (o *Experience) GetConfirmedByOk() (*[]Person, bool)`

GetConfirmedByOk returns a tuple with the ConfirmedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedBy

`func (o *Experience) SetConfirmedBy(v []Person)`

SetConfirmedBy sets ConfirmedBy field to given value.

### HasConfirmedBy

`func (o *Experience) HasConfirmedBy() bool`

HasConfirmedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


