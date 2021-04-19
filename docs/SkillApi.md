# \SkillApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPersonInterest**](SkillApi.md#AddPersonInterest) | **Post** /persons/{personId}/interests/skills/{skillId} | Add an Interest to a Person
[**AddSkillToCertification**](SkillApi.md#AddSkillToCertification) | **Post** /certifications/{certificationId}/skills/{skillId} | 
[**AddSkillToParentSkill**](SkillApi.md#AddSkillToParentSkill) | **Post** /skills/{skillId}/parents/{parentSkillId} | Attach a Skill to a parent Skill, returns the parent Skill
[**ConfirmSkill**](SkillApi.md#ConfirmSkill) | **Post** /persons/{personId}/projects/{projectId}/skills/{skillId}/confirmation/{confirmingPersonId} | Confirm Skill
[**CreateSkill**](SkillApi.md#CreateSkill) | **Post** /skills | Create a Skill
[**DeleteConfirmation**](SkillApi.md#DeleteConfirmation) | **Delete** /persons/{personId}/projects/{projectId}/skills/{skillId}/confirmation/{confirmingPersonId} | Remove a confirmation
[**DeletePersonInterest**](SkillApi.md#DeletePersonInterest) | **Delete** /persons/{personId}/interests/skills/{skillId} | Remove an Interest to a Person
[**DeleteSkill**](SkillApi.md#DeleteSkill) | **Delete** /skills/{skillId} | Delete a Skill
[**DeleteSkillFromCertification**](SkillApi.md#DeleteSkillFromCertification) | **Delete** /certifications/{certificationId}/skills/{skillId} | 
[**GetSkill**](SkillApi.md#GetSkill) | **Get** /skills/{skillId} | Get details for a single skill
[**GetSkillParents**](SkillApi.md#GetSkillParents) | **Get** /skills/{skillId}/parents | Get ghe list of parents for a skill
[**GetSkills**](SkillApi.md#GetSkills) | **Get** /skills | Get a list of all skills, optionally only root
[**RemoveSkillFromParentSkill**](SkillApi.md#RemoveSkillFromParentSkill) | **Delete** /skills/{skillId}/parents/{parentSkillId} | Detaches a Skill from parent Skill, return the parent Skill
[**UpdateSkill**](SkillApi.md#UpdateSkill) | **Put** /skills/{skillId} | Update a Skill
[**UpdateSkillInCertification**](SkillApi.md#UpdateSkillInCertification) | **Put** /certifications/{certificationId}/skills/{skillId} | 



## AddPersonInterest

> PersonDetails AddPersonInterest(ctx, personId, skillId).Execute()

Add an Interest to a Person

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
    skillId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SkillApi.AddPersonInterest(context.Background(), personId, skillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillApi.AddPersonInterest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPersonInterest`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `SkillApi.AddPersonInterest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | [**string**](.md) |  | 
**skillId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPersonInterestRequest struct via the builder pattern


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
    resp, r, err := api_client.SkillApi.AddSkillToCertification(context.Background(), certificationId, skillId).Level(level).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillApi.AddSkillToCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSkillToCertification`: CertificationDetails
    fmt.Fprintf(os.Stdout, "Response from `SkillApi.AddSkillToCertification`: %v\n", resp)
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


## AddSkillToParentSkill

> SkillDetails AddSkillToParentSkill(ctx, skillId, parentSkillId).Execute()

Attach a Skill to a parent Skill, returns the parent Skill

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
    skillId := TODO // string | 
    parentSkillId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SkillApi.AddSkillToParentSkill(context.Background(), skillId, parentSkillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillApi.AddSkillToParentSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSkillToParentSkill`: SkillDetails
    fmt.Fprintf(os.Stdout, "Response from `SkillApi.AddSkillToParentSkill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillId** | [**string**](.md) |  | 
**parentSkillId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSkillToParentSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SkillDetails**](SkillDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
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
    personId := TODO // string | 
    projectId := TODO // string | 
    skillId := TODO // string | 
    confirmingPersonId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SkillApi.ConfirmSkill(context.Background(), personId, projectId, skillId, confirmingPersonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillApi.ConfirmSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmSkill`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `SkillApi.ConfirmSkill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | [**string**](.md) |  | 
**projectId** | [**string**](.md) |  | 
**skillId** | [**string**](.md) |  | 
**confirmingPersonId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmSkillRequest struct via the builder pattern


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


## CreateSkill

> SkillDetails CreateSkill(ctx).Skill(skill).Execute()

Create a Skill

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
    skill := *openapiclient.NewSkill("Id_example", "This is the name", []string{"Java EE"}) // Skill | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SkillApi.CreateSkill(context.Background()).Skill(skill).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillApi.CreateSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSkill`: SkillDetails
    fmt.Fprintf(os.Stdout, "Response from `SkillApi.CreateSkill`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skill** | [**Skill**](Skill.md) |  | 

### Return type

[**SkillDetails**](SkillDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfirmation

> Status DeleteConfirmation(ctx, personId, projectId, skillId, confirmingPersonId).Execute()

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
    personId := TODO // string | 
    projectId := TODO // string | 
    skillId := TODO // string | 
    confirmingPersonId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SkillApi.DeleteConfirmation(context.Background(), personId, projectId, skillId, confirmingPersonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillApi.DeleteConfirmation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteConfirmation`: Status
    fmt.Fprintf(os.Stdout, "Response from `SkillApi.DeleteConfirmation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | [**string**](.md) |  | 
**projectId** | [**string**](.md) |  | 
**skillId** | [**string**](.md) |  | 
**confirmingPersonId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfirmationRequest struct via the builder pattern


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


## DeletePersonInterest

> PersonDetails DeletePersonInterest(ctx, personId, skillId).Execute()

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
    skillId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SkillApi.DeletePersonInterest(context.Background(), personId, skillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillApi.DeletePersonInterest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePersonInterest`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `SkillApi.DeletePersonInterest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | [**string**](.md) |  | 
**skillId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePersonInterestRequest struct via the builder pattern


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


## DeleteSkill

> Status DeleteSkill(ctx, skillId).Execute()

Delete a Skill

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
    skillId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SkillApi.DeleteSkill(context.Background(), skillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillApi.DeleteSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSkill`: Status
    fmt.Fprintf(os.Stdout, "Response from `SkillApi.DeleteSkill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSkillRequest struct via the builder pattern


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
    resp, r, err := api_client.SkillApi.DeleteSkillFromCertification(context.Background(), certificationId, skillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillApi.DeleteSkillFromCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSkillFromCertification`: CertificationDetails
    fmt.Fprintf(os.Stdout, "Response from `SkillApi.DeleteSkillFromCertification`: %v\n", resp)
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


## GetSkill

> SkillDetails GetSkill(ctx, skillId).Execute()

Get details for a single skill

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
    skillId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SkillApi.GetSkill(context.Background(), skillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillApi.GetSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSkill`: SkillDetails
    fmt.Fprintf(os.Stdout, "Response from `SkillApi.GetSkill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SkillDetails**](SkillDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSkillParents

> PagedSkills GetSkillParents(ctx, skillId).Skip(skip).Limit(limit).Execute()

Get ghe list of parents for a skill

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
    skillId := TODO // string | 
    skip := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SkillApi.GetSkillParents(context.Background(), skillId).Skip(skip).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillApi.GetSkillParents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSkillParents`: PagedSkills
    fmt.Fprintf(os.Stdout, "Response from `SkillApi.GetSkillParents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSkillParentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**PagedSkills**](PagedSkills.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSkills

> PagedSkills GetSkills(ctx).RootOnly(rootOnly).Skip(skip).Limit(limit).Execute()

Get a list of all skills, optionally only root

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
    rootOnly := true // bool | Return only root skills (optional) (default to false)
    skip := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SkillApi.GetSkills(context.Background()).RootOnly(rootOnly).Skip(skip).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillApi.GetSkills``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSkills`: PagedSkills
    fmt.Fprintf(os.Stdout, "Response from `SkillApi.GetSkills`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSkillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rootOnly** | **bool** | Return only root skills | [default to false]
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**PagedSkills**](PagedSkills.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveSkillFromParentSkill

> SkillDetails RemoveSkillFromParentSkill(ctx, skillId, parentSkillId).Execute()

Detaches a Skill from parent Skill, return the parent Skill

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
    skillId := TODO // string | 
    parentSkillId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SkillApi.RemoveSkillFromParentSkill(context.Background(), skillId, parentSkillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillApi.RemoveSkillFromParentSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveSkillFromParentSkill`: SkillDetails
    fmt.Fprintf(os.Stdout, "Response from `SkillApi.RemoveSkillFromParentSkill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillId** | [**string**](.md) |  | 
**parentSkillId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSkillFromParentSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SkillDetails**](SkillDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSkill

> SkillDetails UpdateSkill(ctx, skillId).Skill(skill).Execute()

Update a Skill

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
    skillId := TODO // string | 
    skill := *openapiclient.NewSkill("Id_example", "This is the name", []string{"Java EE"}) // Skill | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SkillApi.UpdateSkill(context.Background(), skillId).Skill(skill).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillApi.UpdateSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSkill`: SkillDetails
    fmt.Fprintf(os.Stdout, "Response from `SkillApi.UpdateSkill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skillId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skill** | [**Skill**](Skill.md) |  | 

### Return type

[**SkillDetails**](SkillDetails.md)

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
    resp, r, err := api_client.SkillApi.UpdateSkillInCertification(context.Background(), certificationId, skillId).Level(level).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillApi.UpdateSkillInCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSkillInCertification`: CertificationDetails
    fmt.Fprintf(os.Stdout, "Response from `SkillApi.UpdateSkillInCertification`: %v\n", resp)
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

