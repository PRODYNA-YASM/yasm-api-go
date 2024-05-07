# \CertificationAPI

All URIs are relative to *https://yasm.prodyna.com:443/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPersonCertification**](CertificationAPI.md#AddPersonCertification) | **Post** /persons/{personId}/certifications/{certificationId} | Add Certification to a Person
[**AddSkillToCertification**](CertificationAPI.md#AddSkillToCertification) | **Post** /certifications/{certificationId}/skills/{skillId} | 
[**CreateCertification**](CertificationAPI.md#CreateCertification) | **Post** /certifications | Create a Certification in an Organization
[**DeleteCertification**](CertificationAPI.md#DeleteCertification) | **Delete** /certifications/{certificationId} | Delete a Certification
[**DeletePersonCertification**](CertificationAPI.md#DeletePersonCertification) | **Delete** /persons/{personId}/certifications/{certificationId} | Remove an Interest to a Person
[**DeleteSkillFromCertification**](CertificationAPI.md#DeleteSkillFromCertification) | **Delete** /certifications/{certificationId}/skills/{skillId} | 
[**GetCertification**](CertificationAPI.md#GetCertification) | **Get** /certifications/{certificationId} | Get details about a Certification
[**MoveCertification**](CertificationAPI.md#MoveCertification) | **Put** /certifications/{certificationId}/organizations/{organizationId} | Move a Certification to an Organization
[**SearchCertifications**](CertificationAPI.md#SearchCertifications) | **Post** /certifications/search | Search over certifications
[**UpdateCertification**](CertificationAPI.md#UpdateCertification) | **Put** /certifications/{certificationId} | Update a Certification
[**UpdatePersonCertification**](CertificationAPI.md#UpdatePersonCertification) | **Put** /persons/{personId}/certifications/{certificationId} | Update a Certification of a Person
[**UpdateSkillInCertification**](CertificationAPI.md#UpdateSkillInCertification) | **Put** /certifications/{certificationId}/skills/{skillId} | 



## AddPersonCertification

> PersonDetails AddPersonCertification(ctx, personId, certificationId).Body(body).Execute()

Add Certification to a Person

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    certificationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    body := time.Now() // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationAPI.AddPersonCertification(context.Background(), personId, certificationId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationAPI.AddPersonCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPersonCertification`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificationAPI.AddPersonCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**certificationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPersonCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** |  | 

### Return type

[**PersonDetails**](PersonDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddSkillToCertification

> CertificationDetails AddSkillToCertification(ctx, certificationId, skillId).Level(level).Execute()





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
    certificationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    level := *openapiclient.NewLevel() // Level | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationAPI.AddSkillToCertification(context.Background(), certificationId, skillId).Level(level).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationAPI.AddSkillToCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSkillToCertification`: CertificationDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificationAPI.AddSkillToCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificationId** | **string** |  | 
**skillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSkillToCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **level** | [**Level**](Level.md) |  | 

### Return type

[**CertificationDetails**](CertificationDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCertification

> CertificationDetails CreateCertification(ctx).OrganizationId(organizationId).Certification(certification).Execute()

Create a Certification in an Organization

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
    certification := *openapiclient.NewCertification("Id_example", "Name_example") // Certification | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationAPI.CreateCertification(context.Background()).OrganizationId(organizationId).Certification(certification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationAPI.CreateCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCertification`: CertificationDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificationAPI.CreateCertification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **string** | The ID of the organization | 
 **certification** | [**Certification**](Certification.md) |  | 

### Return type

[**CertificationDetails**](CertificationDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCertification

> Status DeleteCertification(ctx, certificationId).Execute()

Delete a Certification

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
    certificationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationAPI.DeleteCertification(context.Background(), certificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationAPI.DeleteCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCertification`: Status
    fmt.Fprintf(os.Stdout, "Response from `CertificationAPI.DeleteCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertificationRequest struct via the builder pattern


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


## DeletePersonCertification

> PersonDetails DeletePersonCertification(ctx, personId, certificationId).Execute()

Remove an Interest to a Person

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
    certificationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationAPI.DeletePersonCertification(context.Background(), personId, certificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationAPI.DeletePersonCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePersonCertification`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificationAPI.DeletePersonCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**certificationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePersonCertificationRequest struct via the builder pattern


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


## DeleteSkillFromCertification

> CertificationDetails DeleteSkillFromCertification(ctx, certificationId, skillId).Execute()



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
    certificationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationAPI.DeleteSkillFromCertification(context.Background(), certificationId, skillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationAPI.DeleteSkillFromCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSkillFromCertification`: CertificationDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificationAPI.DeleteSkillFromCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificationId** | **string** |  | 
**skillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSkillFromCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CertificationDetails**](CertificationDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertification

> CertificationDetails GetCertification(ctx, certificationId).Execute()

Get details about a Certification

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
    certificationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationAPI.GetCertification(context.Background(), certificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationAPI.GetCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertification`: CertificationDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificationAPI.GetCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CertificationDetails**](CertificationDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveCertification

> CertificationDetails MoveCertification(ctx, certificationId, organizationId).Execute()

Move a Certification to an Organization

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
    certificationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationAPI.MoveCertification(context.Background(), certificationId, organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationAPI.MoveCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoveCertification`: CertificationDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificationAPI.MoveCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificationId** | **string** |  | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CertificationDetails**](CertificationDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCertifications

> PagedCertifications SearchCertifications(ctx).CertificationSearch(certificationSearch).Execute()

Search over certifications

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
    certificationSearch := *openapiclient.NewCertificationSearch(int32(123), int32(123)) // CertificationSearch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationAPI.SearchCertifications(context.Background()).CertificationSearch(certificationSearch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationAPI.SearchCertifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchCertifications`: PagedCertifications
    fmt.Fprintf(os.Stdout, "Response from `CertificationAPI.SearchCertifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCertificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **certificationSearch** | [**CertificationSearch**](CertificationSearch.md) |  | 

### Return type

[**PagedCertifications**](PagedCertifications.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCertification

> CertificationDetails UpdateCertification(ctx, certificationId).Certification(certification).Execute()

Update a Certification

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
    certificationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    certification := *openapiclient.NewCertification("Id_example", "Name_example") // Certification | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationAPI.UpdateCertification(context.Background(), certificationId).Certification(certification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationAPI.UpdateCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCertification`: CertificationDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificationAPI.UpdateCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certification** | [**Certification**](Certification.md) |  | 

### Return type

[**CertificationDetails**](CertificationDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePersonCertification

> PersonDetails UpdatePersonCertification(ctx, personId, certificationId).Body(body).Execute()

Update a Certification of a Person

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    certificationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    body := time.Now() // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationAPI.UpdatePersonCertification(context.Background(), personId, certificationId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationAPI.UpdatePersonCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePersonCertification`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificationAPI.UpdatePersonCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**certificationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePersonCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** |  | 

### Return type

[**PersonDetails**](PersonDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSkillInCertification

> CertificationDetails UpdateSkillInCertification(ctx, certificationId, skillId).Level(level).Execute()





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
    certificationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    level := *openapiclient.NewLevel() // Level | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationAPI.UpdateSkillInCertification(context.Background(), certificationId, skillId).Level(level).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationAPI.UpdateSkillInCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSkillInCertification`: CertificationDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificationAPI.UpdateSkillInCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificationId** | **string** |  | 
**skillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSkillInCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **level** | [**Level**](Level.md) |  | 

### Return type

[**CertificationDetails**](CertificationDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

