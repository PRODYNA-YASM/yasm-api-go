# Person

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Suggestion** | **bool** |  | [default to false]
**Synonyms** | Pointer to **[]string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Geolocation** | Pointer to [**Geolocation**](Geolocation.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EmployeeId** | Pointer to **string** |  | [optional] 
**JobTitle** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**Department** | Pointer to **string** |  | [optional] 
**Education** | Pointer to **string** |  | [optional] 
**Mail** | Pointer to **string** |  | [optional] 
**MobilePhone** | Pointer to **string** |  | [optional] 
**Seniority** | Pointer to **string** |  | [optional] 
**SeniorityEnum** | Pointer to [**Seniority**](Seniority.md) |  | [optional] 
**ExperienceSinceYear** | Pointer to **int32** |  | [optional] 
**OnsiteRatio** | Pointer to **int32** |  | [optional] 
**Picture** | Pointer to **string** | base64 encoded image | [optional] [readonly] 
**Inactive** | Pointer to **bool** | Marks persons not working for the company anymore | [optional] [default to false]

## Methods

### NewPerson

`func NewPerson(suggestion bool, ) *Person`

NewPerson instantiates a new Person object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonWithDefaults

`func NewPersonWithDefaults() *Person`

NewPersonWithDefaults instantiates a new Person object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuggestion

`func (o *Person) GetSuggestion() bool`

GetSuggestion returns the Suggestion field if non-nil, zero value otherwise.

### GetSuggestionOk

`func (o *Person) GetSuggestionOk() (*bool, bool)`

GetSuggestionOk returns a tuple with the Suggestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestion

`func (o *Person) SetSuggestion(v bool)`

SetSuggestion sets Suggestion field to given value.


### GetSynonyms

`func (o *Person) GetSynonyms() []string`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *Person) GetSynonymsOk() (*[]string, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *Person) SetSynonyms(v []string)`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *Person) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.

### GetLocation

`func (o *Person) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Person) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Person) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Person) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetGeolocation

`func (o *Person) GetGeolocation() Geolocation`

GetGeolocation returns the Geolocation field if non-nil, zero value otherwise.

### GetGeolocationOk

`func (o *Person) GetGeolocationOk() (*Geolocation, bool)`

GetGeolocationOk returns a tuple with the Geolocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeolocation

`func (o *Person) SetGeolocation(v Geolocation)`

SetGeolocation sets Geolocation field to given value.

### HasGeolocation

`func (o *Person) HasGeolocation() bool`

HasGeolocation returns a boolean if a field has been set.

### GetDescription

`func (o *Person) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Person) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Person) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Person) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmployeeId

`func (o *Person) GetEmployeeId() string`

GetEmployeeId returns the EmployeeId field if non-nil, zero value otherwise.

### GetEmployeeIdOk

`func (o *Person) GetEmployeeIdOk() (*string, bool)`

GetEmployeeIdOk returns a tuple with the EmployeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeId

`func (o *Person) SetEmployeeId(v string)`

SetEmployeeId sets EmployeeId field to given value.

### HasEmployeeId

`func (o *Person) HasEmployeeId() bool`

HasEmployeeId returns a boolean if a field has been set.

### GetJobTitle

`func (o *Person) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *Person) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *Person) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *Person) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetCompany

`func (o *Person) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Person) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Person) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Person) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetDepartment

`func (o *Person) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *Person) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *Person) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *Person) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetEducation

`func (o *Person) GetEducation() string`

GetEducation returns the Education field if non-nil, zero value otherwise.

### GetEducationOk

`func (o *Person) GetEducationOk() (*string, bool)`

GetEducationOk returns a tuple with the Education field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEducation

`func (o *Person) SetEducation(v string)`

SetEducation sets Education field to given value.

### HasEducation

`func (o *Person) HasEducation() bool`

HasEducation returns a boolean if a field has been set.

### GetMail

`func (o *Person) GetMail() string`

GetMail returns the Mail field if non-nil, zero value otherwise.

### GetMailOk

`func (o *Person) GetMailOk() (*string, bool)`

GetMailOk returns a tuple with the Mail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMail

`func (o *Person) SetMail(v string)`

SetMail sets Mail field to given value.

### HasMail

`func (o *Person) HasMail() bool`

HasMail returns a boolean if a field has been set.

### GetMobilePhone

`func (o *Person) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *Person) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *Person) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *Person) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetSeniority

`func (o *Person) GetSeniority() string`

GetSeniority returns the Seniority field if non-nil, zero value otherwise.

### GetSeniorityOk

`func (o *Person) GetSeniorityOk() (*string, bool)`

GetSeniorityOk returns a tuple with the Seniority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeniority

`func (o *Person) SetSeniority(v string)`

SetSeniority sets Seniority field to given value.

### HasSeniority

`func (o *Person) HasSeniority() bool`

HasSeniority returns a boolean if a field has been set.

### GetSeniorityEnum

`func (o *Person) GetSeniorityEnum() Seniority`

GetSeniorityEnum returns the SeniorityEnum field if non-nil, zero value otherwise.

### GetSeniorityEnumOk

`func (o *Person) GetSeniorityEnumOk() (*Seniority, bool)`

GetSeniorityEnumOk returns a tuple with the SeniorityEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeniorityEnum

`func (o *Person) SetSeniorityEnum(v Seniority)`

SetSeniorityEnum sets SeniorityEnum field to given value.

### HasSeniorityEnum

`func (o *Person) HasSeniorityEnum() bool`

HasSeniorityEnum returns a boolean if a field has been set.

### GetExperienceSinceYear

`func (o *Person) GetExperienceSinceYear() int32`

GetExperienceSinceYear returns the ExperienceSinceYear field if non-nil, zero value otherwise.

### GetExperienceSinceYearOk

`func (o *Person) GetExperienceSinceYearOk() (*int32, bool)`

GetExperienceSinceYearOk returns a tuple with the ExperienceSinceYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperienceSinceYear

`func (o *Person) SetExperienceSinceYear(v int32)`

SetExperienceSinceYear sets ExperienceSinceYear field to given value.

### HasExperienceSinceYear

`func (o *Person) HasExperienceSinceYear() bool`

HasExperienceSinceYear returns a boolean if a field has been set.

### GetOnsiteRatio

`func (o *Person) GetOnsiteRatio() int32`

GetOnsiteRatio returns the OnsiteRatio field if non-nil, zero value otherwise.

### GetOnsiteRatioOk

`func (o *Person) GetOnsiteRatioOk() (*int32, bool)`

GetOnsiteRatioOk returns a tuple with the OnsiteRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnsiteRatio

`func (o *Person) SetOnsiteRatio(v int32)`

SetOnsiteRatio sets OnsiteRatio field to given value.

### HasOnsiteRatio

`func (o *Person) HasOnsiteRatio() bool`

HasOnsiteRatio returns a boolean if a field has been set.

### GetPicture

`func (o *Person) GetPicture() string`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *Person) GetPictureOk() (*string, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *Person) SetPicture(v string)`

SetPicture sets Picture field to given value.

### HasPicture

`func (o *Person) HasPicture() bool`

HasPicture returns a boolean if a field has been set.

### GetInactive

`func (o *Person) GetInactive() bool`

GetInactive returns the Inactive field if non-nil, zero value otherwise.

### GetInactiveOk

`func (o *Person) GetInactiveOk() (*bool, bool)`

GetInactiveOk returns a tuple with the Inactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactive

`func (o *Person) SetInactive(v bool)`

SetInactive sets Inactive field to given value.

### HasInactive

`func (o *Person) HasInactive() bool`

HasInactive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


