# Go API client for client

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.1.0
- Package version: 0.0.1
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import client "github.com/prodyna/yasm-api-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), client.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), client.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), client.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), client.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to */api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AvailabilityApi* | [**CreateAvailability**](docs/AvailabilityApi.md#createavailability) | **Post** /persons/{personId}/availabilities | Create a availability for a person
*AvailabilityApi* | [**DeleteAvailability**](docs/AvailabilityApi.md#deleteavailability) | **Delete** /persons/{personId}/availabilities/{availabilityId} | Delete a person availability
*AvailabilityApi* | [**GetAvailabilities**](docs/AvailabilityApi.md#getavailabilities) | **Get** /persons/{personId}/availabilities | Get a list of all activities for a person
*AvailabilityApi* | [**UpdateAvailability**](docs/AvailabilityApi.md#updateavailability) | **Put** /persons/{personId}/availabilities/{availabilityId} | Update a person availability
*CertificationApi* | [**AddPersonCertification**](docs/CertificationApi.md#addpersoncertification) | **Post** /persons/{personId}/certifications/{certificationId} | Add Certification to a Person
*CertificationApi* | [**AddSkillToCertification**](docs/CertificationApi.md#addskilltocertification) | **Post** /certifications/{certificationId}/skills/{skillId} | 
*CertificationApi* | [**CreateCertification**](docs/CertificationApi.md#createcertification) | **Post** /organizations/{organizationId}/certifications | Create a Certification in an Organization
*CertificationApi* | [**DeleteCertification**](docs/CertificationApi.md#deletecertification) | **Delete** /certifications/{certificationId} | Delete a Certification
*CertificationApi* | [**DeletePersonCertification**](docs/CertificationApi.md#deletepersoncertification) | **Delete** /persons/{personId}/certifications/{certificationId} | Remove an Interest to a Person
*CertificationApi* | [**DeleteSkillFromCertification**](docs/CertificationApi.md#deleteskillfromcertification) | **Delete** /certifications/{certificationId}/skills/{skillId} | 
*CertificationApi* | [**GetCertification**](docs/CertificationApi.md#getcertification) | **Get** /certifications/{certificationId} | Get details about a Certification
*CertificationApi* | [**GetCertifications**](docs/CertificationApi.md#getcertifications) | **Get** /certifications | Get a list of all Certifations indepdenant of the Organization
*CertificationApi* | [**GetCertificationsForOrganization**](docs/CertificationApi.md#getcertificationsfororganization) | **Get** /organizations/{organizationId}/certifications | Get a list of all certifications for a organization
*CertificationApi* | [**MoveCertification**](docs/CertificationApi.md#movecertification) | **Put** /organizations/{organizationId}/certificates/{certificateId} | Move a Certification to an Organization
*CertificationApi* | [**UpdateCertification**](docs/CertificationApi.md#updatecertification) | **Put** /certifications/{certificationId} | Update a Certification
*CertificationApi* | [**UpdatePersonCertification**](docs/CertificationApi.md#updatepersoncertification) | **Put** /persons/{personId}/certifications/{certificationId} | Update a Certification of a Person
*CertificationApi* | [**UpdateSkillInCertification**](docs/CertificationApi.md#updateskillincertification) | **Put** /certifications/{certificationId}/skills/{skillId} | 
*CountryApi* | [**AddLanguageToCountry**](docs/CountryApi.md#addlanguagetocountry) | **Post** /countries/{countryId}/languages/{languageId} | Assign a language to a country
*CountryApi* | [**CreateCountry**](docs/CountryApi.md#createcountry) | **Post** /countries | Create a new Country
*CountryApi* | [**DeleteCountry**](docs/CountryApi.md#deletecountry) | **Delete** /countries/{countryId} | Delete a Country
*CountryApi* | [**GetCountries**](docs/CountryApi.md#getcountries) | **Get** /countries | Get all Countries
*CountryApi* | [**GetCountry**](docs/CountryApi.md#getcountry) | **Get** /countries/{countryId} | Get details about a Country
*CountryApi* | [**RemoveLanguageFromCountry**](docs/CountryApi.md#removelanguagefromcountry) | **Delete** /countries/{countryId}/languages/{languageId} | Assign a language to a country
*CountryApi* | [**UpdateCountry**](docs/CountryApi.md#updatecountry) | **Put** /countries/{countryId} | Update a Country
*IndustryApi* | [**AttachOrganizationToIndustry**](docs/IndustryApi.md#attachorganizationtoindustry) | **Post** /organizations/{organizationId}/industries/{industryId} | Add an Organization to an Industry
*IndustryApi* | [**CreateIndustry**](docs/IndustryApi.md#createindustry) | **Post** /industries | Create an Industry
*IndustryApi* | [**DeleteIndustry**](docs/IndustryApi.md#deleteindustry) | **Delete** /industries/{industryId} | Delete an Industry
*IndustryApi* | [**DetachOrganizationFromIndustry**](docs/IndustryApi.md#detachorganizationfromindustry) | **Delete** /organizations/{organizationId}/industries/{industryId} | Remove an Organization to an Industry
*IndustryApi* | [**GetIndustries**](docs/IndustryApi.md#getindustries) | **Get** /industries | Get all Industries
*IndustryApi* | [**GetIndustry**](docs/IndustryApi.md#getindustry) | **Get** /industries/{industryId} | Get details about an Industry
*IndustryApi* | [**UpdateIndustry**](docs/IndustryApi.md#updateindustry) | **Put** /industries/{industryId} | Update an Industry
*LanguageApi* | [**AddLanguageToCountry**](docs/LanguageApi.md#addlanguagetocountry) | **Post** /countries/{countryId}/languages/{languageId} | Assign a language to a country
*LanguageApi* | [**AddPersonLanguage**](docs/LanguageApi.md#addpersonlanguage) | **Post** /persons/{personId}/languages/{languageId} | Assign a language to the person
*LanguageApi* | [**CreateLanguage**](docs/LanguageApi.md#createlanguage) | **Post** /languages | Create a new language
*LanguageApi* | [**DeleteLanguage**](docs/LanguageApi.md#deletelanguage) | **Delete** /languages/{languageId} | Delete a language
*LanguageApi* | [**GetLanguage**](docs/LanguageApi.md#getlanguage) | **Get** /languages/{languageId} | Get details about a language
*LanguageApi* | [**GetLanguages**](docs/LanguageApi.md#getlanguages) | **Get** /languages | Get a list of Languages
*LanguageApi* | [**RemoveLanguageFromCountry**](docs/LanguageApi.md#removelanguagefromcountry) | **Delete** /countries/{countryId}/languages/{languageId} | Assign a language to a country
*LanguageApi* | [**RemovePersonLanguage**](docs/LanguageApi.md#removepersonlanguage) | **Delete** /persons/{personId}/languages/{languageId} | Remove a language from a person
*LanguageApi* | [**UupdatePersonLanguage**](docs/LanguageApi.md#uupdatepersonlanguage) | **Put** /persons/{personId}/languages/{languageId} | Update a language of a person
*OfficeApi* | [**AddPersonOffice**](docs/OfficeApi.md#addpersonoffice) | **Post** /persons/{personId}/offices/{officeId} | Assing a person to an office
*OfficeApi* | [**CreateOffice**](docs/OfficeApi.md#createoffice) | **Post** /organizations/{organizationId}/offices | Create an Office in an Organization
*OfficeApi* | [**DeleteOffice**](docs/OfficeApi.md#deleteoffice) | **Delete** /organizations/{organizationId}/offices/{officeId} | Delte an Office from an Organization
*OfficeApi* | [**DeletePersonOffice**](docs/OfficeApi.md#deletepersonoffice) | **Delete** /persons/{personId}/offices/{officeId} | Delete the office from a Person
*OfficeApi* | [**GetOffice**](docs/OfficeApi.md#getoffice) | **Get** /organizations/{organizationId}/offices/{officeId} | Get an Office for an Organiaztion
*OfficeApi* | [**UpdateOffice**](docs/OfficeApi.md#updateoffice) | **Put** /organizations/{organizationId}/offices/{officeId} | Update an Office for an Organization
*OrganizationApi* | [**AttachOrganizationToIndustry**](docs/OrganizationApi.md#attachorganizationtoindustry) | **Post** /organizations/{organizationId}/industries/{industryId} | Add an Organization to an Industry
*OrganizationApi* | [**CreateCertification**](docs/OrganizationApi.md#createcertification) | **Post** /organizations/{organizationId}/certifications | Create a Certification in an Organization
*OrganizationApi* | [**CreateOffice**](docs/OrganizationApi.md#createoffice) | **Post** /organizations/{organizationId}/offices | Create an Office in an Organization
*OrganizationApi* | [**CreateOrganization**](docs/OrganizationApi.md#createorganization) | **Post** /organizations | Create an Organization
*OrganizationApi* | [**CreateProject**](docs/OrganizationApi.md#createproject) | **Post** /organizations/{organizationId}/projects | Create a Project in an Organization
*OrganizationApi* | [**DeleteOffice**](docs/OrganizationApi.md#deleteoffice) | **Delete** /organizations/{organizationId}/offices/{officeId} | Delte an Office from an Organization
*OrganizationApi* | [**DeleteOrganization**](docs/OrganizationApi.md#deleteorganization) | **Delete** /organizations/{organizationId} | Delete an organization
*OrganizationApi* | [**DetachOrganizationFromIndustry**](docs/OrganizationApi.md#detachorganizationfromindustry) | **Delete** /organizations/{organizationId}/industries/{industryId} | Remove an Organization to an Industry
*OrganizationApi* | [**GetCertificationsForOrganization**](docs/OrganizationApi.md#getcertificationsfororganization) | **Get** /organizations/{organizationId}/certifications | Get a list of all certifications for a organization
*OrganizationApi* | [**GetOffice**](docs/OrganizationApi.md#getoffice) | **Get** /organizations/{organizationId}/offices/{officeId} | Get an Office for an Organiaztion
*OrganizationApi* | [**GetOrganization**](docs/OrganizationApi.md#getorganization) | **Get** /organizations/{organizationId} | Get details about an Organization
*OrganizationApi* | [**GetOrganizationProjects**](docs/OrganizationApi.md#getorganizationprojects) | **Get** /organizations/{organizationId}/projects | Get a list of all Projects for an Organization
*OrganizationApi* | [**GetOrganizations**](docs/OrganizationApi.md#getorganizations) | **Get** /organizations | Get a list of all Organizations
*OrganizationApi* | [**MergeOrganizations**](docs/OrganizationApi.md#mergeorganizations) | **Put** /organizations/{organizationId}/merge/{otherOrganizationId} | Merge two organizations
*OrganizationApi* | [**MoveCertification**](docs/OrganizationApi.md#movecertification) | **Put** /organizations/{organizationId}/certificates/{certificateId} | Move a Certification to an Organization
*OrganizationApi* | [**UpdateOffice**](docs/OrganizationApi.md#updateoffice) | **Put** /organizations/{organizationId}/offices/{officeId} | Update an Office for an Organization
*OrganizationApi* | [**UpdateOrganization**](docs/OrganizationApi.md#updateorganization) | **Put** /organizations/{organizationId} | Update an Organization
*PersonApi* | [**AddPersonCertification**](docs/PersonApi.md#addpersoncertification) | **Post** /persons/{personId}/certifications/{certificationId} | Add Certification to a Person
*PersonApi* | [**AddPersonInterest**](docs/PersonApi.md#addpersoninterest) | **Post** /persons/{personId}/interests/skills/{skillId} | Add an Interest to a Person
*PersonApi* | [**AddPersonLanguage**](docs/PersonApi.md#addpersonlanguage) | **Post** /persons/{personId}/languages/{languageId} | Assign a language to the person
*PersonApi* | [**AddPersonOffice**](docs/PersonApi.md#addpersonoffice) | **Post** /persons/{personId}/offices/{officeId} | Assing a person to an office
*PersonApi* | [**AddPersonProject**](docs/PersonApi.md#addpersonproject) | **Post** /persons/{personId}/projects/{projectId} | Add Project to a Person
*PersonApi* | [**AddPersonProjectSkill**](docs/PersonApi.md#addpersonprojectskill) | **Post** /persons/{personId}/projects/{projectId}/skills/{skillId} | Add Skill to a Project participation
*PersonApi* | [**AddPersonSkillExperience**](docs/PersonApi.md#addpersonskillexperience) | **Post** /persons/{personId}/experiences/skills/{skillId} | Add an Skill experience to a Person
*PersonApi* | [**AddPersonSkillExperiences**](docs/PersonApi.md#addpersonskillexperiences) | **Post** /persons/{personId}/experiences | Add an Skill experience to a Person (bulk)
*PersonApi* | [**ConfirmSkill**](docs/PersonApi.md#confirmskill) | **Post** /persons/{personId}/projects/{projectId}/skills/{skillId}/confirmation/{confirmingPersonId} | Confirm Skill
*PersonApi* | [**CreateAvailability**](docs/PersonApi.md#createavailability) | **Post** /persons/{personId}/availabilities | Create a availability for a person
*PersonApi* | [**CreatePerson**](docs/PersonApi.md#createperson) | **Post** /persons | Create a new Person
*PersonApi* | [**DeleteAvailability**](docs/PersonApi.md#deleteavailability) | **Delete** /persons/{personId}/availabilities/{availabilityId} | Delete a person availability
*PersonApi* | [**DeleteConfirmation**](docs/PersonApi.md#deleteconfirmation) | **Delete** /persons/{personId}/projects/{projectId}/skills/{skillId}/confirmation/{confirmingPersonId} | Remove a confirmation
*PersonApi* | [**DeletePerson**](docs/PersonApi.md#deleteperson) | **Delete** /persons/{personId} | Delete an existing Person
*PersonApi* | [**DeletePersonCertification**](docs/PersonApi.md#deletepersoncertification) | **Delete** /persons/{personId}/certifications/{certificationId} | Remove an Interest to a Person
*PersonApi* | [**DeletePersonInterest**](docs/PersonApi.md#deletepersoninterest) | **Delete** /persons/{personId}/interests/skills/{skillId} | Remove an Interest to a Person
*PersonApi* | [**DeletePersonOffice**](docs/PersonApi.md#deletepersonoffice) | **Delete** /persons/{personId}/offices/{officeId} | Delete the office from a Person
*PersonApi* | [**DeletePersonProject**](docs/PersonApi.md#deletepersonproject) | **Delete** /persons/{personId}/projects/{projectId} | Remove an Project from a Person
*PersonApi* | [**DeletePersonProjectSkill**](docs/PersonApi.md#deletepersonprojectskill) | **Delete** /persons/{personId}/projects/{projectId}/skills/{skillId} | Remove a Skill from a Project participation
*PersonApi* | [**DeletePersonSkillExperience**](docs/PersonApi.md#deletepersonskillexperience) | **Delete** /persons/{personId}/experiences/skills/{skillId} | Remove an Skill Experience to a Person
*PersonApi* | [**DeletePersonSkillExperiences**](docs/PersonApi.md#deletepersonskillexperiences) | **Delete** /persons/{personId}/experiences | Remove an Skill Experience to a Person
*PersonApi* | [**GeneratePersonProfile**](docs/PersonApi.md#generatepersonprofile) | **Get** /persons/{personid}/profile | Generate a PDF profile from a Person
*PersonApi* | [**GetAvailabilities**](docs/PersonApi.md#getavailabilities) | **Get** /persons/{personId}/availabilities | Get a list of all activities for a person
*PersonApi* | [**GetPerson**](docs/PersonApi.md#getperson) | **Get** /persons/{personId} | Get basic info about a person
*PersonApi* | [**ReadPersonProject**](docs/PersonApi.md#readpersonproject) | **Get** /persons/{personId}/projects/{projectId} | Get a Project Partifipation of a Person
*PersonApi* | [**RemovePersonLanguage**](docs/PersonApi.md#removepersonlanguage) | **Delete** /persons/{personId}/languages/{languageId} | Remove a language from a person
*PersonApi* | [**SearchPersons**](docs/PersonApi.md#searchpersons) | **Post** /persons/search | Complex search over person entities
*PersonApi* | [**UpdateAvailability**](docs/PersonApi.md#updateavailability) | **Put** /persons/{personId}/availabilities/{availabilityId} | Update a person availability
*PersonApi* | [**UpdatePerson**](docs/PersonApi.md#updateperson) | **Put** /persons/{personId} | Update an existing Person
*PersonApi* | [**UpdatePersonCertification**](docs/PersonApi.md#updatepersoncertification) | **Put** /persons/{personId}/certifications/{certificationId} | Update a Certification of a Person
*PersonApi* | [**UpdatePersonProject**](docs/PersonApi.md#updatepersonproject) | **Put** /persons/{personId}/projects/{projectId} | Update a Project of a Person
*PersonApi* | [**UpdatePersonProjectSkill**](docs/PersonApi.md#updatepersonprojectskill) | **Put** /persons/{personId}/projects/{projectId}/skills/{skillId} | Update the level of a Skill in a Project participation
*PersonApi* | [**UpdatePersonSkillExperience**](docs/PersonApi.md#updatepersonskillexperience) | **Put** /persons/{personId}/experiences/skills/{skillId} | Edit an Skill experience to a Person
*PersonApi* | [**UpdatePersonSkillExperiences**](docs/PersonApi.md#updatepersonskillexperiences) | **Put** /persons/{personId}/experiences | Edit an Skill experience to a Person
*PersonApi* | [**UupdatePersonLanguage**](docs/PersonApi.md#uupdatepersonlanguage) | **Put** /persons/{personId}/languages/{languageId} | Update a language of a person
*ProjectApi* | [**AddPersonProject**](docs/ProjectApi.md#addpersonproject) | **Post** /persons/{personId}/projects/{projectId} | Add Project to a Person
*ProjectApi* | [**AddPersonProjectSkill**](docs/ProjectApi.md#addpersonprojectskill) | **Post** /persons/{personId}/projects/{projectId}/skills/{skillId} | Add Skill to a Project participation
*ProjectApi* | [**ConfirmSkill**](docs/ProjectApi.md#confirmskill) | **Post** /persons/{personId}/projects/{projectId}/skills/{skillId}/confirmation/{confirmingPersonId} | Confirm Skill
*ProjectApi* | [**CreateProject**](docs/ProjectApi.md#createproject) | **Post** /organizations/{organizationId}/projects | Create a Project in an Organization
*ProjectApi* | [**DeleteConfirmation**](docs/ProjectApi.md#deleteconfirmation) | **Delete** /persons/{personId}/projects/{projectId}/skills/{skillId}/confirmation/{confirmingPersonId} | Remove a confirmation
*ProjectApi* | [**DeletePersonProject**](docs/ProjectApi.md#deletepersonproject) | **Delete** /persons/{personId}/projects/{projectId} | Remove an Project from a Person
*ProjectApi* | [**DeletePersonProjectSkill**](docs/ProjectApi.md#deletepersonprojectskill) | **Delete** /persons/{personId}/projects/{projectId}/skills/{skillId} | Remove a Skill from a Project participation
*ProjectApi* | [**DeleteProject**](docs/ProjectApi.md#deleteproject) | **Delete** /projects/{projectId} | Delete a project
*ProjectApi* | [**GetOrganizationProjects**](docs/ProjectApi.md#getorganizationprojects) | **Get** /organizations/{organizationId}/projects | Get a list of all Projects for an Organization
*ProjectApi* | [**GetProject**](docs/ProjectApi.md#getproject) | **Get** /projects/{projectId} | Get details about a Project
*ProjectApi* | [**MergeProjects**](docs/ProjectApi.md#mergeprojects) | **Put** /projects/{projectId}/merge/{otherProjectId} | Merge to projects
*ProjectApi* | [**ReadPersonProject**](docs/ProjectApi.md#readpersonproject) | **Get** /persons/{personId}/projects/{projectId} | Get a Project Partifipation of a Person
*ProjectApi* | [**SearchProjects**](docs/ProjectApi.md#searchprojects) | **Post** /projects/search | Complex search over project entities
*ProjectApi* | [**UpdatePersonProject**](docs/ProjectApi.md#updatepersonproject) | **Put** /persons/{personId}/projects/{projectId} | Update a Project of a Person
*ProjectApi* | [**UpdatePersonProjectSkill**](docs/ProjectApi.md#updatepersonprojectskill) | **Put** /persons/{personId}/projects/{projectId}/skills/{skillId} | Update the level of a Skill in a Project participation
*ProjectApi* | [**UpdateProject**](docs/ProjectApi.md#updateproject) | **Put** /projects/{projectId} | Update a Project
*SearchApi* | [**SearchAll**](docs/SearchApi.md#searchall) | **Get** /search/all/{text} | Fulltext search on all kinds of objects
*SkillApi* | [**AddPersonInterest**](docs/SkillApi.md#addpersoninterest) | **Post** /persons/{personId}/interests/skills/{skillId} | Add an Interest to a Person
*SkillApi* | [**AddPersonProjectSkill**](docs/SkillApi.md#addpersonprojectskill) | **Post** /persons/{personId}/projects/{projectId}/skills/{skillId} | Add Skill to a Project participation
*SkillApi* | [**AddPersonSkillExperience**](docs/SkillApi.md#addpersonskillexperience) | **Post** /persons/{personId}/experiences/skills/{skillId} | Add an Skill experience to a Person
*SkillApi* | [**AddPersonSkillExperiences**](docs/SkillApi.md#addpersonskillexperiences) | **Post** /persons/{personId}/experiences | Add an Skill experience to a Person (bulk)
*SkillApi* | [**AddSkillToCertification**](docs/SkillApi.md#addskilltocertification) | **Post** /certifications/{certificationId}/skills/{skillId} | 
*SkillApi* | [**AddSkillToParentSkill**](docs/SkillApi.md#addskilltoparentskill) | **Post** /skills/{skillId}/parents/{parentSkillId} | Attach a Skill to a parent Skill, returns the parent Skill
*SkillApi* | [**ConfirmSkill**](docs/SkillApi.md#confirmskill) | **Post** /persons/{personId}/projects/{projectId}/skills/{skillId}/confirmation/{confirmingPersonId} | Confirm Skill
*SkillApi* | [**CreateSkill**](docs/SkillApi.md#createskill) | **Post** /skills | Create a Skill
*SkillApi* | [**DeleteConfirmation**](docs/SkillApi.md#deleteconfirmation) | **Delete** /persons/{personId}/projects/{projectId}/skills/{skillId}/confirmation/{confirmingPersonId} | Remove a confirmation
*SkillApi* | [**DeletePersonInterest**](docs/SkillApi.md#deletepersoninterest) | **Delete** /persons/{personId}/interests/skills/{skillId} | Remove an Interest to a Person
*SkillApi* | [**DeletePersonProjectSkill**](docs/SkillApi.md#deletepersonprojectskill) | **Delete** /persons/{personId}/projects/{projectId}/skills/{skillId} | Remove a Skill from a Project participation
*SkillApi* | [**DeletePersonSkillExperience**](docs/SkillApi.md#deletepersonskillexperience) | **Delete** /persons/{personId}/experiences/skills/{skillId} | Remove an Skill Experience to a Person
*SkillApi* | [**DeletePersonSkillExperiences**](docs/SkillApi.md#deletepersonskillexperiences) | **Delete** /persons/{personId}/experiences | Remove an Skill Experience to a Person
*SkillApi* | [**DeleteSkill**](docs/SkillApi.md#deleteskill) | **Delete** /skills/{skillId} | Delete a Skill
*SkillApi* | [**DeleteSkillFromCertification**](docs/SkillApi.md#deleteskillfromcertification) | **Delete** /certifications/{certificationId}/skills/{skillId} | 
*SkillApi* | [**GetSkill**](docs/SkillApi.md#getskill) | **Get** /skills/{skillId} | Get details for a single skill
*SkillApi* | [**GetSkillParents**](docs/SkillApi.md#getskillparents) | **Get** /skills/{skillId}/parents | Get ghe list of parents for a skill
*SkillApi* | [**GetSkills**](docs/SkillApi.md#getskills) | **Get** /skills | Get a list of all skills, optionally only root, optionally only kinds
*SkillApi* | [**MergeSkills**](docs/SkillApi.md#mergeskills) | **Put** /skills/{skillId}/merge/{otherSkillId} | Merge two skills
*SkillApi* | [**RemoveSkillFromParentSkill**](docs/SkillApi.md#removeskillfromparentskill) | **Delete** /skills/{skillId}/parents/{parentSkillId} | Detaches a Skill from parent Skill, return the parent Skill
*SkillApi* | [**UpdatePersonProjectSkill**](docs/SkillApi.md#updatepersonprojectskill) | **Put** /persons/{personId}/projects/{projectId}/skills/{skillId} | Update the level of a Skill in a Project participation
*SkillApi* | [**UpdatePersonSkillExperience**](docs/SkillApi.md#updatepersonskillexperience) | **Put** /persons/{personId}/experiences/skills/{skillId} | Edit an Skill experience to a Person
*SkillApi* | [**UpdatePersonSkillExperiences**](docs/SkillApi.md#updatepersonskillexperiences) | **Put** /persons/{personId}/experiences | Edit an Skill experience to a Person
*SkillApi* | [**UpdateSkill**](docs/SkillApi.md#updateskill) | **Put** /skills/{skillId} | Update a Skill
*SkillApi* | [**UpdateSkillInCertification**](docs/SkillApi.md#updateskillincertification) | **Put** /certifications/{certificationId}/skills/{skillId} | 
*StatusApi* | [**GetPing**](docs/StatusApi.md#getping) | **Get** /ping | Server heartbeat operation
*StatusApi* | [**GetVersion**](docs/StatusApi.md#getversion) | **Get** /version | Information about the server


## Documentation For Models

 - [Availability](docs/Availability.md)
 - [AvailabilityAllOf](docs/AvailabilityAllOf.md)
 - [AvailabilityDetail](docs/AvailabilityDetail.md)
 - [AvailabilityFilter](docs/AvailabilityFilter.md)
 - [BasicDomainModel](docs/BasicDomainModel.md)
 - [Certification](docs/Certification.md)
 - [CertificationAllOf](docs/CertificationAllOf.md)
 - [CertificationDetails](docs/CertificationDetails.md)
 - [Country](docs/Country.md)
 - [CountryAllOf](docs/CountryAllOf.md)
 - [CountryDetails](docs/CountryDetails.md)
 - [Descriptable](docs/Descriptable.md)
 - [EntityFilter](docs/EntityFilter.md)
 - [Error](docs/Error.md)
 - [Experience](docs/Experience.md)
 - [ExperienceAllOf](docs/ExperienceAllOf.md)
 - [Geolocation](docs/Geolocation.md)
 - [Industry](docs/Industry.md)
 - [IndustryDetails](docs/IndustryDetails.md)
 - [Language](docs/Language.md)
 - [LanguageDetails](docs/LanguageDetails.md)
 - [LanguageLevel](docs/LanguageLevel.md)
 - [Level](docs/Level.md)
 - [Linkable](docs/Linkable.md)
 - [Locateable](docs/Locateable.md)
 - [MinMax](docs/MinMax.md)
 - [MinMaxPercent](docs/MinMaxPercent.md)
 - [NamedDomainModel](docs/NamedDomainModel.md)
 - [NamedDomainModelAllOf](docs/NamedDomainModelAllOf.md)
 - [Office](docs/Office.md)
 - [Organization](docs/Organization.md)
 - [OrganizationAllOf](docs/OrganizationAllOf.md)
 - [OrganizationDetails](docs/OrganizationDetails.md)
 - [Page](docs/Page.md)
 - [PagedAvailabilities](docs/PagedAvailabilities.md)
 - [PagedAvailabilitiesAllOf](docs/PagedAvailabilitiesAllOf.md)
 - [PagedCertifications](docs/PagedCertifications.md)
 - [PagedCertificationsAllOf](docs/PagedCertificationsAllOf.md)
 - [PagedCountries](docs/PagedCountries.md)
 - [PagedCountriesAllOf](docs/PagedCountriesAllOf.md)
 - [PagedIndustries](docs/PagedIndustries.md)
 - [PagedIndustriesAllOf](docs/PagedIndustriesAllOf.md)
 - [PagedLanguages](docs/PagedLanguages.md)
 - [PagedLanguagesAllOf](docs/PagedLanguagesAllOf.md)
 - [PagedOrganizations](docs/PagedOrganizations.md)
 - [PagedOrganizationsAllOf](docs/PagedOrganizationsAllOf.md)
 - [PagedPersons](docs/PagedPersons.md)
 - [PagedPersonsAllOf](docs/PagedPersonsAllOf.md)
 - [PagedProjects](docs/PagedProjects.md)
 - [PagedProjectsAllOf](docs/PagedProjectsAllOf.md)
 - [PagedSkills](docs/PagedSkills.md)
 - [PagedSkillsAllOf](docs/PagedSkillsAllOf.md)
 - [Person](docs/Person.md)
 - [PersonAllOf](docs/PersonAllOf.md)
 - [PersonCertificationFilter](docs/PersonCertificationFilter.md)
 - [PersonCertificationFilterAllOf](docs/PersonCertificationFilterAllOf.md)
 - [PersonDetails](docs/PersonDetails.md)
 - [PersonIndustryFilter](docs/PersonIndustryFilter.md)
 - [PersonIndustryFilterAllOf](docs/PersonIndustryFilterAllOf.md)
 - [PersonOrganizationFilter](docs/PersonOrganizationFilter.md)
 - [PersonOrganizationFilterAllOf](docs/PersonOrganizationFilterAllOf.md)
 - [PersonProjectFilter](docs/PersonProjectFilter.md)
 - [PersonProjectFilterAllOf](docs/PersonProjectFilterAllOf.md)
 - [PersonScoreDetail](docs/PersonScoreDetail.md)
 - [PersonSearch](docs/PersonSearch.md)
 - [PersonSearchCertificationsInner](docs/PersonSearchCertificationsInner.md)
 - [PersonSearchIndustriesInner](docs/PersonSearchIndustriesInner.md)
 - [PersonSearchOrganizationsInner](docs/PersonSearchOrganizationsInner.md)
 - [PersonSearchProjectsInner](docs/PersonSearchProjectsInner.md)
 - [PersonSearchSkillsInner](docs/PersonSearchSkillsInner.md)
 - [PersonSkillFilter](docs/PersonSkillFilter.md)
 - [PersonSkillFilterAllOf](docs/PersonSkillFilterAllOf.md)
 - [Project](docs/Project.md)
 - [ProjectAllOf](docs/ProjectAllOf.md)
 - [ProjectDetails](docs/ProjectDetails.md)
 - [ProjectParticipation](docs/ProjectParticipation.md)
 - [ProjectParticipationAllOf](docs/ProjectParticipationAllOf.md)
 - [ProjectParticipationUpdate](docs/ProjectParticipationUpdate.md)
 - [ProjectParticipationUpdateTimeframe](docs/ProjectParticipationUpdateTimeframe.md)
 - [ProjectScoreDetail](docs/ProjectScoreDetail.md)
 - [ProjectScoreDetailAllOf](docs/ProjectScoreDetailAllOf.md)
 - [ProjectSearch](docs/ProjectSearch.md)
 - [ProjectSearchSkillsInner](docs/ProjectSearchSkillsInner.md)
 - [ProjectStatus](docs/ProjectStatus.md)
 - [SearchResult](docs/SearchResult.md)
 - [SearchResultAllOf](docs/SearchResultAllOf.md)
 - [SearchResultItem](docs/SearchResultItem.md)
 - [Seniority](docs/Seniority.md)
 - [Skill](docs/Skill.md)
 - [SkillAllOf](docs/SkillAllOf.md)
 - [SkillDetails](docs/SkillDetails.md)
 - [SkillLevel](docs/SkillLevel.md)
 - [SkillLevelUpdate](docs/SkillLevelUpdate.md)
 - [SkillLevelUpdateAllOf](docs/SkillLevelUpdateAllOf.md)
 - [Status](docs/Status.md)
 - [Suggestable](docs/Suggestable.md)
 - [Synonymable](docs/Synonymable.md)
 - [Timeframed](docs/Timeframed.md)
 - [Version](docs/Version.md)


## Documentation For Authorization



### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



