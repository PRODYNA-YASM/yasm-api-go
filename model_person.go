/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the Person type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Person{}

// Person struct for Person
type Person struct {
	NamedDomainModel
	Synonyms []string `json:"synonyms,omitempty"`
	Location *string `json:"location,omitempty"`
	Geolocation *Geolocation `json:"geolocation,omitempty"`
	Description *string `json:"description,omitempty"`
	EmployeeId *string `json:"employeeId,omitempty"`
	JobTitle *string `json:"jobTitle,omitempty"`
	Company *string `json:"company,omitempty"`
	Department *string `json:"department,omitempty"`
	Education *string `json:"education,omitempty"`
	Mail *string `json:"mail,omitempty"`
	MobilePhone *string `json:"mobilePhone,omitempty"`
	Seniority *string `json:"seniority,omitempty"`
	SeniorityEnum *Seniority `json:"seniorityEnum,omitempty"`
	ExperienceSinceYear *int32 `json:"experienceSinceYear,omitempty"`
	OnsiteRatio *int32 `json:"onsiteRatio,omitempty"`
	// base64 encoded image
	Picture *string `json:"picture,omitempty"`
	// SHA256 hash of the full resolution image
	PictureSHA256 *string `json:"pictureSHA256,omitempty"`
	// Marks persons not working for the company anymore
	Inactive *bool `json:"inactive,omitempty"`
}

// NewPerson instantiates a new Person object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerson(id string, name string) *Person {
	this := Person{}
	this.Id = id
	this.Name = name
	var inactive bool = false
	this.Inactive = &inactive
	return &this
}

// NewPersonWithDefaults instantiates a new Person object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonWithDefaults() *Person {
	this := Person{}
	var inactive bool = false
	this.Inactive = &inactive
	return &this
}

// GetSynonyms returns the Synonyms field value if set, zero value otherwise.
func (o *Person) GetSynonyms() []string {
	if o == nil || IsNil(o.Synonyms) {
		var ret []string
		return ret
	}
	return o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetSynonymsOk() ([]string, bool) {
	if o == nil || IsNil(o.Synonyms) {
		return nil, false
	}
	return o.Synonyms, true
}

// HasSynonyms returns a boolean if a field has been set.
func (o *Person) HasSynonyms() bool {
	if o != nil && !IsNil(o.Synonyms) {
		return true
	}

	return false
}

// SetSynonyms gets a reference to the given []string and assigns it to the Synonyms field.
func (o *Person) SetSynonyms(v []string) {
	o.Synonyms = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *Person) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *Person) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *Person) SetLocation(v string) {
	o.Location = &v
}

// GetGeolocation returns the Geolocation field value if set, zero value otherwise.
func (o *Person) GetGeolocation() Geolocation {
	if o == nil || IsNil(o.Geolocation) {
		var ret Geolocation
		return ret
	}
	return *o.Geolocation
}

// GetGeolocationOk returns a tuple with the Geolocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetGeolocationOk() (*Geolocation, bool) {
	if o == nil || IsNil(o.Geolocation) {
		return nil, false
	}
	return o.Geolocation, true
}

// HasGeolocation returns a boolean if a field has been set.
func (o *Person) HasGeolocation() bool {
	if o != nil && !IsNil(o.Geolocation) {
		return true
	}

	return false
}

// SetGeolocation gets a reference to the given Geolocation and assigns it to the Geolocation field.
func (o *Person) SetGeolocation(v Geolocation) {
	o.Geolocation = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Person) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Person) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Person) SetDescription(v string) {
	o.Description = &v
}

// GetEmployeeId returns the EmployeeId field value if set, zero value otherwise.
func (o *Person) GetEmployeeId() string {
	if o == nil || IsNil(o.EmployeeId) {
		var ret string
		return ret
	}
	return *o.EmployeeId
}

// GetEmployeeIdOk returns a tuple with the EmployeeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetEmployeeIdOk() (*string, bool) {
	if o == nil || IsNil(o.EmployeeId) {
		return nil, false
	}
	return o.EmployeeId, true
}

// HasEmployeeId returns a boolean if a field has been set.
func (o *Person) HasEmployeeId() bool {
	if o != nil && !IsNil(o.EmployeeId) {
		return true
	}

	return false
}

