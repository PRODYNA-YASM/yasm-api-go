# \CertificationApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPersonCertification**](CertificationApi.md#AddPersonCertification) | **Post** /persons/{personId}/certifications/{certificationId} | Add Certification to a Person
[**AddSkillToCertification**](CertificationApi.md#AddSkillToCertification) | **Post** /certifications/{certificationId}/skills/{skillId} | 
[**CreateCertification**](CertificationApi.md#CreateCertification) | **Post** /organizations/{organizationId}/certifications | Create a Certification in an Organization
[**DeleteCertification**](CertificationApi.md#DeleteCertification) | **Delete** /certifications/{certificationId} | Delete a Certification
[**DeletePersonCertification**](CertificationApi.md#DeletePersonCertification) | **Delete** /persons/{personId}/certifications/{certificationId} | Remove an Interest to a Person
[**DeleteSkillFromCertification**](CertificationApi.md#DeleteSkillFromCertification) | **Delete** /certifications/{certificationId}/skills/{skillId} | 
[**GetCertification**](CertificationApi.md#GetCertification) | **Get** /certifications/{certificationId} | Get details about a Certification
[**GetCertifications**](CertificationApi.md#GetCertifications) | **Get** /certifications | Get a list of all Certifations indepdenant of the Organization
[**GetCertificationsForOrganization**](CertificationApi.md#GetCertificationsForOrganization) | **Get** /organizations/{organizationId}/certifications | Get a list of all certifications for a organization
[**UpdateCertification**](CertificationApi.md#UpdateCertification) | **Put** /certifications/{certificationId} | Update a Certification
[**UpdatePersonCertification**](CertificationApi.md#UpdatePersonCertification) | **Put** /persons/{personId}/certifications/{certificationId} | Update a Certification of a Person
[**UpdateSkillInCertification**](CertificationApi.md#UpdateSkillInCertification) | **Put** /certifications/{certificationId}/skills/{skillId} | 



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
    openapiclient "./openapi"
)

func main() {
    personId := TODO // string | 
    certificationId := TODO // string | 
    body := time.Now() // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificationApi.AddPersonCertification(context.Background(), personId, certificationId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationApi.AddPersonCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPersonCertification`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificationApi.AddPersonCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | [**string**](.md) |  | 
**certificationId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPersonCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** |  | 

### Return type

[**PersonDetails**](PersonDetails.md)

### Authorization

No authorization required

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
    openapiclient "./openapi"
)

func main() {
    certificationId := TODO // string | 
    skillId := TODO // string | 
    level := *openapiclient.NewLevel() // Level | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificationApi.AddSkillToCertification(context.Background(), certificationId, skillId).Level(level).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationApi.AddSkillToCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSkillToCertification`: CertificationDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificationApi.AddSkillToCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificationId** | [**string**](.md) |  | 
**skillId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSkillToCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **level** | [**Level**](Level.md) |  | 

### Return type

[**CertificationDetails**](CertificationDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCertification

> CertificationDetails CreateCertification(ctx, organizationId).Certification(certification).Execute()

Create a Certification in an Organization

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
    certification := *openapiclient.NewCertification("Id_example", "This is the name") // Certification | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificationApi.CreateCertification(context.Background(), organizationId).Certification(certification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationApi.CreateCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCertification`: CertificationDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificationApi.CreateCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certification** | [**Certification**](Certification.md) |  | 

### Return type

[**CertificationDetails**](CertificationDetails.md)

### Authorization

No authorization required

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
    openapiclient "./openapi"
)

func main() {
    certificationId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificationApi.DeleteCertification(context.Background(), certificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationApi.DeleteCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCertification`: Status
    fmt.Fprintf(os.Stdout, "Response from `CertificationApi.DeleteCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificationId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Status**](Status.md)

### Authorization

No authorization required

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
    openapiclient "./openapi"
)

func main() {
    personId := TODO // string | 
    certificationId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificationApi.DeletePersonCertification(context.Background(), personId, certificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationApi.DeletePersonCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePersonCertification`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificationApi.DeletePersonCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | [**string**](.md) |  | 
**certificationId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePersonCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PersonDetails**](PersonDetails.md)

### Authorization

No authorization required

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
    openapiclient "./openapi"
)

func main() {
    certificationId := TODO // string | 
    skillId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificationApi.DeleteSkillFromCertification(context.Background(), certificationId, skillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationApi.DeleteSkillFromCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSkillFromCertification`: CertificationDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificationApi.DeleteSkillFromCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificationId** | [**string**](.md) |  | 
**skillId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSkillFromCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CertificationDetails**](CertificationDetails.md)

### Authorization

No authorization required

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
    openapiclient "./openapi"
)

func main() {
    certificationId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificationApi.GetCertification(context.Background(), certificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationApi.GetCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertification`: CertificationDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificationApi.GetCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificationId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CertificationDetails**](CertificationDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertifications

> PagedCertifications GetCertifications(ctx).Skip(skip).Limit(limit).Execute()

Get a list of all Certifations indepdenant of the Organization

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificationApi.GetCertifications(context.Background()).Skip(skip).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationApi.GetCertifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertifications`: PagedCertifications
    fmt.Fprintf(os.Stdout, "Response from `CertificationApi.GetCertifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**PagedCertifications**](PagedCertifications.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificationsForOrganization

> PagedCertifications GetCertificationsForOrganization(ctx, organizationId).Skip(skip).Limit(limit).Execute()

Get a list of all certifications for a organization

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
    skip := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificationApi.GetCertificationsForOrganization(context.Background(), organizationId).Skip(skip).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationApi.GetCertificationsForOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificationsForOrganization`: PagedCertifications
    fmt.Fprintf(os.Stdout, "Response from `CertificationApi.GetCertificationsForOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificationsForOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**PagedCertifications**](PagedCertifications.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
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
    openapiclient "./openapi"
)

func main() {
    certificationId := TODO // string | 
    certification := *openapiclient.NewCertification("Id_example", "This is the name") // Certification | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificationApi.UpdateCertification(context.Background(), certificationId).Certification(certification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationApi.UpdateCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCertification`: CertificationDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificationApi.UpdateCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificationId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certification** | [**Certification**](Certification.md) |  | 

### Return type

[**CertificationDetails**](CertificationDetails.md)

### Authorization

No authorization required

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
    openapiclient "./openapi"
)

func main() {
    personId := TODO // string | 
    certificationId := TODO // string | 
    body := time.Now() // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificationApi.UpdatePersonCertification(context.Background(), personId, certificationId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationApi.UpdatePersonCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePersonCertification`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificationApi.UpdatePersonCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | [**string**](.md) |  | 
**certificationId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePersonCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** |  | 

### Return type

[**PersonDetails**](PersonDetails.md)

### Authorization

No authorization required

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
    openapiclient "./openapi"
)

func main() {
    certificationId := TODO // string | 
    skillId := TODO // string | 
    level := *openapiclient.NewLevel() // Level | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CertificationApi.UpdateSkillInCertification(context.Background(), certificationId, skillId).Level(level).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationApi.UpdateSkillInCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSkillInCertification`: CertificationDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificationApi.UpdateSkillInCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificationId** | [**string**](.md) |  | 
**skillId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSkillInCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **level** | [**Level**](Level.md) |  | 

### Return type

[**CertificationDetails**](CertificationDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

