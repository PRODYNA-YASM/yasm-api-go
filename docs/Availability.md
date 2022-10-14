# Availability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkHours** | **float32** |  | 
**PlannedHours** | **float32** |  | 
**Descriptions** | Pointer to **[]string** |  | [optional] 
**Startdate** | **string** |  | 
**Enddate** | **string** |  | 

## Methods

### NewAvailability

`func NewAvailability(workHours float32, plannedHours float32, startdate string, enddate string, ) *Availability`

NewAvailability instantiates a new Availability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilityWithDefaults

`func NewAvailabilityWithDefaults() *Availability`

NewAvailabilityWithDefaults instantiates a new Availability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkHours

`func (o *Availability) GetWorkHours() float32`

GetWorkHours returns the WorkHours field if non-nil, zero value otherwise.

### GetWorkHoursOk

`func (o *Availability) GetWorkHoursOk() (*float32, bool)`

GetWorkHoursOk returns a tuple with the WorkHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkHours

`func (o *Availability) SetWorkHours(v float32)`

SetWorkHours sets WorkHours field to given value.


### GetPlannedHours

`func (o *Availability) GetPlannedHours() float32`

GetPlannedHours returns the PlannedHours field if non-nil, zero value otherwise.

### GetPlannedHoursOk

`func (o *Availability) GetPlannedHoursOk() (*float32, bool)`

GetPlannedHoursOk returns a tuple with the PlannedHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedHours

`func (o *Availability) SetPlannedHours(v float32)`

SetPlannedHours sets PlannedHours field to given value.


### GetDescriptions

`func (o *Availability) GetDescriptions() []string`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *Availability) GetDescriptionsOk() (*[]string, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *Availability) SetDescriptions(v []string)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *Availability) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetStartdate

`func (o *Availability) GetStartdate() string`

GetStartdate returns the Startdate field if non-nil, zero value otherwise.

### GetStartdateOk

`func (o *Availability) GetStartdateOk() (*string, bool)`

GetStartdateOk returns a tuple with the Startdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartdate

`func (o *Availability) SetStartdate(v string)`

SetStartdate sets Startdate field to given value.


### GetEnddate

`func (o *Availability) GetEnddate() string`

GetEnddate returns the Enddate field if non-nil, zero value otherwise.

### GetEnddateOk

`func (o *Availability) GetEnddateOk() (*string, bool)`

GetEnddateOk returns a tuple with the Enddate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnddate

`func (o *Availability) SetEnddate(v string)`

SetEnddate sets Enddate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


