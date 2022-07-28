# V1InstancesScrapeconfigsCreateRequestOauth2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | clientId | 
**ClientSecret** | **string** | clientSecret | 
**TokenUrl** | **string** | The URL to fetch the token from. | 
**Scopes** | Pointer to **[]string** | The URL to fetch the token from. | [optional] 
**TlsConfig** | Pointer to [**V1InstancesScrapeconfigsCreateRequestOauth2TlsConfig**](V1InstancesScrapeconfigsCreateRequestOauth2TlsConfig.md) |  | [optional] 

## Methods

### NewV1InstancesScrapeconfigsCreateRequestOauth2

`func NewV1InstancesScrapeconfigsCreateRequestOauth2(clientId string, clientSecret string, tokenUrl string, ) *V1InstancesScrapeconfigsCreateRequestOauth2`

NewV1InstancesScrapeconfigsCreateRequestOauth2 instantiates a new V1InstancesScrapeconfigsCreateRequestOauth2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesScrapeconfigsCreateRequestOauth2WithDefaults

`func NewV1InstancesScrapeconfigsCreateRequestOauth2WithDefaults() *V1InstancesScrapeconfigsCreateRequestOauth2`

NewV1InstancesScrapeconfigsCreateRequestOauth2WithDefaults instantiates a new V1InstancesScrapeconfigsCreateRequestOauth2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *V1InstancesScrapeconfigsCreateRequestOauth2) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *V1InstancesScrapeconfigsCreateRequestOauth2) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *V1InstancesScrapeconfigsCreateRequestOauth2) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *V1InstancesScrapeconfigsCreateRequestOauth2) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *V1InstancesScrapeconfigsCreateRequestOauth2) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *V1InstancesScrapeconfigsCreateRequestOauth2) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetTokenUrl

`func (o *V1InstancesScrapeconfigsCreateRequestOauth2) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *V1InstancesScrapeconfigsCreateRequestOauth2) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *V1InstancesScrapeconfigsCreateRequestOauth2) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.


### GetScopes

`func (o *V1InstancesScrapeconfigsCreateRequestOauth2) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *V1InstancesScrapeconfigsCreateRequestOauth2) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *V1InstancesScrapeconfigsCreateRequestOauth2) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *V1InstancesScrapeconfigsCreateRequestOauth2) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetTlsConfig

`func (o *V1InstancesScrapeconfigsCreateRequestOauth2) GetTlsConfig() V1InstancesScrapeconfigsCreateRequestOauth2TlsConfig`

GetTlsConfig returns the TlsConfig field if non-nil, zero value otherwise.

### GetTlsConfigOk

`func (o *V1InstancesScrapeconfigsCreateRequestOauth2) GetTlsConfigOk() (*V1InstancesScrapeconfigsCreateRequestOauth2TlsConfig, bool)`

GetTlsConfigOk returns a tuple with the TlsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsConfig

`func (o *V1InstancesScrapeconfigsCreateRequestOauth2) SetTlsConfig(v V1InstancesScrapeconfigsCreateRequestOauth2TlsConfig)`

SetTlsConfig sets TlsConfig field to given value.

### HasTlsConfig

`func (o *V1InstancesScrapeconfigsCreateRequestOauth2) HasTlsConfig() bool`

HasTlsConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


