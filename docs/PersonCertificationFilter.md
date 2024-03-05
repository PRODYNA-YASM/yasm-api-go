# PersonCertificationFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ObjectType** | Pointer to **string** |  | [optional] 
**StartedCertification** | Pointer to **bool** | Include employees who started the certification | [optional] 

## Methods

### NewPersonCertificationFilter

`func NewPersonCertificationFilter(id string, ) *PersonCertificationFilter`

NewPersonCertificationFilter instantiates a new PersonCertificationFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonCertificationFilterWithDefaults

`func NewPersonCertificationFilterWithDefaults() *PersonCertificationFilter`

NewPersonCertificationFilterWithDefaults instantiates a new PersonCertificationFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PersonCertificationFilter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PersonCertificationFilter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PersonCertificationFilter) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *PersonCertificationFilter) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PersonCertificationFilter) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PersonCertificationFilter) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *PersonCertificationFilter) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetStartedCertification

`func (o *PersonCertificationFilter) GetStartedCertification() bool`

GetStartedCertification returns the StartedCertification field if non-nil, zero value otherwise.

### GetStartedCertificationOk

`func (o *PersonCertificationFilter) GetStartedCertificationOk() (*bool, bool)`

GetStartedCertificationOk returns a tuple with the StartedCertification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedCertification

`func (o *PersonCertificationFilter) SetStartedCertification(v bool)`

SetStartedCertification sets StartedCertification field to given value.

### HasStartedCertification

`func (o *PersonCertificationFilter) HasStartedCertification() bool`

HasStartedCertification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


