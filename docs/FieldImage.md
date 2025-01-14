# FieldImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to [**ImageDetails**](ImageDetails.md) |  | [optional] 
**LayoutFieldId** | Pointer to **string** |  | [optional] 

## Methods

### NewFieldImage

`func NewFieldImage() *FieldImage`

NewFieldImage instantiates a new FieldImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldImageWithDefaults

`func NewFieldImageWithDefaults() *FieldImage`

NewFieldImageWithDefaults instantiates a new FieldImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *FieldImage) GetImage() ImageDetails`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *FieldImage) GetImageOk() (*ImageDetails, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *FieldImage) SetImage(v ImageDetails)`

SetImage sets Image field to given value.

### HasImage

`func (o *FieldImage) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetLayoutFieldId

`func (o *FieldImage) GetLayoutFieldId() string`

GetLayoutFieldId returns the LayoutFieldId field if non-nil, zero value otherwise.

### GetLayoutFieldIdOk

`func (o *FieldImage) GetLayoutFieldIdOk() (*string, bool)`

GetLayoutFieldIdOk returns a tuple with the LayoutFieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutFieldId

`func (o *FieldImage) SetLayoutFieldId(v string)`

SetLayoutFieldId sets LayoutFieldId field to given value.

### HasLayoutFieldId

`func (o *FieldImage) HasLayoutFieldId() bool`

HasLayoutFieldId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


