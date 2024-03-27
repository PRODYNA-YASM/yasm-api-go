# \OfficeAPI

All URIs are relative to *https://yasm.prodyna.com:443/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPersonOffice**](OfficeAPI.md#AddPersonOffice) | **Post** /persons/{personId}/offices/{officeId} | Assign a person to an office
[**CreateOffice**](OfficeAPI.md#CreateOffice) | **Post** /offices | Create an Office in an Organization
[**DeleteOffice**](OfficeAPI.md#DeleteOffice) | **Delete** /offices/{officeId} | Delete an Office from an Organization
[**DeletePersonOffice**](OfficeAPI.md#DeletePersonOffice) | **Delete** /persons/{personId}/offices/{officeId} | Delete the office from a Person
[**GetOfficeDetails**](OfficeAPI.md#GetOfficeDetails) | **Get** /offices/{officeId} | Get details about an Office independent of Organization
[**SearchOffices**](OfficeAPI.md#SearchOffices) | **Post** /offices/search | Search over offices
[**UpdateOffice**](OfficeAPI.md#UpdateOffice) | **Put** /offices/{officeId} | Update an Office for an Organization



## AddPersonOffice

> PersonDetails AddPersonOffice(ctx, personId, officeId).Execute()

Assign a person to an office

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
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    officeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OfficeAPI.AddPersonOffice(context.Background(), personId, officeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfficeAPI.AddPersonOffice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPersonOffice`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `OfficeAPI.AddPersonOffice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**officeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPersonOfficeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PersonDetails**](PersonDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOffice

> Office CreateOffice(ctx).OrganizationId(organizationId).Office(office).Execute()

Create an Office in an Organization

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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the organization
    office := *openapiclient.NewOffice("Id_example", "Name_example") // Office | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OfficeAPI.CreateOffice(context.Background()).OrganizationId(organizationId).Office(office).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfficeAPI.CreateOffice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOffice`: Office
    fmt.Fprintf(os.Stdout, "Response from `OfficeAPI.CreateOffice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOfficeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **string** | The ID of the organization | 
 **office** | [**Office**](Office.md) |  | 

### Return type

[**Office**](Office.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOffice

> Status DeleteOffice(ctx, officeId).Execute()

Delete an Office from an Organization

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
    officeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OfficeAPI.DeleteOffice(context.Background(), officeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfficeAPI.DeleteOffice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOffice`: Status
    fmt.Fprintf(os.Stdout, "Response from `OfficeAPI.DeleteOffice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOfficeRequest struct via the builder pattern


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


## DeletePersonOffice

> PersonDetails DeletePersonOffice(ctx, personId, officeId).Execute()

Delete the office from a Person

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
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    officeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OfficeAPI.DeletePersonOffice(context.Background(), personId, officeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfficeAPI.DeletePersonOffice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePersonOffice`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `OfficeAPI.DeletePersonOffice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**officeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePersonOfficeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PersonDetails**](PersonDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOfficeDetails

> Office GetOfficeDetails(ctx, officeId).Execute()

Get details about an Office independent of Organization

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
    officeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OfficeAPI.GetOfficeDetails(context.Background(), officeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfficeAPI.GetOfficeDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOfficeDetails`: Office
    fmt.Fprintf(os.Stdout, "Response from `OfficeAPI.GetOfficeDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOfficeDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Office**](Office.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    resp, r, err := apiClient.OfficeAPI.SearchOffices(context.Background()).OfficeSearch(officeSearch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfficeAPI.SearchOffices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchOffices`: PagedOffices
    fmt.Fprintf(os.Stdout, "Response from `OfficeAPI.SearchOffices`: %v\n", resp)
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


## UpdateOffice

> Office UpdateOffice(ctx, officeId).Office(office).Execute()

Update an Office for an Organization

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
    officeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    office := *openapiclient.NewOffice("Id_example", "Name_example") // Office | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OfficeAPI.UpdateOffice(context.Background(), officeId).Office(office).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfficeAPI.UpdateOffice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOffice`: Office
    fmt.Fprintf(os.Stdout, "Response from `OfficeAPI.UpdateOffice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOfficeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **office** | [**Office**](Office.md) |  | 

### Return type

[**Office**](Office.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

