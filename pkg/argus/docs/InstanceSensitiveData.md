# InstanceSensitiveData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | **string** |  | 
**Instance** | **string** |  | 
**GrafanaUrl** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**DashboardUrl** | **string** |  | 
**GrafanaAdminPassword** | **string** |  | 
**GrafanaAdminUser** | **string** |  | 
**Plan** | [**PlanModel**](PlanModel.md) |  | 
**BucketRetentionTimeRaw** | **int32** |  | 
**BucketRetentionTime5m** | **int32** |  | 
**BucketRetentionTime1h** | **int32** |  | 
**MetricsUrl** | **string** |  | 
**State** | Pointer to **string** |  | [optional] 
**GrafanaPublicReadAccess** | **bool** |  | 
**TargetsUrl** | **string** |  | 
**AlertingUrl** | **string** |  | 

## Methods

### NewInstanceSensitiveData

`func NewInstanceSensitiveData(cluster string, instance string, grafanaUrl string, dashboardUrl string, grafanaAdminPassword string, grafanaAdminUser string, plan PlanModel, bucketRetentionTimeRaw int32, bucketRetentionTime5m int32, bucketRetentionTime1h int32, metricsUrl string, grafanaPublicReadAccess bool, targetsUrl string, alertingUrl string, ) *InstanceSensitiveData`

NewInstanceSensitiveData instantiates a new InstanceSensitiveData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceSensitiveDataWithDefaults

`func NewInstanceSensitiveDataWithDefaults() *InstanceSensitiveData`

NewInstanceSensitiveDataWithDefaults instantiates a new InstanceSensitiveData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *InstanceSensitiveData) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *InstanceSensitiveData) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *InstanceSensitiveData) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetInstance

`func (o *InstanceSensitiveData) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *InstanceSensitiveData) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *InstanceSensitiveData) SetInstance(v string)`

SetInstance sets Instance field to given value.


### GetGrafanaUrl

`func (o *InstanceSensitiveData) GetGrafanaUrl() string`

GetGrafanaUrl returns the GrafanaUrl field if non-nil, zero value otherwise.

### GetGrafanaUrlOk

`func (o *InstanceSensitiveData) GetGrafanaUrlOk() (*string, bool)`

GetGrafanaUrlOk returns a tuple with the GrafanaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaUrl

`func (o *InstanceSensitiveData) SetGrafanaUrl(v string)`

SetGrafanaUrl sets GrafanaUrl field to given value.


### GetName

`func (o *InstanceSensitiveData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceSensitiveData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceSensitiveData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceSensitiveData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDashboardUrl

`func (o *InstanceSensitiveData) GetDashboardUrl() string`

GetDashboardUrl returns the DashboardUrl field if non-nil, zero value otherwise.

### GetDashboardUrlOk

`func (o *InstanceSensitiveData) GetDashboardUrlOk() (*string, bool)`

GetDashboardUrlOk returns a tuple with the DashboardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardUrl

`func (o *InstanceSensitiveData) SetDashboardUrl(v string)`

SetDashboardUrl sets DashboardUrl field to given value.


### GetGrafanaAdminPassword

`func (o *InstanceSensitiveData) GetGrafanaAdminPassword() string`

GetGrafanaAdminPassword returns the GrafanaAdminPassword field if non-nil, zero value otherwise.

### GetGrafanaAdminPasswordOk

`func (o *InstanceSensitiveData) GetGrafanaAdminPasswordOk() (*string, bool)`

GetGrafanaAdminPasswordOk returns a tuple with the GrafanaAdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaAdminPassword

`func (o *InstanceSensitiveData) SetGrafanaAdminPassword(v string)`

SetGrafanaAdminPassword sets GrafanaAdminPassword field to given value.


### GetGrafanaAdminUser

`func (o *InstanceSensitiveData) GetGrafanaAdminUser() string`

GetGrafanaAdminUser returns the GrafanaAdminUser field if non-nil, zero value otherwise.

### GetGrafanaAdminUserOk

`func (o *InstanceSensitiveData) GetGrafanaAdminUserOk() (*string, bool)`

GetGrafanaAdminUserOk returns a tuple with the GrafanaAdminUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaAdminUser

`func (o *InstanceSensitiveData) SetGrafanaAdminUser(v string)`

SetGrafanaAdminUser sets GrafanaAdminUser field to given value.


### GetPlan

`func (o *InstanceSensitiveData) GetPlan() PlanModel`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *InstanceSensitiveData) GetPlanOk() (*PlanModel, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *InstanceSensitiveData) SetPlan(v PlanModel)`

