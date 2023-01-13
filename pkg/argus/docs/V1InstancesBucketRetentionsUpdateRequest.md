# V1InstancesBucketRetentionsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketRetentionTimeRaw** | **string** | Retention time of longtime storage of raw sampled data. Any value &lt; 1d or wrong time string format will result in 14d. After that time the data will be down sampled to 5m. | 
**BucketRetentionTime5m** | **string** | Retention time of longtime storage of 5m sampled data. Any value &lt; 1d or wrong time string format will result in 0d. After that time the data will be down sampled to 1h. | 
**BucketRetentionTime1h** | **string** | Retention time of longtime storage of 1h sampled data. Any value &lt; 1d or wrong time string format will result in 0d. After that time the data will be deleted permanently. | 

## Methods

### NewV1InstancesBucketRetentionsUpdateRequest

`func NewV1InstancesBucketRetentionsUpdateRequest(bucketRetentionTimeRaw string, bucketRetentionTime5m string, bucketRetentionTime1h string, ) *V1InstancesBucketRetentionsUpdateRequest`

NewV1InstancesBucketRetentionsUpdateRequest instantiates a new V1InstancesBucketRetentionsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesBucketRetentionsUpdateRequestWithDefaults

`func NewV1InstancesBucketRetentionsUpdateRequestWithDefaults() *V1InstancesBucketRetentionsUpdateRequest`

NewV1InstancesBucketRetentionsUpdateRequestWithDefaults instantiates a new V1InstancesBucketRetentionsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketRetentionTimeRaw

`func (o *V1InstancesBucketRetentionsUpdateRequest) GetBucketRetentionTimeRaw() string`

GetBucketRetentionTimeRaw returns the BucketRetentionTimeRaw field if non-nil, zero value otherwise.

### GetBucketRetentionTimeRawOk

`func (o *V1InstancesBucketRetentionsUpdateRequest) GetBucketRetentionTimeRawOk() (*string, bool)`

GetBucketRetentionTimeRawOk returns a tuple with the BucketRetentionTimeRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketRetentionTimeRaw

`func (o *V1InstancesBucketRetentionsUpdateRequest) SetBucketRetentionTimeRaw(v string)`

SetBucketRetentionTimeRaw sets BucketRetentionTimeRaw field to given value.


### GetBucketRetentionTime5m

`func (o *V1InstancesBucketRetentionsUpdateRequest) GetBucketRetentionTime5m() string`

GetBucketRetentionTime5m returns the BucketRetentionTime5m field if non-nil, zero value otherwise.

### GetBucketRetentionTime5mOk

`func (o *V1InstancesBucketRetentionsUpdateRequest) GetBucketRetentionTime5mOk() (*string, bool)`

GetBucketRetentionTime5mOk returns a tuple with the BucketRetentionTime5m field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketRetentionTime5m

`func (o *V1InstancesBucketRetentionsUpdateRequest) SetBucketRetentionTime5m(v string)`

SetBucketRetentionTime5m sets BucketRetentionTime5m field to given value.


### GetBucketRetentionTime1h

`func (o *V1InstancesBucketRetentionsUpdateRequest) GetBucketRetentionTime1h() string`

GetBucketRetentionTime1h returns the BucketRetentionTime1h field if non-nil, zero value otherwise.

### GetBucketRetentionTime1hOk

`func (o *V1InstancesBucketRetentionsUpdateRequest) GetBucketRetentionTime1hOk() (*string, bool)`

GetBucketRetentionTime1hOk returns a tuple with the BucketRetentionTime1h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketRetentionTime1h

`func (o *V1InstancesBucketRetentionsUpdateRequest) SetBucketRetentionTime1h(v string)`

SetBucketRetentionTime1h sets BucketRetentionTime1h field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


