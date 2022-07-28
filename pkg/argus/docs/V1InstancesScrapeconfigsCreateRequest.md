# V1InstancesScrapeconfigsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StaticConfigs** | [**[]V1InstancesScrapeconfigsCreateRequestStaticConfigsInner**](V1InstancesScrapeconfigsCreateRequestStaticConfigsInner.md) | A list of scrape configurations. | 
**JobName** | **string** | The job name assigned to scraped metrics by default. &#x60;Additional Validators:&#x60; * must be unique * key and values should only include the characters: a-zA-Z0-9- | 
**Scheme** | **string** | Configures the protocol scheme used for requests. https or http | 
**ScrapeInterval** | **string** | How frequently to scrape targets from this job. E.g. 5m &#x60;Additional Validators:&#x60; * must be a valid time format* must be &gt;&#x3D; 60s | 
**ScrapeTimeout** | **string** | Per-scrape timeout when scraping this job. &#x60;Additional Validators:&#x60; * must be a valid time format* must be smaller than scrapeInterval | 
**MetricsPath** | Pointer to **string** | The HTTP resource path on which to fetch metrics from targets. E.g. /metrics | [optional] [default to "/metrics"]
**SampleLimit** | Pointer to **float32** | Per-scrape limit on number of scraped samples that will be accepted. If more than this number of samples are present after metric relabeling the entire scrape will be treated as failed. The total limit depends on the service plan target limits * samples | [optional] 
**BasicAuth** | Pointer to [**V1InstancesScrapeconfigsCreateRequestBasicAuth**](V1InstancesScrapeconfigsCreateRequestBasicAuth.md) |  | [optional] 
**Oauth2** | Pointer to [**V1InstancesScrapeconfigsCreateRequestOauth2**](V1InstancesScrapeconfigsCreateRequestOauth2.md) |  | [optional] 
**TlsConfig** | Pointer to [**V1InstancesScrapeconfigsCreateRequestOauth2TlsConfig**](V1InstancesScrapeconfigsCreateRequestOauth2TlsConfig.md) |  | [optional] 
**BearerToken** | Pointer to **string** | Sets the &#39;Authorization&#39; header on every scrape request with the configured bearer token. It is mutually exclusive with &#39;bearer_token_file&#39;. &#x60;Additional Validators:&#x60; * needs to be a valid bearer token * if bearerToken is in the body no other authentication method should be in the body | [optional] 
**MetricsRelabelConfigs** | Pointer to [**[]V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner**](V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner.md) | List of metric relabel configurations | [optional] 
**Params** | Pointer to **map[string]interface{}** | Optional http params &#x60;Additional Validators:&#x60; * should not contain more than 5 keys * each key and value should not have more than 200 characters | [optional] 
**HttpSDConfigs** | Pointer to [**[]V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner**](V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner.md) | HTTP-based service discovery provides a more generic way to configure static targets and serves as an interface to plug in custom service discovery mechanisms. | [optional] 
**HonorLabels** | Pointer to **bool** | Note that any globally configured &#39;external_labels&#39; are unaffected by this setting. In communication with external systems, they are always applied only when a time series does not have a given label yet and are ignored otherwise. | [optional] [default to false]
**HonorTimeStamps** | Pointer to **bool** | honor_timestamps controls whether Prometheus respects the timestamps present in scraped data. If honor_timestamps is set to &#39;true&#39;, the timestamps of the metrics exposed by the target will be used. | [optional] [default to false]

## Methods

### NewV1InstancesScrapeconfigsCreateRequest

`func NewV1InstancesScrapeconfigsCreateRequest(staticConfigs []V1InstancesScrapeconfigsCreateRequestStaticConfigsInner, jobName string, scheme string, scrapeInterval string, scrapeTimeout string, ) *V1InstancesScrapeconfigsCreateRequest`

NewV1InstancesScrapeconfigsCreateRequest instantiates a new V1InstancesScrapeconfigsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesScrapeconfigsCreateRequestWithDefaults

`func NewV1InstancesScrapeconfigsCreateRequestWithDefaults() *V1InstancesScrapeconfigsCreateRequest`

NewV1InstancesScrapeconfigsCreateRequestWithDefaults instantiates a new V1InstancesScrapeconfigsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStaticConfigs

`func (o *V1InstancesScrapeconfigsCreateRequest) GetStaticConfigs() []V1InstancesScrapeconfigsCreateRequestStaticConfigsInner`

GetStaticConfigs returns the StaticConfigs field if non-nil, zero value otherwise.

### GetStaticConfigsOk

`func (o *V1InstancesScrapeconfigsCreateRequest) GetStaticConfigsOk() (*[]V1InstancesScrapeconfigsCreateRequestStaticConfigsInner, bool)`

GetStaticConfigsOk returns a tuple with the StaticConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticConfigs

`func (o *V1InstancesScrapeconfigsCreateRequest) SetStaticConfigs(v []V1InstancesScrapeconfigsCreateRequestStaticConfigsInner)`

SetStaticConfigs sets StaticConfigs field to given value.


### GetJobName

