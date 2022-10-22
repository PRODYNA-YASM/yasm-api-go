/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.3.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SearchResult struct for SearchResult
type SearchResult struct {
	Page
	Items []SearchResultItem `json:"items,omitempty"`
}

// NewSearchResult instantiates a new SearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchResult() *SearchResult {
	this := SearchResult{}
	return &this
}

// NewSearchResultWithDefaults instantiates a new SearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchResultWithDefaults() *SearchResult {
	this := SearchResult{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SearchResult) GetItems() []SearchResultItem {
	if o == nil || o.Items == nil {
		var ret []SearchResultItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResult) GetItemsOk() ([]SearchResultItem, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SearchResult) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []SearchResultItem and assigns it to the Items field.
func (o *SearchResult) SetItems(v []SearchResultItem) {
	o.Items = v
}

func (o SearchResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPage, errPage := json.Marshal(o.Page)
	if errPage != nil {
		return []byte{}, errPage
	}
	errPage = json.Unmarshal([]byte(serializedPage), &toSerialize)
	if errPage != nil {
		return []byte{}, errPage
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableSearchResult struct {
	value *SearchResult
	isSet bool
}

func (v NullableSearchResult) Get() *SearchResult {
	return v.value
}

func (v *NullableSearchResult) Set(val *SearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchResult(val *SearchResult) *NullableSearchResult {
	return &NullableSearchResult{value: val, isSet: true}
}

func (v NullableSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


