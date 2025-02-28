# \CountryAPI

All URIs are relative to *https://yasm.prodyna.com:443/api/graph/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLanguageToCountry**](CountryAPI.md#AddLanguageToCountry) | **Post** /countries/{countryId}/languages/{languageId} | Assign a language to a country
[**CreateCountry**](CountryAPI.md#CreateCountry) | **Post** /countries | Create a new Country
[**DeleteCountry**](CountryAPI.md#DeleteCountry) | **Delete** /countries/{countryId} | Delete a Country
[**GetCountry**](CountryAPI.md#GetCountry) | **Get** /countries/{countryId} | Get details about a Country
[**RemoveLanguageFromCountry**](CountryAPI.md#RemoveLanguageFromCountry) | **Delete** /countries/{countryId}/languages/{languageId} | Assign a language to a country
[**SearchCountries**](CountryAPI.md#SearchCountries) | **Post** /countries/search | Search over countries
[**UpdateCountry**](CountryAPI.md#UpdateCountry) | **Put** /countries/{countryId} | Update a Country
[**UpdateOfficeCountry**](CountryAPI.md#UpdateOfficeCountry) | **Put** /offices/{officeId}/countries/{countryId} | Assign a country to an office



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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    countryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    languageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CountryAPI.AddLanguageToCountry(context.Background(), countryId, languageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CountryAPI.AddLanguageToCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLanguageToCountry`: CountryDetails
    fmt.Fprintf(os.Stdout, "Response from `CountryAPI.AddLanguageToCountry`: %v\n", resp)
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

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    country := *openapiclient.NewCountry("Id_example", "Name_example") // Country | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CountryAPI.CreateCountry(context.Background()).Country(country).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CountryAPI.CreateCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCountry`: CountryDetails
    fmt.Fprintf(os.Stdout, "Response from `CountryAPI.CreateCountry`: %v\n", resp)
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

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    countryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CountryAPI.DeleteCountry(context.Background(), countryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CountryAPI.DeleteCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCountry`: Status
    fmt.Fprintf(os.Stdout, "Response from `CountryAPI.DeleteCountry`: %v\n", resp)
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

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    countryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CountryAPI.GetCountry(context.Background(), countryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CountryAPI.GetCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCountry`: CountryDetails
    fmt.Fprintf(os.Stdout, "Response from `CountryAPI.GetCountry`: %v\n", resp)
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

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    countryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    languageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CountryAPI.RemoveLanguageFromCountry(context.Background(), countryId, languageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CountryAPI.RemoveLanguageFromCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveLanguageFromCountry`: CountryDetails
    fmt.Fprintf(os.Stdout, "Response from `CountryAPI.RemoveLanguageFromCountry`: %v\n", resp)
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

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCountries

> PagedCountries SearchCountries(ctx).Search(search).Execute()

Search over countries

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
    resp, r, err := apiClient.CountryAPI.SearchCountries(context.Background()).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CountryAPI.SearchCountries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchCountries`: PagedCountries
    fmt.Fprintf(os.Stdout, "Response from `CountryAPI.SearchCountries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCountriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | [**Search**](Search.md) |  | 

### Return type

[**PagedCountries**](PagedCountries.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    countryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    country := *openapiclient.NewCountry("Id_example", "Name_example") // Country | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CountryAPI.UpdateCountry(context.Background(), countryId).Country(country).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CountryAPI.UpdateCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCountry`: CountryDetails
    fmt.Fprintf(os.Stdout, "Response from `CountryAPI.UpdateCountry`: %v\n", resp)
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

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOfficeCountry

> OfficeDetails UpdateOfficeCountry(ctx, officeId, countryId).Execute()

Assign a country to an office

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
    countryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CountryAPI.UpdateOfficeCountry(context.Background(), officeId, countryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CountryAPI.UpdateOfficeCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOfficeCountry`: OfficeDetails
    fmt.Fprintf(os.Stdout, "Response from `CountryAPI.UpdateOfficeCountry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** |  | 
**countryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOfficeCountryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OfficeDetails**](OfficeDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

