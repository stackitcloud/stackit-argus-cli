# BackupSchedulePutResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Schedule** | [**BackupScheduleModelJson**](BackupScheduleModelJson.md) |  | 

## Methods

### NewBackupSchedulePutResponse

`func NewBackupSchedulePutResponse(message string, schedule BackupScheduleModelJson, ) *BackupSchedulePutResponse`

NewBackupSchedulePutResponse instantiates a new BackupSchedulePutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupSchedulePutResponseWithDefaults

`func NewBackupSchedulePutResponseWithDefaults() *BackupSchedulePutResponse`

NewBackupSchedulePutResponseWithDefaults instantiates a new BackupSchedulePutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *BackupSchedulePutResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BackupSchedulePutResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BackupSchedulePutResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSchedule

`func (o *BackupSchedulePutResponse) GetSchedule() BackupScheduleModelJson`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *BackupSchedulePutResponse) GetScheduleOk() (*BackupScheduleModelJson, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *BackupSchedulePutResponse) SetSchedule(v BackupScheduleModelJson)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