`func (o *V1InstancesScrapeconfigsCreateRequest) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *V1InstancesScrapeconfigsCreateRequest) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *V1InstancesScrapeconfigsCreateRequest) SetJobName(v string)`

SetJobName sets JobName field to given value.


### GetScheme

`func (o *V1InstancesScrapeconfigsCreateRequest) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *V1InstancesScrapeconfigsCreateRequest) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *V1InstancesScrapeconfigsCreateRequest) SetScheme(v string)`

SetScheme sets Scheme field to given value.


### GetScrapeInterval

`func (o *V1InstancesScrapeconfigsCreateRequest) GetScrapeInterval() string`

GetScrapeInterval returns the ScrapeInterval field if non-nil, zero value otherwise.

### GetScrapeIntervalOk

`func (o *V1InstancesScrapeconfigsCreateRequest) GetScrapeIntervalOk() (*string, bool)`

GetScrapeIntervalOk returns a tuple with the ScrapeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrapeInterval

`func (o *V1InstancesScrapeconfigsCreateRequest) SetScrapeInterval(v string)`

SetScrapeInterval sets ScrapeInterval field to given value.


### GetScrapeTimeout

`func (o *V1InstancesScrapeconfigsCreateRequest) GetScrapeTimeout() string`

GetScrapeTimeout returns the ScrapeTimeout field if non-nil, zero value otherwise.

### GetScrapeTimeoutOk

`func (o *V1InstancesScrapeconfigsCreateRequest) GetScrapeTimeoutOk() (*string, bool)`

GetScrapeTimeoutOk returns a tuple with the ScrapeTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrapeTimeout

`func (o *V1InstancesScrapeconfigsCreateRequest) SetScrapeTimeout(v string)`

SetScrapeTimeout sets ScrapeTimeout field to given value.


### GetMetricsPath

`func (o *V1InstancesScrapeconfigsCreateRequest) GetMetricsPath() string`

GetMetricsPath returns the MetricsPath field if non-nil, zero value otherwise.

### GetMetricsPathOk

`func (o *V1InstancesScrapeconfigsCreateRequest) GetMetricsPathOk() (*string, bool)`

GetMetricsPathOk returns a tuple with the MetricsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsPath

`func (o *V1InstancesScrapeconfigsCreateRequest) SetMetricsPath(v string)`

SetMetricsPath sets MetricsPath field to given value.

### HasMetricsPath

`func (o *V1InstancesScrapeconfigsCreateRequest) HasMetricsPath() bool`

HasMetricsPath returns a boolean if a field has been set.

### GetSampleLimit

`func (o *V1InstancesScrapeconfigsCreateRequest) GetSampleLimit() float32`

GetSampleLimit returns the SampleLimit field if non-nil, zero value otherwise.

### GetSampleLimitOk

`func (o *V1InstancesScrapeconfigsCreateRequest) GetSampleLimitOk() (*float32, bool)`

GetSampleLimitOk returns a tuple with the SampleLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleLimit

`func (o *V1InstancesScrapeconfigsCreateRequest) SetSampleLimit(v float32)`

SetSampleLimit sets SampleLimit field to given value.

### HasSampleLimit

`func (o *V1InstancesScrapeconfigsCreateRequest) HasSampleLimit() bool`

HasSampleLimit returns a boolean if a field has been set.

### GetBasicAuth

