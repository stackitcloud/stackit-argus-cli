# \BackupApi

All URIs are relative to *https://api.argus.eu01.stackit.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1InstancesBackupRestoresCreate**](BackupApi.md#V1InstancesBackupRestoresCreate) | **Post** /v1/instances/{instanceId}/backup-restores/{backupDate} | 
[**V1InstancesBackupRetentionsList**](BackupApi.md#V1InstancesBackupRetentionsList) | **Get** /v1/instances/{instanceId}/backup-retentions | 
[**V1InstancesBackupRetentionsUpdate**](BackupApi.md#V1InstancesBackupRetentionsUpdate) | **Put** /v1/instances/{instanceId}/backup-retentions | 
[**V1InstancesBackupSchedulesCreate**](BackupApi.md#V1InstancesBackupSchedulesCreate) | **Post** /v1/instances/{instanceId}/backup-schedules | 
[**V1InstancesBackupSchedulesDelete**](BackupApi.md#V1InstancesBackupSchedulesDelete) | **Delete** /v1/instances/{instanceId}/backup-schedules/{scheduleId} | 
[**V1InstancesBackupSchedulesList**](BackupApi.md#V1InstancesBackupSchedulesList) | **Get** /v1/instances/{instanceId}/backup-schedules | 
[**V1InstancesBackupSchedulesUpdate**](BackupApi.md#V1InstancesBackupSchedulesUpdate) | **Put** /v1/instances/{instanceId}/backup-schedules/{scheduleId} | 
[**V1InstancesBackupsCreate**](BackupApi.md#V1InstancesBackupsCreate) | **Post** /v1/instances/{instanceId}/backups | 
[**V1InstancesBackupsDelete**](BackupApi.md#V1InstancesBackupsDelete) | **Delete** /v1/instances/{instanceId}/backups/{backupDate} | 
[**V1InstancesBackupsList**](BackupApi.md#V1InstancesBackupsList) | **Get** /v1/instances/{instanceId}/backups | 
[**V1ProjectsInstancesBackupRestoresCreate**](BackupApi.md#V1ProjectsInstancesBackupRestoresCreate) | **Post** /v1/projects/{projectId}/instances/{instanceId}/backup-restores/{backupDate} | 
[**V1ProjectsInstancesBackupRetentionsList**](BackupApi.md#V1ProjectsInstancesBackupRetentionsList) | **Get** /v1/projects/{projectId}/instances/{instanceId}/backup-retentions | 
[**V1ProjectsInstancesBackupRetentionsUpdate**](BackupApi.md#V1ProjectsInstancesBackupRetentionsUpdate) | **Put** /v1/projects/{projectId}/instances/{instanceId}/backup-retentions | 
[**V1ProjectsInstancesBackupSchedulesCreate**](BackupApi.md#V1ProjectsInstancesBackupSchedulesCreate) | **Post** /v1/projects/{projectId}/instances/{instanceId}/backup-schedules | 
[**V1ProjectsInstancesBackupSchedulesDelete**](BackupApi.md#V1ProjectsInstancesBackupSchedulesDelete) | **Delete** /v1/projects/{projectId}/instances/{instanceId}/backup-schedules/{scheduleId} | 
[**V1ProjectsInstancesBackupSchedulesList**](BackupApi.md#V1ProjectsInstancesBackupSchedulesList) | **Get** /v1/projects/{projectId}/instances/{instanceId}/backup-schedules | 
[**V1ProjectsInstancesBackupSchedulesUpdate**](BackupApi.md#V1ProjectsInstancesBackupSchedulesUpdate) | **Put** /v1/projects/{projectId}/instances/{instanceId}/backup-schedules/{scheduleId} | 
[**V1ProjectsInstancesBackupsCreate**](BackupApi.md#V1ProjectsInstancesBackupsCreate) | **Post** /v1/projects/{projectId}/instances/{instanceId}/backups | 
[**V1ProjectsInstancesBackupsDelete**](BackupApi.md#V1ProjectsInstancesBackupsDelete) | **Delete** /v1/projects/{projectId}/instances/{instanceId}/backups/{backupDate} | 
[**V1ProjectsInstancesBackupsList**](BackupApi.md#V1ProjectsInstancesBackupsList) | **Get** /v1/projects/{projectId}/instances/{instanceId}/backups | 



## V1InstancesBackupRestoresCreate

> Message V1InstancesBackupRestoresCreate(ctx, backupDate, instanceId).Authorization(authorization).RestoreTarget(restoreTarget).Execute()





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
    backupDate := "backupDate_example" // string | 
    instanceId := "instanceId_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.
    restoreTarget := "restoreTarget_example" // string | List of restore targets

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupApi.V1InstancesBackupRestoresCreate(context.Background(), backupDate, instanceId).Authorization(authorization).RestoreTarget(restoreTarget).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupApi.V1InstancesBackupRestoresCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesBackupRestoresCreate`: Message
    fmt.Fprintf(os.Stdout, "Response from `BackupApi.V1InstancesBackupRestoresCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupDate** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesBackupRestoresCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **restoreTarget** | **string** | List of restore targets | 

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


## V1InstancesBackupRetentionsList

> BackupRetentionResponse V1InstancesBackupRetentionsList(ctx, instanceId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.BackupApi.V1InstancesBackupRetentionsList(context.Background(), instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupApi.V1InstancesBackupRetentionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesBackupRetentionsList`: BackupRetentionResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupApi.V1InstancesBackupRetentionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesBackupRetentionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**BackupRetentionResponse**](BackupRetentionResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesBackupRetentionsUpdate

> Message V1InstancesBackupRetentionsUpdate(ctx, instanceId).Authorization(authorization).Data(data).BackupTarget(backupTarget).Execute()





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
    data := *openapiclient.NewV1InstancesBackupRetentionsUpdateRequest("Retention_example") // V1InstancesBackupRetentionsUpdateRequest | 
    backupTarget := []string{"BackupTarget_example"} // []string | List of backup targets (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupApi.V1InstancesBackupRetentionsUpdate(context.Background(), instanceId).Authorization(authorization).Data(data).BackupTarget(backupTarget).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupApi.V1InstancesBackupRetentionsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesBackupRetentionsUpdate`: Message
    fmt.Fprintf(os.Stdout, "Response from `BackupApi.V1InstancesBackupRetentionsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesBackupRetentionsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesBackupRetentionsUpdateRequest**](V1InstancesBackupRetentionsUpdateRequest.md) |  | 
 **backupTarget** | **[]string** | List of backup targets | 

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


## V1InstancesBackupSchedulesCreate

> BackupSchedulePostResponse V1InstancesBackupSchedulesCreate(ctx, instanceId).Authorization(authorization).Data(data).BackupTarget(backupTarget).Execute()





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
    data := *openapiclient.NewV1InstancesBackupSchedulesCreateRequest("Schedule_example") // V1InstancesBackupSchedulesCreateRequest | 
    backupTarget := []string{"BackupTarget_example"} // []string | List of backup targets (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupApi.V1InstancesBackupSchedulesCreate(context.Background(), instanceId).Authorization(authorization).Data(data).BackupTarget(backupTarget).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupApi.V1InstancesBackupSchedulesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesBackupSchedulesCreate`: BackupSchedulePostResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupApi.V1InstancesBackupSchedulesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesBackupSchedulesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesBackupSchedulesCreateRequest**](V1InstancesBackupSchedulesCreateRequest.md) |  | 
 **backupTarget** | **[]string** | List of backup targets | 

### Return type

[**BackupSchedulePostResponse**](BackupSchedulePostResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesBackupSchedulesDelete

> Message V1InstancesBackupSchedulesDelete(ctx, instanceId, scheduleId).Authorization(authorization).Execute()





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
    scheduleId := "scheduleId_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupApi.V1InstancesBackupSchedulesDelete(context.Background(), instanceId, scheduleId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupApi.V1InstancesBackupSchedulesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesBackupSchedulesDelete`: Message
    fmt.Fprintf(os.Stdout, "Response from `BackupApi.V1InstancesBackupSchedulesDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**scheduleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesBackupSchedulesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

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


## V1InstancesBackupSchedulesList

> BackupScheduleResponse V1InstancesBackupSchedulesList(ctx, instanceId).Authorization(authorization).BackupTarget(backupTarget).Execute()





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
    backupTarget := []string{"BackupTarget_example"} // []string | List of backup targets (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupApi.V1InstancesBackupSchedulesList(context.Background(), instanceId).Authorization(authorization).BackupTarget(backupTarget).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupApi.V1InstancesBackupSchedulesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesBackupSchedulesList`: BackupScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupApi.V1InstancesBackupSchedulesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesBackupSchedulesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **backupTarget** | **[]string** | List of backup targets | 

### Return type

[**BackupScheduleResponse**](BackupScheduleResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesBackupSchedulesUpdate

> BackupSchedulePutResponse V1InstancesBackupSchedulesUpdate(ctx, instanceId, scheduleId).Authorization(authorization).Data(data).Execute()





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
    scheduleId := "scheduleId_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.
    data := *openapiclient.NewV1InstancesBackupSchedulesCreateRequest("Schedule_example") // V1InstancesBackupSchedulesCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupApi.V1InstancesBackupSchedulesUpdate(context.Background(), instanceId, scheduleId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupApi.V1InstancesBackupSchedulesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesBackupSchedulesUpdate`: BackupSchedulePutResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupApi.V1InstancesBackupSchedulesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**scheduleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesBackupSchedulesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesBackupSchedulesCreateRequest**](V1InstancesBackupSchedulesCreateRequest.md) |  | 

### Return type

[**BackupSchedulePutResponse**](BackupSchedulePutResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesBackupsCreate

> Message V1InstancesBackupsCreate(ctx, instanceId).Authorization(authorization).BackupTarget(backupTarget).Execute()





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
    backupTarget := []string{"BackupTarget_example"} // []string | List of backup targets (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupApi.V1InstancesBackupsCreate(context.Background(), instanceId).Authorization(authorization).BackupTarget(backupTarget).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupApi.V1InstancesBackupsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesBackupsCreate`: Message
    fmt.Fprintf(os.Stdout, "Response from `BackupApi.V1InstancesBackupsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesBackupsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **backupTarget** | **[]string** | List of backup targets | 

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


## V1InstancesBackupsDelete

> Message V1InstancesBackupsDelete(ctx, backupDate, instanceId).Authorization(authorization).BackupTarget(backupTarget).Execute()





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
    backupDate := "backupDate_example" // string | 
    instanceId := "instanceId_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.
    backupTarget := "backupTarget_example" // string | Backup target selector

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupApi.V1InstancesBackupsDelete(context.Background(), backupDate, instanceId).Authorization(authorization).BackupTarget(backupTarget).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupApi.V1InstancesBackupsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesBackupsDelete`: Message
    fmt.Fprintf(os.Stdout, "Response from `BackupApi.V1InstancesBackupsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupDate** | **string** |  | 
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesBackupsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **backupTarget** | **string** | Backup target selector | 

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


## V1InstancesBackupsList

> BackupResponse V1InstancesBackupsList(ctx, instanceId).Authorization(authorization).BackupTarget(backupTarget).Execute()





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
    backupTarget := []string{"BackupTarget_example"} // []string | List of backup targets (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupApi.V1InstancesBackupsList(context.Background(), instanceId).Authorization(authorization).BackupTarget(backupTarget).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupApi.V1InstancesBackupsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesBackupsList`: BackupResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupApi.V1InstancesBackupsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesBackupsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **backupTarget** | **[]string** | List of backup targets | 

### Return type

[**BackupResponse**](BackupResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesBackupRestoresCreate

> Message V1ProjectsInstancesBackupRestoresCreate(ctx, backupDate, instanceId, projectId).Authorization(authorization).RestoreTarget(restoreTarget).Execute()





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
    backupDate := "backupDate_example" // string | 
    instanceId := "instanceId_example" // string | 
    projectId := "projectId_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.
    restoreTarget := "restoreTarget_example" // string | List of restore targets

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupApi.V1ProjectsInstancesBackupRestoresCreate(context.Background(), backupDate, instanceId, projectId).Authorization(authorization).RestoreTarget(restoreTarget).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupApi.V1ProjectsInstancesBackupRestoresCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesBackupRestoresCreate`: Message
    fmt.Fprintf(os.Stdout, "Response from `BackupApi.V1ProjectsInstancesBackupRestoresCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupDate** | **string** |  | 
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesBackupRestoresCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **restoreTarget** | **string** | List of restore targets | 

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


## V1ProjectsInstancesBackupRetentionsList

> BackupRetentionResponse V1ProjectsInstancesBackupRetentionsList(ctx, instanceId, projectId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.BackupApi.V1ProjectsInstancesBackupRetentionsList(context.Background(), instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupApi.V1ProjectsInstancesBackupRetentionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesBackupRetentionsList`: BackupRetentionResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupApi.V1ProjectsInstancesBackupRetentionsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesBackupRetentionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

### Return type

[**BackupRetentionResponse**](BackupRetentionResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesBackupRetentionsUpdate

> Message V1ProjectsInstancesBackupRetentionsUpdate(ctx, instanceId, projectId).Authorization(authorization).Data(data).BackupTarget(backupTarget).Execute()





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
    data := *openapiclient.NewV1InstancesBackupRetentionsUpdateRequest("Retention_example") // V1InstancesBackupRetentionsUpdateRequest | 
    backupTarget := []string{"BackupTarget_example"} // []string | List of backup targets (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupApi.V1ProjectsInstancesBackupRetentionsUpdate(context.Background(), instanceId, projectId).Authorization(authorization).Data(data).BackupTarget(backupTarget).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupApi.V1ProjectsInstancesBackupRetentionsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesBackupRetentionsUpdate`: Message
    fmt.Fprintf(os.Stdout, "Response from `BackupApi.V1ProjectsInstancesBackupRetentionsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesBackupRetentionsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesBackupRetentionsUpdateRequest**](V1InstancesBackupRetentionsUpdateRequest.md) |  | 
 **backupTarget** | **[]string** | List of backup targets | 

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


## V1ProjectsInstancesBackupSchedulesCreate

> BackupSchedulePostResponse V1ProjectsInstancesBackupSchedulesCreate(ctx, instanceId, projectId).Authorization(authorization).Data(data).BackupTarget(backupTarget).Execute()





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
    data := *openapiclient.NewV1InstancesBackupSchedulesCreateRequest("Schedule_example") // V1InstancesBackupSchedulesCreateRequest | 
    backupTarget := []string{"BackupTarget_example"} // []string | List of backup targets (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupApi.V1ProjectsInstancesBackupSchedulesCreate(context.Background(), instanceId, projectId).Authorization(authorization).Data(data).BackupTarget(backupTarget).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupApi.V1ProjectsInstancesBackupSchedulesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesBackupSchedulesCreate`: BackupSchedulePostResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupApi.V1ProjectsInstancesBackupSchedulesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesBackupSchedulesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesBackupSchedulesCreateRequest**](V1InstancesBackupSchedulesCreateRequest.md) |  | 
 **backupTarget** | **[]string** | List of backup targets | 

### Return type

[**BackupSchedulePostResponse**](BackupSchedulePostResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesBackupSchedulesDelete

> Message V1ProjectsInstancesBackupSchedulesDelete(ctx, instanceId, projectId, scheduleId).Authorization(authorization).Execute()





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
    scheduleId := "scheduleId_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupApi.V1ProjectsInstancesBackupSchedulesDelete(context.Background(), instanceId, projectId, scheduleId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupApi.V1ProjectsInstancesBackupSchedulesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesBackupSchedulesDelete`: Message
    fmt.Fprintf(os.Stdout, "Response from `BackupApi.V1ProjectsInstancesBackupSchedulesDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 
**scheduleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesBackupSchedulesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 

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


## V1ProjectsInstancesBackupSchedulesList

> BackupScheduleResponse V1ProjectsInstancesBackupSchedulesList(ctx, instanceId, projectId).Authorization(authorization).BackupTarget(backupTarget).Execute()





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
    backupTarget := []string{"BackupTarget_example"} // []string | List of backup targets (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupApi.V1ProjectsInstancesBackupSchedulesList(context.Background(), instanceId, projectId).Authorization(authorization).BackupTarget(backupTarget).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupApi.V1ProjectsInstancesBackupSchedulesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesBackupSchedulesList`: BackupScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupApi.V1ProjectsInstancesBackupSchedulesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesBackupSchedulesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **backupTarget** | **[]string** | List of backup targets | 

### Return type

[**BackupScheduleResponse**](BackupScheduleResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesBackupSchedulesUpdate

> BackupSchedulePutResponse V1ProjectsInstancesBackupSchedulesUpdate(ctx, instanceId, projectId, scheduleId).Authorization(authorization).Data(data).Execute()





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
    scheduleId := "scheduleId_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.
    data := *openapiclient.NewV1InstancesBackupSchedulesCreateRequest("Schedule_example") // V1InstancesBackupSchedulesCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupApi.V1ProjectsInstancesBackupSchedulesUpdate(context.Background(), instanceId, projectId, scheduleId).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupApi.V1ProjectsInstancesBackupSchedulesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesBackupSchedulesUpdate`: BackupSchedulePutResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupApi.V1ProjectsInstancesBackupSchedulesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 
**scheduleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesBackupSchedulesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **data** | [**V1InstancesBackupSchedulesCreateRequest**](V1InstancesBackupSchedulesCreateRequest.md) |  | 

### Return type

[**BackupSchedulePutResponse**](BackupSchedulePutResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesBackupsCreate

> Message V1ProjectsInstancesBackupsCreate(ctx, instanceId, projectId).Authorization(authorization).BackupTarget(backupTarget).Execute()





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
    backupTarget := []string{"BackupTarget_example"} // []string | List of backup targets (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupApi.V1ProjectsInstancesBackupsCreate(context.Background(), instanceId, projectId).Authorization(authorization).BackupTarget(backupTarget).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupApi.V1ProjectsInstancesBackupsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesBackupsCreate`: Message
    fmt.Fprintf(os.Stdout, "Response from `BackupApi.V1ProjectsInstancesBackupsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesBackupsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **backupTarget** | **[]string** | List of backup targets | 

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


## V1ProjectsInstancesBackupsDelete

> Message V1ProjectsInstancesBackupsDelete(ctx, backupDate, instanceId, projectId).Authorization(authorization).BackupTarget(backupTarget).Execute()





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
    backupDate := "backupDate_example" // string | 
    instanceId := "instanceId_example" // string | 
    projectId := "projectId_example" // string | 
    authorization := "authorization_example" // string | Accepts technical credentials and api gateway access.
    backupTarget := "backupTarget_example" // string | Backup target selector

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupApi.V1ProjectsInstancesBackupsDelete(context.Background(), backupDate, instanceId, projectId).Authorization(authorization).BackupTarget(backupTarget).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupApi.V1ProjectsInstancesBackupsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesBackupsDelete`: Message
    fmt.Fprintf(os.Stdout, "Response from `BackupApi.V1ProjectsInstancesBackupsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupDate** | **string** |  | 
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesBackupsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **backupTarget** | **string** | Backup target selector | 

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


## V1ProjectsInstancesBackupsList

> BackupResponse V1ProjectsInstancesBackupsList(ctx, instanceId, projectId).Authorization(authorization).BackupTarget(backupTarget).Execute()





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
    backupTarget := []string{"BackupTarget_example"} // []string | List of backup targets (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BackupApi.V1ProjectsInstancesBackupsList(context.Background(), instanceId, projectId).Authorization(authorization).BackupTarget(backupTarget).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupApi.V1ProjectsInstancesBackupsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesBackupsList`: BackupResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupApi.V1ProjectsInstancesBackupsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesBackupsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts technical credentials and api gateway access. | 
 **backupTarget** | **[]string** | List of backup targets | 

### Return type

[**BackupResponse**](BackupResponse.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

