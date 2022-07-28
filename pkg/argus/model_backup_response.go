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

// BackupResponse struct for BackupResponse
type BackupResponse struct {
	Message string `json:"message"`
	AlertConfigBackups []string `json:"alertConfigBackups"`
	AlertRulesBackups []string `json:"alertRulesBackups"`
	ScrapeConfigBackups []string `json:"scrapeConfigBackups"`
	GrafanaBackups []string `json:"grafanaBackups"`
}

// NewBackupResponse instantiates a new BackupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupResponse(message string, alertConfigBackups []string, alertRulesBackups []string, scrapeConfigBackups []string, grafanaBackups []string) *BackupResponse {
	this := BackupResponse{}
	this.Message = message
	this.AlertConfigBackups = alertConfigBackups
	this.AlertRulesBackups = alertRulesBackups
	this.ScrapeConfigBackups = scrapeConfigBackups
	this.GrafanaBackups = grafanaBackups
	return &this
}

// NewBackupResponseWithDefaults instantiates a new BackupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupResponseWithDefaults() *BackupResponse {
	this := BackupResponse{}
	return &this
}

// GetMessage returns the Message field value
func (o *BackupResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *BackupResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *BackupResponse) SetMessage(v string) {
	o.Message = v
}

// GetAlertConfigBackups returns the AlertConfigBackups field value
func (o *BackupResponse) GetAlertConfigBackups() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AlertConfigBackups
}

// GetAlertConfigBackupsOk returns a tuple with the AlertConfigBackups field value
// and a boolean to check if the value has been set.
func (o *BackupResponse) GetAlertConfigBackupsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlertConfigBackups, true
}

// SetAlertConfigBackups sets field value
func (o *BackupResponse) SetAlertConfigBackups(v []string) {
	o.AlertConfigBackups = v
}

// GetAlertRulesBackups returns the AlertRulesBackups field value
func (o *BackupResponse) GetAlertRulesBackups() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AlertRulesBackups
}

// GetAlertRulesBackupsOk returns a tuple with the AlertRulesBackups field value
// and a boolean to check if the value has been set.
func (o *BackupResponse) GetAlertRulesBackupsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlertRulesBackups, true
}

// SetAlertRulesBackups sets field value
func (o *BackupResponse) SetAlertRulesBackups(v []string) {
	o.AlertRulesBackups = v
}

// GetScrapeConfigBackups returns the ScrapeConfigBackups field value
func (o *BackupResponse) GetScrapeConfigBackups() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ScrapeConfigBackups
}

// GetScrapeConfigBackupsOk returns a tuple with the ScrapeConfigBackups field value
// and a boolean to check if the value has been set.
func (o *BackupResponse) GetScrapeConfigBackupsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScrapeConfigBackups, true
}

// SetScrapeConfigBackups sets field value
func (o *BackupResponse) SetScrapeConfigBackups(v []string) {
	o.ScrapeConfigBackups = v
}

// GetGrafanaBackups returns the GrafanaBackups field value
func (o *BackupResponse) GetGrafanaBackups() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.GrafanaBackups
}

// GetGrafanaBackupsOk returns a tuple with the GrafanaBackups field value
// and a boolean to check if the value has been set.
func (o *BackupResponse) GetGrafanaBackupsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GrafanaBackups, true
}

// SetGrafanaBackups sets field value
func (o *BackupResponse) SetGrafanaBackups(v []string) {
	o.GrafanaBackups = v
}

func (o BackupResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["alertConfigBackups"] = o.AlertConfigBackups
	}
	if true {
		toSerialize["alertRulesBackups"] = o.AlertRulesBackups
	}
	if true {
		toSerialize["scrapeConfigBackups"] = o.ScrapeConfigBackups
	}
	if true {
		toSerialize["grafanaBackups"] = o.GrafanaBackups
	}
	return json.Marshal(toSerialize)
}

type NullableBackupResponse struct {
	value *BackupResponse
	isSet bool
}

func (v NullableBackupResponse) Get() *BackupResponse {
	return v.value
}

func (v *NullableBackupResponse) Set(val *BackupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupResponse(val *BackupResponse) *NullableBackupResponse {
	return &NullableBackupResponse{value: val, isSet: true}
}

func (v NullableBackupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


