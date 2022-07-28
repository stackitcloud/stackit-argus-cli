# V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | URL from which the targets are fetched. | 
**RefreshInterval** | Pointer to **string** | Refresh interval to re-query the endpoint. E.g. 60s &#x60;Additional Validators:&#x60; * must be a valid time format* must be &gt;&#x3D; 60s | [optional] [default to "60s"]
**BasicAuth** | Pointer to [**V1InstancesScrapeconfigsCreateRequestBasicAuth**](V1InstancesScrapeconfigsCreateRequestBasicAuth.md) |  | [optional] 
**TlsConfig** | Pointer to [**V1InstancesScrapeconfigsCreateRequestOauth2TlsConfig**](V1InstancesScrapeconfigsCreateRequestOauth2TlsConfig.md) |  | [optional] 
**Oauth2** | Pointer to [**V1InstancesScrapeconfigsCreateRequestOauth2**](V1InstancesScrapeconfigsCreateRequestOauth2.md) |  | [optional] 

## Methods

### NewV1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner

`func NewV1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner(url string, ) *V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner`

NewV1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner instantiates a new V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesScrapeconfigsCreateRequestHttpSDConfigsInnerWithDefaults

`func NewV1InstancesScrapeconfigsCreateRequestHttpSDConfigsInnerWithDefaults() *V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner`

NewV1InstancesScrapeconfigsCreateRequestHttpSDConfigsInnerWithDefaults instantiates a new V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRefreshInterval

`func (o *V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner) GetRefreshInterval() string`

GetRefreshInterval returns the RefreshInterval field if non-nil, zero value otherwise.

### GetRefreshIntervalOk

`func (o *V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner) GetRefreshIntervalOk() (*string, bool)`

GetRefreshIntervalOk returns a tuple with the RefreshInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshInterval

`func (o *V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner) SetRefreshInterval(v string)`

SetRefreshInterval sets RefreshInterval field to given value.

### HasRefreshInterval

`func (o *V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner) HasRefreshInterval() bool`

HasRefreshInterval returns a boolean if a field has been set.

### GetBasicAuth

`func (o *V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner) GetBasicAuth() V1InstancesScrapeconfigsCreateRequestBasicAuth`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner) GetBasicAuthOk() (*V1InstancesScrapeconfigsCreateRequestBasicAuth, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner) SetBasicAuth(v V1InstancesScrapeconfigsCreateRequestBasicAuth)`

SetBasicAuth sets BasicAuth field to given value.

### HasBasicAuth

`func (o *V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner) HasBasicAuth() bool`

HasBasicAuth returns a boolean if a field has been set.

### GetTlsConfig

`func (o *V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner) GetTlsConfig() V1InstancesScrapeconfigsCreateRequestOauth2TlsConfig`

GetTlsConfig returns the TlsConfig field if non-nil, zero value otherwise.

### GetTlsConfigOk

`func (o *V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner) GetTlsConfigOk() (*V1InstancesScrapeconfigsCreateRequestOauth2TlsConfig, bool)`

GetTlsConfigOk returns a tuple with the TlsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsConfig

`func (o *V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner) SetTlsConfig(v V1InstancesScrapeconfigsCreateRequestOauth2TlsConfig)`

SetTlsConfig sets TlsConfig field to given value.

### HasTlsConfig

`func (o *V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner) HasTlsConfig() bool`

HasTlsConfig returns a boolean if a field has been set.

### GetOauth2

`func (o *V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner) GetOauth2() V1InstancesScrapeconfigsCreateRequestOauth2`

GetOauth2 returns the Oauth2 field if non-nil, zero value otherwise.

### GetOauth2Ok

`func (o *V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner) GetOauth2Ok() (*V1InstancesScrapeconfigsCreateRequestOauth2, bool)`

GetOauth2Ok returns a tuple with the Oauth2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2

`func (o *V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner) SetOauth2(v V1InstancesScrapeconfigsCreateRequestOauth2)`

SetOauth2 sets Oauth2 field to given value.

### HasOauth2

`func (o *V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner) HasOauth2() bool`

HasOauth2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


