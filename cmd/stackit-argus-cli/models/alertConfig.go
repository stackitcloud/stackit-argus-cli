package models

// https://github.com/prometheus/alertmanager/blob/ba8da18fb2b769ace00d270d677980c4d57310e7/config/config.go#L710

//// TimeInterval describes intervals of time. ContainsTime will tell you if a golang time is contained
//// within the interval.
//type TimeInterval struct {
//	Times       []TimeRange       `yaml:"times,omitempty" json:"times,omitempty"`
//	Weekdays    []WeekdayRange    `yaml:"weekdays,flow,omitempty" json:"weekdays,omitempty"`
//	DaysOfMonth []DayOfMonthRange `yaml:"days_of_month,flow,omitempty" json:"days_of_month,omitempty"`
//	Months      []MonthRange      `yaml:"months,flow,omitempty" json:"months,omitempty"`
//	Years       []YearRange       `yaml:"years,flow,omitempty" json:"years,omitempty"`
//	Location    *Location         `yaml:"location,flow,omitempty" json:"location,omitempty"`
//}
//
//// TimeRange represents a range of minutes within a 1440 minute day, exclusive of the End minute. A day consists of 1440 minutes.
//// For example, 4:00PM to End of the day would Begin at 1020 and End at 1440.
//type TimeRange struct {
//	StartMinute int
//	EndMinute   int
//}
//
//// InclusiveRange is used to hold the Beginning and End values of many time interval components.
//type InclusiveRange struct {
//	Begin int
//	End   int
//}
//
//// A WeekdayRange is an inclusive range between [0, 6] where 0 = Sunday.
//type WeekdayRange struct {
//	InclusiveRange
//}
//
//// A DayOfMonthRange is an inclusive range that may have negative Beginning/End values that represent distance from the End of the month Beginning at -1.
//type DayOfMonthRange struct {
//	InclusiveRange
//}
//
//// A MonthRange is an inclusive range between [1, 12] where 1 = January.
//type MonthRange struct {
//	InclusiveRange
//}
//
//// A YearRange is a positive inclusive range.
//type YearRange struct {
//	InclusiveRange
//}
//
//// A Location is a container for a time.Location, used for custom unmarshalling/validation logic.
//type Location struct {
//	*time.Location
//}
//
//// MuteTimeInterval represents a named set of time intervals for which a route should be muted.
//type MuteTimeInterval struct {
//	Name          string         `yaml:"name" json:"name"`
//	TimeIntervals []TimeInterval `yaml:"time_intervals" json:"time_intervals"`
//}
//
//type SlackConfirmationField struct {
//	Text        string `yaml:"text,omitempty"  json:"text,omitempty"`
//	Title       string `yaml:"title,omitempty"  json:"title,omitempty"`
//	OkText      string `yaml:"ok_text,omitempty"  json:"ok_text,omitempty"`
//	DismissText string `yaml:"dismiss_text,omitempty"  json:"dismiss_text,omitempty"`
//}
//
//type SlackAction struct {
//	Type         string                  `yaml:"type,omitempty"  json:"type,omitempty"`
//	Text         string                  `yaml:"text,omitempty"  json:"text,omitempty"`
//	URL          string                  `yaml:"url,omitempty"   json:"url,omitempty"`
//	Style        string                  `yaml:"style,omitempty" json:"style,omitempty"`
//	Name         string                  `yaml:"name,omitempty"  json:"name,omitempty"`
//	Value        string                  `yaml:"value,omitempty"  json:"value,omitempty"`
//	ConfirmField *SlackConfirmationField `yaml:"confirm,omitempty"  json:"confirm,omitempty"`
//}
//
//type SlackField struct {
//	Title string `yaml:"title,omitempty" json:"title,omitempty"`
//	Value string `yaml:"value,omitempty" json:"value,omitempty"`
//	Short *bool  `yaml:"short,omitempty" json:"short,omitempty"`
//}
//
//// SlackConfig configures notifications via Slack.
//type SlackConfig struct {
//	NotifierConfig `yaml:",inline" json:",inline"`
//
//	HTTPConfig string `yaml:"http_config,omitempty" json:"http_config,omitempty"`
//
//	APIURL     string `yaml:"api_url,omitempty" json:"api_url,omitempty"`
//	APIURLFile string `yaml:"api_url_file,omitempty" json:"api_url_file,omitempty"`
//
//	// Slack channel override, (like #other-channel or @username).
//	Channel  string `yaml:"channel,omitempty" json:"channel,omitempty"`
//	Username string `yaml:"username,omitempty" json:"username,omitempty"`
//	Color    string `yaml:"color,omitempty" json:"color,omitempty"`
//
//	Title       string         `yaml:"title,omitempty" json:"title,omitempty"`
//	TitleLink   string         `yaml:"title_link,omitempty" json:"title_link,omitempty"`
//	Pretext     string         `yaml:"pretext,omitempty" json:"pretext,omitempty"`
//	Text        string         `yaml:"text,omitempty" json:"text,omitempty"`
//	Fields      []*SlackField  `yaml:"fields,omitempty" json:"fields,omitempty"`
//	ShortFields bool           `yaml:"short_fields" json:"short_fields,omitempty"`
//	Footer      string         `yaml:"footer,omitempty" json:"footer,omitempty"`
//	Fallback    string         `yaml:"fallback,omitempty" json:"fallback,omitempty"`
//	CallbackID  string         `yaml:"callback_id,omitempty" json:"callback_id,omitempty"`
//	IconEmoji   string         `yaml:"icon_emoji,omitempty" json:"icon_emoji,omitempty"`
//	IconURL     string         `yaml:"icon_url,omitempty" json:"icon_url,omitempty"`
//	ImageURL    string         `yaml:"image_url,omitempty" json:"image_url,omitempty"`
//	ThumbURL    string         `yaml:"thumb_url,omitempty" json:"thumb_url,omitempty"`
//	LinkNames   bool           `yaml:"link_names" json:"link_names,omitempty"`
//	MrkdwnIn    []string       `yaml:"mrkdwn_in,omitempty" json:"mrkdwn_in,omitempty"`
//	Actions     []*SlackAction `yaml:"actions,omitempty" json:"actions,omitempty"`
//}
//
//// PagerdutyLink is a link
//type PagerdutyLink struct {
//	Href string `yaml:"href,omitempty" json:"href,omitempty"`
//	Text string `yaml:"text,omitempty" json:"text,omitempty"`
//}
//
//// PagerdutyImage is an image
//type PagerdutyImage struct {
//	Src  string `yaml:"src,omitempty" json:"src,omitempty"`
//	Alt  string `yaml:"alt,omitempty" json:"alt,omitempty"`
//	Href string `yaml:"href,omitempty" json:"href,omitempty"`
//}
//
//// PagerdutyConfig configures notifications via PagerDuty.
//type PagerdutyConfig struct {
//	NotifierConfig `yaml:",inline" json:",inline"`
//
//	HTTPConfig string `yaml:"http_config,omitempty" json:"http_config,omitempty"`
//
//	ServiceKey     string            `yaml:"service_key,omitempty" json:"service_key,omitempty"`
//	ServiceKeyFile string            `yaml:"service_key_file,omitempty" json:"service_key_file,omitempty"`
//	RoutingKey     string            `yaml:"routing_key,omitempty" json:"routing_key,omitempty"`
//	RoutingKeyFile string            `yaml:"routing_key_file,omitempty" json:"routing_key_file,omitempty"`
//	URL            string            `yaml:"url,omitempty" json:"url,omitempty"`
//	Client         string            `yaml:"client,omitempty" json:"client,omitempty"`
//	ClientURL      string            `yaml:"client_url,omitempty" json:"client_url,omitempty"`
//	Description    string            `yaml:"description,omitempty" json:"description,omitempty"`
//	Details        map[string]string `yaml:"details,omitempty" json:"details,omitempty"`
//	Images         []PagerdutyImage  `yaml:"images,omitempty" json:"images,omitempty"`
//	Links          []PagerdutyLink   `yaml:"links,omitempty" json:"links,omitempty"`
//	Source         string            `yaml:"source,omitempty" json:"source,omitempty"`
//	Severity       string            `yaml:"severity,omitempty" json:"severity,omitempty"`
//	Class          string            `yaml:"class,omitempty" json:"class,omitempty"`
//	Component      string            `yaml:"component,omitempty" json:"component,omitempty"`
//	Group          string            `yaml:"group,omitempty" json:"group,omitempty"`
//}
//
//type TlsConfig struct {
//	InsecureSkipVerify bool `yaml:"insecure_skip_verify"`
//}
//
//// NotifierConfig contains base options common across all notifier configurations.
//type NotifierConfig struct {
//	VSendResolved bool `yaml:"send_resolved" json:"send_resolved"`
//}
//
//// DiscordConfig configures notifications via Discord.
//type DiscordConfig struct {
//	NotifierConfig `yaml:",inline" json:",inline"`
//	HTTPConfig     string `yaml:"http_config,omitempty" json:"http_config,omitempty"`
//	WebhookURL     string `yaml:"webhook_url,omitempty" json:"webhook_url,omitempty"`
//	Title          string `yaml:"title,omitempty" json:"title,omitempty"`
//	Message        string `yaml:"message,omitempty" json:"message,omitempty"`
//}
//
//type OpsGenieConfigResponder struct {
//	// One of those 3 should be filled.
//	ID       string `yaml:"id,omitempty" json:"id,omitempty"`
//	Name     string `yaml:"name,omitempty" json:"name,omitempty"`
//	Username string `yaml:"username,omitempty" json:"username,omitempty"`
//
//	// team, user, escalation, schedule etc.
//	Type string `yaml:"type,omitempty" json:"type,omitempty"`
//}
//
//// WechatConfig configures notifications via Wechat.
//type WechatConfig struct {
//	NotifierConfig `yaml:",inline" json:",inline"`
//
//	HTTPConfig string `yaml:"http_config,omitempty" json:"http_config,omitempty"`
//
//	APISecret   string `yaml:"api_secret,omitempty" json:"api_secret,omitempty"`
//	CorpID      string `yaml:"corp_id,omitempty" json:"corp_id,omitempty"`
//	Message     string `yaml:"message,omitempty" json:"message,omitempty"`
//	APIURL      string `yaml:"api_url,omitempty" json:"api_url,omitempty"`
//	ToUser      string `yaml:"to_user,omitempty" json:"to_user,omitempty"`
//	ToParty     string `yaml:"to_party,omitempty" json:"to_party,omitempty"`
//	ToTag       string `yaml:"to_tag,omitempty" json:"to_tag,omitempty"`
//	AgentID     string `yaml:"agent_id,omitempty" json:"agent_id,omitempty"`
//	MessageType string `yaml:"message_type,omitempty" json:"message_type,omitempty"`
//}
//
//type PushoverConfig struct {
//	NotifierConfig `yaml:",inline" json:",inline"`
//
//	HTTPConfig string `yaml:"http_config,omitempty" json:"http_config,omitempty"`
//
//	UserKey  string `yaml:"user_key,omitempty" json:"user_key,omitempty"`
//	Token    string `yaml:"token,omitempty" json:"token,omitempty"`
//	Title    string `yaml:"title,omitempty" json:"title,omitempty"`
//	Message  string `yaml:"message,omitempty" json:"message,omitempty"`
//	URL      string `yaml:"url,omitempty" json:"url,omitempty"`
//	URLTitle string `yaml:"url_title,omitempty" json:"url_title,omitempty"`
//	Sound    string `yaml:"sound,omitempty" json:"sound,omitempty"`
//	Priority string `yaml:"priority,omitempty" json:"priority,omitempty"`
//	Retry    string `yaml:"retry,omitempty" json:"retry,omitempty"`
//	Expire   string `yaml:"expire,omitempty" json:"expire,omitempty"`
//	HTML     bool   `yaml:"html" json:"html,omitempty"`
//}
//
//// VictorOpsConfig configures notifications via VictorOps.
//type VictorOpsConfig struct {
//	NotifierConfig `yaml:",inline" json:",inline"`
//
//	HTTPConfig string `yaml:"http_config,omitempty" json:"http_config,omitempty"`
//
//	APIKey            string            `yaml:"api_key,omitempty" json:"api_key,omitempty"`
//	APIKeyFile        string            `yaml:"api_key_file,omitempty" json:"api_key_file,omitempty"`
//	APIURL            string            `yaml:"api_url" json:"api_url"`
//	RoutingKey        string            `yaml:"routing_key" json:"routing_key"`
//	MessageType       string            `yaml:"message_type" json:"message_type"`
//	StateMessage      string            `yaml:"state_message" json:"state_message"`
//	EntityDisplayName string            `yaml:"entity_display_name" json:"entity_display_name"`
//	MonitoringTool    string            `yaml:"monitoring_tool" json:"monitoring_tool"`
//	CustomFields      map[string]string `yaml:"custom_fields,omitempty" json:"custom_fields,omitempty"`
//}
//
//// TelegramConfig configures notifications via Telegram.
//type TelegramConfig struct {
//	NotifierConfig `yaml:",inline" json:",inline"`
//
//	HTTPConfig string `yaml:"http_config,omitempty" json:"http_config,omitempty"`
//
//	APIUrl               string `yaml:"api_url" json:"api_url,omitempty"`
//	BotToken             string `yaml:"bot_token,omitempty" json:"token,omitempty"`
//	ChatID               int64  `yaml:"chat_id,omitempty" json:"chat,omitempty"`
//	Message              string `yaml:"message,omitempty" json:"message,omitempty"`
//	DisableNotifications bool   `yaml:"disable_notifications,omitempty" json:"disable_notifications,omitempty"`
//	ParseMode            string `yaml:"parse_mode,omitempty" json:"parse_mode,omitempty"`
//}
//
//type SNSConfig struct {
//	NotifierConfig `yaml:",inline" json:",inline"`
//
//	HTTPConfig string `yaml:"http_config,omitempty" json:"http_config,omitempty"`
//
//	APIUrl      string            `yaml:"api_url,omitempty" json:"api_url,omitempty"`
//	Sigv4       string            `yaml:"sigv4" json:"sigv4"`
//	TopicARN    string            `yaml:"topic_arn,omitempty" json:"topic_arn,omitempty"`
//	PhoneNumber string            `yaml:"phone_number,omitempty" json:"phone_number,omitempty"`
//	TargetARN   string            `yaml:"target_arn,omitempty" json:"target_arn,omitempty"`
//	Subject     string            `yaml:"subject,omitempty" json:"subject,omitempty"`
//	Message     string            `yaml:"message,omitempty" json:"message,omitempty"`
//	Attributes  map[string]string `yaml:"attributes,omitempty" json:"attributes,omitempty"`
//}

