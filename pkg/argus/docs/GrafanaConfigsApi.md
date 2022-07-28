# \GrafanaConfigsApi

All URIs are relative to *https://api.argus.eu01.stackit.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1InstancesGrafanaConfigsList**](GrafanaConfigsApi.md#V1InstancesGrafanaConfigsList) | **Get** /v1/instances/{instanceId}/grafana-configs | 
[**V1InstancesGrafanaConfigsUpdate**](GrafanaConfigsApi.md#V1InstancesGrafanaConfigsUpdate) | **Put** /v1/instances/{instanceId}/grafana-configs | 
[**V1ProjectsInstancesGrafanaConfigsList**](GrafanaConfigsApi.md#V1ProjectsInstancesGrafanaConfigsList) | **Get** /v1/projects/{projectId}/instances/{instanceId}/grafana-configs | 
[**V1ProjectsInstancesGrafanaConfigsUpdate**](GrafanaConfigsApi.md#V1ProjectsInstancesGrafanaConfigsUpdate) | **Put** /v1/projects/{projectId}/instances/{instanceId}/grafana-configs | 



## V1InstancesGrafanaConfigsList

> GrafanaConfigsSerializerRespond V1InstancesGrafanaConfigsList(ctx, instanceId).Authorization(authorization).Execute()





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
    authorization := "authorization_example" // string | Accepts admin basic + token auth and oauth token and authenticated api user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GrafanaConfigsApi.V1InstancesGrafanaConfigsList(context.Background(), instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GrafanaConfigsApi.V1InstancesGrafanaConfigsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesGrafanaConfigsList`: GrafanaConfigsSerializerRespond
    fmt.Fprintf(os.Stdout, "Response from `GrafanaConfigsApi.V1InstancesGrafanaConfigsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesGrafanaConfigsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts admin basic + token auth and oauth token and authenticated api user. | 

### Return type

[**GrafanaConfigsSerializerRespond**](GrafanaConfigsSerializerRespond.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesGrafanaConfigsUpdate

> Message V1InstancesGrafanaConfigsUpdate(ctx, instanceId).Authorization(authorization).V1InstancesGrafanaConfigsUpdateRequest(v1InstancesGrafanaConfigsUpdateRequest).Execute()





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
    authorization := "authorization_example" // string | Accepts admin basic + token auth and oauth token and authenticated api user.
    v1InstancesGrafanaConfigsUpdateRequest := *openapiclient.NewV1InstancesGrafanaConfigsUpdateRequest(false) // V1InstancesGrafanaConfigsUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GrafanaConfigsApi.V1InstancesGrafanaConfigsUpdate(context.Background(), instanceId).Authorization(authorization).V1InstancesGrafanaConfigsUpdateRequest(v1InstancesGrafanaConfigsUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GrafanaConfigsApi.V1InstancesGrafanaConfigsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesGrafanaConfigsUpdate`: Message
    fmt.Fprintf(os.Stdout, "Response from `GrafanaConfigsApi.V1InstancesGrafanaConfigsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesGrafanaConfigsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts admin basic + token auth and oauth token and authenticated api user. | 
 **v1InstancesGrafanaConfigsUpdateRequest** | [**V1InstancesGrafanaConfigsUpdateRequest**](V1InstancesGrafanaConfigsUpdateRequest.md) |  | 

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


## V1ProjectsInstancesGrafanaConfigsList

> GrafanaConfigsSerializerRespond V1ProjectsInstancesGrafanaConfigsList(ctx, instanceId, projectId).Authorization(authorization).Execute()





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
    authorization := "authorization_example" // string | Accepts admin basic + token auth and oauth token and authenticated api user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GrafanaConfigsApi.V1ProjectsInstancesGrafanaConfigsList(context.Background(), instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GrafanaConfigsApi.V1ProjectsInstancesGrafanaConfigsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesGrafanaConfigsList`: GrafanaConfigsSerializerRespond
    fmt.Fprintf(os.Stdout, "Response from `GrafanaConfigsApi.V1ProjectsInstancesGrafanaConfigsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesGrafanaConfigsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts admin basic + token auth and oauth token and authenticated api user. | 

### Return type

[**GrafanaConfigsSerializerRespond**](GrafanaConfigsSerializerRespond.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesGrafanaConfigsUpdate

> Message V1ProjectsInstancesGrafanaConfigsUpdate(ctx, instanceId, projectId).Authorization(authorization).V1InstancesGrafanaConfigsUpdateRequest(v1InstancesGrafanaConfigsUpdateRequest).Execute()





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
    authorization := "authorization_example" // string | Accepts admin basic + token auth and oauth token and authenticated api user.
    v1InstancesGrafanaConfigsUpdateRequest := *openapiclient.NewV1InstancesGrafanaConfigsUpdateRequest(false) // V1InstancesGrafanaConfigsUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GrafanaConfigsApi.V1ProjectsInstancesGrafanaConfigsUpdate(context.Background(), instanceId, projectId).Authorization(authorization).V1InstancesGrafanaConfigsUpdateRequest(v1InstancesGrafanaConfigsUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GrafanaConfigsApi.V1ProjectsInstancesGrafanaConfigsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesGrafanaConfigsUpdate`: Message
    fmt.Fprintf(os.Stdout, "Response from `GrafanaConfigsApi.V1ProjectsInstancesGrafanaConfigsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesGrafanaConfigsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts admin basic + token auth and oauth token and authenticated api user. | 
 **v1InstancesGrafanaConfigsUpdateRequest** | [**V1InstancesGrafanaConfigsUpdateRequest**](V1InstancesGrafanaConfigsUpdateRequest.md) |  | 

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