// SetEmployeeId gets a reference to the given string and assigns it to the EmployeeId field.
func (o *Person) SetEmployeeId(v string) {
	o.EmployeeId = &v
}

// GetJobTitle returns the JobTitle field value if set, zero value otherwise.
func (o *Person) GetJobTitle() string {
	if o == nil || IsNil(o.JobTitle) {
		var ret string
		return ret
	}
	return *o.JobTitle
}

// GetJobTitleOk returns a tuple with the JobTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetJobTitleOk() (*string, bool) {
	if o == nil || IsNil(o.JobTitle) {
		return nil, false
	}
	return o.JobTitle, true
}

// HasJobTitle returns a boolean if a field has been set.
func (o *Person) HasJobTitle() bool {
	if o != nil && !IsNil(o.JobTitle) {
		return true
	}

	return false
}

// SetJobTitle gets a reference to the given string and assigns it to the JobTitle field.
func (o *Person) SetJobTitle(v string) {
	o.JobTitle = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *Person) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *Person) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *Person) SetCompany(v string) {
	o.Company = &v
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *Person) GetDepartment() string {
	if o == nil || IsNil(o.Department) {
		var ret string
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetDepartmentOk() (*string, bool) {
	if o == nil || IsNil(o.Department) {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *Person) HasDepartment() bool {
	if o != nil && !IsNil(o.Department) {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given string and assigns it to the Department field.
func (o *Person) SetDepartment(v string) {
	o.Department = &v
}

// GetEducation returns the Education field value if set, zero value otherwise.
func (o *Person) GetEducation() string {
	if o == nil || IsNil(o.Education) {
		var ret string
		return ret
	}
	return *o.Education
}

// GetEducationOk returns a tuple with the Education field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetEducationOk() (*string, bool) {
	if o == nil || IsNil(o.Education) {
		return nil, false
	}
	return o.Education, true
}

// HasEducation returns a boolean if a field has been set.
func (o *Person) HasEducation() bool {
	if o != nil && !IsNil(o.Education) {
		return true
	}

	return false
}

// SetEducation gets a reference to the given string and assigns it to the Education field.
func (o *Person) SetEducation(v string) {
	o.Education = &v
}

// GetMail returns the Mail field value if set, zero value otherwise.
func (o *Person) GetMail() string {
	if o == nil || IsNil(o.Mail) {
		var ret string
		return ret
	}
	return *o.Mail
}

// GetMailOk returns a tuple with the Mail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetMailOk() (*string, bool) {
	if o == nil || IsNil(o.Mail) {
		return nil, false
	}
	return o.Mail, true
}

// HasMail returns a boolean if a field has been set.
func (o *Person) HasMail() bool {
	if o != nil && !IsNil(o.Mail) {
		return true
	}

	return false
}

// SetMail gets a reference to the given string and assigns it to the Mail field.
func (o *Person) SetMail(v string) {
	o.Mail = &v
}

// GetMobilePhone returns the MobilePhone field value if set, zero value otherwise.
func (o *Person) GetMobilePhone() string {
	if o == nil || IsNil(o.MobilePhone) {
		var ret string
		return ret
	}
	return *o.MobilePhone
}

// GetMobilePhoneOk returns a tuple with the MobilePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetMobilePhoneOk() (*string, bool) {
	if o == nil || IsNil(o.MobilePhone) {
		return nil, false
	}
	return o.MobilePhone, true
}

// HasMobilePhone returns a boolean if a field has been set.
func (o *Person) HasMobilePhone() bool {
	if o != nil && !IsNil(o.MobilePhone) {
		return true
	}

	return false
}

// SetMobilePhone gets a reference to the given string and assigns it to the MobilePhone field.
func (o *Person) SetMobilePhone(v string) {
	o.MobilePhone = &v
}

// GetSeniority returns the Seniority field value if set, zero value otherwise.
func (o *Person) GetSeniority() string {
	if o == nil || IsNil(o.Seniority) {
		var ret string
		return ret
	}
	return *o.Seniority
}

// GetSeniorityOk returns a tuple with the Seniority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetSeniorityOk() (*string, bool) {
	if o == nil || IsNil(o.Seniority) {
		return nil, false
	}
	return o.Seniority, true
}

// HasSeniority returns a boolean if a field has been set.
func (o *Person) HasSeniority() bool {
	if o != nil && !IsNil(o.Seniority) {
		return true
	}

	return false
}

// SetSeniority gets a reference to the given string and assigns it to the Seniority field.
func (o *Person) SetSeniority(v string) {
	o.Seniority = &v
}

// GetSeniorityEnum returns the SeniorityEnum field value if set, zero value otherwise.
func (o *Person) GetSeniorityEnum() Seniority {
	if o == nil || IsNil(o.SeniorityEnum) {
		var ret Seniority
		return ret
	}
	return *o.SeniorityEnum
}

// GetSeniorityEnumOk returns a tuple with the SeniorityEnum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetSeniorityEnumOk() (*Seniority, bool) {
	if o == nil || IsNil(o.SeniorityEnum) {
		return nil, false
	}
	return o.SeniorityEnum, true
}

