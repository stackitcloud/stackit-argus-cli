# Receivers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**EmailConfigs** | Pointer to [**[]EmailConfig**](EmailConfig.md) |  | [optional] 
**OpsgenieConfigs** | Pointer to [**[]OpsgenieConfig**](OpsgenieConfig.md) |  | [optional] 
**WebHookConfigs** | Pointer to [**[]WebHook**](WebHook.md) |  | [optional] 

## Methods

### NewReceivers

`func NewReceivers(name string, ) *Receivers`

NewReceivers instantiates a new Receivers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiversWithDefaults

`func NewReceiversWithDefaults() *Receivers`

NewReceiversWithDefaults instantiates a new Receivers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Receivers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Receivers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Receivers) SetName(v string)`

SetName sets Name field to given value.


### GetEmailConfigs

`func (o *Receivers) GetEmailConfigs() []EmailConfig`

GetEmailConfigs returns the EmailConfigs field if non-nil, zero value otherwise.

### GetEmailConfigsOk

`func (o *Receivers) GetEmailConfigsOk() (*[]EmailConfig, bool)`

GetEmailConfigsOk returns a tuple with the EmailConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConfigs

`func (o *Receivers) SetEmailConfigs(v []EmailConfig)`

SetEmailConfigs sets EmailConfigs field to given value.

### HasEmailConfigs

`func (o *Receivers) HasEmailConfigs() bool`

HasEmailConfigs returns a boolean if a field has been set.

### GetOpsgenieConfigs

`func (o *Receivers) GetOpsgenieConfigs() []OpsgenieConfig`

GetOpsgenieConfigs returns the OpsgenieConfigs field if non-nil, zero value otherwise.

### GetOpsgenieConfigsOk

`func (o *Receivers) GetOpsgenieConfigsOk() (*[]OpsgenieConfig, bool)`

GetOpsgenieConfigsOk returns a tuple with the OpsgenieConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsgenieConfigs

`func (o *Receivers) SetOpsgenieConfigs(v []OpsgenieConfig)`

SetOpsgenieConfigs sets OpsgenieConfigs field to given value.

### HasOpsgenieConfigs

`func (o *Receivers) HasOpsgenieConfigs() bool`

HasOpsgenieConfigs returns a boolean if a field has been set.

### GetWebHookConfigs

`func (o *Receivers) GetWebHookConfigs() []WebHook`

GetWebHookConfigs returns the WebHookConfigs field if non-nil, zero value otherwise.

### GetWebHookConfigsOk

`func (o *Receivers) GetWebHookConfigsOk() (*[]WebHook, bool)`

GetWebHookConfigsOk returns a tuple with the WebHookConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebHookConfigs

`func (o *Receivers) SetWebHookConfigs(v []WebHook)`

SetWebHookConfigs sets WebHookConfigs field to given value.

### HasWebHookConfigs

`func (o *Receivers) HasWebHookConfigs() bool`

HasWebHookConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


