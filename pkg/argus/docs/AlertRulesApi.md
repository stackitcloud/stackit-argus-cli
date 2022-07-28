# \AlertRulesApi

All URIs are relative to *https://api.argus.eu01.stackit.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1InstancesAlertgroupsAlertrulesCreate**](AlertRulesApi.md#V1InstancesAlertgroupsAlertrulesCreate) | **Post** /v1/instances/{instanceId}/alertgroups/{groupName}/alertrules | 
[**V1InstancesAlertgroupsAlertrulesDelete**](AlertRulesApi.md#V1InstancesAlertgroupsAlertrulesDelete) | **Delete** /v1/instances/{instanceId}/alertgroups/{groupName}/alertrules | 
[**V1InstancesAlertgroupsAlertrulesDelete_0**](AlertRulesApi.md#V1InstancesAlertgroupsAlertrulesDelete_0) | **Delete** /v1/instances/{instanceId}/alertgroups/{groupName}/alertrules/{alertName} | 
[**V1InstancesAlertgroupsAlertrulesList**](AlertRulesApi.md#V1InstancesAlertgroupsAlertrulesList) | **Get** /v1/instances/{instanceId}/alertgroups/{groupName}/alertrules | 
[**V1InstancesAlertgroupsAlertrulesPartialUpdate**](AlertRulesApi.md#V1InstancesAlertgroupsAlertrulesPartialUpdate) | **Patch** /v1/instances/{instanceId}/alertgroups/{groupName}/alertrules | 
[**V1InstancesAlertgroupsAlertrulesRead**](AlertRulesApi.md#V1InstancesAlertgroupsAlertrulesRead) | **Get** /v1/instances/{instanceId}/alertgroups/{groupName}/alertrules/{alertName} | 
[**V1InstancesAlertgroupsAlertrulesUpdate**](AlertRulesApi.md#V1InstancesAlertgroupsAlertrulesUpdate) | **Put** /v1/instances/{instanceId}/alertgroups/{groupName}/alertrules/{alertName} | 
[**V1ProjectsInstancesAlertgroupsAlertrulesCreate**](AlertRulesApi.md#V1ProjectsInstancesAlertgroupsAlertrulesCreate) | **Post** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/alertrules | 
[**V1ProjectsInstancesAlertgroupsAlertrulesDelete**](AlertRulesApi.md#V1ProjectsInstancesAlertgroupsAlertrulesDelete) | **Delete** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/alertrules | 
[**V1ProjectsInstancesAlertgroupsAlertrulesDelete_0**](AlertRulesApi.md#V1ProjectsInstancesAlertgroupsAlertrulesDelete_0) | **Delete** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/alertrules/{alertName} | 
[**V1ProjectsInstancesAlertgroupsAlertrulesList**](AlertRulesApi.md#V1ProjectsInstancesAlertgroupsAlertrulesList) | **Get** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/alertrules | 
[**V1ProjectsInstancesAlertgroupsAlertrulesPartialUpdate**](AlertRulesApi.md#V1ProjectsInstancesAlertgroupsAlertrulesPartialUpdate) | **Patch** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/alertrules | 
[**V1ProjectsInstancesAlertgroupsAlertrulesRead**](AlertRulesApi.md#V1ProjectsInstancesAlertgroupsAlertrulesRead) | **Get** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/alertrules/{alertName} | 
[**V1ProjectsInstancesAlertgroupsAlertrulesUpdate**](AlertRulesApi.md#V1ProjectsInstancesAlertgroupsAlertrulesUpdate) | **Put** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/alertrules/{alertName} | 



## V1InstancesAlertgroupsAlertrulesCreate

> AlertRulesResponse V1InstancesAlertgroupsAlertrulesCreate(ctx, groupName, instanceId).Authorization(authorization).Data(data).Execute()





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
    data := *openapiclient.NewV1InstancesAlertgroupsCreateRequestRulesInner("Alert_example", "Expr_example") // V1InstancesAlertgroupsCreateRequestRulesInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRulesApi.V1InstancesAlertgroupsAlertrulesCreate(context.Background(), groupName, instanceId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1InstancesAlertgroupsAlertrulesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsAlertrulesCreate`: AlertRulesResponse
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


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesAlertgroupsCreateRequestRulesInner**](V1InstancesAlertgroupsCreateRequestRulesInner.md) |  | 

### Return type

[**AlertRulesResponse**](AlertRulesResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertgroupsAlertrulesDelete

> AlertRulesResponse V1InstancesAlertgroupsAlertrulesDelete(ctx, groupName, instanceId).Authorization(authorization).AlertName(alertName).Execute()





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
    alertName := []string{"Inner_example"} // []string | Name of the alert rules that should be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRulesApi.V1InstancesAlertgroupsAlertrulesDelete(context.Background(), groupName, instanceId).Authorization(authorization).AlertName(alertName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1InstancesAlertgroupsAlertrulesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsAlertrulesDelete`: AlertRulesResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertRulesApi.V1InstancesAlertgroupsAlertrulesDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertgroupsAlertrulesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **alertName** | **[]string** | Name of the alert rules that should be deleted | 

### Return type

[**AlertRulesResponse**](AlertRulesResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertgroupsAlertrulesDelete_0

> AlertRulesResponse V1InstancesAlertgroupsAlertrulesDelete_0(ctx, alertName, groupName, instanceId).Authorization(authorization).Execute()





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
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRulesApi.V1InstancesAlertgroupsAlertrulesDelete_0(context.Background(), alertName, groupName, instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1InstancesAlertgroupsAlertrulesDelete_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsAlertrulesDelete_0`: AlertRulesResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertRulesApi.V1InstancesAlertgroupsAlertrulesDelete_0`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiV1InstancesAlertgroupsAlertrulesDelete_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**AlertRulesResponse**](AlertRulesResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertgroupsAlertrulesList

> AlertRulesResponse V1InstancesAlertgroupsAlertrulesList(ctx, groupName, instanceId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.AlertRulesApi.V1InstancesAlertgroupsAlertrulesList(context.Background(), groupName, instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1InstancesAlertgroupsAlertrulesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsAlertrulesList`: AlertRulesResponse
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


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**AlertRulesResponse**](AlertRulesResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertgroupsAlertrulesPartialUpdate

> AlertRulesResponse V1InstancesAlertgroupsAlertrulesPartialUpdate(ctx, groupName, instanceId).Authorization(authorization).Data(data).Execute()





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
    data := []openapiclient.V1InstancesAlertgroupsCreateRequestRulesInner{*openapiclient.NewV1InstancesAlertgroupsCreateRequestRulesInner("Alert_example", "Expr_example")} // []V1InstancesAlertgroupsCreateRequestRulesInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRulesApi.V1InstancesAlertgroupsAlertrulesPartialUpdate(context.Background(), groupName, instanceId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1InstancesAlertgroupsAlertrulesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsAlertrulesPartialUpdate`: AlertRulesResponse
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


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**[]V1InstancesAlertgroupsCreateRequestRulesInner**](V1InstancesAlertgroupsCreateRequestRulesInner.md) |  | 

### Return type

[**AlertRulesResponse**](AlertRulesResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertgroupsAlertrulesRead

> AlertRuleResponse V1InstancesAlertgroupsAlertrulesRead(ctx, alertName, groupName, instanceId).Authorization(authorization).Execute()





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
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRulesApi.V1InstancesAlertgroupsAlertrulesRead(context.Background(), alertName, groupName, instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1InstancesAlertgroupsAlertrulesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsAlertrulesRead`: AlertRuleResponse
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



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**AlertRuleResponse**](AlertRuleResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertgroupsAlertrulesUpdate

> AlertRulesResponse V1InstancesAlertgroupsAlertrulesUpdate(ctx, alertName, groupName, instanceId).Authorization(authorization).Data(data).Execute()





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
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.
    data := *openapiclient.NewV1InstancesAlertgroupsAlertrulesUpdateRequest("Expr_example") // V1InstancesAlertgroupsAlertrulesUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRulesApi.V1InstancesAlertgroupsAlertrulesUpdate(context.Background(), alertName, groupName, instanceId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1InstancesAlertgroupsAlertrulesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsAlertrulesUpdate`: AlertRulesResponse
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



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesAlertgroupsAlertrulesUpdateRequest**](V1InstancesAlertgroupsAlertrulesUpdateRequest.md) |  | 

### Return type

[**AlertRulesResponse**](AlertRulesResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsAlertrulesCreate

> AlertRulesResponse V1ProjectsInstancesAlertgroupsAlertrulesCreate(ctx, groupName, instanceId, projectId).Authorization(authorization).Data(data).Execute()





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
    data := *openapiclient.NewV1InstancesAlertgroupsCreateRequestRulesInner("Alert_example", "Expr_example") // V1InstancesAlertgroupsCreateRequestRulesInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesCreate(context.Background(), groupName, instanceId, projectId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsAlertrulesCreate`: AlertRulesResponse
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



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesAlertgroupsCreateRequestRulesInner**](V1InstancesAlertgroupsCreateRequestRulesInner.md) |  | 

### Return type

[**AlertRulesResponse**](AlertRulesResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsAlertrulesDelete

> AlertRulesResponse V1ProjectsInstancesAlertgroupsAlertrulesDelete(ctx, groupName, instanceId, projectId).Authorization(authorization).AlertName(alertName).Execute()





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
    alertName := []string{"Inner_example"} // []string | Name of the alert rules that should be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesDelete(context.Background(), groupName, instanceId, projectId).Authorization(authorization).AlertName(alertName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsAlertrulesDelete`: AlertRulesResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesDelete`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertgroupsAlertrulesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **alertName** | **[]string** | Name of the alert rules that should be deleted | 

### Return type

[**AlertRulesResponse**](AlertRulesResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsAlertrulesDelete_0

> AlertRulesResponse V1ProjectsInstancesAlertgroupsAlertrulesDelete_0(ctx, alertName, groupName, instanceId, projectId).Authorization(authorization).Execute()





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
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesDelete_0(context.Background(), alertName, groupName, instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesDelete_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsAlertrulesDelete_0`: AlertRulesResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesDelete_0`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertgroupsAlertrulesDelete_2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**AlertRulesResponse**](AlertRulesResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsAlertrulesList

> AlertRulesResponse V1ProjectsInstancesAlertgroupsAlertrulesList(ctx, groupName, instanceId, projectId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesList(context.Background(), groupName, instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsAlertrulesList`: AlertRulesResponse
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



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**AlertRulesResponse**](AlertRulesResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsAlertrulesPartialUpdate

> AlertRulesResponse V1ProjectsInstancesAlertgroupsAlertrulesPartialUpdate(ctx, groupName, instanceId, projectId).Authorization(authorization).Data(data).Execute()





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
    data := []openapiclient.V1InstancesAlertgroupsCreateRequestRulesInner{*openapiclient.NewV1InstancesAlertgroupsCreateRequestRulesInner("Alert_example", "Expr_example")} // []V1InstancesAlertgroupsCreateRequestRulesInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesPartialUpdate(context.Background(), groupName, instanceId, projectId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsAlertrulesPartialUpdate`: AlertRulesResponse
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



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**[]V1InstancesAlertgroupsCreateRequestRulesInner**](V1InstancesAlertgroupsCreateRequestRulesInner.md) |  | 

### Return type

[**AlertRulesResponse**](AlertRulesResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsAlertrulesRead

> AlertRuleResponse V1ProjectsInstancesAlertgroupsAlertrulesRead(ctx, alertName, groupName, instanceId, projectId).Authorization(authorization).Execute()





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
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesRead(context.Background(), alertName, groupName, instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsAlertrulesRead`: AlertRuleResponse
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




 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**AlertRuleResponse**](AlertRuleResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsAlertrulesUpdate

> AlertRulesResponse V1ProjectsInstancesAlertgroupsAlertrulesUpdate(ctx, alertName, groupName, instanceId, projectId).Authorization(authorization).Data(data).Execute()





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
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.
    data := *openapiclient.NewV1InstancesAlertgroupsAlertrulesUpdateRequest("Expr_example") // V1InstancesAlertgroupsAlertrulesUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesUpdate(context.Background(), alertName, groupName, instanceId, projectId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRulesApi.V1ProjectsInstancesAlertgroupsAlertrulesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsAlertrulesUpdate`: AlertRulesResponse
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




 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesAlertgroupsAlertrulesUpdateRequest**](V1InstancesAlertgroupsAlertrulesUpdateRequest.md) |  | 

### Return type

[**AlertRulesResponse**](AlertRulesResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

