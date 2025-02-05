# \AwardAPI

All URIs are relative to *https://yasm.prodyna.com:443/api/graph/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAwardToProjectParticipations**](AwardAPI.md#AddAwardToProjectParticipations) | **Post** /project-participations/awards/{awardId} | Add an award to Project Participations
[**CreateAward**](AwardAPI.md#CreateAward) | **Post** /awards | Create a new Award in an Organization
[**DeleteAward**](AwardAPI.md#DeleteAward) | **Delete** /awards/{awardId} | Delete an award
[**DeleteAwardFromProjectParticipations**](AwardAPI.md#DeleteAwardFromProjectParticipations) | **Delete** /project-participations/awards/{awardId} | Remove an Award from Project Participations
[**GetAward**](AwardAPI.md#GetAward) | **Get** /awards/{awardId} | Get details of an award by ID
[**GetPersonsForAwardInProject**](AwardAPI.md#GetPersonsForAwardInProject) | **Get** /projects/{projectId}/awards/{awardId}/persons | Get all persons that have won a certain award in a certain project
[**GetProjectsOfPersonWithAward**](AwardAPI.md#GetProjectsOfPersonWithAward) | **Get** /persons/{personId}/awards/{awardId}/projects | Get all projects where a certain award was won by a certain person
[**MoveAward**](AwardAPI.md#MoveAward) | **Put** /awards/{awardId}/organizations/{organizationId} | Move an award to an Organization
[**SearchAwards**](AwardAPI.md#SearchAwards) | **Post** /awards/search | Search for awards
[**UpdateAward**](AwardAPI.md#UpdateAward) | **Put** /awards/{awardId} | Update an award



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
    resp, r, err := apiClient.AwardAPI.AddAwardToProjectParticipations(context.Background(), awardId).InvolvedProjectParticipations(involvedProjectParticipations).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AwardAPI.AddAwardToProjectParticipations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAwardToProjectParticipations`: []ProjectParticipationDetails
    fmt.Fprintf(os.Stdout, "Response from `AwardAPI.AddAwardToProjectParticipations`: %v\n", resp)
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


## CreateAward

> AwardDetails CreateAward(ctx).OrganizationId(organizationId).Award(award).Execute()

Create a new Award in an Organization

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
    award := *openapiclient.NewAward("Id_example", "Name_example") // Award | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AwardAPI.CreateAward(context.Background()).OrganizationId(organizationId).Award(award).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AwardAPI.CreateAward``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAward`: AwardDetails
    fmt.Fprintf(os.Stdout, "Response from `AwardAPI.CreateAward`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAwardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **string** | The ID of the organization | 
 **award** | [**Award**](Award.md) |  | 

### Return type

[**AwardDetails**](AwardDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAward

> Status DeleteAward(ctx, awardId).Execute()

Delete an award

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AwardAPI.DeleteAward(context.Background(), awardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AwardAPI.DeleteAward``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAward`: Status
    fmt.Fprintf(os.Stdout, "Response from `AwardAPI.DeleteAward`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAwardRequest struct via the builder pattern


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
    resp, r, err := apiClient.AwardAPI.DeleteAwardFromProjectParticipations(context.Background(), awardId).InvolvedProjectParticipations(involvedProjectParticipations).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AwardAPI.DeleteAwardFromProjectParticipations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAwardFromProjectParticipations`: []ProjectParticipationDetails
    fmt.Fprintf(os.Stdout, "Response from `AwardAPI.DeleteAwardFromProjectParticipations`: %v\n", resp)
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


## GetAward

> AwardDetails GetAward(ctx, awardId).Execute()

Get details of an award by ID

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AwardAPI.GetAward(context.Background(), awardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AwardAPI.GetAward``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAward`: AwardDetails
    fmt.Fprintf(os.Stdout, "Response from `AwardAPI.GetAward`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAwardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AwardDetails**](AwardDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPersonsForAwardInProject

> PagedPersons GetPersonsForAwardInProject(ctx, projectId, awardId).Execute()

Get all persons that have won a certain award in a certain project

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
    awardId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AwardAPI.GetPersonsForAwardInProject(context.Background(), projectId, awardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AwardAPI.GetPersonsForAwardInProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPersonsForAwardInProject`: PagedPersons
    fmt.Fprintf(os.Stdout, "Response from `AwardAPI.GetPersonsForAwardInProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**awardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonsForAwardInProjectRequest struct via the builder pattern


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


## GetProjectsOfPersonWithAward

> PagedProjects GetProjectsOfPersonWithAward(ctx, personId, awardId).Execute()

Get all projects where a certain award was won by a certain person

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
    awardId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AwardAPI.GetProjectsOfPersonWithAward(context.Background(), personId, awardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AwardAPI.GetProjectsOfPersonWithAward``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectsOfPersonWithAward`: PagedProjects
    fmt.Fprintf(os.Stdout, "Response from `AwardAPI.GetProjectsOfPersonWithAward`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**awardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsOfPersonWithAwardRequest struct via the builder pattern


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


## MoveAward

> AwardDetails MoveAward(ctx, awardId, organizationId).Execute()

Move an award to an Organization

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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AwardAPI.MoveAward(context.Background(), awardId, organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AwardAPI.MoveAward``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoveAward`: AwardDetails
    fmt.Fprintf(os.Stdout, "Response from `AwardAPI.MoveAward`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awardId** | **string** |  | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveAwardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AwardDetails**](AwardDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAwards

> PagedAwards SearchAwards(ctx).AwardSearch(awardSearch).Execute()

Search for awards

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
    awardSearch := *openapiclient.NewAwardSearch(int32(123), int32(123)) // AwardSearch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AwardAPI.SearchAwards(context.Background()).AwardSearch(awardSearch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AwardAPI.SearchAwards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchAwards`: PagedAwards
    fmt.Fprintf(os.Stdout, "Response from `AwardAPI.SearchAwards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAwardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awardSearch** | [**AwardSearch**](AwardSearch.md) |  | 

### Return type

[**PagedAwards**](PagedAwards.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAward

> AwardDetails UpdateAward(ctx, awardId).Award(award).Execute()

Update an award

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
    award := *openapiclient.NewAward("Id_example", "Name_example") // Award | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AwardAPI.UpdateAward(context.Background(), awardId).Award(award).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AwardAPI.UpdateAward``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAward`: AwardDetails
    fmt.Fprintf(os.Stdout, "Response from `AwardAPI.UpdateAward`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAwardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **award** | [**Award**](Award.md) |  | 

### Return type

[**AwardDetails**](AwardDetails.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

