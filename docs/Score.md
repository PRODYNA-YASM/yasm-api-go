# Score

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | Pointer to **float32** |  | [optional] 
**DirectHit** | Pointer to **bool** | Indicates if the score is a direct hit | [optional] 
**BestMatch** | Pointer to **bool** | Indicates if the score is the best match | [optional] 
**InputEntity** | Pointer to [**NamedDomainModel**](NamedDomainModel.md) |  | [optional] 
**MatchingEntity** | Pointer to [**[]NamedDomainModel**](NamedDomainModel.md) | The matching entities | [optional] 

## Methods

### NewScore

`func NewScore() *Score`

NewScore instantiates a new Score object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScoreWithDefaults

`func NewScoreWithDefaults() *Score`

NewScoreWithDefaults instantiates a new Score object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *Score) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *Score) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *Score) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *Score) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetDirectHit

`func (o *Score) GetDirectHit() bool`

GetDirectHit returns the DirectHit field if non-nil, zero value otherwise.

### GetDirectHitOk

`func (o *Score) GetDirectHitOk() (*bool, bool)`

GetDirectHitOk returns a tuple with the DirectHit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectHit

`func (o *Score) SetDirectHit(v bool)`

SetDirectHit sets DirectHit field to given value.

### HasDirectHit

`func (o *Score) HasDirectHit() bool`

HasDirectHit returns a boolean if a field has been set.

### GetBestMatch

`func (o *Score) GetBestMatch() bool`

GetBestMatch returns the BestMatch field if non-nil, zero value otherwise.

### GetBestMatchOk

`func (o *Score) GetBestMatchOk() (*bool, bool)`

GetBestMatchOk returns a tuple with the BestMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestMatch

`func (o *Score) SetBestMatch(v bool)`

SetBestMatch sets BestMatch field to given value.

### HasBestMatch

`func (o *Score) HasBestMatch() bool`

HasBestMatch returns a boolean if a field has been set.

### GetInputEntity

`func (o *Score) GetInputEntity() NamedDomainModel`

GetInputEntity returns the InputEntity field if non-nil, zero value otherwise.

### GetInputEntityOk

`func (o *Score) GetInputEntityOk() (*NamedDomainModel, bool)`

GetInputEntityOk returns a tuple with the InputEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputEntity

`func (o *Score) SetInputEntity(v NamedDomainModel)`

SetInputEntity sets InputEntity field to given value.

### HasInputEntity

`func (o *Score) HasInputEntity() bool`

HasInputEntity returns a boolean if a field has been set.

### GetMatchingEntity

`func (o *Score) GetMatchingEntity() []NamedDomainModel`

GetMatchingEntity returns the MatchingEntity field if non-nil, zero value otherwise.

### GetMatchingEntityOk

`func (o *Score) GetMatchingEntityOk() (*[]NamedDomainModel, bool)`

GetMatchingEntityOk returns a tuple with the MatchingEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingEntity

`func (o *Score) SetMatchingEntity(v []NamedDomainModel)`

SetMatchingEntity sets MatchingEntity field to given value.

### HasMatchingEntity

`func (o *Score) HasMatchingEntity() bool`

HasMatchingEntity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


