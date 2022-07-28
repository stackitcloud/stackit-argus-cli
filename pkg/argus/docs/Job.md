# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobName** | **string** |  | 
**Scheme** | Pointer to **string** |  | [optional] [default to "http"]
**ScrapeInterval** | **string** |  | 
**ScrapeTimeout** | **string** |  | 
**MetricsPath** | Pointer to **string** |  | [optional] [default to "/metrics"]
**StaticConfigs** | [**[]StaticConfigs**](StaticConfigs.md) |  | 
**SampleLimit** | Pointer to **int32** |  | [optional] 
**BasicAuth** | Pointer to [**BasicAuth**](BasicAuth.md) |  | [optional] 
**Oauth2** | Pointer to [**OAuth2**](OAuth2.md) |  | [optional] 
**TlsConfig** | Pointer to [**TLSConfig**](TLSConfig.md) |  | [optional] 
**BearerToken** | Pointer to **string** |  | [optional] 
**MetricsRelabelConfigs** | Pointer to [**[]MetricsRelabelConfig**](MetricsRelabelConfig.md) |  | [optional] 
**Params** | Pointer to **map[string][]string** |  | [optional] 
**HttpSDConfigs** | Pointer to [**[]HTTPServiceSD**](HTTPServiceSD.md) |  | [optional] 
**HonorLabels** | Pointer to **bool** |  | [optional] [default to false]
**HonorTimeStamps** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewJob

`func NewJob(jobName string, scrapeInterval string, scrapeTimeout string, staticConfigs []StaticConfigs, ) *Job`

NewJob instantiates a new Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithDefaults

`func NewJobWithDefaults() *Job`

NewJobWithDefaults instantiates a new Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobName

`func (o *Job) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *Job) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *Job) SetJobName(v string)`

SetJobName sets JobName field to given value.


### GetScheme

`func (o *Job) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *Job) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *Job) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *Job) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetScrapeInterval

`func (o *Job) GetScrapeInterval() string`

GetScrapeInterval returns the ScrapeInterval field if non-nil, zero value otherwise.

### GetScrapeIntervalOk

`func (o *Job) GetScrapeIntervalOk() (*string, bool)`

GetScrapeIntervalOk returns a tuple with the ScrapeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrapeInterval

`func (o *Job) SetScrapeInterval(v string)`

SetScrapeInterval sets ScrapeInterval field to given value.


### GetScrapeTimeout

`func (o *Job) GetScrapeTimeout() string`

GetScrapeTimeout returns the ScrapeTimeout field if non-nil, zero value otherwise.

### GetScrapeTimeoutOk

`func (o *Job) GetScrapeTimeoutOk() (*string, bool)`

GetScrapeTimeoutOk returns a tuple with the ScrapeTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrapeTimeout

`func (o *Job) SetScrapeTimeout(v string)`

SetScrapeTimeout sets ScrapeTimeout field to given value.


### GetMetricsPath

`func (o *Job) GetMetricsPath() string`

GetMetricsPath returns the MetricsPath field if non-nil, zero value otherwise.

### GetMetricsPathOk

`func (o *Job) GetMetricsPathOk() (*string, bool)`

GetMetricsPathOk returns a tuple with the MetricsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsPath

`func (o *Job) SetMetricsPath(v string)`

SetMetricsPath sets MetricsPath field to given value.

### HasMetricsPath

`func (o *Job) HasMetricsPath() bool`

HasMetricsPath returns a boolean if a field has been set.

### GetStaticConfigs

`func (o *Job) GetStaticConfigs() []StaticConfigs`

GetStaticConfigs returns the StaticConfigs field if non-nil, zero value otherwise.

### GetStaticConfigsOk

`func (o *Job) GetStaticConfigsOk() (*[]StaticConfigs, bool)`

GetStaticConfigsOk returns a tuple with the StaticConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticConfigs

`func (o *Job) SetStaticConfigs(v []StaticConfigs)`

SetStaticConfigs sets StaticConfigs field to given value.


### GetSampleLimit

`func (o *Job) GetSampleLimit() int32`

GetSampleLimit returns the SampleLimit field if non-nil, zero value otherwise.

### GetSampleLimitOk

`func (o *Job) GetSampleLimitOk() (*int32, bool)`

GetSampleLimitOk returns a tuple with the SampleLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleLimit

`func (o *Job) SetSampleLimit(v int32)`

SetSampleLimit sets SampleLimit field to given value.

### HasSampleLimit

`func (o *Job) HasSampleLimit() bool`

HasSampleLimit returns a boolean if a field has been set.

### GetBasicAuth

`func (o *Job) GetBasicAuth() BasicAuth`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *Job) GetBasicAuthOk() (*BasicAuth, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *Job) SetBasicAuth(v BasicAuth)`

SetBasicAuth sets BasicAuth field to given value.

### HasBasicAuth

`func (o *Job) HasBasicAuth() bool`

