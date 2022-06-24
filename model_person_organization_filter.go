/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PersonOrganizationFilter struct for PersonOrganizationFilter
type PersonOrganizationFilter struct {
	Id string `json:"id"`
	AmountOfProjects *MinMax `json:"amountOfProjects,omitempty"`
}

// NewPersonOrganizationFilter instantiates a new PersonOrganizationFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonOrganizationFilter(id string) *PersonOrganizationFilter {
	this := PersonOrganizationFilter{}
	this.Id = id
	return &this
}

// NewPersonOrganizationFilterWithDefaults instantiates a new PersonOrganizationFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonOrganizationFilterWithDefaults() *PersonOrganizationFilter {
	this := PersonOrganizationFilter{}
	return &this
}

// GetId returns the Id field value
func (o *PersonOrganizationFilter) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PersonOrganizationFilter) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PersonOrganizationFilter) SetId(v string) {
	o.Id = v
}

// GetAmountOfProjects returns the AmountOfProjects field value if set, zero value otherwise.
func (o *PersonOrganizationFilter) GetAmountOfProjects() MinMax {
	if o == nil || o.AmountOfProjects == nil {
		var ret MinMax
		return ret
	}
	return *o.AmountOfProjects
}

// GetAmountOfProjectsOk returns a tuple with the AmountOfProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonOrganizationFilter) GetAmountOfProjectsOk() (*MinMax, bool) {
	if o == nil || o.AmountOfProjects == nil {
		return nil, false
	}
	return o.AmountOfProjects, true
}

// HasAmountOfProjects returns a boolean if a field has been set.
func (o *PersonOrganizationFilter) HasAmountOfProjects() bool {
	if o != nil && o.AmountOfProjects != nil {
		return true
	}

	return false
}

// SetAmountOfProjects gets a reference to the given MinMax and assigns it to the AmountOfProjects field.
func (o *PersonOrganizationFilter) SetAmountOfProjects(v MinMax) {
	o.AmountOfProjects = &v
}

func (o PersonOrganizationFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.AmountOfProjects != nil {
		toSerialize["amountOfProjects"] = o.AmountOfProjects
	}
	return json.Marshal(toSerialize)
}

type NullablePersonOrganizationFilter struct {
	value *PersonOrganizationFilter
	isSet bool
}

func (v NullablePersonOrganizationFilter) Get() *PersonOrganizationFilter {
	return v.value
}

func (v *NullablePersonOrganizationFilter) Set(val *PersonOrganizationFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonOrganizationFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonOrganizationFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonOrganizationFilter(val *PersonOrganizationFilter) *NullablePersonOrganizationFilter {
	return &NullablePersonOrganizationFilter{value: val, isSet: true}
}

func (v NullablePersonOrganizationFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonOrganizationFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


