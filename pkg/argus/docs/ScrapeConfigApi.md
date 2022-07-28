# \ScrapeConfigApi

All URIs are relative to *https://api.argus.eu01.stackit.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1InstancesScrapeconfigsAllDelete**](ScrapeConfigApi.md#V1InstancesScrapeconfigsAllDelete) | **Delete** /v1/instances/{instanceId}/scrapeconfigs | 
[**V1InstancesScrapeconfigsCreate**](ScrapeConfigApi.md#V1InstancesScrapeconfigsCreate) | **Post** /v1/instances/{instanceId}/scrapeconfigs | 
[**V1InstancesScrapeconfigsDelete**](ScrapeConfigApi.md#V1InstancesScrapeconfigsDelete) | **Delete** /v1/instances/{instanceId}/scrapeconfigs/{jobName} | 
[**V1InstancesScrapeconfigsList**](ScrapeConfigApi.md#V1InstancesScrapeconfigsList) | **Get** /v1/instances/{instanceId}/scrapeconfigs | 
[**V1InstancesScrapeconfigsPartialUpdate**](ScrapeConfigApi.md#V1InstancesScrapeconfigsPartialUpdate) | **Patch** /v1/instances/{instanceId}/scrapeconfigs | 
[**V1InstancesScrapeconfigsRead**](ScrapeConfigApi.md#V1InstancesScrapeconfigsRead) | **Get** /v1/instances/{instanceId}/scrapeconfigs/{jobName} | 
[**V1InstancesScrapeconfigsUpdate**](ScrapeConfigApi.md#V1InstancesScrapeconfigsUpdate) | **Put** /v1/instances/{instanceId}/scrapeconfigs/{jobName} | 
[**V1ProjectsInstancesScrapeconfigsAllDelete**](ScrapeConfigApi.md#V1ProjectsInstancesScrapeconfigsAllDelete) | **Delete** /v1/projects/{projectId}/instances/{instanceId}/scrapeconfigs | 
[**V1ProjectsInstancesScrapeconfigsCreate**](ScrapeConfigApi.md#V1ProjectsInstancesScrapeconfigsCreate) | **Post** /v1/projects/{projectId}/instances/{instanceId}/scrapeconfigs | 
[**V1ProjectsInstancesScrapeconfigsDelete**](ScrapeConfigApi.md#V1ProjectsInstancesScrapeconfigsDelete) | **Delete** /v1/projects/{projectId}/instances/{instanceId}/scrapeconfigs/{jobName} | 
[**V1ProjectsInstancesScrapeconfigsList**](ScrapeConfigApi.md#V1ProjectsInstancesScrapeconfigsList) | **Get** /v1/projects/{projectId}/instances/{instanceId}/scrapeconfigs | 
[**V1ProjectsInstancesScrapeconfigsPartialUpdate**](ScrapeConfigApi.md#V1ProjectsInstancesScrapeconfigsPartialUpdate) | **Patch** /v1/projects/{projectId}/instances/{instanceId}/scrapeconfigs | 
[**V1ProjectsInstancesScrapeconfigsRead**](ScrapeConfigApi.md#V1ProjectsInstancesScrapeconfigsRead) | **Get** /v1/projects/{projectId}/instances/{instanceId}/scrapeconfigs/{jobName} | 
[**V1ProjectsInstancesScrapeconfigsUpdate**](ScrapeConfigApi.md#V1ProjectsInstancesScrapeconfigsUpdate) | **Put** /v1/projects/{projectId}/instances/{instanceId}/scrapeconfigs/{jobName} | 



## V1InstancesScrapeconfigsAllDelete

> CreateJob V1InstancesScrapeconfigsAllDelete(ctx, instanceId).Authorization(authorization).JobName(jobName).Execute()





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
    jobName := []string{"Inner_example"} // []string | Name of the jobs that should be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScrapeConfigApi.V1InstancesScrapeconfigsAllDelete(context.Background(), instanceId).Authorization(authorization).JobName(jobName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScrapeConfigApi.V1InstancesScrapeconfigsAllDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesScrapeconfigsAllDelete`: CreateJob
    fmt.Fprintf(os.Stdout, "Response from `ScrapeConfigApi.V1InstancesScrapeconfigsAllDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesScrapeconfigsAllDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 
 **jobName** | **[]string** | Name of the jobs that should be deleted | 

### Return type

[**CreateJob**](CreateJob.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesScrapeconfigsCreate

> CreateJob V1InstancesScrapeconfigsCreate(ctx, instanceId).Authorization(authorization).V1InstancesScrapeconfigsCreateRequest(v1InstancesScrapeconfigsCreateRequest).Execute()





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
    v1InstancesScrapeconfigsCreateRequest := *openapiclient.NewV1InstancesScrapeconfigsCreateRequest([]openapiclient.V1InstancesScrapeconfigsCreateRequestStaticConfigsInner{*openapiclient.NewV1InstancesScrapeconfigsCreateRequestStaticConfigsInner([]string{"Targets_example"})}, "JobName_example", "Scheme_example", "ScrapeInterval_example", "ScrapeTimeout_example", "MetricsPath_example") // V1InstancesScrapeconfigsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScrapeConfigApi.V1InstancesScrapeconfigsCreate(context.Background(), instanceId).Authorization(authorization).V1InstancesScrapeconfigsCreateRequest(v1InstancesScrapeconfigsCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScrapeConfigApi.V1InstancesScrapeconfigsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesScrapeconfigsCreate`: CreateJob
    fmt.Fprintf(os.Stdout, "Response from `ScrapeConfigApi.V1InstancesScrapeconfigsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesScrapeconfigsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 
 **v1InstancesScrapeconfigsCreateRequest** | [**V1InstancesScrapeconfigsCreateRequest**](V1InstancesScrapeconfigsCreateRequest.md) |  | 

### Return type

[**CreateJob**](CreateJob.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesScrapeconfigsDelete

> DeleteJob V1InstancesScrapeconfigsDelete(ctx, instanceId, jobName).Authorization(authorization).Execute()





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
    jobName := "jobName_example" // string | 
    authorization := "authorization_example" // string | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScrapeConfigApi.V1InstancesScrapeconfigsDelete(context.Background(), instanceId, jobName).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScrapeConfigApi.V1InstancesScrapeconfigsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesScrapeconfigsDelete`: DeleteJob
    fmt.Fprintf(os.Stdout, "Response from `ScrapeConfigApi.V1InstancesScrapeconfigsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**jobName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesScrapeconfigsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 

### Return type

[**DeleteJob**](DeleteJob.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesScrapeconfigsList

> GetAllJob V1InstancesScrapeconfigsList(ctx, instanceId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.ScrapeConfigApi.V1InstancesScrapeconfigsList(context.Background(), instanceId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScrapeConfigApi.V1InstancesScrapeconfigsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesScrapeconfigsList`: GetAllJob
    fmt.Fprintf(os.Stdout, "Response from `ScrapeConfigApi.V1InstancesScrapeconfigsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesScrapeconfigsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 

### Return type

[**GetAllJob**](GetAllJob.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesScrapeconfigsPartialUpdate

> CreateJob V1InstancesScrapeconfigsPartialUpdate(ctx, instanceId).Authorization(authorization).V1InstancesScrapeconfigsCreateRequest(v1InstancesScrapeconfigsCreateRequest).Execute()





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
    v1InstancesScrapeconfigsCreateRequest := []openapiclient.V1InstancesScrapeconfigsCreateRequest{*openapiclient.NewV1InstancesScrapeconfigsCreateRequest([]openapiclient.V1InstancesScrapeconfigsCreateRequestStaticConfigsInner{*openapiclient.NewV1InstancesScrapeconfigsCreateRequestStaticConfigsInner([]string{"Targets_example"})}, "JobName_example", "Scheme_example", "ScrapeInterval_example", "ScrapeTimeout_example", "MetricsPath_example")} // []V1InstancesScrapeconfigsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScrapeConfigApi.V1InstancesScrapeconfigsPartialUpdate(context.Background(), instanceId).Authorization(authorization).V1InstancesScrapeconfigsCreateRequest(v1InstancesScrapeconfigsCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScrapeConfigApi.V1InstancesScrapeconfigsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesScrapeconfigsPartialUpdate`: CreateJob
    fmt.Fprintf(os.Stdout, "Response from `ScrapeConfigApi.V1InstancesScrapeconfigsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesScrapeconfigsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 
 **v1InstancesScrapeconfigsCreateRequest** | [**[]V1InstancesScrapeconfigsCreateRequest**](V1InstancesScrapeconfigsCreateRequest.md) |  | 

### Return type

[**CreateJob**](CreateJob.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesScrapeconfigsRead

> GetJob V1InstancesScrapeconfigsRead(ctx, instanceId, jobName).Authorization(authorization).Execute()





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
    jobName := "jobName_example" // string | 
    authorization := "authorization_example" // string | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScrapeConfigApi.V1InstancesScrapeconfigsRead(context.Background(), instanceId, jobName).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScrapeConfigApi.V1InstancesScrapeconfigsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesScrapeconfigsRead`: GetJob
    fmt.Fprintf(os.Stdout, "Response from `ScrapeConfigApi.V1InstancesScrapeconfigsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**jobName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesScrapeconfigsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 

### Return type

[**GetJob**](GetJob.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1InstancesScrapeconfigsUpdate

> CreateJob V1InstancesScrapeconfigsUpdate(ctx, instanceId, jobName).Authorization(authorization).V1InstancesScrapeconfigsCreateRequest(v1InstancesScrapeconfigsCreateRequest).Execute()





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
    jobName := "jobName_example" // string | 
    authorization := "authorization_example" // string | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user
    v1InstancesScrapeconfigsCreateRequest := *openapiclient.NewV1InstancesScrapeconfigsCreateRequest([]openapiclient.V1InstancesScrapeconfigsCreateRequestStaticConfigsInner{*openapiclient.NewV1InstancesScrapeconfigsCreateRequestStaticConfigsInner([]string{"Targets_example"})}, "JobName_example", "Scheme_example", "ScrapeInterval_example", "ScrapeTimeout_example", "MetricsPath_example") // V1InstancesScrapeconfigsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScrapeConfigApi.V1InstancesScrapeconfigsUpdate(context.Background(), instanceId, jobName).Authorization(authorization).V1InstancesScrapeconfigsCreateRequest(v1InstancesScrapeconfigsCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScrapeConfigApi.V1InstancesScrapeconfigsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1InstancesScrapeconfigsUpdate`: CreateJob
    fmt.Fprintf(os.Stdout, "Response from `ScrapeConfigApi.V1InstancesScrapeconfigsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**jobName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1InstancesScrapeconfigsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 
 **v1InstancesScrapeconfigsCreateRequest** | [**V1InstancesScrapeconfigsCreateRequest**](V1InstancesScrapeconfigsCreateRequest.md) |  | 

### Return type

[**CreateJob**](CreateJob.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesScrapeconfigsAllDelete

> CreateJob V1ProjectsInstancesScrapeconfigsAllDelete(ctx, instanceId, projectId).Authorization(authorization).JobName(jobName).Execute()





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
    jobName := []string{"Inner_example"} // []string | Name of the jobs that should be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScrapeConfigApi.V1ProjectsInstancesScrapeconfigsAllDelete(context.Background(), instanceId, projectId).Authorization(authorization).JobName(jobName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScrapeConfigApi.V1ProjectsInstancesScrapeconfigsAllDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesScrapeconfigsAllDelete`: CreateJob
    fmt.Fprintf(os.Stdout, "Response from `ScrapeConfigApi.V1ProjectsInstancesScrapeconfigsAllDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesScrapeconfigsAllDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 
 **jobName** | **[]string** | Name of the jobs that should be deleted | 

### Return type

[**CreateJob**](CreateJob.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesScrapeconfigsCreate

> CreateJob V1ProjectsInstancesScrapeconfigsCreate(ctx, instanceId, projectId).Authorization(authorization).V1InstancesScrapeconfigsCreateRequest(v1InstancesScrapeconfigsCreateRequest).Execute()





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
    v1InstancesScrapeconfigsCreateRequest := *openapiclient.NewV1InstancesScrapeconfigsCreateRequest([]openapiclient.V1InstancesScrapeconfigsCreateRequestStaticConfigsInner{*openapiclient.NewV1InstancesScrapeconfigsCreateRequestStaticConfigsInner([]string{"Targets_example"})}, "JobName_example", "Scheme_example", "ScrapeInterval_example", "ScrapeTimeout_example", "MetricsPath_example") // V1InstancesScrapeconfigsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScrapeConfigApi.V1ProjectsInstancesScrapeconfigsCreate(context.Background(), instanceId, projectId).Authorization(authorization).V1InstancesScrapeconfigsCreateRequest(v1InstancesScrapeconfigsCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScrapeConfigApi.V1ProjectsInstancesScrapeconfigsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesScrapeconfigsCreate`: CreateJob
    fmt.Fprintf(os.Stdout, "Response from `ScrapeConfigApi.V1ProjectsInstancesScrapeconfigsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesScrapeconfigsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 
 **v1InstancesScrapeconfigsCreateRequest** | [**V1InstancesScrapeconfigsCreateRequest**](V1InstancesScrapeconfigsCreateRequest.md) |  | 

### Return type

[**CreateJob**](CreateJob.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesScrapeconfigsDelete

> DeleteJob V1ProjectsInstancesScrapeconfigsDelete(ctx, instanceId, jobName, projectId).Authorization(authorization).Execute()





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
    jobName := "jobName_example" // string | 
    projectId := "projectId_example" // string | 
    authorization := "authorization_example" // string | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScrapeConfigApi.V1ProjectsInstancesScrapeconfigsDelete(context.Background(), instanceId, jobName, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScrapeConfigApi.V1ProjectsInstancesScrapeconfigsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesScrapeconfigsDelete`: DeleteJob
    fmt.Fprintf(os.Stdout, "Response from `ScrapeConfigApi.V1ProjectsInstancesScrapeconfigsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**jobName** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesScrapeconfigsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 

### Return type

[**DeleteJob**](DeleteJob.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesScrapeconfigsList

> GetAllJob V1ProjectsInstancesScrapeconfigsList(ctx, instanceId, projectId).Authorization(authorization).Execute()





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
    resp, r, err := apiClient.ScrapeConfigApi.V1ProjectsInstancesScrapeconfigsList(context.Background(), instanceId, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScrapeConfigApi.V1ProjectsInstancesScrapeconfigsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesScrapeconfigsList`: GetAllJob
    fmt.Fprintf(os.Stdout, "Response from `ScrapeConfigApi.V1ProjectsInstancesScrapeconfigsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesScrapeconfigsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 

### Return type

[**GetAllJob**](GetAllJob.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesScrapeconfigsPartialUpdate

> CreateJob V1ProjectsInstancesScrapeconfigsPartialUpdate(ctx, instanceId, projectId).Authorization(authorization).V1InstancesScrapeconfigsCreateRequest(v1InstancesScrapeconfigsCreateRequest).Execute()





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
    v1InstancesScrapeconfigsCreateRequest := []openapiclient.V1InstancesScrapeconfigsCreateRequest{*openapiclient.NewV1InstancesScrapeconfigsCreateRequest([]openapiclient.V1InstancesScrapeconfigsCreateRequestStaticConfigsInner{*openapiclient.NewV1InstancesScrapeconfigsCreateRequestStaticConfigsInner([]string{"Targets_example"})}, "JobName_example", "Scheme_example", "ScrapeInterval_example", "ScrapeTimeout_example", "MetricsPath_example")} // []V1InstancesScrapeconfigsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScrapeConfigApi.V1ProjectsInstancesScrapeconfigsPartialUpdate(context.Background(), instanceId, projectId).Authorization(authorization).V1InstancesScrapeconfigsCreateRequest(v1InstancesScrapeconfigsCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScrapeConfigApi.V1ProjectsInstancesScrapeconfigsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesScrapeconfigsPartialUpdate`: CreateJob
    fmt.Fprintf(os.Stdout, "Response from `ScrapeConfigApi.V1ProjectsInstancesScrapeconfigsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesScrapeconfigsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 
 **v1InstancesScrapeconfigsCreateRequest** | [**[]V1InstancesScrapeconfigsCreateRequest**](V1InstancesScrapeconfigsCreateRequest.md) |  | 

### Return type

[**CreateJob**](CreateJob.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesScrapeconfigsRead

> GetJob V1ProjectsInstancesScrapeconfigsRead(ctx, instanceId, jobName, projectId).Authorization(authorization).Execute()





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
    jobName := "jobName_example" // string | 
    projectId := "projectId_example" // string | 
    authorization := "authorization_example" // string | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScrapeConfigApi.V1ProjectsInstancesScrapeconfigsRead(context.Background(), instanceId, jobName, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScrapeConfigApi.V1ProjectsInstancesScrapeconfigsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesScrapeconfigsRead`: GetJob
    fmt.Fprintf(os.Stdout, "Response from `ScrapeConfigApi.V1ProjectsInstancesScrapeconfigsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**jobName** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesScrapeconfigsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 

### Return type

[**GetJob**](GetJob.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ProjectsInstancesScrapeconfigsUpdate

> CreateJob V1ProjectsInstancesScrapeconfigsUpdate(ctx, instanceId, jobName, projectId).Authorization(authorization).V1InstancesScrapeconfigsCreateRequest(v1InstancesScrapeconfigsCreateRequest).Execute()





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
    jobName := "jobName_example" // string | 
    projectId := "projectId_example" // string | 
    authorization := "authorization_example" // string | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user
    v1InstancesScrapeconfigsCreateRequest := *openapiclient.NewV1InstancesScrapeconfigsCreateRequest([]openapiclient.V1InstancesScrapeconfigsCreateRequestStaticConfigsInner{*openapiclient.NewV1InstancesScrapeconfigsCreateRequestStaticConfigsInner([]string{"Targets_example"})}, "JobName_example", "Scheme_example", "ScrapeInterval_example", "ScrapeTimeout_example", "MetricsPath_example") // V1InstancesScrapeconfigsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScrapeConfigApi.V1ProjectsInstancesScrapeconfigsUpdate(context.Background(), instanceId, jobName, projectId).Authorization(authorization).V1InstancesScrapeconfigsCreateRequest(v1InstancesScrapeconfigsCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScrapeConfigApi.V1ProjectsInstancesScrapeconfigsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1ProjectsInstancesScrapeconfigsUpdate`: CreateJob
    fmt.Fprintf(os.Stdout, "Response from `ScrapeConfigApi.V1ProjectsInstancesScrapeconfigsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**jobName** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1ProjectsInstancesScrapeconfigsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | Accepts basic auth and bearer token from admins and jwt token from oauth and basic auth from api user | 
 **v1InstancesScrapeconfigsCreateRequest** | [**V1InstancesScrapeconfigsCreateRequest**](V1InstancesScrapeconfigsCreateRequest.md) |  | 

### Return type

[**CreateJob**](CreateJob.md)

### Authorization

[Basic](../README.md#Basic), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

