# V1InstancesTracesConfigsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retention** | **string** | How long to keep the traces &#x60;Additional Validators:&#x60; * Should be a valid time string * Should not be bigger than 30 days | 

## Methods

### NewV1InstancesTracesConfigsUpdateRequest

`func NewV1InstancesTracesConfigsUpdateRequest(retention string, ) *V1InstancesTracesConfigsUpdateRequest`

NewV1InstancesTracesConfigsUpdateRequest instantiates a new V1InstancesTracesConfigsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesTracesConfigsUpdateRequestWithDefaults

`func NewV1InstancesTracesConfigsUpdateRequestWithDefaults() *V1InstancesTracesConfigsUpdateRequest`

NewV1InstancesTracesConfigsUpdateRequestWithDefaults instantiates a new V1InstancesTracesConfigsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetention

`func (o *V1InstancesTracesConfigsUpdateRequest) GetRetention() string`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *V1InstancesTracesConfigsUpdateRequest) GetRetentionOk() (*string, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *V1InstancesTracesConfigsUpdateRequest) SetRetention(v string)`

SetRetention sets Retention field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