// HasSeniorityEnum returns a boolean if a field has been set.
func (o *Person) HasSeniorityEnum() bool {
	if o != nil && !IsNil(o.SeniorityEnum) {
		return true
	}

	return false
}

// SetSeniorityEnum gets a reference to the given Seniority and assigns it to the SeniorityEnum field.
func (o *Person) SetSeniorityEnum(v Seniority) {
	o.SeniorityEnum = &v
}

// GetExperienceSinceYear returns the ExperienceSinceYear field value if set, zero value otherwise.
func (o *Person) GetExperienceSinceYear() int32 {
	if o == nil || IsNil(o.ExperienceSinceYear) {
		var ret int32
		return ret
	}
	return *o.ExperienceSinceYear
}

// GetExperienceSinceYearOk returns a tuple with the ExperienceSinceYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetExperienceSinceYearOk() (*int32, bool) {
	if o == nil || IsNil(o.ExperienceSinceYear) {
		return nil, false
	}
	return o.ExperienceSinceYear, true
}

// HasExperienceSinceYear returns a boolean if a field has been set.
func (o *Person) HasExperienceSinceYear() bool {
	if o != nil && !IsNil(o.ExperienceSinceYear) {
		return true
	}

	return false
}

// SetExperienceSinceYear gets a reference to the given int32 and assigns it to the ExperienceSinceYear field.
func (o *Person) SetExperienceSinceYear(v int32) {
	o.ExperienceSinceYear = &v
}

// GetOnsiteRatio returns the OnsiteRatio field value if set, zero value otherwise.
func (o *Person) GetOnsiteRatio() int32 {
	if o == nil || IsNil(o.OnsiteRatio) {
		var ret int32
		return ret
	}
	return *o.OnsiteRatio
}

// GetOnsiteRatioOk returns a tuple with the OnsiteRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetOnsiteRatioOk() (*int32, bool) {
	if o == nil || IsNil(o.OnsiteRatio) {
		return nil, false
	}
	return o.OnsiteRatio, true
}

// HasOnsiteRatio returns a boolean if a field has been set.
func (o *Person) HasOnsiteRatio() bool {
	if o != nil && !IsNil(o.OnsiteRatio) {
		return true
	}

	return false
}

// SetOnsiteRatio gets a reference to the given int32 and assigns it to the OnsiteRatio field.
func (o *Person) SetOnsiteRatio(v int32) {
	o.OnsiteRatio = &v
}

// GetPicture returns the Picture field value if set, zero value otherwise.
func (o *Person) GetPicture() string {
	if o == nil || IsNil(o.Picture) {
		var ret string
		return ret
	}
	return *o.Picture
}

// GetPictureOk returns a tuple with the Picture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetPictureOk() (*string, bool) {
	if o == nil || IsNil(o.Picture) {
		return nil, false
	}
	return o.Picture, true
}

// HasPicture returns a boolean if a field has been set.
func (o *Person) HasPicture() bool {
	if o != nil && !IsNil(o.Picture) {
		return true
	}

	return false
}

// SetPicture gets a reference to the given string and assigns it to the Picture field.
func (o *Person) SetPicture(v string) {
	o.Picture = &v
}

