# \LogsApi

All URIs are relative to *https://api.argus.eu01.stackit.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1InstancesLogsAlertgroupsCreate**](LogsApi.md#V1InstancesLogsAlertgroupsCreate) | **Post** /v1/instances/{instanceId}/logs-alertgroups | 
[**V1InstancesLogsAlertgroupsDelete**](LogsApi.md#V1InstancesLogsAlertgroupsDelete) | **Delete** /v1/instances/{instanceId}/logs-alertgroups/{groupName} | 
[**V1InstancesLogsAlertgroupsList**](LogsApi.md#V1InstancesLogsAlertgroupsList) | **Get** /v1/instances/{instanceId}/logs-alertgroups | 
[**V1InstancesLogsAlertgroupsRead**](LogsApi.md#V1InstancesLogsAlertgroupsRead) | **Get** /v1/instances/{instanceId}/logs-alertgroups/{groupName} | 
[**V1InstancesLogsAlertgroupsUpdate**](LogsApi.md#V1InstancesLogsAlertgroupsUpdate) | **Put** /v1/instances/{instanceId}/logs-alertgroups/{groupName} | 
[**V1InstancesLogsConfigsList**](LogsApi.md#V1InstancesLogsConfigsList) | **Get** /v1/instances/{instanceId}/logs-configs | 
[**V1InstancesLogsConfigsUpdate**](LogsApi.md#V1InstancesLogsConfigsUpdate) | **Put** /v1/instances/{instanceId}/logs-configs | 
[**V1ProjectsInstancesLogsAlertgroupsCreate**](LogsApi.md#V1ProjectsInstancesLogsAlertgroupsCreate) | **Post** /v1/projects/{projectId}/instances/{instanceId}/logs-alertgroups | 
[**V1ProjectsInstancesLogsAlertgroupsDelete**](LogsApi.md#V1ProjectsInstancesLogsAlertgroupsDelete) | **Delete** /v1/projects/{projectId}/instances/{instanceId}/logs-alertgroups/{groupName} | 
[**V1ProjectsInstancesLogsAlertgroupsList**](LogsApi.md#V1ProjectsInstancesLogsAlertgroupsList) | **Get** /v1/projects/{projectId}/instances/{instanceId}/logs-alertgroups | 
[**V1ProjectsInstancesLogsAlertgroupsRead**](LogsApi.md#V1ProjectsInstancesLogsAlertgroupsRead) | **Get** /v1/projects/{projectId}/instances/{instanceId}/logs-alertgroups/{groupName} | 
[**V1ProjectsInstancesLogsAlertgroupsUpdate**](LogsApi.md#V1ProjectsInstancesLogsAlertgroupsUpdate) | **Put** /v1/projects/{projectId}/instances/{instanceId}/logs-alertgroups/{groupName} | 
[**V1ProjectsInstancesLogsConfigsList**](LogsApi.md#V1ProjectsInstancesLogsConfigsList) | **Get** /v1/projects/{projectId}/instances/{instanceId}/logs-configs | 
[**V1ProjectsInstancesLogsConfigsUpdate**](LogsApi.md#V1ProjectsInstancesLogsConfigsUpdate) | **Put** /v1/projects/{projectId}/instances/{instanceId}/logs-configs | 



## V1InstancesLogsAlertgroupsCreate

> AlertGroupsResponse V1InstancesLogsAlertgroupsCreate(ctx, instanceId).Authorization(authorization).Data(data).Execute()





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
    data := *openapiclient.NewV1InstancesAlertgroupsCreateRequest("Name_example", []openapiclient.V1InstancesAlertgroupsCreateRequestRulesInner{*openapiclient.NewV1InstancesAlertgroupsCreateRequestRulesInner("Alert_example", "Expr_example")}) // V1InstancesAlertgroupsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsApi.V1InstancesLogsAlertgroupsCreate(context.Background(), instanceId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.V1InstancesLogsAlertgroupsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesLogsAlertgroupsCreate`: AlertGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.V1InstancesLogsAlertgroupsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesLogsAlertgroupsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesAlertgroupsCreateRequest**](V1InstancesAlertgroupsCreateRequest.md) |  | 

### Return type

[**AlertGroupsResponse**](AlertGroupsResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesLogsAlertgroupsDelete

> AlertGroupsResponse V1InstancesLogsAlertgroupsDelete(ctx, groupName, instanceId).Authorization(authorization).Execute()





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
    groupName := "groupName_example" // string | 
    instanceId := "instanceId_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsApi.V1InstancesLogsAlertgroupsDelete(context.Background(), groupName, instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.V1InstancesLogsAlertgroupsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesLogsAlertgroupsDelete`: AlertGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.V1InstancesLogsAlertgroupsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesLogsAlertgroupsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**AlertGroupsResponse**](AlertGroupsResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesLogsAlertgroupsList

> AlertGroupsResponse V1InstancesLogsAlertgroupsList(ctx, instanceId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.LogsApi.V1InstancesLogsAlertgroupsList(context.Background(), instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.V1InstancesLogsAlertgroupsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesLogsAlertgroupsList`: AlertGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.V1InstancesLogsAlertgroupsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesLogsAlertgroupsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**AlertGroupsResponse**](AlertGroupsResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesLogsAlertgroupsRead

> AlertGroupResponse V1InstancesLogsAlertgroupsRead(ctx, groupName, instanceId).Authorization(authorization).Execute()





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
    groupName := "groupName_example" // string | 
    instanceId := "instanceId_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsApi.V1InstancesLogsAlertgroupsRead(context.Background(), groupName, instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.V1InstancesLogsAlertgroupsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesLogsAlertgroupsRead`: AlertGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.V1InstancesLogsAlertgroupsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesLogsAlertgroupsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**AlertGroupResponse**](AlertGroupResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesLogsAlertgroupsUpdate

> AlertGroupsResponse V1InstancesLogsAlertgroupsUpdate(ctx, groupName, instanceId).Authorization(authorization).Data(data).Execute()





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
    groupName := "groupName_example" // string | 
    instanceId := "instanceId_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.
    data := *openapiclient.NewV1InstancesAlertgroupsUpdateRequest([]openapiclient.V1InstancesAlertgroupsCreateRequestRulesInner{*openapiclient.NewV1InstancesAlertgroupsCreateRequestRulesInner("Alert_example", "Expr_example")}) // V1InstancesAlertgroupsUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsApi.V1InstancesLogsAlertgroupsUpdate(context.Background(), groupName, instanceId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.V1InstancesLogsAlertgroupsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesLogsAlertgroupsUpdate`: AlertGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.V1InstancesLogsAlertgroupsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesLogsAlertgroupsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesAlertgroupsUpdateRequest**](V1InstancesAlertgroupsUpdateRequest.md) |  | 

### Return type

[**AlertGroupsResponse**](AlertGroupsResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesLogsConfigsList

> LogsConfigResponse V1InstancesLogsConfigsList(ctx, instanceId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.LogsApi.V1InstancesLogsConfigsList(context.Background(), instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.V1InstancesLogsConfigsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesLogsConfigsList`: LogsConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.V1InstancesLogsConfigsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesLogsConfigsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**LogsConfigResponse**](LogsConfigResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesLogsConfigsUpdate

> Message V1InstancesLogsConfigsUpdate(ctx, instanceId).Authorization(authorization).Data(data).Execute()





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
    data := *openapiclient.NewV1InstancesLogsConfigsUpdateRequest("Retention_example") // V1InstancesLogsConfigsUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsApi.V1InstancesLogsConfigsUpdate(context.Background(), instanceId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.V1InstancesLogsConfigsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesLogsConfigsUpdate`: Message
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.V1InstancesLogsConfigsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesLogsConfigsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesLogsConfigsUpdateRequest**](V1InstancesLogsConfigsUpdateRequest.md) |  | 

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


## V1ProjectsInstancesLogsAlertgroupsCreate

> AlertGroupsResponse V1ProjectsInstancesLogsAlertgroupsCreate(ctx, instanceId, projectId).Authorization(authorization).Data(data).Execute()





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
    data := *openapiclient.NewV1InstancesAlertgroupsCreateRequest("Name_example", []openapiclient.V1InstancesAlertgroupsCreateRequestRulesInner{*openapiclient.NewV1InstancesAlertgroupsCreateRequestRulesInner("Alert_example", "Expr_example")}) // V1InstancesAlertgroupsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsApi.V1ProjectsInstancesLogsAlertgroupsCreate(context.Background(), instanceId, projectId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.V1ProjectsInstancesLogsAlertgroupsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesLogsAlertgroupsCreate`: AlertGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.V1ProjectsInstancesLogsAlertgroupsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesLogsAlertgroupsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesAlertgroupsCreateRequest**](V1InstancesAlertgroupsCreateRequest.md) |  | 

### Return type

[**AlertGroupsResponse**](AlertGroupsResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesLogsAlertgroupsDelete

> AlertGroupsResponse V1ProjectsInstancesLogsAlertgroupsDelete(ctx, groupName, instanceId, projectId).Authorization(authorization).Execute()





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
    groupName := "groupName_example" // string | 
    instanceId := "instanceId_example" // string | 
    projectId := "projectId_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsApi.V1ProjectsInstancesLogsAlertgroupsDelete(context.Background(), groupName, instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.V1ProjectsInstancesLogsAlertgroupsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesLogsAlertgroupsDelete`: AlertGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.V1ProjectsInstancesLogsAlertgroupsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** |  | 
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesLogsAlertgroupsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**AlertGroupsResponse**](AlertGroupsResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesLogsAlertgroupsList

> AlertGroupsResponse V1ProjectsInstancesLogsAlertgroupsList(ctx, instanceId, projectId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.LogsApi.V1ProjectsInstancesLogsAlertgroupsList(context.Background(), instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.V1ProjectsInstancesLogsAlertgroupsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesLogsAlertgroupsList`: AlertGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.V1ProjectsInstancesLogsAlertgroupsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesLogsAlertgroupsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**AlertGroupsResponse**](AlertGroupsResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesLogsAlertgroupsRead

> AlertGroupResponse V1ProjectsInstancesLogsAlertgroupsRead(ctx, groupName, instanceId, projectId).Authorization(authorization).Execute()





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
    groupName := "groupName_example" // string | 
    instanceId := "instanceId_example" // string | 
    projectId := "projectId_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsApi.V1ProjectsInstancesLogsAlertgroupsRead(context.Background(), groupName, instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.V1ProjectsInstancesLogsAlertgroupsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesLogsAlertgroupsRead`: AlertGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.V1ProjectsInstancesLogsAlertgroupsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** |  | 
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesLogsAlertgroupsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**AlertGroupResponse**](AlertGroupResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesLogsAlertgroupsUpdate

> AlertGroupsResponse V1ProjectsInstancesLogsAlertgroupsUpdate(ctx, groupName, instanceId, projectId).Authorization(authorization).Data(data).Execute()





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
    groupName := "groupName_example" // string | 
    instanceId := "instanceId_example" // string | 
    projectId := "projectId_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.
    data := *openapiclient.NewV1InstancesAlertgroupsUpdateRequest([]openapiclient.V1InstancesAlertgroupsCreateRequestRulesInner{*openapiclient.NewV1InstancesAlertgroupsCreateRequestRulesInner("Alert_example", "Expr_example")}) // V1InstancesAlertgroupsUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsApi.V1ProjectsInstancesLogsAlertgroupsUpdate(context.Background(), groupName, instanceId, projectId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.V1ProjectsInstancesLogsAlertgroupsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesLogsAlertgroupsUpdate`: AlertGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.V1ProjectsInstancesLogsAlertgroupsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** |  | 
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesLogsAlertgroupsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesAlertgroupsUpdateRequest**](V1InstancesAlertgroupsUpdateRequest.md) |  | 

### Return type

[**AlertGroupsResponse**](AlertGroupsResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesLogsConfigsList

> LogsConfigResponse V1ProjectsInstancesLogsConfigsList(ctx, instanceId, projectId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.LogsApi.V1ProjectsInstancesLogsConfigsList(context.Background(), instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.V1ProjectsInstancesLogsConfigsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesLogsConfigsList`: LogsConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.V1ProjectsInstancesLogsConfigsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesLogsConfigsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**LogsConfigResponse**](LogsConfigResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesLogsConfigsUpdate

> Message V1ProjectsInstancesLogsConfigsUpdate(ctx, instanceId, projectId).Authorization(authorization).Data(data).Execute()





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
    data := *openapiclient.NewV1InstancesLogsConfigsUpdateRequest("Retention_example") // V1InstancesLogsConfigsUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsApi.V1ProjectsInstancesLogsConfigsUpdate(context.Background(), instanceId, projectId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.V1ProjectsInstancesLogsConfigsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesLogsConfigsUpdate`: Message
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.V1ProjectsInstancesLogsConfigsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesLogsConfigsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesLogsConfigsUpdateRequest**](V1InstancesLogsConfigsUpdateRequest.md) |  | 

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

