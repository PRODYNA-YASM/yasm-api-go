/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.72.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the AwardDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwardDetails{}

// AwardDetails struct for AwardDetails
type AwardDetails struct {
	Audit *Audit `json:"audit,omitempty"`
	Award *Award `json:"award,omitempty"`
	Organization *Organization `json:"organization,omitempty"`
}

// NewAwardDetails instantiates a new AwardDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwardDetails() *AwardDetails {
	this := AwardDetails{}
	return &this
}

// NewAwardDetailsWithDefaults instantiates a new AwardDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwardDetailsWithDefaults() *AwardDetails {
	this := AwardDetails{}
	return &this
}

// GetAudit returns the Audit field value if set, zero value otherwise.
func (o *AwardDetails) GetAudit() Audit {
	if o == nil || IsNil(o.Audit) {
		var ret Audit
		return ret
	}
	return *o.Audit
}

// GetAuditOk returns a tuple with the Audit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwardDetails) GetAuditOk() (*Audit, bool) {
	if o == nil || IsNil(o.Audit) {
		return nil, false
	}
	return o.Audit, true
}

// HasAudit returns a boolean if a field has been set.
func (o *AwardDetails) HasAudit() bool {
	if o != nil && !IsNil(o.Audit) {
		return true
	}

	return false
}

// SetAudit gets a reference to the given Audit and assigns it to the Audit field.
func (o *AwardDetails) SetAudit(v Audit) {
	o.Audit = &v
}

// GetAward returns the Award field value if set, zero value otherwise.
func (o *AwardDetails) GetAward() Award {
	if o == nil || IsNil(o.Award) {
		var ret Award
		return ret
	}
	return *o.Award
}

// GetAwardOk returns a tuple with the Award field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwardDetails) GetAwardOk() (*Award, bool) {
	if o == nil || IsNil(o.Award) {
		return nil, false
	}
	return o.Award, true
}

// HasAward returns a boolean if a field has been set.
func (o *AwardDetails) HasAward() bool {
	if o != nil && !IsNil(o.Award) {
		return true
	}

	return false
}

// SetAward gets a reference to the given Award and assigns it to the Award field.
func (o *AwardDetails) SetAward(v Award) {
	o.Award = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *AwardDetails) GetOrganization() Organization {
	if o == nil || IsNil(o.Organization) {
		var ret Organization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwardDetails) GetOrganizationOk() (*Organization, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *AwardDetails) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Organization and assigns it to the Organization field.
func (o *AwardDetails) SetOrganization(v Organization) {
	o.Organization = &v
}

func (o AwardDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwardDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Audit) {
		toSerialize["audit"] = o.Audit
	}
	if !IsNil(o.Award) {
		toSerialize["award"] = o.Award
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	return toSerialize, nil
}

type NullableAwardDetails struct {
	value *AwardDetails
	isSet bool
}

func (v NullableAwardDetails) Get() *AwardDetails {
	return v.value
}

func (v *NullableAwardDetails) Set(val *AwardDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAwardDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAwardDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwardDetails(val *AwardDetails) *NullableAwardDetails {
	return &NullableAwardDetails{value: val, isSet: true}
}

func (v NullableAwardDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwardDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


