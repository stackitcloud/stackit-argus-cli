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

// Instance struct for Instance
type Instance struct {
	Cluster                 string    `json:"cluster"`
	Instance                string    `json:"instance"`
	Name                    *string   `json:"name,omitempty"`
	Plan                    PlanModel `json:"plan"`
	BucketRetentionTimeRaw  int32     `json:"bucketRetentionTimeRaw"`
	BucketRetentionTime5m   int32     `json:"bucketRetentionTime5m"`
	BucketRetentionTime1h   int32     `json:"bucketRetentionTime1h"`
	State                   *string   `json:"state,omitempty"`
	GrafanaPublicReadAccess bool      `json:"grafanaPublicReadAccess"`
}

// NewInstance instantiates a new Instance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstance(cluster string, instance string, plan PlanModel, bucketRetentionTimeRaw int32, bucketRetentionTime5m int32, bucketRetentionTime1h int32, grafanaPublicReadAccess bool) *Instance {
	this := Instance{}
	this.Cluster = cluster
	this.Instance = instance
	this.Plan = plan
	this.BucketRetentionTimeRaw = bucketRetentionTimeRaw
	this.BucketRetentionTime5m = bucketRetentionTime5m
	this.BucketRetentionTime1h = bucketRetentionTime1h
	this.GrafanaPublicReadAccess = grafanaPublicReadAccess
	return &this
}

// NewInstanceWithDefaults instantiates a new Instance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceWithDefaults() *Instance {
	this := Instance{}
	return &this
}

// GetCluster returns the Cluster field value
func (o *Instance) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *Instance) GetClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value
func (o *Instance) SetCluster(v string) {
	o.Cluster = v
}

// GetInstance returns the Instance field value
func (o *Instance) GetInstance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value
// and a boolean to check if the value has been set.
func (o *Instance) GetInstanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Instance, true
}

// SetInstance sets field value
func (o *Instance) SetInstance(v string) {
	o.Instance = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Instance) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Instance) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Instance) SetName(v string) {
	o.Name = &v
}

// GetPlan returns the Plan field value
func (o *Instance) GetPlan() PlanModel {
	if o == nil {
		var ret PlanModel
		return ret
	}

	return o.Plan
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
func (o *Instance) GetPlanOk() (*PlanModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plan, true
}

// SetPlan sets field value
func (o *Instance) SetPlan(v PlanModel) {
	o.Plan = v
}

// GetBucketRetentionTimeRaw returns the BucketRetentionTimeRaw field value
func (o *Instance) GetBucketRetentionTimeRaw() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BucketRetentionTimeRaw
}

// GetBucketRetentionTimeRawOk returns a tuple with the BucketRetentionTimeRaw field value
// and a boolean to check if the value has been set.
func (o *Instance) GetBucketRetentionTimeRawOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketRetentionTimeRaw, true
}

// SetBucketRetentionTimeRaw sets field value
func (o *Instance) SetBucketRetentionTimeRaw(v int32) {
	o.BucketRetentionTimeRaw = v
}

// GetBucketRetentionTime5m returns the BucketRetentionTime5m field value
func (o *Instance) GetBucketRetentionTime5m() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BucketRetentionTime5m
}

// GetBucketRetentionTime5mOk returns a tuple with the BucketRetentionTime5m field value
// and a boolean to check if the value has been set.
func (o *Instance) GetBucketRetentionTime5mOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketRetentionTime5m, true
}

// SetBucketRetentionTime5m sets field value
func (o *Instance) SetBucketRetentionTime5m(v int32) {
	o.BucketRetentionTime5m = v
}

// GetBucketRetentionTime1h returns the BucketRetentionTime1h field value
func (o *Instance) GetBucketRetentionTime1h() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BucketRetentionTime1h
}

// GetBucketRetentionTime1hOk returns a tuple with the BucketRetentionTime1h field value
// and a boolean to check if the value has been set.
func (o *Instance) GetBucketRetentionTime1hOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketRetentionTime1h, true
}

// SetBucketRetentionTime1h sets field value
func (o *Instance) SetBucketRetentionTime1h(v int32) {
	o.BucketRetentionTime1h = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Instance) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Instance) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Instance) SetState(v string) {
	o.State = &v
}

// GetGrafanaPublicReadAccess returns the GrafanaPublicReadAccess field value
func (o *Instance) GetGrafanaPublicReadAccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.GrafanaPublicReadAccess
}

// GetGrafanaPublicReadAccessOk returns a tuple with the GrafanaPublicReadAccess field value
// and a boolean to check if the value has been set.
func (o *Instance) GetGrafanaPublicReadAccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrafanaPublicReadAccess, true
}

// SetGrafanaPublicReadAccess sets field value
func (o *Instance) SetGrafanaPublicReadAccess(v bool) {
	o.GrafanaPublicReadAccess = v
}

func (o Instance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cluster"] = o.Cluster
	}
	if true {
		toSerialize["instance"] = o.Instance
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["plan"] = o.Plan
	}
	if true {
		toSerialize["bucketRetentionTimeRaw"] = o.BucketRetentionTimeRaw
	}
	if true {
		toSerialize["bucketRetentionTime5m"] = o.BucketRetentionTime5m
	}
	if true {
		toSerialize["bucketRetentionTime1h"] = o.BucketRetentionTime1h
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["grafanaPublicReadAccess"] = o.GrafanaPublicReadAccess
	}
	return json.Marshal(toSerialize)
}

type NullableInstance struct {
	value *Instance
	isSet bool
}

func (v NullableInstance) Get() *Instance {
	return v.value
}

func (v *NullableInstance) Set(val *Instance) {
	v.value = val
	v.isSet = true
}

func (v NullableInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstance(val *Instance) *NullableInstance {
	return &NullableInstance{value: val, isSet: true}
}

func (v NullableInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}