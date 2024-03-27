# \IndustriesAPI

All URIs are relative to *https://yasm.prodyna.com:443/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchIndustries**](IndustriesAPI.md#SearchIndustries) | **Post** /industries/search | Search over industries



## SearchIndustries

> PagedIndustries SearchIndustries(ctx).Search(search).Execute()

Search over industries

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
    resp, r, err := apiClient.IndustriesAPI.SearchIndustries(context.Background()).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndustriesAPI.SearchIndustries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchIndustries`: PagedIndustries
    fmt.Fprintf(os.Stdout, "Response from `IndustriesAPI.SearchIndustries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchIndustriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | [**Search**](Search.md) |  | 

### Return type

[**PagedIndustries**](PagedIndustries.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

