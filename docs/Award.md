# Award

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Synonyms** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**YearOfAward** | Pointer to **NullableInt32** | The year in which the award was received | [optional] 

## Methods

### NewAward

`func NewAward() *Award`

NewAward instantiates a new Award object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwardWithDefaults

`func NewAwardWithDefaults() *Award`

NewAwardWithDefaults instantiates a new Award object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSynonyms

`func (o *Award) GetSynonyms() []string`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *Award) GetSynonymsOk() (*[]string, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *Award) SetSynonyms(v []string)`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *Award) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.

### GetDescription

`func (o *Award) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Award) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Award) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Award) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetYearOfAward

`func (o *Award) GetYearOfAward() int32`

GetYearOfAward returns the YearOfAward field if non-nil, zero value otherwise.

### GetYearOfAwardOk

`func (o *Award) GetYearOfAwardOk() (*int32, bool)`

GetYearOfAwardOk returns a tuple with the YearOfAward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearOfAward

`func (o *Award) SetYearOfAward(v int32)`

SetYearOfAward sets YearOfAward field to given value.

### HasYearOfAward

`func (o *Award) HasYearOfAward() bool`

HasYearOfAward returns a boolean if a field has been set.

### SetYearOfAwardNil

`func (o *Award) SetYearOfAwardNil(b bool)`

 SetYearOfAwardNil sets the value for YearOfAward to be an explicit nil

### UnsetYearOfAward
`func (o *Award) UnsetYearOfAward()`

UnsetYearOfAward ensures that no value is present for YearOfAward, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


