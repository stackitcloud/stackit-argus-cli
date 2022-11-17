package models

// https://github.com/prometheus/prometheus/blob/fa6e05903fd3ce52e374a6e1bf4eb98c9f1f45a7/config/config.go#L405

// MetricRelabelConfigs is the configuration for relabeling of target label sets.
type MetricRelabelConfigs struct {
	SourceLabels []string `yaml:"source_labels" json:"sourceLabels"`
	Separator    string   `yaml:"separator" json:"separator"`
	TargetLabel  string   `yaml:"target_label" json:"targetLabel"`
	Regex        string   `yaml:"regex" json:"regex"`
	Modulus      uint64   `yaml:"modulus" json:"modulus"`
	Replacement  string   `yaml:"replacement" json:"replacement"`
	Action       string   `yaml:"action" json:"action"`
}

type StaticConfig struct {
	Targets []string            `yaml:"targets" json:"targets"`
	Labels  map[string][]string `yaml:"labels" json:"labels"`
}

type BasicAuth struct {
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}

type TlsConfig struct {
	InsecureSkipVerify bool `yaml:"insecure_skip_verify" json:"insecureSkipVerify"`
}

type Oauth2 struct {
	ClientId     string    `yaml:"client_id" json:"clientId"`
	ClientSecret string    `yaml:"client_secret" json:"clientSecret"`
	TokenUrl     string    `yaml:"token_url" json:"tokenUrl"`
	Scopes       []string  `yaml:"scopes" json:"scopes"`
	TlsConfig    TlsConfig `yaml:"tls_config" json:"tlsConfig"`
}

type HttpSdConfig struct {
	Url             string    `yaml:"url" json:"url"`
	RefreshInterval string    `yaml:"refresh_interval" json:"refreshInterval"`
	BasicAuth       BasicAuth `yaml:"basic_auth" json:"basicAuth"`
	Oauth2          Oauth2    `yaml:"oauth2" json:"oauth2"`
	TlsConfig       TlsConfig `yaml:"tls_config" json:"tlsConfig"`
}

// ScrapeConfig configures a scraping unit for Prometheus.
type ScrapeConfig struct {
	StaticConfigs        []*StaticConfig         `yaml:"static_configs" json:"staticConfigs"`
	JobName              string                  `yaml:"job_name" json:"jobName"`
	Scheme               string                  `yaml:"scheme" json:"scheme"`
	ScrapeInterval       string                  `yaml:"scrape_interval" json:"scrapeInterval"`
	ScrapeTimeout        string                  `yaml:"scrape_timeout" json:"scrapeTimeout"`
	MetricsPath          string                  `yaml:"metrics_path" json:"metricsPath"`
	SampleLimit          uint                    `yaml:"sample_limit" json:"sampleLimit"`
	BasicAuth            BasicAuth               `yaml:"basic_auth" json:"basicAuth"`
	Oauth2               Oauth2                  `yaml:"oauth2" json:"oauth2"`
	TlsConfig            TlsConfig               `yaml:"tls_config" json:"tlsConfig"`
	BearerToken          string                  `yaml:"bearer_token" json:"bearerToken"`
	MetricRelabelConfigs []*MetricRelabelConfigs `yaml:"metric_relabel_configs" json:"metricRelabelConfigs"`
	HttpSdConfigs        []*HttpSdConfig         `yaml:"http_sd_configs" json:"httpSdConfigs"`
	Params               map[string][]string     `yaml:"params" json:"params"`
	HonorLabels          bool                    `yaml:"honor_labels" json:"honorLabels"`
	HonorTimestamps      bool                    `yaml:"honor_timestamps" json:"honorTimestamps"`
}
