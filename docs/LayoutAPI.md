# \LayoutAPI

All URIs are relative to *https://yasm.prodyna.com:443/api/graph/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLayout**](LayoutAPI.md#CreateLayout) | **Post** /layouts | Create a Layout
[**DeleteLayout**](LayoutAPI.md#DeleteLayout) | **Delete** /layouts/{layoutId} | Delete a Layout
[**GetLayout**](LayoutAPI.md#GetLayout) | **Get** /layouts/{layoutId} | Read a Layout
[**SearchLayouts**](LayoutAPI.md#SearchLayouts) | **Post** /layouts/search | Search over Layouts



## CreateLayout

> LayoutDetails CreateLayout(ctx).SvgFile(svgFile).Id(id).Name(name).Execute()

Create a Layout

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
    svgFile := os.NewFile(1234, "some_file") // *os.File | SVG layout file
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    name := "name_example" // string | Layout JSON object (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LayoutAPI.CreateLayout(context.Background()).SvgFile(svgFile).Id(id).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LayoutAPI.CreateLayout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLayout`: LayoutDetails
    fmt.Fprintf(os.Stdout, "Response from `LayoutAPI.CreateLayout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svgFile** | ***os.File** | SVG layout file | 
 **id** | **string** |  | 
 **name** | **string** | Layout JSON object | 

### Return type

[**LayoutDetails**](LayoutDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLayout

> Status DeleteLayout(ctx, layoutId).Execute()

Delete a Layout

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
    layoutId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LayoutAPI.DeleteLayout(context.Background(), layoutId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LayoutAPI.DeleteLayout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLayout`: Status
    fmt.Fprintf(os.Stdout, "Response from `LayoutAPI.DeleteLayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**layoutId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLayoutRequest struct via the builder pattern


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


## GetLayout

> LayoutDetails GetLayout(ctx, layoutId).Execute()

Read a Layout

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
    layoutId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LayoutAPI.GetLayout(context.Background(), layoutId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LayoutAPI.GetLayout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLayout`: LayoutDetails
    fmt.Fprintf(os.Stdout, "Response from `LayoutAPI.GetLayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**layoutId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LayoutDetails**](LayoutDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchLayouts

> PagedLayouts SearchLayouts(ctx).Search(search).Execute()

Search over Layouts

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
    search := *openapiclient.NewSearch(int32(123), int32(123)) // Search |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LayoutAPI.SearchLayouts(context.Background()).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LayoutAPI.SearchLayouts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchLayouts`: PagedLayouts
    fmt.Fprintf(os.Stdout, "Response from `LayoutAPI.SearchLayouts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchLayoutsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | [**Search**](Search.md) |  | 

### Return type

[**PagedLayouts**](PagedLayouts.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

