/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the Audit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Audit{}

// Audit struct for Audit
type Audit struct {
	Created *time.Time `json:"created,omitempty"`
	CreatedBy *string `json:"createdBy,omitempty"`
	CreatedById *string `json:"createdById,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
	UpdatedBy *string `json:"updatedBy,omitempty"`
	UpdatedById *string `json:"updatedById,omitempty"`
}

// NewAudit instantiates a new Audit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAudit() *Audit {
	this := Audit{}
	return &this
}

// NewAuditWithDefaults instantiates a new Audit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditWithDefaults() *Audit {
	this := Audit{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Audit) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audit) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Audit) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Audit) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Audit) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audit) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Audit) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *Audit) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedById returns the CreatedById field value if set, zero value otherwise.
func (o *Audit) GetCreatedById() string {
	if o == nil || IsNil(o.CreatedById) {
		var ret string
		return ret
	}
	return *o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audit) GetCreatedByIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedById) {
		return nil, false
	}
	return o.CreatedById, true
}

// HasCreatedById returns a boolean if a field has been set.
func (o *Audit) HasCreatedById() bool {
	if o != nil && !IsNil(o.CreatedById) {
		return true
	}

	return false
}

// SetCreatedById gets a reference to the given string and assigns it to the CreatedById field.
func (o *Audit) SetCreatedById(v string) {
	o.CreatedById = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *Audit) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audit) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *Audit) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *Audit) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *Audit) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audit) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *Audit) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *Audit) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetUpdatedById returns the UpdatedById field value if set, zero value otherwise.
func (o *Audit) GetUpdatedById() string {
	if o == nil || IsNil(o.UpdatedById) {
		var ret string
		return ret
	}
	return *o.UpdatedById
}

// GetUpdatedByIdOk returns a tuple with the UpdatedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audit) GetUpdatedByIdOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedById) {
		return nil, false
	}
	return o.UpdatedById, true
}

// HasUpdatedById returns a boolean if a field has been set.
func (o *Audit) HasUpdatedById() bool {
	if o != nil && !IsNil(o.UpdatedById) {
		return true
	}

	return false
}

// SetUpdatedById gets a reference to the given string and assigns it to the UpdatedById field.
func (o *Audit) SetUpdatedById(v string) {
	o.UpdatedById = &v
}

func (o Audit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Audit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.CreatedById) {
		toSerialize["createdById"] = o.CreatedById
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	if !IsNil(o.UpdatedById) {
		toSerialize["updatedById"] = o.UpdatedById
	}
	return toSerialize, nil
}

type NullableAudit struct {
	value *Audit
	isSet bool
}

func (v NullableAudit) Get() *Audit {
	return v.value
}

func (v *NullableAudit) Set(val *Audit) {
	v.value = val
	v.isSet = true
}

func (v NullableAudit) IsSet() bool {
	return v.isSet
}

func (v *NullableAudit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudit(val *Audit) *NullableAudit {
	return &NullableAudit{value: val, isSet: true}
}

func (v NullableAudit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

