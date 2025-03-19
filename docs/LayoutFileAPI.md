# \LayoutFileAPI

All URIs are relative to *https://yasm.prodyna.com:443/api/graph/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLayoutFile**](LayoutFileAPI.md#DeleteLayoutFile) | **Delete** /layout/{layoutFileId}/file | Delete a Layout SVG file by layoutFileId
[**ReadLayoutFile**](LayoutFileAPI.md#ReadLayoutFile) | **Get** /layout/{layoutFileId}/file | read a Layout SVG file
[**SearchLayoutFileDetails**](LayoutFileAPI.md#SearchLayoutFileDetails) | **Post** /layout/file/search | Search over Layout File
[**UploadLayoutFile**](LayoutFileAPI.md#UploadLayoutFile) | **Post** /layout/{layoutId}/file/upload | Upload a Layout SVG file behind a layoutId



## DeleteLayoutFile

> Status DeleteLayoutFile(ctx, layoutFileId).Execute()

Delete a Layout SVG file by layoutFileId

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
    layoutFileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LayoutFileAPI.DeleteLayoutFile(context.Background(), layoutFileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LayoutFileAPI.DeleteLayoutFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLayoutFile`: Status
    fmt.Fprintf(os.Stdout, "Response from `LayoutFileAPI.DeleteLayoutFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**layoutFileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLayoutFileRequest struct via the builder pattern


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


## ReadLayoutFile

> *os.File ReadLayoutFile(ctx, layoutFileId).Execute()

read a Layout SVG file

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
    layoutFileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LayoutFileAPI.ReadLayoutFile(context.Background(), layoutFileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LayoutFileAPI.ReadLayoutFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadLayoutFile`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `LayoutFileAPI.ReadLayoutFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**layoutFileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadLayoutFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/svg+xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchLayoutFileDetails

> PagedLayoutFileDetails SearchLayoutFileDetails(ctx).LayoutFileDetailsSearch(layoutFileDetailsSearch).Execute()

Search over Layout File

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
    layoutFileDetailsSearch := *openapiclient.NewLayoutFileDetailsSearch(int32(123), int32(123)) // LayoutFileDetailsSearch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LayoutFileAPI.SearchLayoutFileDetails(context.Background()).LayoutFileDetailsSearch(layoutFileDetailsSearch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LayoutFileAPI.SearchLayoutFileDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchLayoutFileDetails`: PagedLayoutFileDetails
    fmt.Fprintf(os.Stdout, "Response from `LayoutFileAPI.SearchLayoutFileDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchLayoutFileDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **layoutFileDetailsSearch** | [**LayoutFileDetailsSearch**](LayoutFileDetailsSearch.md) |  | 

### Return type

[**PagedLayoutFileDetails**](PagedLayoutFileDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadLayoutFile

> LayoutFileDetails UploadLayoutFile(ctx, layoutId).Body(body).Execute()

Upload a Layout SVG file behind a layoutId

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
    body := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LayoutFileAPI.UploadLayoutFile(context.Background(), layoutId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LayoutFileAPI.UploadLayoutFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadLayoutFile`: LayoutFileDetails
    fmt.Fprintf(os.Stdout, "Response from `LayoutFileAPI.UploadLayoutFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**layoutId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadLayoutFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** |  | 

### Return type

[**LayoutFileDetails**](LayoutFileDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

