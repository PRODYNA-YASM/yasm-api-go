# ProjectScoreDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | Pointer to **float32** |  | [optional] 
**DirectHit** | Pointer to **bool** | Indicates if the score is a direct hit | [optional] 
**Scores** | Pointer to [**[]Score**](Score.md) |  | [optional] 

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

### GetDirectHit

`func (o *ProjectScoreDetail) GetDirectHit() bool`

GetDirectHit returns the DirectHit field if non-nil, zero value otherwise.

### GetDirectHitOk

`func (o *ProjectScoreDetail) GetDirectHitOk() (*bool, bool)`

GetDirectHitOk returns a tuple with the DirectHit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectHit

`func (o *ProjectScoreDetail) SetDirectHit(v bool)`

SetDirectHit sets DirectHit field to given value.

### HasDirectHit

`func (o *ProjectScoreDetail) HasDirectHit() bool`

HasDirectHit returns a boolean if a field has been set.

### GetScores

`func (o *ProjectScoreDetail) GetScores() []Score`

GetScores returns the Scores field if non-nil, zero value otherwise.

### GetScoresOk

`func (o *ProjectScoreDetail) GetScoresOk() (*[]Score, bool)`

GetScoresOk returns a tuple with the Scores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScores

`func (o *ProjectScoreDetail) SetScores(v []Score)`

SetScores sets Scores field to given value.

### HasScores

`func (o *ProjectScoreDetail) HasScores() bool`

HasScores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


