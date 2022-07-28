# OAuth2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** |  | 
**ClientSecret** | **string** |  | 
**TokenUrl** | **string** |  | 
**Scopes** | Pointer to **[]string** |  | [optional] 
**TlsConfig** | Pointer to [**TLSConfig**](TLSConfig.md) |  | [optional] 

## Methods

### NewOAuth2

`func NewOAuth2(clientId string, clientSecret string, tokenUrl string, ) *OAuth2`

NewOAuth2 instantiates a new OAuth2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2WithDefaults

`func NewOAuth2WithDefaults() *OAuth2`

NewOAuth2WithDefaults instantiates a new OAuth2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OAuth2) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuth2) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuth2) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *OAuth2) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuth2) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuth2) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetTokenUrl

`func (o *OAuth2) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *OAuth2) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *OAuth2) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.


### GetScopes

`func (o *OAuth2) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OAuth2) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OAuth2) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *OAuth2) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetTlsConfig

`func (o *OAuth2) GetTlsConfig() TLSConfig`

GetTlsConfig returns the TlsConfig field if non-nil, zero value otherwise.

### GetTlsConfigOk

`func (o *OAuth2) GetTlsConfigOk() (*TLSConfig, bool)`

GetTlsConfigOk returns a tuple with the TlsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsConfig

`func (o *OAuth2) SetTlsConfig(v TLSConfig)`

SetTlsConfig sets TlsConfig field to given value.

### HasTlsConfig

`func (o *OAuth2) HasTlsConfig() bool`

HasTlsConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


