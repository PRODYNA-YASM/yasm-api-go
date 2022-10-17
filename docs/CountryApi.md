# \CountryApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLanguageToCountry**](CountryApi.md#AddLanguageToCountry) | **Post** /countries/{countryId}/languages/{languageId} | Assign a language to a country
[**CreateCountry**](CountryApi.md#CreateCountry) | **Post** /countries | Create a new Country
[**DeleteCountry**](CountryApi.md#DeleteCountry) | **Delete** /countries/{countryId} | Delete a Country
[**GetCountries**](CountryApi.md#GetCountries) | **Get** /countries | Get all Countries
[**GetCountry**](CountryApi.md#GetCountry) | **Get** /countries/{countryId} | Get details about a Country
[**RemoveLanguageFromCountry**](CountryApi.md#RemoveLanguageFromCountry) | **Delete** /countries/{countryId}/languages/{languageId} | Assign a language to a country
[**UpdateCountry**](CountryApi.md#UpdateCountry) | **Put** /countries/{countryId} | Update a Country



## AddLanguageToCountry

> CountryDetails AddLanguageToCountry(ctx, countryId, languageId).Execute()

Assign a language to a country

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
    countryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    languageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CountryApi.AddLanguageToCountry(context.Background(), countryId, languageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CountryApi.AddLanguageToCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLanguageToCountry`: CountryDetails
    fmt.Fprintf(os.Stdout, "Response from `CountryApi.AddLanguageToCountry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryId** | **string** |  | 
**languageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLanguageToCountryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CountryDetails**](CountryDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCountry

> CountryDetails CreateCountry(ctx).Country(country).Execute()

Create a new Country

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
    country := *openapiclient.NewCountry("Id_example") // Country | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CountryApi.CreateCountry(context.Background()).Country(country).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CountryApi.CreateCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCountry`: CountryDetails
    fmt.Fprintf(os.Stdout, "Response from `CountryApi.CreateCountry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCountryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | [**Country**](Country.md) |  | 

### Return type

[**CountryDetails**](CountryDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCountry

> Status DeleteCountry(ctx, countryId).Execute()

Delete a Country

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
    countryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CountryApi.DeleteCountry(context.Background(), countryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CountryApi.DeleteCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCountry`: Status
    fmt.Fprintf(os.Stdout, "Response from `CountryApi.DeleteCountry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCountryRequest struct via the builder pattern


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


## GetCountries

> PagedCountries GetCountries(ctx).Skip(skip).Limit(limit).Execute()

Get all Countries

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
    skip := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CountryApi.GetCountries(context.Background()).Skip(skip).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CountryApi.GetCountries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCountries`: PagedCountries
    fmt.Fprintf(os.Stdout, "Response from `CountryApi.GetCountries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCountriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**PagedCountries**](PagedCountries.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountry

> CountryDetails GetCountry(ctx, countryId).Execute()

Get details about a Country

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
    countryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CountryApi.GetCountry(context.Background(), countryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CountryApi.GetCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCountry`: CountryDetails
    fmt.Fprintf(os.Stdout, "Response from `CountryApi.GetCountry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCountryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CountryDetails**](CountryDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveLanguageFromCountry

> CountryDetails RemoveLanguageFromCountry(ctx, countryId, languageId).Execute()

Assign a language to a country

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
    countryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    languageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CountryApi.RemoveLanguageFromCountry(context.Background(), countryId, languageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CountryApi.RemoveLanguageFromCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveLanguageFromCountry`: CountryDetails
    fmt.Fprintf(os.Stdout, "Response from `CountryApi.RemoveLanguageFromCountry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryId** | **string** |  | 
**languageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveLanguageFromCountryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CountryDetails**](CountryDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCountry

> CountryDetails UpdateCountry(ctx, countryId).Country(country).Execute()

Update a Country

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
    countryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    country := *openapiclient.NewCountry("Id_example") // Country | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CountryApi.UpdateCountry(context.Background(), countryId).Country(country).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CountryApi.UpdateCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCountry`: CountryDetails
    fmt.Fprintf(os.Stdout, "Response from `CountryApi.UpdateCountry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCountryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **country** | [**Country**](Country.md) |  | 

### Return type

[**CountryDetails**](CountryDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

