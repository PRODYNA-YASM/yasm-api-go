# PersonDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Person** | Pointer to [**Person**](Person.md) |  | [optional] 
**Projects** | Pointer to [**[]ProjectParticipation**](ProjectParticipation.md) |  | [optional] 
**Interests** | Pointer to [**[]Skill**](Skill.md) |  | [optional] 
**Certifications** | Pointer to [**[]CertificationDetails**](CertificationDetails.md) |  | [optional] 
**Languages** | Pointer to [**[]LanguageLevel**](LanguageLevel.md) |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Office** | Pointer to [**Office**](Office.md) |  | [optional] 

## Methods

### NewPersonDetails

`func NewPersonDetails() *PersonDetails`

NewPersonDetails instantiates a new PersonDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonDetailsWithDefaults

`func NewPersonDetailsWithDefaults() *PersonDetails`

NewPersonDetailsWithDefaults instantiates a new PersonDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPerson

`func (o *PersonDetails) GetPerson() Person`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *PersonDetails) GetPersonOk() (*Person, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *PersonDetails) SetPerson(v Person)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *PersonDetails) HasPerson() bool`

HasPerson returns a boolean if a field has been set.

### GetProjects

`func (o *PersonDetails) GetProjects() []ProjectParticipation`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *PersonDetails) GetProjectsOk() (*[]ProjectParticipation, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *PersonDetails) SetProjects(v []ProjectParticipation)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *PersonDetails) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetInterests

`func (o *PersonDetails) GetInterests() []Skill`

GetInterests returns the Interests field if non-nil, zero value otherwise.

### GetInterestsOk

`func (o *PersonDetails) GetInterestsOk() (*[]Skill, bool)`

GetInterestsOk returns a tuple with the Interests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterests

`func (o *PersonDetails) SetInterests(v []Skill)`

SetInterests sets Interests field to given value.

### HasInterests

`func (o *PersonDetails) HasInterests() bool`

HasInterests returns a boolean if a field has been set.

### GetCertifications

`func (o *PersonDetails) GetCertifications() []CertificationDetails`

GetCertifications returns the Certifications field if non-nil, zero value otherwise.

### GetCertificationsOk

`func (o *PersonDetails) GetCertificationsOk() (*[]CertificationDetails, bool)`

GetCertificationsOk returns a tuple with the Certifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifications

`func (o *PersonDetails) SetCertifications(v []CertificationDetails)`

SetCertifications sets Certifications field to given value.

### HasCertifications

`func (o *PersonDetails) HasCertifications() bool`

HasCertifications returns a boolean if a field has been set.

### GetLanguages

`func (o *PersonDetails) GetLanguages() []LanguageLevel`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *PersonDetails) GetLanguagesOk() (*[]LanguageLevel, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *PersonDetails) SetLanguages(v []LanguageLevel)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *PersonDetails) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetLocation

`func (o *PersonDetails) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PersonDetails) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PersonDetails) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PersonDetails) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetOffice

`func (o *PersonDetails) GetOffice() Office`

GetOffice returns the Office field if non-nil, zero value otherwise.

### GetOfficeOk

`func (o *PersonDetails) GetOfficeOk() (*Office, bool)`

GetOfficeOk returns a tuple with the Office field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice

`func (o *PersonDetails) SetOffice(v Office)`

SetOffice sets Office field to given value.

### HasOffice

`func (o *PersonDetails) HasOffice() bool`

HasOffice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


