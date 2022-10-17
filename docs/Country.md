# Country

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Picture** | Pointer to **string** | base64 encoded image | [optional] 
**Id** | **string** |  | 
**ObjectType** | Pointer to **string** |  | [optional] 

## Methods

### NewCountry

`func NewCountry(id string, ) *Country`

NewCountry instantiates a new Country object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryWithDefaults

`func NewCountryWithDefaults() *Country`

NewCountryWithDefaults instantiates a new Country object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPicture

`func (o *Country) GetPicture() string`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *Country) GetPictureOk() (*string, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *Country) SetPicture(v string)`

SetPicture sets Picture field to given value.

### HasPicture

`func (o *Country) HasPicture() bool`

HasPicture returns a boolean if a field has been set.

### GetId

`func (o *Country) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Country) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Country) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *Country) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *Country) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *Country) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *Country) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


