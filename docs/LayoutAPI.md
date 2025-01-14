# \LayoutAPI

All URIs are relative to *https://yasm.prodyna.com:443/api/graph/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLayout**](LayoutAPI.md#CreateLayout) | **Post** /layouts | Create a Layout
[**DeleteLayout**](LayoutAPI.md#DeleteLayout) | **Delete** /layouts/{layoutId} | Delete a Layout
[**ReadLayout**](LayoutAPI.md#ReadLayout) | **Get** /layouts/{layoutId} | Read a Layout
[**SearchLayouts**](LayoutAPI.md#SearchLayouts) | **Post** /layouts/search | Search over Layouts
[**UpdateLayout**](LayoutAPI.md#UpdateLayout) | **Put** /layouts/{layoutId} | Update a Layout



## CreateLayout

> LayoutDetails CreateLayout(ctx).Layout(layout).Execute()

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
    layout := *openapiclient.NewLayout("Id_example") // Layout | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LayoutAPI.CreateLayout(context.Background()).Layout(layout).Execute()
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
 **layout** | [**Layout**](Layout.md) |  | 

### Return type

[**LayoutDetails**](LayoutDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
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


## ReadLayout

> string ReadLayout(ctx, layoutId).Execute()

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
    resp, r, err := apiClient.LayoutAPI.ReadLayout(context.Background(), layoutId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LayoutAPI.ReadLayout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadLayout`: string
    fmt.Fprintf(os.Stdout, "Response from `LayoutAPI.ReadLayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**layoutId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadLayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png, application/json

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


## UpdateLayout

> LayoutDetails UpdateLayout(ctx, layoutId).Layout(layout).Execute()

Update a Layout

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
    layout := *openapiclient.NewLayout("Id_example") // Layout | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LayoutAPI.UpdateLayout(context.Background(), layoutId).Layout(layout).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LayoutAPI.UpdateLayout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLayout`: LayoutDetails
    fmt.Fprintf(os.Stdout, "Response from `LayoutAPI.UpdateLayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**layoutId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **layout** | [**Layout**](Layout.md) |  | 

### Return type

[**LayoutDetails**](LayoutDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

