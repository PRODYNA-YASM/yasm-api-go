# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Partner** | Pointer to **bool** |  | [optional] [default to false]
**Customer** | Pointer to **bool** |  | [optional] [default to false]
**Picture** | Pointer to **string** | base64 encoded image | [optional] 
**Id** | **string** |  | 
**ObjectType** | Pointer to **string** |  | [optional] 
**Suggestion** | **bool** |  | [default to false]
**Synonyms** | Pointer to **[]string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Geolocation** | Pointer to [**Geolocation**](Geolocation.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewOrganization

`func NewOrganization(id string, suggestion bool, ) *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartner

`func (o *Organization) GetPartner() bool`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *Organization) GetPartnerOk() (*bool, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *Organization) SetPartner(v bool)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *Organization) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetCustomer

`func (o *Organization) GetCustomer() bool`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *Organization) GetCustomerOk() (*bool, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *Organization) SetCustomer(v bool)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *Organization) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetPicture

`func (o *Organization) GetPicture() string`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *Organization) GetPictureOk() (*string, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *Organization) SetPicture(v string)`

SetPicture sets Picture field to given value.

### HasPicture

`func (o *Organization) HasPicture() bool`

HasPicture returns a boolean if a field has been set.

### GetId

`func (o *Organization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Organization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Organization) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *Organization) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *Organization) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *Organization) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *Organization) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetSuggestion

`func (o *Organization) GetSuggestion() bool`

GetSuggestion returns the Suggestion field if non-nil, zero value otherwise.

### GetSuggestionOk

`func (o *Organization) GetSuggestionOk() (*bool, bool)`

GetSuggestionOk returns a tuple with the Suggestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestion

`func (o *Organization) SetSuggestion(v bool)`

SetSuggestion sets Suggestion field to given value.


### GetSynonyms

`func (o *Organization) GetSynonyms() []string`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *Organization) GetSynonymsOk() (*[]string, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *Organization) SetSynonyms(v []string)`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *Organization) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.

### GetLocation

`func (o *Organization) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Organization) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Organization) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Organization) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetGeolocation

`func (o *Organization) GetGeolocation() Geolocation`

GetGeolocation returns the Geolocation field if non-nil, zero value otherwise.

### GetGeolocationOk

`func (o *Organization) GetGeolocationOk() (*Geolocation, bool)`

GetGeolocationOk returns a tuple with the Geolocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeolocation

`func (o *Organization) SetGeolocation(v Geolocation)`

SetGeolocation sets Geolocation field to given value.

### HasGeolocation

`func (o *Organization) HasGeolocation() bool`

HasGeolocation returns a boolean if a field has been set.

### GetDescription

`func (o *Organization) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Organization) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Organization) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Organization) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


