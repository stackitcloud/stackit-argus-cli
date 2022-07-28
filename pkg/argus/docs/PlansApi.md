# \PlansApi

All URIs are relative to *https://api.argus.eu01.stackit.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1PlansList**](PlansApi.md#V1PlansList) | **Get** /v1/plans | 



## V1PlansList

> Plan V1PlansList(ctx).Authorization(authorization).Execute()





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
    authorization := "authorization_example" // string | Accepts admin auth and broker auth.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlansApi.V1PlansList(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlansApi.V1PlansList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1PlansList`: Plan
    fmt.Fprintf(os.Stdout, "Response from `PlansApi.V1PlansList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1PlansListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Accepts admin auth and broker auth. | 

### Return type

[**Plan**](Plan.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

