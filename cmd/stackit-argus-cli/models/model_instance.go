package models

import (
	"go/types"
)

type InstanceList struct {
	Message   string             `json:"message"`
	Instances []InstanceListItem `json:"instances"`
}

type InstanceListItem struct {
	Id          string `json:"id"`
	PlanName    string `json:"planName"`
	Instance    string `json:"instance"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	Error       string `json:"error"`
	ServiceName string `json:"serviceName"`
}

type ProjectInstance struct {
	Message      string           `json:"message"`
	DashboardUrl string           `json:"dashboardUrl"`
	IsUpdatable  bool             `json:"isUpdatable"`
	Name         string           `json:"name"`
	Parameters   Parameters       `json:"parameters"`
	Id           string           `json:"id"`
	ServiceName  string           `json:"serviceName"`
	PlanId       string           `json:"planId"`
	PlanName     string           `json:"planName"`
	PlanSchema   string           `json:"planSchema"`
	Status       string           `json:"status"`
	Instance     InstanceResponse `json:"instance"`
}

type Parameters struct {
}

type InstanceResponse struct {
	Instance                string      `json:"instance"`
	Cluster                 string      `json:"cluster"`
	GrafanaUrl              string      `json:"grafanaUrl"`
	DashboardUrl            string      `json:"dashboardUrl"`
	GrafanaPlugins          types.Array `json:"grafanaPlugins"`
	Name                    string      `json:"name"`
	GrafanaAdminPassword    string      `json:"grafanaAdminPassword"`
	GrafanaAdminUser        string      `json:"grafanaAdminUser"`
	MetricsRetentionTimeRaw int         `json:"metricsRetentionTimeRaw"`
	MetricsRetentionTime5m  int         `json:"metricsRetentionTime5m"`
	MetricsRetentionTime1h  int         `json:"metricsRetentionTime1h"`
	MetricsUrl              string      `json:"metricsUrl"`
	PushMetricsUrl          string      `json:"pushMetricsUrl"`
	GrafanaPublicReadAccess bool        `json:"grafanaPublicReadAccess"`
	TargetsUrl              string      `json:"targetsUrl"`
	AlertingUrl             string      `json:"alertingUrl"`
	Plan                    Plan        `json:"plan"`
	LogsUrl                 string      `json:"logsUrl"`
	LogsPushUrl             string      `json:"logsPushUrl"`
	JaegerTracesUrl         string      `json:"jaegerTracesUrl"`
	OtlpTracesUrl           string      `json:"otlpTracesUrl"`
	ZipkinSpansUrl          string      `json:"zipkinSpansUrl"`
	JaegerUiUrl             string      `json:"jaegerUiUrl"`
}

type Plan struct {
	PlanId                  string  `json:"planId"`
	Id                      string  `json:"id"`
	Description             string  `json:"description"`
	Name                    string  `json:"name"`
	BucketSize              int     `json:"bucketSize"`
	GrafanaGlobalUsers      int     `json:"grafanaGlobalUsers"`
	GrafanaGlobalOrgs       int     `json:"grafanaGlobalOrgs"`
	GrafanaGlobalDashboards int     `json:"grafanaGlobalDashboards"`
	AlertRules              int     `json:"alertRules"`
	TargetNumber            int     `json:"target_number"`
	SamplesPerScrape        int     `json:"samples_per_scrape"`
	GrafanaGlobalSessions   int     `json:"grafana_global_sessions"`
	Amount                  float32 `json:"amount"`
	AlertReceivers          int     `json:"alert_receivers"`
	AlertMatchers           int     `json:"alert_matchers"`
	TracesStorage           int     `json:"traces_storage"`
	LogsStorage             int     `json:"logs_storage"`
	LogsAlert               int     `json:"logs_alert"`
	IsFree                  bool    `json:"is_free"`
	IsPublic                bool    `json:"is_public"`
	Schema                  string  `json:"schema"`
}
