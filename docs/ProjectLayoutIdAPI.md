# \ProjectLayoutIdAPI

All URIs are relative to *https://yasm.prodyna.com:443/api/graph/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchProjectLayouts**](ProjectLayoutIdAPI.md#SearchProjectLayouts) | **Post** /projects/{projectId}/layouts/search | Search for Project Layouts



## SearchProjectLayouts

> PagedProjectLayouts SearchProjectLayouts(ctx, projectId).ProjectLayoutSearch(projectLayoutSearch).Execute()

Search for Project Layouts

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
    projectLayoutSearch := *openapiclient.NewProjectLayoutSearch(int32(123), int32(123)) // ProjectLayoutSearch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectLayoutIdAPI.SearchProjectLayouts(context.Background(), projectId).ProjectLayoutSearch(projectLayoutSearch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectLayoutIdAPI.SearchProjectLayouts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchProjectLayouts`: PagedProjectLayouts
    fmt.Fprintf(os.Stdout, "Response from `ProjectLayoutIdAPI.SearchProjectLayouts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchProjectLayoutsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectLayoutSearch** | [**ProjectLayoutSearch**](ProjectLayoutSearch.md) |  | 

### Return type

[**PagedProjectLayouts**](PagedProjectLayouts.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

