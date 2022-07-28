# Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | **string** |  | 
**Instance** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Plan** | [**PlanModel**](PlanModel.md) |  | 
**MetricsRetentionTimeRaw** | **int32** |  | 
**MetricsRetentionTime5m** | **int32** |  | 
**MetricsRetentionTime1h** | **int32** |  | 
**State** | Pointer to **string** |  | [optional] 
**GrafanaPublicReadAccess** | **bool** |  | 

## Methods

### NewInstance

`func NewInstance(cluster string, instance string, plan PlanModel, metricsRetentionTimeRaw int32, metricsRetentionTime5m int32, metricsRetentionTime1h int32, grafanaPublicReadAccess bool, ) *Instance`

NewInstance instantiates a new Instance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceWithDefaults

`func NewInstanceWithDefaults() *Instance`

NewInstanceWithDefaults instantiates a new Instance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *Instance) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Instance) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Instance) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetInstance

`func (o *Instance) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *Instance) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *Instance) SetInstance(v string)`

SetInstance sets Instance field to given value.


### GetName

`func (o *Instance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Instance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Instance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Instance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlan

`func (o *Instance) GetPlan() PlanModel`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Instance) GetPlanOk() (*PlanModel, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Instance) SetPlan(v PlanModel)`

SetPlan sets Plan field to given value.


### GetMetricsRetentionTimeRaw

`func (o *Instance) GetMetricsRetentionTimeRaw() int32`

GetMetricsRetentionTimeRaw returns the MetricsRetentionTimeRaw field if non-nil, zero value otherwise.

### GetMetricsRetentionTimeRawOk

`func (o *Instance) GetMetricsRetentionTimeRawOk() (*int32, bool)`

GetMetricsRetentionTimeRawOk returns a tuple with the MetricsRetentionTimeRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsRetentionTimeRaw

`func (o *Instance) SetMetricsRetentionTimeRaw(v int32)`

SetMetricsRetentionTimeRaw sets MetricsRetentionTimeRaw field to given value.


### GetMetricsRetentionTime5m

`func (o *Instance) GetMetricsRetentionTime5m() int32`

GetMetricsRetentionTime5m returns the MetricsRetentionTime5m field if non-nil, zero value otherwise.

### GetMetricsRetentionTime5mOk

`func (o *Instance) GetMetricsRetentionTime5mOk() (*int32, bool)`

GetMetricsRetentionTime5mOk returns a tuple with the MetricsRetentionTime5m field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsRetentionTime5m

`func (o *Instance) SetMetricsRetentionTime5m(v int32)`

SetMetricsRetentionTime5m sets MetricsRetentionTime5m field to given value.


### GetMetricsRetentionTime1h

`func (o *Instance) GetMetricsRetentionTime1h() int32`

GetMetricsRetentionTime1h returns the MetricsRetentionTime1h field if non-nil, zero value otherwise.

### GetMetricsRetentionTime1hOk

`func (o *Instance) GetMetricsRetentionTime1hOk() (*int32, bool)`

GetMetricsRetentionTime1hOk returns a tuple with the MetricsRetentionTime1h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsRetentionTime1h

`func (o *Instance) SetMetricsRetentionTime1h(v int32)`

SetMetricsRetentionTime1h sets MetricsRetentionTime1h field to given value.


### GetState

`func (o *Instance) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Instance) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Instance) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Instance) HasState() bool`

HasState returns a boolean if a field has been set.

### GetGrafanaPublicReadAccess

`func (o *Instance) GetGrafanaPublicReadAccess() bool`

GetGrafanaPublicReadAccess returns the GrafanaPublicReadAccess field if non-nil, zero value otherwise.

### GetGrafanaPublicReadAccessOk

`func (o *Instance) GetGrafanaPublicReadAccessOk() (*bool, bool)`

GetGrafanaPublicReadAccessOk returns a tuple with the GrafanaPublicReadAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaPublicReadAccess

`func (o *Instance) SetGrafanaPublicReadAccess(v bool)`

SetGrafanaPublicReadAccess sets GrafanaPublicReadAccess field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


