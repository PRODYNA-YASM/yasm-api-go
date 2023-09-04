# ProjectAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**External** | Pointer to **bool** | true if project was done outside of the organization | [optional] [default to false]
**ProjectType** | Pointer to **string** |  | [optional] 
**ProjectTypeEnum** | Pointer to [**ProjectType**](ProjectType.md) |  | [optional] 

## Methods

### NewProjectAllOf

`func NewProjectAllOf() *ProjectAllOf`

NewProjectAllOf instantiates a new ProjectAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectAllOfWithDefaults

`func NewProjectAllOfWithDefaults() *ProjectAllOf`

NewProjectAllOfWithDefaults instantiates a new ProjectAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternal

`func (o *ProjectAllOf) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *ProjectAllOf) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *ProjectAllOf) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *ProjectAllOf) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetProjectType

`func (o *ProjectAllOf) GetProjectType() string`

GetProjectType returns the ProjectType field if non-nil, zero value otherwise.

### GetProjectTypeOk

`func (o *ProjectAllOf) GetProjectTypeOk() (*string, bool)`

GetProjectTypeOk returns a tuple with the ProjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectType

`func (o *ProjectAllOf) SetProjectType(v string)`

SetProjectType sets ProjectType field to given value.

### HasProjectType

`func (o *ProjectAllOf) HasProjectType() bool`

HasProjectType returns a boolean if a field has been set.

### GetProjectTypeEnum

`func (o *ProjectAllOf) GetProjectTypeEnum() ProjectType`

GetProjectTypeEnum returns the ProjectTypeEnum field if non-nil, zero value otherwise.

### GetProjectTypeEnumOk

`func (o *ProjectAllOf) GetProjectTypeEnumOk() (*ProjectType, bool)`

GetProjectTypeEnumOk returns a tuple with the ProjectTypeEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTypeEnum

`func (o *ProjectAllOf) SetProjectTypeEnum(v ProjectType)`

SetProjectTypeEnum sets ProjectTypeEnum field to given value.

### HasProjectTypeEnum

`func (o *ProjectAllOf) HasProjectTypeEnum() bool`

HasProjectTypeEnum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


