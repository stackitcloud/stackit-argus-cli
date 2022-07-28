/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: v1
Contact: stackit-argus@mail.schwarz
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

import (
	"encoding/json"
)

// V1InstancesScrapeconfigsCreateRequest struct for V1InstancesScrapeconfigsCreateRequest
type V1InstancesScrapeconfigsCreateRequest struct {
	// A list of scrape configurations.
	StaticConfigs []V1InstancesScrapeconfigsCreateRequestStaticConfigsInner `json:"staticConfigs"`
	// The job name assigned to scraped metrics by default. `Additional Validators:` * must be unique * key and values should only include the characters: a-zA-Z0-9-
	JobName string `json:"jobName"`
	// Configures the protocol scheme used for requests. https or http
	Scheme string `json:"scheme"`
	// How frequently to scrape targets from this job. E.g. 5m `Additional Validators:` * must be a valid time format* must be >= 60s
	ScrapeInterval string `json:"scrapeInterval"`
	// Per-scrape timeout when scraping this job. `Additional Validators:` * must be a valid time format* must be smaller than scrapeInterval
	ScrapeTimeout string `json:"scrapeTimeout"`
	// The HTTP resource path on which to fetch metrics from targets. E.g. /metrics
	MetricsPath *string `json:"metricsPath,omitempty"`
	// Per-scrape limit on number of scraped samples that will be accepted. If more than this number of samples are present after metric relabeling the entire scrape will be treated as failed. The total limit depends on the service plan target limits * samples
	SampleLimit *float32 `json:"sampleLimit,omitempty"`
	BasicAuth *V1InstancesScrapeconfigsCreateRequestBasicAuth `json:"basicAuth,omitempty"`
	Oauth2 *V1InstancesScrapeconfigsCreateRequestOauth2 `json:"oauth2,omitempty"`
	TlsConfig *V1InstancesScrapeconfigsCreateRequestOauth2TlsConfig `json:"tlsConfig,omitempty"`
	// Sets the 'Authorization' header on every scrape request with the configured bearer token. It is mutually exclusive with 'bearer_token_file'. `Additional Validators:` * needs to be a valid bearer token * if bearerToken is in the body no other authentication method should be in the body
	BearerToken *string `json:"bearerToken,omitempty"`
	// List of metric relabel configurations
	MetricsRelabelConfigs []V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner `json:"metricsRelabelConfigs,omitempty"`
	// Optional http params `Additional Validators:` * should not contain more than 5 keys * each key and value should not have more than 200 characters
	Params map[string]interface{} `json:"params,omitempty"`
	// HTTP-based service discovery provides a more generic way to configure static targets and serves as an interface to plug in custom service discovery mechanisms.
	HttpSDConfigs []V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner `json:"httpSDConfigs,omitempty"`
	// Note that any globally configured 'external_labels' are unaffected by this setting. In communication with external systems, they are always applied only when a time series does not have a given label yet and are ignored otherwise.
	HonorLabels *bool `json:"honorLabels,omitempty"`
	// honor_timestamps controls whether Prometheus respects the timestamps present in scraped data. If honor_timestamps is set to 'true', the timestamps of the metrics exposed by the target will be used.
	HonorTimeStamps *bool `json:"honorTimeStamps,omitempty"`
}

// NewV1InstancesScrapeconfigsCreateRequest instantiates a new V1InstancesScrapeconfigsCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1InstancesScrapeconfigsCreateRequest(staticConfigs []V1InstancesScrapeconfigsCreateRequestStaticConfigsInner, jobName string, scheme string, scrapeInterval string, scrapeTimeout string) *V1InstancesScrapeconfigsCreateRequest {
	this := V1InstancesScrapeconfigsCreateRequest{}
	this.StaticConfigs = staticConfigs
	this.JobName = jobName
	this.Scheme = scheme
	this.ScrapeInterval = scrapeInterval
	this.ScrapeTimeout = scrapeTimeout
	var metricsPath string = "/metrics"
	this.MetricsPath = &metricsPath
	var honorLabels bool = false
	this.HonorLabels = &honorLabels
	var honorTimeStamps bool = false
	this.HonorTimeStamps = &honorTimeStamps
	return &this
}

