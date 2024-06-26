# AvailabilityDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**Availability** | Pointer to [**Availability**](Availability.md) |  | [optional] 
**Percent** | Pointer to **float32** |  | [optional] 

## Methods

### NewAvailabilityDetail

`func NewAvailabilityDetail() *AvailabilityDetail`

NewAvailabilityDetail instantiates a new AvailabilityDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilityDetailWithDefaults

`func NewAvailabilityDetailWithDefaults() *AvailabilityDetail`

NewAvailabilityDetailWithDefaults instantiates a new AvailabilityDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *AvailabilityDetail) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *AvailabilityDetail) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *AvailabilityDetail) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *AvailabilityDetail) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetAvailability

`func (o *AvailabilityDetail) GetAvailability() Availability`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *AvailabilityDetail) GetAvailabilityOk() (*Availability, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *AvailabilityDetail) SetAvailability(v Availability)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *AvailabilityDetail) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetPercent

`func (o *AvailabilityDetail) GetPercent() float32`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *AvailabilityDetail) GetPercentOk() (*float32, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *AvailabilityDetail) SetPercent(v float32)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *AvailabilityDetail) HasPercent() bool`

HasPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


