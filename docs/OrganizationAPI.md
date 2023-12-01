# \OrganizationAPI

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrganizationToParentOrganization**](OrganizationAPI.md#AddOrganizationToParentOrganization) | **Post** /organizations/{organizationId}/parents/{parentOrganizationId} | Attach an Organization to a parent Organization, returns the parent Organization
[**AttachOrganizationToIndustry**](OrganizationAPI.md#AttachOrganizationToIndustry) | **Post** /organizations/{organizationId}/industries/{industryId} | Add an Organization to an Industry
[**CreateCertification**](OrganizationAPI.md#CreateCertification) | **Post** /organizations/{organizationId}/certifications | Create a Certification in an Organization
[**CreateOffice**](OrganizationAPI.md#CreateOffice) | **Post** /organizations/{organizationId}/offices | Create an Office in an Organization
[**CreateOrganization**](OrganizationAPI.md#CreateOrganization) | **Post** /organizations | Create an Organization
[**CreateProject**](OrganizationAPI.md#CreateProject) | **Post** /organizations/{organizationId}/projects | Create a Project in an Organization
[**DeleteOffice**](OrganizationAPI.md#DeleteOffice) | **Delete** /organizations/{organizationId}/offices/{officeId} | Delete an Office from an Organization
[**DeleteOrganization**](OrganizationAPI.md#DeleteOrganization) | **Delete** /organizations/{organizationId} | Delete an organization
[**DetachOrganizationFromIndustry**](OrganizationAPI.md#DetachOrganizationFromIndustry) | **Delete** /organizations/{organizationId}/industries/{industryId} | Remove an Organization to an Industry
[**GetCertificationsForOrganization**](OrganizationAPI.md#GetCertificationsForOrganization) | **Get** /organizations/{organizationId}/certifications | Get a list of all certifications for a organization
[**GetOffice**](OrganizationAPI.md#GetOffice) | **Get** /organizations/{organizationId}/offices/{officeId} | Get an Office for an Organization
[**GetOrganization**](OrganizationAPI.md#GetOrganization) | **Get** /organizations/{organizationId} | Get details about an Organization
[**GetOrganizationProjects**](OrganizationAPI.md#GetOrganizationProjects) | **Get** /organizations/{organizationId}/projects | Get a list of all Projects for an Organization
[**GetOrganizations**](OrganizationAPI.md#GetOrganizations) | **Get** /organizations | Get a list of all Organizations
[**MoveCertification**](OrganizationAPI.md#MoveCertification) | **Put** /organizations/{organizationId}/certificates/{certificateId} | Move a Certification to an Organization
[**RemoveOrganizationFromParentOrganization**](OrganizationAPI.md#RemoveOrganizationFromParentOrganization) | **Delete** /organizations/{organizationId}/parents/{parentOrganizationId} | Detaches an Organization from parent Organization, return the parent Organization
[**UpdateOffice**](OrganizationAPI.md#UpdateOffice) | **Put** /organizations/{organizationId}/offices/{officeId} | Update an Office for an Organization
[**UpdateOrganization**](OrganizationAPI.md#UpdateOrganization) | **Put** /organizations/{organizationId} | Update an Organization
[**UpdateProjectOrganization**](OrganizationAPI.md#UpdateProjectOrganization) | **Put** /organizations/{organizationId}/projects/{projectId} | project is now point to the new organization



## AddOrganizationToParentOrganization

> OrganizationDetails AddOrganizationToParentOrganization(ctx, organizationId, parentOrganizationId).Execute()

Attach an Organization to a parent Organization, returns the parent Organization

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
    parentOrganizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.AddOrganizationToParentOrganization(context.Background(), organizationId, parentOrganizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.AddOrganizationToParentOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddOrganizationToParentOrganization`: OrganizationDetails
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.AddOrganizationToParentOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**parentOrganizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOrganizationToParentOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationDetails**](OrganizationDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    resp, r, err := apiClient.OrganizationAPI.AttachOrganizationToIndustry(context.Background(), organizationId, industryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.AttachOrganizationToIndustry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachOrganizationToIndustry`: OrganizationDetails
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.AttachOrganizationToIndustry`: %v\n", resp)
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

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    certification := *openapiclient.NewCertification(false, "Id_example", "Name_example") // Certification | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.CreateCertification(context.Background(), organizationId).Certification(certification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.CreateCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCertification`: CertificationDetails
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.CreateCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certification** | [**Certification**](Certification.md) |  | 

### Return type

[**CertificationDetails**](CertificationDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    office := *openapiclient.NewOffice("Id_example", "Name_example") // Office | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.CreateOffice(context.Background(), organizationId).Office(office).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.CreateOffice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOffice`: OrganizationDetails
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.CreateOffice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

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


## CreateOrganization

> OrganizationDetails CreateOrganization(ctx).Organization(organization).Execute()

Create an Organization

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
    organization := *openapiclient.NewOrganization(false, "Id_example", "Name_example") // Organization | Details about the Organization

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.CreateOrganization(context.Background()).Organization(organization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.CreateOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganization`: OrganizationDetails
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.CreateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organization** | [**Organization**](Organization.md) | Details about the Organization | 

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


## CreateProject

> ProjectDetails CreateProject(ctx, organizationId).Project(project).Execute()

Create a Project in an Organization

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
    organizationId := "70ee6f30-d7c1-4f91-a653-9819ecbfa667" // string | 
    project := *openapiclient.NewProject(false, "Id_example", "Name_example") // Project | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.CreateProject(context.Background(), organizationId).Project(project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.CreateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProject`: ProjectDetails
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.CreateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | [**Project**](Project.md) |  | 

### Return type

[**ProjectDetails**](ProjectDetails.md)

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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    officeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.DeleteOffice(context.Background(), organizationId, officeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.DeleteOffice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOffice`: Status
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.DeleteOffice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**officeId** | **string** |  | 

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


## DeleteOrganization

> Status DeleteOrganization(ctx, organizationId).Execute()

Delete an organization

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.DeleteOrganization(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.DeleteOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrganization`: Status
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.DeleteOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationRequest struct via the builder pattern


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
    resp, r, err := apiClient.OrganizationAPI.DetachOrganizationFromIndustry(context.Background(), organizationId, industryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.DetachOrganizationFromIndustry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DetachOrganizationFromIndustry`: OrganizationDetails
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.DetachOrganizationFromIndustry`: %v\n", resp)
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

[bearerAuth](../README.md#bearerAuth)

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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skip := int32(0) // int32 |  (optional) (default to 0)
    limit := int32(20) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.GetCertificationsForOrganization(context.Background(), organizationId).Skip(skip).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetCertificationsForOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificationsForOrganization`: PagedCertifications
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetCertificationsForOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificationsForOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**PagedCertifications**](PagedCertifications.md)

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

Get an Office for an Organization

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
    officeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.GetOffice(context.Background(), organizationId, officeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetOffice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOffice`: Office
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetOffice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**officeId** | **string** |  | 

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


## GetOrganization

> OrganizationDetails GetOrganization(ctx, organizationId).Execute()

Get details about an Organization

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.GetOrganization(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganization`: OrganizationDetails
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationDetails**](OrganizationDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationProjects

> PagedProjects GetOrganizationProjects(ctx, organizationId).Skip(skip).Limit(limit).Execute()

Get a list of all Projects for an Organization

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
    organizationId := "70ee6f30-d7c1-4f91-a653-9819ecbfa667" // string | 
    skip := int32(0) // int32 |  (optional) (default to 0)
    limit := int32(20) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.GetOrganizationProjects(context.Background(), organizationId).Skip(skip).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetOrganizationProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationProjects`: PagedProjects
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetOrganizationProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**PagedProjects**](PagedProjects.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizations

> PagedOrganizations GetOrganizations(ctx).Skip(skip).Limit(limit).Term(term).Execute()

Get a list of all Organizations

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
    term := "term_example" // string | Optionally filter organizations by search term (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.GetOrganizations(context.Background()).Skip(skip).Limit(limit).Term(term).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.GetOrganizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizations`: PagedOrganizations
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.GetOrganizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]
 **term** | **string** | Optionally filter organizations by search term | [default to &quot;&quot;]

### Return type

[**PagedOrganizations**](PagedOrganizations.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveCertification

> CertificationDetails MoveCertification(ctx, organizationId, certificateId).Execute()

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
    organizationId := "70ee6f30-d7c1-4f91-a653-9819ecbfa667" // string | 
    certificateId := "70ee6f30-d7c1-4f91-a653-9819ecbfa667" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.MoveCertification(context.Background(), organizationId, certificateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.MoveCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoveCertification`: CertificationDetails
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.MoveCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**certificateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CertificationDetails**](CertificationDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOrganizationFromParentOrganization

> OrganizationDetails RemoveOrganizationFromParentOrganization(ctx, organizationId, parentOrganizationId).Execute()

Detaches an Organization from parent Organization, return the parent Organization

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
    parentOrganizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.RemoveOrganizationFromParentOrganization(context.Background(), organizationId, parentOrganizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.RemoveOrganizationFromParentOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveOrganizationFromParentOrganization`: OrganizationDetails
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.RemoveOrganizationFromParentOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**parentOrganizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOrganizationFromParentOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationDetails**](OrganizationDetails.md)

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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    officeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    office := *openapiclient.NewOffice("Id_example", "Name_example") // Office | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.UpdateOffice(context.Background(), organizationId, officeId).Office(office).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.UpdateOffice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOffice`: Office
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.UpdateOffice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**officeId** | **string** |  | 

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


## UpdateOrganization

> OrganizationDetails UpdateOrganization(ctx, organizationId).Organization(organization).Execute()

Update an Organization

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
    organization := *openapiclient.NewOrganization(false, "Id_example", "Name_example") // Organization | Details about the organization

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.UpdateOrganization(context.Background(), organizationId).Organization(organization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.UpdateOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganization`: OrganizationDetails
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.UpdateOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organization** | [**Organization**](Organization.md) | Details about the organization | 

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


## UpdateProjectOrganization

> ProjectDetails UpdateProjectOrganization(ctx, organizationId, projectId).Execute()

project is now point to the new organization

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
    organizationId := "70ee6f30-d7c1-4f91-a653-9819ecbfa667" // string | 
    projectId := "70ee6f30-d7c1-4f91-a653-9819ecbfa667" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.UpdateProjectOrganization(context.Background(), organizationId, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.UpdateProjectOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProjectOrganization`: ProjectDetails
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.UpdateProjectOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectDetails**](ProjectDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

