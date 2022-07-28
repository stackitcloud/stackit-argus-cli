# \GrafanaPluginsApi

All URIs are relative to *https://api.argus.eu01.stackit.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1GrafanaPluginsCreate**](GrafanaPluginsApi.md#V1GrafanaPluginsCreate) | **Post** /v1/grafana-plugins | 
[**V1GrafanaPluginsDelete**](GrafanaPluginsApi.md#V1GrafanaPluginsDelete) | **Delete** /v1/grafana-plugins/{pluginName} | 
[**V1GrafanaPluginsList**](GrafanaPluginsApi.md#V1GrafanaPluginsList) | **Get** /v1/grafana-plugins | 
[**V1GrafanaPluginsRead**](GrafanaPluginsApi.md#V1GrafanaPluginsRead) | **Get** /v1/grafana-plugins/{pluginName} | 
[**V1GrafanaPluginsUpdate**](GrafanaPluginsApi.md#V1GrafanaPluginsUpdate) | **Put** /v1/grafana-plugins/{pluginName} | 



## V1GrafanaPluginsCreate

> GrafanaPluginSingle V1GrafanaPluginsCreate(ctx).Authorization(authorization).Data(data).Execute()





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
    authorization := "authorization_example" // string | Accepts admin auth.
    data := *openapiclient.NewV1GrafanaPluginsCreateRequest("Name_example", "Description_example") // V1GrafanaPluginsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GrafanaPluginsApi.V1GrafanaPluginsCreate(context.Background()).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GrafanaPluginsApi.V1GrafanaPluginsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GrafanaPluginsCreate`: GrafanaPluginSingle
    fmt.Fprintf(os.Stdout, "Response from `GrafanaPluginsApi.V1GrafanaPluginsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1GrafanaPluginsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Accepts admin auth. | 
 **data** | [**V1GrafanaPluginsCreateRequest**](V1GrafanaPluginsCreateRequest.md) |  | 

### Return type

[**GrafanaPluginSingle**](GrafanaPluginSingle.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GrafanaPluginsDelete

> Message V1GrafanaPluginsDelete(ctx, pluginName).Authorization(authorization).Execute()





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
    pluginName := "pluginName_example" // string | 
    authorization := "authorization_example" // string | Accepts admin auth.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GrafanaPluginsApi.V1GrafanaPluginsDelete(context.Background(), pluginName).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GrafanaPluginsApi.V1GrafanaPluginsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GrafanaPluginsDelete`: Message
    fmt.Fprintf(os.Stdout, "Response from `GrafanaPluginsApi.V1GrafanaPluginsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GrafanaPluginsDeleteRequest struct via the builder pattern


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


## V1GrafanaPluginsList

> GrafanaPlugin V1GrafanaPluginsList(ctx).Authorization(authorization).Execute()





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
    authorization := "authorization_example" // string | Accepts admin auth.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GrafanaPluginsApi.V1GrafanaPluginsList(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GrafanaPluginsApi.V1GrafanaPluginsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GrafanaPluginsList`: GrafanaPlugin
    fmt.Fprintf(os.Stdout, "Response from `GrafanaPluginsApi.V1GrafanaPluginsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1GrafanaPluginsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Accepts admin auth. | 

### Return type

[**GrafanaPlugin**](GrafanaPlugin.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GrafanaPluginsRead

> GrafanaPluginSingle V1GrafanaPluginsRead(ctx, pluginName).Authorization(authorization).Execute()





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
    pluginName := "pluginName_example" // string | 
    authorization := "authorization_example" // string | Accepts admin auth.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GrafanaPluginsApi.V1GrafanaPluginsRead(context.Background(), pluginName).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GrafanaPluginsApi.V1GrafanaPluginsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GrafanaPluginsRead`: GrafanaPluginSingle
    fmt.Fprintf(os.Stdout, "Response from `GrafanaPluginsApi.V1GrafanaPluginsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GrafanaPluginsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts admin auth. | 

### Return type

[**GrafanaPluginSingle**](GrafanaPluginSingle.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GrafanaPluginsUpdate

> GrafanaPluginSingle V1GrafanaPluginsUpdate(ctx, pluginName).Authorization(authorization).Data(data).Execute()





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
    pluginName := "pluginName_example" // string | 
    authorization := "authorization_example" // string | Accepts admin auth.
    data := *openapiclient.NewV1GrafanaPluginsUpdateRequest("Description_example") // V1GrafanaPluginsUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GrafanaPluginsApi.V1GrafanaPluginsUpdate(context.Background(), pluginName).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GrafanaPluginsApi.V1GrafanaPluginsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1GrafanaPluginsUpdate`: GrafanaPluginSingle
    fmt.Fprintf(os.Stdout, "Response from `GrafanaPluginsApi.V1GrafanaPluginsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GrafanaPluginsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts admin auth. | 
 **data** | [**V1GrafanaPluginsUpdateRequest**](V1GrafanaPluginsUpdateRequest.md) |  | 

### Return type

[**GrafanaPluginSingle**](GrafanaPluginSingle.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