// InhibitRule defines an inhibition rule that mutes alerts that match the
// target labels if an alert matching the source labels exists.
// Both alerts have to have a set of labels being equal.
type InhibitRule struct {
	SourceMatch   map[string]string `yaml:"source_match" json:"sourceMatch"`
	SourceMatchRE map[string]string `yaml:"source_match_re" json:"sourceMatchRe"`
	TargetMatch   map[string]string `yaml:"target_match" json:"targetMatch"`
	TargetMatchRE map[string]string `yaml:"target_match_re" json:"targetMatchRe"`
	Equal         []string          `yaml:"equal" json:"equal"`

	// not supported

	//SourceMatchers []string `yaml:"source_matchers"`
	//TargetMatchers []string `yaml:"target_matchers"`
}

// OpsGenieConfig configures notifications via OpsGenie.
type OpsGenieConfig struct {
	APIKey string `yaml:"api_key" json:"apiKey"`
	APIURL string `yaml:"api_url" json:"apiUrl"`
	Tags   string `yaml:"tags" json:"tags"`

	// not supported

	//NotifierConfig `yaml:",inline"`
	//HTTPConfig     string                    `yaml:"http_config"`
	//APIKeyFile     string                    `yaml:"api_key_file"`
	//Message        string                    `yaml:"message"`
	//Description    string                    `yaml:"description"`
	//Source         string                    `yaml:"source"`
	//Details        map[string]string         `yaml:"details"`
	//Entity         string                    `yaml:"entity"`
	//Responders     []OpsGenieConfigResponder `yaml:"responders"`
	//Actions        string                    `yaml:"actions"`
	//Note           string                    `yaml:"note"`
	//Priority       string                    `yaml:"priority"`
	//UpdateAlerts   bool                      `yaml:"update_alerts"`
}

