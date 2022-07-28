# V1InstancesAlertconfigsUpdateRequestReceiversInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | &#x60;Additional Validators:&#x60; * must be unique * should only include the characters: a-zA-Z0-9- | 
**EmailConfigs** | Pointer to [**[]V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner**](V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner.md) | Email configurations | [optional] 
**OpsgenieConfigs** | Pointer to [**[]V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner**](V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner.md) | Configuration for ops genie. | [optional] 
**WebHookConfigs** | Pointer to [**[]V1InstancesAlertconfigsUpdateRequestReceiversInnerWebHookConfigsInner**](V1InstancesAlertconfigsUpdateRequestReceiversInnerWebHookConfigsInner.md) |  | [optional] 

## Methods

### NewV1InstancesAlertconfigsUpdateRequestReceiversInner

`func NewV1InstancesAlertconfigsUpdateRequestReceiversInner(name string, ) *V1InstancesAlertconfigsUpdateRequestReceiversInner`

NewV1InstancesAlertconfigsUpdateRequestReceiversInner instantiates a new V1InstancesAlertconfigsUpdateRequestReceiversInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesAlertconfigsUpdateRequestReceiversInnerWithDefaults

`func NewV1InstancesAlertconfigsUpdateRequestReceiversInnerWithDefaults() *V1InstancesAlertconfigsUpdateRequestReceiversInner`

NewV1InstancesAlertconfigsUpdateRequestReceiversInnerWithDefaults instantiates a new V1InstancesAlertconfigsUpdateRequestReceiversInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V1InstancesAlertconfigsUpdateRequestReceiversInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1InstancesAlertconfigsUpdateRequestReceiversInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1InstancesAlertconfigsUpdateRequestReceiversInner) SetName(v string)`

SetName sets Name field to given value.


### GetEmailConfigs

`func (o *V1InstancesAlertconfigsUpdateRequestReceiversInner) GetEmailConfigs() []V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner`

GetEmailConfigs returns the EmailConfigs field if non-nil, zero value otherwise.

### GetEmailConfigsOk

`func (o *V1InstancesAlertconfigsUpdateRequestReceiversInner) GetEmailConfigsOk() (*[]V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner, bool)`

GetEmailConfigsOk returns a tuple with the EmailConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConfigs

`func (o *V1InstancesAlertconfigsUpdateRequestReceiversInner) SetEmailConfigs(v []V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner)`

SetEmailConfigs sets EmailConfigs field to given value.

### HasEmailConfigs

`func (o *V1InstancesAlertconfigsUpdateRequestReceiversInner) HasEmailConfigs() bool`

HasEmailConfigs returns a boolean if a field has been set.

### GetOpsgenieConfigs

`func (o *V1InstancesAlertconfigsUpdateRequestReceiversInner) GetOpsgenieConfigs() []V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner`

GetOpsgenieConfigs returns the OpsgenieConfigs field if non-nil, zero value otherwise.

### GetOpsgenieConfigsOk

`func (o *V1InstancesAlertconfigsUpdateRequestReceiversInner) GetOpsgenieConfigsOk() (*[]V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner, bool)`

GetOpsgenieConfigsOk returns a tuple with the OpsgenieConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsgenieConfigs

`func (o *V1InstancesAlertconfigsUpdateRequestReceiversInner) SetOpsgenieConfigs(v []V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner)`

SetOpsgenieConfigs sets OpsgenieConfigs field to given value.

### HasOpsgenieConfigs

`func (o *V1InstancesAlertconfigsUpdateRequestReceiversInner) HasOpsgenieConfigs() bool`

HasOpsgenieConfigs returns a boolean if a field has been set.

### GetWebHookConfigs

`func (o *V1InstancesAlertconfigsUpdateRequestReceiversInner) GetWebHookConfigs() []V1InstancesAlertconfigsUpdateRequestReceiversInnerWebHookConfigsInner`

GetWebHookConfigs returns the WebHookConfigs field if non-nil, zero value otherwise.

### GetWebHookConfigsOk

`func (o *V1InstancesAlertconfigsUpdateRequestReceiversInner) GetWebHookConfigsOk() (*[]V1InstancesAlertconfigsUpdateRequestReceiversInnerWebHookConfigsInner, bool)`

GetWebHookConfigsOk returns a tuple with the WebHookConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebHookConfigs

`func (o *V1InstancesAlertconfigsUpdateRequestReceiversInner) SetWebHookConfigs(v []V1InstancesAlertconfigsUpdateRequestReceiversInnerWebHookConfigsInner)`

SetWebHookConfigs sets WebHookConfigs field to given value.

### HasWebHookConfigs

`func (o *V1InstancesAlertconfigsUpdateRequestReceiversInner) HasWebHookConfigs() bool`

HasWebHookConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


