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

// PutAlertRule struct for PutAlertRule
type PutAlertRule struct {
	Message string      `json:"message"`
	Data    []AlertRule `json:"data"`
}

// NewPutAlertRule instantiates a new PutAlertRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutAlertRule(message string, data []AlertRule) *PutAlertRule {
	this := PutAlertRule{}
	this.Message = message
	this.Data = data
	return &this
}

// NewPutAlertRuleWithDefaults instantiates a new PutAlertRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutAlertRuleWithDefaults() *PutAlertRule {
	this := PutAlertRule{}
	return &this
}

// GetMessage returns the Message field value
func (o *PutAlertRule) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *PutAlertRule) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *PutAlertRule) SetMessage(v string) {
	o.Message = v
}

// GetData returns the Data field value
func (o *PutAlertRule) GetData() []AlertRule {
	if o == nil {
		var ret []AlertRule
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PutAlertRule) GetDataOk() ([]AlertRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PutAlertRule) SetData(v []AlertRule) {
	o.Data = v
}

func (o PutAlertRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePutAlertRule struct {
	value *PutAlertRule
	isSet bool
}

func (v NullablePutAlertRule) Get() *PutAlertRule {
	return v.value
}

func (v *NullablePutAlertRule) Set(val *PutAlertRule) {
	v.value = val
	v.isSet = true
}

func (v NullablePutAlertRule) IsSet() bool {
	return v.isSet
}

func (v *NullablePutAlertRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutAlertRule(val *PutAlertRule) *NullablePutAlertRule {
	return &NullablePutAlertRule{value: val, isSet: true}
}

func (v NullablePutAlertRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutAlertRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}