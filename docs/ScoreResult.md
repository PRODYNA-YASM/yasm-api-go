# ScoreResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | Pointer to **float32** |  | [optional] 
**DirectHit** | Pointer to **bool** |  | [optional] 

## Methods

### NewScoreResult

`func NewScoreResult() *ScoreResult`

NewScoreResult instantiates a new ScoreResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScoreResultWithDefaults

`func NewScoreResultWithDefaults() *ScoreResult`

NewScoreResultWithDefaults instantiates a new ScoreResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *ScoreResult) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ScoreResult) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ScoreResult) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *ScoreResult) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetDirectHit

`func (o *ScoreResult) GetDirectHit() bool`

GetDirectHit returns the DirectHit field if non-nil, zero value otherwise.

### GetDirectHitOk

`func (o *ScoreResult) GetDirectHitOk() (*bool, bool)`

GetDirectHitOk returns a tuple with the DirectHit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectHit

`func (o *ScoreResult) SetDirectHit(v bool)`

SetDirectHit sets DirectHit field to given value.

### HasDirectHit

`func (o *ScoreResult) HasDirectHit() bool`

HasDirectHit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


