/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.0
Contact: patrick.koss@mail.schwarz
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

import (
	"encoding/json"
)

// Dashboard struct for Dashboard
type Dashboard struct {
	Message      string `json:"message"`
	DashboardUrl string `json:"dashboardUrl"`
}

// NewDashboard instantiates a new Dashboard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboard(message string, dashboardUrl string) *Dashboard {
	this := Dashboard{}
	this.Message = message
	this.DashboardUrl = dashboardUrl
	return &this
}

// NewDashboardWithDefaults instantiates a new Dashboard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardWithDefaults() *Dashboard {
	this := Dashboard{}
	return &this
}

// GetMessage returns the Message field value
func (o *Dashboard) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *Dashboard) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *Dashboard) SetMessage(v string) {
	o.Message = v
}

// GetDashboardUrl returns the DashboardUrl field value
func (o *Dashboard) GetDashboardUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DashboardUrl
}

// GetDashboardUrlOk returns a tuple with the DashboardUrl field value
// and a boolean to check if the value has been set.
func (o *Dashboard) GetDashboardUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DashboardUrl, true
}

// SetDashboardUrl sets field value
func (o *Dashboard) SetDashboardUrl(v string) {
	o.DashboardUrl = v
}

func (o Dashboard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["dashboardUrl"] = o.DashboardUrl
	}
	return json.Marshal(toSerialize)
}

type NullableDashboard struct {
	value *Dashboard
	isSet bool
}

func (v NullableDashboard) Get() *Dashboard {
	return v.value
}

func (v *NullableDashboard) Set(val *Dashboard) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboard) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboard(val *Dashboard) *NullableDashboard {
	return &NullableDashboard{value: val, isSet: true}
}

func (v NullableDashboard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}