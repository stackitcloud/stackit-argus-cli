# Global

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResolveTimeout** | Pointer to **string** |  | [optional] [default to "5m"]
**SmtpFrom** | Pointer to **string** |  | [optional] 
**SmtpSmarthost** | Pointer to **string** |  | [optional] 
**SmtpAuthUsername** | Pointer to **string** |  | [optional] 
**SmtpAuthPassword** | Pointer to **string** |  | [optional] 
**SmtpAuthIdentity** | Pointer to **string** |  | [optional] 
**OpsgenieApiKey** | Pointer to **string** |  | [optional] 
**OpsgenieApiUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewGlobal

`func NewGlobal() *Global`

NewGlobal instantiates a new Global object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalWithDefaults

`func NewGlobalWithDefaults() *Global`

NewGlobalWithDefaults instantiates a new Global object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResolveTimeout

`func (o *Global) GetResolveTimeout() string`

GetResolveTimeout returns the ResolveTimeout field if non-nil, zero value otherwise.

### GetResolveTimeoutOk

`func (o *Global) GetResolveTimeoutOk() (*string, bool)`

GetResolveTimeoutOk returns a tuple with the ResolveTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveTimeout

`func (o *Global) SetResolveTimeout(v string)`

SetResolveTimeout sets ResolveTimeout field to given value.

### HasResolveTimeout

`func (o *Global) HasResolveTimeout() bool`

HasResolveTimeout returns a boolean if a field has been set.

### GetSmtpFrom

`func (o *Global) GetSmtpFrom() string`

GetSmtpFrom returns the SmtpFrom field if non-nil, zero value otherwise.

### GetSmtpFromOk

`func (o *Global) GetSmtpFromOk() (*string, bool)`

GetSmtpFromOk returns a tuple with the SmtpFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpFrom

`func (o *Global) SetSmtpFrom(v string)`

SetSmtpFrom sets SmtpFrom field to given value.

### HasSmtpFrom

`func (o *Global) HasSmtpFrom() bool`

HasSmtpFrom returns a boolean if a field has been set.

### GetSmtpSmarthost

`func (o *Global) GetSmtpSmarthost() string`

GetSmtpSmarthost returns the SmtpSmarthost field if non-nil, zero value otherwise.

### GetSmtpSmarthostOk

`func (o *Global) GetSmtpSmarthostOk() (*string, bool)`

GetSmtpSmarthostOk returns a tuple with the SmtpSmarthost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpSmarthost

`func (o *Global) SetSmtpSmarthost(v string)`

SetSmtpSmarthost sets SmtpSmarthost field to given value.

### HasSmtpSmarthost

`func (o *Global) HasSmtpSmarthost() bool`

HasSmtpSmarthost returns a boolean if a field has been set.

### GetSmtpAuthUsername

`func (o *Global) GetSmtpAuthUsername() string`

GetSmtpAuthUsername returns the SmtpAuthUsername field if non-nil, zero value otherwise.

### GetSmtpAuthUsernameOk

`func (o *Global) GetSmtpAuthUsernameOk() (*string, bool)`

GetSmtpAuthUsernameOk returns a tuple with the SmtpAuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpAuthUsername

`func (o *Global) SetSmtpAuthUsername(v string)`

SetSmtpAuthUsername sets SmtpAuthUsername field to given value.

### HasSmtpAuthUsername

`func (o *Global) HasSmtpAuthUsername() bool`

HasSmtpAuthUsername returns a boolean if a field has been set.

### GetSmtpAuthPassword

`func (o *Global) GetSmtpAuthPassword() string`

GetSmtpAuthPassword returns the SmtpAuthPassword field if non-nil, zero value otherwise.

### GetSmtpAuthPasswordOk

`func (o *Global) GetSmtpAuthPasswordOk() (*string, bool)`

GetSmtpAuthPasswordOk returns a tuple with the SmtpAuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpAuthPassword

`func (o *Global) SetSmtpAuthPassword(v string)`

SetSmtpAuthPassword sets SmtpAuthPassword field to given value.

### HasSmtpAuthPassword

`func (o *Global) HasSmtpAuthPassword() bool`

HasSmtpAuthPassword returns a boolean if a field has been set.

### GetSmtpAuthIdentity

`func (o *Global) GetSmtpAuthIdentity() string`

GetSmtpAuthIdentity returns the SmtpAuthIdentity field if non-nil, zero value otherwise.

### GetSmtpAuthIdentityOk

`func (o *Global) GetSmtpAuthIdentityOk() (*string, bool)`

GetSmtpAuthIdentityOk returns a tuple with the SmtpAuthIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpAuthIdentity

`func (o *Global) SetSmtpAuthIdentity(v string)`

SetSmtpAuthIdentity sets SmtpAuthIdentity field to given value.

### HasSmtpAuthIdentity

`func (o *Global) HasSmtpAuthIdentity() bool`

HasSmtpAuthIdentity returns a boolean if a field has been set.

### GetOpsgenieApiKey

`func (o *Global) GetOpsgenieApiKey() string`

GetOpsgenieApiKey returns the OpsgenieApiKey field if non-nil, zero value otherwise.

### GetOpsgenieApiKeyOk

`func (o *Global) GetOpsgenieApiKeyOk() (*string, bool)`

GetOpsgenieApiKeyOk returns a tuple with the OpsgenieApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsgenieApiKey

`func (o *Global) SetOpsgenieApiKey(v string)`

SetOpsgenieApiKey sets OpsgenieApiKey field to given value.

### HasOpsgenieApiKey

`func (o *Global) HasOpsgenieApiKey() bool`

HasOpsgenieApiKey returns a boolean if a field has been set.

### GetOpsgenieApiUrl

`func (o *Global) GetOpsgenieApiUrl() string`

GetOpsgenieApiUrl returns the OpsgenieApiUrl field if non-nil, zero value otherwise.

### GetOpsgenieApiUrlOk

`func (o *Global) GetOpsgenieApiUrlOk() (*string, bool)`

GetOpsgenieApiUrlOk returns a tuple with the OpsgenieApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsgenieApiUrl

`func (o *Global) SetOpsgenieApiUrl(v string)`

SetOpsgenieApiUrl sets OpsgenieApiUrl field to given value.

### HasOpsgenieApiUrl

`func (o *Global) HasOpsgenieApiUrl() bool`

HasOpsgenieApiUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


