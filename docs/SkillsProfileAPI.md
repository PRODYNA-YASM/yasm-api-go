# \SkillsProfileAPI

All URIs are relative to *https://yasm.prodyna.com:443/api/graph/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSkillsProfileForPerson**](SkillsProfileAPI.md#CreateSkillsProfileForPerson) | **Post** /persons/{personId}/skills-profiles | Create a SkillsProfile
[**DeleteSkillsProfile**](SkillsProfileAPI.md#DeleteSkillsProfile) | **Delete** /skillsProfiles/{skillsProfileId} | Delete a skills profile
[**GetPersonSkillsProfiles**](SkillsProfileAPI.md#GetPersonSkillsProfiles) | **Get** /persons/{personId}/skills-profiles | Get all SkillsProfiles of a single person
[**GetSkillsProfile**](SkillsProfileAPI.md#GetSkillsProfile) | **Get** /skillsProfiles/{skillsProfileId} | Get a SkillsProfile
[**UpdateSkillsProfile**](SkillsProfileAPI.md#UpdateSkillsProfile) | **Put** /skillsProfiles/{skillsProfileId} | Update a SkillsProfile



## CreateSkillsProfileForPerson

> SkillsProfileDetails CreateSkillsProfileForPerson(ctx, personId).SkillsProfileRequest(skillsProfileRequest).Execute()

Create a SkillsProfile

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillsProfileRequest := *openapiclient.NewSkillsProfileRequest() // SkillsProfileRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SkillsProfileAPI.CreateSkillsProfileForPerson(context.Background(), personId).SkillsProfileRequest(skillsProfileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillsProfileAPI.CreateSkillsProfileForPerson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSkillsProfileForPerson`: SkillsProfileDetails
    fmt.Fprintf(os.Stdout, "Response from `SkillsProfileAPI.CreateSkillsProfileForPerson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSkillsProfileForPersonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skillsProfileRequest** | [**SkillsProfileRequest**](SkillsProfileRequest.md) |  | 

### Return type

[**SkillsProfileDetails**](SkillsProfileDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSkillsProfile

> Status DeleteSkillsProfile(ctx, skillsProfileId).Execute()

Delete a skills profile

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    skillsProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SkillsProfileAPI.DeleteSkillsProfile(context.Background(), skillsProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillsProfileAPI.DeleteSkillsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSkillsProfile`: Status
    fmt.Fprintf(os.Stdout, "Response from `SkillsProfileAPI.DeleteSkillsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillsProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSkillsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Status**](Status.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPersonSkillsProfiles

> PagedSkillsProfiles GetPersonSkillsProfiles(ctx, personId).Execute()

Get all SkillsProfiles of a single person

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SkillsProfileAPI.GetPersonSkillsProfiles(context.Background(), personId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillsProfileAPI.GetPersonSkillsProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPersonSkillsProfiles`: PagedSkillsProfiles
    fmt.Fprintf(os.Stdout, "Response from `SkillsProfileAPI.GetPersonSkillsProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonSkillsProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PagedSkillsProfiles**](PagedSkillsProfiles.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSkillsProfile

> SkillsProfileDetails GetSkillsProfile(ctx, skillsProfileId).Execute()

Get a SkillsProfile

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    skillsProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SkillsProfileAPI.GetSkillsProfile(context.Background(), skillsProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillsProfileAPI.GetSkillsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSkillsProfile`: SkillsProfileDetails
    fmt.Fprintf(os.Stdout, "Response from `SkillsProfileAPI.GetSkillsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillsProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSkillsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SkillsProfileDetails**](SkillsProfileDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSkillsProfile

> SkillsProfileDetails UpdateSkillsProfile(ctx, skillsProfileId).SkillsProfileRequest(skillsProfileRequest).Execute()

Update a SkillsProfile

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    skillsProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillsProfileRequest := *openapiclient.NewSkillsProfileRequest() // SkillsProfileRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SkillsProfileAPI.UpdateSkillsProfile(context.Background(), skillsProfileId).SkillsProfileRequest(skillsProfileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillsProfileAPI.UpdateSkillsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSkillsProfile`: SkillsProfileDetails
    fmt.Fprintf(os.Stdout, "Response from `SkillsProfileAPI.UpdateSkillsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillsProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSkillsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skillsProfileRequest** | [**SkillsProfileRequest**](SkillsProfileRequest.md) |  | 

### Return type

[**SkillsProfileDetails**](SkillsProfileDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