// GetPictureSHA256 returns the PictureSHA256 field value if set, zero value otherwise.
func (o *Person) GetPictureSHA256() string {
	if o == nil || IsNil(o.PictureSHA256) {
		var ret string
		return ret
	}
	return *o.PictureSHA256
}

// GetPictureSHA256Ok returns a tuple with the PictureSHA256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetPictureSHA256Ok() (*string, bool) {
	if o == nil || IsNil(o.PictureSHA256) {
		return nil, false
	}
	return o.PictureSHA256, true
}

// HasPictureSHA256 returns a boolean if a field has been set.
func (o *Person) HasPictureSHA256() bool {
	if o != nil && !IsNil(o.PictureSHA256) {
		return true
	}

	return false
}

// SetPictureSHA256 gets a reference to the given string and assigns it to the PictureSHA256 field.
func (o *Person) SetPictureSHA256(v string) {
	o.PictureSHA256 = &v
}

// GetInactive returns the Inactive field value if set, zero value otherwise.
func (o *Person) GetInactive() bool {
	if o == nil || IsNil(o.Inactive) {
		var ret bool
		return ret
	}
	return *o.Inactive
}

// GetInactiveOk returns a tuple with the Inactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetInactiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Inactive) {
		return nil, false
	}
	return o.Inactive, true
}

// HasInactive returns a boolean if a field has been set.
func (o *Person) HasInactive() bool {
	if o != nil && !IsNil(o.Inactive) {
		return true
	}

	return false
}

// SetInactive gets a reference to the given bool and assigns it to the Inactive field.
func (o *Person) SetInactive(v bool) {
	o.Inactive = &v
}

func (o Person) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Person) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedNamedDomainModel, errNamedDomainModel := json.Marshal(o.NamedDomainModel)
	if errNamedDomainModel != nil {
		return map[string]interface{}{}, errNamedDomainModel
	}
	errNamedDomainModel = json.Unmarshal([]byte(serializedNamedDomainModel), &toSerialize)
	if errNamedDomainModel != nil {
		return map[string]interface{}{}, errNamedDomainModel
	}
	if !IsNil(o.Synonyms) {
		toSerialize["synonyms"] = o.Synonyms
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Geolocation) {
		toSerialize["geolocation"] = o.Geolocation
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.EmployeeId) {
		toSerialize["employeeId"] = o.EmployeeId
	}
	if !IsNil(o.JobTitle) {
		toSerialize["jobTitle"] = o.JobTitle
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.Department) {
		toSerialize["department"] = o.Department
	}
	if !IsNil(o.Education) {
		toSerialize["education"] = o.Education
	}
	if !IsNil(o.Mail) {
		toSerialize["mail"] = o.Mail
	}
	if !IsNil(o.MobilePhone) {
		toSerialize["mobilePhone"] = o.MobilePhone
	}
	if !IsNil(o.Seniority) {
		toSerialize["seniority"] = o.Seniority
	}
	if !IsNil(o.SeniorityEnum) {
		toSerialize["seniorityEnum"] = o.SeniorityEnum
	}
	if !IsNil(o.ExperienceSinceYear) {
		toSerialize["experienceSinceYear"] = o.ExperienceSinceYear
	}
	if !IsNil(o.OnsiteRatio) {
		toSerialize["onsiteRatio"] = o.OnsiteRatio
	}
	if !IsNil(o.Picture) {
		toSerialize["picture"] = o.Picture
	}
	if !IsNil(o.PictureSHA256) {
		toSerialize["pictureSHA256"] = o.PictureSHA256
	}
	if !IsNil(o.Inactive) {
		toSerialize["inactive"] = o.Inactive
	}
	return toSerialize, nil
}

type NullablePerson struct {
	value *Person
	isSet bool
}

func (v NullablePerson) Get() *Person {
	return v.value
}

func (v *NullablePerson) Set(val *Person) {
	v.value = val
	v.isSet = true
}

func (v NullablePerson) IsSet() bool {
	return v.isSet
}

func (v *NullablePerson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerson(val *Person) *NullablePerson {
	return &NullablePerson{value: val, isSet: true}
}

func (v NullablePerson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


