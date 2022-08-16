package main

import (
	"go/types"
)

type ApiStandardPayload struct {
	ProjectId     string
	Authorization string
}

type CreateInstance struct {
	IdAndHeader ApiStandardPayload
	Body        CreateInstanceBody
}

type CreateInstanceBody struct {
	Name      string    `json:"name"`
	PlanId    string    `json:"plan_id"`
	Parameter types.Nil `json:"parameter"`
}

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
	Message      string    `json:"message"`
	DashboardUrl string    `json:"dashboardUrl"`
	IsUpdatable  bool      `json:"isUpdatable"`
	Name         string    `json:"name"`
	Parameters   types.Nil `json:"parameters"`
	Id           string    `json:"id"`
	ServiceName  string    `json:"serviceName"`
	PlanId       string    `json:"planId"`
	PlanName     string    `json:"planName"`
	PlanSchema   types.Nil `json:"planSchema"`
	Status       string    `json:"status"`
	Instance     Instance  `json:"instance"`
}

type Instance struct {
	Instance                string      `json:"instance"`
	Cluster                 string      `json:"cluster"`
	GrafanaUrl              string      `json:"grafanaUrl"`
	DashboardUrl            string      `json:"dashboardUrl"`
	GrafanaPlugins          types.Array `json:"grafanaPlugins"`
	Name                    string      `json:"name"`
	GrafanaAdminPassword    string      `json:"grafanaAdminPassword"`
	GrafanaAdminUser        string      `json:"grafanaAdminUser"`
	MetricsRetentionTimeRaw string      `json:"metricsRetentionTimeRaw"`
	MetricsRetentionTime5m  string      `json:"metricsRetentionTime5m"`
	MetricsRetentionTime1h  string      `json:"metricsRetentionTime1h"`
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
	PlanId                  string `json:"planId"`
	Id                      string `json:"id"`
	Description             string `json:"description"`
	Name                    string `json:"name"`
	BucketSize              int    `json:"bucketSize"`
	GrafanaGlobalUsers      int    `json:"grafanaGlobalUsers"`
	GrafanaGlobalOrgs       int    `json:"grafanaGlobalOrgs"`
	GrafanaGlobalDashboards int    `json:"grafanaGlobalDashboards"`
	AlertRules              int    `json:"alertRules"`
	TargetNumber            int    `json:"target_number"`
	SamplesPerScrape        int    `json:"samples_per_scrape"`
	GrafanaGlobalSessions   int    `json:"grafana_global_sessions"`
	Amount                  int    `json:"amount"`
	AlertReceivers          int    `json:"alert_receivers"`
	AlertMatchers           int    `json:"alert_matchers"`
	TracesStorage           int    `json:"traces_storage"`
	LogsStorage             int    `json:"logs_storage"`
	LogsAlert               int    `json:"logs_alert"`
	IsFree                  bool   `json:"is_free"`
	IsPublic                bool   `json:"is_public"`
	Schema                  string `json:"schema"`
}

type InstanceCreateResponse struct {
	Message      string `json:"message"`
	InstanceId   string `json:"instance_id"`
	DashboardUrl string `json:"dashboard_url"`
}
