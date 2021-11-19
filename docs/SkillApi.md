# \SkillApi

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPersonInterest**](SkillApi.md#AddPersonInterest) | **Post** /persons/{personId}/interests/skills/{skillId} | Add an Interest to a Person
[**AddPersonProjectSkill**](SkillApi.md#AddPersonProjectSkill) | **Post** /persons/{personId}/projects/{projectId}/skills/{skillId} | Add Skill to a Project participation
[**AddPersonSkillExperience**](SkillApi.md#AddPersonSkillExperience) | **Post** /persons/{personId}/experiences/skills/{skillId} | Add an Skill experience to a Person
[**AddPersonSkillExperiences**](SkillApi.md#AddPersonSkillExperiences) | **Post** /persons/{personId}/experiences | Add an Skill experience to a Person (bulk)
[**AddSkillToCertification**](SkillApi.md#AddSkillToCertification) | **Post** /certifications/{certificationId}/skills/{skillId} | 
[**AddSkillToParentSkill**](SkillApi.md#AddSkillToParentSkill) | **Post** /skills/{skillId}/parents/{parentSkillId} | Attach a Skill to a parent Skill, returns the parent Skill
[**ConfirmSkill**](SkillApi.md#ConfirmSkill) | **Post** /persons/{personId}/projects/{projectId}/skills/{skillId}/confirmation/{confirmingPersonId} | Confirm Skill
[**CreateSkill**](SkillApi.md#CreateSkill) | **Post** /skills | Create a Skill
[**DeleteConfirmation**](SkillApi.md#DeleteConfirmation) | **Delete** /persons/{personId}/projects/{projectId}/skills/{skillId}/confirmation/{confirmingPersonId} | Remove a confirmation
[**DeletePersonInterest**](SkillApi.md#DeletePersonInterest) | **Delete** /persons/{personId}/interests/skills/{skillId} | Remove an Interest to a Person
[**DeletePersonProjectSkill**](SkillApi.md#DeletePersonProjectSkill) | **Delete** /persons/{personId}/projects/{projectId}/skills/{skillId} | Remove a Skill from a Project participation
[**DeletePersonSkillExperience**](SkillApi.md#DeletePersonSkillExperience) | **Delete** /persons/{personId}/experiences/skills/{skillId} | Remove an Skill Experience to a Person
[**DeletePersonSkillExperiences**](SkillApi.md#DeletePersonSkillExperiences) | **Delete** /persons/{personId}/experiences | Remove an Skill Experience to a Person
[**DeleteSkill**](SkillApi.md#DeleteSkill) | **Delete** /skills/{skillId} | Delete a Skill
[**DeleteSkillFromCertification**](SkillApi.md#DeleteSkillFromCertification) | **Delete** /certifications/{certificationId}/skills/{skillId} | 
[**GetSkill**](SkillApi.md#GetSkill) | **Get** /skills/{skillId} | Get details for a single skill
[**GetSkillParents**](SkillApi.md#GetSkillParents) | **Get** /skills/{skillId}/parents | Get ghe list of parents for a skill
[**GetSkills**](SkillApi.md#GetSkills) | **Get** /skills | Get a list of all skills, optionally only root, optionally only kinds
[**RemoveSkillFromParentSkill**](SkillApi.md#RemoveSkillFromParentSkill) | **Delete** /skills/{skillId}/parents/{parentSkillId} | Detaches a Skill from parent Skill, return the parent Skill
[**UpdatePersonProjectSkill**](SkillApi.md#UpdatePersonProjectSkill) | **Put** /persons/{personId}/projects/{projectId}/skills/{skillId} | Update the level of a Skill in a Project participation
[**UpdatePersonSkillExperience**](SkillApi.md#UpdatePersonSkillExperience) | **Put** /persons/{personId}/experiences/skills/{skillId} | Edit an Skill experience to a Person
[**UpdatePersonSkillExperiences**](SkillApi.md#UpdatePersonSkillExperiences) | **Put** /persons/{personId}/experiences | Edit an Skill experience to a Person
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
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

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
**personId** | **string** |  | 
**skillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPersonInterestRequest struct via the builder pattern


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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SkillApi.AddPersonProjectSkill(context.Background(), personId, projectId, skillId).Level(level).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillApi.AddPersonProjectSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPersonProjectSkill`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `SkillApi.AddPersonProjectSkill`: %v\n", resp)
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


## AddPersonSkillExperience

> PersonDetails AddPersonSkillExperience(ctx, personId, skillId).Level(level).Execute()

Add an Skill experience to a Person

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
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    level := *openapiclient.NewLevel() // Level | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SkillApi.AddPersonSkillExperience(context.Background(), personId, skillId).Level(level).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillApi.AddPersonSkillExperience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPersonSkillExperience`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `SkillApi.AddPersonSkillExperience`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**skillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPersonSkillExperienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **level** | [**Level**](Level.md) |  | 

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


## AddPersonSkillExperiences

> PersonDetails AddPersonSkillExperiences(ctx, personId).SkillLevelUpdate(skillLevelUpdate).Execute()

Add an Skill experience to a Person (bulk)

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
    skillLevelUpdate := []openapiclient.SkillLevelUpdate{*openapiclient.NewSkillLevelUpdate()} // []SkillLevelUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SkillApi.AddPersonSkillExperiences(context.Background(), personId).SkillLevelUpdate(skillLevelUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillApi.AddPersonSkillExperiences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPersonSkillExperiences`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `SkillApi.AddPersonSkillExperiences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPersonSkillExperiencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skillLevelUpdate** | [**[]SkillLevelUpdate**](SkillLevelUpdate.md) |  | 

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
    certificationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
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

[bearerAuth](../README.md#bearerAuth)

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
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    parentSkillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

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
**skillId** | **string** |  | 
**parentSkillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSkillToParentSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SkillDetails**](SkillDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

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
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    confirmingPersonId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

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
    skill := *openapiclient.NewSkill("Id_example", "This is the name", false) // Skill | 

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SkillApi.DeleteConfirmation(context.Background(), personId, projectId, skillId, confirmingPersonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillApi.DeleteConfirmation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteConfirmation`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `SkillApi.DeleteConfirmation`: %v\n", resp)
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
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

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
**personId** | **string** |  | 
**skillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePersonInterestRequest struct via the builder pattern


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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SkillApi.DeletePersonProjectSkill(context.Background(), personId, projectId, skillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillApi.DeletePersonProjectSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePersonProjectSkill`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `SkillApi.DeletePersonProjectSkill`: %v\n", resp)
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


## DeletePersonSkillExperience

> PersonDetails DeletePersonSkillExperience(ctx, personId, skillId).Execute()

Remove an Skill Experience to a Person

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
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SkillApi.DeletePersonSkillExperience(context.Background(), personId, skillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillApi.DeletePersonSkillExperience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePersonSkillExperience`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `SkillApi.DeletePersonSkillExperience`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**skillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePersonSkillExperienceRequest struct via the builder pattern


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


## DeletePersonSkillExperiences

> PersonDetails DeletePersonSkillExperiences(ctx, personId).RequestBody(requestBody).Execute()

Remove an Skill Experience to a Person

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
    requestBody := []string{"Property_example"} // []string | A list of skillIds

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SkillApi.DeletePersonSkillExperiences(context.Background(), personId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillApi.DeletePersonSkillExperiences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePersonSkillExperiences`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `SkillApi.DeletePersonSkillExperiences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePersonSkillExperiencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** | A list of skillIds | 

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
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

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
**skillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSkillRequest struct via the builder pattern


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
    certificationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

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
**certificationId** | **string** |  | 
**skillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSkillFromCertificationRequest struct via the builder pattern


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
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

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
**skillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SkillDetails**](SkillDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

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
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
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
**skillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSkillParentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**PagedSkills**](PagedSkills.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSkills

> PagedSkills GetSkills(ctx).Types(types).Suggestions(suggestions).Skip(skip).Limit(limit).Execute()

Get a list of all skills, optionally only root, optionally only kinds

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
    types := "types_example" // string | Gives you either all skills, only the root kills or those which are defining kinds (optional) (default to "all")
    suggestions := "suggestions_example" // string | Optionally filter skills based on suggestion (optional) (default to "all")
    skip := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SkillApi.GetSkills(context.Background()).Types(types).Suggestions(suggestions).Skip(skip).Limit(limit).Execute()
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
 **types** | **string** | Gives you either all skills, only the root kills or those which are defining kinds | [default to &quot;all&quot;]
 **suggestions** | **string** | Optionally filter skills based on suggestion | [default to &quot;all&quot;]
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**PagedSkills**](PagedSkills.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

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
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    parentSkillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

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
**skillId** | **string** |  | 
**parentSkillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSkillFromParentSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SkillDetails**](SkillDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SkillApi.UpdatePersonProjectSkill(context.Background(), personId, projectId, skillId).Level(level).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillApi.UpdatePersonProjectSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePersonProjectSkill`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `SkillApi.UpdatePersonProjectSkill`: %v\n", resp)
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


## UpdatePersonSkillExperience

> PersonDetails UpdatePersonSkillExperience(ctx, personId, skillId).Level(level).Execute()

Edit an Skill experience to a Person

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
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    level := *openapiclient.NewLevel() // Level | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SkillApi.UpdatePersonSkillExperience(context.Background(), personId, skillId).Level(level).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillApi.UpdatePersonSkillExperience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePersonSkillExperience`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `SkillApi.UpdatePersonSkillExperience`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**skillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePersonSkillExperienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **level** | [**Level**](Level.md) |  | 

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


## UpdatePersonSkillExperiences

> PersonDetails UpdatePersonSkillExperiences(ctx, personId).SkillLevelUpdate(skillLevelUpdate).Execute()

Edit an Skill experience to a Person

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
    skillLevelUpdate := []openapiclient.SkillLevelUpdate{*openapiclient.NewSkillLevelUpdate()} // []SkillLevelUpdate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SkillApi.UpdatePersonSkillExperiences(context.Background(), personId).SkillLevelUpdate(skillLevelUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkillApi.UpdatePersonSkillExperiences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePersonSkillExperiences`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `SkillApi.UpdatePersonSkillExperiences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePersonSkillExperiencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skillLevelUpdate** | [**[]SkillLevelUpdate**](SkillLevelUpdate.md) |  | 

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
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skill := *openapiclient.NewSkill("Id_example", "This is the name", false) // Skill | 

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
**skillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skill** | [**Skill**](Skill.md) |  | 

### Return type

[**SkillDetails**](SkillDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

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
    certificationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
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

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

