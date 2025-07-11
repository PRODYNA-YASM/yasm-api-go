# ProjectLayoutDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**ProjectLayout** | Pointer to [**ProjectLayout**](ProjectLayout.md) |  | [optional] 
**LayoutId** | Pointer to **string** |  | [optional] 
**Overlays** | Pointer to [**[]LayoutOverlay**](LayoutOverlay.md) |  | [optional] 

## Methods

### NewProjectLayoutDetails

`func NewProjectLayoutDetails() *ProjectLayoutDetails`

NewProjectLayoutDetails instantiates a new ProjectLayoutDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectLayoutDetailsWithDefaults

`func NewProjectLayoutDetailsWithDefaults() *ProjectLayoutDetails`

NewProjectLayoutDetailsWithDefaults instantiates a new ProjectLayoutDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *ProjectLayoutDetails) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *ProjectLayoutDetails) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *ProjectLayoutDetails) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *ProjectLayoutDetails) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetProjectLayout

`func (o *ProjectLayoutDetails) GetProjectLayout() ProjectLayout`

GetProjectLayout returns the ProjectLayout field if non-nil, zero value otherwise.

### GetProjectLayoutOk

`func (o *ProjectLayoutDetails) GetProjectLayoutOk() (*ProjectLayout, bool)`

GetProjectLayoutOk returns a tuple with the ProjectLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectLayout

`func (o *ProjectLayoutDetails) SetProjectLayout(v ProjectLayout)`

SetProjectLayout sets ProjectLayout field to given value.

### HasProjectLayout

`func (o *ProjectLayoutDetails) HasProjectLayout() bool`

HasProjectLayout returns a boolean if a field has been set.

### GetLayoutId

`func (o *ProjectLayoutDetails) GetLayoutId() string`

GetLayoutId returns the LayoutId field if non-nil, zero value otherwise.

### GetLayoutIdOk

`func (o *ProjectLayoutDetails) GetLayoutIdOk() (*string, bool)`

GetLayoutIdOk returns a tuple with the LayoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutId

`func (o *ProjectLayoutDetails) SetLayoutId(v string)`

SetLayoutId sets LayoutId field to given value.

### HasLayoutId

`func (o *ProjectLayoutDetails) HasLayoutId() bool`

HasLayoutId returns a boolean if a field has been set.

### GetOverlays

`func (o *ProjectLayoutDetails) GetOverlays() []LayoutOverlay`

GetOverlays returns the Overlays field if non-nil, zero value otherwise.

### GetOverlaysOk

`func (o *ProjectLayoutDetails) GetOverlaysOk() (*[]LayoutOverlay, bool)`

GetOverlaysOk returns a tuple with the Overlays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlays

`func (o *ProjectLayoutDetails) SetOverlays(v []LayoutOverlay)`

SetOverlays sets Overlays field to given value.

### HasOverlays

`func (o *ProjectLayoutDetails) HasOverlays() bool`

HasOverlays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


