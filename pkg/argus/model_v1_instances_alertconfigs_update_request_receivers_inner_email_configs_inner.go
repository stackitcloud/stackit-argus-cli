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

// V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner
type V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner struct {
	// The email address to send notifications to.
	To *string `json:"to,omitempty"`
	// The sender address.
	From *string `json:"from,omitempty"`
	// The SMTP host through which emails are sent.
	Smarthost *string `json:"smarthost,omitempty"`
	// SMTP authentication information.
	AuthUsername *string `json:"authUsername,omitempty"`
	// SMTP authentication information.
	AuthPassword *string `json:"authPassword,omitempty"`
	// SMTP authentication information.
	AuthIdentity *string `json:"authIdentity,omitempty"`
}

// NewV1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner instantiates a new V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner() *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner {
	this := V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner{}
	return &this
}

// NewV1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInnerWithDefaults instantiates a new V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInnerWithDefaults() *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner {
	this := V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner{}
	return &this
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) SetTo(v string) {
	o.To = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) GetFrom() string {
	if o == nil || o.From == nil {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) GetFromOk() (*string, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) SetFrom(v string) {
	o.From = &v
}

// GetSmarthost returns the Smarthost field value if set, zero value otherwise.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) GetSmarthost() string {
	if o == nil || o.Smarthost == nil {
		var ret string
		return ret
	}
	return *o.Smarthost
}

// GetSmarthostOk returns a tuple with the Smarthost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) GetSmarthostOk() (*string, bool) {
	if o == nil || o.Smarthost == nil {
		return nil, false
	}
	return o.Smarthost, true
}

// HasSmarthost returns a boolean if a field has been set.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) HasSmarthost() bool {
	if o != nil && o.Smarthost != nil {
		return true
	}

	return false
}

// SetSmarthost gets a reference to the given string and assigns it to the Smarthost field.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) SetSmarthost(v string) {
	o.Smarthost = &v
}

// GetAuthUsername returns the AuthUsername field value if set, zero value otherwise.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) GetAuthUsername() string {
	if o == nil || o.AuthUsername == nil {
		var ret string
		return ret
	}
	return *o.AuthUsername
}

// GetAuthUsernameOk returns a tuple with the AuthUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) GetAuthUsernameOk() (*string, bool) {
	if o == nil || o.AuthUsername == nil {
		return nil, false
	}
	return o.AuthUsername, true
}

// HasAuthUsername returns a boolean if a field has been set.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) HasAuthUsername() bool {
	if o != nil && o.AuthUsername != nil {
		return true
	}

	return false
}

// SetAuthUsername gets a reference to the given string and assigns it to the AuthUsername field.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) SetAuthUsername(v string) {
	o.AuthUsername = &v
}

// GetAuthPassword returns the AuthPassword field value if set, zero value otherwise.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) GetAuthPassword() string {
	if o == nil || o.AuthPassword == nil {
		var ret string
		return ret
	}
	return *o.AuthPassword
}

// GetAuthPasswordOk returns a tuple with the AuthPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) GetAuthPasswordOk() (*string, bool) {
	if o == nil || o.AuthPassword == nil {
		return nil, false
	}
	return o.AuthPassword, true
}

// HasAuthPassword returns a boolean if a field has been set.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) HasAuthPassword() bool {
	if o != nil && o.AuthPassword != nil {
		return true
	}

	return false
}

// SetAuthPassword gets a reference to the given string and assigns it to the AuthPassword field.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) SetAuthPassword(v string) {
	o.AuthPassword = &v
}

// GetAuthIdentity returns the AuthIdentity field value if set, zero value otherwise.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) GetAuthIdentity() string {
	if o == nil || o.AuthIdentity == nil {
		var ret string
		return ret
	}
	return *o.AuthIdentity
}

// GetAuthIdentityOk returns a tuple with the AuthIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) GetAuthIdentityOk() (*string, bool) {
	if o == nil || o.AuthIdentity == nil {
		return nil, false
	}
	return o.AuthIdentity, true
}

// HasAuthIdentity returns a boolean if a field has been set.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) HasAuthIdentity() bool {
	if o != nil && o.AuthIdentity != nil {
		return true
	}

	return false
}

// SetAuthIdentity gets a reference to the given string and assigns it to the AuthIdentity field.
func (o *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) SetAuthIdentity(v string) {
	o.AuthIdentity = &v
}

func (o V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.Smarthost != nil {
		toSerialize["smarthost"] = o.Smarthost
	}
	if o.AuthUsername != nil {
		toSerialize["authUsername"] = o.AuthUsername
	}
	if o.AuthPassword != nil {
		toSerialize["authPassword"] = o.AuthPassword
	}
	if o.AuthIdentity != nil {
		toSerialize["authIdentity"] = o.AuthIdentity
	}
	return json.Marshal(toSerialize)
}

type NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner struct {
	value *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner
	isSet bool
}

func (v NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) Get() *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner {
	return v.value
}

func (v *NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) Set(val *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner(val *V1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) *NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner {
	return &NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner{value: val, isSet: true}
}

func (v NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1InstancesAlertconfigsUpdateRequestReceiversInnerEmailConfigsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}