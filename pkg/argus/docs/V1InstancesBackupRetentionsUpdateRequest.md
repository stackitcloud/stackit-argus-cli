# V1InstancesBackupRetentionsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retention** | **string** | How long the backups should be stored. &#x60;Additional Validators:&#x60; * must be a valid time string | 

## Methods

### NewV1InstancesBackupRetentionsUpdateRequest

`func NewV1InstancesBackupRetentionsUpdateRequest(retention string, ) *V1InstancesBackupRetentionsUpdateRequest`

NewV1InstancesBackupRetentionsUpdateRequest instantiates a new V1InstancesBackupRetentionsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesBackupRetentionsUpdateRequestWithDefaults

`func NewV1InstancesBackupRetentionsUpdateRequestWithDefaults() *V1InstancesBackupRetentionsUpdateRequest`

NewV1InstancesBackupRetentionsUpdateRequestWithDefaults instantiates a new V1InstancesBackupRetentionsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetention

`func (o *V1InstancesBackupRetentionsUpdateRequest) GetRetention() string`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *V1InstancesBackupRetentionsUpdateRequest) GetRetentionOk() (*string, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *V1InstancesBackupRetentionsUpdateRequest) SetRetention(v string)`

SetRetention sets Retention field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


