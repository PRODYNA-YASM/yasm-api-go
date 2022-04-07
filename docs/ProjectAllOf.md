# ProjectAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeframe** | Pointer to [**Timeframed**](Timeframed.md) |  | [optional] 
**External** | Pointer to **bool** | true if project was done outside of the organization | [optional] [default to false]

## Methods

### NewProjectAllOf

`func NewProjectAllOf() *ProjectAllOf`

NewProjectAllOf instantiates a new ProjectAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectAllOfWithDefaults

`func NewProjectAllOfWithDefaults() *ProjectAllOf`

NewProjectAllOfWithDefaults instantiates a new ProjectAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeframe

`func (o *ProjectAllOf) GetTimeframe() Timeframed`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *ProjectAllOf) GetTimeframeOk() (*Timeframed, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *ProjectAllOf) SetTimeframe(v Timeframed)`

SetTimeframe sets Timeframe field to given value.

### HasTimeframe

`func (o *ProjectAllOf) HasTimeframe() bool`

HasTimeframe returns a boolean if a field has been set.

### GetExternal

`func (o *ProjectAllOf) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *ProjectAllOf) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *ProjectAllOf) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *ProjectAllOf) HasExternal() bool`

HasExternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


