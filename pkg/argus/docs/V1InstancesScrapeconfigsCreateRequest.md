# V1InstancesScrapeconfigsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StaticConfigs** | [**[]V1InstancesScrapeconfigsCreateRequestStaticConfigsInner**](V1InstancesScrapeconfigsCreateRequestStaticConfigsInner.md) | A list of scrape configurations. | 
**ProxyUrl** | Pointer to **string** | Proxy url | [optional] 
**JobName** | **string** | The job name assigned to scraped metrics by default. | 
**Scheme** | **string** | Configures the protocol scheme used for requests. https or http | 
**ScrapeInterval** | **string** | How frequently to scrape targets from this job. E.g. 5m | 
**ScrapeTimeout** | **string** | Per-scrape timeout when scraping this job. | 
**MetricsPath** | **string** | The HTTP resource path on which to fetch metrics from targets. E.g. /metrics | [default to "/metrics"]
**BasicAuth** | Pointer to [**V1InstancesScrapeconfigsCreateRequestBasicAuth**](V1InstancesScrapeconfigsCreateRequestBasicAuth.md) |  | [optional] 
**TlsConfig** | Pointer to [**V1InstancesScrapeconfigsCreateRequestTlsConfig**](V1InstancesScrapeconfigsCreateRequestTlsConfig.md) |  | [optional] 
**BearerToken** | Pointer to **string** | Sets the &#x60;Authorization&#x60; header on every scrape request with the configured bearer token. It is mutually exclusive with &#x60;bearer_token_file&#x60;. | [optional] 
**MetricsRelabelConfigs** | Pointer to [**[]V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner**](V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner.md) | List of metric relabel configurations | [optional] 
**Params** | Pointer to [**V1InstancesScrapeconfigsCreateRequestParams**](V1InstancesScrapeconfigsCreateRequestParams.md) |  | [optional] 
**HonorLabels** | Pointer to **bool** | Note that any globally configured &#39;external_labels&#39; are unaffected by this setting. In communication with external systems, they are always applied only when a time series does not have a given label yet and are ignored otherwise. | [optional] [default to false]
**HonorTimeStamps** | Pointer to **bool** | honor_timestamps controls whether Prometheus respects the timestamps present in scraped data. If honor_timestamps is set to &#39;true&#39;, the timestamps of the metrics exposed by the target will be used. | [optional] [default to false]

## Methods

### NewV1InstancesScrapeconfigsCreateRequest

`func NewV1InstancesScrapeconfigsCreateRequest(staticConfigs []V1InstancesScrapeconfigsCreateRequestStaticConfigsInner, jobName string, scheme string, scrapeInterval string, scrapeTimeout string, metricsPath string, ) *V1InstancesScrapeconfigsCreateRequest`

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


### GetProxyUrl

`func (o *V1InstancesScrapeconfigsCreateRequest) GetProxyUrl() string`

GetProxyUrl returns the ProxyUrl field if non-nil, zero value otherwise.

### GetProxyUrlOk

`func (o *V1InstancesScrapeconfigsCreateRequest) GetProxyUrlOk() (*string, bool)`

GetProxyUrlOk returns a tuple with the ProxyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUrl

`func (o *V1InstancesScrapeconfigsCreateRequest) SetProxyUrl(v string)`

SetProxyUrl sets ProxyUrl field to given value.

### HasProxyUrl

`func (o *V1InstancesScrapeconfigsCreateRequest) HasProxyUrl() bool`

HasProxyUrl returns a boolean if a field has been set.

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

### GetTlsConfig

`func (o *V1InstancesScrapeconfigsCreateRequest) GetTlsConfig() V1InstancesScrapeconfigsCreateRequestTlsConfig`

GetTlsConfig returns the TlsConfig field if non-nil, zero value otherwise.

### GetTlsConfigOk

`func (o *V1InstancesScrapeconfigsCreateRequest) GetTlsConfigOk() (*V1InstancesScrapeconfigsCreateRequestTlsConfig, bool)`

GetTlsConfigOk returns a tuple with the TlsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsConfig

`func (o *V1InstancesScrapeconfigsCreateRequest) SetTlsConfig(v V1InstancesScrapeconfigsCreateRequestTlsConfig)`

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

`func (o *V1InstancesScrapeconfigsCreateRequest) GetParams() V1InstancesScrapeconfigsCreateRequestParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *V1InstancesScrapeconfigsCreateRequest) GetParamsOk() (*V1InstancesScrapeconfigsCreateRequestParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *V1InstancesScrapeconfigsCreateRequest) SetParams(v V1InstancesScrapeconfigsCreateRequestParams)`

SetParams sets Params field to given value.

### HasParams

`func (o *V1InstancesScrapeconfigsCreateRequest) HasParams() bool`

HasParams returns a boolean if a field has been set.

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


