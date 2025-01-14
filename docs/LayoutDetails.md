# LayoutDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**Layout** | Pointer to [**Layout**](Layout.md) |  | [optional] 

## Methods

### NewLayoutDetails

`func NewLayoutDetails() *LayoutDetails`

NewLayoutDetails instantiates a new LayoutDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayoutDetailsWithDefaults

`func NewLayoutDetailsWithDefaults() *LayoutDetails`

NewLayoutDetailsWithDefaults instantiates a new LayoutDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *LayoutDetails) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *LayoutDetails) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *LayoutDetails) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *LayoutDetails) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetLayout

`func (o *LayoutDetails) GetLayout() Layout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *LayoutDetails) GetLayoutOk() (*Layout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *LayoutDetails) SetLayout(v Layout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *LayoutDetails) HasLayout() bool`

HasLayout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


