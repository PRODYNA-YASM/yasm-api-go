# \StatisticAPI

All URIs are relative to *https://yasm.prodyna.com:443/api/graph/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SkillsEvaluation**](StatisticAPI.md#SkillsEvaluation) | **Post** /statistics/skill | Percentage skills evaluation



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