SetPlan sets Plan field to given value.


### GetBucketRetentionTimeRaw

`func (o *InstanceSensitiveData) GetBucketRetentionTimeRaw() int32`

GetBucketRetentionTimeRaw returns the BucketRetentionTimeRaw field if non-nil, zero value otherwise.

### GetBucketRetentionTimeRawOk

`func (o *InstanceSensitiveData) GetBucketRetentionTimeRawOk() (*int32, bool)`

GetBucketRetentionTimeRawOk returns a tuple with the BucketRetentionTimeRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketRetentionTimeRaw

`func (o *InstanceSensitiveData) SetBucketRetentionTimeRaw(v int32)`

SetBucketRetentionTimeRaw sets BucketRetentionTimeRaw field to given value.


### GetBucketRetentionTime5m

`func (o *InstanceSensitiveData) GetBucketRetentionTime5m() int32`

GetBucketRetentionTime5m returns the BucketRetentionTime5m field if non-nil, zero value otherwise.

### GetBucketRetentionTime5mOk

`func (o *InstanceSensitiveData) GetBucketRetentionTime5mOk() (*int32, bool)`

GetBucketRetentionTime5mOk returns a tuple with the BucketRetentionTime5m field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketRetentionTime5m

`func (o *InstanceSensitiveData) SetBucketRetentionTime5m(v int32)`

SetBucketRetentionTime5m sets BucketRetentionTime5m field to given value.


### GetBucketRetentionTime1h

`func (o *InstanceSensitiveData) GetBucketRetentionTime1h() int32`

GetBucketRetentionTime1h returns the BucketRetentionTime1h field if non-nil, zero value otherwise.

### GetBucketRetentionTime1hOk

`func (o *InstanceSensitiveData) GetBucketRetentionTime1hOk() (*int32, bool)`

GetBucketRetentionTime1hOk returns a tuple with the BucketRetentionTime1h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketRetentionTime1h

`func (o *InstanceSensitiveData) SetBucketRetentionTime1h(v int32)`

SetBucketRetentionTime1h sets BucketRetentionTime1h field to given value.


### GetMetricsUrl

`func (o *InstanceSensitiveData) GetMetricsUrl() string`

GetMetricsUrl returns the MetricsUrl field if non-nil, zero value otherwise.

### GetMetricsUrlOk

`func (o *InstanceSensitiveData) GetMetricsUrlOk() (*string, bool)`

GetMetricsUrlOk returns a tuple with the MetricsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsUrl

`func (o *InstanceSensitiveData) SetMetricsUrl(v string)`

SetMetricsUrl sets MetricsUrl field to given value.


### GetState

`func (o *InstanceSensitiveData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InstanceSensitiveData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InstanceSensitiveData) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *InstanceSensitiveData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetGrafanaPublicReadAccess

`func (o *InstanceSensitiveData) GetGrafanaPublicReadAccess() bool`

GetGrafanaPublicReadAccess returns the GrafanaPublicReadAccess field if non-nil, zero value otherwise.

### GetGrafanaPublicReadAccessOk

`func (o *InstanceSensitiveData) GetGrafanaPublicReadAccessOk() (*bool, bool)`

GetGrafanaPublicReadAccessOk returns a tuple with the GrafanaPublicReadAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaPublicReadAccess

`func (o *InstanceSensitiveData) SetGrafanaPublicReadAccess(v bool)`

SetGrafanaPublicReadAccess sets GrafanaPublicReadAccess field to given value.


### GetTargetsUrl

`func (o *InstanceSensitiveData) GetTargetsUrl() string`

GetTargetsUrl returns the TargetsUrl field if non-nil, zero value otherwise.

### GetTargetsUrlOk

`func (o *InstanceSensitiveData) GetTargetsUrlOk() (*string, bool)`

GetTargetsUrlOk returns a tuple with the TargetsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetsUrl

`func (o *InstanceSensitiveData) SetTargetsUrl(v string)`

SetTargetsUrl sets TargetsUrl field to given value.


### GetAlertingUrl

`func (o *InstanceSensitiveData) GetAlertingUrl() string`

GetAlertingUrl returns the AlertingUrl field if non-nil, zero value otherwise.

### GetAlertingUrlOk

`func (o *InstanceSensitiveData) GetAlertingUrlOk() (*string, bool)`

GetAlertingUrlOk returns a tuple with the AlertingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingUrl

`func (o *InstanceSensitiveData) SetAlertingUrl(v string)`

SetAlertingUrl sets AlertingUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


