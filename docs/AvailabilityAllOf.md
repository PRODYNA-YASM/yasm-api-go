# AvailabilityAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkHours** | **float32** |  | 
**PlannedHours** | **float32** |  | 
**Descriptions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAvailabilityAllOf

`func NewAvailabilityAllOf(workHours float32, plannedHours float32, ) *AvailabilityAllOf`

NewAvailabilityAllOf instantiates a new AvailabilityAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilityAllOfWithDefaults

`func NewAvailabilityAllOfWithDefaults() *AvailabilityAllOf`

NewAvailabilityAllOfWithDefaults instantiates a new AvailabilityAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkHours

`func (o *AvailabilityAllOf) GetWorkHours() float32`

GetWorkHours returns the WorkHours field if non-nil, zero value otherwise.

### GetWorkHoursOk

`func (o *AvailabilityAllOf) GetWorkHoursOk() (*float32, bool)`

GetWorkHoursOk returns a tuple with the WorkHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkHours

`func (o *AvailabilityAllOf) SetWorkHours(v float32)`

SetWorkHours sets WorkHours field to given value.


### GetPlannedHours

`func (o *AvailabilityAllOf) GetPlannedHours() float32`

GetPlannedHours returns the PlannedHours field if non-nil, zero value otherwise.

### GetPlannedHoursOk

`func (o *AvailabilityAllOf) GetPlannedHoursOk() (*float32, bool)`

GetPlannedHoursOk returns a tuple with the PlannedHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedHours

`func (o *AvailabilityAllOf) SetPlannedHours(v float32)`

SetPlannedHours sets PlannedHours field to given value.


### GetDescriptions

`func (o *AvailabilityAllOf) GetDescriptions() []string`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *AvailabilityAllOf) GetDescriptionsOk() (*[]string, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *AvailabilityAllOf) SetDescriptions(v []string)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *AvailabilityAllOf) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


