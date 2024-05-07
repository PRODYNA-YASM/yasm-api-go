# SkillSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Types** | Pointer to **string** | Gives you either all skills, only the root kills | [optional] [default to "all"]
**SkillIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSkillSearch

`func NewSkillSearch() *SkillSearch`

NewSkillSearch instantiates a new SkillSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillSearchWithDefaults

`func NewSkillSearchWithDefaults() *SkillSearch`

NewSkillSearchWithDefaults instantiates a new SkillSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypes

`func (o *SkillSearch) GetTypes() string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *SkillSearch) GetTypesOk() (*string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *SkillSearch) SetTypes(v string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *SkillSearch) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetSkillIds

`func (o *SkillSearch) GetSkillIds() []string`

GetSkillIds returns the SkillIds field if non-nil, zero value otherwise.

### GetSkillIdsOk

`func (o *SkillSearch) GetSkillIdsOk() (*[]string, bool)`

GetSkillIdsOk returns a tuple with the SkillIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillIds

`func (o *SkillSearch) SetSkillIds(v []string)`

SetSkillIds sets SkillIds field to given value.

### HasSkillIds

`func (o *SkillSearch) HasSkillIds() bool`

HasSkillIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


