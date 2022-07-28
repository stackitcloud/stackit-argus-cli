# TracesConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Config** | [**TraceConfig**](TraceConfig.md) |  | 

## Methods

### NewTracesConfigResponse

`func NewTracesConfigResponse(message string, config TraceConfig, ) *TracesConfigResponse`

NewTracesConfigResponse instantiates a new TracesConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTracesConfigResponseWithDefaults

`func NewTracesConfigResponseWithDefaults() *TracesConfigResponse`

NewTracesConfigResponseWithDefaults instantiates a new TracesConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *TracesConfigResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TracesConfigResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TracesConfigResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetConfig

`func (o *TracesConfigResponse) GetConfig() TraceConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *TracesConfigResponse) GetConfigOk() (*TraceConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *TracesConfigResponse) SetConfig(v TraceConfig)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


