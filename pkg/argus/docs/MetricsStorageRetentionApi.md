# \MetricsStorageRetentionApi

All URIs are relative to *https://api.argus.eu01.stackit.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1InstancesMetricsStorageRetentionsList**](MetricsStorageRetentionApi.md#V1InstancesMetricsStorageRetentionsList) | **Get** /v1/instances/{instanceId}/metrics-storage-retentions | 
[**V1InstancesMetricsStorageRetentionsUpdate**](MetricsStorageRetentionApi.md#V1InstancesMetricsStorageRetentionsUpdate) | **Put** /v1/instances/{instanceId}/metrics-storage-retentions | 
[**V1ProjectsInstancesMetricsStorageRetentionsList**](MetricsStorageRetentionApi.md#V1ProjectsInstancesMetricsStorageRetentionsList) | **Get** /v1/projects/{projectId}/instances/{instanceId}/metrics-storage-retentions | 
[**V1ProjectsInstancesMetricsStorageRetentionsUpdate**](MetricsStorageRetentionApi.md#V1ProjectsInstancesMetricsStorageRetentionsUpdate) | **Put** /v1/projects/{projectId}/instances/{instanceId}/metrics-storage-retentions | 



## V1InstancesMetricsStorageRetentionsList

> BucketRetentionTimeRespond V1InstancesMetricsStorageRetentionsList(ctx, instanceId).Authorization(authorization).Execute()





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
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsStorageRetentionApi.V1InstancesMetricsStorageRetentionsList(context.Background(), instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsStorageRetentionApi.V1InstancesMetricsStorageRetentionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesMetricsStorageRetentionsList`: BucketRetentionTimeRespond
    fmt.Fprintf(os.Stdout, "Response from `MetricsStorageRetentionApi.V1InstancesMetricsStorageRetentionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesMetricsStorageRetentionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**BucketRetentionTimeRespond**](BucketRetentionTimeRespond.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesMetricsStorageRetentionsUpdate

> Message V1InstancesMetricsStorageRetentionsUpdate(ctx, instanceId).Authorization(authorization).Data(data).Execute()





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
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.
    data := *openapiclient.NewV1InstancesMetricsStorageRetentionsUpdateRequest("MetricsRetentionTimeRaw_example", "MetricsRetentionTime5m_example", "MetricsRetentionTime1h_example") // V1InstancesMetricsStorageRetentionsUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsStorageRetentionApi.V1InstancesMetricsStorageRetentionsUpdate(context.Background(), instanceId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsStorageRetentionApi.V1InstancesMetricsStorageRetentionsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesMetricsStorageRetentionsUpdate`: Message
    fmt.Fprintf(os.Stdout, "Response from `MetricsStorageRetentionApi.V1InstancesMetricsStorageRetentionsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesMetricsStorageRetentionsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesMetricsStorageRetentionsUpdateRequest**](V1InstancesMetricsStorageRetentionsUpdateRequest.md) |  | 

### Return type

[**Message**](Message.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesMetricsStorageRetentionsList

> BucketRetentionTimeRespond V1ProjectsInstancesMetricsStorageRetentionsList(ctx, instanceId, projectId).Authorization(authorization).Execute()





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
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsStorageRetentionApi.V1ProjectsInstancesMetricsStorageRetentionsList(context.Background(), instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsStorageRetentionApi.V1ProjectsInstancesMetricsStorageRetentionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesMetricsStorageRetentionsList`: BucketRetentionTimeRespond
    fmt.Fprintf(os.Stdout, "Response from `MetricsStorageRetentionApi.V1ProjectsInstancesMetricsStorageRetentionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesMetricsStorageRetentionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**BucketRetentionTimeRespond**](BucketRetentionTimeRespond.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesMetricsStorageRetentionsUpdate

> Message V1ProjectsInstancesMetricsStorageRetentionsUpdate(ctx, instanceId, projectId).Authorization(authorization).Data(data).Execute()





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
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.
    data := *openapiclient.NewV1InstancesMetricsStorageRetentionsUpdateRequest("MetricsRetentionTimeRaw_example", "MetricsRetentionTime5m_example", "MetricsRetentionTime1h_example") // V1InstancesMetricsStorageRetentionsUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsStorageRetentionApi.V1ProjectsInstancesMetricsStorageRetentionsUpdate(context.Background(), instanceId, projectId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsStorageRetentionApi.V1ProjectsInstancesMetricsStorageRetentionsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesMetricsStorageRetentionsUpdate`: Message
    fmt.Fprintf(os.Stdout, "Response from `MetricsStorageRetentionApi.V1ProjectsInstancesMetricsStorageRetentionsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesMetricsStorageRetentionsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesMetricsStorageRetentionsUpdateRequest**](V1InstancesMetricsStorageRetentionsUpdateRequest.md) |  | 

### Return type

[**Message**](Message.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