// EmailConfig configures notifications via mail.
type EmailConfig struct {
	To           string `yaml:"to" json:"to"`
	From         string `yaml:"from" json:"from"`
	Smarthost    string `yaml:"smarthost" json:"smarthost"`
	AuthUsername string `yaml:"auth_username" json:"authUsername"`
	AuthPassword string `yaml:"auth_password" json:"authPassword"`
	AuthIdentity string `yaml:"auth_identity" json:"authIdentity"`

	// not supported

	//Hello            string            `yaml:"hello"`
	//AuthPasswordFile string            `yaml:"auth_password_file"`
	//AuthSecret       string            `yaml:"auth_secret"`
	//Headers          map[string]string `yaml:"headers"`
	//HTML             string            `yaml:"html"`
	//Text             string            `yaml:"text"`
	//RequireTLS       *bool             `yaml:"require_tls"`
	//TLSConfig        TlsConfig         `yaml:"tls_config"`
}

// WebhookConfig configures notifications via a generic webhook.
type WebhookConfig struct {
	URL     string `yaml:"url" json:"url"`
	MsTeams bool   `yaml:"ms_teams" json:"msTeams"`

	// not supported

	//NotifierConfig `yaml:",inline"`
	//HTTPConfig     string `yaml:"http_config"`
	//MaxAlerts      uint64 `yaml:"max_alerts"`
}

