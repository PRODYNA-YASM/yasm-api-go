# \PersonApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPersonCertification**](PersonApi.md#AddPersonCertification) | **Post** /persons/{personId}/certifications/{certificationId} | Add Certification to a Person
[**AddPersonInterest**](PersonApi.md#AddPersonInterest) | **Post** /persons/{personId}/interests/skills/{skillId} | Add an Interest to a Person
[**AddPersonProject**](PersonApi.md#AddPersonProject) | **Post** /persons/{personId}/projects/{projectId} | Add Project to a Person
[**ConfirmSkill**](PersonApi.md#ConfirmSkill) | **Post** /persons/{personId}/projects/{projectId}/skills/{skillId}/confirmation/{confirmingPersonId} | Confirm Skill
[**CreatePerson**](PersonApi.md#CreatePerson) | **Post** /persons | Create a new Person
[**DeleteConfirmation**](PersonApi.md#DeleteConfirmation) | **Delete** /persons/{personId}/projects/{projectId}/skills/{skillId}/confirmation/{confirmingPersonId} | Remove a confirmation
[**DeletePersonCertification**](PersonApi.md#DeletePersonCertification) | **Delete** /persons/{personId}/certifications/{certificationId} | Remove an Interest to a Person
[**DeletePersonInterest**](PersonApi.md#DeletePersonInterest) | **Delete** /persons/{personId}/interests/skills/{skillId} | Remove an Interest to a Person
[**DeletePersonProject**](PersonApi.md#DeletePersonProject) | **Delete** /persons/{personId}/projects/{projectId} | Remove an Project from a Person
[**GeneratePersonProfile**](PersonApi.md#GeneratePersonProfile) | **Get** /persons/{personid}/profile | Generate a PDF profile from a Person
[**GetPerson**](PersonApi.md#GetPerson) | **Get** /persons/{personId} | Get basic info about a person
[**GetPersons**](PersonApi.md#GetPersons) | **Get** /persons | Get a list of all persons
[**UpdatePerson**](PersonApi.md#UpdatePerson) | **Put** /persons/{personId} | Update an existing Person
[**UpdatePersonCertification**](PersonApi.md#UpdatePersonCertification) | **Put** /persons/{personId}/certifications/{certificationId} | Update a Certification of a Person
[**UpdatePersonProject**](PersonApi.md#UpdatePersonProject) | **Put** /persons/{personId}/projects/{projectId} | Update a Project of a Person



## AddPersonCertification

> PersonDetails AddPersonCertification(ctx, personId, certificationId).Date(date).Execute()

Add Certification to a Person

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
    date := *openapiclient.NewDate() // Date | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PersonApi.AddPersonCertification(context.Background(), personId, certificationId).Date(date).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonApi.AddPersonCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPersonCertification`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonApi.AddPersonCertification`: %v\n", resp)
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


 **date** | [**Date**](Date.md) |  | 

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


## AddPersonInterest

> PersonDetails AddPersonInterest(ctx, personId, skillId).Person(person).Execute()

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
    person := *openapiclient.NewPerson("Id_example", "Darko Krizic") // Person | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PersonApi.AddPersonInterest(context.Background(), personId, skillId).Person(person).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonApi.AddPersonInterest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPersonInterest`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonApi.AddPersonInterest`: %v\n", resp)
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


 **person** | [**Person**](Person.md) |  | 

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


## AddPersonProject

> PersonDetails AddPersonProject(ctx, personId, projectId).SkillLevel(skillLevel).Execute()

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
    personId := TODO // string | 
    projectId := TODO // string | 
    skillLevel := []openapiclient.SkillLevel{*openapiclient.NewSkillLevel()} // []SkillLevel | A list of Skills each with a level

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PersonApi.AddPersonProject(context.Background(), personId, projectId).SkillLevel(skillLevel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonApi.AddPersonProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPersonProject`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonApi.AddPersonProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | [**string**](.md) |  | 
**projectId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPersonProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skillLevel** | [**[]SkillLevel**](SkillLevel.md) | A list of Skills each with a level | 

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
    resp, r, err := api_client.PersonApi.ConfirmSkill(context.Background(), personId, projectId, skillId, confirmingPersonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonApi.ConfirmSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmSkill`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonApi.ConfirmSkill`: %v\n", resp)
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


## CreatePerson

> PersonDetails CreatePerson(ctx).Person(person).Execute()

Create a new Person

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
    person := *openapiclient.NewPerson("Id_example", "Darko Krizic") // Person | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PersonApi.CreatePerson(context.Background()).Person(person).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonApi.CreatePerson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePerson`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonApi.CreatePerson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePersonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **person** | [**Person**](Person.md) |  | 

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
    resp, r, err := api_client.PersonApi.DeleteConfirmation(context.Background(), personId, projectId, skillId, confirmingPersonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonApi.DeleteConfirmation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteConfirmation`: Status
    fmt.Fprintf(os.Stdout, "Response from `PersonApi.DeleteConfirmation`: %v\n", resp)
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


## DeletePersonCertification

> Status DeletePersonCertification(ctx, personId, certificationId).Execute()

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
    resp, r, err := api_client.PersonApi.DeletePersonCertification(context.Background(), personId, certificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonApi.DeletePersonCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePersonCertification`: Status
    fmt.Fprintf(os.Stdout, "Response from `PersonApi.DeletePersonCertification`: %v\n", resp)
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

> Status DeletePersonInterest(ctx, personId, skillId).Execute()

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
    resp, r, err := api_client.PersonApi.DeletePersonInterest(context.Background(), personId, skillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonApi.DeletePersonInterest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePersonInterest`: Status
    fmt.Fprintf(os.Stdout, "Response from `PersonApi.DeletePersonInterest`: %v\n", resp)
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

[**Status**](Status.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePersonProject

> Status DeletePersonProject(ctx, personId, projectId).Execute()

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
    personId := TODO // string | 
    projectId := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PersonApi.DeletePersonProject(context.Background(), personId, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonApi.DeletePersonProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePersonProject`: Status
    fmt.Fprintf(os.Stdout, "Response from `PersonApi.DeletePersonProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | [**string**](.md) |  | 
**projectId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePersonProjectRequest struct via the builder pattern


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


## GeneratePersonProfile

> *os.File GeneratePersonProfile(ctx, personid).Execute()

Generate a PDF profile from a Person

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
    personid := TODO // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PersonApi.GeneratePersonProfile(context.Background(), personid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonApi.GeneratePersonProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneratePersonProfile`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `PersonApi.GeneratePersonProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personid** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGeneratePersonProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPerson

> PersonDetails GetPerson(ctx, personId).Execute()

Get basic info about a person

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PersonApi.GetPerson(context.Background(), personId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonApi.GetPerson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPerson`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonApi.GetPerson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonRequest struct via the builder pattern


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


## GetPersons

> PagedPersons GetPersons(ctx).Skip(skip).Limit(limit).Execute()

Get a list of all persons

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
    resp, r, err := api_client.PersonApi.GetPersons(context.Background()).Skip(skip).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonApi.GetPersons``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPersons`: PagedPersons
    fmt.Fprintf(os.Stdout, "Response from `PersonApi.GetPersons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**PagedPersons**](PagedPersons.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePerson

> PersonDetails UpdatePerson(ctx, personId).Person(person).Execute()

Update an existing Person

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
    person := *openapiclient.NewPerson("Id_example", "Darko Krizic") // Person | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PersonApi.UpdatePerson(context.Background(), personId).Person(person).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonApi.UpdatePerson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePerson`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonApi.UpdatePerson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePersonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **person** | [**Person**](Person.md) |  | 

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


## UpdatePersonCertification

> PersonDetails UpdatePersonCertification(ctx, personId, certificationId).Date(date).Execute()

Update a Certification of a Person

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
    date := *openapiclient.NewDate() // Date | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PersonApi.UpdatePersonCertification(context.Background(), personId, certificationId).Date(date).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonApi.UpdatePersonCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePersonCertification`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonApi.UpdatePersonCertification`: %v\n", resp)
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


 **date** | [**Date**](Date.md) |  | 

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


## UpdatePersonProject

> PersonDetails UpdatePersonProject(ctx, personId, projectId).SkillLevel(skillLevel).Execute()

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
    personId := TODO // string | 
    projectId := TODO // string | 
    skillLevel := []openapiclient.SkillLevel{*openapiclient.NewSkillLevel()} // []SkillLevel | A list of Skills each with a level

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PersonApi.UpdatePersonProject(context.Background(), personId, projectId).SkillLevel(skillLevel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonApi.UpdatePersonProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePersonProject`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonApi.UpdatePersonProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | [**string**](.md) |  | 
**projectId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePersonProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skillLevel** | [**[]SkillLevel**](SkillLevel.md) | A list of Skills each with a level | 

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

