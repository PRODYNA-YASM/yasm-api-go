# Skill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Suggestion** | Pointer to **bool** |  | [optional] [default to false]
**Synonyms** | **[]string** |  | 
**Invest** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSkill

`func NewSkill(id string, name string, synonyms []string, ) *Skill`

NewSkill instantiates a new Skill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillWithDefaults

`func NewSkillWithDefaults() *Skill`

NewSkillWithDefaults instantiates a new Skill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Skill) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Skill) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Skill) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Skill) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Skill) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Skill) SetName(v string)`

SetName sets Name field to given value.


### GetSuggestion

`func (o *Skill) GetSuggestion() bool`

GetSuggestion returns the Suggestion field if non-nil, zero value otherwise.

### GetSuggestionOk

`func (o *Skill) GetSuggestionOk() (*bool, bool)`

GetSuggestionOk returns a tuple with the Suggestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestion

`func (o *Skill) SetSuggestion(v bool)`

SetSuggestion sets Suggestion field to given value.

### HasSuggestion

`func (o *Skill) HasSuggestion() bool`

HasSuggestion returns a boolean if a field has been set.

### GetSynonyms

`func (o *Skill) GetSynonyms() []string`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *Skill) GetSynonymsOk() (*[]string, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *Skill) SetSynonyms(v []string)`

SetSynonyms sets Synonyms field to given value.


### GetInvest

`func (o *Skill) GetInvest() bool`

GetInvest returns the Invest field if non-nil, zero value otherwise.

### GetInvestOk

`func (o *Skill) GetInvestOk() (*bool, bool)`

GetInvestOk returns a tuple with the Invest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvest

`func (o *Skill) SetInvest(v bool)`

SetInvest sets Invest field to given value.

### HasInvest

`func (o *Skill) HasInvest() bool`

HasInvest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


