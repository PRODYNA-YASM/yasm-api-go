/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the OfficeSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfficeSearch{}

// OfficeSearch struct for OfficeSearch
type OfficeSearch struct {
	Search
	OrganizationIds []string `json:"organizationIds,omitempty"`
}

// NewOfficeSearch instantiates a new OfficeSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfficeSearch(skip int32, limit int32) *OfficeSearch {
	this := OfficeSearch{}
	this.Skip = skip
	this.Limit = limit
	return &this
}

// NewOfficeSearchWithDefaults instantiates a new OfficeSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfficeSearchWithDefaults() *OfficeSearch {
	this := OfficeSearch{}
	return &this
}

// GetOrganizationIds returns the OrganizationIds field value if set, zero value otherwise.
func (o *OfficeSearch) GetOrganizationIds() []string {
	if o == nil || IsNil(o.OrganizationIds) {
		var ret []string
		return ret
	}
	return o.OrganizationIds
}

// GetOrganizationIdsOk returns a tuple with the OrganizationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfficeSearch) GetOrganizationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OrganizationIds) {
		return nil, false
	}
	return o.OrganizationIds, true
}

// HasOrganizationIds returns a boolean if a field has been set.
func (o *OfficeSearch) HasOrganizationIds() bool {
	if o != nil && !IsNil(o.OrganizationIds) {
		return true
	}

	return false
}

// SetOrganizationIds gets a reference to the given []string and assigns it to the OrganizationIds field.
func (o *OfficeSearch) SetOrganizationIds(v []string) {
	o.OrganizationIds = v
}

func (o OfficeSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfficeSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSearch, errSearch := json.Marshal(o.Search)
	if errSearch != nil {
		return map[string]interface{}{}, errSearch
	}
	errSearch = json.Unmarshal([]byte(serializedSearch), &toSerialize)
	if errSearch != nil {
		return map[string]interface{}{}, errSearch
	}
	if !IsNil(o.OrganizationIds) {
		toSerialize["organizationIds"] = o.OrganizationIds
	}
	return toSerialize, nil
}

type NullableOfficeSearch struct {
	value *OfficeSearch
	isSet bool
}

func (v NullableOfficeSearch) Get() *OfficeSearch {
	return v.value
}

func (v *NullableOfficeSearch) Set(val *OfficeSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableOfficeSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableOfficeSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfficeSearch(val *OfficeSearch) *NullableOfficeSearch {
	return &NullableOfficeSearch{value: val, isSet: true}
}

func (v NullableOfficeSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfficeSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

