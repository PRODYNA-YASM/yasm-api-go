# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Suggestion** | **bool** |  | [default to false]
**Synonyms** | Pointer to **[]string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Geolocation** | Pointer to [**Geolocation**](Geolocation.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**External** | Pointer to **bool** | true if project was done outside of the organization | [optional] [default to false]
**ProjectType** | Pointer to **string** |  | [optional] 
**ProjectTypeEnum** | Pointer to [**ProjectType**](ProjectType.md) |  | [optional] 

## Methods

### NewProject

`func NewProject(suggestion bool, ) *Project`

NewProject instantiates a new Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDefaults

`func NewProjectWithDefaults() *Project`

NewProjectWithDefaults instantiates a new Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuggestion

`func (o *Project) GetSuggestion() bool`

GetSuggestion returns the Suggestion field if non-nil, zero value otherwise.

### GetSuggestionOk

`func (o *Project) GetSuggestionOk() (*bool, bool)`

GetSuggestionOk returns a tuple with the Suggestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestion

`func (o *Project) SetSuggestion(v bool)`

SetSuggestion sets Suggestion field to given value.


### GetSynonyms

`func (o *Project) GetSynonyms() []string`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *Project) GetSynonymsOk() (*[]string, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *Project) SetSynonyms(v []string)`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *Project) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.

### GetLocation

`func (o *Project) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Project) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Project) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Project) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetGeolocation

`func (o *Project) GetGeolocation() Geolocation`

GetGeolocation returns the Geolocation field if non-nil, zero value otherwise.

### GetGeolocationOk

`func (o *Project) GetGeolocationOk() (*Geolocation, bool)`

GetGeolocationOk returns a tuple with the Geolocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeolocation

`func (o *Project) SetGeolocation(v Geolocation)`

SetGeolocation sets Geolocation field to given value.

### HasGeolocation

`func (o *Project) HasGeolocation() bool`

HasGeolocation returns a boolean if a field has been set.

### GetDescription

`func (o *Project) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Project) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Project) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Project) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternal

`func (o *Project) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *Project) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *Project) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *Project) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetProjectType

`func (o *Project) GetProjectType() string`

GetProjectType returns the ProjectType field if non-nil, zero value otherwise.

### GetProjectTypeOk

`func (o *Project) GetProjectTypeOk() (*string, bool)`

GetProjectTypeOk returns a tuple with the ProjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectType

`func (o *Project) SetProjectType(v string)`

SetProjectType sets ProjectType field to given value.

### HasProjectType

`func (o *Project) HasProjectType() bool`

HasProjectType returns a boolean if a field has been set.

### GetProjectTypeEnum

`func (o *Project) GetProjectTypeEnum() ProjectType`

GetProjectTypeEnum returns the ProjectTypeEnum field if non-nil, zero value otherwise.

### GetProjectTypeEnumOk

`func (o *Project) GetProjectTypeEnumOk() (*ProjectType, bool)`

GetProjectTypeEnumOk returns a tuple with the ProjectTypeEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTypeEnum

`func (o *Project) SetProjectTypeEnum(v ProjectType)`

SetProjectTypeEnum sets ProjectTypeEnum field to given value.

### HasProjectTypeEnum

`func (o *Project) HasProjectTypeEnum() bool`

HasProjectTypeEnum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