`func (o *V1InstancesScrapeconfigsCreateRequest) GetBasicAuth() V1InstancesScrapeconfigsCreateRequestBasicAuth`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *V1InstancesScrapeconfigsCreateRequest) GetBasicAuthOk() (*V1InstancesScrapeconfigsCreateRequestBasicAuth, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *V1InstancesScrapeconfigsCreateRequest) SetBasicAuth(v V1InstancesScrapeconfigsCreateRequestBasicAuth)`

SetBasicAuth sets BasicAuth field to given value.

### HasBasicAuth

`func (o *V1InstancesScrapeconfigsCreateRequest) HasBasicAuth() bool`

HasBasicAuth returns a boolean if a field has been set.

### GetOauth2

`func (o *V1InstancesScrapeconfigsCreateRequest) GetOauth2() V1InstancesScrapeconfigsCreateRequestOauth2`

GetOauth2 returns the Oauth2 field if non-nil, zero value otherwise.

### GetOauth2Ok

`func (o *V1InstancesScrapeconfigsCreateRequest) GetOauth2Ok() (*V1InstancesScrapeconfigsCreateRequestOauth2, bool)`

GetOauth2Ok returns a tuple with the Oauth2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2

`func (o *V1InstancesScrapeconfigsCreateRequest) SetOauth2(v V1InstancesScrapeconfigsCreateRequestOauth2)`

SetOauth2 sets Oauth2 field to given value.

### HasOauth2

`func (o *V1InstancesScrapeconfigsCreateRequest) HasOauth2() bool`

HasOauth2 returns a boolean if a field has been set.

### GetTlsConfig

`func (o *V1InstancesScrapeconfigsCreateRequest) GetTlsConfig() V1InstancesScrapeconfigsCreateRequestOauth2TlsConfig`

GetTlsConfig returns the TlsConfig field if non-nil, zero value otherwise.

### GetTlsConfigOk

`func (o *V1InstancesScrapeconfigsCreateRequest) GetTlsConfigOk() (*V1InstancesScrapeconfigsCreateRequestOauth2TlsConfig, bool)`

GetTlsConfigOk returns a tuple with the TlsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsConfig

`func (o *V1InstancesScrapeconfigsCreateRequest) SetTlsConfig(v V1InstancesScrapeconfigsCreateRequestOauth2TlsConfig)`

SetTlsConfig sets TlsConfig field to given value.

### HasTlsConfig

`func (o *V1InstancesScrapeconfigsCreateRequest) HasTlsConfig() bool`

HasTlsConfig returns a boolean if a field has been set.

### GetBearerToken

`func (o *V1InstancesScrapeconfigsCreateRequest) GetBearerToken() string`

GetBearerToken returns the BearerToken field if non-nil, zero value otherwise.

### GetBearerTokenOk

`func (o *V1InstancesScrapeconfigsCreateRequest) GetBearerTokenOk() (*string, bool)`

GetBearerTokenOk returns a tuple with the BearerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerToken

`func (o *V1InstancesScrapeconfigsCreateRequest) SetBearerToken(v string)`

SetBearerToken sets BearerToken field to given value.

### HasBearerToken

`func (o *V1InstancesScrapeconfigsCreateRequest) HasBearerToken() bool`

HasBearerToken returns a boolean if a field has been set.

### GetMetricsRelabelConfigs

`func (o *V1InstancesScrapeconfigsCreateRequest) GetMetricsRelabelConfigs() []V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner`

GetMetricsRelabelConfigs returns the MetricsRelabelConfigs field if non-nil, zero value otherwise.

### GetMetricsRelabelConfigsOk

`func (o *V1InstancesScrapeconfigsCreateRequest) GetMetricsRelabelConfigsOk() (*[]V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner, bool)`

GetMetricsRelabelConfigsOk returns a tuple with the MetricsRelabelConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsRelabelConfigs

`func (o *V1InstancesScrapeconfigsCreateRequest) SetMetricsRelabelConfigs(v []V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner)`

SetMetricsRelabelConfigs sets MetricsRelabelConfigs field to given value.

### HasMetricsRelabelConfigs

`func (o *V1InstancesScrapeconfigsCreateRequest) HasMetricsRelabelConfigs() bool`

HasMetricsRelabelConfigs returns a boolean if a field has been set.

### GetParams

`func (o *V1InstancesScrapeconfigsCreateRequest) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *V1InstancesScrapeconfigsCreateRequest) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *V1InstancesScrapeconfigsCreateRequest) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *V1InstancesScrapeconfigsCreateRequest) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetHttpSDConfigs

`func (o *V1InstancesScrapeconfigsCreateRequest) GetHttpSDConfigs() []V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner`

GetHttpSDConfigs returns the HttpSDConfigs field if non-nil, zero value otherwise.

### GetHttpSDConfigsOk

`func (o *V1InstancesScrapeconfigsCreateRequest) GetHttpSDConfigsOk() (*[]V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner, bool)`

GetHttpSDConfigsOk returns a tuple with the HttpSDConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpSDConfigs

`func (o *V1InstancesScrapeconfigsCreateRequest) SetHttpSDConfigs(v []V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner)`

SetHttpSDConfigs sets HttpSDConfigs field to given value.

### HasHttpSDConfigs

`func (o *V1InstancesScrapeconfigsCreateRequest) HasHttpSDConfigs() bool`

HasHttpSDConfigs returns a boolean if a field has been set.

### GetHonorLabels

`func (o *V1InstancesScrapeconfigsCreateRequest) GetHonorLabels() bool`

GetHonorLabels returns the HonorLabels field if non-nil, zero value otherwise.

### GetHonorLabelsOk

`func (o *V1InstancesScrapeconfigsCreateRequest) GetHonorLabelsOk() (*bool, bool)`

GetHonorLabelsOk returns a tuple with the HonorLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHonorLabels

`func (o *V1InstancesScrapeconfigsCreateRequest) SetHonorLabels(v bool)`

SetHonorLabels sets HonorLabels field to given value.

### HasHonorLabels

`func (o *V1InstancesScrapeconfigsCreateRequest) HasHonorLabels() bool`

HasHonorLabels returns a boolean if a field has been set.

### GetHonorTimeStamps

`func (o *V1InstancesScrapeconfigsCreateRequest) GetHonorTimeStamps() bool`

GetHonorTimeStamps returns the HonorTimeStamps field if non-nil, zero value otherwise.

### GetHonorTimeStampsOk

`func (o *V1InstancesScrapeconfigsCreateRequest) GetHonorTimeStampsOk() (*bool, bool)`

GetHonorTimeStampsOk returns a tuple with the HonorTimeStamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHonorTimeStamps

`func (o *V1InstancesScrapeconfigsCreateRequest) SetHonorTimeStamps(v bool)`

SetHonorTimeStamps sets HonorTimeStamps field to given value.

### HasHonorTimeStamps

`func (o *V1InstancesScrapeconfigsCreateRequest) HasHonorTimeStamps() bool`

HasHonorTimeStamps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


