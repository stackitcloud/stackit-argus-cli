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

// V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner
type V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner struct {
	// The API key to use when talking to the OpsGenie API.
	ApiKey *string `json:"apiKey,omitempty"`
	// The host to send OpsGenie API requests to.
	ApiUrl *string `json:"apiUrl,omitempty"`
	// Comma separated list of tags attached to the notifications.
	Tags       *string                                                                           `json:"tags,omitempty"`
	HttpConfig *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig `json:"httpConfig,omitempty"`
}

// NewV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner instantiates a new V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner() *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner {
	this := V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner{}
	return &this
}

// NewV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerWithDefaults instantiates a new V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerWithDefaults() *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner {
	this := V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner) GetApiKey() string {
	if o == nil || o.ApiKey == nil {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner) GetApiKeyOk() (*string, bool) {
	if o == nil || o.ApiKey == nil {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner) HasApiKey() bool {
	if o != nil && o.ApiKey != nil {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetApiUrl returns the ApiUrl field value if set, zero value otherwise.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner) GetApiUrl() string {
	if o == nil || o.ApiUrl == nil {
		var ret string
		return ret
	}
	return *o.ApiUrl
}

// GetApiUrlOk returns a tuple with the ApiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner) GetApiUrlOk() (*string, bool) {
	if o == nil || o.ApiUrl == nil {
		return nil, false
	}
	return o.ApiUrl, true
}

// HasApiUrl returns a boolean if a field has been set.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner) HasApiUrl() bool {
	if o != nil && o.ApiUrl != nil {
		return true
	}

	return false
}

// SetApiUrl gets a reference to the given string and assigns it to the ApiUrl field.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner) SetApiUrl(v string) {
	o.ApiUrl = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner) GetTags() string {
	if o == nil || o.Tags == nil {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner) GetTagsOk() (*string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner) SetTags(v string) {
	o.Tags = &v
}

// GetHttpConfig returns the HttpConfig field value if set, zero value otherwise.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner) GetHttpConfig() V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig {
	if o == nil || o.HttpConfig == nil {
		var ret V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig
		return ret
	}
	return *o.HttpConfig
}

// GetHttpConfigOk returns a tuple with the HttpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner) GetHttpConfigOk() (*V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig, bool) {
	if o == nil || o.HttpConfig == nil {
		return nil, false
	}
	return o.HttpConfig, true
}

// HasHttpConfig returns a boolean if a field has been set.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner) HasHttpConfig() bool {
	if o != nil && o.HttpConfig != nil {
		return true
	}

	return false
}

// SetHttpConfig gets a reference to the given V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig and assigns it to the HttpConfig field.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner) SetHttpConfig(v V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInnerHttpConfig) {
	o.HttpConfig = &v
}

func (o V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiKey != nil {
		toSerialize["apiKey"] = o.ApiKey
	}
	if o.ApiUrl != nil {
		toSerialize["apiUrl"] = o.ApiUrl
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.HttpConfig != nil {
		toSerialize["httpConfig"] = o.HttpConfig
	}
	return json.Marshal(toSerialize)
}

type NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner struct {
	value *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner
	isSet bool
}

func (v NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner) Get() *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner {
	return v.value
}

func (v *NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner) Set(val *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner(val *V1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner) *NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner {
	return &NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner{value: val, isSet: true}
}

func (v NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerOpsgenieConfigsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
