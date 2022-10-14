# \ProjectApi

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPersonProject**](ProjectApi.md#AddPersonProject) | **Post** /persons/{personId}/projects/{projectId} | Add Project to a Person
[**AddPersonProjectSkill**](ProjectApi.md#AddPersonProjectSkill) | **Post** /persons/{personId}/projects/{projectId}/skills/{skillId} | Add Skill to a Project participation
[**ConfirmSkill**](ProjectApi.md#ConfirmSkill) | **Post** /persons/{personId}/projects/{projectId}/skills/{skillId}/confirmation/{confirmingPersonId} | Confirm Skill
[**CreateProject**](ProjectApi.md#CreateProject) | **Post** /organizations/{organizationId}/projects | Create a Project in an Organization
[**DeleteConfirmation**](ProjectApi.md#DeleteConfirmation) | **Delete** /persons/{personId}/projects/{projectId}/skills/{skillId}/confirmation/{confirmingPersonId} | Remove a confirmation
[**DeletePersonProject**](ProjectApi.md#DeletePersonProject) | **Delete** /persons/{personId}/projects/{projectId} | Remove an Project from a Person
[**DeletePersonProjectSkill**](ProjectApi.md#DeletePersonProjectSkill) | **Delete** /persons/{personId}/projects/{projectId}/skills/{skillId} | Remove a Skill from a Project participation
[**DeleteProject**](ProjectApi.md#DeleteProject) | **Delete** /projects/{projectId} | Delete a project
[**GetOrganizationProjects**](ProjectApi.md#GetOrganizationProjects) | **Get** /organizations/{organizationId}/projects | Get a list of all Projects for an Organization
[**GetProject**](ProjectApi.md#GetProject) | **Get** /projects/{projectId} | Get details about a Project
[**MergeProjects**](ProjectApi.md#MergeProjects) | **Put** /projects/{projectId}/merge/{otherProjectId} | Merge to projects
[**ReadPersonProject**](ProjectApi.md#ReadPersonProject) | **Get** /persons/{personId}/projects/{projectId} | Get a Project Partifipation of a Person
[**SearchProjects**](ProjectApi.md#SearchProjects) | **Post** /projects/search | Complex search over project entities
[**UpdatePersonProject**](ProjectApi.md#UpdatePersonProject) | **Put** /persons/{personId}/projects/{projectId} | Update a Project of a Person
[**UpdatePersonProjectSkill**](ProjectApi.md#UpdatePersonProjectSkill) | **Put** /persons/{personId}/projects/{projectId}/skills/{skillId} | Update the level of a Skill in a Project participation
[**UpdateProject**](ProjectApi.md#UpdateProject) | **Put** /projects/{projectId} | Update a Project



## AddPersonProject

> PersonDetails AddPersonProject(ctx, personId, projectId).ProjectParticipationUpdate(projectParticipationUpdate).Execute()

Add Project to a Person

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
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    projectParticipationUpdate := *openapiclient.NewProjectParticipationUpdate() // ProjectParticipationUpdate | List of Skills with level and timeframe

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.AddPersonProject(context.Background(), personId, projectId).ProjectParticipationUpdate(projectParticipationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.AddPersonProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPersonProject`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.AddPersonProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPersonProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectParticipationUpdate** | [**ProjectParticipationUpdate**](ProjectParticipationUpdate.md) | List of Skills with level and timeframe | 

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


## AddPersonProjectSkill

> PersonDetails AddPersonProjectSkill(ctx, personId, projectId, skillId).Level(level).Execute()

Add Skill to a Project participation

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
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    level := *openapiclient.NewLevel() // Level | The Skill Level

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.AddPersonProjectSkill(context.Background(), personId, projectId, skillId).Level(level).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.AddPersonProjectSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPersonProjectSkill`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.AddPersonProjectSkill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**projectId** | **string** |  | 
**skillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPersonProjectSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **level** | [**Level**](Level.md) | The Skill Level | 

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


## ConfirmSkill

> PersonDetails ConfirmSkill(ctx, personId, projectId, skillId, confirmingPersonId).Execute()

Confirm Skill

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
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    confirmingPersonId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ConfirmSkill(context.Background(), personId, projectId, skillId, confirmingPersonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ConfirmSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmSkill`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ConfirmSkill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**projectId** | **string** |  | 
**skillId** | **string** |  | 
**confirmingPersonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmSkillRequest struct via the builder pattern


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
    openapiclient "./openapi"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    project := *openapiclient.NewProject(false, "Id_example") // Project | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.CreateProject(context.Background(), organizationId).Project(project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.CreateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProject`: ProjectDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.CreateProject`: %v\n", resp)
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


## DeleteConfirmation

> PersonDetails DeleteConfirmation(ctx, personId, projectId, skillId, confirmingPersonId).Execute()

Remove a confirmation

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
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    confirmingPersonId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.DeleteConfirmation(context.Background(), personId, projectId, skillId, confirmingPersonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.DeleteConfirmation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteConfirmation`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.DeleteConfirmation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**projectId** | **string** |  | 
**skillId** | **string** |  | 
**confirmingPersonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfirmationRequest struct via the builder pattern


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


## DeletePersonProject

> PersonDetails DeletePersonProject(ctx, personId, projectId).Execute()

Remove an Project from a Person

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
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.DeletePersonProject(context.Background(), personId, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.DeletePersonProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePersonProject`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.DeletePersonProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePersonProjectRequest struct via the builder pattern


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


## DeletePersonProjectSkill

> PersonDetails DeletePersonProjectSkill(ctx, personId, projectId, skillId).Execute()

Remove a Skill from a Project participation

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
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.DeletePersonProjectSkill(context.Background(), personId, projectId, skillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.DeletePersonProjectSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePersonProjectSkill`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.DeletePersonProjectSkill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**projectId** | **string** |  | 
**skillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePersonProjectSkillRequest struct via the builder pattern


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


## DeleteProject

> Status DeleteProject(ctx, projectId).Execute()

Delete a project

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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.DeleteProject(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.DeleteProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteProject`: Status
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.DeleteProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectRequest struct via the builder pattern


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
    openapiclient "./openapi"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skip := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.GetOrganizationProjects(context.Background(), organizationId).Skip(skip).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.GetOrganizationProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationProjects`: PagedProjects
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.GetOrganizationProjects`: %v\n", resp)
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


## GetProject

> ProjectDetails GetProject(ctx, projectId).Execute()

Get details about a Project

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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.GetProject(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.GetProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProject`: ProjectDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.GetProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectRequest struct via the builder pattern


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


## MergeProjects

> ProjectDetails MergeProjects(ctx, projectId, otherProjectId).Execute()

Merge to projects

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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    otherProjectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.MergeProjects(context.Background(), projectId, otherProjectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.MergeProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MergeProjects`: ProjectDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.MergeProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**otherProjectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMergeProjectsRequest struct via the builder pattern


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


## ReadPersonProject

> ProjectParticipation ReadPersonProject(ctx, personId, projectId).Execute()

Get a Project Partifipation of a Person

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
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.ReadPersonProject(context.Background(), personId, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.ReadPersonProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadPersonProject`: ProjectParticipation
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.ReadPersonProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadPersonProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectParticipation**](ProjectParticipation.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchProjects

> PagedProjects SearchProjects(ctx).ProjectSearch(projectSearch).Skip(skip).Limit(limit).Execute()

Complex search over project entities

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
    projectSearch := *openapiclient.NewProjectSearch() // ProjectSearch | 
    skip := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.SearchProjects(context.Background()).ProjectSearch(projectSearch).Skip(skip).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.SearchProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchProjects`: PagedProjects
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.SearchProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectSearch** | [**ProjectSearch**](ProjectSearch.md) |  | 
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**PagedProjects**](PagedProjects.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePersonProject

> PersonDetails UpdatePersonProject(ctx, personId, projectId).ProjectParticipationUpdate(projectParticipationUpdate).Execute()

Update a Project of a Person

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
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    projectParticipationUpdate := *openapiclient.NewProjectParticipationUpdate() // ProjectParticipationUpdate | A time frame and a list of Skills

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.UpdatePersonProject(context.Background(), personId, projectId).ProjectParticipationUpdate(projectParticipationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.UpdatePersonProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePersonProject`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.UpdatePersonProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePersonProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectParticipationUpdate** | [**ProjectParticipationUpdate**](ProjectParticipationUpdate.md) | A time frame and a list of Skills | 

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


## UpdatePersonProjectSkill

> PersonDetails UpdatePersonProjectSkill(ctx, personId, projectId, skillId).Level(level).Execute()

Update the level of a Skill in a Project participation

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
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    level := *openapiclient.NewLevel() // Level | The Skill Level

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.UpdatePersonProjectSkill(context.Background(), personId, projectId, skillId).Level(level).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.UpdatePersonProjectSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePersonProjectSkill`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.UpdatePersonProjectSkill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**projectId** | **string** |  | 
**skillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePersonProjectSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **level** | [**Level**](Level.md) | The Skill Level | 

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


## UpdateProject

> ProjectDetails UpdateProject(ctx, projectId).Project(project).Execute()

Update a Project

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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    project := *openapiclient.NewProject(false, "Id_example") // Project | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectApi.UpdateProject(context.Background(), projectId).Project(project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.UpdateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProject`: ProjectDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.UpdateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectRequest struct via the builder pattern


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

