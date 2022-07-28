# BackupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**AlertConfigBackups** | **[]string** |  | 
**AlertRulesBackups** | **[]string** |  | 
**ScrapeConfigBackups** | **[]string** |  | 
**GrafanaBackups** | **[]string** |  | 

## Methods

### NewBackupResponse

`func NewBackupResponse(message string, alertConfigBackups []string, alertRulesBackups []string, scrapeConfigBackups []string, grafanaBackups []string, ) *BackupResponse`

NewBackupResponse instantiates a new BackupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupResponseWithDefaults

`func NewBackupResponseWithDefaults() *BackupResponse`

NewBackupResponseWithDefaults instantiates a new BackupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *BackupResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BackupResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BackupResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetAlertConfigBackups

`func (o *BackupResponse) GetAlertConfigBackups() []string`

GetAlertConfigBackups returns the AlertConfigBackups field if non-nil, zero value otherwise.

### GetAlertConfigBackupsOk

`func (o *BackupResponse) GetAlertConfigBackupsOk() (*[]string, bool)`

GetAlertConfigBackupsOk returns a tuple with the AlertConfigBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertConfigBackups

`func (o *BackupResponse) SetAlertConfigBackups(v []string)`

SetAlertConfigBackups sets AlertConfigBackups field to given value.


### GetAlertRulesBackups

`func (o *BackupResponse) GetAlertRulesBackups() []string`

GetAlertRulesBackups returns the AlertRulesBackups field if non-nil, zero value otherwise.

### GetAlertRulesBackupsOk

`func (o *BackupResponse) GetAlertRulesBackupsOk() (*[]string, bool)`

GetAlertRulesBackupsOk returns a tuple with the AlertRulesBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertRulesBackups

`func (o *BackupResponse) SetAlertRulesBackups(v []string)`

SetAlertRulesBackups sets AlertRulesBackups field to given value.


### GetScrapeConfigBackups

`func (o *BackupResponse) GetScrapeConfigBackups() []string`

GetScrapeConfigBackups returns the ScrapeConfigBackups field if non-nil, zero value otherwise.

### GetScrapeConfigBackupsOk

`func (o *BackupResponse) GetScrapeConfigBackupsOk() (*[]string, bool)`

GetScrapeConfigBackupsOk returns a tuple with the ScrapeConfigBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrapeConfigBackups

`func (o *BackupResponse) SetScrapeConfigBackups(v []string)`

SetScrapeConfigBackups sets ScrapeConfigBackups field to given value.


### GetGrafanaBackups

`func (o *BackupResponse) GetGrafanaBackups() []string`

GetGrafanaBackups returns the GrafanaBackups field if non-nil, zero value otherwise.

### GetGrafanaBackupsOk

`func (o *BackupResponse) GetGrafanaBackupsOk() (*[]string, bool)`

GetGrafanaBackupsOk returns a tuple with the GrafanaBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaBackups

`func (o *BackupResponse) SetGrafanaBackups(v []string)`

SetGrafanaBackups sets GrafanaBackups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


