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

// Services struct for Services
type Services struct {
	Message string `json:"message"`
	Instances []Instance `json:"instances"`
}

// NewServices instantiates a new Services object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServices(message string, instances []Instance) *Services {
	this := Services{}
	this.Message = message
	this.Instances = instances
	return &this
}

// NewServicesWithDefaults instantiates a new Services object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesWithDefaults() *Services {
	this := Services{}
	return &this
}

// GetMessage returns the Message field value
func (o *Services) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *Services) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *Services) SetMessage(v string) {
	o.Message = v
}

// GetInstances returns the Instances field value
func (o *Services) GetInstances() []Instance {
	if o == nil {
		var ret []Instance
		return ret
	}

	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value
// and a boolean to check if the value has been set.
func (o *Services) GetInstancesOk() ([]Instance, bool) {
	if o == nil {
		return nil, false
	}
	return o.Instances, true
}

// SetInstances sets field value
func (o *Services) SetInstances(v []Instance) {
	o.Instances = v
}

func (o Services) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["instances"] = o.Instances
	}
	return json.Marshal(toSerialize)
}

type NullableServices struct {
	value *Services
	isSet bool
}

func (v NullableServices) Get() *Services {
	return v.value
}

func (v *NullableServices) Set(val *Services) {
	v.value = val
	v.isSet = true
}

func (v NullableServices) IsSet() bool {
	return v.isSet
}

func (v *NullableServices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServices(val *Services) *NullableServices {
	return &NullableServices{value: val, isSet: true}
}

func (v NullableServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


