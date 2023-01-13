# \AlertRulesApi

All URIs are relative to *https://api.argus.eu01.stackit.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1InstancesAlertgroupsAlertrulesAllDelete**](AlertRulesApi.md#V1InstancesAlertgroupsAlertrulesAllDelete) | **Delete** /v1/instances/{instanceId}/alertgroups/{groupName}/alertrules | 
[**V1InstancesAlertgroupsAlertrulesCreate**](AlertRulesApi.md#V1InstancesAlertgroupsAlertrulesCreate) | **Post** /v1/instances/{instanceId}/alertgroups/{groupName}/alertrules | 
[**V1InstancesAlertgroupsAlertrulesDelete**](AlertRulesApi.md#V1InstancesAlertgroupsAlertrulesDelete) | **Delete** /v1/instances/{instanceId}/alertgroups/{groupName}/alertrules/{alertName} | 
[**V1InstancesAlertgroupsAlertrulesList**](AlertRulesApi.md#V1InstancesAlertgroupsAlertrulesList) | **Get** /v1/instances/{instanceId}/alertgroups/{groupName}/alertrules | 
[**V1InstancesAlertgroupsAlertrulesPartialUpdate**](AlertRulesApi.md#V1InstancesAlertgroupsAlertrulesPartialUpdate) | **Patch** /v1/instances/{instanceId}/alertgroups/{groupName}/alertrules | 
[**V1InstancesAlertgroupsAlertrulesRead**](AlertRulesApi.md#V1InstancesAlertgroupsAlertrulesRead) | **Get** /v1/instances/{instanceId}/alertgroups/{groupName}/alertrules/{alertName} | 
[**V1InstancesAlertgroupsAlertrulesUpdate**](AlertRulesApi.md#V1InstancesAlertgroupsAlertrulesUpdate) | **Put** /v1/instances/{instanceId}/alertgroups/{groupName}/alertrules/{alertName} | 
[**V1ProjectsInstancesAlertgroupsAlertrulesAllDelete**](AlertRulesApi.md#V1ProjectsInstancesAlertgroupsAlertrulesAllDelete) | **Delete** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/alertrules | 
[**V1ProjectsInstancesAlertgroupsAlertrulesCreate**](AlertRulesApi.md#V1ProjectsInstancesAlertgroupsAlertrulesCreate) | **Post** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/alertrules | 
[**V1ProjectsInstancesAlertgroupsAlertrulesDelete**](AlertRulesApi.md#V1ProjectsInstancesAlertgroupsAlertrulesDelete) | **Delete** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/alertrules/{alertName} | 
[**V1ProjectsInstancesAlertgroupsAlertrulesList**](AlertRulesApi.md#V1ProjectsInstancesAlertgroupsAlertrulesList) | **Get** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/alertrules | 
[**V1ProjectsInstancesAlertgroupsAlertrulesPartialUpdate**](AlertRulesApi.md#V1ProjectsInstancesAlertgroupsAlertrulesPartialUpdate) | **Patch** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/alertrules | 
[**V1ProjectsInstancesAlertgroupsAlertrulesRead**](AlertRulesApi.md#V1ProjectsInstancesAlertgroupsAlertrulesRead) | **Get** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/alertrules/{alertName} | 
[**V1ProjectsInstancesAlertgroupsAlertrulesUpdate**](AlertRulesApi.md#V1ProjectsInstancesAlertgroupsAlertrulesUpdate) | **Put** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/alertrules/{alertName} | 



## V1InstancesAlertgroupsAlertrulesAllDelete

> PostAlertRule V1InstancesAlertgroupsAlertrulesAllDelete(ctx, groupName, instanceId).Authorization(authorization).AlertName(alertName).Execute()





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
    alertName := []string{"Inner_example"} // []string | Name of the alert rules that should be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRulesApi.V1InstancesAlertgroupsAlertrulesAllDelete(context.Background(), groupName, instanceId).Authorization(authorization).AlertName(alertName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1InstancesAlertgroupsAlertrulesAllDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsAlertrulesAllDelete`: PostAlertRule
    fmt.Fprintf(os.Stdout, "Response from `AlertRulesApi.V1InstancesAlertgroupsAlertrulesAllDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertgroupsAlertrulesAllDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 
 **alertName** | **[]string** | Name of the alert rules that should be deleted | 

### Return type

[**PostAlertRule**](PostAlertRule.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertgroupsAlertrulesCreate

> PostAlertRule V1InstancesAlertgroupsAlertrulesCreate(ctx, groupName, instanceId).Authorization(authorization).V1InstancesAlertgroupsAlertrulesCreateRequest(v1InstancesAlertgroupsAlertrulesCreateRequest).Execute()





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
    v1InstancesAlertgroupsAlertrulesCreateRequest := *openapiclient.NewV1InstancesAlertgroupsAlertrulesCreateRequest("Alert_example", "Expr_example", map[string]interface{}(123), map[string]interface{}(123)) // V1InstancesAlertgroupsAlertrulesCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRulesApi.V1InstancesAlertgroupsAlertrulesCreate(context.Background(), groupName, instanceId).Authorization(authorization).V1InstancesAlertgroupsAlertrulesCreateRequest(v1InstancesAlertgroupsAlertrulesCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1InstancesAlertgroupsAlertrulesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsAlertrulesCreate`: PostAlertRule
    fmt.Fprintf(os.Stdout, "Response from `AlertRulesApi.V1InstancesAlertgroupsAlertrulesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertgroupsAlertrulesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 
 **v1InstancesAlertgroupsAlertrulesCreateRequest** | [**V1InstancesAlertgroupsAlertrulesCreateRequest**](V1InstancesAlertgroupsAlertrulesCreateRequest.md) |  | 

### Return type

[**PostAlertRule**](PostAlertRule.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertgroupsAlertrulesDelete

> DeleteAlertRule V1InstancesAlertgroupsAlertrulesDelete(ctx, alertName, groupName, instanceId).Authorization(authorization).Execute()





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
    alertName := "alertName_example" // string | 
    groupName := "groupName_example" // string | 
    instanceId := "instanceId_example" // string | 
    authorization := "authorization_example" // string | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRulesApi.V1InstancesAlertgroupsAlertrulesDelete(context.Background(), alertName, groupName, instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1InstancesAlertgroupsAlertrulesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsAlertrulesDelete`: DeleteAlertRule
    fmt.Fprintf(os.Stdout, "Response from `AlertRulesApi.V1InstancesAlertgroupsAlertrulesDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertName** | **string** |  | 
**groupName** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertgroupsAlertrulesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 

### Return type

[**DeleteAlertRule**](DeleteAlertRule.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertgroupsAlertrulesList

> GetAllAlertRules V1InstancesAlertgroupsAlertrulesList(ctx, groupName, instanceId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.AlertRulesApi.V1InstancesAlertgroupsAlertrulesList(context.Background(), groupName, instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1InstancesAlertgroupsAlertrulesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsAlertrulesList`: GetAllAlertRules
    fmt.Fprintf(os.Stdout, "Response from `AlertRulesApi.V1InstancesAlertgroupsAlertrulesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertgroupsAlertrulesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 

### Return type

[**GetAllAlertRules**](GetAllAlertRules.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertgroupsAlertrulesPartialUpdate

> PostAlertRule V1InstancesAlertgroupsAlertrulesPartialUpdate(ctx, groupName, instanceId).Authorization(authorization).V1InstancesAlertgroupsAlertrulesCreateRequest(v1InstancesAlertgroupsAlertrulesCreateRequest).Execute()





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
    v1InstancesAlertgroupsAlertrulesCreateRequest := []openapiclient.V1InstancesAlertgroupsAlertrulesCreateRequest{*openapiclient.NewV1InstancesAlertgroupsAlertrulesCreateRequest("Alert_example", "Expr_example", map[string]interface{}(123), map[string]interface{}(123))} // []V1InstancesAlertgroupsAlertrulesCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRulesApi.V1InstancesAlertgroupsAlertrulesPartialUpdate(context.Background(), groupName, instanceId).Authorization(authorization).V1InstancesAlertgroupsAlertrulesCreateRequest(v1InstancesAlertgroupsAlertrulesCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1InstancesAlertgroupsAlertrulesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsAlertrulesPartialUpdate`: PostAlertRule
    fmt.Fprintf(os.Stdout, "Response from `AlertRulesApi.V1InstancesAlertgroupsAlertrulesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertgroupsAlertrulesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 
 **v1InstancesAlertgroupsAlertrulesCreateRequest** | [**[]V1InstancesAlertgroupsAlertrulesCreateRequest**](V1InstancesAlertgroupsAlertrulesCreateRequest.md) |  | 

### Return type

[**PostAlertRule**](PostAlertRule.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertgroupsAlertrulesRead

> GetAlertRule V1InstancesAlertgroupsAlertrulesRead(ctx, alertName, groupName, instanceId).Authorization(authorization).Execute()





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
    alertName := "alertName_example" // string | 
    groupName := "groupName_example" // string | 
    instanceId := "instanceId_example" // string | 
    authorization := "authorization_example" // string | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRulesApi.V1InstancesAlertgroupsAlertrulesRead(context.Background(), alertName, groupName, instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1InstancesAlertgroupsAlertrulesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsAlertrulesRead`: GetAlertRule
    fmt.Fprintf(os.Stdout, "Response from `AlertRulesApi.V1InstancesAlertgroupsAlertrulesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertName** | **string** |  | 
**groupName** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertgroupsAlertrulesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 

### Return type

[**GetAlertRule**](GetAlertRule.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertgroupsAlertrulesUpdate

> PutAlertRule V1InstancesAlertgroupsAlertrulesUpdate(ctx, alertName, groupName, instanceId).Authorization(authorization).V1InstancesAlertgroupsAlertrulesCreateRequest(v1InstancesAlertgroupsAlertrulesCreateRequest).Execute()





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
    alertName := "alertName_example" // string | 
    groupName := "groupName_example" // string | 
    instanceId := "instanceId_example" // string | 
    authorization := "authorization_example" // string | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user
    v1InstancesAlertgroupsAlertrulesCreateRequest := *openapiclient.NewV1InstancesAlertgroupsAlertrulesCreateRequest("Alert_example", "Expr_example", map[string]interface{}(123), map[string]interface{}(123)) // V1InstancesAlertgroupsAlertrulesCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRulesApi.V1InstancesAlertgroupsAlertrulesUpdate(context.Background(), alertName, groupName, instanceId).Authorization(authorization).V1InstancesAlertgroupsAlertrulesCreateRequest(v1InstancesAlertgroupsAlertrulesCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1InstancesAlertgroupsAlertrulesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsAlertrulesUpdate`: PutAlertRule
    fmt.Fprintf(os.Stdout, "Response from `AlertRulesApi.V1InstancesAlertgroupsAlertrulesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertName** | **string** |  | 
**groupName** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertgroupsAlertrulesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 
 **v1InstancesAlertgroupsAlertrulesCreateRequest** | [**V1InstancesAlertgroupsAlertrulesCreateRequest**](V1InstancesAlertgroupsAlertrulesCreateRequest.md) |  | 

### Return type

[**PutAlertRule**](PutAlertRule.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsAlertrulesAllDelete

> PostAlertRule V1ProjectsInstancesAlertgroupsAlertrulesAllDelete(ctx, groupName, instanceId, projectId).Authorization(authorization).AlertName(alertName).Execute()





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
    alertName := []string{"Inner_example"} // []string | Name of the alert rules that should be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesAllDelete(context.Background(), groupName, instanceId, projectId).Authorization(authorization).AlertName(alertName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesAllDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsAlertrulesAllDelete`: PostAlertRule
    fmt.Fprintf(os.Stdout, "Response from `AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesAllDelete`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertgroupsAlertrulesAllDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 
 **alertName** | **[]string** | Name of the alert rules that should be deleted | 

### Return type

[**PostAlertRule**](PostAlertRule.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsAlertrulesCreate

> PostAlertRule V1ProjectsInstancesAlertgroupsAlertrulesCreate(ctx, groupName, instanceId, projectId).Authorization(authorization).V1InstancesAlertgroupsAlertrulesCreateRequest(v1InstancesAlertgroupsAlertrulesCreateRequest).Execute()





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
    v1InstancesAlertgroupsAlertrulesCreateRequest := *openapiclient.NewV1InstancesAlertgroupsAlertrulesCreateRequest("Alert_example", "Expr_example", map[string]interface{}(123), map[string]interface{}(123)) // V1InstancesAlertgroupsAlertrulesCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesCreate(context.Background(), groupName, instanceId, projectId).Authorization(authorization).V1InstancesAlertgroupsAlertrulesCreateRequest(v1InstancesAlertgroupsAlertrulesCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsAlertrulesCreate`: PostAlertRule
    fmt.Fprintf(os.Stdout, "Response from `AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesCreate`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertgroupsAlertrulesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 
 **v1InstancesAlertgroupsAlertrulesCreateRequest** | [**V1InstancesAlertgroupsAlertrulesCreateRequest**](V1InstancesAlertgroupsAlertrulesCreateRequest.md) |  | 

### Return type

[**PostAlertRule**](PostAlertRule.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsAlertrulesDelete

> DeleteAlertRule V1ProjectsInstancesAlertgroupsAlertrulesDelete(ctx, alertName, groupName, instanceId, projectId).Authorization(authorization).Execute()





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
    alertName := "alertName_example" // string | 
    groupName := "groupName_example" // string | 
    instanceId := "instanceId_example" // string | 
    projectId := "projectId_example" // string | 
    authorization := "authorization_example" // string | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesDelete(context.Background(), alertName, groupName, instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsAlertrulesDelete`: DeleteAlertRule
    fmt.Fprintf(os.Stdout, "Response from `AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertName** | **string** |  | 
**groupName** | **string** |  | 
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertgroupsAlertrulesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 

### Return type

[**DeleteAlertRule**](DeleteAlertRule.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsAlertrulesList

> GetAllAlertRules V1ProjectsInstancesAlertgroupsAlertrulesList(ctx, groupName, instanceId, projectId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesList(context.Background(), groupName, instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsAlertrulesList`: GetAllAlertRules
    fmt.Fprintf(os.Stdout, "Response from `AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertgroupsAlertrulesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 

### Return type

[**GetAllAlertRules**](GetAllAlertRules.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsAlertrulesPartialUpdate

> PostAlertRule V1ProjectsInstancesAlertgroupsAlertrulesPartialUpdate(ctx, groupName, instanceId, projectId).Authorization(authorization).V1InstancesAlertgroupsAlertrulesCreateRequest(v1InstancesAlertgroupsAlertrulesCreateRequest).Execute()





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
    v1InstancesAlertgroupsAlertrulesCreateRequest := []openapiclient.V1InstancesAlertgroupsAlertrulesCreateRequest{*openapiclient.NewV1InstancesAlertgroupsAlertrulesCreateRequest("Alert_example", "Expr_example", map[string]interface{}(123), map[string]interface{}(123))} // []V1InstancesAlertgroupsAlertrulesCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesPartialUpdate(context.Background(), groupName, instanceId, projectId).Authorization(authorization).V1InstancesAlertgroupsAlertrulesCreateRequest(v1InstancesAlertgroupsAlertrulesCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsAlertrulesPartialUpdate`: PostAlertRule
    fmt.Fprintf(os.Stdout, "Response from `AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesPartialUpdate`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertgroupsAlertrulesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 
 **v1InstancesAlertgroupsAlertrulesCreateRequest** | [**[]V1InstancesAlertgroupsAlertrulesCreateRequest**](V1InstancesAlertgroupsAlertrulesCreateRequest.md) |  | 

### Return type

[**PostAlertRule**](PostAlertRule.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsAlertrulesRead

> GetAlertRule V1ProjectsInstancesAlertgroupsAlertrulesRead(ctx, alertName, groupName, instanceId, projectId).Authorization(authorization).Execute()





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
    alertName := "alertName_example" // string | 
    groupName := "groupName_example" // string | 
    instanceId := "instanceId_example" // string | 
    projectId := "projectId_example" // string | 
    authorization := "authorization_example" // string | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesRead(context.Background(), alertName, groupName, instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsAlertrulesRead`: GetAlertRule
    fmt.Fprintf(os.Stdout, "Response from `AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertName** | **string** |  | 
**groupName** | **string** |  | 
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertgroupsAlertrulesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 

### Return type

[**GetAlertRule**](GetAlertRule.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsAlertrulesUpdate

> PutAlertRule V1ProjectsInstancesAlertgroupsAlertrulesUpdate(ctx, alertName, groupName, instanceId, projectId).Authorization(authorization).V1InstancesAlertgroupsAlertrulesCreateRequest(v1InstancesAlertgroupsAlertrulesCreateRequest).Execute()





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
    alertName := "alertName_example" // string | 
    groupName := "groupName_example" // string | 
    instanceId := "instanceId_example" // string | 
    projectId := "projectId_example" // string | 
    authorization := "authorization_example" // string | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user
    v1InstancesAlertgroupsAlertrulesCreateRequest := *openapiclient.NewV1InstancesAlertgroupsAlertrulesCreateRequest("Alert_example", "Expr_example", map[string]interface{}(123), map[string]interface{}(123)) // V1InstancesAlertgroupsAlertrulesCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesUpdate(context.Background(), alertName, groupName, instanceId, projectId).Authorization(authorization).V1InstancesAlertgroupsAlertrulesCreateRequest(v1InstancesAlertgroupsAlertrulesCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsAlertrulesUpdate`: PutAlertRule
    fmt.Fprintf(os.Stdout, "Response from `AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertName** | **string** |  | 
**groupName** | **string** |  | 
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertgroupsAlertrulesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 
 **v1InstancesAlertgroupsAlertrulesCreateRequest** | [**V1InstancesAlertgroupsAlertrulesCreateRequest**](V1InstancesAlertgroupsAlertrulesCreateRequest.md) |  | 

### Return type

[**PutAlertRule**](PutAlertRule.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

