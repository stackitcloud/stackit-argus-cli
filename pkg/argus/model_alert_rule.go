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

// AlertRule struct for AlertRule
type AlertRule struct {
	Alert string `json:"alert"`
	Expr string `json:"expr"`
	Labels *map[string]string `json:"labels,omitempty"`
	Annotations *map[string]string `json:"annotations,omitempty"`
	For *string `json:"for,omitempty"`
}

// NewAlertRule instantiates a new AlertRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertRule(alert string, expr string) *AlertRule {
	this := AlertRule{}
	this.Alert = alert
	this.Expr = expr
	var for_ string = "0s"
	this.For = &for_
	return &this
}

// NewAlertRuleWithDefaults instantiates a new AlertRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertRuleWithDefaults() *AlertRule {
	this := AlertRule{}
	var for_ string = "0s"
	this.For = &for_
	return &this
}

// GetAlert returns the Alert field value
func (o *AlertRule) GetAlert() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alert
}

// GetAlertOk returns a tuple with the Alert field value
// and a boolean to check if the value has been set.
func (o *AlertRule) GetAlertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alert, true
}

// SetAlert sets field value
func (o *AlertRule) SetAlert(v string) {
	o.Alert = v
}

// GetExpr returns the Expr field value
func (o *AlertRule) GetExpr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expr
}

// GetExprOk returns a tuple with the Expr field value
// and a boolean to check if the value has been set.
func (o *AlertRule) GetExprOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expr, true
}

// SetExpr sets field value
func (o *AlertRule) SetExpr(v string) {
	o.Expr = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *AlertRule) GetLabels() map[string]string {
	if o == nil || o.Labels == nil {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRule) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *AlertRule) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *AlertRule) SetLabels(v map[string]string) {
	o.Labels = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *AlertRule) GetAnnotations() map[string]string {
	if o == nil || o.Annotations == nil {
		var ret map[string]string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRule) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *AlertRule) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *AlertRule) SetAnnotations(v map[string]string) {
	o.Annotations = &v
}

// GetFor returns the For field value if set, zero value otherwise.
func (o *AlertRule) GetFor() string {
	if o == nil || o.For == nil {
		var ret string
		return ret
	}
	return *o.For
}

// GetForOk returns a tuple with the For field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRule) GetForOk() (*string, bool) {
	if o == nil || o.For == nil {
		return nil, false
	}
	return o.For, true
}

// HasFor returns a boolean if a field has been set.
func (o *AlertRule) HasFor() bool {
	if o != nil && o.For != nil {
		return true
	}

	return false
}

// SetFor gets a reference to the given string and assigns it to the For field.
func (o *AlertRule) SetFor(v string) {
	o.For = &v
}

func (o AlertRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["alert"] = o.Alert
	}
	if true {
		toSerialize["expr"] = o.Expr
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	if o.For != nil {
		toSerialize["for"] = o.For
	}
	return json.Marshal(toSerialize)
}

type NullableAlertRule struct {
	value *AlertRule
	isSet bool
}

func (v NullableAlertRule) Get() *AlertRule {
	return v.value
}

func (v *NullableAlertRule) Set(val *AlertRule) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertRule) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertRule(val *AlertRule) *NullableAlertRule {
	return &NullableAlertRule{value: val, isSet: true}
}

func (v NullableAlertRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


