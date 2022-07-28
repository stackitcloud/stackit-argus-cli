# HTTPServiceSD

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**RefreshInterval** | Pointer to **string** |  | [optional] [default to "60s"]
**BasicAuth** | Pointer to [**BasicAuth**](BasicAuth.md) |  | [optional] 
**Oauth2** | Pointer to [**OAuth2**](OAuth2.md) |  | [optional] 
**TlsConfig** | Pointer to [**TLSConfig**](TLSConfig.md) |  | [optional] 

## Methods

### NewHTTPServiceSD

`func NewHTTPServiceSD(url string, ) *HTTPServiceSD`

NewHTTPServiceSD instantiates a new HTTPServiceSD object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHTTPServiceSDWithDefaults

`func NewHTTPServiceSDWithDefaults() *HTTPServiceSD`

NewHTTPServiceSDWithDefaults instantiates a new HTTPServiceSD object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *HTTPServiceSD) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HTTPServiceSD) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HTTPServiceSD) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRefreshInterval

`func (o *HTTPServiceSD) GetRefreshInterval() string`

GetRefreshInterval returns the RefreshInterval field if non-nil, zero value otherwise.

### GetRefreshIntervalOk

`func (o *HTTPServiceSD) GetRefreshIntervalOk() (*string, bool)`

GetRefreshIntervalOk returns a tuple with the RefreshInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshInterval

`func (o *HTTPServiceSD) SetRefreshInterval(v string)`

SetRefreshInterval sets RefreshInterval field to given value.

### HasRefreshInterval

`func (o *HTTPServiceSD) HasRefreshInterval() bool`

HasRefreshInterval returns a boolean if a field has been set.

### GetBasicAuth

`func (o *HTTPServiceSD) GetBasicAuth() BasicAuth`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *HTTPServiceSD) GetBasicAuthOk() (*BasicAuth, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *HTTPServiceSD) SetBasicAuth(v BasicAuth)`

SetBasicAuth sets BasicAuth field to given value.

### HasBasicAuth

`func (o *HTTPServiceSD) HasBasicAuth() bool`

HasBasicAuth returns a boolean if a field has been set.

### GetOauth2

`func (o *HTTPServiceSD) GetOauth2() OAuth2`

GetOauth2 returns the Oauth2 field if non-nil, zero value otherwise.

### GetOauth2Ok

`func (o *HTTPServiceSD) GetOauth2Ok() (*OAuth2, bool)`

GetOauth2Ok returns a tuple with the Oauth2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2

`func (o *HTTPServiceSD) SetOauth2(v OAuth2)`

SetOauth2 sets Oauth2 field to given value.

### HasOauth2

`func (o *HTTPServiceSD) HasOauth2() bool`

HasOauth2 returns a boolean if a field has been set.

### GetTlsConfig

`func (o *HTTPServiceSD) GetTlsConfig() TLSConfig`

GetTlsConfig returns the TlsConfig field if non-nil, zero value otherwise.

### GetTlsConfigOk

`func (o *HTTPServiceSD) GetTlsConfigOk() (*TLSConfig, bool)`

GetTlsConfigOk returns a tuple with the TlsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsConfig

`func (o *HTTPServiceSD) SetTlsConfig(v TLSConfig)`

SetTlsConfig sets TlsConfig field to given value.

### HasTlsConfig

`func (o *HTTPServiceSD) HasTlsConfig() bool`

HasTlsConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


