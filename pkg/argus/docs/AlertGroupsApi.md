# \AlertGroupsApi

All URIs are relative to *https://api.argus.eu01.stackit.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1InstancesAlertgroupsCreate**](AlertGroupsApi.md#V1InstancesAlertgroupsCreate) | **Post** /v1/instances/{instanceId}/alertgroups | 
[**V1InstancesAlertgroupsDelete**](AlertGroupsApi.md#V1InstancesAlertgroupsDelete) | **Delete** /v1/instances/{instanceId}/alertgroups | 
[**V1InstancesAlertgroupsDelete_0**](AlertGroupsApi.md#V1InstancesAlertgroupsDelete_0) | **Delete** /v1/instances/{instanceId}/alertgroups/{groupName} | 
[**V1InstancesAlertgroupsList**](AlertGroupsApi.md#V1InstancesAlertgroupsList) | **Get** /v1/instances/{instanceId}/alertgroups | 
[**V1InstancesAlertgroupsPartialUpdate**](AlertGroupsApi.md#V1InstancesAlertgroupsPartialUpdate) | **Patch** /v1/instances/{instanceId}/alertgroups | 
[**V1InstancesAlertgroupsRead**](AlertGroupsApi.md#V1InstancesAlertgroupsRead) | **Get** /v1/instances/{instanceId}/alertgroups/{groupName} | 
[**V1InstancesAlertgroupsUpdate**](AlertGroupsApi.md#V1InstancesAlertgroupsUpdate) | **Put** /v1/instances/{instanceId}/alertgroups/{groupName} | 
[**V1ProjectsInstancesAlertgroupsCreate**](AlertGroupsApi.md#V1ProjectsInstancesAlertgroupsCreate) | **Post** /v1/projects/{projectId}/instances/{instanceId}/alertgroups | 
[**V1ProjectsInstancesAlertgroupsDelete**](AlertGroupsApi.md#V1ProjectsInstancesAlertgroupsDelete) | **Delete** /v1/projects/{projectId}/instances/{instanceId}/alertgroups | 
[**V1ProjectsInstancesAlertgroupsDelete_0**](AlertGroupsApi.md#V1ProjectsInstancesAlertgroupsDelete_0) | **Delete** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName} | 
[**V1ProjectsInstancesAlertgroupsList**](AlertGroupsApi.md#V1ProjectsInstancesAlertgroupsList) | **Get** /v1/projects/{projectId}/instances/{instanceId}/alertgroups | 
[**V1ProjectsInstancesAlertgroupsPartialUpdate**](AlertGroupsApi.md#V1ProjectsInstancesAlertgroupsPartialUpdate) | **Patch** /v1/projects/{projectId}/instances/{instanceId}/alertgroups | 
[**V1ProjectsInstancesAlertgroupsRead**](AlertGroupsApi.md#V1ProjectsInstancesAlertgroupsRead) | **Get** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName} | 
[**V1ProjectsInstancesAlertgroupsUpdate**](AlertGroupsApi.md#V1ProjectsInstancesAlertgroupsUpdate) | **Put** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName} | 



## V1InstancesAlertgroupsCreate

> AlertGroupsResponse V1InstancesAlertgroupsCreate(ctx, instanceId).Authorization(authorization).Data(data).Execute()





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
    resp, r, err := apiClient.AlertGroupsApi.V1InstancesAlertgroupsCreate(context.Background(), instanceId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1InstancesAlertgroupsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsCreate`: AlertGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertGroupsApi.V1InstancesAlertgroupsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertgroupsCreateRequest struct via the builder pattern


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


## V1InstancesAlertgroupsDelete

> AlertGroupsResponse V1InstancesAlertgroupsDelete(ctx, instanceId).Authorization(authorization).GroupName(groupName).Execute()





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
    groupName := []string{"Inner_example"} // []string | Name of the groups that should be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertGroupsApi.V1InstancesAlertgroupsDelete(context.Background(), instanceId).Authorization(authorization).GroupName(groupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1InstancesAlertgroupsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsDelete`: AlertGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertGroupsApi.V1InstancesAlertgroupsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertgroupsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **groupName** | **[]string** | Name of the groups that should be deleted | 

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


## V1InstancesAlertgroupsDelete_0

> AlertGroupsResponse V1InstancesAlertgroupsDelete_0(ctx, groupName, instanceId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.AlertGroupsApi.V1InstancesAlertgroupsDelete_0(context.Background(), groupName, instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1InstancesAlertgroupsDelete_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsDelete_0`: AlertGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertGroupsApi.V1InstancesAlertgroupsDelete_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertgroupsDelete_1Request struct via the builder pattern


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


## V1InstancesAlertgroupsList

> AlertGroupsResponse V1InstancesAlertgroupsList(ctx, instanceId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.AlertGroupsApi.V1InstancesAlertgroupsList(context.Background(), instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1InstancesAlertgroupsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsList`: AlertGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertGroupsApi.V1InstancesAlertgroupsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertgroupsListRequest struct via the builder pattern


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


## V1InstancesAlertgroupsPartialUpdate

> AlertGroupsResponse V1InstancesAlertgroupsPartialUpdate(ctx, instanceId).Authorization(authorization).Data(data).Execute()





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
    data := []openapiclient.V1InstancesAlertgroupsCreateRequest{*openapiclient.NewV1InstancesAlertgroupsCreateRequest("Name_example", []openapiclient.V1InstancesAlertgroupsCreateRequestRulesInner{*openapiclient.NewV1InstancesAlertgroupsCreateRequestRulesInner("Alert_example", "Expr_example")})} // []V1InstancesAlertgroupsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertGroupsApi.V1InstancesAlertgroupsPartialUpdate(context.Background(), instanceId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1InstancesAlertgroupsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsPartialUpdate`: AlertGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertGroupsApi.V1InstancesAlertgroupsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertgroupsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**[]V1InstancesAlertgroupsCreateRequest**](V1InstancesAlertgroupsCreateRequest.md) |  | 

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


## V1InstancesAlertgroupsRead

> AlertGroupResponse V1InstancesAlertgroupsRead(ctx, groupName, instanceId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.AlertGroupsApi.V1InstancesAlertgroupsRead(context.Background(), groupName, instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1InstancesAlertgroupsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsRead`: AlertGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertGroupsApi.V1InstancesAlertgroupsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertgroupsReadRequest struct via the builder pattern


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


## V1InstancesAlertgroupsUpdate

> AlertGroupsResponse V1InstancesAlertgroupsUpdate(ctx, groupName, instanceId).Authorization(authorization).Data(data).Execute()





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
    resp, r, err := apiClient.AlertGroupsApi.V1InstancesAlertgroupsUpdate(context.Background(), groupName, instanceId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1InstancesAlertgroupsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsUpdate`: AlertGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertGroupsApi.V1InstancesAlertgroupsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertgroupsUpdateRequest struct via the builder pattern


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


## V1ProjectsInstancesAlertgroupsCreate

> AlertGroupsResponse V1ProjectsInstancesAlertgroupsCreate(ctx, instanceId, projectId).Authorization(authorization).Data(data).Execute()





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
    resp, r, err := apiClient.AlertGroupsApi.V1ProjectsInstancesAlertgroupsCreate(context.Background(), instanceId, projectId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1ProjectsInstancesAlertgroupsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsCreate`: AlertGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertGroupsApi.V1ProjectsInstancesAlertgroupsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertgroupsCreateRequest struct via the builder pattern


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


## V1ProjectsInstancesAlertgroupsDelete

> AlertGroupsResponse V1ProjectsInstancesAlertgroupsDelete(ctx, instanceId, projectId).Authorization(authorization).GroupName(groupName).Execute()





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
    groupName := []string{"Inner_example"} // []string | Name of the groups that should be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertGroupsApi.V1ProjectsInstancesAlertgroupsDelete(context.Background(), instanceId, projectId).Authorization(authorization).GroupName(groupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1ProjectsInstancesAlertgroupsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsDelete`: AlertGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertGroupsApi.V1ProjectsInstancesAlertgroupsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertgroupsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **groupName** | **[]string** | Name of the groups that should be deleted | 

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


## V1ProjectsInstancesAlertgroupsDelete_0

> AlertGroupsResponse V1ProjectsInstancesAlertgroupsDelete_0(ctx, groupName, instanceId, projectId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.AlertGroupsApi.V1ProjectsInstancesAlertgroupsDelete_0(context.Background(), groupName, instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1ProjectsInstancesAlertgroupsDelete_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsDelete_0`: AlertGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertGroupsApi.V1ProjectsInstancesAlertgroupsDelete_0`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertgroupsDelete_2Request struct via the builder pattern


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


## V1ProjectsInstancesAlertgroupsList

> AlertGroupsResponse V1ProjectsInstancesAlertgroupsList(ctx, instanceId, projectId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.AlertGroupsApi.V1ProjectsInstancesAlertgroupsList(context.Background(), instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1ProjectsInstancesAlertgroupsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsList`: AlertGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertGroupsApi.V1ProjectsInstancesAlertgroupsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertgroupsListRequest struct via the builder pattern


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


## V1ProjectsInstancesAlertgroupsPartialUpdate

> AlertGroupsResponse V1ProjectsInstancesAlertgroupsPartialUpdate(ctx, instanceId, projectId).Authorization(authorization).Data(data).Execute()





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
    data := []openapiclient.V1InstancesAlertgroupsCreateRequest{*openapiclient.NewV1InstancesAlertgroupsCreateRequest("Name_example", []openapiclient.V1InstancesAlertgroupsCreateRequestRulesInner{*openapiclient.NewV1InstancesAlertgroupsCreateRequestRulesInner("Alert_example", "Expr_example")})} // []V1InstancesAlertgroupsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertGroupsApi.V1ProjectsInstancesAlertgroupsPartialUpdate(context.Background(), instanceId, projectId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1ProjectsInstancesAlertgroupsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsPartialUpdate`: AlertGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertGroupsApi.V1ProjectsInstancesAlertgroupsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertgroupsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**[]V1InstancesAlertgroupsCreateRequest**](V1InstancesAlertgroupsCreateRequest.md) |  | 

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


## V1ProjectsInstancesAlertgroupsRead

> AlertGroupResponse V1ProjectsInstancesAlertgroupsRead(ctx, groupName, instanceId, projectId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.AlertGroupsApi.V1ProjectsInstancesAlertgroupsRead(context.Background(), groupName, instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1ProjectsInstancesAlertgroupsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsRead`: AlertGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertGroupsApi.V1ProjectsInstancesAlertgroupsRead`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertgroupsReadRequest struct via the builder pattern


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


## V1ProjectsInstancesAlertgroupsUpdate

> AlertGroupsResponse V1ProjectsInstancesAlertgroupsUpdate(ctx, groupName, instanceId, projectId).Authorization(authorization).Data(data).Execute()





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
    resp, r, err := apiClient.AlertGroupsApi.V1ProjectsInstancesAlertgroupsUpdate(context.Background(), groupName, instanceId, projectId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1ProjectsInstancesAlertgroupsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsUpdate`: AlertGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertGroupsApi.V1ProjectsInstancesAlertgroupsUpdate`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertgroupsUpdateRequest struct via the builder pattern


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

