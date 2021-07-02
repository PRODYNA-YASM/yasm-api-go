# SearchResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | Pointer to **float32** |  | [optional] 
**Type** | **string** |  | 
**Item** | Pointer to [**NamedDomainModel**](NamedDomainModel.md) |  | [optional] 

## Methods

### NewSearchResultItem

`func NewSearchResultItem(type_ string, ) *SearchResultItem`

NewSearchResultItem instantiates a new SearchResultItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultItemWithDefaults

`func NewSearchResultItemWithDefaults() *SearchResultItem`

NewSearchResultItemWithDefaults instantiates a new SearchResultItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *SearchResultItem) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *SearchResultItem) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *SearchResultItem) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *SearchResultItem) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetType

`func (o *SearchResultItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchResultItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchResultItem) SetType(v string)`

SetType sets Type field to given value.


### GetItem

`func (o *SearchResultItem) GetItem() NamedDomainModel`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *SearchResultItem) GetItemOk() (*NamedDomainModel, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *SearchResultItem) SetItem(v NamedDomainModel)`

SetItem sets Item field to given value.

### HasItem

`func (o *SearchResultItem) HasItem() bool`

HasItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