// NewV1InstancesScrapeconfigsCreateRequestWithDefaults instantiates a new V1InstancesScrapeconfigsCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1InstancesScrapeconfigsCreateRequestWithDefaults() *V1InstancesScrapeconfigsCreateRequest {
	this := V1InstancesScrapeconfigsCreateRequest{}
	var metricsPath string = "/metrics"
	this.MetricsPath = &metricsPath
	var honorLabels bool = false
	this.HonorLabels = &honorLabels
	var honorTimeStamps bool = false
	this.HonorTimeStamps = &honorTimeStamps
	return &this
}

// GetStaticConfigs returns the StaticConfigs field value
func (o *V1InstancesScrapeconfigsCreateRequest) GetStaticConfigs() []V1InstancesScrapeconfigsCreateRequestStaticConfigsInner {
	if o == nil {
		var ret []V1InstancesScrapeconfigsCreateRequestStaticConfigsInner
		return ret
	}

	return o.StaticConfigs
}

// GetStaticConfigsOk returns a tuple with the StaticConfigs field value
// and a boolean to check if the value has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) GetStaticConfigsOk() ([]V1InstancesScrapeconfigsCreateRequestStaticConfigsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.StaticConfigs, true
}

// SetStaticConfigs sets field value
func (o *V1InstancesScrapeconfigsCreateRequest) SetStaticConfigs(v []V1InstancesScrapeconfigsCreateRequestStaticConfigsInner) {
	o.StaticConfigs = v
}

// GetJobName returns the JobName field value
func (o *V1InstancesScrapeconfigsCreateRequest) GetJobName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobName
}

// GetJobNameOk returns a tuple with the JobName field value
// and a boolean to check if the value has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) GetJobNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobName, true
}

// SetJobName sets field value
func (o *V1InstancesScrapeconfigsCreateRequest) SetJobName(v string) {
	o.JobName = v
}

// GetScheme returns the Scheme field value
func (o *V1InstancesScrapeconfigsCreateRequest) GetScheme() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value
// and a boolean to check if the value has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) GetSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scheme, true
}

// SetScheme sets field value
func (o *V1InstancesScrapeconfigsCreateRequest) SetScheme(v string) {
	o.Scheme = v
}

// GetScrapeInterval returns the ScrapeInterval field value
func (o *V1InstancesScrapeconfigsCreateRequest) GetScrapeInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScrapeInterval
}

// GetScrapeIntervalOk returns a tuple with the ScrapeInterval field value
// and a boolean to check if the value has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) GetScrapeIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScrapeInterval, true
}

// SetScrapeInterval sets field value
func (o *V1InstancesScrapeconfigsCreateRequest) SetScrapeInterval(v string) {
	o.ScrapeInterval = v
}

// GetScrapeTimeout returns the ScrapeTimeout field value
func (o *V1InstancesScrapeconfigsCreateRequest) GetScrapeTimeout() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScrapeTimeout
}

// GetScrapeTimeoutOk returns a tuple with the ScrapeTimeout field value
// and a boolean to check if the value has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) GetScrapeTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScrapeTimeout, true
}

// SetScrapeTimeout sets field value
func (o *V1InstancesScrapeconfigsCreateRequest) SetScrapeTimeout(v string) {
	o.ScrapeTimeout = v
}

// GetMetricsPath returns the MetricsPath field value if set, zero value otherwise.
func (o *V1InstancesScrapeconfigsCreateRequest) GetMetricsPath() string {
	if o == nil || o.MetricsPath == nil {
		var ret string
		return ret
	}
	return *o.MetricsPath
}

// GetMetricsPathOk returns a tuple with the MetricsPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) GetMetricsPathOk() (*string, bool) {
	if o == nil || o.MetricsPath == nil {
		return nil, false
	}
	return o.MetricsPath, true
}

// HasMetricsPath returns a boolean if a field has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) HasMetricsPath() bool {
	if o != nil && o.MetricsPath != nil {
		return true
	}

	return false
}

// SetMetricsPath gets a reference to the given string and assigns it to the MetricsPath field.
func (o *V1InstancesScrapeconfigsCreateRequest) SetMetricsPath(v string) {
	o.MetricsPath = &v
}

