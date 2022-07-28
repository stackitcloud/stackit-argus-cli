# \ClusterApi

All URIs are relative to *https://api.argus.eu01.stackit.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ClustersCreate**](ClusterApi.md#V1ClustersCreate) | **Post** /v1/clusters/{clusterId} | 
[**V1ClustersDelete**](ClusterApi.md#V1ClustersDelete) | **Delete** /v1/clusters/{clusterId} | 
[**V1ClustersList**](ClusterApi.md#V1ClustersList) | **Get** /v1/clusters | 
[**V1MonitoringAuthList**](ClusterApi.md#V1MonitoringAuthList) | **Get** /v1/monitoring-auth | 
[**V1MonitoringSdList**](ClusterApi.md#V1MonitoringSdList) | **Get** /v1/monitoring-sd | 



## V1ClustersCreate

> Message V1ClustersCreate(ctx, clusterId).Authorization(authorization).Data(data).Execute()





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
    authorization := "authorization_example" // string | Accepts admin auth.
    data := *openapiclient.NewV1ClustersCreateRequest("Kubeconfig_example") // V1ClustersCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterApi.V1ClustersCreate(context.Background(), clusterId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterApi.V1ClustersCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ClustersCreate`: Message
    fmt.Fprintf(os.Stdout, "Response from `ClusterApi.V1ClustersCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ClustersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts admin auth. | 
 **data** | [**V1ClustersCreateRequest**](V1ClustersCreateRequest.md) |  | 

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


## V1ClustersDelete

> Message V1ClustersDelete(ctx, clusterId).Authorization(authorization).Execute()





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
    authorization := "authorization_example" // string | Accepts admin auth.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterApi.V1ClustersDelete(context.Background(), clusterId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterApi.V1ClustersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ClustersDelete`: Message
    fmt.Fprintf(os.Stdout, "Response from `ClusterApi.V1ClustersDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ClustersDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts admin auth. | 

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


## V1ClustersList

> ClusterList V1ClustersList(ctx).Authorization(authorization).Search(search).PageIdx(pageIdx).Execute()





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
    authorization := "authorization_example" // string | Accepts admin auth
    search := "search_example" // string | Search string that should be included in the cluster name (optional)
    pageIdx := float32(8.14) // float32 | Page index if the clusters are having paging. 4 clusters per page are returned (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterApi.V1ClustersList(context.Background()).Authorization(authorization).Search(search).PageIdx(pageIdx).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterApi.V1ClustersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ClustersList`: ClusterList
    fmt.Fprintf(os.Stdout, "Response from `ClusterApi.V1ClustersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ClustersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Accepts admin auth | 
 **search** | **string** | Search string that should be included in the cluster name | 
 **pageIdx** | **float32** | Page index if the clusters are having paging. 4 clusters per page are returned | 

### Return type

[**ClusterList**](ClusterList.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MonitoringAuthList

> Message V1MonitoringAuthList(ctx).Authorization(authorization).Execute()





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
    authorization := "authorization_example" // string | Accepts admin auth

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterApi.V1MonitoringAuthList(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterApi.V1MonitoringAuthList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1MonitoringAuthList`: Message
    fmt.Fprintf(os.Stdout, "Response from `ClusterApi.V1MonitoringAuthList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MonitoringAuthListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Accepts admin auth | 

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


## V1MonitoringSdList

> map[string]interface{} V1MonitoringSdList(ctx).Authorization(authorization).Execute()





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
    authorization := "authorization_example" // string | Accepts admin auth

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterApi.V1MonitoringSdList(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterApi.V1MonitoringSdList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1MonitoringSdList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClusterApi.V1MonitoringSdList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1MonitoringSdListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Accepts admin auth | 

### Return type

**map[string]interface{}**

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

