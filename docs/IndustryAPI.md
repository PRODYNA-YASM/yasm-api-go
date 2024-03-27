# \IndustryAPI

All URIs are relative to *https://yasm.prodyna.com:443/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachOrganizationToIndustry**](IndustryAPI.md#AttachOrganizationToIndustry) | **Post** /organizations/{organizationId}/industries/{industryId} | Add an Organization to an Industry
[**CreateIndustry**](IndustryAPI.md#CreateIndustry) | **Post** /industries | Create an Industry
[**DeleteIndustry**](IndustryAPI.md#DeleteIndustry) | **Delete** /industries/{industryId} | Delete an Industry
[**DetachOrganizationFromIndustry**](IndustryAPI.md#DetachOrganizationFromIndustry) | **Delete** /organizations/{organizationId}/industries/{industryId} | Remove an Organization to an Industry
[**GetIndustry**](IndustryAPI.md#GetIndustry) | **Get** /industries/{industryId} | Get details about an Industry
[**UpdateIndustry**](IndustryAPI.md#UpdateIndustry) | **Put** /industries/{industryId} | Update an Industry



## AttachOrganizationToIndustry

> OrganizationDetails AttachOrganizationToIndustry(ctx, organizationId, industryId).Execute()

Add an Organization to an Industry

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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    industryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndustryAPI.AttachOrganizationToIndustry(context.Background(), organizationId, industryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndustryAPI.AttachOrganizationToIndustry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachOrganizationToIndustry`: OrganizationDetails
    fmt.Fprintf(os.Stdout, "Response from `IndustryAPI.AttachOrganizationToIndustry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**industryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachOrganizationToIndustryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationDetails**](OrganizationDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIndustry

> IndustryDetails CreateIndustry(ctx).Industry(industry).Execute()

Create an Industry

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
    industry := *openapiclient.NewIndustry(false, "Id_example", "Name_example") // Industry | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndustryAPI.CreateIndustry(context.Background()).Industry(industry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndustryAPI.CreateIndustry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIndustry`: IndustryDetails
    fmt.Fprintf(os.Stdout, "Response from `IndustryAPI.CreateIndustry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndustryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **industry** | [**Industry**](Industry.md) |  | 

### Return type

[**IndustryDetails**](IndustryDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIndustry

> Status DeleteIndustry(ctx, industryId).Execute()

Delete an Industry

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
    industryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndustryAPI.DeleteIndustry(context.Background(), industryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndustryAPI.DeleteIndustry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIndustry`: Status
    fmt.Fprintf(os.Stdout, "Response from `IndustryAPI.DeleteIndustry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**industryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndustryRequest struct via the builder pattern


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


## DetachOrganizationFromIndustry

> OrganizationDetails DetachOrganizationFromIndustry(ctx, organizationId, industryId).Execute()

Remove an Organization to an Industry

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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    industryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndustryAPI.DetachOrganizationFromIndustry(context.Background(), organizationId, industryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndustryAPI.DetachOrganizationFromIndustry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DetachOrganizationFromIndustry`: OrganizationDetails
    fmt.Fprintf(os.Stdout, "Response from `IndustryAPI.DetachOrganizationFromIndustry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**industryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachOrganizationFromIndustryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationDetails**](OrganizationDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndustry

> IndustryDetails GetIndustry(ctx, industryId).Execute()

Get details about an Industry

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
    industryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndustryAPI.GetIndustry(context.Background(), industryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndustryAPI.GetIndustry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIndustry`: IndustryDetails
    fmt.Fprintf(os.Stdout, "Response from `IndustryAPI.GetIndustry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**industryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndustryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IndustryDetails**](IndustryDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndustry

> IndustryDetails UpdateIndustry(ctx, industryId).Industry(industry).Execute()

Update an Industry

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
    industryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    industry := *openapiclient.NewIndustry(false, "Id_example", "Name_example") // Industry | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndustryAPI.UpdateIndustry(context.Background(), industryId).Industry(industry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndustryAPI.UpdateIndustry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndustry`: IndustryDetails
    fmt.Fprintf(os.Stdout, "Response from `IndustryAPI.UpdateIndustry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**industryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndustryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **industry** | [**Industry**](Industry.md) |  | 

### Return type

[**IndustryDetails**](IndustryDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

