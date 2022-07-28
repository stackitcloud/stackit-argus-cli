# V1InstancesBackupSchedulesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | **string** | The schedule for how often to create a backup. &#x60;Additional Validators:&#x60; * must be a valid cronjob format * must run less than hourly | 

## Methods

### NewV1InstancesBackupSchedulesCreateRequest

`func NewV1InstancesBackupSchedulesCreateRequest(schedule string, ) *V1InstancesBackupSchedulesCreateRequest`

NewV1InstancesBackupSchedulesCreateRequest instantiates a new V1InstancesBackupSchedulesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesBackupSchedulesCreateRequestWithDefaults

`func NewV1InstancesBackupSchedulesCreateRequestWithDefaults() *V1InstancesBackupSchedulesCreateRequest`

NewV1InstancesBackupSchedulesCreateRequestWithDefaults instantiates a new V1InstancesBackupSchedulesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedule

`func (o *V1InstancesBackupSchedulesCreateRequest) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *V1InstancesBackupSchedulesCreateRequest) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *V1InstancesBackupSchedulesCreateRequest) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


