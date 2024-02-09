# \PersonAPI

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPersonCertification**](PersonAPI.md#AddPersonCertification) | **Post** /persons/{personId}/certifications/{certificationId} | Add Certification to a Person
[**AddPersonInterest**](PersonAPI.md#AddPersonInterest) | **Post** /persons/{personId}/interests/skills/{skillId} | Add an Interest to a Person
[**AddPersonLanguage**](PersonAPI.md#AddPersonLanguage) | **Post** /persons/{personId}/languages/{languageId} | Assign a language to the person
[**AddPersonOffice**](PersonAPI.md#AddPersonOffice) | **Post** /persons/{personId}/offices/{officeId} | Assign a person to an office
[**AddPersonProfile**](PersonAPI.md#AddPersonProfile) | **Post** /persons/{personId}/profiles/{profileId} | Add a Profile to a Person
[**AddPersonSkillExperience**](PersonAPI.md#AddPersonSkillExperience) | **Post** /persons/{personId}/experiences/skills/{skillId} | Add an Skill experience to a Person
[**AddPersonSkillExperiences**](PersonAPI.md#AddPersonSkillExperiences) | **Post** /persons/{personId}/experiences | Add an Skill experience to a Person (bulk)
[**AddProjectParticipation**](PersonAPI.md#AddProjectParticipation) | **Post** /project-participations | Add Project to a Person
[**AddSkillConfirmation**](PersonAPI.md#AddSkillConfirmation) | **Post** /project-participations/{projectParticipationId}/skills/{skillId}/confirmation/{confirmingPersonId} | Confirm Skill
[**CreateAvailability**](PersonAPI.md#CreateAvailability) | **Post** /persons/{personId}/availabilities | Create a availability for a person
[**CreatePerson**](PersonAPI.md#CreatePerson) | **Post** /persons | Create a new Person
[**DeleteAvailability**](PersonAPI.md#DeleteAvailability) | **Delete** /persons/{personId}/availabilities/{availabilityId} | Delete a person availability
[**DeletePerson**](PersonAPI.md#DeletePerson) | **Delete** /persons/{personId} | Delete an existing Person
[**DeletePersonCertification**](PersonAPI.md#DeletePersonCertification) | **Delete** /persons/{personId}/certifications/{certificationId} | Remove an Interest to a Person
[**DeletePersonInterest**](PersonAPI.md#DeletePersonInterest) | **Delete** /persons/{personId}/interests/skills/{skillId} | Remove an Interest to a Person
[**DeletePersonOffice**](PersonAPI.md#DeletePersonOffice) | **Delete** /persons/{personId}/offices/{officeId} | Delete the office from a Person
[**DeletePersonPicture**](PersonAPI.md#DeletePersonPicture) | **Delete** /persons/{personId}/picture | Delete person image
[**DeletePersonProfile**](PersonAPI.md#DeletePersonProfile) | **Delete** /persons/{personId}/profiles/{profileId} | Remove a Profile from a Person
[**DeletePersonSkillExperience**](PersonAPI.md#DeletePersonSkillExperience) | **Delete** /persons/{personId}/experiences/skills/{skillId} | Remove an Skill Experience to a Person
[**DeletePersonSkillExperiences**](PersonAPI.md#DeletePersonSkillExperiences) | **Delete** /persons/{personId}/experiences | Remove an Skill Experience to a Person
[**DeleteProjectParticipation**](PersonAPI.md#DeleteProjectParticipation) | **Delete** /project-participations/{projectParticipationId} | Remove an Project from a Person
[**GeneratePersonProfile**](PersonAPI.md#GeneratePersonProfile) | **Get** /persons/{personId}/pdf-profile | Generate a PDF profile from a Person
[**GetAllBusinessDepartments**](PersonAPI.md#GetAllBusinessDepartments) | **Get** /persons/departments | Get all unique business departments
[**GetAvailabilities**](PersonAPI.md#GetAvailabilities) | **Get** /persons/{personId}/availabilities | Get a list of all activities for a person
[**GetPerson**](PersonAPI.md#GetPerson) | **Get** /persons/{personId} | Get basic info about a person
[**ReadPersonPicture**](PersonAPI.md#ReadPersonPicture) | **Get** /persons/{personId}/picture | Read person image
[**ReadPersonProjectParticipation**](PersonAPI.md#ReadPersonProjectParticipation) | **Get** /persons/{personId}/project-participation | Get a Project Participation of a Person
[**ReadProjectParticipation**](PersonAPI.md#ReadProjectParticipation) | **Get** /project-participations/{projectParticipationId} | Get a project participation
[**RemovePersonLanguage**](PersonAPI.md#RemovePersonLanguage) | **Delete** /persons/{personId}/languages/{languageId} | Remove a language from a person
[**RemoveSkillConfirmation**](PersonAPI.md#RemoveSkillConfirmation) | **Delete** /project-participations/{projectParticipationId}/skills/{skillId}/confirmation/{confirmingPersonId} | Remove a confirmation
[**SearchPersons**](PersonAPI.md#SearchPersons) | **Post** /persons/search | Complex search over person entities
[**UpdateAvailability**](PersonAPI.md#UpdateAvailability) | **Put** /persons/{personId}/availabilities/{availabilityId} | Update a person availability
[**UpdatePerson**](PersonAPI.md#UpdatePerson) | **Put** /persons/{personId} | Update an existing Person
[**UpdatePersonCertification**](PersonAPI.md#UpdatePersonCertification) | **Put** /persons/{personId}/certifications/{certificationId} | Update a Certification of a Person
[**UpdatePersonLanguage**](PersonAPI.md#UpdatePersonLanguage) | **Put** /persons/{personId}/languages/{languageId} | Update a language of a person
[**UpdatePersonPicture**](PersonAPI.md#UpdatePersonPicture) | **Put** /persons/{personId}/picture | Update person image
[**UpdatePersonSkillExperience**](PersonAPI.md#UpdatePersonSkillExperience) | **Put** /persons/{personId}/experiences/skills/{skillId} | Edit an Skill experience to a Person
[**UpdatePersonSkillExperiences**](PersonAPI.md#UpdatePersonSkillExperiences) | **Put** /persons/{personId}/experiences | Edit an Skill experience to a Person
[**UpdateProjectParticipation**](PersonAPI.md#UpdateProjectParticipation) | **Put** /project-participations/{projectParticipationId} | Update a Project of a Person



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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    projectParticipationCreate := *openapiclient.NewProjectParticipationCreate("Id_example", "ProjectId_example", "PersonId_example") // ProjectParticipationCreate | List of Skills with level and timeframe

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.AddProjectParticipation(context.Background()).ProjectParticipationCreate(projectParticipationCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.AddProjectParticipation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddProjectParticipation`: ProjectParticipationDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.AddProjectParticipation`: %v\n", resp)
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

[bearerAuth](../README.md#bearerAuth)

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
    resp, r, err := apiClient.PersonAPI.AddSkillConfirmation(context.Background(), projectParticipationId, skillId, confirmingPersonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.AddSkillConfirmation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSkillConfirmation`: ProjectParticipationDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.AddSkillConfirmation`: %v\n", resp)
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


## DeletePersonPicture

> Status DeletePersonPicture(ctx, personId).Execute()

Delete person image

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
    resp, r, err := apiClient.PersonAPI.DeletePersonPicture(context.Background(), personId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.DeletePersonPicture``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePersonPicture`: Status
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.DeletePersonPicture`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePersonPictureRequest struct via the builder pattern


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
    resp, r, err := apiClient.PersonAPI.DeleteProjectParticipation(context.Background(), projectParticipationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.DeleteProjectParticipation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteProjectParticipation`: Status
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.DeleteProjectParticipation`: %v\n", resp)
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

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneratePersonProfile

> *os.File GeneratePersonProfile(ctx, personId).ProjectIds(projectIds).SkillIds(skillIds).CertificationIds(certificationIds).ProfileIds(profileIds).Seniority(seniority).Template(template).Execute()

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
    projectIds := []string{"Inner_example"} // []string |  (optional)
    skillIds := []string{"Inner_example"} // []string |  (optional)
    certificationIds := []string{"Inner_example"} // []string |  (optional)
    profileIds := []string{"Inner_example"} // []string |  (optional)
    seniority := openapiclient.Seniority("UNKNOWN") // Seniority |  (optional)
    template := "template_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.GeneratePersonProfile(context.Background(), personId).ProjectIds(projectIds).SkillIds(skillIds).CertificationIds(certificationIds).ProfileIds(profileIds).Seniority(seniority).Template(template).Execute()
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

 **projectIds** | **[]string** |  | 
 **skillIds** | **[]string** |  | 
 **certificationIds** | **[]string** |  | 
 **profileIds** | **[]string** |  | 
 **seniority** | [**Seniority**](Seniority.md) |  | 
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
    limit := int32(20) // int32 |  (optional) (default to 20)

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


## ReadPersonProjectParticipation

> []PersonProjectParticipationDetails ReadPersonProjectParticipation(ctx, personId).ProjectIds(projectIds).Execute()

Get a Project Participation of a Person

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
    projectIds := []string{"Inner_example"} // []string | Filter by project ids, if not set, all projects are returned (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.ReadPersonProjectParticipation(context.Background(), personId).ProjectIds(projectIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.ReadPersonProjectParticipation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadPersonProjectParticipation`: []PersonProjectParticipationDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.ReadPersonProjectParticipation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadPersonProjectParticipationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectIds** | **[]string** | Filter by project ids, if not set, all projects are returned | 

### Return type

[**[]PersonProjectParticipationDetails**](PersonProjectParticipationDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

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
    resp, r, err := apiClient.PersonAPI.ReadProjectParticipation(context.Background(), projectParticipationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.ReadProjectParticipation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadProjectParticipation`: ProjectParticipationDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.ReadProjectParticipation`: %v\n", resp)
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
    resp, r, err := apiClient.PersonAPI.RemoveSkillConfirmation(context.Background(), projectParticipationId, skillId, confirmingPersonId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.RemoveSkillConfirmation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveSkillConfirmation`: ProjectParticipationDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.RemoveSkillConfirmation`: %v\n", resp)
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
    limit := int32(20) // int32 |  (optional) (default to 20)

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


## UpdatePersonLanguage

> PersonDetails UpdatePersonLanguage(ctx, personId, languageId).Level(level).Execute()

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
    resp, r, err := apiClient.PersonAPI.UpdatePersonLanguage(context.Background(), personId, languageId).Level(level).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.UpdatePersonLanguage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePersonLanguage`: PersonDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.UpdatePersonLanguage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **string** |  | 
**languageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePersonLanguageRequest struct via the builder pattern


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
    openapiclient "github.com/prodyna-yasm/yasm-api-go"
)

func main() {
    projectParticipationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    projectParticipationUpdate := *openapiclient.NewProjectParticipationUpdate() // ProjectParticipationUpdate | A time frame and a list of Skills

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonAPI.UpdateProjectParticipation(context.Background(), projectParticipationId).ProjectParticipationUpdate(projectParticipationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.UpdateProjectParticipation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProjectParticipation`: ProjectParticipationDetails
    fmt.Fprintf(os.Stdout, "Response from `PersonAPI.UpdateProjectParticipation`: %v\n", resp)
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

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

