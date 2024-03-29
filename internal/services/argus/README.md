# Go API client for argus

API endpoints for Argus on STACKIT

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import argus "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), argus.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), argus.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), argus.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), argus.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.argus.eu01.stackit.cloud*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AlertConfigApi* | [**V1InstancesAlertconfigsList**](docs/AlertConfigApi.md#v1instancesalertconfigslist) | **Get** /v1/instances/{instanceId}/alertconfigs | 
*AlertConfigApi* | [**V1InstancesAlertconfigsUpdate**](docs/AlertConfigApi.md#v1instancesalertconfigsupdate) | **Put** /v1/instances/{instanceId}/alertconfigs | 
*AlertConfigApi* | [**V1ProjectsInstancesAlertconfigsList**](docs/AlertConfigApi.md#v1projectsinstancesalertconfigslist) | **Get** /v1/projects/{projectId}/instances/{instanceId}/alertconfigs | 
*AlertConfigApi* | [**V1ProjectsInstancesAlertconfigsUpdate**](docs/AlertConfigApi.md#v1projectsinstancesalertconfigsupdate) | **Put** /v1/projects/{projectId}/instances/{instanceId}/alertconfigs | 
*AlertGroupsApi* | [**V1InstancesAlertgroupsAllDelete**](docs/AlertGroupsApi.md#v1instancesalertgroupsalldelete) | **Delete** /v1/instances/{instanceId}/alertgroups | 
*AlertGroupsApi* | [**V1InstancesAlertgroupsCreate**](docs/AlertGroupsApi.md#v1instancesalertgroupscreate) | **Post** /v1/instances/{instanceId}/alertgroups | 
*AlertGroupsApi* | [**V1InstancesAlertgroupsDelete**](docs/AlertGroupsApi.md#v1instancesalertgroupsdelete) | **Delete** /v1/instances/{instanceId}/alertgroups/{groupName} | 
*AlertGroupsApi* | [**V1InstancesAlertgroupsList**](docs/AlertGroupsApi.md#v1instancesalertgroupslist) | **Get** /v1/instances/{instanceId}/alertgroups | 
*AlertGroupsApi* | [**V1InstancesAlertgroupsPartialUpdate**](docs/AlertGroupsApi.md#v1instancesalertgroupspartialupdate) | **Patch** /v1/instances/{instanceId}/alertgroups | 
*AlertGroupsApi* | [**V1InstancesAlertgroupsRead**](docs/AlertGroupsApi.md#v1instancesalertgroupsread) | **Get** /v1/instances/{instanceId}/alertgroups/{groupName} | 
*AlertGroupsApi* | [**V1InstancesAlertgroupsUpdate**](docs/AlertGroupsApi.md#v1instancesalertgroupsupdate) | **Put** /v1/instances/{instanceId}/alertgroups/{groupName} | 
*AlertGroupsApi* | [**V1ProjectsInstancesAlertgroupsAllDelete**](docs/AlertGroupsApi.md#v1projectsinstancesalertgroupsalldelete) | **Delete** /v1/projects/{projectId}/instances/{instanceId}/alertgroups | 
*AlertGroupsApi* | [**V1ProjectsInstancesAlertgroupsCreate**](docs/AlertGroupsApi.md#v1projectsinstancesalertgroupscreate) | **Post** /v1/projects/{projectId}/instances/{instanceId}/alertgroups | 
*AlertGroupsApi* | [**V1ProjectsInstancesAlertgroupsDelete**](docs/AlertGroupsApi.md#v1projectsinstancesalertgroupsdelete) | **Delete** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName} | 
*AlertGroupsApi* | [**V1ProjectsInstancesAlertgroupsList**](docs/AlertGroupsApi.md#v1projectsinstancesalertgroupslist) | **Get** /v1/projects/{projectId}/instances/{instanceId}/alertgroups | 
*AlertGroupsApi* | [**V1ProjectsInstancesAlertgroupsPartialUpdate**](docs/AlertGroupsApi.md#v1projectsinstancesalertgroupspartialupdate) | **Patch** /v1/projects/{projectId}/instances/{instanceId}/alertgroups | 
*AlertGroupsApi* | [**V1ProjectsInstancesAlertgroupsRead**](docs/AlertGroupsApi.md#v1projectsinstancesalertgroupsread) | **Get** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName} | 
*AlertGroupsApi* | [**V1ProjectsInstancesAlertgroupsUpdate**](docs/AlertGroupsApi.md#v1projectsinstancesalertgroupsupdate) | **Put** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName} | 
*AlertRulesApi* | [**V1InstancesAlertgroupsAlertrulesAllDelete**](docs/AlertRulesApi.md#v1instancesalertgroupsalertrulesalldelete) | **Delete** /v1/instances/{instanceId}/alertgroups/{groupName}/alertrules | 
*AlertRulesApi* | [**V1InstancesAlertgroupsAlertrulesCreate**](docs/AlertRulesApi.md#v1instancesalertgroupsalertrulescreate) | **Post** /v1/instances/{instanceId}/alertgroups/{groupName}/alertrules | 
*AlertRulesApi* | [**V1InstancesAlertgroupsAlertrulesDelete**](docs/AlertRulesApi.md#v1instancesalertgroupsalertrulesdelete) | **Delete** /v1/instances/{instanceId}/alertgroups/{groupName}/alertrules/{alertName} | 
*AlertRulesApi* | [**V1InstancesAlertgroupsAlertrulesList**](docs/AlertRulesApi.md#v1instancesalertgroupsalertruleslist) | **Get** /v1/instances/{instanceId}/alertgroups/{groupName}/alertrules | 
*AlertRulesApi* | [**V1InstancesAlertgroupsAlertrulesPartialUpdate**](docs/AlertRulesApi.md#v1instancesalertgroupsalertrulespartialupdate) | **Patch** /v1/instances/{instanceId}/alertgroups/{groupName}/alertrules | 
*AlertRulesApi* | [**V1InstancesAlertgroupsAlertrulesRead**](docs/AlertRulesApi.md#v1instancesalertgroupsalertrulesread) | **Get** /v1/instances/{instanceId}/alertgroups/{groupName}/alertrules/{alertName} | 
*AlertRulesApi* | [**V1InstancesAlertgroupsAlertrulesUpdate**](docs/AlertRulesApi.md#v1instancesalertgroupsalertrulesupdate) | **Put** /v1/instances/{instanceId}/alertgroups/{groupName}/alertrules/{alertName} | 
*AlertRulesApi* | [**V1ProjectsInstancesAlertgroupsAlertrulesAllDelete**](docs/AlertRulesApi.md#v1projectsinstancesalertgroupsalertrulesalldelete) | **Delete** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/alertrules | 
*AlertRulesApi* | [**V1ProjectsInstancesAlertgroupsAlertrulesCreate**](docs/AlertRulesApi.md#v1projectsinstancesalertgroupsalertrulescreate) | **Post** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/alertrules | 
*AlertRulesApi* | [**V1ProjectsInstancesAlertgroupsAlertrulesDelete**](docs/AlertRulesApi.md#v1projectsinstancesalertgroupsalertrulesdelete) | **Delete** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/alertrules/{alertName} | 
*AlertRulesApi* | [**V1ProjectsInstancesAlertgroupsAlertrulesList**](docs/AlertRulesApi.md#v1projectsinstancesalertgroupsalertruleslist) | **Get** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/alertrules | 
*AlertRulesApi* | [**V1ProjectsInstancesAlertgroupsAlertrulesPartialUpdate**](docs/AlertRulesApi.md#v1projectsinstancesalertgroupsalertrulespartialupdate) | **Patch** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/alertrules | 
*AlertRulesApi* | [**V1ProjectsInstancesAlertgroupsAlertrulesRead**](docs/AlertRulesApi.md#v1projectsinstancesalertgroupsalertrulesread) | **Get** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/alertrules/{alertName} | 
*AlertRulesApi* | [**V1ProjectsInstancesAlertgroupsAlertrulesUpdate**](docs/AlertRulesApi.md#v1projectsinstancesalertgroupsalertrulesupdate) | **Put** /v1/projects/{projectId}/instances/{instanceId}/alertgroups/{groupName}/alertrules/{alertName} | 
*BucketRetentionApi* | [**V1InstancesBucketRetentionsList**](docs/BucketRetentionApi.md#v1instancesbucketretentionslist) | **Get** /v1/instances/{instanceId}/bucket-retentions | 
*BucketRetentionApi* | [**V1InstancesBucketRetentionsUpdate**](docs/BucketRetentionApi.md#v1instancesbucketretentionsupdate) | **Put** /v1/instances/{instanceId}/bucket-retentions | 
*BucketRetentionApi* | [**V1ProjectsInstancesBucketRetentionsList**](docs/BucketRetentionApi.md#v1projectsinstancesbucketretentionslist) | **Get** /v1/projects/{projectId}/instances/{instanceId}/bucket-retentions | 
*BucketRetentionApi* | [**V1ProjectsInstancesBucketRetentionsUpdate**](docs/BucketRetentionApi.md#v1projectsinstancesbucketretentionsupdate) | **Put** /v1/projects/{projectId}/instances/{instanceId}/bucket-retentions | 
*GrafanaConfigsApi* | [**V1InstancesGrafanaConfigsList**](docs/GrafanaConfigsApi.md#v1instancesgrafanaconfigslist) | **Get** /v1/instances/{instanceId}/grafana-configs | 
*GrafanaConfigsApi* | [**V1InstancesGrafanaConfigsUpdate**](docs/GrafanaConfigsApi.md#v1instancesgrafanaconfigsupdate) | **Put** /v1/instances/{instanceId}/grafana-configs | 
*GrafanaConfigsApi* | [**V1ProjectsInstancesGrafanaConfigsList**](docs/GrafanaConfigsApi.md#v1projectsinstancesgrafanaconfigslist) | **Get** /v1/projects/{projectId}/instances/{instanceId}/grafana-configs | 
*GrafanaConfigsApi* | [**V1ProjectsInstancesGrafanaConfigsUpdate**](docs/GrafanaConfigsApi.md#v1projectsinstancesgrafanaconfigsupdate) | **Put** /v1/projects/{projectId}/instances/{instanceId}/grafana-configs | 
*InstancesApi* | [**V1InstancesRead**](docs/InstancesApi.md#v1instancesread) | **Get** /v1/instances/{instanceId} | 
*InstancesApi* | [**V1ProjectsInstancesRead**](docs/InstancesApi.md#v1projectsinstancesread) | **Get** /v1/projects/{projectId}/instances/{instanceId} | 
*ScrapeConfigApi* | [**V1InstancesScrapeconfigsAllDelete**](docs/ScrapeConfigApi.md#v1instancesscrapeconfigsalldelete) | **Delete** /v1/instances/{instanceId}/scrapeconfigs | 
*ScrapeConfigApi* | [**V1InstancesScrapeconfigsCreate**](docs/ScrapeConfigApi.md#v1instancesscrapeconfigscreate) | **Post** /v1/instances/{instanceId}/scrapeconfigs | 
*ScrapeConfigApi* | [**V1InstancesScrapeconfigsDelete**](docs/ScrapeConfigApi.md#v1instancesscrapeconfigsdelete) | **Delete** /v1/instances/{instanceId}/scrapeconfigs/{jobName} | 
*ScrapeConfigApi* | [**V1InstancesScrapeconfigsList**](docs/ScrapeConfigApi.md#v1instancesscrapeconfigslist) | **Get** /v1/instances/{instanceId}/scrapeconfigs | 
*ScrapeConfigApi* | [**V1InstancesScrapeconfigsPartialUpdate**](docs/ScrapeConfigApi.md#v1instancesscrapeconfigspartialupdate) | **Patch** /v1/instances/{instanceId}/scrapeconfigs | 
*ScrapeConfigApi* | [**V1InstancesScrapeconfigsRead**](docs/ScrapeConfigApi.md#v1instancesscrapeconfigsread) | **Get** /v1/instances/{instanceId}/scrapeconfigs/{jobName} | 
*ScrapeConfigApi* | [**V1InstancesScrapeconfigsUpdate**](docs/ScrapeConfigApi.md#v1instancesscrapeconfigsupdate) | **Put** /v1/instances/{instanceId}/scrapeconfigs/{jobName} | 
*ScrapeConfigApi* | [**V1ProjectsInstancesScrapeconfigsAllDelete**](docs/ScrapeConfigApi.md#v1projectsinstancesscrapeconfigsalldelete) | **Delete** /v1/projects/{projectId}/instances/{instanceId}/scrapeconfigs | 
*ScrapeConfigApi* | [**V1ProjectsInstancesScrapeconfigsCreate**](docs/ScrapeConfigApi.md#v1projectsinstancesscrapeconfigscreate) | **Post** /v1/projects/{projectId}/instances/{instanceId}/scrapeconfigs | 
*ScrapeConfigApi* | [**V1ProjectsInstancesScrapeconfigsDelete**](docs/ScrapeConfigApi.md#v1projectsinstancesscrapeconfigsdelete) | **Delete** /v1/projects/{projectId}/instances/{instanceId}/scrapeconfigs/{jobName} | 
*ScrapeConfigApi* | [**V1ProjectsInstancesScrapeconfigsList**](docs/ScrapeConfigApi.md#v1projectsinstancesscrapeconfigslist) | **Get** /v1/projects/{projectId}/instances/{instanceId}/scrapeconfigs | 
*ScrapeConfigApi* | [**V1ProjectsInstancesScrapeconfigsPartialUpdate**](docs/ScrapeConfigApi.md#v1projectsinstancesscrapeconfigspartialupdate) | **Patch** /v1/projects/{projectId}/instances/{instanceId}/scrapeconfigs | 
*ScrapeConfigApi* | [**V1ProjectsInstancesScrapeconfigsRead**](docs/ScrapeConfigApi.md#v1projectsinstancesscrapeconfigsread) | **Get** /v1/projects/{projectId}/instances/{instanceId}/scrapeconfigs/{jobName} | 
*ScrapeConfigApi* | [**V1ProjectsInstancesScrapeconfigsUpdate**](docs/ScrapeConfigApi.md#v1projectsinstancesscrapeconfigsupdate) | **Put** /v1/projects/{projectId}/instances/{instanceId}/scrapeconfigs/{jobName} | 


## Documentation For Models

 - [Alert](docs/Alert.md)
 - [AlertGroup](docs/AlertGroup.md)
 - [AlertRule](docs/AlertRule.md)
 - [ApiUser](docs/ApiUser.md)
 - [BasicAuth](docs/BasicAuth.md)
 - [BucketRetentionTimeRespond](docs/BucketRetentionTimeRespond.md)
 - [ClusterList](docs/ClusterList.md)
 - [CreateJob](docs/CreateJob.md)
 - [Dashboard](docs/Dashboard.md)
 - [DeleteAlertGroup](docs/DeleteAlertGroup.md)
 - [DeleteAlertRule](docs/DeleteAlertRule.md)
 - [DeleteJob](docs/DeleteJob.md)
 - [EmailConfig](docs/EmailConfig.md)
 - [GetAlert](docs/GetAlert.md)
 - [GetAlertGroup](docs/GetAlertGroup.md)
 - [GetAlertRule](docs/GetAlertRule.md)
 - [GetAllAlertGroups](docs/GetAllAlertGroups.md)
 - [GetAllAlertRules](docs/GetAllAlertRules.md)
 - [GetAllJob](docs/GetAllJob.md)
 - [GetJob](docs/GetJob.md)
 - [Global](docs/Global.md)
 - [GrafanaConfigsSerializerRespond](docs/GrafanaConfigsSerializerRespond.md)
 - [HttpConfig](docs/HttpConfig.md)
 - [InhibitRules](docs/InhibitRules.md)
 - [Instance](docs/Instance.md)
 - [InstanceSensitiveData](docs/InstanceSensitiveData.md)
 - [Job](docs/Job.md)
 - [Message](docs/Message.md)
 - [MetricsRelabelConfig](docs/MetricsRelabelConfig.md)
 - [OpsgenieConfig](docs/OpsgenieConfig.md)
 - [PermissionDenied](docs/PermissionDenied.md)
 - [Plan](docs/Plan.md)
 - [PlanModel](docs/PlanModel.md)
 - [PlanSingle](docs/PlanSingle.md)
 - [PostAlertGroup](docs/PostAlertGroup.md)
 - [PostAlertRule](docs/PostAlertRule.md)
 - [PutAlert](docs/PutAlert.md)
 - [PutAlertGroup](docs/PutAlertGroup.md)
 - [PutAlertRule](docs/PutAlertRule.md)
 - [Receivers](docs/Receivers.md)
 - [Route](docs/Route.md)
 - [RouteSerializer2](docs/RouteSerializer2.md)
 - [Service](docs/Service.md)
 - [Services](docs/Services.md)
 - [StaticConfigs](docs/StaticConfigs.md)
 - [Status](docs/Status.md)
 - [TLSConfig](docs/TLSConfig.md)
 - [Tag](docs/Tag.md)
 - [V1InstancesAlertconfigsUpdateRequest](docs/V1InstancesAlertconfigsUpdateRequest.md)
 - [V1InstancesAlertconfigsUpdateRequestGlobal](docs/V1InstancesAlertconfigsUpdateRequestGlobal.md)
 - [V1InstancesAlertconfigsUpdateRequestInhibitRules](docs/V1InstancesAlertconfigsUpdateRequestInhibitRules.md)
 - [V1InstancesAlertconfigsUpdateRequestReceiversInner](docs/V1InstancesAlertconfigsUpdateRequestReceiversInner.md)
 - [V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner](docs/V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner.md)
 - [V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner](docs/V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner.md)
 - [V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig](docs/V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig.md)
 - [V1InstancesAlertconfigsUpdateRequestReceiversInnerWebHookConfigsInner](docs/V1InstancesAlertconfigsUpdateRequestReceiversInnerWebHookConfigsInner.md)
 - [V1InstancesAlertconfigsUpdateRequestRoute](docs/V1InstancesAlertconfigsUpdateRequestRoute.md)
 - [V1InstancesAlertconfigsUpdateRequestRouteRoutesInner](docs/V1InstancesAlertconfigsUpdateRequestRouteRoutesInner.md)
 - [V1InstancesAlertgroupsAlertrulesCreateRequest](docs/V1InstancesAlertgroupsAlertrulesCreateRequest.md)
 - [V1InstancesAlertgroupsCreateRequest](docs/V1InstancesAlertgroupsCreateRequest.md)
 - [V1InstancesAlertgroupsCreateRequestRulesInner](docs/V1InstancesAlertgroupsCreateRequestRulesInner.md)
 - [V1InstancesBucketRetentionsUpdateRequest](docs/V1InstancesBucketRetentionsUpdateRequest.md)
 - [V1InstancesGrafanaConfigsUpdateRequest](docs/V1InstancesGrafanaConfigsUpdateRequest.md)
 - [V1InstancesScrapeconfigsCreateRequest](docs/V1InstancesScrapeconfigsCreateRequest.md)
 - [V1InstancesScrapeconfigsCreateRequestBasicAuth](docs/V1InstancesScrapeconfigsCreateRequestBasicAuth.md)
 - [V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner](docs/V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner.md)
 - [V1InstancesScrapeconfigsCreateRequestParams](docs/V1InstancesScrapeconfigsCreateRequestParams.md)
 - [V1InstancesScrapeconfigsCreateRequestStaticConfigsInner](docs/V1InstancesScrapeconfigsCreateRequestStaticConfigsInner.md)
 - [V1InstancesScrapeconfigsCreateRequestTlsConfig](docs/V1InstancesScrapeconfigsCreateRequestTlsConfig.md)
 - [WebHook](docs/WebHook.md)


## Documentation For Authorization



### Basic

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

patrick.koss@mail.schwarz

