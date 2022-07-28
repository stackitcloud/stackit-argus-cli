# V1InstancesScrapeconfigsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StaticConfigs** | [**[]V1InstancesScrapeconfigsUpdateRequestStaticConfigsInner**](V1InstancesScrapeconfigsUpdateRequestStaticConfigsInner.md) | A list of scrape configurations. | 
**Scheme** | **string** | Configures the protocol scheme used for requests. https or http | 
**ScrapeInterval** | **string** | How frequently to scrape targets from this job. E.g. 5m &#x60;Additional Validators:&#x60; * must be a valid time format* must be &gt;&#x3D; 60s | 
**ScrapeTimeout** | **string** | Per-scrape timeout when scraping this job. &#x60;Additional Validators:&#x60; * must be a valid time format* must be smaller than scrapeInterval | 
**MetricsPath** | **string** | The HTTP resource path on which to fetch metrics from targets. E.g. /metrics | [default to "/metrics"]
**BasicAuth** | Pointer to [**V1InstancesScrapeconfigsCreateRequestBasicAuth**](V1InstancesScrapeconfigsCreateRequestBasicAuth.md) |  | [optional] 
**TlsConfig** | Pointer to [**V1InstancesScrapeconfigsCreateRequestOauth2TlsConfig**](V1InstancesScrapeconfigsCreateRequestOauth2TlsConfig.md) |  | [optional] 
**BearerToken** | Pointer to **string** | Sets the &#39;Authorization&#39; header on every scrape request with the configured bearer token. It is mutually exclusive with &#39;bearer_token_file&#39;. &#x60;Additional Validators:&#x60; * needs to be a valid bearer token * if bearerToken is in the body no other authentication method should be in the body | [optional] 
**MetricsRelabelConfigs** | Pointer to [**[]V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner**](V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner.md) | List of metric relabel configurations | [optional] 
**Params** | Pointer to **map[string]interface{}** | Optional http params &#x60;Additional Validators:&#x60; * should not contain more than 5 keys * each key and value should not have more than 200 characters | [optional] 
**HonorLabels** | Pointer to **bool** | Note that any globally configured &#39;external_labels&#39; are unaffected by this setting. In communication with external systems, they are always applied only when a time series does not have a given label yet and are ignored otherwise. | [optional] [default to false]
**HonorTimeStamps** | Pointer to **bool** | honor_timestamps controls whether Prometheus respects the timestamps present in scraped data. If honor_timestamps is set to &#39;true&#39;, the timestamps of the metrics exposed by the target will be used. | [optional] [default to false]

## Methods

### NewV1InstancesScrapeconfigsUpdateRequest

`func NewV1InstancesScrapeconfigsUpdateRequest(staticConfigs []V1InstancesScrapeconfigsUpdateRequestStaticConfigsInner, scheme string, scrapeInterval string, scrapeTimeout string, metricsPath string, ) *V1InstancesScrapeconfigsUpdateRequest`

NewV1InstancesScrapeconfigsUpdateRequest instantiates a new V1InstancesScrapeconfigsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesScrapeconfigsUpdateRequestWithDefaults

`func NewV1InstancesScrapeconfigsUpdateRequestWithDefaults() *V1InstancesScrapeconfigsUpdateRequest`

NewV1InstancesScrapeconfigsUpdateRequestWithDefaults instantiates a new V1InstancesScrapeconfigsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStaticConfigs

`func (o *V1InstancesScrapeconfigsUpdateRequest) GetStaticConfigs() []V1InstancesScrapeconfigsUpdateRequestStaticConfigsInner`

GetStaticConfigs returns the StaticConfigs field if non-nil, zero value otherwise.

### GetStaticConfigsOk

`func (o *V1InstancesScrapeconfigsUpdateRequest) GetStaticConfigsOk() (*[]V1InstancesScrapeconfigsUpdateRequestStaticConfigsInner, bool)`

