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

// AlertGroup struct for AlertGroup
type AlertGroup struct {
	Name     string      `json:"name"`
	Interval string      `json:"interval"`
	Rules    []AlertRule `json:"rules"`
}

// NewAlertGroup instantiates a new AlertGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertGroup(name string, interval string, rules []AlertRule) *AlertGroup {
	this := AlertGroup{}
	this.Name = name
	this.Interval = interval
	this.Rules = rules
	return &this
}

// NewAlertGroupWithDefaults instantiates a new AlertGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertGroupWithDefaults() *AlertGroup {
	this := AlertGroup{}
	return &this
}

// GetName returns the Name field value
func (o *AlertGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AlertGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AlertGroup) SetName(v string) {
	o.Name = v
}

// GetInterval returns the Interval field value
func (o *AlertGroup) GetInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *AlertGroup) GetIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *AlertGroup) SetInterval(v string) {
	o.Interval = v
}

// GetRules returns the Rules field value
func (o *AlertGroup) GetRules() []AlertRule {
	if o == nil {
		var ret []AlertRule
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *AlertGroup) GetRulesOk() ([]AlertRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *AlertGroup) SetRules(v []AlertRule) {
	o.Rules = v
}

func (o AlertGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["interval"] = o.Interval
	}
	if true {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableAlertGroup struct {
	value *AlertGroup
	isSet bool
}

func (v NullableAlertGroup) Get() *AlertGroup {
	return v.value
}

func (v *NullableAlertGroup) Set(val *AlertGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertGroup(val *AlertGroup) *NullableAlertGroup {
	return &NullableAlertGroup{value: val, isSet: true}
}

func (v NullableAlertGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
