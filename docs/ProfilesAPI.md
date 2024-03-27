# \ProfilesAPI

All URIs are relative to *https://yasm.prodyna.com:443/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchProfiles**](ProfilesAPI.md#SearchProfiles) | **Post** /profiles/search | Search over profiles



## SearchProfiles

> PagedProfiles SearchProfiles(ctx).Search(search).Execute()

Search over profiles

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
    resp, r, err := apiClient.ProfilesAPI.SearchProfiles(context.Background()).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.SearchProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchProfiles`: PagedProfiles
    fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.SearchProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | [**Search**](Search.md) |  | 

### Return type

[**PagedProfiles**](PagedProfiles.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

