# \OfficeApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOffice**](OfficeApi.md#CreateOffice) | **Post** /organizations/{organizationId}/offices | Create an Office in an Organization
[**DeleteOffice**](OfficeApi.md#DeleteOffice) | **Delete** /organizations/{organizationId}/offices/{officeId} | Delte an Office from an Organization
[**GetOffice**](OfficeApi.md#GetOffice) | **Get** /organizations/{organizationId}/offices/{officeId} | Get an Office for an Organiaztion
[**UpdateOffice**](OfficeApi.md#UpdateOffice) | **Put** /organizations/{organizationId}/offices/{officeId} | Update an Office for an Organization



## CreateOffice

> OrganizationDetails CreateOffice(ctx, organizationId).Office(office).Execute()

Create an Office in an Organization

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := TODO // string | 
    office := *openapiclient.NewOffice("Id_example", "This is the name") // Office | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OfficeApi.CreateOffice(context.Background(), organizationId).Office(office).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfficeApi.CreateOffice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOffice`: OrganizationDetails
    fmt.Fprintf(os.Stdout, "Response from `OfficeApi.CreateOffice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOfficeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **office** | [**Office**](Office.md) |  | 

### Return type

[**OrganizationDetails**](OrganizationDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOffice

> Status DeleteOffice(ctx, organizationId, officeId).Execute()

Delte an Office from an Organization

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := TODO // string | 
    officeId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OfficeApi.DeleteOffice(context.Background(), organizationId, officeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfficeApi.DeleteOffice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOffice`: Status
    fmt.Fprintf(os.Stdout, "Response from `OfficeApi.DeleteOffice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) |  | 
**officeId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOfficeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Status**](Status.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOffice

> Office GetOffice(ctx, organizationId, officeId).Execute()

Get an Office for an Organiaztion

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := TODO // string | 
    officeId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OfficeApi.GetOffice(context.Background(), organizationId, officeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfficeApi.GetOffice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOffice`: Office
    fmt.Fprintf(os.Stdout, "Response from `OfficeApi.GetOffice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) |  | 
**officeId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOfficeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Office**](Office.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOffice

> Office UpdateOffice(ctx, organizationId, officeId).Office(office).Execute()

Update an Office for an Organization

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := TODO // string | 
    officeId := TODO // string | 
    office := *openapiclient.NewOffice("Id_example", "This is the name") // Office | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OfficeApi.UpdateOffice(context.Background(), organizationId, officeId).Office(office).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfficeApi.UpdateOffice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOffice`: Office
    fmt.Fprintf(os.Stdout, "Response from `OfficeApi.UpdateOffice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) |  | 
**officeId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOfficeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **office** | [**Office**](Office.md) |  | 

### Return type

[**Office**](Office.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

