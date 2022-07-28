# EmailConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **string** |  | 
**FromPerson** | Pointer to **string** |  | [optional] 
**Smarthost** | Pointer to **string** |  | [optional] 
**AuthUsername** | Pointer to **string** |  | [optional] 
**AuthPassword** | Pointer to **string** |  | [optional] 
**AuthIdentity** | Pointer to **string** |  | [optional] 

## Methods

### NewEmailConfig

`func NewEmailConfig(to string, ) *EmailConfig`

NewEmailConfig instantiates a new EmailConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailConfigWithDefaults

`func NewEmailConfigWithDefaults() *EmailConfig`

NewEmailConfigWithDefaults instantiates a new EmailConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *EmailConfig) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailConfig) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailConfig) SetTo(v string)`

SetTo sets To field to given value.


### GetFromPerson

`func (o *EmailConfig) GetFromPerson() string`

GetFromPerson returns the FromPerson field if non-nil, zero value otherwise.

### GetFromPersonOk

`func (o *EmailConfig) GetFromPersonOk() (*string, bool)`

GetFromPersonOk returns a tuple with the FromPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPerson

`func (o *EmailConfig) SetFromPerson(v string)`

SetFromPerson sets FromPerson field to given value.

### HasFromPerson

`func (o *EmailConfig) HasFromPerson() bool`

HasFromPerson returns a boolean if a field has been set.

### GetSmarthost

`func (o *EmailConfig) GetSmarthost() string`

GetSmarthost returns the Smarthost field if non-nil, zero value otherwise.

### GetSmarthostOk

`func (o *EmailConfig) GetSmarthostOk() (*string, bool)`

GetSmarthostOk returns a tuple with the Smarthost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmarthost

`func (o *EmailConfig) SetSmarthost(v string)`

SetSmarthost sets Smarthost field to given value.

### HasSmarthost

`func (o *EmailConfig) HasSmarthost() bool`

HasSmarthost returns a boolean if a field has been set.

### GetAuthUsername

`func (o *EmailConfig) GetAuthUsername() string`

GetAuthUsername returns the AuthUsername field if non-nil, zero value otherwise.

### GetAuthUsernameOk

`func (o *EmailConfig) GetAuthUsernameOk() (*string, bool)`

GetAuthUsernameOk returns a tuple with the AuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsername

`func (o *EmailConfig) SetAuthUsername(v string)`

SetAuthUsername sets AuthUsername field to given value.

### HasAuthUsername

`func (o *EmailConfig) HasAuthUsername() bool`

HasAuthUsername returns a boolean if a field has been set.

### GetAuthPassword

`func (o *EmailConfig) GetAuthPassword() string`

GetAuthPassword returns the AuthPassword field if non-nil, zero value otherwise.

### GetAuthPasswordOk

`func (o *EmailConfig) GetAuthPasswordOk() (*string, bool)`

GetAuthPasswordOk returns a tuple with the AuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPassword

`func (o *EmailConfig) SetAuthPassword(v string)`

SetAuthPassword sets AuthPassword field to given value.

### HasAuthPassword

`func (o *EmailConfig) HasAuthPassword() bool`

HasAuthPassword returns a boolean if a field has been set.

### GetAuthIdentity

`func (o *EmailConfig) GetAuthIdentity() string`

GetAuthIdentity returns the AuthIdentity field if non-nil, zero value otherwise.

### GetAuthIdentityOk

`func (o *EmailConfig) GetAuthIdentityOk() (*string, bool)`

GetAuthIdentityOk returns a tuple with the AuthIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthIdentity

`func (o *EmailConfig) SetAuthIdentity(v string)`

SetAuthIdentity sets AuthIdentity field to given value.

### HasAuthIdentity

`func (o *EmailConfig) HasAuthIdentity() bool`

HasAuthIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


