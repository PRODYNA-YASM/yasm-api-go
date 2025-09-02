# \PDFProfileAPI

All URIs are relative to *https://yasm.prodyna.com:443/api/graph/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCreatePdfProfiles**](PDFProfileAPI.md#BulkCreatePdfProfiles) | **Post** /pdf-profile/bulk-create | Generate PDF profiles wrapped in a ZIP File from a list of Persons
[**CreatePdfProfile**](PDFProfileAPI.md#CreatePdfProfile) | **Post** /pdf-profile/create | Generate a PDF profile from a Person
[**PublishPdfProfile**](PDFProfileAPI.md#PublishPdfProfile) | **Post** /pdf-profile/publish | Generate a PDF profile from a Person



## BulkCreatePdfProfiles

> *os.File BulkCreatePdfProfiles(ctx).ProfileRequests(profileRequests).Execute()

Generate PDF profiles wrapped in a ZIP File from a list of Persons

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
    profileRequests := *openapiclient.NewProfileRequests() // ProfileRequests | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PDFProfileAPI.BulkCreatePdfProfiles(context.Background()).ProfileRequests(profileRequests).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PDFProfileAPI.BulkCreatePdfProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreatePdfProfiles`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `PDFProfileAPI.BulkCreatePdfProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreatePdfProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileRequests** | [**ProfileRequests**](ProfileRequests.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePdfProfile

> *os.File CreatePdfProfile(ctx).ProfileRequest(profileRequest).Execute()

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
    resp, r, err := apiClient.PDFProfileAPI.CreatePdfProfile(context.Background()).ProfileRequest(profileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PDFProfileAPI.CreatePdfProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePdfProfile`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `PDFProfileAPI.CreatePdfProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePdfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileRequest** | [**ProfileRequest**](ProfileRequest.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishPdfProfile

> PdfProfile PublishPdfProfile(ctx).ProfileRequest(profileRequest).Execute()

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
    resp, r, err := apiClient.PDFProfileAPI.PublishPdfProfile(context.Background()).ProfileRequest(profileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PDFProfileAPI.PublishPdfProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublishPdfProfile`: PdfProfile
    fmt.Fprintf(os.Stdout, "Response from `PDFProfileAPI.PublishPdfProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublishPdfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileRequest** | [**ProfileRequest**](ProfileRequest.md) |  | 

### Return type

[**PdfProfile**](PdfProfile.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

