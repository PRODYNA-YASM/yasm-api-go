# ProfileDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audit** | Pointer to [**Audit**](Audit.md) |  | [optional] 
**Profile** | Pointer to [**Profile**](Profile.md) |  | [optional] 

## Methods

### NewProfileDetails

`func NewProfileDetails() *ProfileDetails`

NewProfileDetails instantiates a new ProfileDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileDetailsWithDefaults

`func NewProfileDetailsWithDefaults() *ProfileDetails`

NewProfileDetailsWithDefaults instantiates a new ProfileDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudit

`func (o *ProfileDetails) GetAudit() Audit`

GetAudit returns the Audit field if non-nil, zero value otherwise.

### GetAuditOk

`func (o *ProfileDetails) GetAuditOk() (*Audit, bool)`

GetAuditOk returns a tuple with the Audit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudit

`func (o *ProfileDetails) SetAudit(v Audit)`

SetAudit sets Audit field to given value.

### HasAudit

`func (o *ProfileDetails) HasAudit() bool`

HasAudit returns a boolean if a field has been set.

### GetProfile

`func (o *ProfileDetails) GetProfile() Profile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ProfileDetails) GetProfileOk() (*Profile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ProfileDetails) SetProfile(v Profile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ProfileDetails) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


