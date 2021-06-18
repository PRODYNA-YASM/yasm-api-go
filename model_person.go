/*
 * YASM (Yet Another Skill Management) API
 *
 * This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.
 *
 * API version: 0.6.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Person struct for Person
type Person struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Suggestion *bool `json:"suggestion,omitempty"`
	Synonyms *[]string `json:"synonyms,omitempty"`
	Location *string `json:"location,omitempty"`
}

// NewPerson instantiates a new Person object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerson(id string, name string) *Person {
	this := Person{}
	this.Id = id
	this.Name = name
	var suggestion bool = false
	this.Suggestion = &suggestion
	return &this
}

// NewPersonWithDefaults instantiates a new Person object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonWithDefaults() *Person {
	this := Person{}
	var suggestion bool = false
	this.Suggestion = &suggestion
	return &this
}

// GetId returns the Id field value
func (o *Person) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Person) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Person) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Person) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Person) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Person) SetName(v string) {
	o.Name = v
}

// GetSuggestion returns the Suggestion field value if set, zero value otherwise.
func (o *Person) GetSuggestion() bool {
	if o == nil || o.Suggestion == nil {
		var ret bool
		return ret
	}
	return *o.Suggestion
}

// GetSuggestionOk returns a tuple with the Suggestion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetSuggestionOk() (*bool, bool) {
	if o == nil || o.Suggestion == nil {
		return nil, false
	}
	return o.Suggestion, true
}

// HasSuggestion returns a boolean if a field has been set.
func (o *Person) HasSuggestion() bool {
	if o != nil && o.Suggestion != nil {
		return true
	}

	return false
}

// SetSuggestion gets a reference to the given bool and assigns it to the Suggestion field.
func (o *Person) SetSuggestion(v bool) {
	o.Suggestion = &v
}

// GetSynonyms returns the Synonyms field value if set, zero value otherwise.
func (o *Person) GetSynonyms() []string {
	if o == nil || o.Synonyms == nil {
		var ret []string
		return ret
	}
	return *o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetSynonymsOk() (*[]string, bool) {
	if o == nil || o.Synonyms == nil {
		return nil, false
	}
	return o.Synonyms, true
}

// HasSynonyms returns a boolean if a field has been set.
func (o *Person) HasSynonyms() bool {
	if o != nil && o.Synonyms != nil {
		return true
	}

	return false
}

// SetSynonyms gets a reference to the given []string and assigns it to the Synonyms field.
func (o *Person) SetSynonyms(v []string) {
	o.Synonyms = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *Person) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *Person) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *Person) SetLocation(v string) {
	o.Location = &v
}

func (o Person) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Suggestion != nil {
		toSerialize["suggestion"] = o.Suggestion
	}
	if o.Synonyms != nil {
		toSerialize["synonyms"] = o.Synonyms
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	return json.Marshal(toSerialize)
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


