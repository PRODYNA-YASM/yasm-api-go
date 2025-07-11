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

// checks if the ProjectView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectView{}

// ProjectView struct for ProjectView
type ProjectView struct {
	Project *Project `json:"project,omitempty"`
	Organization *Organization `json:"organization,omitempty"`
	Timeframe *Timeframed `json:"timeframe,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
}

// NewProjectView instantiates a new ProjectView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectView() *ProjectView {
	this := ProjectView{}
	return &this
}

// NewProjectViewWithDefaults instantiates a new ProjectView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectViewWithDefaults() *ProjectView {
	this := ProjectView{}
	return &this
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ProjectView) GetProject() Project {
	if o == nil || IsNil(o.Project) {
		var ret Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectView) GetProjectOk() (*Project, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ProjectView) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Project and assigns it to the Project field.
func (o *ProjectView) SetProject(v Project) {
	o.Project = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ProjectView) GetOrganization() Organization {
	if o == nil || IsNil(o.Organization) {
		var ret Organization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectView) GetOrganizationOk() (*Organization, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ProjectView) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Organization and assigns it to the Organization field.
func (o *ProjectView) SetOrganization(v Organization) {
	o.Organization = &v
}

// GetTimeframe returns the Timeframe field value if set, zero value otherwise.
func (o *ProjectView) GetTimeframe() Timeframed {
	if o == nil || IsNil(o.Timeframe) {
		var ret Timeframed
		return ret
	}
	return *o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectView) GetTimeframeOk() (*Timeframed, bool) {
	if o == nil || IsNil(o.Timeframe) {
		return nil, false
	}
	return o.Timeframe, true
}

// HasTimeframe returns a boolean if a field has been set.
func (o *ProjectView) HasTimeframe() bool {
	if o != nil && !IsNil(o.Timeframe) {
		return true
	}

	return false
}

// SetTimeframe gets a reference to the given Timeframed and assigns it to the Timeframe field.
func (o *ProjectView) SetTimeframe(v Timeframed) {
	o.Timeframe = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *ProjectView) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectView) GetObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *ProjectView) HasObjectType() bool {
	if o != nil && !IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *ProjectView) SetObjectType(v string) {
	o.ObjectType = &v
}

func (o ProjectView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.Timeframe) {
		toSerialize["timeframe"] = o.Timeframe
	}
	if !IsNil(o.ObjectType) {
		toSerialize["objectType"] = o.ObjectType
	}
	return toSerialize, nil
}

type NullableProjectView struct {
	value *ProjectView
	isSet bool
}

func (v NullableProjectView) Get() *ProjectView {
	return v.value
}

func (v *NullableProjectView) Set(val *ProjectView) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectView) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectView(val *ProjectView) *NullableProjectView {
	return &NullableProjectView{value: val, isSet: true}
}

func (v NullableProjectView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


