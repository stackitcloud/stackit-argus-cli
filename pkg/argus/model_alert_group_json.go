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

// AlertGroupJson struct for AlertGroupJson
type AlertGroupJson struct {
	Name string `json:"name"`
	Interval *string `json:"interval,omitempty"`
	Rules []AlertRuleRecordJson `json:"rules"`
}

// NewAlertGroupJson instantiates a new AlertGroupJson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertGroupJson(name string, rules []AlertRuleRecordJson) *AlertGroupJson {
	this := AlertGroupJson{}
	this.Name = name
	var interval string = "60s"
	this.Interval = &interval
	this.Rules = rules
	return &this
}

// NewAlertGroupJsonWithDefaults instantiates a new AlertGroupJson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertGroupJsonWithDefaults() *AlertGroupJson {
	this := AlertGroupJson{}
	var interval string = "60s"
	this.Interval = &interval
	return &this
}

// GetName returns the Name field value
func (o *AlertGroupJson) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AlertGroupJson) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AlertGroupJson) SetName(v string) {
	o.Name = v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *AlertGroupJson) GetInterval() string {
	if o == nil || o.Interval == nil {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertGroupJson) GetIntervalOk() (*string, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *AlertGroupJson) HasInterval() bool {
	if o != nil && o.Interval != nil {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *AlertGroupJson) SetInterval(v string) {
	o.Interval = &v
}

// GetRules returns the Rules field value
func (o *AlertGroupJson) GetRules() []AlertRuleRecordJson {
	if o == nil {
		var ret []AlertRuleRecordJson
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *AlertGroupJson) GetRulesOk() ([]AlertRuleRecordJson, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *AlertGroupJson) SetRules(v []AlertRuleRecordJson) {
	o.Rules = v
}

func (o AlertGroupJson) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	if true {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableAlertGroupJson struct {
	value *AlertGroupJson
	isSet bool
}

func (v NullableAlertGroupJson) Get() *AlertGroupJson {
	return v.value
}

func (v *NullableAlertGroupJson) Set(val *AlertGroupJson) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertGroupJson) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertGroupJson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertGroupJson(val *AlertGroupJson) *NullableAlertGroupJson {
	return &NullableAlertGroupJson{value: val, isSet: true}
}

func (v NullableAlertGroupJson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertGroupJson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