// GetSampleLimit returns the SampleLimit field value if set, zero value otherwise.
func (o *V1InstancesScrapeconfigsCreateRequest) GetSampleLimit() float32 {
	if o == nil || o.SampleLimit == nil {
		var ret float32
		return ret
	}
	return *o.SampleLimit
}

// GetSampleLimitOk returns a tuple with the SampleLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) GetSampleLimitOk() (*float32, bool) {
	if o == nil || o.SampleLimit == nil {
		return nil, false
	}
	return o.SampleLimit, true
}

// HasSampleLimit returns a boolean if a field has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) HasSampleLimit() bool {
	if o != nil && o.SampleLimit != nil {
		return true
	}

	return false
}

// SetSampleLimit gets a reference to the given float32 and assigns it to the SampleLimit field.
func (o *V1InstancesScrapeconfigsCreateRequest) SetSampleLimit(v float32) {
	o.SampleLimit = &v
}

// GetBasicAuth returns the BasicAuth field value if set, zero value otherwise.
func (o *V1InstancesScrapeconfigsCreateRequest) GetBasicAuth() V1InstancesScrapeconfigsCreateRequestBasicAuth {
	if o == nil || o.BasicAuth == nil {
		var ret V1InstancesScrapeconfigsCreateRequestBasicAuth
		return ret
	}
	return *o.BasicAuth
}

// GetBasicAuthOk returns a tuple with the BasicAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) GetBasicAuthOk() (*V1InstancesScrapeconfigsCreateRequestBasicAuth, bool) {
	if o == nil || o.BasicAuth == nil {
		return nil, false
	}
	return o.BasicAuth, true
}

// HasBasicAuth returns a boolean if a field has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) HasBasicAuth() bool {
	if o != nil && o.BasicAuth != nil {
		return true
	}

	return false
}

// SetBasicAuth gets a reference to the given V1InstancesScrapeconfigsCreateRequestBasicAuth and assigns it to the BasicAuth field.
func (o *V1InstancesScrapeconfigsCreateRequest) SetBasicAuth(v V1InstancesScrapeconfigsCreateRequestBasicAuth) {
	o.BasicAuth = &v
}

// GetOauth2 returns the Oauth2 field value if set, zero value otherwise.
func (o *V1InstancesScrapeconfigsCreateRequest) GetOauth2() V1InstancesScrapeconfigsCreateRequestOauth2 {
	if o == nil || o.Oauth2 == nil {
		var ret V1InstancesScrapeconfigsCreateRequestOauth2
		return ret
	}
	return *o.Oauth2
}

// GetOauth2Ok returns a tuple with the Oauth2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) GetOauth2Ok() (*V1InstancesScrapeconfigsCreateRequestOauth2, bool) {
	if o == nil || o.Oauth2 == nil {
		return nil, false
	}
	return o.Oauth2, true
}

// HasOauth2 returns a boolean if a field has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) HasOauth2() bool {
	if o != nil && o.Oauth2 != nil {
		return true
	}

	return false
}

// SetOauth2 gets a reference to the given V1InstancesScrapeconfigsCreateRequestOauth2 and assigns it to the Oauth2 field.
func (o *V1InstancesScrapeconfigsCreateRequest) SetOauth2(v V1InstancesScrapeconfigsCreateRequestOauth2) {
	o.Oauth2 = &v
}

// GetTlsConfig returns the TlsConfig field value if set, zero value otherwise.
func (o *V1InstancesScrapeconfigsCreateRequest) GetTlsConfig() V1InstancesScrapeconfigsCreateRequestOauth2TlsConfig {
	if o == nil || o.TlsConfig == nil {
		var ret V1InstancesScrapeconfigsCreateRequestOauth2TlsConfig
		return ret
	}
	return *o.TlsConfig
}

// GetTlsConfigOk returns a tuple with the TlsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) GetTlsConfigOk() (*V1InstancesScrapeconfigsCreateRequestOauth2TlsConfig, bool) {
	if o == nil || o.TlsConfig == nil {
		return nil, false
	}
	return o.TlsConfig, true
}

// HasTlsConfig returns a boolean if a field has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) HasTlsConfig() bool {
	if o != nil && o.TlsConfig != nil {
		return true
	}

	return false
}

