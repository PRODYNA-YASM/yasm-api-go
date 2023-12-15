# \LanguageAPI

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLanguageToCountry**](LanguageAPI.md#AddLanguageToCountry) | **Post** /countries/{countryId}/languages/{languageId} | Assign a language to a country
[**AddPersonLanguage**](LanguageAPI.md#AddPersonLanguage) | **Post** /persons/{personId}/languages/{languageId} | Assign a language to the person
[**CreateLanguage**](LanguageAPI.md#CreateLanguage) | **Post** /languages | Create a new language
[**DeleteLanguage**](LanguageAPI.md#DeleteLanguage) | **Delete** /languages/{languageId} | Delete a language
[**GetLanguage**](LanguageAPI.md#GetLanguage) | **Get** /languages/{languageId} | Get details about a language
[**GetLanguages**](LanguageAPI.md#GetLanguages) | **Get** /languages | Get a list of Languages
[**RemoveLanguageFromCountry**](LanguageAPI.md#RemoveLanguageFromCountry) | **Delete** /countries/{countryId}/languages/{languageId} | Assign a language to a country
[**RemovePersonLanguage**](LanguageAPI.md#RemovePersonLanguage) | **Delete** /persons/{personId}/languages/{languageId} | Remove a language from a person
[**UpdatePersonLanguage**](LanguageAPI.md#UpdatePersonLanguage) | **Put** /persons/{personId}/languages/{languageId} | Update a language of a person



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
    resp, r, err := apiClient.LanguageAPI.AddLanguageToCountry(context.Background(), countryId, languageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageAPI.AddLanguageToCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLanguageToCountry`: CountryDetails
    fmt.Fprintf(os.Stdout, "Response from `LanguageAPI.AddLanguageToCountry`: %v\n", resp)
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


## AddPersonLanguage

> PersonDetails AddPersonLanguage(ctx, personId, languageId).Level(level).Execute()

Assign a language to the person

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
    languageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    level := *openapiclient.NewLevel() // Level | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LanguageAPI.AddPersonLanguage(context.Background(), personId, languageId).Level(level).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageAPI.AddPersonLanguage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPersonLanguage`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `LanguageAPI.AddPersonLanguage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**languageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPersonLanguageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **level** | [**Level**](Level.md) |  | 

### Return type

[**PersonDetails**](PersonDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLanguage

> LanguageDetails CreateLanguage(ctx).Language(language).Execute()

Create a new language

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
    language := *openapiclient.NewLanguage("Id_example", "Name_example") // Language | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LanguageAPI.CreateLanguage(context.Background()).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageAPI.CreateLanguage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLanguage`: LanguageDetails
    fmt.Fprintf(os.Stdout, "Response from `LanguageAPI.CreateLanguage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLanguageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **language** | [**Language**](Language.md) |  | 

### Return type

[**LanguageDetails**](LanguageDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLanguage

> Status DeleteLanguage(ctx, languageId).Execute()

Delete a language

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
    languageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LanguageAPI.DeleteLanguage(context.Background(), languageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageAPI.DeleteLanguage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLanguage`: Status
    fmt.Fprintf(os.Stdout, "Response from `LanguageAPI.DeleteLanguage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**languageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLanguageRequest struct via the builder pattern


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


## GetLanguage

> LanguageDetails GetLanguage(ctx, languageId).Execute()

Get details about a language

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
    languageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LanguageAPI.GetLanguage(context.Background(), languageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageAPI.GetLanguage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLanguage`: LanguageDetails
    fmt.Fprintf(os.Stdout, "Response from `LanguageAPI.GetLanguage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**languageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanguageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LanguageDetails**](LanguageDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLanguages

> PagedLanguages GetLanguages(ctx).Skip(skip).Limit(limit).Execute()

Get a list of Languages

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
    skip := int32(0) // int32 |  (optional) (default to 0)
    limit := int32(20) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LanguageAPI.GetLanguages(context.Background()).Skip(skip).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageAPI.GetLanguages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLanguages`: PagedLanguages
    fmt.Fprintf(os.Stdout, "Response from `LanguageAPI.GetLanguages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLanguagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**PagedLanguages**](PagedLanguages.md)

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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    countryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    languageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LanguageAPI.RemoveLanguageFromCountry(context.Background(), countryId, languageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageAPI.RemoveLanguageFromCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveLanguageFromCountry`: CountryDetails
    fmt.Fprintf(os.Stdout, "Response from `LanguageAPI.RemoveLanguageFromCountry`: %v\n", resp)
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


## RemovePersonLanguage

> PersonDetails RemovePersonLanguage(ctx, personId, languageId).Execute()

Remove a language from a person

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
    languageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LanguageAPI.RemovePersonLanguage(context.Background(), personId, languageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageAPI.RemovePersonLanguage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemovePersonLanguage`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `LanguageAPI.RemovePersonLanguage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**languageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePersonLanguageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PersonDetails**](PersonDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePersonLanguage

> PersonDetails UpdatePersonLanguage(ctx, personId, languageId).Level(level).Execute()

Update a language of a person

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
    languageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    level := *openapiclient.NewLevel() // Level | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LanguageAPI.UpdatePersonLanguage(context.Background(), personId, languageId).Level(level).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageAPI.UpdatePersonLanguage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePersonLanguage`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `LanguageAPI.UpdatePersonLanguage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**languageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePersonLanguageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **level** | [**Level**](Level.md) |  | 

### Return type

[**PersonDetails**](PersonDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

