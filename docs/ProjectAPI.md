# \ProjectAPI

All URIs are relative to *https://yasm.prodyna.com:443/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddExecutiveOrganizationToProject**](ProjectAPI.md#AddExecutiveOrganizationToProject) | **Post** /projects/{projectId}/executive-organizations/{organizationId} | Add an Organization to a Project as executive organization
[**AddProjectParticipation**](ProjectAPI.md#AddProjectParticipation) | **Post** /project-participations | Add Project to a Person
[**AddSkillConfirmation**](ProjectAPI.md#AddSkillConfirmation) | **Post** /project-participations/{projectParticipationId}/skills/{skillId}/confirmation/{confirmingPersonId} | Confirm Skill
[**CreateProject**](ProjectAPI.md#CreateProject) | **Post** /projects | Create a Project in an Organization
[**DeleteProject**](ProjectAPI.md#DeleteProject) | **Delete** /projects/{projectId} | Delete a project
[**DeleteProjectParticipation**](ProjectAPI.md#DeleteProjectParticipation) | **Delete** /project-participations/{projectParticipationId} | Remove an Project from a Person
[**GetProject**](ProjectAPI.md#GetProject) | **Get** /projects/{projectId} | Get details about a Project
[**MoveProject**](ProjectAPI.md#MoveProject) | **Put** /projects/{projectId}/organizations/{organizationId} | Move a Project to an Organization
[**ReadProjectParticipation**](ProjectAPI.md#ReadProjectParticipation) | **Get** /project-participations/{projectParticipationId} | Get a project participation
[**RemoveExecutiveOrganizationFromProject**](ProjectAPI.md#RemoveExecutiveOrganizationFromProject) | **Delete** /projects/{projectId}/executive-organizations/{organizationId} | Remove an Organization from a Project as executive organization
[**RemoveSkillConfirmation**](ProjectAPI.md#RemoveSkillConfirmation) | **Delete** /project-participations/{projectParticipationId}/skills/{skillId}/confirmation/{confirmingPersonId} | Remove a confirmation
[**SearchProjectParticipations**](ProjectAPI.md#SearchProjectParticipations) | **Post** /project-participations/search | Search over project participations
[**SearchProjects**](ProjectAPI.md#SearchProjects) | **Post** /projects/search | Complex search over project entities
[**UpdateProject**](ProjectAPI.md#UpdateProject) | **Put** /projects/{projectId} | Update a Project
[**UpdateProjectParticipation**](ProjectAPI.md#UpdateProjectParticipation) | **Put** /project-participations/{projectParticipationId} | Update a Project of a Person



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
    resp, r, err := apiClient.ProjectAPI.AddExecutiveOrganizationToProject(context.Background(), projectId, organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.AddExecutiveOrganizationToProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddExecutiveOrganizationToProject`: ProjectDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.AddExecutiveOrganizationToProject`: %v\n", resp)
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


## AddProjectParticipation

> ProjectParticipationDetails AddProjectParticipation(ctx).ProjectParticipationCreate(projectParticipationCreate).Execute()

Add Project to a Person

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
    projectParticipationCreate := *openapiclient.NewProjectParticipationCreate(*openapiclient.NewTimeframed(time.Now()), "Id_example", "ProjectId_example", "PersonId_example") // ProjectParticipationCreate | List of Skills with level and timeframe

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectAPI.AddProjectParticipation(context.Background()).ProjectParticipationCreate(projectParticipationCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.AddProjectParticipation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddProjectParticipation`: ProjectParticipationDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.AddProjectParticipation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddProjectParticipationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectParticipationCreate** | [**ProjectParticipationCreate**](ProjectParticipationCreate.md) | List of Skills with level and timeframe | 

### Return type

[**ProjectParticipationDetails**](ProjectParticipationDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddSkillConfirmation

> ProjectParticipationDetails AddSkillConfirmation(ctx, projectParticipationId, skillId, confirmingPersonId).Execute()

Confirm Skill

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
    projectParticipationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    confirmingPersonId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectAPI.AddSkillConfirmation(context.Background(), projectParticipationId, skillId, confirmingPersonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.AddSkillConfirmation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSkillConfirmation`: ProjectParticipationDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.AddSkillConfirmation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectParticipationId** | **string** |  | 
**skillId** | **string** |  | 
**confirmingPersonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSkillConfirmationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ProjectParticipationDetails**](ProjectParticipationDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
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
    resp, r, err := apiClient.ProjectAPI.CreateProject(context.Background()).OrganizationId(organizationId).Project(project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.CreateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProject`: ProjectDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.CreateProject`: %v\n", resp)
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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectAPI.DeleteProject(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.DeleteProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteProject`: Status
    fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.DeleteProject`: %v\n", resp)
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

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProjectParticipation

> Status DeleteProjectParticipation(ctx, projectParticipationId).Execute()

Remove an Project from a Person

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
    projectParticipationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectAPI.DeleteProjectParticipation(context.Background(), projectParticipationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.DeleteProjectParticipation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteProjectParticipation`: Status
    fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.DeleteProjectParticipation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectParticipationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectParticipationRequest struct via the builder pattern


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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectAPI.GetProject(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProject`: ProjectDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetProject`: %v\n", resp)
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
    resp, r, err := apiClient.ProjectAPI.MoveProject(context.Background(), projectId, organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.MoveProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoveProject`: ProjectDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.MoveProject`: %v\n", resp)
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


## ReadProjectParticipation

> ProjectParticipationDetails ReadProjectParticipation(ctx, projectParticipationId).Execute()

Get a project participation

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
    projectParticipationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectAPI.ReadProjectParticipation(context.Background(), projectParticipationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.ReadProjectParticipation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadProjectParticipation`: ProjectParticipationDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.ReadProjectParticipation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectParticipationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadProjectParticipationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectParticipationDetails**](ProjectParticipationDetails.md)

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
    resp, r, err := apiClient.ProjectAPI.RemoveExecutiveOrganizationFromProject(context.Background(), projectId, organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.RemoveExecutiveOrganizationFromProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveExecutiveOrganizationFromProject`: ProjectDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.RemoveExecutiveOrganizationFromProject`: %v\n", resp)
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


## RemoveSkillConfirmation

> ProjectParticipationDetails RemoveSkillConfirmation(ctx, projectParticipationId, skillId, confirmingPersonId).Execute()

Remove a confirmation

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
    projectParticipationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    confirmingPersonId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectAPI.RemoveSkillConfirmation(context.Background(), projectParticipationId, skillId, confirmingPersonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.RemoveSkillConfirmation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveSkillConfirmation`: ProjectParticipationDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.RemoveSkillConfirmation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectParticipationId** | **string** |  | 
**skillId** | **string** |  | 
**confirmingPersonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSkillConfirmationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ProjectParticipationDetails**](ProjectParticipationDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchProjectParticipations

> PagedProjectParticipation SearchProjectParticipations(ctx).ProjectParticipationSearch(projectParticipationSearch).Execute()

Search over project participations

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
    projectParticipationSearch := *openapiclient.NewProjectParticipationSearch(int32(123), int32(123)) // ProjectParticipationSearch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectAPI.SearchProjectParticipations(context.Background()).ProjectParticipationSearch(projectParticipationSearch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.SearchProjectParticipations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchProjectParticipations`: PagedProjectParticipation
    fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.SearchProjectParticipations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchProjectParticipationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectParticipationSearch** | [**ProjectParticipationSearch**](ProjectParticipationSearch.md) |  | 

### Return type

[**PagedProjectParticipation**](PagedProjectParticipation.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchProjects

> PagedProjects SearchProjects(ctx).ProjectSearch(projectSearch).Execute()

Complex search over project entities

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
    projectSearch := *openapiclient.NewProjectSearch(int32(123), int32(123)) // ProjectSearch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectAPI.SearchProjects(context.Background()).ProjectSearch(projectSearch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.SearchProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchProjects`: PagedProjects
    fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.SearchProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectSearch** | [**ProjectSearch**](ProjectSearch.md) |  | 

### Return type

[**PagedProjects**](PagedProjects.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    project := *openapiclient.NewProject(false, "Id_example", "Name_example") // Project | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectAPI.UpdateProject(context.Background(), projectId).Project(project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.UpdateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProject`: ProjectDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.UpdateProject`: %v\n", resp)
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

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProjectParticipation

> ProjectParticipationDetails UpdateProjectParticipation(ctx, projectParticipationId).ProjectParticipationUpdate(projectParticipationUpdate).Execute()

Update a Project of a Person

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
    projectParticipationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    projectParticipationUpdate := *openapiclient.NewProjectParticipationUpdate(*openapiclient.NewTimeframed(time.Now())) // ProjectParticipationUpdate | A time frame and a list of Skills

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectAPI.UpdateProjectParticipation(context.Background(), projectParticipationId).ProjectParticipationUpdate(projectParticipationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.UpdateProjectParticipation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProjectParticipation`: ProjectParticipationDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.UpdateProjectParticipation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectParticipationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectParticipationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectParticipationUpdate** | [**ProjectParticipationUpdate**](ProjectParticipationUpdate.md) | A time frame and a list of Skills | 

### Return type

[**ProjectParticipationDetails**](ProjectParticipationDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

