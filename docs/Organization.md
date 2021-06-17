# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Suggestion** | Pointer to **bool** |  | [optional] [default to false]
**Partner** | Pointer to **bool** |  | [optional] [default to false]
**Customer** | Pointer to **bool** |  | [optional] [default to false]
**Geolocation** | Pointer to [**Geolocation**](Geolocation.md) |  | [optional] 

## Methods

### NewOrganization

`func NewOrganization(id string, name string, ) *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetName

`func (o *Organization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Organization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Organization) SetName(v string)`

SetName sets Name field to given value.


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

### HasSuggestion

`func (o *Organization) HasSuggestion() bool`

HasSuggestion returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


