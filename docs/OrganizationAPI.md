# \OrganizationAPI

All URIs are relative to *https://yasm.prodyna.com:443/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddExecutiveOrganizationToProject**](OrganizationAPI.md#AddExecutiveOrganizationToProject) | **Post** /projects/{projectId}/executive-organizations/{organizationId} | Add an Organization to a Project as executive organization
[**AddOrganizationToParentOrganization**](OrganizationAPI.md#AddOrganizationToParentOrganization) | **Post** /organizations/{organizationId}/parents/{parentOrganizationId} | Attach an Organization to a parent Organization, returns the parent Organization
[**AttachOrganizationToIndustry**](OrganizationAPI.md#AttachOrganizationToIndustry) | **Post** /organizations/{organizationId}/industries/{industryId} | Add an Organization to an Industry
[**CreateOffice**](OrganizationAPI.md#CreateOffice) | **Post** /offices | Create an Office in an Organization
[**CreateOrganization**](OrganizationAPI.md#CreateOrganization) | **Post** /organizations | Create an Organization
[**CreateProject**](OrganizationAPI.md#CreateProject) | **Post** /projects | Create a Project in an Organization
[**DeleteOrganization**](OrganizationAPI.md#DeleteOrganization) | **Delete** /organizations/{organizationId} | Delete an organization
[**DetachOrganizationFromIndustry**](OrganizationAPI.md#DetachOrganizationFromIndustry) | **Delete** /organizations/{organizationId}/industries/{industryId} | Remove an Organization to an Industry
[**GetOrganization**](OrganizationAPI.md#GetOrganization) | **Get** /organizations/{organizationId} | Get details about an Organization
[**MoveCertification**](OrganizationAPI.md#MoveCertification) | **Put** /certifications/{certificationId}/organizations/{organizationId} | Move a Certification to an Organization
[**MoveProject**](OrganizationAPI.md#MoveProject) | **Put** /projects/{projectId}/organizations/{organizationId} | Move a Project to an Organization
[**RemoveExecutiveOrganizationFromProject**](OrganizationAPI.md#RemoveExecutiveOrganizationFromProject) | **Delete** /projects/{projectId}/executive-organizations/{organizationId} | Remove an Organization from a Project as executive organization
[**RemoveOrganizationFromParentOrganization**](OrganizationAPI.md#RemoveOrganizationFromParentOrganization) | **Delete** /organizations/{organizationId}/parents/{parentOrganizationId} | Detaches an Organization from parent Organization, return the parent Organization
[**SearchOrganizations**](OrganizationAPI.md#SearchOrganizations) | **Post** /organizations/search | Search over organizations
[**UpdateOrganization**](OrganizationAPI.md#UpdateOrganization) | **Put** /organizations/{organizationId} | Update an Organization



## AddExecutiveOrganizationToProject

> ProjectDetails AddExecutiveOrganizationToProject(ctx, projectId, organizationId).Execute()

Add an Organization to a Project as executive organization

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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.AddExecutiveOrganizationToProject(context.Background(), projectId, organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.AddExecutiveOrganizationToProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddExecutiveOrganizationToProject`: ProjectDetails
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.AddExecutiveOrganizationToProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddExecutiveOrganizationToProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectDetails**](ProjectDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

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
    resp, r, err := apiClient.OrganizationAPI.CreateOffice(context.Background()).OrganizationId(organizationId).Office(office).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.CreateOffice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOffice`: Office
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.CreateOffice`: %v\n", resp)
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

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProject

> ProjectDetails CreateProject(ctx).OrganizationId(organizationId).Project(project).Execute()

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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the organization
    project := *openapiclient.NewProject(false, "Id_example", "Name_example") // Project | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.CreateProject(context.Background()).OrganizationId(organizationId).Project(project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.CreateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProject`: ProjectDetails
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.CreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **string** | The ID of the organization | 
 **project** | [**Project**](Project.md) |  | 

### Return type

[**ProjectDetails**](ProjectDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
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

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

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
    resp, r, err := apiClient.OrganizationAPI.MoveCertification(context.Background(), certificationId, organizationId).Execute()
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


## MoveProject

> ProjectDetails MoveProject(ctx, projectId, organizationId).Execute()

Move a Project to an Organization

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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.MoveProject(context.Background(), projectId, organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.MoveProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoveProject`: ProjectDetails
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.MoveProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectDetails**](ProjectDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveExecutiveOrganizationFromProject

> ProjectDetails RemoveExecutiveOrganizationFromProject(ctx, projectId, organizationId).Execute()

Remove an Organization from a Project as executive organization

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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationAPI.RemoveExecutiveOrganizationFromProject(context.Background(), projectId, organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.RemoveExecutiveOrganizationFromProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveExecutiveOrganizationFromProject`: ProjectDetails
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.RemoveExecutiveOrganizationFromProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveExecutiveOrganizationFromProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectDetails**](ProjectDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

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

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchOrganizations

> PagedOrganizations SearchOrganizations(ctx).Search(search).Execute()

Search over organizations

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
    resp, r, err := apiClient.OrganizationAPI.SearchOrganizations(context.Background()).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.SearchOrganizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchOrganizations`: PagedOrganizations
    fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.SearchOrganizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | [**Search**](Search.md) |  | 

### Return type

[**PagedOrganizations**](PagedOrganizations.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

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

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