// SetTlsConfig gets a reference to the given V1InstancesScrapeconfigsCreateRequestOauth2TlsConfig and assigns it to the TlsConfig field.
func (o *V1InstancesScrapeconfigsCreateRequest) SetTlsConfig(v V1InstancesScrapeconfigsCreateRequestOauth2TlsConfig) {
	o.TlsConfig = &v
}

// GetBearerToken returns the BearerToken field value if set, zero value otherwise.
func (o *V1InstancesScrapeconfigsCreateRequest) GetBearerToken() string {
	if o == nil || o.BearerToken == nil {
		var ret string
		return ret
	}
	return *o.BearerToken
}

// GetBearerTokenOk returns a tuple with the BearerToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) GetBearerTokenOk() (*string, bool) {
	if o == nil || o.BearerToken == nil {
		return nil, false
	}
	return o.BearerToken, true
}

// HasBearerToken returns a boolean if a field has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) HasBearerToken() bool {
	if o != nil && o.BearerToken != nil {
		return true
	}

	return false
}

// SetBearerToken gets a reference to the given string and assigns it to the BearerToken field.
func (o *V1InstancesScrapeconfigsCreateRequest) SetBearerToken(v string) {
	o.BearerToken = &v
}

// GetMetricsRelabelConfigs returns the MetricsRelabelConfigs field value if set, zero value otherwise.
func (o *V1InstancesScrapeconfigsCreateRequest) GetMetricsRelabelConfigs() []V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner {
	if o == nil || o.MetricsRelabelConfigs == nil {
		var ret []V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner
		return ret
	}
	return o.MetricsRelabelConfigs
}

// GetMetricsRelabelConfigsOk returns a tuple with the MetricsRelabelConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) GetMetricsRelabelConfigsOk() ([]V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner, bool) {
	if o == nil || o.MetricsRelabelConfigs == nil {
		return nil, false
	}
	return o.MetricsRelabelConfigs, true
}

// HasMetricsRelabelConfigs returns a boolean if a field has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) HasMetricsRelabelConfigs() bool {
	if o != nil && o.MetricsRelabelConfigs != nil {
		return true
	}

	return false
}

// SetMetricsRelabelConfigs gets a reference to the given []V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner and assigns it to the MetricsRelabelConfigs field.
func (o *V1InstancesScrapeconfigsCreateRequest) SetMetricsRelabelConfigs(v []V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) {
	o.MetricsRelabelConfigs = v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *V1InstancesScrapeconfigsCreateRequest) GetParams() map[string]interface{} {
	if o == nil || o.Params == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) HasParams() bool {
	if o != nil && o.Params != nil {
		return true
	}

	return false
}

// SetParams gets a reference to the given map[string]interface{} and assigns it to the Params field.
func (o *V1InstancesScrapeconfigsCreateRequest) SetParams(v map[string]interface{}) {
	o.Params = v
}

// GetHttpSDConfigs returns the HttpSDConfigs field value if set, zero value otherwise.
func (o *V1InstancesScrapeconfigsCreateRequest) GetHttpSDConfigs() []V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner {
	if o == nil || o.HttpSDConfigs == nil {
		var ret []V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner
		return ret
	}
	return o.HttpSDConfigs
}

// GetHttpSDConfigsOk returns a tuple with the HttpSDConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) GetHttpSDConfigsOk() ([]V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner, bool) {
	if o == nil || o.HttpSDConfigs == nil {
		return nil, false
	}
	return o.HttpSDConfigs, true
}

// HasHttpSDConfigs returns a boolean if a field has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) HasHttpSDConfigs() bool {
	if o != nil && o.HttpSDConfigs != nil {
		return true
	}

	return false
}

// SetHttpSDConfigs gets a reference to the given []V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner and assigns it to the HttpSDConfigs field.
func (o *V1InstancesScrapeconfigsCreateRequest) SetHttpSDConfigs(v []V1InstancesScrapeconfigsCreateRequestHttpSDConfigsInner) {
	o.HttpSDConfigs = v
}

// GetHonorLabels returns the HonorLabels field value if set, zero value otherwise.
func (o *V1InstancesScrapeconfigsCreateRequest) GetHonorLabels() bool {
	if o == nil || o.HonorLabels == nil {
		var ret bool
		return ret
	}
	return *o.HonorLabels
}

