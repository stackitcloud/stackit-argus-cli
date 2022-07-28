# V1InstancesMetricsStorageRetentionsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricsRetentionTimeRaw** | **string** | Retention time of longtime storage of raw sampled data. After that time the data will be down sampled to 5m. Keep in mind, that the initial goal of downsampling is not saving disk or object storage space. In fact, downsampling doesn&#39;t save you any space but instead, it adds 2 more blocks for each raw block which are only slightly smaller or relatively similar size to raw block. This is done by internal downsampling implementation which to be mathematically correct holds various aggregations. This means that downsampling can increase the size of your storage a bit (~3x), if you choose to store all resolutions (recommended). The goal of downsampling is to provide an opportunity to get fast results for range queries of big time intervals like months or years. &#x60;Additional Validators:&#x60; * Should be a valid time string * Should not be bigger than 13 months | 
**MetricsRetentionTime5m** | **string** | Retention time of longtime storage of 5m sampled data. After that time the data will be down sampled to 1h. &#x60;Additional Validators:&#x60; * Should be a valid time string * Should not be bigger than metricsRetentionTimeRaw | 
**MetricsRetentionTime1h** | **string** | Retention time of longtime storage of 1h sampled data. After that time the data will be deleted permanently. &#x60;Additional Validators:&#x60; * Should be a valid time string * Should not be bigger than metricsRetentionTime5m | 

## Methods

### NewV1InstancesMetricsStorageRetentionsUpdateRequest

`func NewV1InstancesMetricsStorageRetentionsUpdateRequest(metricsRetentionTimeRaw string, metricsRetentionTime5m string, metricsRetentionTime1h string, ) *V1InstancesMetricsStorageRetentionsUpdateRequest`

NewV1InstancesMetricsStorageRetentionsUpdateRequest instantiates a new V1InstancesMetricsStorageRetentionsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesMetricsStorageRetentionsUpdateRequestWithDefaults

`func NewV1InstancesMetricsStorageRetentionsUpdateRequestWithDefaults() *V1InstancesMetricsStorageRetentionsUpdateRequest`

NewV1InstancesMetricsStorageRetentionsUpdateRequestWithDefaults instantiates a new V1InstancesMetricsStorageRetentionsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricsRetentionTimeRaw

`func (o *V1InstancesMetricsStorageRetentionsUpdateRequest) GetMetricsRetentionTimeRaw() string`

GetMetricsRetentionTimeRaw returns the MetricsRetentionTimeRaw field if non-nil, zero value otherwise.

### GetMetricsRetentionTimeRawOk

`func (o *V1InstancesMetricsStorageRetentionsUpdateRequest) GetMetricsRetentionTimeRawOk() (*string, bool)`

GetMetricsRetentionTimeRawOk returns a tuple with the MetricsRetentionTimeRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsRetentionTimeRaw

`func (o *V1InstancesMetricsStorageRetentionsUpdateRequest) SetMetricsRetentionTimeRaw(v string)`

SetMetricsRetentionTimeRaw sets MetricsRetentionTimeRaw field to given value.


### GetMetricsRetentionTime5m

`func (o *V1InstancesMetricsStorageRetentionsUpdateRequest) GetMetricsRetentionTime5m() string`

GetMetricsRetentionTime5m returns the MetricsRetentionTime5m field if non-nil, zero value otherwise.

### GetMetricsRetentionTime5mOk

`func (o *V1InstancesMetricsStorageRetentionsUpdateRequest) GetMetricsRetentionTime5mOk() (*string, bool)`

GetMetricsRetentionTime5mOk returns a tuple with the MetricsRetentionTime5m field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsRetentionTime5m

`func (o *V1InstancesMetricsStorageRetentionsUpdateRequest) SetMetricsRetentionTime5m(v string)`

SetMetricsRetentionTime5m sets MetricsRetentionTime5m field to given value.


### GetMetricsRetentionTime1h

`func (o *V1InstancesMetricsStorageRetentionsUpdateRequest) GetMetricsRetentionTime1h() string`

GetMetricsRetentionTime1h returns the MetricsRetentionTime1h field if non-nil, zero value otherwise.

### GetMetricsRetentionTime1hOk

`func (o *V1InstancesMetricsStorageRetentionsUpdateRequest) GetMetricsRetentionTime1hOk() (*string, bool)`

GetMetricsRetentionTime1hOk returns a tuple with the MetricsRetentionTime1h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsRetentionTime1h

`func (o *V1InstancesMetricsStorageRetentionsUpdateRequest) SetMetricsRetentionTime1h(v string)`

SetMetricsRetentionTime1h sets MetricsRetentionTime1h field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


