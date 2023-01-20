package models

// https://github.com/prometheus/prometheus/blob/fa6e05903fd3ce52e374a6e1bf4eb98c9f1f45a7/config/config.go#L405

// MetricRelabelConfigs is the configuration for relabeling of target label sets.
type MetricRelabelConfigs struct {
	SourceLabels []string `yaml:"source_labels,omitempty" json:"sourceLabels,omitempty"`
	Separator    string   `yaml:"separator,omitempty" json:"separator,omitempty"`
	TargetLabel  string   `yaml:"target_label,omitempty" json:"targetLabel,omitempty"`
	Regex        string   `yaml:"regex,omitempty" json:"regex,omitempty"`
	Modulus      uint64   `yaml:"modulus,omitempty" json:"modulus,omitempty"`
	Replacement  string   `yaml:"replacement,omitempty" json:"replacement,omitempty"`
	Action       string   `yaml:"action,omitempty" json:"action,omitempty"`
}

type StaticConfig struct {
	Targets []string            `yaml:"targets,omitempty" json:"targets,omitempty"`
	Labels  map[string][]string `yaml:"labels,omitempty" json:"labels,omitempty"`
}

type BasicAuth struct {
	Username string `yaml:"username,omitempty" json:"username,omitempty"`
	Password string `yaml:"password,omitempty" json:"password,omitempty"`
}

type TlsConfig struct {
	InsecureSkipVerify bool `yaml:"insecure_skip_verify,omitempty" json:"insecureSkipVerify,omitempty"`
}

type Oauth2 struct {
	ClientId     string     `yaml:"client_id,omitempty" json:"clientId,omitempty"`
	ClientSecret string     `yaml:"client_secret,omitempty" json:"clientSecret,omitempty"`
	TokenUrl     string     `yaml:"token_url,omitempty" json:"tokenUrl,omitempty"`
	Scopes       []string   `yaml:"scopes,omitempty" json:"scopes,omitempty"`
	TlsConfig    *TlsConfig `yaml:"tls_config,omitempty" json:"tlsConfig,omitempty"`
}

type HttpSdConfig struct {
	Url             string     `yaml:"url,omitempty" json:"url,omitempty"`
	RefreshInterval string     `yaml:"refresh_interval,omitempty" json:"refreshInterval,omitempty"`
	BasicAuth       *BasicAuth `yaml:"basic_auth,omitempty" json:"basicAuth,omitempty"`
	Oauth2          *Oauth2    `yaml:"oauth2,omitempty" json:"oauth2,omitempty"`
	TlsConfig       *TlsConfig `yaml:"tls_config,omitempty" json:"tlsConfig,omitempty"`
}

// ScrapeConfig configures a scraping unit for Prometheus.
type ScrapeConfig struct {
	StaticConfigs        []*StaticConfig         `yaml:"static_configs,omitempty" json:"staticConfigs,omitempty"`
	JobName              string                  `yaml:"job_name,omitempty" json:"jobName,omitempty"`
	Scheme               string                  `yaml:"scheme,omitempty" json:"scheme,omitempty"`
	ScrapeInterval       string                  `yaml:"scrape_interval,omitempty" json:"scrapeInterval,omitempty"`
	ScrapeTimeout        string                  `yaml:"scrape_timeout,omitempty" json:"scrapeTimeout,omitempty"`
	MetricsPath          string                  `yaml:"metrics_path,omitempty" json:"metricsPath,omitempty"`
	SampleLimit          uint                    `yaml:"sample_limit,omitempty" json:"sampleLimit,omitempty"`
	BasicAuth            *BasicAuth              `yaml:"basic_auth,omitempty" json:"basicAuth,omitempty"`
	Oauth2               *Oauth2                 `yaml:"oauth2,omitempty" json:"oauth2,omitempty"`
	TlsConfig            *TlsConfig              `yaml:"tls_config,omitempty" json:"tlsConfig,omitempty"`
	BearerToken          string                  `yaml:"bearer_token,omitempty" json:"bearerToken,omitempty"`
	MetricRelabelConfigs []*MetricRelabelConfigs `yaml:"metric_relabel_configs,omitempty" json:"metricRelabelConfigs,omitempty"`
	HttpSdConfigs        []*HttpSdConfig         `yaml:"http_sd_configs,omitempty" json:"httpSdConfigs,omitempty"`
	Params               map[string][]string     `yaml:"params,omitempty" json:"params,omitempty"`
	HonorLabels          bool                    `yaml:"honor_labels,omitempty" json:"honorLabels,omitempty"`
	HonorTimestamps      bool                    `yaml:"honor_timestamps,omitempty" json:"honorTimestamps,omitempty"`
}