// Receiver configuration provides configuration on how to contact a receiver.
type Receiver struct {
	Name            string            `yaml:"name" json:"name"`
	EmailConfigs    []*EmailConfig    `yaml:"email_configs" json:"emailConfigs"`
	OpsGenieConfigs []*OpsGenieConfig `yaml:"opsgenie_configs" json:"opsgenieConfigs"`
	WebhookConfigs  []*WebhookConfig  `yaml:"webhook_configs" json:"webhookConfigs"`

	// not supported

	//DiscordConfigs   []*DiscordConfig   `yaml:"discord_configs"`
	//PagerdutyConfigs []*PagerdutyConfig `yaml:"pagerduty_configs"`
	//SlackConfigs     []*SlackConfig     `yaml:"slack_configs"`
	//WechatConfigs    []*WechatConfig    `yaml:"wechat_configs"`
	//PushoverConfigs  []*PushoverConfig  `yaml:"pushover_configs"`
	//VictorOpsConfigs []*VictorOpsConfig `yaml:"victorops_configs"`
	//SNSConfigs       []*SNSConfig       `yaml:"sns_configs"`
	//TelegramConfigs  []*TelegramConfig  `yaml:"telegram_configs"`
}

// A Route is a node that contains definitions of how to handle alerts.
type Route struct {
	Receiver       string            `yaml:"receiver" json:"receiver"`
	GroupBy        []string          `yaml:"group_by" json:"groupBy"`
	GroupWait      string            `yaml:"group_wait" json:"groupWait"`
	GroupInterval  string            `yaml:"group_interval" json:"groupInterval"`
	RepeatInterval string            `yaml:"repeat_interval" json:"repeatInterval"`
	Match          map[string]string `yaml:"match" json:"match"`
	MatchRE        map[string]string `yaml:"match_re" json:"matchRe"`
	Matchers       []string          `yaml:"matchers" json:"matchers"`
	Routes         []*Route          `yaml:"routes" json:"routes"`

	// not supported

	//Continue            bool     `yaml:"continue"`
	//MuteTimeIntervals   []string `yaml:"mute_time_intervals"`
	//ActiveTimeIntervals []string `yaml:"active_time_intervals"`
}

