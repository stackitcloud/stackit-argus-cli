# V1InstancesAlertconfigsUpdateRequestGlobal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResolveTimeout** | Pointer to **string** | ResolveTimeout is the default value used by alertmanager if the alert does not include EndsAt, after this time passes it can declare the alert as resolved if it has not been updated. This has no impact on alerts from Prometheus, as they always include EndsAt. | [optional] [default to "5m"]
**SmtpFrom** | Pointer to **string** | The default SMTP From header field. | [optional] 
**SmtpSmarthost** | Pointer to **string** | The default SMTP smarthost used for sending emails, including port number. Port number usually is 25, or 587 for SMTP over TLS (sometimes referred to as STARTTLS). Example: smtp.example.org:587 | [optional] 
**SmtpAuthUsername** | Pointer to **string** | SMTP Auth using CRAM-MD5, LOGIN and PLAIN. If empty, Alertmanager doesn&#39;t authenticate to the SMTP server. | [optional] 
**SmtpAuthPassword** | Pointer to **string** | SMTP Auth using LOGIN and PLAIN. | [optional] 
**OpsgenieApiKey** | Pointer to **string** | Opsgenie api key | [optional] 
**OpsgenieApiUrl** | Pointer to **string** | Opsgenie api url | [optional] 

## Methods

### NewV1InstancesAlertconfigsUpdateRequestGlobal

`func NewV1InstancesAlertconfigsUpdateRequestGlobal() *V1InstancesAlertconfigsUpdateRequestGlobal`

NewV1InstancesAlertconfigsUpdateRequestGlobal instantiates a new V1InstancesAlertconfigsUpdateRequestGlobal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesAlertconfigsUpdateRequestGlobalWithDefaults

`func NewV1InstancesAlertconfigsUpdateRequestGlobalWithDefaults() *V1InstancesAlertconfigsUpdateRequestGlobal`

NewV1InstancesAlertconfigsUpdateRequestGlobalWithDefaults instantiates a new V1InstancesAlertconfigsUpdateRequestGlobal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResolveTimeout

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) GetResolveTimeout() string`

GetResolveTimeout returns the ResolveTimeout field if non-nil, zero value otherwise.

### GetResolveTimeoutOk

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) GetResolveTimeoutOk() (*string, bool)`

GetResolveTimeoutOk returns a tuple with the ResolveTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveTimeout

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) SetResolveTimeout(v string)`

SetResolveTimeout sets ResolveTimeout field to given value.

### HasResolveTimeout

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) HasResolveTimeout() bool`

HasResolveTimeout returns a boolean if a field has been set.

### GetSmtpFrom

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) GetSmtpFrom() string`

GetSmtpFrom returns the SmtpFrom field if non-nil, zero value otherwise.

### GetSmtpFromOk

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) GetSmtpFromOk() (*string, bool)`

GetSmtpFromOk returns a tuple with the SmtpFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpFrom

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) SetSmtpFrom(v string)`

SetSmtpFrom sets SmtpFrom field to given value.

### HasSmtpFrom

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) HasSmtpFrom() bool`

HasSmtpFrom returns a boolean if a field has been set.

### GetSmtpSmarthost

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) GetSmtpSmarthost() string`

GetSmtpSmarthost returns the SmtpSmarthost field if non-nil, zero value otherwise.

### GetSmtpSmarthostOk

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) GetSmtpSmarthostOk() (*string, bool)`

GetSmtpSmarthostOk returns a tuple with the SmtpSmarthost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpSmarthost

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) SetSmtpSmarthost(v string)`

SetSmtpSmarthost sets SmtpSmarthost field to given value.

### HasSmtpSmarthost

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) HasSmtpSmarthost() bool`

HasSmtpSmarthost returns a boolean if a field has been set.

### GetSmtpAuthUsername

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) GetSmtpAuthUsername() string`

GetSmtpAuthUsername returns the SmtpAuthUsername field if non-nil, zero value otherwise.

### GetSmtpAuthUsernameOk

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) GetSmtpAuthUsernameOk() (*string, bool)`

GetSmtpAuthUsernameOk returns a tuple with the SmtpAuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpAuthUsername

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) SetSmtpAuthUsername(v string)`

SetSmtpAuthUsername sets SmtpAuthUsername field to given value.

### HasSmtpAuthUsername

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) HasSmtpAuthUsername() bool`

HasSmtpAuthUsername returns a boolean if a field has been set.

### GetSmtpAuthPassword

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) GetSmtpAuthPassword() string`

GetSmtpAuthPassword returns the SmtpAuthPassword field if non-nil, zero value otherwise.

### GetSmtpAuthPasswordOk

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) GetSmtpAuthPasswordOk() (*string, bool)`

GetSmtpAuthPasswordOk returns a tuple with the SmtpAuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpAuthPassword

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) SetSmtpAuthPassword(v string)`

SetSmtpAuthPassword sets SmtpAuthPassword field to given value.

### HasSmtpAuthPassword

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) HasSmtpAuthPassword() bool`

HasSmtpAuthPassword returns a boolean if a field has been set.

### GetOpsgenieApiKey

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) GetOpsgenieApiKey() string`

GetOpsgenieApiKey returns the OpsgenieApiKey field if non-nil, zero value otherwise.

### GetOpsgenieApiKeyOk

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) GetOpsgenieApiKeyOk() (*string, bool)`

GetOpsgenieApiKeyOk returns a tuple with the OpsgenieApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsgenieApiKey

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) SetOpsgenieApiKey(v string)`

SetOpsgenieApiKey sets OpsgenieApiKey field to given value.

### HasOpsgenieApiKey

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) HasOpsgenieApiKey() bool`

HasOpsgenieApiKey returns a boolean if a field has been set.

### GetOpsgenieApiUrl

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) GetOpsgenieApiUrl() string`

GetOpsgenieApiUrl returns the OpsgenieApiUrl field if non-nil, zero value otherwise.

### GetOpsgenieApiUrlOk

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) GetOpsgenieApiUrlOk() (*string, bool)`

GetOpsgenieApiUrlOk returns a tuple with the OpsgenieApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsgenieApiUrl

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) SetOpsgenieApiUrl(v string)`

SetOpsgenieApiUrl sets OpsgenieApiUrl field to given value.

### HasOpsgenieApiUrl

`func (o *V1InstancesAlertconfigsUpdateRequestGlobal) HasOpsgenieApiUrl() bool`

HasOpsgenieApiUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