// GetHonorLabelsOk returns a tuple with the HonorLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) GetHonorLabelsOk() (*bool, bool) {
	if o == nil || o.HonorLabels == nil {
		return nil, false
	}
	return o.HonorLabels, true
}

// HasHonorLabels returns a boolean if a field has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) HasHonorLabels() bool {
	if o != nil && o.HonorLabels != nil {
		return true
	}

	return false
}

// SetHonorLabels gets a reference to the given bool and assigns it to the HonorLabels field.
func (o *V1InstancesScrapeconfigsCreateRequest) SetHonorLabels(v bool) {
	o.HonorLabels = &v
}

// GetHonorTimeStamps returns the HonorTimeStamps field value if set, zero value otherwise.
func (o *V1InstancesScrapeconfigsCreateRequest) GetHonorTimeStamps() bool {
	if o == nil || o.HonorTimeStamps == nil {
		var ret bool
		return ret
	}
	return *o.HonorTimeStamps
}

// GetHonorTimeStampsOk returns a tuple with the HonorTimeStamps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) GetHonorTimeStampsOk() (*bool, bool) {
	if o == nil || o.HonorTimeStamps == nil {
		return nil, false
	}
	return o.HonorTimeStamps, true
}

// HasHonorTimeStamps returns a boolean if a field has been set.
func (o *V1InstancesScrapeconfigsCreateRequest) HasHonorTimeStamps() bool {
	if o != nil && o.HonorTimeStamps != nil {
		return true
	}

	return false
}

// SetHonorTimeStamps gets a reference to the given bool and assigns it to the HonorTimeStamps field.
func (o *V1InstancesScrapeconfigsCreateRequest) SetHonorTimeStamps(v bool) {
	o.HonorTimeStamps = &v
}

func (o V1InstancesScrapeconfigsCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["staticConfigs"] = o.StaticConfigs
	}
	if true {
		toSerialize["jobName"] = o.JobName
	}
	if true {
		toSerialize["scheme"] = o.Scheme
	}
	if true {
		toSerialize["scrapeInterval"] = o.ScrapeInterval
	}
	if true {
		toSerialize["scrapeTimeout"] = o.ScrapeTimeout
	}
	if o.MetricsPath != nil {
		toSerialize["metricsPath"] = o.MetricsPath
	}
	if o.SampleLimit != nil {
		toSerialize["sampleLimit"] = o.SampleLimit
	}
	if o.BasicAuth != nil {
		toSerialize["basicAuth"] = o.BasicAuth
	}
	if o.Oauth2 != nil {
		toSerialize["oauth2"] = o.Oauth2
	}
	if o.TlsConfig != nil {
		toSerialize["tlsConfig"] = o.TlsConfig
	}
	if o.BearerToken != nil {
		toSerialize["bearerToken"] = o.BearerToken
	}
	if o.MetricsRelabelConfigs != nil {
		toSerialize["metricsRelabelConfigs"] = o.MetricsRelabelConfigs
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	if o.HttpSDConfigs != nil {
		toSerialize["httpSDConfigs"] = o.HttpSDConfigs
	}
	if o.HonorLabels != nil {
		toSerialize["honorLabels"] = o.HonorLabels
	}
	if o.HonorTimeStamps != nil {
		toSerialize["honorTimeStamps"] = o.HonorTimeStamps
	}
	return json.Marshal(toSerialize)
}

type NullableV1InstancesScrapeconfigsCreateRequest struct {
	value *V1InstancesScrapeconfigsCreateRequest
	isSet bool
}

func (v NullableV1InstancesScrapeconfigsCreateRequest) Get() *V1InstancesScrapeconfigsCreateRequest {
	return v.value
}

func (v *NullableV1InstancesScrapeconfigsCreateRequest) Set(val *V1InstancesScrapeconfigsCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1InstancesScrapeconfigsCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1InstancesScrapeconfigsCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1InstancesScrapeconfigsCreateRequest(val *V1InstancesScrapeconfigsCreateRequest) *NullableV1InstancesScrapeconfigsCreateRequest {
	return &NullableV1InstancesScrapeconfigsCreateRequest{value: val, isSet: true}
}

func (v NullableV1InstancesScrapeconfigsCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1InstancesScrapeconfigsCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


