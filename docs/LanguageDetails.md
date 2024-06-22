# LanguageDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**Language** | Pointer to [**Language**](Language.md) |  | [optional] 
**Countries** | Pointer to [**[]Country**](Country.md) |  | [optional] 

## Methods

### NewLanguageDetails

`func NewLanguageDetails() *LanguageDetails`

NewLanguageDetails instantiates a new LanguageDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanguageDetailsWithDefaults

`func NewLanguageDetailsWithDefaults() *LanguageDetails`

NewLanguageDetailsWithDefaults instantiates a new LanguageDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *LanguageDetails) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *LanguageDetails) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *LanguageDetails) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *LanguageDetails) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetLanguage

`func (o *LanguageDetails) GetLanguage() Language`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *LanguageDetails) GetLanguageOk() (*Language, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *LanguageDetails) SetLanguage(v Language)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *LanguageDetails) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCountries

`func (o *LanguageDetails) GetCountries() []Country`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *LanguageDetails) GetCountriesOk() (*[]Country, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *LanguageDetails) SetCountries(v []Country)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *LanguageDetails) HasCountries() bool`

HasCountries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


