# \ProjectAPI

All URIs are relative to *https://yasm.prodyna.com:443/api/graph/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAwardToProjectParticipations**](ProjectAPI.md#AddAwardToProjectParticipations) | **Post** /project-participations/awards/{awardId} | Add an award to Project Participations
[**AddExecutiveOrganizationToProject**](ProjectAPI.md#AddExecutiveOrganizationToProject) | **Post** /projects/{projectId}/executive-organizations/{organizationId} | Add an Organization to a Project as executive organization
[**AddProjectParticipation**](ProjectAPI.md#AddProjectParticipation) | **Post** /project-participations | Add Project to a Person
[**CreateProject**](ProjectAPI.md#CreateProject) | **Post** /projects | Create a Project in an Organization
[**DeleteAwardFromProjectParticipations**](ProjectAPI.md#DeleteAwardFromProjectParticipations) | **Delete** /project-participations/awards/{awardId} | Remove an Award from Project Participations
[**DeleteProject**](ProjectAPI.md#DeleteProject) | **Delete** /projects/{projectId} | Delete a project
[**DeleteProjectParticipation**](ProjectAPI.md#DeleteProjectParticipation) | **Delete** /project-participations/{projectParticipationId} | Remove an Project from a Person
[**GetPersonsForSkillInProject**](ProjectAPI.md#GetPersonsForSkillInProject) | **Get** /projects/{projectId}/skills/{skillId}/persons | Get all persons that use a certain skill in a certain project
[**GetProject**](ProjectAPI.md#GetProject) | **Get** /projects/{projectId} | Get details about a Project
[**GetProjectsOfPersonWithSkill**](ProjectAPI.md#GetProjectsOfPersonWithSkill) | **Get** /persons/{personId}/skills/{skillId}/projects | Get all projects where a certain skill was used by a certain person
[**MergeProjects**](ProjectAPI.md#MergeProjects) | **Put** /projects/{projectId}/merge-with/{sourceProjectId} | Merge source project into main project
[**MoveProject**](ProjectAPI.md#MoveProject) | **Put** /projects/{projectId}/organizations/{organizationId} | Move a Project to an Organization
[**ReadProjectParticipation**](ProjectAPI.md#ReadProjectParticipation) | **Get** /project-participations/{projectParticipationId} | Get a project participation
[**RemoveExecutiveOrganizationFromProject**](ProjectAPI.md#RemoveExecutiveOrganizationFromProject) | **Delete** /projects/{projectId}/executive-organizations/{organizationId} | Remove an Organization from a Project as executive organization
[**SearchProjectParticipations**](ProjectAPI.md#SearchProjectParticipations) | **Post** /project-participations/search | Search over project participations
[**SearchProjects**](ProjectAPI.md#SearchProjects) | **Post** /projects/search | Complex search over project entities
[**UpdateProject**](ProjectAPI.md#UpdateProject) | **Put** /projects/{projectId} | Update a Project
[**UpdateProjectParticipation**](ProjectAPI.md#UpdateProjectParticipation) | **Put** /project-participations/{projectParticipationId} | Update a Project of a Person



## AddAwardToProjectParticipations

> []ProjectParticipationDetails AddAwardToProjectParticipations(ctx, awardId).InvolvedProjectParticipations(involvedProjectParticipations).Execute()

Add an award to Project Participations

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
    awardId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    involvedProjectParticipations := *openapiclient.NewInvolvedProjectParticipations() // InvolvedProjectParticipations | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectAPI.AddAwardToProjectParticipations(context.Background(), awardId).InvolvedProjectParticipations(involvedProjectParticipations).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.AddAwardToProjectParticipations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAwardToProjectParticipations`: []ProjectParticipationDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.AddAwardToProjectParticipations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAwardToProjectParticipationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **involvedProjectParticipations** | [**InvolvedProjectParticipations**](InvolvedProjectParticipations.md) |  | 

### Return type

[**[]ProjectParticipationDetails**](ProjectParticipationDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    project := *openapiclient.NewProject("Id_example", "Name_example") // Project | 

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


## DeleteAwardFromProjectParticipations

> []ProjectParticipationDetails DeleteAwardFromProjectParticipations(ctx, awardId).InvolvedProjectParticipations(involvedProjectParticipations).Execute()

Remove an Award from Project Participations

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
    awardId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    involvedProjectParticipations := *openapiclient.NewInvolvedProjectParticipations() // InvolvedProjectParticipations | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectAPI.DeleteAwardFromProjectParticipations(context.Background(), awardId).InvolvedProjectParticipations(involvedProjectParticipations).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.DeleteAwardFromProjectParticipations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAwardFromProjectParticipations`: []ProjectParticipationDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.DeleteAwardFromProjectParticipations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAwardFromProjectParticipationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **involvedProjectParticipations** | [**InvolvedProjectParticipations**](InvolvedProjectParticipations.md) |  | 

### Return type

[**[]ProjectParticipationDetails**](ProjectParticipationDetails.md)

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


## GetPersonsForSkillInProject

> PagedPersons GetPersonsForSkillInProject(ctx, projectId, skillId).Execute()

Get all persons that use a certain skill in a certain project

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
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectAPI.GetPersonsForSkillInProject(context.Background(), projectId, skillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetPersonsForSkillInProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPersonsForSkillInProject`: PagedPersons
    fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetPersonsForSkillInProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**skillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonsForSkillInProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PagedPersons**](PagedPersons.md)

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


## GetProjectsOfPersonWithSkill

> PagedProjects GetProjectsOfPersonWithSkill(ctx, personId, skillId).Execute()

Get all projects where a certain skill was used by a certain person

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
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectAPI.GetProjectsOfPersonWithSkill(context.Background(), personId, skillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProjectsOfPersonWithSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectsOfPersonWithSkill`: PagedProjects
    fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetProjectsOfPersonWithSkill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**skillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsOfPersonWithSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PagedProjects**](PagedProjects.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MergeProjects

> ProjectDetails MergeProjects(ctx, projectId, sourceProjectId).Project(project).Execute()

Merge source project into main project

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
    sourceProjectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    project := *openapiclient.NewProject("Id_example", "Name_example") // Project | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectAPI.MergeProjects(context.Background(), projectId, sourceProjectId).Project(project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.MergeProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MergeProjects`: ProjectDetails
    fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.MergeProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**sourceProjectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMergeProjectsRequest struct via the builder pattern


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
    project := *openapiclient.NewProject("Id_example", "Name_example") // Project | 

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

