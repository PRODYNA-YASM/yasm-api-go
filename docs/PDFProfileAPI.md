# \PDFProfileAPI

All URIs are relative to *https://yasm.prodyna.com:443/api/graph/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GeneratePdfProfile**](PDFProfileAPI.md#GeneratePdfProfile) | **Post** /pdf-profile | Generate a PDF profile from a Person



## GeneratePdfProfile

> *os.File GeneratePdfProfile(ctx).ProfileRequest(profileRequest).Execute()

Generate a PDF profile from a Person

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
    profileRequest := *openapiclient.NewProfileRequest() // ProfileRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PDFProfileAPI.GeneratePdfProfile(context.Background()).ProfileRequest(profileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PDFProfileAPI.GeneratePdfProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneratePdfProfile`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `PDFProfileAPI.GeneratePdfProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGeneratePdfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileRequest** | [**ProfileRequest**](ProfileRequest.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

