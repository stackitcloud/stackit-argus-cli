# GrafanaPluginSingle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Plugin** | [**GrafanaPluginModel**](GrafanaPluginModel.md) |  | 

## Methods

### NewGrafanaPluginSingle

`func NewGrafanaPluginSingle(message string, plugin GrafanaPluginModel, ) *GrafanaPluginSingle`

NewGrafanaPluginSingle instantiates a new GrafanaPluginSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrafanaPluginSingleWithDefaults

`func NewGrafanaPluginSingleWithDefaults() *GrafanaPluginSingle`

NewGrafanaPluginSingleWithDefaults instantiates a new GrafanaPluginSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GrafanaPluginSingle) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GrafanaPluginSingle) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GrafanaPluginSingle) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetPlugin

`func (o *GrafanaPluginSingle) GetPlugin() GrafanaPluginModel`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *GrafanaPluginSingle) GetPluginOk() (*GrafanaPluginModel, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *GrafanaPluginSingle) SetPlugin(v GrafanaPluginModel)`

SetPlugin sets Plugin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


