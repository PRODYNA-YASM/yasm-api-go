/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.3.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Organization struct for Organization
type Organization struct {
	NamedDomainModel
	Suggestion bool `json:"suggestion"`
	Synonyms []string `json:"synonyms,omitempty"`
	Location *string `json:"location,omitempty"`
	Geolocation *Geolocation `json:"geolocation,omitempty"`
	Description *string `json:"description,omitempty"`
	Partner *bool `json:"partner,omitempty"`
	Customer *bool `json:"customer,omitempty"`
	// base64 encoded image
	Picture *string `json:"picture,omitempty"`
}

// NewOrganization instantiates a new Organization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganization(suggestion bool, id string, name string) *Organization {
	this := Organization{}
	this.Id = id
	this.Name = name
	this.Suggestion = suggestion
	var partner bool = false
	this.Partner = &partner
	var customer bool = false
	this.Customer = &customer
	return &this
}

// NewOrganizationWithDefaults instantiates a new Organization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationWithDefaults() *Organization {
	this := Organization{}
	var suggestion bool = false
	this.Suggestion = suggestion
	var partner bool = false
	this.Partner = &partner
	var customer bool = false
	this.Customer = &customer
	return &this
}

// GetSuggestion returns the Suggestion field value
func (o *Organization) GetSuggestion() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Suggestion
}

// GetSuggestionOk returns a tuple with the Suggestion field value
// and a boolean to check if the value has been set.
func (o *Organization) GetSuggestionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Suggestion, true
}

// SetSuggestion sets field value
func (o *Organization) SetSuggestion(v bool) {
	o.Suggestion = v
}

// GetSynonyms returns the Synonyms field value if set, zero value otherwise.
func (o *Organization) GetSynonyms() []string {
	if o == nil || o.Synonyms == nil {
		var ret []string
		return ret
	}
	return o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetSynonymsOk() ([]string, bool) {
	if o == nil || o.Synonyms == nil {
		return nil, false
	}
	return o.Synonyms, true
}

// HasSynonyms returns a boolean if a field has been set.
func (o *Organization) HasSynonyms() bool {
	if o != nil && o.Synonyms != nil {
		return true
	}

	return false
}

// SetSynonyms gets a reference to the given []string and assigns it to the Synonyms field.
func (o *Organization) SetSynonyms(v []string) {
	o.Synonyms = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *Organization) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *Organization) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *Organization) SetLocation(v string) {
	o.Location = &v
}

// GetGeolocation returns the Geolocation field value if set, zero value otherwise.
func (o *Organization) GetGeolocation() Geolocation {
	if o == nil || o.Geolocation == nil {
		var ret Geolocation
		return ret
	}
	return *o.Geolocation
}

// GetGeolocationOk returns a tuple with the Geolocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetGeolocationOk() (*Geolocation, bool) {
	if o == nil || o.Geolocation == nil {
		return nil, false
	}
	return o.Geolocation, true
}

// HasGeolocation returns a boolean if a field has been set.
func (o *Organization) HasGeolocation() bool {
	if o != nil && o.Geolocation != nil {
		return true
	}

	return false
}

// SetGeolocation gets a reference to the given Geolocation and assigns it to the Geolocation field.
func (o *Organization) SetGeolocation(v Geolocation) {
	o.Geolocation = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Organization) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Organization) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Organization) SetDescription(v string) {
	o.Description = &v
}

// GetPartner returns the Partner field value if set, zero value otherwise.
func (o *Organization) GetPartner() bool {
	if o == nil || o.Partner == nil {
		var ret bool
		return ret
	}
	return *o.Partner
}

// GetPartnerOk returns a tuple with the Partner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetPartnerOk() (*bool, bool) {
	if o == nil || o.Partner == nil {
		return nil, false
	}
	return o.Partner, true
}

// HasPartner returns a boolean if a field has been set.
func (o *Organization) HasPartner() bool {
	if o != nil && o.Partner != nil {
		return true
	}

	return false
}

// SetPartner gets a reference to the given bool and assigns it to the Partner field.
func (o *Organization) SetPartner(v bool) {
	o.Partner = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *Organization) GetCustomer() bool {
	if o == nil || o.Customer == nil {
		var ret bool
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetCustomerOk() (*bool, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *Organization) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given bool and assigns it to the Customer field.
func (o *Organization) SetCustomer(v bool) {
	o.Customer = &v
}

// GetPicture returns the Picture field value if set, zero value otherwise.
func (o *Organization) GetPicture() string {
	if o == nil || o.Picture == nil {
		var ret string
		return ret
	}
	return *o.Picture
}

// GetPictureOk returns a tuple with the Picture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetPictureOk() (*string, bool) {
	if o == nil || o.Picture == nil {
		return nil, false
	}
	return o.Picture, true
}

// HasPicture returns a boolean if a field has been set.
func (o *Organization) HasPicture() bool {
	if o != nil && o.Picture != nil {
		return true
	}

	return false
}

// SetPicture gets a reference to the given string and assigns it to the Picture field.
func (o *Organization) SetPicture(v string) {
	o.Picture = &v
}

func (o Organization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedNamedDomainModel, errNamedDomainModel := json.Marshal(o.NamedDomainModel)
	if errNamedDomainModel != nil {
		return []byte{}, errNamedDomainModel
	}
	errNamedDomainModel = json.Unmarshal([]byte(serializedNamedDomainModel), &toSerialize)
	if errNamedDomainModel != nil {
		return []byte{}, errNamedDomainModel
	}
	if true {
		toSerialize["suggestion"] = o.Suggestion
	}
	if o.Synonyms != nil {
		toSerialize["synonyms"] = o.Synonyms
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Geolocation != nil {
		toSerialize["geolocation"] = o.Geolocation
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Partner != nil {
		toSerialize["partner"] = o.Partner
	}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.Picture != nil {
		toSerialize["picture"] = o.Picture
	}
	return json.Marshal(toSerialize)
}

type NullableOrganization struct {
	value *Organization
	isSet bool
}

func (v NullableOrganization) Get() *Organization {
	return v.value
}

func (v *NullableOrganization) Set(val *Organization) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganization(val *Organization) *NullableOrganization {
	return &NullableOrganization{value: val, isSet: true}
}

func (v NullableOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


