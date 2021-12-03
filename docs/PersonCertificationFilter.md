# PersonCertificationFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**StartedCertificaiton** | Pointer to **bool** | Include employees who started the certification | [optional] 

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


### GetStartedCertificaiton

`func (o *PersonCertificationFilter) GetStartedCertificaiton() bool`

GetStartedCertificaiton returns the StartedCertificaiton field if non-nil, zero value otherwise.

### GetStartedCertificaitonOk

`func (o *PersonCertificationFilter) GetStartedCertificaitonOk() (*bool, bool)`

GetStartedCertificaitonOk returns a tuple with the StartedCertificaiton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedCertificaiton

`func (o *PersonCertificationFilter) SetStartedCertificaiton(v bool)`

SetStartedCertificaiton sets StartedCertificaiton field to given value.

### HasStartedCertificaiton

`func (o *PersonCertificationFilter) HasStartedCertificaiton() bool`

HasStartedCertificaiton returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


