# \TracesApi

All URIs are relative to *https://api.argus.eu01.stackit.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1InstancesTracesConfigsList**](TracesApi.md#V1InstancesTracesConfigsList) | **Get** /v1/instances/{instanceId}/traces-configs | 
[**V1InstancesTracesConfigsUpdate**](TracesApi.md#V1InstancesTracesConfigsUpdate) | **Put** /v1/instances/{instanceId}/traces-configs | 
[**V1ProjectsInstancesTracesConfigsList**](TracesApi.md#V1ProjectsInstancesTracesConfigsList) | **Get** /v1/projects/{projectId}/instances/{instanceId}/traces-configs | 
[**V1ProjectsInstancesTracesConfigsUpdate**](TracesApi.md#V1ProjectsInstancesTracesConfigsUpdate) | **Put** /v1/projects/{projectId}/instances/{instanceId}/traces-configs | 



## V1InstancesTracesConfigsList

> TracesConfigResponse V1InstancesTracesConfigsList(ctx, instanceId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.TracesApi.V1InstancesTracesConfigsList(context.Background(), instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesApi.V1InstancesTracesConfigsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesTracesConfigsList`: TracesConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `TracesApi.V1InstancesTracesConfigsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesTracesConfigsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**TracesConfigResponse**](TracesConfigResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesTracesConfigsUpdate

> Message V1InstancesTracesConfigsUpdate(ctx, instanceId).Authorization(authorization).Data(data).Execute()





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
    data := *openapiclient.NewV1InstancesTracesConfigsUpdateRequest("Retention_example") // V1InstancesTracesConfigsUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracesApi.V1InstancesTracesConfigsUpdate(context.Background(), instanceId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesApi.V1InstancesTracesConfigsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesTracesConfigsUpdate`: Message
    fmt.Fprintf(os.Stdout, "Response from `TracesApi.V1InstancesTracesConfigsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesTracesConfigsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesTracesConfigsUpdateRequest**](V1InstancesTracesConfigsUpdateRequest.md) |  | 

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


## V1ProjectsInstancesTracesConfigsList

> TracesConfigResponse V1ProjectsInstancesTracesConfigsList(ctx, instanceId, projectId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.TracesApi.V1ProjectsInstancesTracesConfigsList(context.Background(), instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesApi.V1ProjectsInstancesTracesConfigsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesTracesConfigsList`: TracesConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `TracesApi.V1ProjectsInstancesTracesConfigsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesTracesConfigsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**TracesConfigResponse**](TracesConfigResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesTracesConfigsUpdate

> Message V1ProjectsInstancesTracesConfigsUpdate(ctx, instanceId, projectId).Authorization(authorization).Data(data).Execute()





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
    data := *openapiclient.NewV1InstancesTracesConfigsUpdateRequest("Retention_example") // V1InstancesTracesConfigsUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracesApi.V1ProjectsInstancesTracesConfigsUpdate(context.Background(), instanceId, projectId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracesApi.V1ProjectsInstancesTracesConfigsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesTracesConfigsUpdate`: Message
    fmt.Fprintf(os.Stdout, "Response from `TracesApi.V1ProjectsInstancesTracesConfigsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesTracesConfigsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesTracesConfigsUpdateRequest**](V1InstancesTracesConfigsUpdateRequest.md) |  | 

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

