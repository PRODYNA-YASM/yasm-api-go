/*
 * YASM (Yet Another Skill Management) API
 *
 * This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.
 *
 * API version: 0.5.10
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PagedCertifications struct for PagedCertifications
type PagedCertifications struct {
	Skip *int32 `json:"skip,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	Count *int32 `json:"count,omitempty"`
	Certifications *[]Certification `json:"certifications,omitempty"`
}

// NewPagedCertifications instantiates a new PagedCertifications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedCertifications() *PagedCertifications {
	this := PagedCertifications{}
	return &this
}

// NewPagedCertificationsWithDefaults instantiates a new PagedCertifications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedCertificationsWithDefaults() *PagedCertifications {
	this := PagedCertifications{}
	return &this
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *PagedCertifications) GetSkip() int32 {
	if o == nil || o.Skip == nil {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedCertifications) GetSkipOk() (*int32, bool) {
	if o == nil || o.Skip == nil {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *PagedCertifications) HasSkip() bool {
	if o != nil && o.Skip != nil {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *PagedCertifications) SetSkip(v int32) {
	o.Skip = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *PagedCertifications) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedCertifications) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *PagedCertifications) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *PagedCertifications) SetLimit(v int32) {
	o.Limit = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PagedCertifications) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedCertifications) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PagedCertifications) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PagedCertifications) SetCount(v int32) {
	o.Count = &v
}

// GetCertifications returns the Certifications field value if set, zero value otherwise.
func (o *PagedCertifications) GetCertifications() []Certification {
	if o == nil || o.Certifications == nil {
		var ret []Certification
		return ret
	}
	return *o.Certifications
}

// GetCertificationsOk returns a tuple with the Certifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedCertifications) GetCertificationsOk() (*[]Certification, bool) {
	if o == nil || o.Certifications == nil {
		return nil, false
	}
	return o.Certifications, true
}

// HasCertifications returns a boolean if a field has been set.
func (o *PagedCertifications) HasCertifications() bool {
	if o != nil && o.Certifications != nil {
		return true
	}

	return false
}

// SetCertifications gets a reference to the given []Certification and assigns it to the Certifications field.
func (o *PagedCertifications) SetCertifications(v []Certification) {
	o.Certifications = &v
}

func (o PagedCertifications) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Skip != nil {
		toSerialize["skip"] = o.Skip
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Certifications != nil {
		toSerialize["certifications"] = o.Certifications
	}
	return json.Marshal(toSerialize)
}

type NullablePagedCertifications struct {
	value *PagedCertifications
	isSet bool
}

func (v NullablePagedCertifications) Get() *PagedCertifications {
	return v.value
}

func (v *NullablePagedCertifications) Set(val *PagedCertifications) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedCertifications) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedCertifications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedCertifications(val *PagedCertifications) *NullablePagedCertifications {
	return &NullablePagedCertifications{value: val, isSet: true}
}

func (v NullablePagedCertifications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedCertifications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


