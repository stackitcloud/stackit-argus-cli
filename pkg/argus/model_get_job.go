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

// GetJob struct for GetJob
type GetJob struct {
	Message string `json:"message"`
	Data    Job    `json:"data"`
}

// NewGetJob instantiates a new GetJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetJob(message string, data Job) *GetJob {
	this := GetJob{}
	this.Message = message
	this.Data = data
	return &this
}

// NewGetJobWithDefaults instantiates a new GetJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetJobWithDefaults() *GetJob {
	this := GetJob{}
	return &this
}

// GetMessage returns the Message field value
func (o *GetJob) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *GetJob) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *GetJob) SetMessage(v string) {
	o.Message = v
}

// GetData returns the Data field value
func (o *GetJob) GetData() Job {
	if o == nil {
		var ret Job
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetJob) GetDataOk() (*Job, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetJob) SetData(v Job) {
	o.Data = v
}

func (o GetJob) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetJob struct {
	value *GetJob
	isSet bool
}

func (v NullableGetJob) Get() *GetJob {
	return v.value
}

func (v *NullableGetJob) Set(val *GetJob) {
	v.value = val
	v.isSet = true
}

func (v NullableGetJob) IsSet() bool {
	return v.isSet
}

func (v *NullableGetJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetJob(val *GetJob) *NullableGetJob {
	return &NullableGetJob{value: val, isSet: true}
}

func (v NullableGetJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
