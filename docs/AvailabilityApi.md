# \AvailabilityApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAvailability**](AvailabilityApi.md#CreateAvailability) | **Post** /persons/{personId}/availabilities | Create a availability for a person
[**DeleteAvailability**](AvailabilityApi.md#DeleteAvailability) | **Delete** /persons/{personId}/availabilities/{availabilityId} | Delete a person availability
[**GetAvailabilities**](AvailabilityApi.md#GetAvailabilities) | **Get** /persons/{personId}/availabilities | Get a list of all activities for a person
[**UpdateAvailability**](AvailabilityApi.md#UpdateAvailability) | **Put** /persons/{personId}/availabilities/{availabilityId} | Update a person availability



## CreateAvailability

> AvailabilityDetail CreateAvailability(ctx, personId).Availability(availability).Execute()

Create a availability for a person

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    availability := *openapiclient.NewAvailability(float32(38.5), float32(22.5), "Id_example", time.Now(), time.Now()) // Availability | The availability

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvailabilityApi.CreateAvailability(context.Background(), personId).Availability(availability).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvailabilityApi.CreateAvailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAvailability`: AvailabilityDetail
    fmt.Fprintf(os.Stdout, "Response from `AvailabilityApi.CreateAvailability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **availability** | [**Availability**](Availability.md) | The availability | 

### Return type

[**AvailabilityDetail**](AvailabilityDetail.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAvailability

> Status DeleteAvailability(ctx, personId, availabilityId).Execute()

Delete a person availability

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    availabilityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvailabilityApi.DeleteAvailability(context.Background(), personId, availabilityId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvailabilityApi.DeleteAvailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAvailability`: Status
    fmt.Fprintf(os.Stdout, "Response from `AvailabilityApi.DeleteAvailability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**availabilityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Status**](Status.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailabilities

> PagedAvailabilities GetAvailabilities(ctx, personId).Execute()

Get a list of all activities for a person

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvailabilityApi.GetAvailabilities(context.Background(), personId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvailabilityApi.GetAvailabilities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvailabilities`: PagedAvailabilities
    fmt.Fprintf(os.Stdout, "Response from `AvailabilityApi.GetAvailabilities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PagedAvailabilities**](PagedAvailabilities.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAvailability

> AvailabilityDetail UpdateAvailability(ctx, personId, availabilityId).Availability(availability).Execute()

Update a person availability

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    availabilityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    availability := *openapiclient.NewAvailability(float32(38.5), float32(22.5), "Id_example", time.Now(), time.Now()) // Availability | The availability

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvailabilityApi.UpdateAvailability(context.Background(), personId, availabilityId).Availability(availability).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvailabilityApi.UpdateAvailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAvailability`: AvailabilityDetail
    fmt.Fprintf(os.Stdout, "Response from `AvailabilityApi.UpdateAvailability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**availabilityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **availability** | [**Availability**](Availability.md) | The availability | 

### Return type

[**AvailabilityDetail**](AvailabilityDetail.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

