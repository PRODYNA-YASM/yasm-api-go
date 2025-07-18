/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.74.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ProfileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileRequest{}

// ProfileRequest struct for ProfileRequest
type ProfileRequest struct {
	Template *string `json:"template,omitempty"`
	Person *PersonDetails `json:"person,omitempty"`
	ProjectParticipation []PersonProjectParticipationDetails `json:"projectParticipation,omitempty"`
	Anonymize *bool `json:"anonymize,omitempty"`
}

// NewProfileRequest instantiates a new ProfileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileRequest() *ProfileRequest {
	this := ProfileRequest{}
	return &this
}

// NewProfileRequestWithDefaults instantiates a new ProfileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileRequestWithDefaults() *ProfileRequest {
	this := ProfileRequest{}
	return &this
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *ProfileRequest) GetTemplate() string {
	if o == nil || IsNil(o.Template) {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileRequest) GetTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *ProfileRequest) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *ProfileRequest) SetTemplate(v string) {
	o.Template = &v
}

// GetPerson returns the Person field value if set, zero value otherwise.
func (o *ProfileRequest) GetPerson() PersonDetails {
	if o == nil || IsNil(o.Person) {
		var ret PersonDetails
		return ret
	}
	return *o.Person
}

// GetPersonOk returns a tuple with the Person field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileRequest) GetPersonOk() (*PersonDetails, bool) {
	if o == nil || IsNil(o.Person) {
		return nil, false
	}
	return o.Person, true
}

// HasPerson returns a boolean if a field has been set.
func (o *ProfileRequest) HasPerson() bool {
	if o != nil && !IsNil(o.Person) {
		return true
	}

	return false
}

// SetPerson gets a reference to the given PersonDetails and assigns it to the Person field.
func (o *ProfileRequest) SetPerson(v PersonDetails) {
	o.Person = &v
}

// GetProjectParticipation returns the ProjectParticipation field value if set, zero value otherwise.
func (o *ProfileRequest) GetProjectParticipation() []PersonProjectParticipationDetails {
	if o == nil || IsNil(o.ProjectParticipation) {
		var ret []PersonProjectParticipationDetails
		return ret
	}
	return o.ProjectParticipation
}

// GetProjectParticipationOk returns a tuple with the ProjectParticipation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileRequest) GetProjectParticipationOk() ([]PersonProjectParticipationDetails, bool) {
	if o == nil || IsNil(o.ProjectParticipation) {
		return nil, false
	}
	return o.ProjectParticipation, true
}

// HasProjectParticipation returns a boolean if a field has been set.
func (o *ProfileRequest) HasProjectParticipation() bool {
	if o != nil && !IsNil(o.ProjectParticipation) {
		return true
	}

	return false
}

// SetProjectParticipation gets a reference to the given []PersonProjectParticipationDetails and assigns it to the ProjectParticipation field.
func (o *ProfileRequest) SetProjectParticipation(v []PersonProjectParticipationDetails) {
	o.ProjectParticipation = v
}

// GetAnonymize returns the Anonymize field value if set, zero value otherwise.
func (o *ProfileRequest) GetAnonymize() bool {
	if o == nil || IsNil(o.Anonymize) {
		var ret bool
		return ret
	}
	return *o.Anonymize
}

// GetAnonymizeOk returns a tuple with the Anonymize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileRequest) GetAnonymizeOk() (*bool, bool) {
	if o == nil || IsNil(o.Anonymize) {
		return nil, false
	}
	return o.Anonymize, true
}

// HasAnonymize returns a boolean if a field has been set.
func (o *ProfileRequest) HasAnonymize() bool {
	if o != nil && !IsNil(o.Anonymize) {
		return true
	}

	return false
}

// SetAnonymize gets a reference to the given bool and assigns it to the Anonymize field.
func (o *ProfileRequest) SetAnonymize(v bool) {
	o.Anonymize = &v
}

func (o ProfileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	if !IsNil(o.Person) {
		toSerialize["person"] = o.Person
	}
	if !IsNil(o.ProjectParticipation) {
		toSerialize["projectParticipation"] = o.ProjectParticipation
	}
	if !IsNil(o.Anonymize) {
		toSerialize["anonymize"] = o.Anonymize
	}
	return toSerialize, nil
}

type NullableProfileRequest struct {
	value *ProfileRequest
	isSet bool
}

func (v NullableProfileRequest) Get() *ProfileRequest {
	return v.value
}

func (v *NullableProfileRequest) Set(val *ProfileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileRequest(val *ProfileRequest) *NullableProfileRequest {
	return &NullableProfileRequest{value: val, isSet: true}
}

func (v NullableProfileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