HasBasicAuth returns a boolean if a field has been set.

### GetOauth2

`func (o *Job) GetOauth2() OAuth2`

GetOauth2 returns the Oauth2 field if non-nil, zero value otherwise.

### GetOauth2Ok

`func (o *Job) GetOauth2Ok() (*OAuth2, bool)`

GetOauth2Ok returns a tuple with the Oauth2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2

`func (o *Job) SetOauth2(v OAuth2)`

SetOauth2 sets Oauth2 field to given value.

### HasOauth2

`func (o *Job) HasOauth2() bool`

HasOauth2 returns a boolean if a field has been set.

### GetTlsConfig

`func (o *Job) GetTlsConfig() TLSConfig`

GetTlsConfig returns the TlsConfig field if non-nil, zero value otherwise.

### GetTlsConfigOk

`func (o *Job) GetTlsConfigOk() (*TLSConfig, bool)`

GetTlsConfigOk returns a tuple with the TlsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsConfig

`func (o *Job) SetTlsConfig(v TLSConfig)`

SetTlsConfig sets TlsConfig field to given value.

### HasTlsConfig

`func (o *Job) HasTlsConfig() bool`

HasTlsConfig returns a boolean if a field has been set.

### GetBearerToken

`func (o *Job) GetBearerToken() string`

GetBearerToken returns the BearerToken field if non-nil, zero value otherwise.

### GetBearerTokenOk

`func (o *Job) GetBearerTokenOk() (*string, bool)`

GetBearerTokenOk returns a tuple with the BearerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerToken

`func (o *Job) SetBearerToken(v string)`

SetBearerToken sets BearerToken field to given value.

### HasBearerToken

`func (o *Job) HasBearerToken() bool`

HasBearerToken returns a boolean if a field has been set.

### GetMetricsRelabelConfigs

`func (o *Job) GetMetricsRelabelConfigs() []MetricsRelabelConfig`

GetMetricsRelabelConfigs returns the MetricsRelabelConfigs field if non-nil, zero value otherwise.

### GetMetricsRelabelConfigsOk

`func (o *Job) GetMetricsRelabelConfigsOk() (*[]MetricsRelabelConfig, bool)`

GetMetricsRelabelConfigsOk returns a tuple with the MetricsRelabelConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsRelabelConfigs

`func (o *Job) SetMetricsRelabelConfigs(v []MetricsRelabelConfig)`

SetMetricsRelabelConfigs sets MetricsRelabelConfigs field to given value.

### HasMetricsRelabelConfigs

`func (o *Job) HasMetricsRelabelConfigs() bool`

HasMetricsRelabelConfigs returns a boolean if a field has been set.

### GetParams

`func (o *Job) GetParams() map[string][]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *Job) GetParamsOk() (*map[string][]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *Job) SetParams(v map[string][]string)`

SetParams sets Params field to given value.

### HasParams

`func (o *Job) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetHttpSDConfigs

`func (o *Job) GetHttpSDConfigs() []HTTPServiceSD`

GetHttpSDConfigs returns the HttpSDConfigs field if non-nil, zero value otherwise.

### GetHttpSDConfigsOk

`func (o *Job) GetHttpSDConfigsOk() (*[]HTTPServiceSD, bool)`

GetHttpSDConfigsOk returns a tuple with the HttpSDConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpSDConfigs

`func (o *Job) SetHttpSDConfigs(v []HTTPServiceSD)`

SetHttpSDConfigs sets HttpSDConfigs field to given value.

### HasHttpSDConfigs

`func (o *Job) HasHttpSDConfigs() bool`

HasHttpSDConfigs returns a boolean if a field has been set.

### GetHonorLabels

`func (o *Job) GetHonorLabels() bool`

GetHonorLabels returns the HonorLabels field if non-nil, zero value otherwise.

### GetHonorLabelsOk

`func (o *Job) GetHonorLabelsOk() (*bool, bool)`

GetHonorLabelsOk returns a tuple with the HonorLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHonorLabels

`func (o *Job) SetHonorLabels(v bool)`

SetHonorLabels sets HonorLabels field to given value.

### HasHonorLabels

`func (o *Job) HasHonorLabels() bool`

HasHonorLabels returns a boolean if a field has been set.

### GetHonorTimeStamps

`func (o *Job) GetHonorTimeStamps() bool`

GetHonorTimeStamps returns the HonorTimeStamps field if non-nil, zero value otherwise.

### GetHonorTimeStampsOk

`func (o *Job) GetHonorTimeStampsOk() (*bool, bool)`

GetHonorTimeStampsOk returns a tuple with the HonorTimeStamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHonorTimeStamps

`func (o *Job) SetHonorTimeStamps(v bool)`

SetHonorTimeStamps sets HonorTimeStamps field to given value.

### HasHonorTimeStamps

`func (o *Job) HasHonorTimeStamps() bool`

HasHonorTimeStamps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


