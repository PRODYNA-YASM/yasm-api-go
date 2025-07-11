# \ProjectLayoutAPI

All URIs are relative to *https://yasm.prodyna.com:443/api/graph/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLayoutOverlay**](ProjectLayoutAPI.md#CreateLayoutOverlay) | **Post** /projects/{projectId}/layouts/{projectLayoutId}/overlay | Create a Layout Overlay
[**CreateProjectLayout**](ProjectLayoutAPI.md#CreateProjectLayout) | **Post** /projects/{projectId}/layouts | Create a Project Layout
[**DeleteOverlayFromLayoutInProject**](ProjectLayoutAPI.md#DeleteOverlayFromLayoutInProject) | **Delete** /projects/{projectId}/images/{projectLayoutId}/overlay/{overlayId} | Delete an overlay from a layout in a project
[**RenderOverlayToPng**](ProjectLayoutAPI.md#RenderOverlayToPng) | **Get** /projects/{projectId}/layouts/{projectLayoutId}.png | Render a Layout Overlay to png



## CreateLayoutOverlay

> ProjectLayoutDetails CreateLayoutOverlay(ctx, projectId, projectLayoutId).FieldId(fieldId).PngFile(pngFile).Execute()

Create a Layout Overlay

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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    projectLayoutId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    fieldId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    pngFile := os.NewFile(1234, "some_file") // *os.File | png layout file

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectLayoutAPI.CreateLayoutOverlay(context.Background(), projectId, projectLayoutId).FieldId(fieldId).PngFile(pngFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectLayoutAPI.CreateLayoutOverlay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLayoutOverlay`: ProjectLayoutDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectLayoutAPI.CreateLayoutOverlay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**projectLayoutId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLayoutOverlayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fieldId** | **string** |  | 
 **pngFile** | ***os.File** | png layout file | 

### Return type

[**ProjectLayoutDetails**](ProjectLayoutDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProjectLayout

> ProjectLayoutDetails CreateProjectLayout(ctx, projectId).LayoutId(layoutId).ProjectLayout(projectLayout).Execute()

Create a Project Layout

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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    layoutId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the layout to use for the image
    projectLayout := *openapiclient.NewProjectLayout("Id_example", "Name_example") // ProjectLayout | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectLayoutAPI.CreateProjectLayout(context.Background(), projectId).LayoutId(layoutId).ProjectLayout(projectLayout).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectLayoutAPI.CreateProjectLayout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProjectLayout`: ProjectLayoutDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectLayoutAPI.CreateProjectLayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectLayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **layoutId** | **string** | The ID of the layout to use for the image | 
 **projectLayout** | [**ProjectLayout**](ProjectLayout.md) |  | 

### Return type

[**ProjectLayoutDetails**](ProjectLayoutDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOverlayFromLayoutInProject

> Status DeleteOverlayFromLayoutInProject(ctx, projectId, projectLayoutId, overlayId).Execute()

Delete an overlay from a layout in a project

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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    projectLayoutId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    overlayId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectLayoutAPI.DeleteOverlayFromLayoutInProject(context.Background(), projectId, projectLayoutId, overlayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectLayoutAPI.DeleteOverlayFromLayoutInProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOverlayFromLayoutInProject`: Status
    fmt.Fprintf(os.Stdout, "Response from `ProjectLayoutAPI.DeleteOverlayFromLayoutInProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**projectLayoutId** | **string** |  | 
**overlayId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOverlayFromLayoutInProjectRequest struct via the builder pattern


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


## RenderOverlayToPng

> *os.File RenderOverlayToPng(ctx, projectId, projectLayoutId).Execute()

Render a Layout Overlay to png

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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    projectLayoutId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectLayoutAPI.RenderOverlayToPng(context.Background(), projectId, projectLayoutId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectLayoutAPI.RenderOverlayToPng``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RenderOverlayToPng`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ProjectLayoutAPI.RenderOverlayToPng`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**projectLayoutId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenderOverlayToPngRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

