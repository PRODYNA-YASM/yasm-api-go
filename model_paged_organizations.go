/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PagedOrganizations struct for PagedOrganizations
type PagedOrganizations struct {
	Page
	Organizations []OrganizationDetails `json:"organizations,omitempty"`
}

// NewPagedOrganizations instantiates a new PagedOrganizations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedOrganizations() *PagedOrganizations {
	this := PagedOrganizations{}
	return &this
}

// NewPagedOrganizationsWithDefaults instantiates a new PagedOrganizations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedOrganizationsWithDefaults() *PagedOrganizations {
	this := PagedOrganizations{}
	return &this
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *PagedOrganizations) GetOrganizations() []OrganizationDetails {
	if o == nil || o.Organizations == nil {
		var ret []OrganizationDetails
		return ret
	}
	return o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedOrganizations) GetOrganizationsOk() ([]OrganizationDetails, bool) {
	if o == nil || o.Organizations == nil {
		return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *PagedOrganizations) HasOrganizations() bool {
	if o != nil && o.Organizations != nil {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given []OrganizationDetails and assigns it to the Organizations field.
func (o *PagedOrganizations) SetOrganizations(v []OrganizationDetails) {
	o.Organizations = v
}

func (o PagedOrganizations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPage, errPage := json.Marshal(o.Page)
	if errPage != nil {
		return []byte{}, errPage
	}
	errPage = json.Unmarshal([]byte(serializedPage), &toSerialize)
	if errPage != nil {
		return []byte{}, errPage
	}
	if o.Organizations != nil {
		toSerialize["organizations"] = o.Organizations
	}
	return json.Marshal(toSerialize)
}

type NullablePagedOrganizations struct {
	value *PagedOrganizations
	isSet bool
}

func (v NullablePagedOrganizations) Get() *PagedOrganizations {
	return v.value
}

func (v *NullablePagedOrganizations) Set(val *PagedOrganizations) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedOrganizations) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedOrganizations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedOrganizations(val *PagedOrganizations) *NullablePagedOrganizations {
	return &NullablePagedOrganizations{value: val, isSet: true}
}

func (v NullablePagedOrganizations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedOrganizations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


