# \ProjectParticipationAPI

All URIs are relative to *https://yasm.prodyna.com:443/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadPersonProjectParticipation**](ProjectParticipationAPI.md#ReadPersonProjectParticipation) | **Get** /persons/{personId}/project-participation | Get a Project Participation of a Person



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
	resp, r, err := apiClient.ProjectParticipationAPI.ReadPersonProjectParticipation(context.Background(), personId).ProjectIds(projectIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectParticipationAPI.ReadPersonProjectParticipation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadPersonProjectParticipation`: []PersonProjectParticipationDetails
	fmt.Fprintf(os.Stdout, "Response from `ProjectParticipationAPI.ReadPersonProjectParticipation`: %v\n", resp)
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

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

