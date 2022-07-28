# \AlertConfigApi

All URIs are relative to *https://api.argus.eu01.stackit.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1InstancesAlertconfigsList**](AlertConfigApi.md#V1InstancesAlertconfigsList) | **Get** /v1/instances/{instanceId}/alertconfigs | 
[**V1InstancesAlertconfigsUpdate**](AlertConfigApi.md#V1InstancesAlertconfigsUpdate) | **Put** /v1/instances/{instanceId}/alertconfigs | 
[**V1ProjectsInstancesAlertconfigsList**](AlertConfigApi.md#V1ProjectsInstancesAlertconfigsList) | **Get** /v1/projects/{projectId}/instances/{instanceId}/alertconfigs | 
[**V1ProjectsInstancesAlertconfigsUpdate**](AlertConfigApi.md#V1ProjectsInstancesAlertconfigsUpdate) | **Put** /v1/projects/{projectId}/instances/{instanceId}/alertconfigs | 



## V1InstancesAlertconfigsList

> GetAlert V1InstancesAlertconfigsList(ctx, instanceId).Authorization(authorization).Execute()





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
    authorization := "authorization_example" // string | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigApi.V1InstancesAlertconfigsList(context.Background(), instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigApi.V1InstancesAlertconfigsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertconfigsList`: GetAlert
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigApi.V1InstancesAlertconfigsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertconfigsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 

### Return type

[**GetAlert**](GetAlert.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertconfigsUpdate

> PutAlert V1InstancesAlertconfigsUpdate(ctx, instanceId).Authorization(authorization).V1InstancesAlertconfigsUpdateRequest(v1InstancesAlertconfigsUpdateRequest).Execute()





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
    authorization := "authorization_example" // string | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user
    v1InstancesAlertconfigsUpdateRequest := *openapiclient.NewV1InstancesAlertconfigsUpdateRequest(*openapiclient.NewV1InstancesAlertconfigsUpdateRequestRoute("Receiver_example"), []openapiclient.V1InstancesAlertconfigsUpdateRequestReceiversInner{*openapiclient.NewV1InstancesAlertconfigsUpdateRequestReceiversInner("Name_example")}) // V1InstancesAlertconfigsUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigApi.V1InstancesAlertconfigsUpdate(context.Background(), instanceId).Authorization(authorization).V1InstancesAlertconfigsUpdateRequest(v1InstancesAlertconfigsUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigApi.V1InstancesAlertconfigsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertconfigsUpdate`: PutAlert
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigApi.V1InstancesAlertconfigsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertconfigsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 
 **v1InstancesAlertconfigsUpdateRequest** | [**V1InstancesAlertconfigsUpdateRequest**](V1InstancesAlertconfigsUpdateRequest.md) |  | 

### Return type

[**PutAlert**](PutAlert.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertconfigsList

> GetAlert V1ProjectsInstancesAlertconfigsList(ctx, instanceId, projectId).Authorization(authorization).Execute()





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
    authorization := "authorization_example" // string | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigApi.V1ProjectsInstancesAlertconfigsList(context.Background(), instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigApi.V1ProjectsInstancesAlertconfigsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertconfigsList`: GetAlert
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigApi.V1ProjectsInstancesAlertconfigsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertconfigsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 

### Return type

[**GetAlert**](GetAlert.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertconfigsUpdate

> PutAlert V1ProjectsInstancesAlertconfigsUpdate(ctx, instanceId, projectId).Authorization(authorization).V1InstancesAlertconfigsUpdateRequest(v1InstancesAlertconfigsUpdateRequest).Execute()





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
    authorization := "authorization_example" // string | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user
    v1InstancesAlertconfigsUpdateRequest := *openapiclient.NewV1InstancesAlertconfigsUpdateRequest(*openapiclient.NewV1InstancesAlertconfigsUpdateRequestRoute("Receiver_example"), []openapiclient.V1InstancesAlertconfigsUpdateRequestReceiversInner{*openapiclient.NewV1InstancesAlertconfigsUpdateRequestReceiversInner("Name_example")}) // V1InstancesAlertconfigsUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigApi.V1ProjectsInstancesAlertconfigsUpdate(context.Background(), instanceId, projectId).Authorization(authorization).V1InstancesAlertconfigsUpdateRequest(v1InstancesAlertconfigsUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigApi.V1ProjectsInstancesAlertconfigsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertconfigsUpdate`: PutAlert
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigApi.V1ProjectsInstancesAlertconfigsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertconfigsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 
 **v1InstancesAlertconfigsUpdateRequest** | [**V1InstancesAlertconfigsUpdateRequest**](V1InstancesAlertconfigsUpdateRequest.md) |  | 

### Return type

[**PutAlert**](PutAlert.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

