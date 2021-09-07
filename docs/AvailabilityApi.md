# \AvailabilityApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAvailability**](AvailabilityApi.md#CreateAvailability) | **Post** /persons/{personId}/availabilities | Create a availability for a person
[**DeleteAvailability**](AvailabilityApi.md#DeleteAvailability) | **Delete** /person/{personId}/availabilities/{availabilityId} | Delete a person availability
[**GetAvailabilities**](AvailabilityApi.md#GetAvailabilities) | **Get** /persons/{personId}/availabilities | Get a list of all activities for a person
[**GetAvailabilitiesCalculated**](AvailabilityApi.md#GetAvailabilitiesCalculated) | **Get** /persons/{personId}/availabilities/calculated | Get a list of all activities for a person for a given time persion, calculated on server side
[**UpdateAvailability**](AvailabilityApi.md#UpdateAvailability) | **Put** /person/{personId}/availabilities/{availabilityId} | Update a person availability



## CreateAvailability

> PersonDetails CreateAvailability(ctx, personId).Availability(availability).Execute()

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
    personId := TODO // string | 
    availability := *openapiclient.NewAvailability("Id_example", "This is the name", time.Now(), time.Now(), float32(38.5), float32(22.5)) // Availability | The availability

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AvailabilityApi.CreateAvailability(context.Background(), personId).Availability(availability).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvailabilityApi.CreateAvailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAvailability`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `AvailabilityApi.CreateAvailability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **availability** | [**Availability**](Availability.md) | The availability | 

### Return type

[**PersonDetails**](PersonDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAvailability

> AvailabilityDetail DeleteAvailability(ctx, personId, availabilityId).Execute()

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
    personId := TODO // string | 
    availabilityId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AvailabilityApi.DeleteAvailability(context.Background(), personId, availabilityId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvailabilityApi.DeleteAvailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAvailability`: AvailabilityDetail
    fmt.Fprintf(os.Stdout, "Response from `AvailabilityApi.DeleteAvailability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | [**string**](.md) |  | 
**availabilityId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AvailabilityDetail**](AvailabilityDetail.md)

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
    personId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AvailabilityApi.GetAvailabilities(context.Background(), personId).Execute()
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
**personId** | [**string**](.md) |  | 

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


## GetAvailabilitiesCalculated

> PagedAvailabilities GetAvailabilitiesCalculated(ctx, personId).StartDate(startDate).PeriodDays(periodDays).NumberOfPeriods(numberOfPeriods).Execute()

Get a list of all activities for a person for a given time persion, calculated on server side

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
    personId := TODO // string | 
    startDate := time.Now() // string | 
    periodDays := int32(7) // int32 |  (optional) (default to 7)
    numberOfPeriods := int32(10) // int32 |  (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AvailabilityApi.GetAvailabilitiesCalculated(context.Background(), personId).StartDate(startDate).PeriodDays(periodDays).NumberOfPeriods(numberOfPeriods).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvailabilityApi.GetAvailabilitiesCalculated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvailabilitiesCalculated`: PagedAvailabilities
    fmt.Fprintf(os.Stdout, "Response from `AvailabilityApi.GetAvailabilitiesCalculated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailabilitiesCalculatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** |  | 
 **periodDays** | **int32** |  | [default to 7]
 **numberOfPeriods** | **int32** |  | [default to 10]

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
    personId := TODO // string | 
    availabilityId := TODO // string | 
    availability := *openapiclient.NewAvailability("Id_example", "This is the name", time.Now(), time.Now(), float32(38.5), float32(22.5)) // Availability | The availability

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AvailabilityApi.UpdateAvailability(context.Background(), personId, availabilityId).Availability(availability).Execute()
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
**personId** | [**string**](.md) |  | 
**availabilityId** | [**string**](.md) |  | 

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
