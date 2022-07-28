# GrafanaOauth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**ApiUrl** | **string** |  | 
**AuthUrl** | **string** |  | 
**Scopes** | Pointer to **string** |  | [optional] [default to "openid profile email"]
**TokenUrl** | **string** |  | 
**OauthClientId** | **string** |  | 
**OauthClientSecret** | **string** |  | 
**RoleAttributeStrict** | Pointer to **bool** |  | [optional] [default to true]
**RoleAttributePath** | **string** |  | 

## Methods

### NewGrafanaOauth

`func NewGrafanaOauth(enabled bool, apiUrl string, authUrl string, tokenUrl string, oauthClientId string, oauthClientSecret string, roleAttributePath string, ) *GrafanaOauth`

NewGrafanaOauth instantiates a new GrafanaOauth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrafanaOauthWithDefaults

`func NewGrafanaOauthWithDefaults() *GrafanaOauth`

NewGrafanaOauthWithDefaults instantiates a new GrafanaOauth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GrafanaOauth) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GrafanaOauth) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GrafanaOauth) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetApiUrl

`func (o *GrafanaOauth) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *GrafanaOauth) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *GrafanaOauth) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.


### GetAuthUrl

`func (o *GrafanaOauth) GetAuthUrl() string`

GetAuthUrl returns the AuthUrl field if non-nil, zero value otherwise.

### GetAuthUrlOk

`func (o *GrafanaOauth) GetAuthUrlOk() (*string, bool)`

GetAuthUrlOk returns a tuple with the AuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUrl

`func (o *GrafanaOauth) SetAuthUrl(v string)`

SetAuthUrl sets AuthUrl field to given value.


### GetScopes

`func (o *GrafanaOauth) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *GrafanaOauth) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *GrafanaOauth) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *GrafanaOauth) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetTokenUrl

`func (o *GrafanaOauth) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *GrafanaOauth) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *GrafanaOauth) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.


### GetOauthClientId

`func (o *GrafanaOauth) GetOauthClientId() string`

GetOauthClientId returns the OauthClientId field if non-nil, zero value otherwise.

### GetOauthClientIdOk

`func (o *GrafanaOauth) GetOauthClientIdOk() (*string, bool)`

GetOauthClientIdOk returns a tuple with the OauthClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientId

`func (o *GrafanaOauth) SetOauthClientId(v string)`

SetOauthClientId sets OauthClientId field to given value.


### GetOauthClientSecret

`func (o *GrafanaOauth) GetOauthClientSecret() string`

GetOauthClientSecret returns the OauthClientSecret field if non-nil, zero value otherwise.

### GetOauthClientSecretOk

`func (o *GrafanaOauth) GetOauthClientSecretOk() (*string, bool)`

GetOauthClientSecretOk returns a tuple with the OauthClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientSecret

`func (o *GrafanaOauth) SetOauthClientSecret(v string)`

SetOauthClientSecret sets OauthClientSecret field to given value.


### GetRoleAttributeStrict

`func (o *GrafanaOauth) GetRoleAttributeStrict() bool`

GetRoleAttributeStrict returns the RoleAttributeStrict field if non-nil, zero value otherwise.

### GetRoleAttributeStrictOk

`func (o *GrafanaOauth) GetRoleAttributeStrictOk() (*bool, bool)`

GetRoleAttributeStrictOk returns a tuple with the RoleAttributeStrict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAttributeStrict

`func (o *GrafanaOauth) SetRoleAttributeStrict(v bool)`

SetRoleAttributeStrict sets RoleAttributeStrict field to given value.

### HasRoleAttributeStrict

`func (o *GrafanaOauth) HasRoleAttributeStrict() bool`

HasRoleAttributeStrict returns a boolean if a field has been set.

### GetRoleAttributePath

`func (o *GrafanaOauth) GetRoleAttributePath() string`

GetRoleAttributePath returns the RoleAttributePath field if non-nil, zero value otherwise.

### GetRoleAttributePathOk

`func (o *GrafanaOauth) GetRoleAttributePathOk() (*string, bool)`

GetRoleAttributePathOk returns a tuple with the RoleAttributePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAttributePath

`func (o *GrafanaOauth) SetRoleAttributePath(v string)`

SetRoleAttributePath sets RoleAttributePath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


