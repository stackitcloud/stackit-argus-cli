# V1InstancesGrafanaConfigsUpdateRequestGenericOauth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | enable or disable generic oauth login | 
**ApiUrl** | **string** | Set api_url to the resource that returns OpenID UserInfo compatible information. | 
**AuthUrl** | **string** | Authentication endpoint of idp. | 
**Scopes** | Pointer to **string** | Space seperated list of scopes of the token | [optional] [default to "openid profile email"]
**TokenUrl** | **string** | Token endpoint of the idp. | 
**OauthClientId** | **string** | Oauth client id for auth endpoint. | 
**OauthClientSecret** | **string** | Oauth client secret for auth endpoint. | 
**RoleAttributeStrict** | Pointer to **bool** | If  therole_attribute_path property does not return a role, then the user is assigned the Viewer role by default. You can disable the role assignment by setting role_attribute_strict &#x3D; true. It denies user access if no role or an invalid role is returned. | [optional] [default to true]
**RoleAttributePath** | **string** | Grafana checks for the presence of a role using the JMESPath specified via the role_attribute_path configuration option. The JMESPath is applied to the id_token first. If there is no match, then the UserInfo endpoint specified via the api_url configuration option is tried next. The result after evaluation of the role_attribute_path JMESPath expression should be a valid Grafana role, for example, Viewer, Editor or Admin For example: contains(roles[\\*], &#39;grafana-admin&#39;) &amp;&amp; &#39;Admin&#39; || contains(roles[\\*], &#39;grafana-editor&#39;) &amp;&amp; &#39;Editor&#39; || contains(roles[\\*], &#39;grafana-viewer&#39;) &amp;&amp; &#39;Viewer&#39; | 

## Methods

### NewV1InstancesGrafanaConfigsUpdateRequestGenericOauth

`func NewV1InstancesGrafanaConfigsUpdateRequestGenericOauth(enabled bool, apiUrl string, authUrl string, tokenUrl string, oauthClientId string, oauthClientSecret string, roleAttributePath string, ) *V1InstancesGrafanaConfigsUpdateRequestGenericOauth`

NewV1InstancesGrafanaConfigsUpdateRequestGenericOauth instantiates a new V1InstancesGrafanaConfigsUpdateRequestGenericOauth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesGrafanaConfigsUpdateRequestGenericOauthWithDefaults

`func NewV1InstancesGrafanaConfigsUpdateRequestGenericOauthWithDefaults() *V1InstancesGrafanaConfigsUpdateRequestGenericOauth`

NewV1InstancesGrafanaConfigsUpdateRequestGenericOauthWithDefaults instantiates a new V1InstancesGrafanaConfigsUpdateRequestGenericOauth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetApiUrl

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.


### GetAuthUrl

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) GetAuthUrl() string`

GetAuthUrl returns the AuthUrl field if non-nil, zero value otherwise.

### GetAuthUrlOk

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) GetAuthUrlOk() (*string, bool)`

GetAuthUrlOk returns a tuple with the AuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUrl

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) SetAuthUrl(v string)`

SetAuthUrl sets AuthUrl field to given value.


### GetScopes

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetTokenUrl

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.


### GetOauthClientId

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) GetOauthClientId() string`

GetOauthClientId returns the OauthClientId field if non-nil, zero value otherwise.

### GetOauthClientIdOk

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) GetOauthClientIdOk() (*string, bool)`

GetOauthClientIdOk returns a tuple with the OauthClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientId

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) SetOauthClientId(v string)`

SetOauthClientId sets OauthClientId field to given value.


### GetOauthClientSecret

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) GetOauthClientSecret() string`

GetOauthClientSecret returns the OauthClientSecret field if non-nil, zero value otherwise.

### GetOauthClientSecretOk

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) GetOauthClientSecretOk() (*string, bool)`

GetOauthClientSecretOk returns a tuple with the OauthClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClientSecret

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) SetOauthClientSecret(v string)`

SetOauthClientSecret sets OauthClientSecret field to given value.


### GetRoleAttributeStrict

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) GetRoleAttributeStrict() bool`

GetRoleAttributeStrict returns the RoleAttributeStrict field if non-nil, zero value otherwise.

### GetRoleAttributeStrictOk

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) GetRoleAttributeStrictOk() (*bool, bool)`

GetRoleAttributeStrictOk returns a tuple with the RoleAttributeStrict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAttributeStrict

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) SetRoleAttributeStrict(v bool)`

SetRoleAttributeStrict sets RoleAttributeStrict field to given value.

### HasRoleAttributeStrict

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) HasRoleAttributeStrict() bool`

HasRoleAttributeStrict returns a boolean if a field has been set.

### GetRoleAttributePath

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) GetRoleAttributePath() string`

GetRoleAttributePath returns the RoleAttributePath field if non-nil, zero value otherwise.

### GetRoleAttributePathOk

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) GetRoleAttributePathOk() (*string, bool)`

GetRoleAttributePathOk returns a tuple with the RoleAttributePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAttributePath

`func (o *V1InstancesGrafanaConfigsUpdateRequestGenericOauth) SetRoleAttributePath(v string)`

SetRoleAttributePath sets RoleAttributePath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


