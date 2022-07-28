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

// TraceConfig struct for TraceConfig
type TraceConfig struct {
	Retention string `json:"retention"`
}

// NewTraceConfig instantiates a new TraceConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraceConfig(retention string) *TraceConfig {
	this := TraceConfig{}
	this.Retention = retention
	return &this
}

// NewTraceConfigWithDefaults instantiates a new TraceConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceConfigWithDefaults() *TraceConfig {
	this := TraceConfig{}
	return &this
}

// GetRetention returns the Retention field value
func (o *TraceConfig) GetRetention() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value
// and a boolean to check if the value has been set.
func (o *TraceConfig) GetRetentionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Retention, true
}

// SetRetention sets field value
func (o *TraceConfig) SetRetention(v string) {
	o.Retention = v
}

func (o TraceConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["retention"] = o.Retention
	}
	return json.Marshal(toSerialize)
}

type NullableTraceConfig struct {
	value *TraceConfig
	isSet bool
}

func (v NullableTraceConfig) Get() *TraceConfig {
	return v.value
}

func (v *NullableTraceConfig) Set(val *TraceConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceConfig(val *TraceConfig) *NullableTraceConfig {
	return &NullableTraceConfig{value: val, isSet: true}
}

func (v NullableTraceConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


