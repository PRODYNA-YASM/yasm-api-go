# OfficeDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**Office** | Pointer to [**Office**](Office.md) |  | [optional] 
**Country** | Pointer to [**Country**](Country.md) |  | [optional] 

## Methods

### NewOfficeDetails

`func NewOfficeDetails() *OfficeDetails`

NewOfficeDetails instantiates a new OfficeDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfficeDetailsWithDefaults

`func NewOfficeDetailsWithDefaults() *OfficeDetails`

NewOfficeDetailsWithDefaults instantiates a new OfficeDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *OfficeDetails) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *OfficeDetails) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *OfficeDetails) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *OfficeDetails) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetOffice

`func (o *OfficeDetails) GetOffice() Office`

GetOffice returns the Office field if non-nil, zero value otherwise.

### GetOfficeOk

`func (o *OfficeDetails) GetOfficeOk() (*Office, bool)`

GetOfficeOk returns a tuple with the Office field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice

`func (o *OfficeDetails) SetOffice(v Office)`

SetOffice sets Office field to given value.

### HasOffice

`func (o *OfficeDetails) HasOffice() bool`

HasOffice returns a boolean if a field has been set.

### GetCountry

`func (o *OfficeDetails) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *OfficeDetails) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *OfficeDetails) SetCountry(v Country)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *OfficeDetails) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


