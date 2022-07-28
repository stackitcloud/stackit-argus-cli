# LogsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retention** | **string** |  | 

## Methods

### NewLogsConfig

`func NewLogsConfig(retention string, ) *LogsConfig`

NewLogsConfig instantiates a new LogsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsConfigWithDefaults

`func NewLogsConfigWithDefaults() *LogsConfig`

NewLogsConfigWithDefaults instantiates a new LogsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetention

`func (o *LogsConfig) GetRetention() string`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *LogsConfig) GetRetentionOk() (*string, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *LogsConfig) SetRetention(v string)`

SetRetention sets Retention field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


