# LogsConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Config** | [**LogsConfig**](LogsConfig.md) |  | 

## Methods

### NewLogsConfigResponse

`func NewLogsConfigResponse(message string, config LogsConfig, ) *LogsConfigResponse`

NewLogsConfigResponse instantiates a new LogsConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsConfigResponseWithDefaults

`func NewLogsConfigResponseWithDefaults() *LogsConfigResponse`

NewLogsConfigResponseWithDefaults instantiates a new LogsConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *LogsConfigResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LogsConfigResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LogsConfigResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetConfig

`func (o *LogsConfigResponse) GetConfig() LogsConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *LogsConfigResponse) GetConfigOk() (*LogsConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *LogsConfigResponse) SetConfig(v LogsConfig)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


