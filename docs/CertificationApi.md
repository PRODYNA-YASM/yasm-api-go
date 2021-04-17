# \CertificationApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSkillToCertification**](CertificationApi.md#AddSkillToCertification) | **Post** /certifications/{certificationId}/skills/{skillId} | 
[**CertificationsCertificationIdSkillsSkillIdDelete**](CertificationApi.md#CertificationsCertificationIdSkillsSkillIdDelete) | **Delete** /certifications/{certificationId}/skills/{skillId} | 
[**CertificationsCertificationIdSkillsSkillIdPut**](CertificationApi.md#CertificationsCertificationIdSkillsSkillIdPut) | **Put** /certifications/{certificationId}/skills/{skillId} | 
[**DeleteCertification**](CertificationApi.md#DeleteCertification) | **Delete** /certifications/{certificationId} | Delete a Certification
[**GetCertification**](CertificationApi.md#GetCertification) | **Get** /certifications/{certificationId} | Get details about a Certification
[**GetCertifications**](CertificationApi.md#GetCertifications) | **Get** /certifications | Get a list of all Certifations indepdenant of the Organization
[**UpdateCertification**](CertificationApi.md#UpdateCertification) | **Put** /certifications/{certificationId} | Update a Certification



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


## CertificationsCertificationIdSkillsSkillIdDelete

> Status CertificationsCertificationIdSkillsSkillIdDelete(ctx, certificationId, skillId).Execute()



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
    resp, r, err := api_client.CertificationApi.CertificationsCertificationIdSkillsSkillIdDelete(context.Background(), certificationId, skillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationApi.CertificationsCertificationIdSkillsSkillIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificationsCertificationIdSkillsSkillIdDelete`: Status
    fmt.Fprintf(os.Stdout, "Response from `CertificationApi.CertificationsCertificationIdSkillsSkillIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificationId** | [**string**](.md) |  | 
**skillId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificationsCertificationIdSkillsSkillIdDeleteRequest struct via the builder pattern


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


## CertificationsCertificationIdSkillsSkillIdPut

> CertificationDetails CertificationsCertificationIdSkillsSkillIdPut(ctx, certificationId, skillId).Level(level).Execute()





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
    resp, r, err := api_client.CertificationApi.CertificationsCertificationIdSkillsSkillIdPut(context.Background(), certificationId, skillId).Level(level).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationApi.CertificationsCertificationIdSkillsSkillIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CertificationsCertificationIdSkillsSkillIdPut`: CertificationDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificationApi.CertificationsCertificationIdSkillsSkillIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificationId** | [**string**](.md) |  | 
**skillId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCertificationsCertificationIdSkillsSkillIdPutRequest struct via the builder pattern


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
    certification := *openapiclient.NewCertification("Id_example", "Name_example") // Certification | 

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

