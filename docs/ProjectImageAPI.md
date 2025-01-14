# \ProjectImageAPI

All URIs are relative to *https://yasm.prodyna.com:443/api/graph/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProjectImage**](ProjectImageAPI.md#CreateProjectImage) | **Post** /project-images | Create a Project Image
[**DeleteProjectImage**](ProjectImageAPI.md#DeleteProjectImage) | **Delete** /project-images/{projectImageId} | Delete Project Image
[**ReadProjectImage**](ProjectImageAPI.md#ReadProjectImage) | **Get** /project-images/{projectImageId} | Read a Project Image
[**SearchProjectImages**](ProjectImageAPI.md#SearchProjectImages) | **Post** /project-images/search | Search over Project Images



## CreateProjectImage

> string CreateProjectImage(ctx).ProjectImageCreate(projectImageCreate).Execute()

Create a Project Image

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
    projectImageCreate := *openapiclient.NewProjectImageCreate("Id_example") // ProjectImageCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectImageAPI.CreateProjectImage(context.Background()).ProjectImageCreate(projectImageCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectImageAPI.CreateProjectImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProjectImage`: string
    fmt.Fprintf(os.Stdout, "Response from `ProjectImageAPI.CreateProjectImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectImageCreate** | [**ProjectImageCreate**](ProjectImageCreate.md) |  | 

### Return type

**string**

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: image/png, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProjectImage

> Status DeleteProjectImage(ctx, projectImageId).Execute()

Delete Project Image

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
    projectImageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectImageAPI.DeleteProjectImage(context.Background(), projectImageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectImageAPI.DeleteProjectImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteProjectImage`: Status
    fmt.Fprintf(os.Stdout, "Response from `ProjectImageAPI.DeleteProjectImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectImageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectImageRequest struct via the builder pattern


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


## ReadProjectImage

> string ReadProjectImage(ctx, projectImageId).Execute()

Read a Project Image

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
    projectImageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectImageAPI.ReadProjectImage(context.Background(), projectImageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectImageAPI.ReadProjectImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadProjectImage`: string
    fmt.Fprintf(os.Stdout, "Response from `ProjectImageAPI.ReadProjectImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectImageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadProjectImageRequest struct via the builder pattern


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


## SearchProjectImages

> PagedProjectImages SearchProjectImages(ctx).ProjectImageSearch(projectImageSearch).Execute()

Search over Project Images

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
    projectImageSearch := *openapiclient.NewProjectImageSearch(int32(123), int32(123)) // ProjectImageSearch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectImageAPI.SearchProjectImages(context.Background()).ProjectImageSearch(projectImageSearch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectImageAPI.SearchProjectImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchProjectImages`: PagedProjectImages
    fmt.Fprintf(os.Stdout, "Response from `ProjectImageAPI.SearchProjectImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchProjectImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectImageSearch** | [**ProjectImageSearch**](ProjectImageSearch.md) |  | 

### Return type

[**PagedProjectImages**](PagedProjectImages.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

