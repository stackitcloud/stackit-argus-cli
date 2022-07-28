# \OsbApi

All URIs are relative to *https://api.argus.eu01.stackit.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ClustersInstancesStatusList**](OsbApi.md#V1ClustersInstancesStatusList) | **Get** /v1/clusters/{clusterId}/instances/{instanceId}/status | 
[**V1OsbClustersInstancesCreate**](OsbApi.md#V1OsbClustersInstancesCreate) | **Post** /v1/osb/clusters/{clusterId}/instances/{instanceId} | 
[**V1OsbClustersInstancesDelete**](OsbApi.md#V1OsbClustersInstancesDelete) | **Delete** /v1/osb/clusters/{clusterId}/instances/{instanceId} | 
[**V1OsbClustersInstancesUpdate**](OsbApi.md#V1OsbClustersInstancesUpdate) | **Put** /v1/osb/clusters/{clusterId}/instances/{instanceId} | 
[**V1OsbClustersInstancesUsersCreate**](OsbApi.md#V1OsbClustersInstancesUsersCreate) | **Post** /v1/osb/clusters/{clusterId}/instances/{instanceId}/users/{serviceBindingId} | 
[**V1OsbClustersInstancesUsersDelete**](OsbApi.md#V1OsbClustersInstancesUsersDelete) | **Delete** /v1/osb/clusters/{clusterId}/instances/{instanceId}/users/{serviceBindingId} | 
[**V1OsbClustersInstancesUsersList**](OsbApi.md#V1OsbClustersInstancesUsersList) | **Get** /v1/osb/clusters/{clusterId}/instances/{instanceId}/users | 



## V1ClustersInstancesStatusList

> Status V1ClustersInstancesStatusList(ctx, clusterId, instanceId).Authorization(authorization).Execute()





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
    clusterId := "clusterId_example" // string | 
    instanceId := "instanceId_example" // string | 
    authorization := "authorization_example" // string | Accepts admin auth and broker.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OsbApi.V1ClustersInstancesStatusList(context.Background(), clusterId, instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OsbApi.V1ClustersInstancesStatusList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ClustersInstancesStatusList`: Status
    fmt.Fprintf(os.Stdout, "Response from `OsbApi.V1ClustersInstancesStatusList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ClustersInstancesStatusListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts admin auth and broker. | 

### Return type

[**Status**](Status.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OsbClustersInstancesCreate

> Dashboard V1OsbClustersInstancesCreate(ctx, clusterId, instanceId).Authorization(authorization).Data(data).Execute()





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
    clusterId := "clusterId_example" // string | 
    instanceId := "instanceId_example" // string | 
    authorization := "authorization_example" // string | Accepts admin auth and broker auth.
    data := *openapiclient.NewV1OsbClustersInstancesUpdateRequest() // V1OsbClustersInstancesUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OsbApi.V1OsbClustersInstancesCreate(context.Background(), clusterId, instanceId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OsbApi.V1OsbClustersInstancesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OsbClustersInstancesCreate`: Dashboard
    fmt.Fprintf(os.Stdout, "Response from `OsbApi.V1OsbClustersInstancesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OsbClustersInstancesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts admin auth and broker auth. | 
 **data** | [**V1OsbClustersInstancesUpdateRequest**](V1OsbClustersInstancesUpdateRequest.md) |  | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OsbClustersInstancesDelete

> Message V1OsbClustersInstancesDelete(ctx, clusterId, instanceId).Authorization(authorization).Execute()





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
    clusterId := "clusterId_example" // string | 
    instanceId := "instanceId_example" // string | 
    authorization := "authorization_example" // string | Accepts admin auth and broker auth.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OsbApi.V1OsbClustersInstancesDelete(context.Background(), clusterId, instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OsbApi.V1OsbClustersInstancesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OsbClustersInstancesDelete`: Message
    fmt.Fprintf(os.Stdout, "Response from `OsbApi.V1OsbClustersInstancesDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OsbClustersInstancesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts admin auth and broker auth. | 

### Return type

[**Message**](Message.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OsbClustersInstancesUpdate

> Dashboard V1OsbClustersInstancesUpdate(ctx, clusterId, instanceId).Authorization(authorization).Data(data).Execute()





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
    clusterId := "clusterId_example" // string | 
    instanceId := "instanceId_example" // string | 
    authorization := "authorization_example" // string | Accepts admin auth and broker auth.
    data := *openapiclient.NewV1OsbClustersInstancesUpdateRequest() // V1OsbClustersInstancesUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OsbApi.V1OsbClustersInstancesUpdate(context.Background(), clusterId, instanceId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OsbApi.V1OsbClustersInstancesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OsbClustersInstancesUpdate`: Dashboard
    fmt.Fprintf(os.Stdout, "Response from `OsbApi.V1OsbClustersInstancesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OsbClustersInstancesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts admin auth and broker auth. | 
 **data** | [**V1OsbClustersInstancesUpdateRequest**](V1OsbClustersInstancesUpdateRequest.md) |  | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OsbClustersInstancesUsersCreate

> ApiUserCreated V1OsbClustersInstancesUsersCreate(ctx, clusterId, instanceId, serviceBindingId).Authorization(authorization).Execute()





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
    clusterId := "clusterId_example" // string | 
    instanceId := "instanceId_example" // string | 
    serviceBindingId := "serviceBindingId_example" // string | 
    authorization := "authorization_example" // string | Accepts system permissions.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OsbApi.V1OsbClustersInstancesUsersCreate(context.Background(), clusterId, instanceId, serviceBindingId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OsbApi.V1OsbClustersInstancesUsersCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OsbClustersInstancesUsersCreate`: ApiUserCreated
    fmt.Fprintf(os.Stdout, "Response from `OsbApi.V1OsbClustersInstancesUsersCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**instanceId** | **string** |  | 
**serviceBindingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OsbClustersInstancesUsersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts system permissions. | 

### Return type

[**ApiUserCreated**](ApiUserCreated.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OsbClustersInstancesUsersDelete

> Message V1OsbClustersInstancesUsersDelete(ctx, clusterId, instanceId, serviceBindingId).Authorization(authorization).Execute()





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
    clusterId := "clusterId_example" // string | 
    instanceId := "instanceId_example" // string | 
    serviceBindingId := "serviceBindingId_example" // string | 
    authorization := "authorization_example" // string | Accepts system permission.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OsbApi.V1OsbClustersInstancesUsersDelete(context.Background(), clusterId, instanceId, serviceBindingId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OsbApi.V1OsbClustersInstancesUsersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OsbClustersInstancesUsersDelete`: Message
    fmt.Fprintf(os.Stdout, "Response from `OsbApi.V1OsbClustersInstancesUsersDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**instanceId** | **string** |  | 
**serviceBindingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OsbClustersInstancesUsersDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts system permission. | 

### Return type

[**Message**](Message.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OsbClustersInstancesUsersList

> ApiUser V1OsbClustersInstancesUsersList(ctx, clusterId, instanceId).Authorization(authorization).Execute()





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
    clusterId := "clusterId_example" // string | 
    instanceId := "instanceId_example" // string | 
    authorization := "authorization_example" // string | Accepts system permissions.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OsbApi.V1OsbClustersInstancesUsersList(context.Background(), clusterId, instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OsbApi.V1OsbClustersInstancesUsersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1OsbClustersInstancesUsersList`: ApiUser
    fmt.Fprintf(os.Stdout, "Response from `OsbApi.V1OsbClustersInstancesUsersList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OsbClustersInstancesUsersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts system permissions. | 

### Return type

[**ApiUser**](ApiUser.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

