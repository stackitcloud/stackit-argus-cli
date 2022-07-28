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

// EmailConfig struct for EmailConfig
type EmailConfig struct {
	To           string  `json:"to"`
	FromPerson   *string `json:"fromPerson,omitempty"`
	Smarthost    *string `json:"smarthost,omitempty"`
	AuthUsername *string `json:"authUsername,omitempty"`
	AuthPassword *string `json:"authPassword,omitempty"`
	AuthIdentity *string `json:"authIdentity,omitempty"`
}

// NewEmailConfig instantiates a new EmailConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailConfig(to string) *EmailConfig {
	this := EmailConfig{}
	this.To = to
	return &this
}

// NewEmailConfigWithDefaults instantiates a new EmailConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailConfigWithDefaults() *EmailConfig {
	this := EmailConfig{}
	return &this
}

// GetTo returns the To field value
func (o *EmailConfig) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *EmailConfig) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *EmailConfig) SetTo(v string) {
	o.To = v
}

// GetFromPerson returns the FromPerson field value if set, zero value otherwise.
func (o *EmailConfig) GetFromPerson() string {
	if o == nil || o.FromPerson == nil {
		var ret string
		return ret
	}
	return *o.FromPerson
}

// GetFromPersonOk returns a tuple with the FromPerson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConfig) GetFromPersonOk() (*string, bool) {
	if o == nil || o.FromPerson == nil {
		return nil, false
	}
	return o.FromPerson, true
}

// HasFromPerson returns a boolean if a field has been set.
func (o *EmailConfig) HasFromPerson() bool {
	if o != nil && o.FromPerson != nil {
		return true
	}

	return false
}

// SetFromPerson gets a reference to the given string and assigns it to the FromPerson field.
func (o *EmailConfig) SetFromPerson(v string) {
	o.FromPerson = &v
}

// GetSmarthost returns the Smarthost field value if set, zero value otherwise.
func (o *EmailConfig) GetSmarthost() string {
	if o == nil || o.Smarthost == nil {
		var ret string
		return ret
	}
	return *o.Smarthost
}

// GetSmarthostOk returns a tuple with the Smarthost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConfig) GetSmarthostOk() (*string, bool) {
	if o == nil || o.Smarthost == nil {
		return nil, false
	}
	return o.Smarthost, true
}

// HasSmarthost returns a boolean if a field has been set.
func (o *EmailConfig) HasSmarthost() bool {
	if o != nil && o.Smarthost != nil {
		return true
	}

	return false
}

// SetSmarthost gets a reference to the given string and assigns it to the Smarthost field.
func (o *EmailConfig) SetSmarthost(v string) {
	o.Smarthost = &v
}

// GetAuthUsername returns the AuthUsername field value if set, zero value otherwise.
func (o *EmailConfig) GetAuthUsername() string {
	if o == nil || o.AuthUsername == nil {
		var ret string
		return ret
	}
	return *o.AuthUsername
}

// GetAuthUsernameOk returns a tuple with the AuthUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConfig) GetAuthUsernameOk() (*string, bool) {
	if o == nil || o.AuthUsername == nil {
		return nil, false
	}
	return o.AuthUsername, true
}

// HasAuthUsername returns a boolean if a field has been set.
func (o *EmailConfig) HasAuthUsername() bool {
	if o != nil && o.AuthUsername != nil {
		return true
	}

	return false
}

// SetAuthUsername gets a reference to the given string and assigns it to the AuthUsername field.
func (o *EmailConfig) SetAuthUsername(v string) {
	o.AuthUsername = &v
}

// GetAuthPassword returns the AuthPassword field value if set, zero value otherwise.
func (o *EmailConfig) GetAuthPassword() string {
	if o == nil || o.AuthPassword == nil {
		var ret string
		return ret
	}
	return *o.AuthPassword
}

// GetAuthPasswordOk returns a tuple with the AuthPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConfig) GetAuthPasswordOk() (*string, bool) {
	if o == nil || o.AuthPassword == nil {
		return nil, false
	}
	return o.AuthPassword, true
}

// HasAuthPassword returns a boolean if a field has been set.
func (o *EmailConfig) HasAuthPassword() bool {
	if o != nil && o.AuthPassword != nil {
		return true
	}

	return false
}

// SetAuthPassword gets a reference to the given string and assigns it to the AuthPassword field.
func (o *EmailConfig) SetAuthPassword(v string) {
	o.AuthPassword = &v
}

// GetAuthIdentity returns the AuthIdentity field value if set, zero value otherwise.
func (o *EmailConfig) GetAuthIdentity() string {
	if o == nil || o.AuthIdentity == nil {
		var ret string
		return ret
	}
	return *o.AuthIdentity
}

// GetAuthIdentityOk returns a tuple with the AuthIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailConfig) GetAuthIdentityOk() (*string, bool) {
	if o == nil || o.AuthIdentity == nil {
		return nil, false
	}
	return o.AuthIdentity, true
}

// HasAuthIdentity returns a boolean if a field has been set.
func (o *EmailConfig) HasAuthIdentity() bool {
	if o != nil && o.AuthIdentity != nil {
		return true
	}

	return false
}

// SetAuthIdentity gets a reference to the given string and assigns it to the AuthIdentity field.
func (o *EmailConfig) SetAuthIdentity(v string) {
	o.AuthIdentity = &v
}

func (o EmailConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["to"] = o.To
	}
	if o.FromPerson != nil {
		toSerialize["fromPerson"] = o.FromPerson
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

type NullableEmailConfig struct {
	value *EmailConfig
	isSet bool
}

func (v NullableEmailConfig) Get() *EmailConfig {
	return v.value
}

func (v *NullableEmailConfig) Set(val *EmailConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailConfig(val *EmailConfig) *NullableEmailConfig {
	return &NullableEmailConfig{value: val, isSet: true}
}

func (v NullableEmailConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}