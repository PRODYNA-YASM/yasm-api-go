# Score

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | Pointer to **float32** |  | [optional] 
**ScoreType** | Pointer to **string** | Specifies the type of score | [optional] [default to "skillScore"]

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

### GetScoreType

`func (o *Score) GetScoreType() string`

GetScoreType returns the ScoreType field if non-nil, zero value otherwise.

### GetScoreTypeOk

`func (o *Score) GetScoreTypeOk() (*string, bool)`

GetScoreTypeOk returns a tuple with the ScoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreType

`func (o *Score) SetScoreType(v string)`

SetScoreType sets ScoreType field to given value.

### HasScoreType

`func (o *Score) HasScoreType() bool`

HasScoreType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


