# PersonAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**PictureSHA256** | Pointer to **string** | SHA256 hash of the full resolution image | [optional] [readonly] 
**Inactive** | Pointer to **bool** | Marks persons not working for the company anymore | [optional] [default to false]

## Methods

### NewPersonAllOf

`func NewPersonAllOf() *PersonAllOf`

NewPersonAllOf instantiates a new PersonAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonAllOfWithDefaults

`func NewPersonAllOfWithDefaults() *PersonAllOf`

NewPersonAllOfWithDefaults instantiates a new PersonAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployeeId

`func (o *PersonAllOf) GetEmployeeId() string`

GetEmployeeId returns the EmployeeId field if non-nil, zero value otherwise.

### GetEmployeeIdOk

`func (o *PersonAllOf) GetEmployeeIdOk() (*string, bool)`

GetEmployeeIdOk returns a tuple with the EmployeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeId

`func (o *PersonAllOf) SetEmployeeId(v string)`

SetEmployeeId sets EmployeeId field to given value.

### HasEmployeeId

`func (o *PersonAllOf) HasEmployeeId() bool`

HasEmployeeId returns a boolean if a field has been set.

### GetJobTitle

`func (o *PersonAllOf) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *PersonAllOf) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *PersonAllOf) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *PersonAllOf) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetCompany

`func (o *PersonAllOf) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PersonAllOf) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PersonAllOf) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *PersonAllOf) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetDepartment

`func (o *PersonAllOf) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *PersonAllOf) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *PersonAllOf) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *PersonAllOf) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetEducation

`func (o *PersonAllOf) GetEducation() string`

GetEducation returns the Education field if non-nil, zero value otherwise.

### GetEducationOk

`func (o *PersonAllOf) GetEducationOk() (*string, bool)`

GetEducationOk returns a tuple with the Education field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEducation

`func (o *PersonAllOf) SetEducation(v string)`

SetEducation sets Education field to given value.

### HasEducation

`func (o *PersonAllOf) HasEducation() bool`

HasEducation returns a boolean if a field has been set.

### GetMail

`func (o *PersonAllOf) GetMail() string`

GetMail returns the Mail field if non-nil, zero value otherwise.

### GetMailOk

`func (o *PersonAllOf) GetMailOk() (*string, bool)`

GetMailOk returns a tuple with the Mail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMail

`func (o *PersonAllOf) SetMail(v string)`

SetMail sets Mail field to given value.

### HasMail

`func (o *PersonAllOf) HasMail() bool`

HasMail returns a boolean if a field has been set.

### GetMobilePhone

`func (o *PersonAllOf) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *PersonAllOf) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *PersonAllOf) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *PersonAllOf) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetSeniority

`func (o *PersonAllOf) GetSeniority() string`

GetSeniority returns the Seniority field if non-nil, zero value otherwise.

### GetSeniorityOk

`func (o *PersonAllOf) GetSeniorityOk() (*string, bool)`

GetSeniorityOk returns a tuple with the Seniority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeniority

`func (o *PersonAllOf) SetSeniority(v string)`

SetSeniority sets Seniority field to given value.

### HasSeniority

`func (o *PersonAllOf) HasSeniority() bool`

HasSeniority returns a boolean if a field has been set.

### GetSeniorityEnum

`func (o *PersonAllOf) GetSeniorityEnum() Seniority`

GetSeniorityEnum returns the SeniorityEnum field if non-nil, zero value otherwise.

### GetSeniorityEnumOk

`func (o *PersonAllOf) GetSeniorityEnumOk() (*Seniority, bool)`

GetSeniorityEnumOk returns a tuple with the SeniorityEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeniorityEnum

`func (o *PersonAllOf) SetSeniorityEnum(v Seniority)`

SetSeniorityEnum sets SeniorityEnum field to given value.

### HasSeniorityEnum

`func (o *PersonAllOf) HasSeniorityEnum() bool`

HasSeniorityEnum returns a boolean if a field has been set.

### GetExperienceSinceYear

`func (o *PersonAllOf) GetExperienceSinceYear() int32`

GetExperienceSinceYear returns the ExperienceSinceYear field if non-nil, zero value otherwise.

### GetExperienceSinceYearOk

`func (o *PersonAllOf) GetExperienceSinceYearOk() (*int32, bool)`

GetExperienceSinceYearOk returns a tuple with the ExperienceSinceYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperienceSinceYear

`func (o *PersonAllOf) SetExperienceSinceYear(v int32)`

SetExperienceSinceYear sets ExperienceSinceYear field to given value.

### HasExperienceSinceYear

`func (o *PersonAllOf) HasExperienceSinceYear() bool`

HasExperienceSinceYear returns a boolean if a field has been set.

### GetOnsiteRatio

`func (o *PersonAllOf) GetOnsiteRatio() int32`

GetOnsiteRatio returns the OnsiteRatio field if non-nil, zero value otherwise.

### GetOnsiteRatioOk

`func (o *PersonAllOf) GetOnsiteRatioOk() (*int32, bool)`

GetOnsiteRatioOk returns a tuple with the OnsiteRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnsiteRatio

`func (o *PersonAllOf) SetOnsiteRatio(v int32)`

SetOnsiteRatio sets OnsiteRatio field to given value.

### HasOnsiteRatio

`func (o *PersonAllOf) HasOnsiteRatio() bool`

HasOnsiteRatio returns a boolean if a field has been set.

### GetPicture

`func (o *PersonAllOf) GetPicture() string`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *PersonAllOf) GetPictureOk() (*string, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *PersonAllOf) SetPicture(v string)`

SetPicture sets Picture field to given value.

### HasPicture

`func (o *PersonAllOf) HasPicture() bool`

HasPicture returns a boolean if a field has been set.

### GetPictureSHA256

`func (o *PersonAllOf) GetPictureSHA256() string`

GetPictureSHA256 returns the PictureSHA256 field if non-nil, zero value otherwise.

### GetPictureSHA256Ok

`func (o *PersonAllOf) GetPictureSHA256Ok() (*string, bool)`

GetPictureSHA256Ok returns a tuple with the PictureSHA256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictureSHA256

`func (o *PersonAllOf) SetPictureSHA256(v string)`

SetPictureSHA256 sets PictureSHA256 field to given value.

### HasPictureSHA256

`func (o *PersonAllOf) HasPictureSHA256() bool`

HasPictureSHA256 returns a boolean if a field has been set.

### GetInactive

`func (o *PersonAllOf) GetInactive() bool`

GetInactive returns the Inactive field if non-nil, zero value otherwise.

### GetInactiveOk

`func (o *PersonAllOf) GetInactiveOk() (*bool, bool)`

GetInactiveOk returns a tuple with the Inactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactive

`func (o *PersonAllOf) SetInactive(v bool)`

SetInactive sets Inactive field to given value.

### HasInactive

`func (o *PersonAllOf) HasInactive() bool`

HasInactive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