// GlobalConfig defines configuration parameters that are valid globally
// unless overwritten.
type GlobalConfig struct {
	ResolveTimeout   string `yaml:"resolve_timeout" json:"resolveTimeout"`
	SMTPFrom         string `yaml:"smtp_from" json:"smtpFrom"`
	SMTPSmarthost    string `yaml:"smtp_smarthost" json:"smtpSmarthost"`
	SMTPAuthUsername string `yaml:"smtp_auth_username" json:"smtpAuthUsername"`
	SMTPAuthPassword string `yaml:"smtp_auth_password" json:"smtpAuthPassword"`
	SMTPAuthIdentity string `yaml:"smtp_auth_identity" json:"smtpAuthIdentity"`
	OpsGenieAPIKey   string `yaml:"opsgenie_api_key" json:"opsgenieApiKey"`
	OpsGenieAPIURL   string `yaml:"opsgenie_api_url" json:"opsgenieApiUrl"`

	// not supported

	//HTTPConfig           string `yaml:"http_config"`
	//SMTPHello            string `yaml:"smtp_hello"`
	//SMTPAuthPasswordFile string `yaml:"smtp_auth_password_file"`
	//SMTPAuthSecret       string `yaml:"smtp_auth_secret"`
	//SMTPRequireTLS       bool   `yaml:"smtp_require_tls"`
	//SlackAPIURL          string `yaml:"slack_api_url"`
	//SlackAPIURLFile      string `yaml:"slack_api_url_file"`
	//PagerdutyURL         string `yaml:"pagerduty_url"`
	//OpsGenieAPIKeyFile   string `yaml:"opsgenie_api_key_file"`
	//WeChatAPIURL         string `yaml:"wechat_api_url"`
	//WeChatAPISecret      string `yaml:"wechat_api_secret"`
	//WeChatAPICorpID      string `yaml:"wechat_api_corp_id"`
	//VictorOpsAPIURL      string `yaml:"victorops_api_url"`
	//VictorOpsAPIKey      string `yaml:"victorops_api_key"`
	//VictorOpsAPIKeyFile  string `yaml:"victorops_api_key_file"`
	//TelegramAPIUrl       string `yaml:"telegram_api_url"`
}

// Config is the top-level configuration for Alertmanager's config files.
type Config struct {
	Global       *GlobalConfig  `yaml:"global" json:"global"`
	Route        *Route         `yaml:"route" json:"route"`
	Receivers    []*Receiver    `yaml:"receivers" json:"receivers"`
	InhibitRules []*InhibitRule `yaml:"inhibit_rules" json:"inhibitRules"`

	// not supported

	//Templates         []string           `yaml:"templates"`
	//MuteTimeIntervals []MuteTimeInterval `yaml:"mute_time_intervals,omitempty"`
	//TimeIntervals     []TimeInterval     `yaml:"time_intervals,omitempty"`
}
