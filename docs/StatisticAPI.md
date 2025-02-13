# \StatisticAPI

All URIs are relative to *https://yasm.prodyna.com:443/api/graph/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProfileStatistics**](StatisticAPI.md#ProfileStatistics) | **Post** /statistics/profile | Profiles statistics over regions (country, office)
[**SkillsEvaluation**](StatisticAPI.md#SkillsEvaluation) | **Post** /statistics/skill | Percentage skills evaluation



## ProfileStatistics

> []ProfileStatistic ProfileStatistics(ctx).Execute()

Profiles statistics over regions (country, office)

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatisticAPI.ProfileStatistics(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatisticAPI.ProfileStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProfileStatistics`: []ProfileStatistic
    fmt.Fprintf(os.Stdout, "Response from `StatisticAPI.ProfileStatistics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProfileStatisticsRequest struct via the builder pattern


### Return type

[**[]ProfileStatistic**](ProfileStatistic.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SkillsEvaluation

> []SkillPercentageDetail SkillsEvaluation(ctx).SkillEvaluationFilter(skillEvaluationFilter).Execute()

Percentage skills evaluation

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
    skillEvaluationFilter := *openapiclient.NewSkillEvaluationFilter(time.Now(), time.Now()) // SkillEvaluationFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatisticAPI.SkillsEvaluation(context.Background()).SkillEvaluationFilter(skillEvaluationFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatisticAPI.SkillsEvaluation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SkillsEvaluation`: []SkillPercentageDetail
    fmt.Fprintf(os.Stdout, "Response from `StatisticAPI.SkillsEvaluation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSkillsEvaluationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skillEvaluationFilter** | [**SkillEvaluationFilter**](SkillEvaluationFilter.md) |  | 

### Return type

[**[]SkillPercentageDetail**](SkillPercentageDetail.md)

### Authorization

[oidcScheme](../README.md#oidcScheme), [bearerScheme](../README.md#bearerScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

