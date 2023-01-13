# \BucketRetentionApi

All URIs are relative to *https://api.argus.eu01.stackit.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1InstancesBucketRetentionsList**](BucketRetentionApi.md#V1InstancesBucketRetentionsList) | **Get** /v1/instances/{instanceId}/bucket-retentions | 
[**V1InstancesBucketRetentionsUpdate**](BucketRetentionApi.md#V1InstancesBucketRetentionsUpdate) | **Put** /v1/instances/{instanceId}/bucket-retentions | 
[**V1ProjectsInstancesBucketRetentionsList**](BucketRetentionApi.md#V1ProjectsInstancesBucketRetentionsList) | **Get** /v1/projects/{projectId}/instances/{instanceId}/bucket-retentions | 
[**V1ProjectsInstancesBucketRetentionsUpdate**](BucketRetentionApi.md#V1ProjectsInstancesBucketRetentionsUpdate) | **Put** /v1/projects/{projectId}/instances/{instanceId}/bucket-retentions | 



## V1InstancesBucketRetentionsList

> BucketRetentionTimeRespond V1InstancesBucketRetentionsList(ctx, instanceId).Authorization(authorization).Execute()





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
    instanceId := "instanceId_example" // string | 
    authorization := "authorization_example" // string | Accepts admin basic + token auth and oauth token.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketRetentionApi.V1InstancesBucketRetentionsList(context.Background(), instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketRetentionApi.V1InstancesBucketRetentionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesBucketRetentionsList`: BucketRetentionTimeRespond
    fmt.Fprintf(os.Stdout, "Response from `BucketRetentionApi.V1InstancesBucketRetentionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesBucketRetentionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts admin basic + token auth and oauth token. | 

### Return type

[**BucketRetentionTimeRespond**](BucketRetentionTimeRespond.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesBucketRetentionsUpdate

> Message V1InstancesBucketRetentionsUpdate(ctx, instanceId).Authorization(authorization).V1InstancesBucketRetentionsUpdateRequest(v1InstancesBucketRetentionsUpdateRequest).Execute()





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
    instanceId := "instanceId_example" // string | 
    authorization := "authorization_example" // string | Accepts admin basic + token auth and oauth token.
    v1InstancesBucketRetentionsUpdateRequest := *openapiclient.NewV1InstancesBucketRetentionsUpdateRequest("BucketRetentionTimeRaw_example", "BucketRetentionTime5m_example", "BucketRetentionTime1h_example") // V1InstancesBucketRetentionsUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketRetentionApi.V1InstancesBucketRetentionsUpdate(context.Background(), instanceId).Authorization(authorization).V1InstancesBucketRetentionsUpdateRequest(v1InstancesBucketRetentionsUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketRetentionApi.V1InstancesBucketRetentionsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesBucketRetentionsUpdate`: Message
    fmt.Fprintf(os.Stdout, "Response from `BucketRetentionApi.V1InstancesBucketRetentionsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesBucketRetentionsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts admin basic + token auth and oauth token. | 
 **v1InstancesBucketRetentionsUpdateRequest** | [**V1InstancesBucketRetentionsUpdateRequest**](V1InstancesBucketRetentionsUpdateRequest.md) |  | 

### Return type

[**Message**](Message.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesBucketRetentionsList

> BucketRetentionTimeRespond V1ProjectsInstancesBucketRetentionsList(ctx, instanceId, projectId).Authorization(authorization).Execute()





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
    instanceId := "instanceId_example" // string | 
    projectId := "projectId_example" // string | 
    authorization := "authorization_example" // string | Accepts admin basic + token auth and oauth token.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketRetentionApi.V1ProjectsInstancesBucketRetentionsList(context.Background(), instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketRetentionApi.V1ProjectsInstancesBucketRetentionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesBucketRetentionsList`: BucketRetentionTimeRespond
    fmt.Fprintf(os.Stdout, "Response from `BucketRetentionApi.V1ProjectsInstancesBucketRetentionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesBucketRetentionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts admin basic + token auth and oauth token. | 

### Return type

[**BucketRetentionTimeRespond**](BucketRetentionTimeRespond.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesBucketRetentionsUpdate

> Message V1ProjectsInstancesBucketRetentionsUpdate(ctx, instanceId, projectId).Authorization(authorization).V1InstancesBucketRetentionsUpdateRequest(v1InstancesBucketRetentionsUpdateRequest).Execute()





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
    instanceId := "instanceId_example" // string | 
    projectId := "projectId_example" // string | 
    authorization := "authorization_example" // string | Accepts admin basic + token auth and oauth token.
    v1InstancesBucketRetentionsUpdateRequest := *openapiclient.NewV1InstancesBucketRetentionsUpdateRequest("BucketRetentionTimeRaw_example", "BucketRetentionTime5m_example", "BucketRetentionTime1h_example") // V1InstancesBucketRetentionsUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketRetentionApi.V1ProjectsInstancesBucketRetentionsUpdate(context.Background(), instanceId, projectId).Authorization(authorization).V1InstancesBucketRetentionsUpdateRequest(v1InstancesBucketRetentionsUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketRetentionApi.V1ProjectsInstancesBucketRetentionsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesBucketRetentionsUpdate`: Message
    fmt.Fprintf(os.Stdout, "Response from `BucketRetentionApi.V1ProjectsInstancesBucketRetentionsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesBucketRetentionsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts admin basic + token auth and oauth token. | 
 **v1InstancesBucketRetentionsUpdateRequest** | [**V1InstancesBucketRetentionsUpdateRequest**](V1InstancesBucketRetentionsUpdateRequest.md) |  | 

### Return type

[**Message**](Message.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

