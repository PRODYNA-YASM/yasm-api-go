# Date

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Year** | **float32** |  | 
**Month** | **int32** |  | 
**Day** | Pointer to **int32** |  | [optional] 

## Methods

### NewDate

`func NewDate(year float32, month int32, ) *Date`

NewDate instantiates a new Date object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateWithDefaults

`func NewDateWithDefaults() *Date`

NewDateWithDefaults instantiates a new Date object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetYear

`func (o *Date) GetYear() float32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *Date) GetYearOk() (*float32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *Date) SetYear(v float32)`

SetYear sets Year field to given value.


### GetMonth

`func (o *Date) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *Date) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *Date) SetMonth(v int32)`

SetMonth sets Month field to given value.


### GetDay

`func (o *Date) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *Date) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *Date) SetDay(v int32)`

SetDay sets Day field to given value.

### HasDay

`func (o *Date) HasDay() bool`

HasDay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


