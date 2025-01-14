# ProjectImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **string** |  | [optional] 
**FieldImages** | Pointer to [**[]FieldImage**](FieldImage.md) |  | [optional] 

## Methods

### NewProjectImage

`func NewProjectImage() *ProjectImage`

NewProjectImage instantiates a new ProjectImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectImageWithDefaults

`func NewProjectImageWithDefaults() *ProjectImage`

NewProjectImageWithDefaults instantiates a new ProjectImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ProjectImage) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectImage) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectImage) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ProjectImage) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetFieldImages

`func (o *ProjectImage) GetFieldImages() []FieldImage`

GetFieldImages returns the FieldImages field if non-nil, zero value otherwise.

### GetFieldImagesOk

`func (o *ProjectImage) GetFieldImagesOk() (*[]FieldImage, bool)`

GetFieldImagesOk returns a tuple with the FieldImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldImages

`func (o *ProjectImage) SetFieldImages(v []FieldImage)`

SetFieldImages sets FieldImages field to given value.

### HasFieldImages

`func (o *ProjectImage) HasFieldImages() bool`

HasFieldImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