GetStaticConfigsOk returns a tuple with the StaticConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticConfigs

`func (o *V1InstancesScrapeconfigsUpdateRequest) SetStaticConfigs(v []V1InstancesScrapeconfigsUpdateRequestStaticConfigsInner)`

SetStaticConfigs sets StaticConfigs field to given value.


### GetScheme

`func (o *V1InstancesScrapeconfigsUpdateRequest) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *V1InstancesScrapeconfigsUpdateRequest) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *V1InstancesScrapeconfigsUpdateRequest) SetScheme(v string)`

SetScheme sets Scheme field to given value.


### GetScrapeInterval

`func (o *V1InstancesScrapeconfigsUpdateRequest) GetScrapeInterval() string`

GetScrapeInterval returns the ScrapeInterval field if non-nil, zero value otherwise.

### GetScrapeIntervalOk

`func (o *V1InstancesScrapeconfigsUpdateRequest) GetScrapeIntervalOk() (*string, bool)`

GetScrapeIntervalOk returns a tuple with the ScrapeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrapeInterval

`func (o *V1InstancesScrapeconfigsUpdateRequest) SetScrapeInterval(v string)`

SetScrapeInterval sets ScrapeInterval field to given value.


### GetScrapeTimeout

`func (o *V1InstancesScrapeconfigsUpdateRequest) GetScrapeTimeout() string`

GetScrapeTimeout returns the ScrapeTimeout field if non-nil, zero value otherwise.

### GetScrapeTimeoutOk

`func (o *V1InstancesScrapeconfigsUpdateRequest) GetScrapeTimeoutOk() (*string, bool)`

GetScrapeTimeoutOk returns a tuple with the ScrapeTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrapeTimeout

`func (o *V1InstancesScrapeconfigsUpdateRequest) SetScrapeTimeout(v string)`

SetScrapeTimeout sets ScrapeTimeout field to given value.


### GetMetricsPath

`func (o *V1InstancesScrapeconfigsUpdateRequest) GetMetricsPath() string`

GetMetricsPath returns the MetricsPath field if non-nil, zero value otherwise.

### GetMetricsPathOk

`func (o *V1InstancesScrapeconfigsUpdateRequest) GetMetricsPathOk() (*string, bool)`

GetMetricsPathOk returns a tuple with the MetricsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsPath

`func (o *V1InstancesScrapeconfigsUpdateRequest) SetMetricsPath(v string)`

SetMetricsPath sets MetricsPath field to given value.


### GetBasicAuth

