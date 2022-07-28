# \AlertRecordsApi

All URIs are relative to *https://api.argus.eu01.stackit.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1InstancesAlertgroupsRecordsCreate**](AlertRecordsApi.md#V1InstancesAlertgroupsRecordsCreate) | **Post** /v1/instances/{instanceId}/alertgroups/{groupName}/records | 
[**V1InstancesAlertgroupsRecordsDelete**](AlertRecordsApi.md#V1InstancesAlertgroupsRecordsDelete) | **Delete** /v1/instances/{instanceId}/alertgroups/{groupName}/records | 
[**V1InstancesAlertgroupsRecordsDelete_0**](AlertRecordsApi.md#V1InstancesAlertgroupsRecordsDelete_0) | **Delete** /v1/instances/{instanceId}/alertgroups/{groupName}/records/{alertRecord} | 
[**V1InstancesAlertgroupsRecordsList**](AlertRecordsApi.md#V1InstancesAlertgroupsRecordsList) | **Get** /v1/instances/{instanceId}/alertgroups/{groupName}/records | 
[**V1InstancesAlertgroupsRecordsPartialUpdate**](AlertRecordsApi.md#V1InstancesAlertgroupsRecordsPartialUpdate) | **Patch** /v1/instances/{instanceId}/alertgroups/{groupName}/records | 
[**V1InstancesAlertgroupsRecordsRead**](AlertRecordsApi.md#V1InstancesAlertgroupsRecordsRead) | **Get** /v1/instances/{instanceId}/alertgroups/{groupName}/records/{alertRecord} | 
[**V1InstancesAlertgroupsRecordsUpdate**](AlertRecordsApi.md#V1InstancesAlertgroupsRecordsUpdate) | **Put** /v1/instances/{instanceId}/alertgroups/{groupName}/records/{alertRecord} | 
[**V1ProjectsInstancesAlertgroupsRecordsCreate**](AlertRecordsApi.md#V1ProjectsInstancesAlertgroupsRecordsCreate) | **Post** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/records | 
[**V1ProjectsInstancesAlertgroupsRecordsDelete**](AlertRecordsApi.md#V1ProjectsInstancesAlertgroupsRecordsDelete) | **Delete** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/records | 
[**V1ProjectsInstancesAlertgroupsRecordsDelete_0**](AlertRecordsApi.md#V1ProjectsInstancesAlertgroupsRecordsDelete_0) | **Delete** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/records/{alertRecord} | 
[**V1ProjectsInstancesAlertgroupsRecordsList**](AlertRecordsApi.md#V1ProjectsInstancesAlertgroupsRecordsList) | **Get** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/records | 
[**V1ProjectsInstancesAlertgroupsRecordsPartialUpdate**](AlertRecordsApi.md#V1ProjectsInstancesAlertgroupsRecordsPartialUpdate) | **Patch** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/records | 
[**V1ProjectsInstancesAlertgroupsRecordsRead**](AlertRecordsApi.md#V1ProjectsInstancesAlertgroupsRecordsRead) | **Get** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/records/{alertRecord} | 
[**V1ProjectsInstancesAlertgroupsRecordsUpdate**](AlertRecordsApi.md#V1ProjectsInstancesAlertgroupsRecordsUpdate) | **Put** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/records/{alertRecord} | 



## V1InstancesAlertgroupsRecordsCreate

> AlertRecordsResponse V1InstancesAlertgroupsRecordsCreate(ctx, groupName, instanceId).Authorization(authorization).Data(data).Execute()





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
    data := *openapiclient.NewV1InstancesAlertgroupsRecordsCreateRequest("Record_example", "Expr_example") // V1InstancesAlertgroupsRecordsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRecordsApi.V1InstancesAlertgroupsRecordsCreate(context.Background(), groupName, instanceId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRecordsApi.V1InstancesAlertgroupsRecordsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsRecordsCreate`: AlertRecordsResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertRecordsApi.V1InstancesAlertgroupsRecordsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertgroupsRecordsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesAlertgroupsRecordsCreateRequest**](V1InstancesAlertgroupsRecordsCreateRequest.md) |  | 

### Return type

[**AlertRecordsResponse**](AlertRecordsResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertgroupsRecordsDelete

> AlertRecordsResponse V1InstancesAlertgroupsRecordsDelete(ctx, groupName, instanceId).Authorization(authorization).AlertRecord(alertRecord).Execute()





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
    alertRecord := []string{"Inner_example"} // []string | Name of the records that should be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRecordsApi.V1InstancesAlertgroupsRecordsDelete(context.Background(), groupName, instanceId).Authorization(authorization).AlertRecord(alertRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRecordsApi.V1InstancesAlertgroupsRecordsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsRecordsDelete`: AlertRecordsResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertRecordsApi.V1InstancesAlertgroupsRecordsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertgroupsRecordsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **alertRecord** | **[]string** | Name of the records that should be deleted | 

### Return type

[**AlertRecordsResponse**](AlertRecordsResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertgroupsRecordsDelete_0

> AlertRecordsResponse V1InstancesAlertgroupsRecordsDelete_0(ctx, alertRecord, groupName, instanceId).Authorization(authorization).Execute()





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
    alertRecord := "alertRecord_example" // string | 
    groupName := "groupName_example" // string | 
    instanceId := "instanceId_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRecordsApi.V1InstancesAlertgroupsRecordsDelete_0(context.Background(), alertRecord, groupName, instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRecordsApi.V1InstancesAlertgroupsRecordsDelete_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsRecordsDelete_0`: AlertRecordsResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertRecordsApi.V1InstancesAlertgroupsRecordsDelete_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertRecord** | **string** |  | 
**groupName** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertgroupsRecordsDelete_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**AlertRecordsResponse**](AlertRecordsResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertgroupsRecordsList

> AlertRecordsResponse V1InstancesAlertgroupsRecordsList(ctx, groupName, instanceId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.AlertRecordsApi.V1InstancesAlertgroupsRecordsList(context.Background(), groupName, instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRecordsApi.V1InstancesAlertgroupsRecordsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsRecordsList`: AlertRecordsResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertRecordsApi.V1InstancesAlertgroupsRecordsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertgroupsRecordsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**AlertRecordsResponse**](AlertRecordsResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertgroupsRecordsPartialUpdate

> AlertRecordsResponse V1InstancesAlertgroupsRecordsPartialUpdate(ctx, groupName, instanceId).Authorization(authorization).Data(data).Execute()





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
    data := []openapiclient.V1InstancesAlertgroupsRecordsCreateRequest{*openapiclient.NewV1InstancesAlertgroupsRecordsCreateRequest("Record_example", "Expr_example")} // []V1InstancesAlertgroupsRecordsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRecordsApi.V1InstancesAlertgroupsRecordsPartialUpdate(context.Background(), groupName, instanceId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRecordsApi.V1InstancesAlertgroupsRecordsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsRecordsPartialUpdate`: AlertRecordsResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertRecordsApi.V1InstancesAlertgroupsRecordsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertgroupsRecordsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**[]V1InstancesAlertgroupsRecordsCreateRequest**](V1InstancesAlertgroupsRecordsCreateRequest.md) |  | 

### Return type

[**AlertRecordsResponse**](AlertRecordsResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertgroupsRecordsRead

> AlertRecordResponse V1InstancesAlertgroupsRecordsRead(ctx, alertRecord, groupName, instanceId).Authorization(authorization).Execute()





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
    alertRecord := "alertRecord_example" // string | 
    groupName := "groupName_example" // string | 
    instanceId := "instanceId_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRecordsApi.V1InstancesAlertgroupsRecordsRead(context.Background(), alertRecord, groupName, instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRecordsApi.V1InstancesAlertgroupsRecordsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsRecordsRead`: AlertRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertRecordsApi.V1InstancesAlertgroupsRecordsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertRecord** | **string** |  | 
**groupName** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertgroupsRecordsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**AlertRecordResponse**](AlertRecordResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesAlertgroupsRecordsUpdate

> AlertRecordsResponse V1InstancesAlertgroupsRecordsUpdate(ctx, alertRecord, groupName, instanceId).Authorization(authorization).Data(data).Execute()





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
    alertRecord := "alertRecord_example" // string | 
    groupName := "groupName_example" // string | 
    instanceId := "instanceId_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.
    data := *openapiclient.NewV1InstancesAlertgroupsRecordsUpdateRequest("Expr_example") // V1InstancesAlertgroupsRecordsUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRecordsApi.V1InstancesAlertgroupsRecordsUpdate(context.Background(), alertRecord, groupName, instanceId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRecordsApi.V1InstancesAlertgroupsRecordsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesAlertgroupsRecordsUpdate`: AlertRecordsResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertRecordsApi.V1InstancesAlertgroupsRecordsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertRecord** | **string** |  | 
**groupName** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesAlertgroupsRecordsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesAlertgroupsRecordsUpdateRequest**](V1InstancesAlertgroupsRecordsUpdateRequest.md) |  | 

### Return type

[**AlertRecordsResponse**](AlertRecordsResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsRecordsCreate

> AlertRecordsResponse V1ProjectsInstancesAlertgroupsRecordsCreate(ctx, groupName, instanceId, projectId).Authorization(authorization).Data(data).Execute()





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
    data := *openapiclient.NewV1InstancesAlertgroupsRecordsCreateRequest("Record_example", "Expr_example") // V1InstancesAlertgroupsRecordsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRecordsApi.V1ProjectsInstancesAlertgroupsRecordsCreate(context.Background(), groupName, instanceId, projectId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRecordsApi.V1ProjectsInstancesAlertgroupsRecordsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsRecordsCreate`: AlertRecordsResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertRecordsApi.V1ProjectsInstancesAlertgroupsRecordsCreate`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertgroupsRecordsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesAlertgroupsRecordsCreateRequest**](V1InstancesAlertgroupsRecordsCreateRequest.md) |  | 

### Return type

[**AlertRecordsResponse**](AlertRecordsResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsRecordsDelete

> AlertRecordsResponse V1ProjectsInstancesAlertgroupsRecordsDelete(ctx, groupName, instanceId, projectId).Authorization(authorization).AlertRecord(alertRecord).Execute()





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
    alertRecord := []string{"Inner_example"} // []string | Name of the records that should be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRecordsApi.V1ProjectsInstancesAlertgroupsRecordsDelete(context.Background(), groupName, instanceId, projectId).Authorization(authorization).AlertRecord(alertRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRecordsApi.V1ProjectsInstancesAlertgroupsRecordsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsRecordsDelete`: AlertRecordsResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertRecordsApi.V1ProjectsInstancesAlertgroupsRecordsDelete`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertgroupsRecordsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **alertRecord** | **[]string** | Name of the records that should be deleted | 

### Return type

[**AlertRecordsResponse**](AlertRecordsResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsRecordsDelete_0

> AlertRecordsResponse V1ProjectsInstancesAlertgroupsRecordsDelete_0(ctx, alertRecord, groupName, instanceId, projectId).Authorization(authorization).Execute()





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
    alertRecord := "alertRecord_example" // string | 
    groupName := "groupName_example" // string | 
    instanceId := "instanceId_example" // string | 
    projectId := "projectId_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRecordsApi.V1ProjectsInstancesAlertgroupsRecordsDelete_0(context.Background(), alertRecord, groupName, instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRecordsApi.V1ProjectsInstancesAlertgroupsRecordsDelete_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsRecordsDelete_0`: AlertRecordsResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertRecordsApi.V1ProjectsInstancesAlertgroupsRecordsDelete_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertRecord** | **string** |  | 
**groupName** | **string** |  | 
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertgroupsRecordsDelete_2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**AlertRecordsResponse**](AlertRecordsResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsRecordsList

> AlertRecordsResponse V1ProjectsInstancesAlertgroupsRecordsList(ctx, groupName, instanceId, projectId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.AlertRecordsApi.V1ProjectsInstancesAlertgroupsRecordsList(context.Background(), groupName, instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRecordsApi.V1ProjectsInstancesAlertgroupsRecordsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsRecordsList`: AlertRecordsResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertRecordsApi.V1ProjectsInstancesAlertgroupsRecordsList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertgroupsRecordsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**AlertRecordsResponse**](AlertRecordsResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsRecordsPartialUpdate

> AlertRecordsResponse V1ProjectsInstancesAlertgroupsRecordsPartialUpdate(ctx, groupName, instanceId, projectId).Authorization(authorization).Data(data).Execute()





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
    data := []openapiclient.V1InstancesAlertgroupsRecordsCreateRequest{*openapiclient.NewV1InstancesAlertgroupsRecordsCreateRequest("Record_example", "Expr_example")} // []V1InstancesAlertgroupsRecordsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRecordsApi.V1ProjectsInstancesAlertgroupsRecordsPartialUpdate(context.Background(), groupName, instanceId, projectId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRecordsApi.V1ProjectsInstancesAlertgroupsRecordsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsRecordsPartialUpdate`: AlertRecordsResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertRecordsApi.V1ProjectsInstancesAlertgroupsRecordsPartialUpdate`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertgroupsRecordsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**[]V1InstancesAlertgroupsRecordsCreateRequest**](V1InstancesAlertgroupsRecordsCreateRequest.md) |  | 

### Return type

[**AlertRecordsResponse**](AlertRecordsResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsRecordsRead

> AlertRecordResponse V1ProjectsInstancesAlertgroupsRecordsRead(ctx, alertRecord, groupName, instanceId, projectId).Authorization(authorization).Execute()





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
    alertRecord := "alertRecord_example" // string | 
    groupName := "groupName_example" // string | 
    instanceId := "instanceId_example" // string | 
    projectId := "projectId_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRecordsApi.V1ProjectsInstancesAlertgroupsRecordsRead(context.Background(), alertRecord, groupName, instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRecordsApi.V1ProjectsInstancesAlertgroupsRecordsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsRecordsRead`: AlertRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertRecordsApi.V1ProjectsInstancesAlertgroupsRecordsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertRecord** | **string** |  | 
**groupName** | **string** |  | 
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertgroupsRecordsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**AlertRecordResponse**](AlertRecordResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesAlertgroupsRecordsUpdate

> AlertRecordsResponse V1ProjectsInstancesAlertgroupsRecordsUpdate(ctx, alertRecord, groupName, instanceId, projectId).Authorization(authorization).Data(data).Execute()





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
    alertRecord := "alertRecord_example" // string | 
    groupName := "groupName_example" // string | 
    instanceId := "instanceId_example" // string | 
    projectId := "projectId_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.
    data := *openapiclient.NewV1InstancesAlertgroupsRecordsUpdateRequest("Expr_example") // V1InstancesAlertgroupsRecordsUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertRecordsApi.V1ProjectsInstancesAlertgroupsRecordsUpdate(context.Background(), alertRecord, groupName, instanceId, projectId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertRecordsApi.V1ProjectsInstancesAlertgroupsRecordsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesAlertgroupsRecordsUpdate`: AlertRecordsResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertRecordsApi.V1ProjectsInstancesAlertgroupsRecordsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertRecord** | **string** |  | 
**groupName** | **string** |  | 
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesAlertgroupsRecordsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesAlertgroupsRecordsUpdateRequest**](V1InstancesAlertgroupsRecordsUpdateRequest.md) |  | 

### Return type

[**AlertRecordsResponse**](AlertRecordsResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

