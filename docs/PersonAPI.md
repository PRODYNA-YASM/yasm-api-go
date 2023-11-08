# \PersonAPI

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPersonCertification**](PersonAPI.md#AddPersonCertification) | **Post** /persons/{personId}/certifications/{certificationId} | Add Certification to a Person
[**AddPersonInterest**](PersonAPI.md#AddPersonInterest) | **Post** /persons/{personId}/interests/skills/{skillId} | Add an Interest to a Person
[**AddPersonLanguage**](PersonAPI.md#AddPersonLanguage) | **Post** /persons/{personId}/languages/{languageId} | Assign a language to the person
[**AddPersonOffice**](PersonAPI.md#AddPersonOffice) | **Post** /persons/{personId}/offices/{officeId} | Assign a person to an office
[**AddPersonProfile**](PersonAPI.md#AddPersonProfile) | **Post** /persons/{personId}/profiles/{profileId} | Add a Profile to a Person
[**AddPersonProject**](PersonAPI.md#AddPersonProject) | **Post** /persons/{personId}/projects/{projectId} | Add Project to a Person
[**AddPersonProjectSkill**](PersonAPI.md#AddPersonProjectSkill) | **Post** /persons/{personId}/projects/{projectId}/skills/{skillId} | Add Skill to a Project participation
[**AddPersonSkillExperience**](PersonAPI.md#AddPersonSkillExperience) | **Post** /persons/{personId}/experiences/skills/{skillId} | Add an Skill experience to a Person
[**AddPersonSkillExperiences**](PersonAPI.md#AddPersonSkillExperiences) | **Post** /persons/{personId}/experiences | Add an Skill experience to a Person (bulk)
[**ConfirmSkill**](PersonAPI.md#ConfirmSkill) | **Post** /persons/{personId}/projects/{projectId}/skills/{skillId}/confirmation/{confirmingPersonId} | Confirm Skill
[**CreateAvailability**](PersonAPI.md#CreateAvailability) | **Post** /persons/{personId}/availabilities | Create a availability for a person
[**CreatePerson**](PersonAPI.md#CreatePerson) | **Post** /persons | Create a new Person
[**DeleteAvailability**](PersonAPI.md#DeleteAvailability) | **Delete** /persons/{personId}/availabilities/{availabilityId} | Delete a person availability
[**DeleteConfirmation**](PersonAPI.md#DeleteConfirmation) | **Delete** /persons/{personId}/projects/{projectId}/skills/{skillId}/confirmation/{confirmingPersonId} | Remove a confirmation
[**DeletePerson**](PersonAPI.md#DeletePerson) | **Delete** /persons/{personId} | Delete an existing Person
[**DeletePersonCertification**](PersonAPI.md#DeletePersonCertification) | **Delete** /persons/{personId}/certifications/{certificationId} | Remove an Interest to a Person
[**DeletePersonInterest**](PersonAPI.md#DeletePersonInterest) | **Delete** /persons/{personId}/interests/skills/{skillId} | Remove an Interest to a Person
[**DeletePersonOffice**](PersonAPI.md#DeletePersonOffice) | **Delete** /persons/{personId}/offices/{officeId} | Delete the office from a Person
[**DeletePersonProfile**](PersonAPI.md#DeletePersonProfile) | **Delete** /persons/{personId}/profiles/{profileId} | Remove a Profile from a Person
[**DeletePersonProject**](PersonAPI.md#DeletePersonProject) | **Delete** /persons/{personId}/projects/{projectId} | Remove an Project from a Person
[**DeletePersonProjectSkill**](PersonAPI.md#DeletePersonProjectSkill) | **Delete** /persons/{personId}/projects/{projectId}/skills/{skillId} | Remove a Skill from a Project participation
[**DeletePersonSkillExperience**](PersonAPI.md#DeletePersonSkillExperience) | **Delete** /persons/{personId}/experiences/skills/{skillId} | Remove an Skill Experience to a Person
[**DeletePersonSkillExperiences**](PersonAPI.md#DeletePersonSkillExperiences) | **Delete** /persons/{personId}/experiences | Remove an Skill Experience to a Person
[**GeneratePersonProfile**](PersonAPI.md#GeneratePersonProfile) | **Get** /persons/{personId}/pdf-profile | Generate a PDF profile from a Person
[**GetAllBusinessDepartments**](PersonAPI.md#GetAllBusinessDepartments) | **Get** /persons/departments | Get all unique business departments
[**GetAvailabilities**](PersonAPI.md#GetAvailabilities) | **Get** /persons/{personId}/availabilities | Get a list of all activities for a person
[**GetPerson**](PersonAPI.md#GetPerson) | **Get** /persons/{personId} | Get basic info about a person
[**ReadPersonPicture**](PersonAPI.md#ReadPersonPicture) | **Get** /persons/{personId}/picture | Read person image
[**ReadPersonProject**](PersonAPI.md#ReadPersonProject) | **Get** /persons/{personId}/projects/{projectId} | Get a Project Partifipation of a Person
[**RemovePersonLanguage**](PersonAPI.md#RemovePersonLanguage) | **Delete** /persons/{personId}/languages/{languageId} | Remove a language from a person
[**SearchPersons**](PersonAPI.md#SearchPersons) | **Post** /persons/search | Complex search over person entities
[**UpdateAvailability**](PersonAPI.md#UpdateAvailability) | **Put** /persons/{personId}/availabilities/{availabilityId} | Update a person availability
[**UpdatePerson**](PersonAPI.md#UpdatePerson) | **Put** /persons/{personId} | Update an existing Person
[**UpdatePersonCertification**](PersonAPI.md#UpdatePersonCertification) | **Put** /persons/{personId}/certifications/{certificationId} | Update a Certification of a Person
[**UpdatePersonPicture**](PersonAPI.md#UpdatePersonPicture) | **Put** /persons/{personId}/picture | Update person image
[**UpdatePersonProject**](PersonAPI.md#UpdatePersonProject) | **Put** /persons/{personId}/projects/{projectId} | Update a Project of a Person
[**UpdatePersonProjectSkill**](PersonAPI.md#UpdatePersonProjectSkill) | **Put** /persons/{personId}/projects/{projectId}/skills/{skillId} | Update the level of a Skill in a Project participation
[**UpdatePersonSkillExperience**](PersonAPI.md#UpdatePersonSkillExperience) | **Put** /persons/{personId}/experiences/skills/{skillId} | Edit an Skill experience to a Person
[**UpdatePersonSkillExperiences**](PersonAPI.md#UpdatePersonSkillExperiences) | **Put** /persons/{personId}/experiences | Edit an Skill experience to a Person
[**UupdatePersonLanguage**](PersonAPI.md#UupdatePersonLanguage) | **Put** /persons/{personId}/languages/{languageId} | Update a language of a person



## AddPersonCertification

> PersonDetails AddPersonCertification(ctx, personId, certificationId).Body(body).Execute()

Add Certification to a Person

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
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    certificationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    body := time.Now() // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.AddPersonCertification(context.Background(), personId, certificationId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.AddPersonCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPersonCertification`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.AddPersonCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**certificationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPersonCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** |  | 

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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.AddPersonInterest(context.Background(), personId, skillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.AddPersonInterest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPersonInterest`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.AddPersonInterest`: %v\n", resp)
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


## AddPersonLanguage

> PersonDetails AddPersonLanguage(ctx, personId, languageId).Level(level).Execute()

Assign a language to the person

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
    languageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    level := *openapiclient.NewLevel() // Level | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.AddPersonLanguage(context.Background(), personId, languageId).Level(level).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.AddPersonLanguage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPersonLanguage`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.AddPersonLanguage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**languageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPersonLanguageRequest struct via the builder pattern


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


## AddPersonOffice

> PersonDetails AddPersonOffice(ctx, personId, officeId).Execute()

Assign a person to an office

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
    officeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.AddPersonOffice(context.Background(), personId, officeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.AddPersonOffice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPersonOffice`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.AddPersonOffice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**officeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPersonOfficeRequest struct via the builder pattern


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


## AddPersonProfile

> PersonDetails AddPersonProfile(ctx, personId, profileId).Execute()

Add a Profile to a Person

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
    profileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.AddPersonProfile(context.Background(), personId, profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.AddPersonProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPersonProfile`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.AddPersonProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**profileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPersonProfileRequest struct via the builder pattern


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


## AddPersonProject

> ProjectParticipation AddPersonProject(ctx, personId, projectId).ProjectParticipationUpdate(projectParticipationUpdate).Execute()

Add Project to a Person

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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    projectParticipationUpdate := *openapiclient.NewProjectParticipationUpdate() // ProjectParticipationUpdate | List of Skills with level and timeframe

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.AddPersonProject(context.Background(), personId, projectId).ProjectParticipationUpdate(projectParticipationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.AddPersonProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPersonProject`: ProjectParticipation
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.AddPersonProject`: %v\n", resp)
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

[**ProjectParticipation**](ProjectParticipation.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddPersonProjectSkill

> ProjectParticipation AddPersonProjectSkill(ctx, personId, projectId, skillId).Level(level).Execute()

Add Skill to a Project participation

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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    level := *openapiclient.NewLevel() // Level | The Skill Level

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.AddPersonProjectSkill(context.Background(), personId, projectId, skillId).Level(level).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.AddPersonProjectSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPersonProjectSkill`: ProjectParticipation
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.AddPersonProjectSkill`: %v\n", resp)
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

[**ProjectParticipation**](ProjectParticipation.md)

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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    level := *openapiclient.NewLevel() // Level | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.AddPersonSkillExperience(context.Background(), personId, skillId).Level(level).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.AddPersonSkillExperience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPersonSkillExperience`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.AddPersonSkillExperience`: %v\n", resp)
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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillLevelUpdate := []openapiclient.SkillLevelUpdate{*openapiclient.NewSkillLevelUpdate()} // []SkillLevelUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.AddPersonSkillExperiences(context.Background(), personId).SkillLevelUpdate(skillLevelUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.AddPersonSkillExperiences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPersonSkillExperiences`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.AddPersonSkillExperiences`: %v\n", resp)
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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    confirmingPersonId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.ConfirmSkill(context.Background(), personId, projectId, skillId, confirmingPersonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.ConfirmSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmSkill`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.ConfirmSkill`: %v\n", resp)
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


## CreateAvailability

> AvailabilityDetail CreateAvailability(ctx, personId).Availability(availability).Execute()

Create a availability for a person

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
    availability := *openapiclient.NewAvailability(float32(38.5), float32(22.5), "Id_example", "Name_example") // Availability | The availability

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.CreateAvailability(context.Background(), personId).Availability(availability).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.CreateAvailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAvailability`: AvailabilityDetail
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.CreateAvailability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **availability** | [**Availability**](Availability.md) | The availability | 

### Return type

[**AvailabilityDetail**](AvailabilityDetail.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    person := *openapiclient.NewPerson(false, "Id_example", "Name_example") // Person | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.CreatePerson(context.Background()).Person(person).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.CreatePerson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePerson`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.CreatePerson`: %v\n", resp)
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

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAvailability

> Status DeleteAvailability(ctx, personId, availabilityId).Execute()

Delete a person availability

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
    availabilityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.DeleteAvailability(context.Background(), personId, availabilityId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.DeleteAvailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAvailability`: Status
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.DeleteAvailability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**availabilityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAvailabilityRequest struct via the builder pattern


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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    confirmingPersonId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.DeleteConfirmation(context.Background(), personId, projectId, skillId, confirmingPersonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.DeleteConfirmation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteConfirmation`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.DeleteConfirmation`: %v\n", resp)
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


## DeletePerson

> Status DeletePerson(ctx, personId).Execute()

Delete an existing Person

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.DeletePerson(context.Background(), personId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.DeletePerson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePerson`: Status
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.DeletePerson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePersonRequest struct via the builder pattern


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


## DeletePersonCertification

> PersonDetails DeletePersonCertification(ctx, personId, certificationId).Execute()

Remove an Interest to a Person

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
    certificationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.DeletePersonCertification(context.Background(), personId, certificationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.DeletePersonCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePersonCertification`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.DeletePersonCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**certificationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePersonCertificationRequest struct via the builder pattern


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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.DeletePersonInterest(context.Background(), personId, skillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.DeletePersonInterest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePersonInterest`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.DeletePersonInterest`: %v\n", resp)
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


## DeletePersonOffice

> PersonDetails DeletePersonOffice(ctx, personId, officeId).Execute()

Delete the office from a Person

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
    officeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.DeletePersonOffice(context.Background(), personId, officeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.DeletePersonOffice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePersonOffice`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.DeletePersonOffice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**officeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePersonOfficeRequest struct via the builder pattern


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


## DeletePersonProfile

> PersonDetails DeletePersonProfile(ctx, personId, profileId).Execute()

Remove a Profile from a Person

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
    profileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.DeletePersonProfile(context.Background(), personId, profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.DeletePersonProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePersonProfile`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.DeletePersonProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**profileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePersonProfileRequest struct via the builder pattern


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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.DeletePersonProject(context.Background(), personId, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.DeletePersonProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePersonProject`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.DeletePersonProject`: %v\n", resp)
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

> ProjectParticipation DeletePersonProjectSkill(ctx, personId, projectId, skillId).Execute()

Remove a Skill from a Project participation

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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.DeletePersonProjectSkill(context.Background(), personId, projectId, skillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.DeletePersonProjectSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePersonProjectSkill`: ProjectParticipation
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.DeletePersonProjectSkill`: %v\n", resp)
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

[**ProjectParticipation**](ProjectParticipation.md)

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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.DeletePersonSkillExperience(context.Background(), personId, skillId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.DeletePersonSkillExperience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePersonSkillExperience`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.DeletePersonSkillExperience`: %v\n", resp)
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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    requestBody := []string{"Property_example"} // []string | A list of skillIds

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.DeletePersonSkillExperiences(context.Background(), personId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.DeletePersonSkillExperiences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePersonSkillExperiences`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.DeletePersonSkillExperiences`: %v\n", resp)
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


## GeneratePersonProfile

> *os.File GeneratePersonProfile(ctx, personId).ProjectParticipationIds(projectParticipationIds).SkillIds(skillIds).CertificationIds(certificationIds).ProfileIds(profileIds).Template(template).Execute()

Generate a PDF profile from a Person

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
    projectParticipationIds := []string{"Inner_example"} // []string |  (optional)
    skillIds := []string{"Inner_example"} // []string |  (optional)
    certificationIds := []string{"Inner_example"} // []string |  (optional)
    profileIds := []string{"Inner_example"} // []string |  (optional)
    template := "template_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.GeneratePersonProfile(context.Background(), personId).ProjectParticipationIds(projectParticipationIds).SkillIds(skillIds).CertificationIds(certificationIds).ProfileIds(profileIds).Template(template).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.GeneratePersonProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneratePersonProfile`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.GeneratePersonProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGeneratePersonProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectParticipationIds** | **[]string** |  | 
 **skillIds** | **[]string** |  | 
 **certificationIds** | **[]string** |  | 
 **profileIds** | **[]string** |  | 
 **template** | **string** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllBusinessDepartments

> []string GetAllBusinessDepartments(ctx).Skip(skip).Limit(limit).Execute()

Get all unique business departments

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
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.GetAllBusinessDepartments(context.Background()).Skip(skip).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.GetAllBusinessDepartments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllBusinessDepartments`: []string
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.GetAllBusinessDepartments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllBusinessDepartmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

**[]string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailabilities

> PagedAvailabilities GetAvailabilities(ctx, personId).Execute()

Get a list of all activities for a person

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.GetAvailabilities(context.Background(), personId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.GetAvailabilities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvailabilities`: PagedAvailabilities
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.GetAvailabilities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PagedAvailabilities**](PagedAvailabilities.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.GetPerson(context.Background(), personId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.GetPerson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPerson`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.GetPerson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonRequest struct via the builder pattern


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


## ReadPersonPicture

> *os.File ReadPersonPicture(ctx, personId).Execute()

Read person image

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.ReadPersonPicture(context.Background(), personId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.ReadPersonPicture``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadPersonPicture`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.ReadPersonPicture`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadPersonPictureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/jpeg, application/json

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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.ReadPersonProject(context.Background(), personId, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.ReadPersonProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadPersonProject`: ProjectParticipation
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.ReadPersonProject`: %v\n", resp)
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


## RemovePersonLanguage

> PersonDetails RemovePersonLanguage(ctx, personId, languageId).Execute()

Remove a language from a person

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
    languageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.RemovePersonLanguage(context.Background(), personId, languageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.RemovePersonLanguage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemovePersonLanguage`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.RemovePersonLanguage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**languageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePersonLanguageRequest struct via the builder pattern


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


## SearchPersons

> PagedPersons SearchPersons(ctx).PersonSearch(personSearch).Skip(skip).Limit(limit).Execute()

Complex search over person entities

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
    personSearch := *openapiclient.NewPersonSearch() // PersonSearch | 
    skip := int32(0) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.SearchPersons(context.Background()).PersonSearch(personSearch).Skip(skip).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.SearchPersons``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchPersons`: PagedPersons
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.SearchPersons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPersonsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **personSearch** | [**PersonSearch**](PersonSearch.md) |  | 
 **skip** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**PagedPersons**](PagedPersons.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAvailability

> AvailabilityDetail UpdateAvailability(ctx, personId, availabilityId).Availability(availability).Execute()

Update a person availability

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
    availabilityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    availability := *openapiclient.NewAvailability(float32(38.5), float32(22.5), "Id_example", "Name_example") // Availability | The availability

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.UpdateAvailability(context.Background(), personId, availabilityId).Availability(availability).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.UpdateAvailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAvailability`: AvailabilityDetail
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.UpdateAvailability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**availabilityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **availability** | [**Availability**](Availability.md) | The availability | 

### Return type

[**AvailabilityDetail**](AvailabilityDetail.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    person := *openapiclient.NewPerson(false, "Id_example", "Name_example") // Person | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.UpdatePerson(context.Background(), personId).Person(person).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.UpdatePerson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePerson`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.UpdatePerson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePersonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **person** | [**Person**](Person.md) |  | 

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


## UpdatePersonCertification

> PersonDetails UpdatePersonCertification(ctx, personId, certificationId).Body(body).Execute()

Update a Certification of a Person

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
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    certificationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    body := time.Now() // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.UpdatePersonCertification(context.Background(), personId, certificationId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.UpdatePersonCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePersonCertification`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.UpdatePersonCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**certificationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePersonCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** |  | 

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


## UpdatePersonPicture

> Status UpdatePersonPicture(ctx, personId).Body(body).Execute()

Update person image

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
    body := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.UpdatePersonPicture(context.Background(), personId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.UpdatePersonPicture``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePersonPicture`: Status
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.UpdatePersonPicture`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePersonPictureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** |  | 

### Return type

[**Status**](Status.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePersonProject

> ProjectParticipation UpdatePersonProject(ctx, personId, projectId).ProjectParticipationUpdate(projectParticipationUpdate).Execute()

Update a Project of a Person

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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    projectParticipationUpdate := *openapiclient.NewProjectParticipationUpdate() // ProjectParticipationUpdate | A time frame and a list of Skills

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.UpdatePersonProject(context.Background(), personId, projectId).ProjectParticipationUpdate(projectParticipationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.UpdatePersonProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePersonProject`: ProjectParticipation
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.UpdatePersonProject`: %v\n", resp)
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

[**ProjectParticipation**](ProjectParticipation.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePersonProjectSkill

> ProjectParticipation UpdatePersonProjectSkill(ctx, personId, projectId, skillId).Level(level).Execute()

Update the level of a Skill in a Project participation

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
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    level := *openapiclient.NewLevel() // Level | The Skill Level

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.UpdatePersonProjectSkill(context.Background(), personId, projectId, skillId).Level(level).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.UpdatePersonProjectSkill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePersonProjectSkill`: ProjectParticipation
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.UpdatePersonProjectSkill`: %v\n", resp)
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

[**ProjectParticipation**](ProjectParticipation.md)

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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    level := *openapiclient.NewLevel() // Level | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.UpdatePersonSkillExperience(context.Background(), personId, skillId).Level(level).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.UpdatePersonSkillExperience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePersonSkillExperience`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.UpdatePersonSkillExperience`: %v\n", resp)
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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    personId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    skillLevelUpdate := []openapiclient.SkillLevelUpdate{*openapiclient.NewSkillLevelUpdate()} // []SkillLevelUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.UpdatePersonSkillExperiences(context.Background(), personId).SkillLevelUpdate(skillLevelUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.UpdatePersonSkillExperiences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePersonSkillExperiences`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.UpdatePersonSkillExperiences`: %v\n", resp)
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


## UupdatePersonLanguage

> PersonDetails UupdatePersonLanguage(ctx, personId, languageId).Level(level).Execute()

Update a language of a person

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
    languageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    level := *openapiclient.NewLevel() // Level | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.UupdatePersonLanguage(context.Background(), personId, languageId).Level(level).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.UupdatePersonLanguage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UupdatePersonLanguage`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.UupdatePersonLanguage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**languageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUupdatePersonLanguageRequest struct via the builder pattern


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

