# V1InstancesLogsConfigsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retention** | **string** | How long to keep the logs &#x60;Additional Validators:&#x60; * Should be a valid time string * Should not be longer than 30 days | 

## Methods

### NewV1InstancesLogsConfigsUpdateRequest

`func NewV1InstancesLogsConfigsUpdateRequest(retention string, ) *V1InstancesLogsConfigsUpdateRequest`

NewV1InstancesLogsConfigsUpdateRequest instantiates a new V1InstancesLogsConfigsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesLogsConfigsUpdateRequestWithDefaults

`func NewV1InstancesLogsConfigsUpdateRequestWithDefaults() *V1InstancesLogsConfigsUpdateRequest`

NewV1InstancesLogsConfigsUpdateRequestWithDefaults instantiates a new V1InstancesLogsConfigsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetention

`func (o *V1InstancesLogsConfigsUpdateRequest) GetRetention() string`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *V1InstancesLogsConfigsUpdateRequest) GetRetentionOk() (*string, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *V1InstancesLogsConfigsUpdateRequest) SetRetention(v string)`

SetRetention sets Retention field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


