package models

// https://github.com/prometheus/alertmanager/blob/ba8da18fb2b769ace00d270d677980c4d57310e7/config/config.go#L710

// InhibitRule defines an inhibition rule that mutes alerts that match the
// target labels if an alert matching the source labels exists.
// Both alerts have to have a set of labels being equal.
type InhibitRule struct {
	SourceMatch   map[string]string `yaml:"source_match,omitempty" json:"sourceMatch,omitempty"`
	SourceMatchRE map[string]string `yaml:"source_match_re,omitempty" json:"sourceMatchRe,omitempty"`
	TargetMatch   map[string]string `yaml:"target_match,omitempty" json:"targetMatch,omitempty"`
	TargetMatchRE map[string]string `yaml:"target_match_re,omitempty" json:"targetMatchRe,omitempty"`
	Equal         []string          `yaml:"equal,omitempty" json:"equal,omitempty"`
}

// OpsGenieConfig configures notifications via OpsGenie.
type OpsGenieConfig struct {
	APIKey string `yaml:"api_key,omitempty" json:"apiKey,omitempty"`
	APIURL string `yaml:"api_url,omitempty" json:"apiUrl,omitempty"`
	Tags   string `yaml:"tags,omitempty" json:"tags,omitempty"`
}

// EmailConfig configures notifications via mail.
type EmailConfig struct {
	To           string `yaml:"to,omitempty" json:"to,omitempty"`
	From         string `yaml:"from,omitempty" json:"from,omitempty"`
	Smarthost    string `yaml:"smarthost,omitempty" json:"smarthost,omitempty"`
	AuthUsername string `yaml:"auth_username,omitempty" json:"authUsername,omitempty"`
	AuthPassword string `yaml:"auth_password,omitempty" json:"authPassword,omitempty"`
	AuthIdentity string `yaml:"auth_identity,omitempty" json:"authIdentity,omitempty"`
}

// WebhookConfig configures notifications via a generic webhook.
type WebhookConfig struct {
	URL     string `yaml:"url,omitempty" json:"url,omitempty"`
	MsTeams bool   `yaml:"ms_teams,omitempty" json:"msTeams,omitempty"`
}

// Receiver configuration provides configuration on how to contact a receiver.
type Receiver struct {
	Name            string            `yaml:"name,omitempty" json:"name,omitempty"`
	EmailConfigs    []*EmailConfig    `yaml:"email_configs,omitempty" json:"emailConfigs,omitempty"`
	OpsGenieConfigs []*OpsGenieConfig `yaml:"opsgenie_configs,omitempty" json:"opsgenieConfigs,omitempty"`
	WebhookConfigs  []*WebhookConfig  `yaml:"webhook_configs,omitempty" json:"webhookConfigs,omitempty"`
}

// A Route is a node that contains definitions of how to handle alerts.
type Route struct {
	Receiver       string            `yaml:"receiver,omitempty" json:"receiver,omitempty"`
	GroupBy        []string          `yaml:"group_by,omitempty" json:"groupBy,omitempty"`
	GroupWait      string            `yaml:"group_wait,omitempty" json:"groupWait,omitempty"`
	GroupInterval  string            `yaml:"group_interval,omitempty" json:"groupInterval,omitempty"`
	RepeatInterval string            `yaml:"repeat_interval,omitempty" json:"repeatInterval,omitempty"`
	Match          map[string]string `yaml:"match,omitempty" json:"match,omitempty"`
	MatchRE        map[string]string `yaml:"match_re,omitempty" json:"matchRe,omitempty"`
	Matchers       []string          `yaml:"matchers,omitempty" json:"matchers,omitempty"`
	Routes         []*Route          `yaml:"routes,omitempty" json:"routes,omitempty"`
}

// GlobalConfig defines configuration parameters that are valid globally
// unless overwritten.
type GlobalConfig struct {
	ResolveTimeout   string `yaml:"resolve_timeout,omitempty" json:"resolveTimeout,omitempty"`
	SMTPFrom         string `yaml:"smtp_from,omitempty" json:"smtpFrom,omitempty"`
	SMTPSmarthost    string `yaml:"smtp_smarthost,omitempty" json:"smtpSmarthost,omitempty"`
	SMTPAuthUsername string `yaml:"smtp_auth_username,omitempty" json:"smtpAuthUsername,omitempty"`
	SMTPAuthPassword string `yaml:"smtp_auth_password,omitempty" json:"smtpAuthPassword,omitempty"`
	SMTPAuthIdentity string `yaml:"smtp_auth_identity,omitempty" json:"smtpAuthIdentity,omitempty"`
	OpsGenieAPIKey   string `yaml:"opsgenie_api_key,omitempty" json:"opsgenieApiKey,omitempty"`
	OpsGenieAPIURL   string `yaml:"opsgenie_api_url,omitempty" json:"opsgenieApiUrl,omitempty"`
}

// Config is the top-level configuration for Alertmanager's config files.
type Config struct {
	Global       *GlobalConfig  `yaml:"global,omitempty" json:"global,omitempty"`
	Route        *Route         `yaml:"route,omitempty" json:"route,omitempty"`
	Receivers    []*Receiver    `yaml:"receivers,omitempty" json:"receivers,omitempty"`
	InhibitRules []*InhibitRule `yaml:"inhibit_rules,omitempty" json:"inhibitRules,omitempty"`
}
