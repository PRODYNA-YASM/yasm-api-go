# \ProfileAPI

All URIs are relative to *https://yasm.prodyna.com:443/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPersonProfile**](ProfileAPI.md#AddPersonProfile) | **Post** /persons/{personId}/profiles/{profileId} | Add a Profile to a Person
[**CreateProfile**](ProfileAPI.md#CreateProfile) | **Post** /profiles | Create a new Profile
[**DeletePersonProfile**](ProfileAPI.md#DeletePersonProfile) | **Delete** /persons/{personId}/profiles/{profileId} | Remove a Profile from a Person
[**DeleteProfile**](ProfileAPI.md#DeleteProfile) | **Delete** /profiles/{profileId} | Delete a Profile
[**GetProfile**](ProfileAPI.md#GetProfile) | **Get** /profiles/{profileId} | Get details about a Profile
[**UpdateProfile**](ProfileAPI.md#UpdateProfile) | **Put** /profiles/{profileId} | Update a Profile



## AddPersonProfile

> PersonDetails AddPersonProfile(ctx, personId, profileId).Execute()

Add a Profile to a Person

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
    profileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileAPI.AddPersonProfile(context.Background(), personId, profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileAPI.AddPersonProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPersonProfile`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `ProfileAPI.AddPersonProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**profileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPersonProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PersonDetails**](PersonDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProfile

> ProfileDetails CreateProfile(ctx).Profile(profile).Execute()

Create a new Profile

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
    profile := *openapiclient.NewProfile("Id_example", "Name_example") // Profile | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileAPI.CreateProfile(context.Background()).Profile(profile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileAPI.CreateProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProfile`: ProfileDetails
    fmt.Fprintf(os.Stdout, "Response from `ProfileAPI.CreateProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profile** | [**Profile**](Profile.md) |  | 

### Return type

[**ProfileDetails**](ProfileDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePersonProfile

> PersonDetails DeletePersonProfile(ctx, personId, profileId).Execute()

Remove a Profile from a Person

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
    profileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileAPI.DeletePersonProfile(context.Background(), personId, profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileAPI.DeletePersonProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePersonProfile`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `ProfileAPI.DeletePersonProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**profileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePersonProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PersonDetails**](PersonDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProfile

> Status DeleteProfile(ctx, profileId).Execute()

Delete a Profile

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
    profileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileAPI.DeleteProfile(context.Background(), profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileAPI.DeleteProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteProfile`: Status
    fmt.Fprintf(os.Stdout, "Response from `ProfileAPI.DeleteProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProfileRequest struct via the builder pattern


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


## GetProfile

> ProfileDetails GetProfile(ctx, profileId).Execute()

Get details about a Profile

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
    profileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileAPI.GetProfile(context.Background(), profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileAPI.GetProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfile`: ProfileDetails
    fmt.Fprintf(os.Stdout, "Response from `ProfileAPI.GetProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProfileDetails**](ProfileDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProfile

> ProfileDetails UpdateProfile(ctx, profileId).Profile(profile).Execute()

Update a Profile

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
    profileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    profile := *openapiclient.NewProfile("Id_example", "Name_example") // Profile | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileAPI.UpdateProfile(context.Background(), profileId).Profile(profile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileAPI.UpdateProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProfile`: ProfileDetails
    fmt.Fprintf(os.Stdout, "Response from `ProfileAPI.UpdateProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **profile** | [**Profile**](Profile.md) |  | 

### Return type

[**ProfileDetails**](ProfileDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

