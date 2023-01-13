# \AlertGroupsApi

All URIs are relative to *https://api.argus.eu01.stackit.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1InstancesAlertgroupsAllDelete**](AlertGroupsApi.md#V1InstancesAlertgroupsAllDelete) | **Delete** /v1/instances/{instanceId}/alertgroups | 
[**V1InstancesAlertgroupsCreate**](AlertGroupsApi.md#V1InstancesAlertgroupsCreate) | **Post** /v1/instances/{instanceId}/alertgroups | 
[**V1InstancesAlertgroupsDelete**](AlertGroupsApi.md#V1InstancesAlertgroupsDelete) | **Delete** /v1/instances/{instanceId}/alertgroups/{groupName} | 
[**V1InstancesAlertgroupsList**](AlertGroupsApi.md#V1InstancesAlertgroupsList) | **Get** /v1/instances/{instanceId}/alertgroups | 
[**V1InstancesAlertgroupsPartialUpdate**](AlertGroupsApi.md#V1InstancesAlertgroupsPartialUpdate) | **Patch** /v1/instances/{instanceId}/alertgroups | 
[**V1InstancesAlertgroupsRead**](AlertGroupsApi.md#V1InstancesAlertgroupsRead) | **Get** /v1/instances/{instanceId}/alertgroups/{groupName} | 
[**V1InstancesAlertgroupsUpdate**](AlertGroupsApi.md#V1InstancesAlertgroupsUpdate) | **Put** /v1/instances/{instanceId}/alertgroups/{groupName} | 
[**V1ProjectsInstancesAlertgroupsAllDelete**](AlertGroupsApi.md#V1ProjectsInstancesAlertgroupsAllDelete) | **Delete** /v1/projects/{projectId}/instances/{instanceId}/alertgroups | 
[**V1ProjectsInstancesAlertgroupsCreate**](AlertGroupsApi.md#V1ProjectsInstancesAlertgroupsCreate) | **Post** /v1/projects/{projectId}/instances/{instanceId}/alertgroups | 
[**V1ProjectsInstancesAlertgroupsDelete**](AlertGroupsApi.md#V1ProjectsInstancesAlertgroupsDelete) | **Delete** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName} | 
[**V1ProjectsInstancesAlertgroupsList**](AlertGroupsApi.md#V1ProjectsInstancesAlertgroupsList) | **Get** /v1/projects/{projectId}/instances/{instanceId}/alertgroups | 
[**V1ProjectsInstancesAlertgroupsPartialUpdate**](AlertGroupsApi.md#V1ProjectsInstancesAlertgroupsPartialUpdate) | **Patch** /v1/projects/{projectId}/instances/{instanceId}/alertgroups | 
[**V1ProjectsInstancesAlertgroupsRead**](AlertGroupsApi.md#V1ProjectsInstancesAlertgroupsRead) | **Get** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName} | 
[**V1ProjectsInstancesAlertgroupsUpdate**](AlertGroupsApi.md#V1ProjectsInstancesAlertgroupsUpdate) | **Put** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName} | 



## V1InstancesAlertgroupsAllDelete

> PostAlertGroup V1InstancesAlertgroupsAllDelete(ctx, instanceId).Authorization(authorization).GroupName(groupName).Execute()





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
    groupName := []string{"Inner_example"} // []string | Name of the groups that should be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertGroupsApi.V1InstancesAlertgroupsAllDelete(context.Background(), instanceId).Authorization(authorization).GroupName(groupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1InstancesAlertgroupsAllDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsAllDelete`: PostAlertGroup
    fmt.Fprintf(os.Stdout, "Response from `AlertGroupsApi.V1InstancesAlertgroupsAllDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertgroupsAllDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 
 **groupName** | **[]string** | Name of the groups that should be deleted | 

### Return type

[**PostAlertGroup**](PostAlertGroup.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertgroupsCreate

> PostAlertGroup V1InstancesAlertgroupsCreate(ctx, instanceId).Authorization(authorization).V1InstancesAlertgroupsCreateRequest(v1InstancesAlertgroupsCreateRequest).Execute()





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
    v1InstancesAlertgroupsCreateRequest := *openapiclient.NewV1InstancesAlertgroupsCreateRequest("Name_example", "Interval_example", []openapiclient.V1InstancesAlertgroupsCreateRequestRulesInner{*openapiclient.NewV1InstancesAlertgroupsCreateRequestRulesInner("Alert_example", "Expr_example", map[string]interface{}(123), map[string]interface{}(123))}) // V1InstancesAlertgroupsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertGroupsApi.V1InstancesAlertgroupsCreate(context.Background(), instanceId).Authorization(authorization).V1InstancesAlertgroupsCreateRequest(v1InstancesAlertgroupsCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1InstancesAlertgroupsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsCreate`: PostAlertGroup
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

 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 
 **v1InstancesAlertgroupsCreateRequest** | [**V1InstancesAlertgroupsCreateRequest**](V1InstancesAlertgroupsCreateRequest.md) |  | 

### Return type

[**PostAlertGroup**](PostAlertGroup.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertgroupsDelete

> DeleteAlertGroup V1InstancesAlertgroupsDelete(ctx, groupName, instanceId).Authorization(authorization).Execute()





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
    authorization := "authorization_example" // string | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertGroupsApi.V1InstancesAlertgroupsDelete(context.Background(), groupName, instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1InstancesAlertgroupsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsDelete`: DeleteAlertGroup
    fmt.Fprintf(os.Stdout, "Response from `AlertGroupsApi.V1InstancesAlertgroupsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertgroupsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 

### Return type

[**DeleteAlertGroup**](DeleteAlertGroup.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertgroupsList

> GetAllAlertGroups V1InstancesAlertgroupsList(ctx, instanceId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.AlertGroupsApi.V1InstancesAlertgroupsList(context.Background(), instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1InstancesAlertgroupsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsList`: GetAllAlertGroups
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

 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 

### Return type

[**GetAllAlertGroups**](GetAllAlertGroups.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertgroupsPartialUpdate

> PostAlertGroup V1InstancesAlertgroupsPartialUpdate(ctx, instanceId).Authorization(authorization).V1InstancesAlertgroupsCreateRequest(v1InstancesAlertgroupsCreateRequest).Execute()





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
    v1InstancesAlertgroupsCreateRequest := []openapiclient.V1InstancesAlertgroupsCreateRequest{*openapiclient.NewV1InstancesAlertgroupsCreateRequest("Name_example", "Interval_example", []openapiclient.V1InstancesAlertgroupsCreateRequestRulesInner{*openapiclient.NewV1InstancesAlertgroupsCreateRequestRulesInner("Alert_example", "Expr_example", map[string]interface{}(123), map[string]interface{}(123))})} // []V1InstancesAlertgroupsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertGroupsApi.V1InstancesAlertgroupsPartialUpdate(context.Background(), instanceId).Authorization(authorization).V1InstancesAlertgroupsCreateRequest(v1InstancesAlertgroupsCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1InstancesAlertgroupsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsPartialUpdate`: PostAlertGroup
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

 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 
 **v1InstancesAlertgroupsCreateRequest** | [**[]V1InstancesAlertgroupsCreateRequest**](V1InstancesAlertgroupsCreateRequest.md) |  | 

### Return type

[**PostAlertGroup**](PostAlertGroup.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertgroupsRead

> GetAlertGroup V1InstancesAlertgroupsRead(ctx, groupName, instanceId).Authorization(authorization).Execute()





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
    authorization := "authorization_example" // string | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertGroupsApi.V1InstancesAlertgroupsRead(context.Background(), groupName, instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1InstancesAlertgroupsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsRead`: GetAlertGroup
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


 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 

### Return type

[**GetAlertGroup**](GetAlertGroup.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertgroupsUpdate

> PutAlertGroup V1InstancesAlertgroupsUpdate(ctx, groupName, instanceId).Authorization(authorization).V1InstancesAlertgroupsCreateRequest(v1InstancesAlertgroupsCreateRequest).Execute()





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
    authorization := "authorization_example" // string | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user
    v1InstancesAlertgroupsCreateRequest := *openapiclient.NewV1InstancesAlertgroupsCreateRequest("Name_example", "Interval_example", []openapiclient.V1InstancesAlertgroupsCreateRequestRulesInner{*openapiclient.NewV1InstancesAlertgroupsCreateRequestRulesInner("Alert_example", "Expr_example", map[string]interface{}(123), map[string]interface{}(123))}) // V1InstancesAlertgroupsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertGroupsApi.V1InstancesAlertgroupsUpdate(context.Background(), groupName, instanceId).Authorization(authorization).V1InstancesAlertgroupsCreateRequest(v1InstancesAlertgroupsCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1InstancesAlertgroupsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsUpdate`: PutAlertGroup
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


 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 
 **v1InstancesAlertgroupsCreateRequest** | [**V1InstancesAlertgroupsCreateRequest**](V1InstancesAlertgroupsCreateRequest.md) |  | 

### Return type

[**PutAlertGroup**](PutAlertGroup.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsAllDelete

> PostAlertGroup V1ProjectsInstancesAlertgroupsAllDelete(ctx, instanceId, projectId).Authorization(authorization).GroupName(groupName).Execute()





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
    groupName := []string{"Inner_example"} // []string | Name of the groups that should be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertGroupsApi.V1ProjectsInstancesAlertgroupsAllDelete(context.Background(), instanceId, projectId).Authorization(authorization).GroupName(groupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1ProjectsInstancesAlertgroupsAllDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsAllDelete`: PostAlertGroup
    fmt.Fprintf(os.Stdout, "Response from `AlertGroupsApi.V1ProjectsInstancesAlertgroupsAllDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertgroupsAllDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 
 **groupName** | **[]string** | Name of the groups that should be deleted | 

### Return type

[**PostAlertGroup**](PostAlertGroup.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsCreate

> PostAlertGroup V1ProjectsInstancesAlertgroupsCreate(ctx, instanceId, projectId).Authorization(authorization).V1InstancesAlertgroupsCreateRequest(v1InstancesAlertgroupsCreateRequest).Execute()





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
    v1InstancesAlertgroupsCreateRequest := *openapiclient.NewV1InstancesAlertgroupsCreateRequest("Name_example", "Interval_example", []openapiclient.V1InstancesAlertgroupsCreateRequestRulesInner{*openapiclient.NewV1InstancesAlertgroupsCreateRequestRulesInner("Alert_example", "Expr_example", map[string]interface{}(123), map[string]interface{}(123))}) // V1InstancesAlertgroupsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertGroupsApi.V1ProjectsInstancesAlertgroupsCreate(context.Background(), instanceId, projectId).Authorization(authorization).V1InstancesAlertgroupsCreateRequest(v1InstancesAlertgroupsCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1ProjectsInstancesAlertgroupsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsCreate`: PostAlertGroup
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


 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 
 **v1InstancesAlertgroupsCreateRequest** | [**V1InstancesAlertgroupsCreateRequest**](V1InstancesAlertgroupsCreateRequest.md) |  | 

### Return type

[**PostAlertGroup**](PostAlertGroup.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsDelete

> DeleteAlertGroup V1ProjectsInstancesAlertgroupsDelete(ctx, groupName, instanceId, projectId).Authorization(authorization).Execute()





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
    authorization := "authorization_example" // string | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertGroupsApi.V1ProjectsInstancesAlertgroupsDelete(context.Background(), groupName, instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1ProjectsInstancesAlertgroupsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsDelete`: DeleteAlertGroup
    fmt.Fprintf(os.Stdout, "Response from `AlertGroupsApi.V1ProjectsInstancesAlertgroupsDelete`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertgroupsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 

### Return type

[**DeleteAlertGroup**](DeleteAlertGroup.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsList

> GetAllAlertGroups V1ProjectsInstancesAlertgroupsList(ctx, instanceId, projectId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.AlertGroupsApi.V1ProjectsInstancesAlertgroupsList(context.Background(), instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1ProjectsInstancesAlertgroupsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsList`: GetAllAlertGroups
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


 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 

### Return type

[**GetAllAlertGroups**](GetAllAlertGroups.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsPartialUpdate

> PostAlertGroup V1ProjectsInstancesAlertgroupsPartialUpdate(ctx, instanceId, projectId).Authorization(authorization).V1InstancesAlertgroupsCreateRequest(v1InstancesAlertgroupsCreateRequest).Execute()





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
    v1InstancesAlertgroupsCreateRequest := []openapiclient.V1InstancesAlertgroupsCreateRequest{*openapiclient.NewV1InstancesAlertgroupsCreateRequest("Name_example", "Interval_example", []openapiclient.V1InstancesAlertgroupsCreateRequestRulesInner{*openapiclient.NewV1InstancesAlertgroupsCreateRequestRulesInner("Alert_example", "Expr_example", map[string]interface{}(123), map[string]interface{}(123))})} // []V1InstancesAlertgroupsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertGroupsApi.V1ProjectsInstancesAlertgroupsPartialUpdate(context.Background(), instanceId, projectId).Authorization(authorization).V1InstancesAlertgroupsCreateRequest(v1InstancesAlertgroupsCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1ProjectsInstancesAlertgroupsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsPartialUpdate`: PostAlertGroup
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


 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 
 **v1InstancesAlertgroupsCreateRequest** | [**[]V1InstancesAlertgroupsCreateRequest**](V1InstancesAlertgroupsCreateRequest.md) |  | 

### Return type

[**PostAlertGroup**](PostAlertGroup.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsRead

> GetAlertGroup V1ProjectsInstancesAlertgroupsRead(ctx, groupName, instanceId, projectId).Authorization(authorization).Execute()





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
    authorization := "authorization_example" // string | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertGroupsApi.V1ProjectsInstancesAlertgroupsRead(context.Background(), groupName, instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1ProjectsInstancesAlertgroupsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsRead`: GetAlertGroup
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



 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 

### Return type

[**GetAlertGroup**](GetAlertGroup.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsUpdate

> PutAlertGroup V1ProjectsInstancesAlertgroupsUpdate(ctx, groupName, instanceId, projectId).Authorization(authorization).V1InstancesAlertgroupsCreateRequest(v1InstancesAlertgroupsCreateRequest).Execute()





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
    authorization := "authorization_example" // string | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user
    v1InstancesAlertgroupsCreateRequest := *openapiclient.NewV1InstancesAlertgroupsCreateRequest("Name_example", "Interval_example", []openapiclient.V1InstancesAlertgroupsCreateRequestRulesInner{*openapiclient.NewV1InstancesAlertgroupsCreateRequestRulesInner("Alert_example", "Expr_example", map[string]interface{}(123), map[string]interface{}(123))}) // V1InstancesAlertgroupsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertGroupsApi.V1ProjectsInstancesAlertgroupsUpdate(context.Background(), groupName, instanceId, projectId).Authorization(authorization).V1InstancesAlertgroupsCreateRequest(v1InstancesAlertgroupsCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertGroupsApi.V1ProjectsInstancesAlertgroupsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsUpdate`: PutAlertGroup
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



 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 
 **v1InstancesAlertgroupsCreateRequest** | [**V1InstancesAlertgroupsCreateRequest**](V1InstancesAlertgroupsCreateRequest.md) |  | 

### Return type

[**PutAlertGroup**](PutAlertGroup.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

