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

// V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig The HTTP client's configuration.
type V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig struct {
	//
	ProxyUrl *string `json:"proxyUrl,omitempty"`
}

// NewV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig instantiates a new V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig() *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig {
	this := V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig{}
	return &this
}

// NewV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfigWithDefaults instantiates a new V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfigWithDefaults() *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig {
	this := V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig{}
	return &this
}

// GetProxyUrl returns the ProxyUrl field value if set, zero value otherwise.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig) GetProxyUrl() string {
	if o == nil || o.ProxyUrl == nil {
		var ret string
		return ret
	}
	return *o.ProxyUrl
}

// GetProxyUrlOk returns a tuple with the ProxyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig) GetProxyUrlOk() (*string, bool) {
	if o == nil || o.ProxyUrl == nil {
		return nil, false
	}
	return o.ProxyUrl, true
}

// HasProxyUrl returns a boolean if a field has been set.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig) HasProxyUrl() bool {
	if o != nil && o.ProxyUrl != nil {
		return true
	}

	return false
}

// SetProxyUrl gets a reference to the given string and assigns it to the ProxyUrl field.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig) SetProxyUrl(v string) {
	o.ProxyUrl = &v
}

func (o V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProxyUrl != nil {
		toSerialize["proxyUrl"] = o.ProxyUrl
	}
	return json.Marshal(toSerialize)
}

type NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig struct {
	value *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig
	isSet bool
}

func (v NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig) Get() *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig {
	return v.value
}

func (v *NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig) Set(val *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig(val *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig) *NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig {
	return &NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig{value: val, isSet: true}
}

func (v NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
