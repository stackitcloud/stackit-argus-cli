# GrafanaPlugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Plugins** | [**[]GrafanaPluginModel**](GrafanaPluginModel.md) |  | 

## Methods

### NewGrafanaPlugin

`func NewGrafanaPlugin(message string, plugins []GrafanaPluginModel, ) *GrafanaPlugin`

NewGrafanaPlugin instantiates a new GrafanaPlugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrafanaPluginWithDefaults

`func NewGrafanaPluginWithDefaults() *GrafanaPlugin`

NewGrafanaPluginWithDefaults instantiates a new GrafanaPlugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GrafanaPlugin) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GrafanaPlugin) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GrafanaPlugin) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetPlugins

`func (o *GrafanaPlugin) GetPlugins() []GrafanaPluginModel`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *GrafanaPlugin) GetPluginsOk() (*[]GrafanaPluginModel, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *GrafanaPlugin) SetPlugins(v []GrafanaPluginModel)`

SetPlugins sets Plugins field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


