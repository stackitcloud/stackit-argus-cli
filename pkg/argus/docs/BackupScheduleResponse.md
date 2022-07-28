# BackupScheduleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**AlertConfigBackupSchedules** | Pointer to [**[]BackupScheduleModelJson**](BackupScheduleModelJson.md) |  | [optional] 
**AlertRulesBackupSchedules** | Pointer to [**[]BackupScheduleModelJson**](BackupScheduleModelJson.md) |  | [optional] 
**ScrapeConfigBackupSchedules** | Pointer to [**[]BackupScheduleModelJson**](BackupScheduleModelJson.md) |  | [optional] 
**GrafanaBackupSchedules** | Pointer to [**[]BackupScheduleModelJson**](BackupScheduleModelJson.md) |  | [optional] 

## Methods

### NewBackupScheduleResponse

`func NewBackupScheduleResponse(message string, ) *BackupScheduleResponse`

NewBackupScheduleResponse instantiates a new BackupScheduleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupScheduleResponseWithDefaults

`func NewBackupScheduleResponseWithDefaults() *BackupScheduleResponse`

NewBackupScheduleResponseWithDefaults instantiates a new BackupScheduleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *BackupScheduleResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BackupScheduleResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BackupScheduleResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetAlertConfigBackupSchedules

`func (o *BackupScheduleResponse) GetAlertConfigBackupSchedules() []BackupScheduleModelJson`

GetAlertConfigBackupSchedules returns the AlertConfigBackupSchedules field if non-nil, zero value otherwise.

### GetAlertConfigBackupSchedulesOk

`func (o *BackupScheduleResponse) GetAlertConfigBackupSchedulesOk() (*[]BackupScheduleModelJson, bool)`

GetAlertConfigBackupSchedulesOk returns a tuple with the AlertConfigBackupSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertConfigBackupSchedules

`func (o *BackupScheduleResponse) SetAlertConfigBackupSchedules(v []BackupScheduleModelJson)`

SetAlertConfigBackupSchedules sets AlertConfigBackupSchedules field to given value.

### HasAlertConfigBackupSchedules

`func (o *BackupScheduleResponse) HasAlertConfigBackupSchedules() bool`

HasAlertConfigBackupSchedules returns a boolean if a field has been set.

### GetAlertRulesBackupSchedules

`func (o *BackupScheduleResponse) GetAlertRulesBackupSchedules() []BackupScheduleModelJson`

GetAlertRulesBackupSchedules returns the AlertRulesBackupSchedules field if non-nil, zero value otherwise.

### GetAlertRulesBackupSchedulesOk

`func (o *BackupScheduleResponse) GetAlertRulesBackupSchedulesOk() (*[]BackupScheduleModelJson, bool)`

GetAlertRulesBackupSchedulesOk returns a tuple with the AlertRulesBackupSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertRulesBackupSchedules

`func (o *BackupScheduleResponse) SetAlertRulesBackupSchedules(v []BackupScheduleModelJson)`

SetAlertRulesBackupSchedules sets AlertRulesBackupSchedules field to given value.

### HasAlertRulesBackupSchedules

`func (o *BackupScheduleResponse) HasAlertRulesBackupSchedules() bool`

HasAlertRulesBackupSchedules returns a boolean if a field has been set.

### GetScrapeConfigBackupSchedules

`func (o *BackupScheduleResponse) GetScrapeConfigBackupSchedules() []BackupScheduleModelJson`

GetScrapeConfigBackupSchedules returns the ScrapeConfigBackupSchedules field if non-nil, zero value otherwise.

### GetScrapeConfigBackupSchedulesOk

`func (o *BackupScheduleResponse) GetScrapeConfigBackupSchedulesOk() (*[]BackupScheduleModelJson, bool)`

GetScrapeConfigBackupSchedulesOk returns a tuple with the ScrapeConfigBackupSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrapeConfigBackupSchedules

`func (o *BackupScheduleResponse) SetScrapeConfigBackupSchedules(v []BackupScheduleModelJson)`

SetScrapeConfigBackupSchedules sets ScrapeConfigBackupSchedules field to given value.

### HasScrapeConfigBackupSchedules

`func (o *BackupScheduleResponse) HasScrapeConfigBackupSchedules() bool`

HasScrapeConfigBackupSchedules returns a boolean if a field has been set.

### GetGrafanaBackupSchedules

`func (o *BackupScheduleResponse) GetGrafanaBackupSchedules() []BackupScheduleModelJson`

GetGrafanaBackupSchedules returns the GrafanaBackupSchedules field if non-nil, zero value otherwise.

### GetGrafanaBackupSchedulesOk

`func (o *BackupScheduleResponse) GetGrafanaBackupSchedulesOk() (*[]BackupScheduleModelJson, bool)`

GetGrafanaBackupSchedulesOk returns a tuple with the GrafanaBackupSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaBackupSchedules

`func (o *BackupScheduleResponse) SetGrafanaBackupSchedules(v []BackupScheduleModelJson)`

SetGrafanaBackupSchedules sets GrafanaBackupSchedules field to given value.

### HasGrafanaBackupSchedules

`func (o *BackupScheduleResponse) HasGrafanaBackupSchedules() bool`

HasGrafanaBackupSchedules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


