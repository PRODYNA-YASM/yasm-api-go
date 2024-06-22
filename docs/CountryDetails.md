# CountryDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**Country** | Pointer to [**Country**](Country.md) |  | [optional] 
**Languages** | Pointer to [**[]Language**](Language.md) |  | [optional] 

## Methods

### NewCountryDetails

`func NewCountryDetails() *CountryDetails`

NewCountryDetails instantiates a new CountryDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryDetailsWithDefaults

`func NewCountryDetailsWithDefaults() *CountryDetails`

NewCountryDetailsWithDefaults instantiates a new CountryDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *CountryDetails) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *CountryDetails) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *CountryDetails) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *CountryDetails) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetCountry

`func (o *CountryDetails) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CountryDetails) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CountryDetails) SetCountry(v Country)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CountryDetails) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetLanguages

`func (o *CountryDetails) GetLanguages() []Language`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *CountryDetails) GetLanguagesOk() (*[]Language, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *CountryDetails) SetLanguages(v []Language)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *CountryDetails) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