`func (o *V1InstancesScrapeconfigsUpdateRequest) GetBasicAuth() V1InstancesScrapeconfigsCreateRequestBasicAuth`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *V1InstancesScrapeconfigsUpdateRequest) GetBasicAuthOk() (*V1InstancesScrapeconfigsCreateRequestBasicAuth, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *V1InstancesScrapeconfigsUpdateRequest) SetBasicAuth(v V1InstancesScrapeconfigsCreateRequestBasicAuth)`

SetBasicAuth sets BasicAuth field to given value.

### HasBasicAuth

`func (o *V1InstancesScrapeconfigsUpdateRequest) HasBasicAuth() bool`

HasBasicAuth returns a boolean if a field has been set.

### GetTlsConfig

`func (o *V1InstancesScrapeconfigsUpdateRequest) GetTlsConfig() V1InstancesScrapeconfigsCreateRequestOauth2TlsConfig`

GetTlsConfig returns the TlsConfig field if non-nil, zero value otherwise.

### GetTlsConfigOk

`func (o *V1InstancesScrapeconfigsUpdateRequest) GetTlsConfigOk() (*V1InstancesScrapeconfigsCreateRequestOauth2TlsConfig, bool)`

GetTlsConfigOk returns a tuple with the TlsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsConfig

`func (o *V1InstancesScrapeconfigsUpdateRequest) SetTlsConfig(v V1InstancesScrapeconfigsCreateRequestOauth2TlsConfig)`

SetTlsConfig sets TlsConfig field to given value.

### HasTlsConfig

`func (o *V1InstancesScrapeconfigsUpdateRequest) HasTlsConfig() bool`

HasTlsConfig returns a boolean if a field has been set.

### GetBearerToken

`func (o *V1InstancesScrapeconfigsUpdateRequest) GetBearerToken() string`

GetBearerToken returns the BearerToken field if non-nil, zero value otherwise.

### GetBearerTokenOk

`func (o *V1InstancesScrapeconfigsUpdateRequest) GetBearerTokenOk() (*string, bool)`

GetBearerTokenOk returns a tuple with the BearerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerToken

`func (o *V1InstancesScrapeconfigsUpdateRequest) SetBearerToken(v string)`

SetBearerToken sets BearerToken field to given value.

### HasBearerToken

`func (o *V1InstancesScrapeconfigsUpdateRequest) HasBearerToken() bool`

HasBearerToken returns a boolean if a field has been set.

### GetMetricsRelabelConfigs

`func (o *V1InstancesScrapeconfigsUpdateRequest) GetMetricsRelabelConfigs() []V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner`

GetMetricsRelabelConfigs returns the MetricsRelabelConfigs field if non-nil, zero value otherwise.

### GetMetricsRelabelConfigsOk

`func (o *V1InstancesScrapeconfigsUpdateRequest) GetMetricsRelabelConfigsOk() (*[]V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner, bool)`

GetMetricsRelabelConfigsOk returns a tuple with the MetricsRelabelConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsRelabelConfigs

`func (o *V1InstancesScrapeconfigsUpdateRequest) SetMetricsRelabelConfigs(v []V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner)`

SetMetricsRelabelConfigs sets MetricsRelabelConfigs field to given value.

### HasMetricsRelabelConfigs

`func (o *V1InstancesScrapeconfigsUpdateRequest) HasMetricsRelabelConfigs() bool`

HasMetricsRelabelConfigs returns a boolean if a field has been set.

### GetParams

`func (o *V1InstancesScrapeconfigsUpdateRequest) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *V1InstancesScrapeconfigsUpdateRequest) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *V1InstancesScrapeconfigsUpdateRequest) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *V1InstancesScrapeconfigsUpdateRequest) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetHonorLabels

`func (o *V1InstancesScrapeconfigsUpdateRequest) GetHonorLabels() bool`

GetHonorLabels returns the HonorLabels field if non-nil, zero value otherwise.

### GetHonorLabelsOk

`func (o *V1InstancesScrapeconfigsUpdateRequest) GetHonorLabelsOk() (*bool, bool)`

GetHonorLabelsOk returns a tuple with the HonorLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHonorLabels

`func (o *V1InstancesScrapeconfigsUpdateRequest) SetHonorLabels(v bool)`

SetHonorLabels sets HonorLabels field to given value.

### HasHonorLabels

`func (o *V1InstancesScrapeconfigsUpdateRequest) HasHonorLabels() bool`

HasHonorLabels returns a boolean if a field has been set.

### GetHonorTimeStamps

`func (o *V1InstancesScrapeconfigsUpdateRequest) GetHonorTimeStamps() bool`

GetHonorTimeStamps returns the HonorTimeStamps field if non-nil, zero value otherwise.

### GetHonorTimeStampsOk

`func (o *V1InstancesScrapeconfigsUpdateRequest) GetHonorTimeStampsOk() (*bool, bool)`

GetHonorTimeStampsOk returns a tuple with the HonorTimeStamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHonorTimeStamps

`func (o *V1InstancesScrapeconfigsUpdateRequest) SetHonorTimeStamps(v bool)`

SetHonorTimeStamps sets HonorTimeStamps field to given value.

### HasHonorTimeStamps

`func (o *V1InstancesScrapeconfigsUpdateRequest) HasHonorTimeStamps() bool`

HasHonorTimeStamps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


