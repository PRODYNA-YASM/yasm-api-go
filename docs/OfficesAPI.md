# \OfficesAPI

All URIs are relative to *https://yasm.prodyna.com:443/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchOffices**](OfficesAPI.md#SearchOffices) | **Post** /offices/search | Search over offices



## SearchOffices

> PagedOffices SearchOffices(ctx).OfficeSearch(officeSearch).Execute()

Search over offices

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
    officeSearch := *openapiclient.NewOfficeSearch(int32(123), int32(123)) // OfficeSearch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OfficesAPI.SearchOffices(context.Background()).OfficeSearch(officeSearch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfficesAPI.SearchOffices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchOffices`: PagedOffices
    fmt.Fprintf(os.Stdout, "Response from `OfficesAPI.SearchOffices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchOfficesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **officeSearch** | [**OfficeSearch**](OfficeSearch.md) |  | 

### Return type

[**PagedOffices**](PagedOffices.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

