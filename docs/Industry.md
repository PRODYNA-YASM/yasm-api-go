# Industry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Suggestion** | **bool** |  | [default to false]
**Synonyms** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIndustry

`func NewIndustry(suggestion bool, ) *Industry`

NewIndustry instantiates a new Industry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndustryWithDefaults

`func NewIndustryWithDefaults() *Industry`

NewIndustryWithDefaults instantiates a new Industry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuggestion

`func (o *Industry) GetSuggestion() bool`

GetSuggestion returns the Suggestion field if non-nil, zero value otherwise.

### GetSuggestionOk

`func (o *Industry) GetSuggestionOk() (*bool, bool)`

GetSuggestionOk returns a tuple with the Suggestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestion

`func (o *Industry) SetSuggestion(v bool)`

SetSuggestion sets Suggestion field to given value.


### GetSynonyms

`func (o *Industry) GetSynonyms() []string`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *Industry) GetSynonymsOk() (*[]string, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *Industry) SetSynonyms(v []string)`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *Industry) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


