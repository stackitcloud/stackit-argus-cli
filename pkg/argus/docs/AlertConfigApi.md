# \AlertConfigApi

All URIs are relative to *https://api.argus.eu01.stackit.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1InstancesAlertconfigsList**](AlertConfigApi.md#V1InstancesAlertconfigsList) | **Get** /v1/instances/{instanceId}/alertconfigs | 
[**V1InstancesAlertconfigsReceiversCreate**](AlertConfigApi.md#V1InstancesAlertconfigsReceiversCreate) | **Post** /v1/instances/{instanceId}/alertconfigs/receivers | 
[**V1InstancesAlertconfigsReceiversDelete**](AlertConfigApi.md#V1InstancesAlertconfigsReceiversDelete) | **Delete** /v1/instances/{instanceId}/alertconfigs/receivers/{receiver} | 
[**V1InstancesAlertconfigsReceiversList**](AlertConfigApi.md#V1InstancesAlertconfigsReceiversList) | **Get** /v1/instances/{instanceId}/alertconfigs/receivers | 
[**V1InstancesAlertconfigsReceiversRead**](AlertConfigApi.md#V1InstancesAlertconfigsReceiversRead) | **Get** /v1/instances/{instanceId}/alertconfigs/receivers/{receiver} | 
[**V1InstancesAlertconfigsReceiversUpdate**](AlertConfigApi.md#V1InstancesAlertconfigsReceiversUpdate) | **Put** /v1/instances/{instanceId}/alertconfigs/receivers/{receiver} | 
[**V1InstancesAlertconfigsRoutesCreate**](AlertConfigApi.md#V1InstancesAlertconfigsRoutesCreate) | **Post** /v1/instances/{instanceId}/alertconfigs/routes | 
[**V1InstancesAlertconfigsRoutesDelete**](AlertConfigApi.md#V1InstancesAlertconfigsRoutesDelete) | **Delete** /v1/instances/{instanceId}/alertconfigs/routes/{receiver} | 
[**V1InstancesAlertconfigsRoutesList**](AlertConfigApi.md#V1InstancesAlertconfigsRoutesList) | **Get** /v1/instances/{instanceId}/alertconfigs/routes | 
[**V1InstancesAlertconfigsRoutesRead**](AlertConfigApi.md#V1InstancesAlertconfigsRoutesRead) | **Get** /v1/instances/{instanceId}/alertconfigs/routes/{receiver} | 
[**V1InstancesAlertconfigsRoutesUpdate**](AlertConfigApi.md#V1InstancesAlertconfigsRoutesUpdate) | **Put** /v1/instances/{instanceId}/alertconfigs/routes/{receiver} | 
[**V1InstancesAlertconfigsUpdate**](AlertConfigApi.md#V1InstancesAlertconfigsUpdate) | **Put** /v1/instances/{instanceId}/alertconfigs | 
[**V1ProjectsInstancesAlertconfigsList**](AlertConfigApi.md#V1ProjectsInstancesAlertconfigsList) | **Get** /v1/projects/{projectId}/instances/{instanceId}/alertconfigs | 
[**V1ProjectsInstancesAlertconfigsReceiversCreate**](AlertConfigApi.md#V1ProjectsInstancesAlertconfigsReceiversCreate) | **Post** /v1/projects/{projectId}/instances/{instanceId}/alertconfigs/receivers | 
[**V1ProjectsInstancesAlertconfigsReceiversDelete**](AlertConfigApi.md#V1ProjectsInstancesAlertconfigsReceiversDelete) | **Delete** /v1/projects/{projectId}/instances/{instanceId}/alertconfigs/receivers/{receiver} | 
[**V1ProjectsInstancesAlertconfigsReceiversList**](AlertConfigApi.md#V1ProjectsInstancesAlertconfigsReceiversList) | **Get** /v1/projects/{projectId}/instances/{instanceId}/alertconfigs/receivers | 
[**V1ProjectsInstancesAlertconfigsReceiversRead**](AlertConfigApi.md#V1ProjectsInstancesAlertconfigsReceiversRead) | **Get** /v1/projects/{projectId}/instances/{instanceId}/alertconfigs/receivers/{receiver} | 
[**V1ProjectsInstancesAlertconfigsReceiversUpdate**](AlertConfigApi.md#V1ProjectsInstancesAlertconfigsReceiversUpdate) | **Put** /v1/projects/{projectId}/instances/{instanceId}/alertconfigs/receivers/{receiver} | 
[**V1ProjectsInstancesAlertconfigsRoutesCreate**](AlertConfigApi.md#V1ProjectsInstancesAlertconfigsRoutesCreate) | **Post** /v1/projects/{projectId}/instances/{instanceId}/alertconfigs/routes | 
[**V1ProjectsInstancesAlertconfigsRoutesDelete**](AlertConfigApi.md#V1ProjectsInstancesAlertconfigsRoutesDelete) | **Delete** /v1/projects/{projectId}/instances/{instanceId}/alertconfigs/routes/{receiver} | 
[**V1ProjectsInstancesAlertconfigsRoutesList**](AlertConfigApi.md#V1ProjectsInstancesAlertconfigsRoutesList) | **Get** /v1/projects/{projectId}/instances/{instanceId}/alertconfigs/routes | 
[**V1ProjectsInstancesAlertconfigsRoutesRead**](AlertConfigApi.md#V1ProjectsInstancesAlertconfigsRoutesRead) | **Get** /v1/projects/{projectId}/instances/{instanceId}/alertconfigs/routes/{receiver} | 
[**V1ProjectsInstancesAlertconfigsRoutesUpdate**](AlertConfigApi.md#V1ProjectsInstancesAlertconfigsRoutesUpdate) | **Put** /v1/projects/{projectId}/instances/{instanceId}/alertconfigs/routes/{receiver} | 
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
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.

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

 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**GetAlert**](GetAlert.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertconfigsReceiversCreate

> ReceiversResponse V1InstancesAlertconfigsReceiversCreate(ctx, instanceId).Authorization(authorization).Data(data).Execute()





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
    data := *openapiclient.NewV1InstancesAlertconfigsUpdateRequestReceiversInner("Name_example") // V1InstancesAlertconfigsUpdateRequestReceiversInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigApi.V1InstancesAlertconfigsReceiversCreate(context.Background(), instanceId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigApi.V1InstancesAlertconfigsReceiversCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertconfigsReceiversCreate`: ReceiversResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigApi.V1InstancesAlertconfigsReceiversCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertconfigsReceiversCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesAlertconfigsUpdateRequestReceiversInner**](V1InstancesAlertconfigsUpdateRequestReceiversInner.md) |  | 

### Return type

[**ReceiversResponse**](ReceiversResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertconfigsReceiversDelete

> ReceiversResponse V1InstancesAlertconfigsReceiversDelete(ctx, instanceId, receiver).Authorization(authorization).Execute()





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
    receiver := "receiver_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigApi.V1InstancesAlertconfigsReceiversDelete(context.Background(), instanceId, receiver).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigApi.V1InstancesAlertconfigsReceiversDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertconfigsReceiversDelete`: ReceiversResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigApi.V1InstancesAlertconfigsReceiversDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**receiver** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertconfigsReceiversDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**ReceiversResponse**](ReceiversResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertconfigsReceiversList

> ReceiversResponse V1InstancesAlertconfigsReceiversList(ctx, instanceId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.AlertConfigApi.V1InstancesAlertconfigsReceiversList(context.Background(), instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigApi.V1InstancesAlertconfigsReceiversList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertconfigsReceiversList`: ReceiversResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigApi.V1InstancesAlertconfigsReceiversList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertconfigsReceiversListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**ReceiversResponse**](ReceiversResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertconfigsReceiversRead

> ReceiversResponseSerializerSingle V1InstancesAlertconfigsReceiversRead(ctx, instanceId, receiver).Authorization(authorization).Execute()





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
    receiver := "receiver_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigApi.V1InstancesAlertconfigsReceiversRead(context.Background(), instanceId, receiver).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigApi.V1InstancesAlertconfigsReceiversRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertconfigsReceiversRead`: ReceiversResponseSerializerSingle
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigApi.V1InstancesAlertconfigsReceiversRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**receiver** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertconfigsReceiversReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**ReceiversResponseSerializerSingle**](ReceiversResponseSerializerSingle.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertconfigsReceiversUpdate

> ReceiversResponse V1InstancesAlertconfigsReceiversUpdate(ctx, instanceId, receiver).Authorization(authorization).Data(data).Execute()





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
    receiver := "receiver_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.
    data := *openapiclient.NewV1InstancesAlertconfigsUpdateRequestReceiversInner("Name_example") // V1InstancesAlertconfigsUpdateRequestReceiversInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigApi.V1InstancesAlertconfigsReceiversUpdate(context.Background(), instanceId, receiver).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigApi.V1InstancesAlertconfigsReceiversUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertconfigsReceiversUpdate`: ReceiversResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigApi.V1InstancesAlertconfigsReceiversUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**receiver** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertconfigsReceiversUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesAlertconfigsUpdateRequestReceiversInner**](V1InstancesAlertconfigsUpdateRequestReceiversInner.md) |  | 

### Return type

[**ReceiversResponse**](ReceiversResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertconfigsRoutesCreate

> RouteResponse V1InstancesAlertconfigsRoutesCreate(ctx, instanceId).Authorization(authorization).Data(data).Execute()





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
    data := *openapiclient.NewV1InstancesAlertconfigsUpdateRequestRoute("Receiver_example") // V1InstancesAlertconfigsUpdateRequestRoute | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigApi.V1InstancesAlertconfigsRoutesCreate(context.Background(), instanceId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigApi.V1InstancesAlertconfigsRoutesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertconfigsRoutesCreate`: RouteResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigApi.V1InstancesAlertconfigsRoutesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertconfigsRoutesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesAlertconfigsUpdateRequestRoute**](V1InstancesAlertconfigsUpdateRequestRoute.md) |  | 

### Return type

[**RouteResponse**](RouteResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertconfigsRoutesDelete

> RouteResponse V1InstancesAlertconfigsRoutesDelete(ctx, instanceId, receiver).Authorization(authorization).Execute()





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
    receiver := "receiver_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigApi.V1InstancesAlertconfigsRoutesDelete(context.Background(), instanceId, receiver).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigApi.V1InstancesAlertconfigsRoutesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertconfigsRoutesDelete`: RouteResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigApi.V1InstancesAlertconfigsRoutesDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**receiver** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertconfigsRoutesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**RouteResponse**](RouteResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertconfigsRoutesList

> RouteResponse V1InstancesAlertconfigsRoutesList(ctx, instanceId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.AlertConfigApi.V1InstancesAlertconfigsRoutesList(context.Background(), instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigApi.V1InstancesAlertconfigsRoutesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertconfigsRoutesList`: RouteResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigApi.V1InstancesAlertconfigsRoutesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertconfigsRoutesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**RouteResponse**](RouteResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertconfigsRoutesRead

> RouteResponse V1InstancesAlertconfigsRoutesRead(ctx, instanceId, receiver).Authorization(authorization).Execute()





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
    receiver := "receiver_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigApi.V1InstancesAlertconfigsRoutesRead(context.Background(), instanceId, receiver).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigApi.V1InstancesAlertconfigsRoutesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertconfigsRoutesRead`: RouteResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigApi.V1InstancesAlertconfigsRoutesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**receiver** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertconfigsRoutesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**RouteResponse**](RouteResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertconfigsRoutesUpdate

> RouteResponse V1InstancesAlertconfigsRoutesUpdate(ctx, instanceId, receiver).Authorization(authorization).Data(data).Execute()





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
    receiver := "receiver_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.
    data := *openapiclient.NewV1InstancesAlertconfigsUpdateRequestRoute("Receiver_example") // V1InstancesAlertconfigsUpdateRequestRoute | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigApi.V1InstancesAlertconfigsRoutesUpdate(context.Background(), instanceId, receiver).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigApi.V1InstancesAlertconfigsRoutesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertconfigsRoutesUpdate`: RouteResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigApi.V1InstancesAlertconfigsRoutesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**receiver** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertconfigsRoutesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesAlertconfigsUpdateRequestRoute**](V1InstancesAlertconfigsUpdateRequestRoute.md) |  | 

### Return type

[**RouteResponse**](RouteResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertconfigsUpdate

> PutAlert V1InstancesAlertconfigsUpdate(ctx, instanceId).Authorization(authorization).Data(data).Execute()





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
    data := *openapiclient.NewV1InstancesAlertconfigsUpdateRequest(*openapiclient.NewV1InstancesAlertconfigsUpdateRequestRoute("Receiver_example"), []openapiclient.V1InstancesAlertconfigsUpdateRequestReceiversInner{*openapiclient.NewV1InstancesAlertconfigsUpdateRequestReceiversInner("Name_example")}) // V1InstancesAlertconfigsUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigApi.V1InstancesAlertconfigsUpdate(context.Background(), instanceId).Authorization(authorization).Data(data).Execute()
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

 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesAlertconfigsUpdateRequest**](V1InstancesAlertconfigsUpdateRequest.md) |  | 

### Return type

[**PutAlert**](PutAlert.md)

### Authorization

[Basic](../README.md#Basic)

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
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.

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


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**GetAlert**](GetAlert.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertconfigsReceiversCreate

> ReceiversResponse V1ProjectsInstancesAlertconfigsReceiversCreate(ctx, instanceId, projectId).Authorization(authorization).Data(data).Execute()





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
    data := *openapiclient.NewV1InstancesAlertconfigsUpdateRequestReceiversInner("Name_example") // V1InstancesAlertconfigsUpdateRequestReceiversInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigApi.V1ProjectsInstancesAlertconfigsReceiversCreate(context.Background(), instanceId, projectId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigApi.V1ProjectsInstancesAlertconfigsReceiversCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertconfigsReceiversCreate`: ReceiversResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigApi.V1ProjectsInstancesAlertconfigsReceiversCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertconfigsReceiversCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesAlertconfigsUpdateRequestReceiversInner**](V1InstancesAlertconfigsUpdateRequestReceiversInner.md) |  | 

### Return type

[**ReceiversResponse**](ReceiversResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertconfigsReceiversDelete

> ReceiversResponse V1ProjectsInstancesAlertconfigsReceiversDelete(ctx, instanceId, projectId, receiver).Authorization(authorization).Execute()





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
    receiver := "receiver_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigApi.V1ProjectsInstancesAlertconfigsReceiversDelete(context.Background(), instanceId, projectId, receiver).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigApi.V1ProjectsInstancesAlertconfigsReceiversDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertconfigsReceiversDelete`: ReceiversResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigApi.V1ProjectsInstancesAlertconfigsReceiversDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 
**receiver** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertconfigsReceiversDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**ReceiversResponse**](ReceiversResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertconfigsReceiversList

> ReceiversResponse V1ProjectsInstancesAlertconfigsReceiversList(ctx, instanceId, projectId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.AlertConfigApi.V1ProjectsInstancesAlertconfigsReceiversList(context.Background(), instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigApi.V1ProjectsInstancesAlertconfigsReceiversList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertconfigsReceiversList`: ReceiversResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigApi.V1ProjectsInstancesAlertconfigsReceiversList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertconfigsReceiversListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**ReceiversResponse**](ReceiversResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertconfigsReceiversRead

> ReceiversResponseSerializerSingle V1ProjectsInstancesAlertconfigsReceiversRead(ctx, instanceId, projectId, receiver).Authorization(authorization).Execute()





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
    receiver := "receiver_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigApi.V1ProjectsInstancesAlertconfigsReceiversRead(context.Background(), instanceId, projectId, receiver).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigApi.V1ProjectsInstancesAlertconfigsReceiversRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertconfigsReceiversRead`: ReceiversResponseSerializerSingle
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigApi.V1ProjectsInstancesAlertconfigsReceiversRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 
**receiver** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertconfigsReceiversReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**ReceiversResponseSerializerSingle**](ReceiversResponseSerializerSingle.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertconfigsReceiversUpdate

> ReceiversResponse V1ProjectsInstancesAlertconfigsReceiversUpdate(ctx, instanceId, projectId, receiver).Authorization(authorization).Data(data).Execute()





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
    receiver := "receiver_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.
    data := *openapiclient.NewV1InstancesAlertconfigsUpdateRequestReceiversInner("Name_example") // V1InstancesAlertconfigsUpdateRequestReceiversInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigApi.V1ProjectsInstancesAlertconfigsReceiversUpdate(context.Background(), instanceId, projectId, receiver).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigApi.V1ProjectsInstancesAlertconfigsReceiversUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertconfigsReceiversUpdate`: ReceiversResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigApi.V1ProjectsInstancesAlertconfigsReceiversUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 
**receiver** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertconfigsReceiversUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesAlertconfigsUpdateRequestReceiversInner**](V1InstancesAlertconfigsUpdateRequestReceiversInner.md) |  | 

### Return type

[**ReceiversResponse**](ReceiversResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertconfigsRoutesCreate

> RouteResponse V1ProjectsInstancesAlertconfigsRoutesCreate(ctx, instanceId, projectId).Authorization(authorization).Data(data).Execute()





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
    data := *openapiclient.NewV1InstancesAlertconfigsUpdateRequestRoute("Receiver_example") // V1InstancesAlertconfigsUpdateRequestRoute | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigApi.V1ProjectsInstancesAlertconfigsRoutesCreate(context.Background(), instanceId, projectId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigApi.V1ProjectsInstancesAlertconfigsRoutesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertconfigsRoutesCreate`: RouteResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigApi.V1ProjectsInstancesAlertconfigsRoutesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertconfigsRoutesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesAlertconfigsUpdateRequestRoute**](V1InstancesAlertconfigsUpdateRequestRoute.md) |  | 

### Return type

[**RouteResponse**](RouteResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertconfigsRoutesDelete

> RouteResponse V1ProjectsInstancesAlertconfigsRoutesDelete(ctx, instanceId, projectId, receiver).Authorization(authorization).Execute()





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
    receiver := "receiver_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigApi.V1ProjectsInstancesAlertconfigsRoutesDelete(context.Background(), instanceId, projectId, receiver).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigApi.V1ProjectsInstancesAlertconfigsRoutesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertconfigsRoutesDelete`: RouteResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigApi.V1ProjectsInstancesAlertconfigsRoutesDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 
**receiver** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertconfigsRoutesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**RouteResponse**](RouteResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertconfigsRoutesList

> RouteResponse V1ProjectsInstancesAlertconfigsRoutesList(ctx, instanceId, projectId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.AlertConfigApi.V1ProjectsInstancesAlertconfigsRoutesList(context.Background(), instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigApi.V1ProjectsInstancesAlertconfigsRoutesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertconfigsRoutesList`: RouteResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigApi.V1ProjectsInstancesAlertconfigsRoutesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertconfigsRoutesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**RouteResponse**](RouteResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertconfigsRoutesRead

> RouteResponse V1ProjectsInstancesAlertconfigsRoutesRead(ctx, instanceId, projectId, receiver).Authorization(authorization).Execute()





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
    receiver := "receiver_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigApi.V1ProjectsInstancesAlertconfigsRoutesRead(context.Background(), instanceId, projectId, receiver).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigApi.V1ProjectsInstancesAlertconfigsRoutesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertconfigsRoutesRead`: RouteResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigApi.V1ProjectsInstancesAlertconfigsRoutesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 
**receiver** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertconfigsRoutesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**RouteResponse**](RouteResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertconfigsRoutesUpdate

> RouteResponse V1ProjectsInstancesAlertconfigsRoutesUpdate(ctx, instanceId, projectId, receiver).Authorization(authorization).Data(data).Execute()





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
    receiver := "receiver_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.
    data := *openapiclient.NewV1InstancesAlertconfigsUpdateRequestRoute("Receiver_example") // V1InstancesAlertconfigsUpdateRequestRoute | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigApi.V1ProjectsInstancesAlertconfigsRoutesUpdate(context.Background(), instanceId, projectId, receiver).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertConfigApi.V1ProjectsInstancesAlertconfigsRoutesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertconfigsRoutesUpdate`: RouteResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertConfigApi.V1ProjectsInstancesAlertconfigsRoutesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 
**receiver** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertconfigsRoutesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesAlertconfigsUpdateRequestRoute**](V1InstancesAlertconfigsUpdateRequestRoute.md) |  | 

### Return type

[**RouteResponse**](RouteResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertconfigsUpdate

> PutAlert V1ProjectsInstancesAlertconfigsUpdate(ctx, instanceId, projectId).Authorization(authorization).Data(data).Execute()





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
    data := *openapiclient.NewV1InstancesAlertconfigsUpdateRequest(*openapiclient.NewV1InstancesAlertconfigsUpdateRequestRoute("Receiver_example"), []openapiclient.V1InstancesAlertconfigsUpdateRequestReceiversInner{*openapiclient.NewV1InstancesAlertconfigsUpdateRequestReceiversInner("Name_example")}) // V1InstancesAlertconfigsUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertConfigApi.V1ProjectsInstancesAlertconfigsUpdate(context.Background(), instanceId, projectId).Authorization(authorization).Data(data).Execute()
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


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesAlertconfigsUpdateRequest**](V1InstancesAlertconfigsUpdateRequest.md) |  | 

### Return type

[**PutAlert**](PutAlert.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

