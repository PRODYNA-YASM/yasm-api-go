# \IndustryApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIndustry**](IndustryApi.md#CreateIndustry) | **Post** /industries | Create an Industry
[**DeleteIndustry**](IndustryApi.md#DeleteIndustry) | **Delete** /industries/{industryId} | Delete an Industry
[**GetIndustries**](IndustryApi.md#GetIndustries) | **Get** /industries | Get all Industries
[**GetIndustry**](IndustryApi.md#GetIndustry) | **Get** /industries/{industryId} | Get details about an Industry
[**UpdateIndustry**](IndustryApi.md#UpdateIndustry) | **Put** /industries/{industryId} | Update an Industry



## CreateIndustry

> Industry CreateIndustry(ctx).Industry(industry).Execute()

Create an Industry

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
    industry := *openapiclient.NewIndustry("Id_example", interface{}(Aviation)) // Industry | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IndustryApi.CreateIndustry(context.Background()).Industry(industry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndustryApi.CreateIndustry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIndustry`: Industry
    fmt.Fprintf(os.Stdout, "Response from `IndustryApi.CreateIndustry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndustryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **industry** | [**Industry**](Industry.md) |  | 

### Return type

[**Industry**](Industry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIndustry

> Status DeleteIndustry(ctx, industryId).Execute()

Delete an Industry

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
    industryId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IndustryApi.DeleteIndustry(context.Background(), industryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndustryApi.DeleteIndustry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIndustry`: Status
    fmt.Fprintf(os.Stdout, "Response from `IndustryApi.DeleteIndustry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**industryId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndustryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Status**](Status.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndustries

> PagedIndustries GetIndustries(ctx).Skip(skip).Limit(limit).Execute()

Get all Industries

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
    skip := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IndustryApi.GetIndustries(context.Background()).Skip(skip).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndustryApi.GetIndustries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIndustries`: PagedIndustries
    fmt.Fprintf(os.Stdout, "Response from `IndustryApi.GetIndustries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIndustriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**PagedIndustries**](PagedIndustries.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndustry

> Industry GetIndustry(ctx, industryId).Execute()

Get details about an Industry

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
    industryId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IndustryApi.GetIndustry(context.Background(), industryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndustryApi.GetIndustry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIndustry`: Industry
    fmt.Fprintf(os.Stdout, "Response from `IndustryApi.GetIndustry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**industryId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndustryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Industry**](Industry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndustry

> Industry UpdateIndustry(ctx, industryId).Industry(industry).Execute()

Update an Industry

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
    industryId := TODO // string | 
    industry := *openapiclient.NewIndustry("Id_example", interface{}(Aviation)) // Industry | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IndustryApi.UpdateIndustry(context.Background(), industryId).Industry(industry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndustryApi.UpdateIndustry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndustry`: Industry
    fmt.Fprintf(os.Stdout, "Response from `IndustryApi.UpdateIndustry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**industryId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndustryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **industry** | [**Industry**](Industry.md) |  | 

### Return type

[**Industry**](Industry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
