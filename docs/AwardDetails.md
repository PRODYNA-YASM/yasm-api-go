# AwardDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**Award** | Pointer to [**Award**](Award.md) |  | [optional] 
**Organization** | Pointer to [**Organization**](Organization.md) |  | [optional] 

## Methods

### NewAwardDetails

`func NewAwardDetails() *AwardDetails`

NewAwardDetails instantiates a new AwardDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwardDetailsWithDefaults

`func NewAwardDetailsWithDefaults() *AwardDetails`

NewAwardDetailsWithDefaults instantiates a new AwardDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *AwardDetails) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *AwardDetails) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *AwardDetails) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *AwardDetails) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetAward

`func (o *AwardDetails) GetAward() Award`

GetAward returns the Award field if non-nil, zero value otherwise.

### GetAwardOk

`func (o *AwardDetails) GetAwardOk() (*Award, bool)`

GetAwardOk returns a tuple with the Award field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAward

`func (o *AwardDetails) SetAward(v Award)`

SetAward sets Award field to given value.

### HasAward

`func (o *AwardDetails) HasAward() bool`

HasAward returns a boolean if a field has been set.

### GetOrganization

`func (o *AwardDetails) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *AwardDetails) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *AwardDetails) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *AwardDetails) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


