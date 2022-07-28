# V1OsbClustersInstancesUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** | String array of tags | [optional] 
**ServiceName** | Pointer to **string** | Name of the service | [optional] 
**OrgId** | Pointer to **string** | Marketplace organization id | [optional] 
**MetricsRetentionTimeRaw** | Pointer to **string** | Retention time of longtime storage of raw sampled data. Any value &lt; 1d or wrong time string format will result in 14d. After that time the data will be down sampled to 5m. | [optional] 
**MetricsRetentionTime5m** | Pointer to **string** | Retention time of longtime storage of 5m sampled data. Any value &lt; 1d or wrong time string format will result in 0d. After that time the data will be down sampled to 1h. | [optional] 
**MetricsRetentionTime1h** | Pointer to **string** | Retention time of longtime storage of 1h sampled data. Any value &lt; 1d or wrong time string format will result in 0d. After that time the data will be deleted permanently. | [optional] 

## Methods

### NewV1OsbClustersInstancesUpdateRequest

`func NewV1OsbClustersInstancesUpdateRequest() *V1OsbClustersInstancesUpdateRequest`

NewV1OsbClustersInstancesUpdateRequest instantiates a new V1OsbClustersInstancesUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1OsbClustersInstancesUpdateRequestWithDefaults

`func NewV1OsbClustersInstancesUpdateRequestWithDefaults() *V1OsbClustersInstancesUpdateRequest`

NewV1OsbClustersInstancesUpdateRequestWithDefaults instantiates a new V1OsbClustersInstancesUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *V1OsbClustersInstancesUpdateRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *V1OsbClustersInstancesUpdateRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *V1OsbClustersInstancesUpdateRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *V1OsbClustersInstancesUpdateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetServiceName

`func (o *V1OsbClustersInstancesUpdateRequest) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *V1OsbClustersInstancesUpdateRequest) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *V1OsbClustersInstancesUpdateRequest) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *V1OsbClustersInstancesUpdateRequest) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetOrgId

`func (o *V1OsbClustersInstancesUpdateRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *V1OsbClustersInstancesUpdateRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *V1OsbClustersInstancesUpdateRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *V1OsbClustersInstancesUpdateRequest) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetMetricsRetentionTimeRaw

`func (o *V1OsbClustersInstancesUpdateRequest) GetMetricsRetentionTimeRaw() string`

GetMetricsRetentionTimeRaw returns the MetricsRetentionTimeRaw field if non-nil, zero value otherwise.

### GetMetricsRetentionTimeRawOk

`func (o *V1OsbClustersInstancesUpdateRequest) GetMetricsRetentionTimeRawOk() (*string, bool)`

GetMetricsRetentionTimeRawOk returns a tuple with the MetricsRetentionTimeRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsRetentionTimeRaw

`func (o *V1OsbClustersInstancesUpdateRequest) SetMetricsRetentionTimeRaw(v string)`

SetMetricsRetentionTimeRaw sets MetricsRetentionTimeRaw field to given value.

### HasMetricsRetentionTimeRaw

`func (o *V1OsbClustersInstancesUpdateRequest) HasMetricsRetentionTimeRaw() bool`

HasMetricsRetentionTimeRaw returns a boolean if a field has been set.

### GetMetricsRetentionTime5m

`func (o *V1OsbClustersInstancesUpdateRequest) GetMetricsRetentionTime5m() string`

GetMetricsRetentionTime5m returns the MetricsRetentionTime5m field if non-nil, zero value otherwise.

### GetMetricsRetentionTime5mOk

`func (o *V1OsbClustersInstancesUpdateRequest) GetMetricsRetentionTime5mOk() (*string, bool)`

GetMetricsRetentionTime5mOk returns a tuple with the MetricsRetentionTime5m field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsRetentionTime5m

`func (o *V1OsbClustersInstancesUpdateRequest) SetMetricsRetentionTime5m(v string)`

SetMetricsRetentionTime5m sets MetricsRetentionTime5m field to given value.

### HasMetricsRetentionTime5m

`func (o *V1OsbClustersInstancesUpdateRequest) HasMetricsRetentionTime5m() bool`

HasMetricsRetentionTime5m returns a boolean if a field has been set.

### GetMetricsRetentionTime1h

`func (o *V1OsbClustersInstancesUpdateRequest) GetMetricsRetentionTime1h() string`

GetMetricsRetentionTime1h returns the MetricsRetentionTime1h field if non-nil, zero value otherwise.

### GetMetricsRetentionTime1hOk

`func (o *V1OsbClustersInstancesUpdateRequest) GetMetricsRetentionTime1hOk() (*string, bool)`

GetMetricsRetentionTime1hOk returns a tuple with the MetricsRetentionTime1h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsRetentionTime1h

`func (o *V1OsbClustersInstancesUpdateRequest) SetMetricsRetentionTime1h(v string)`

SetMetricsRetentionTime1h sets MetricsRetentionTime1h field to given value.

### HasMetricsRetentionTime1h

`func (o *V1OsbClustersInstancesUpdateRequest) HasMetricsRetentionTime1h() bool`

HasMetricsRetentionTime1h returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


