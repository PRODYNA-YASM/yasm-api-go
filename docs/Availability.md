# Availability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Startdate** | **string** |  | 
**Enddate** | **string** |  | 
**WorkHours** | **float32** |  | 
**PlannedHours** | **float32** |  | 

## Methods

### NewAvailability

`func NewAvailability(id string, name string, startdate string, enddate string, workHours float32, plannedHours float32, ) *Availability`

NewAvailability instantiates a new Availability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilityWithDefaults

`func NewAvailabilityWithDefaults() *Availability`

NewAvailabilityWithDefaults instantiates a new Availability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Availability) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Availability) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Availability) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Availability) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Availability) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Availability) SetName(v string)`

SetName sets Name field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


