# V1InstancesAlertconfigsUpdateRequestReceiversInnerWebHookConfigsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The endpoint to send HTTP POST requests to. &#x60;Additional Validators:&#x60; * must be a syntactically valid url address | [optional] 
**MsTeams** | Pointer to **bool** | Microsoft Teams webhooks require special handling. If you set this property to true, it is treated as such | [optional] [default to false]

## Methods

### NewV1InstancesAlertconfigsUpdateRequestReceiversInnerWebHookConfigsInner

`func NewV1InstancesAlertconfigsUpdateRequestReceiversInnerWebHookConfigsInner() *V1InstancesAlertconfigsUpdateRequestReceiversInnerWebHookConfigsInner`

NewV1InstancesAlertconfigsUpdateRequestReceiversInnerWebHookConfigsInner instantiates a new V1InstancesAlertconfigsUpdateRequestReceiversInnerWebHookConfigsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesAlertconfigsUpdateRequestReceiversInnerWebHookConfigsInnerWithDefaults

`func NewV1InstancesAlertconfigsUpdateRequestReceiversInnerWebHookConfigsInnerWithDefaults() *V1InstancesAlertconfigsUpdateRequestReceiversInnerWebHookConfigsInner`

NewV1InstancesAlertconfigsUpdateRequestReceiversInnerWebHookConfigsInnerWithDefaults instantiates a new V1InstancesAlertconfigsUpdateRequestReceiversInnerWebHookConfigsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerWebHookConfigsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerWebHookConfigsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerWebHookConfigsInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerWebHookConfigsInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMsTeams

`func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerWebHookConfigsInner) GetMsTeams() bool`

GetMsTeams returns the MsTeams field if non-nil, zero value otherwise.

### GetMsTeamsOk

`func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerWebHookConfigsInner) GetMsTeamsOk() (*bool, bool)`

GetMsTeamsOk returns a tuple with the MsTeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsTeams

`func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerWebHookConfigsInner) SetMsTeams(v bool)`

SetMsTeams sets MsTeams field to given value.

### HasMsTeams

`func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerWebHookConfigsInner) HasMsTeams() bool`

HasMsTeams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


