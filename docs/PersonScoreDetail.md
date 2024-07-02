# PersonScoreDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectHit** | Pointer to **bool** |  | [optional] 
**Scores** | Pointer to [**[]Score**](Score.md) |  | [optional] 

## Methods

### NewPersonScoreDetail

`func NewPersonScoreDetail() *PersonScoreDetail`

NewPersonScoreDetail instantiates a new PersonScoreDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonScoreDetailWithDefaults

`func NewPersonScoreDetailWithDefaults() *PersonScoreDetail`

NewPersonScoreDetailWithDefaults instantiates a new PersonScoreDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectHit

`func (o *PersonScoreDetail) GetDirectHit() bool`

GetDirectHit returns the DirectHit field if non-nil, zero value otherwise.

### GetDirectHitOk

`func (o *PersonScoreDetail) GetDirectHitOk() (*bool, bool)`

GetDirectHitOk returns a tuple with the DirectHit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectHit

`func (o *PersonScoreDetail) SetDirectHit(v bool)`

SetDirectHit sets DirectHit field to given value.

### HasDirectHit

`func (o *PersonScoreDetail) HasDirectHit() bool`

HasDirectHit returns a boolean if a field has been set.

### GetScores

`func (o *PersonScoreDetail) GetScores() []Score`

GetScores returns the Scores field if non-nil, zero value otherwise.

### GetScoresOk

`func (o *PersonScoreDetail) GetScoresOk() (*[]Score, bool)`

GetScoresOk returns a tuple with the Scores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScores

`func (o *PersonScoreDetail) SetScores(v []Score)`

SetScores sets Scores field to given value.

### HasScores

`func (o *PersonScoreDetail) HasScores() bool`

HasScores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


