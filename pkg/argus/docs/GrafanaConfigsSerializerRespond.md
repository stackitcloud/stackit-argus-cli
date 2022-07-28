# GrafanaConfigsSerializerRespond

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicReadAccess** | Pointer to **bool** |  | [optional] 
**GenericOauth** | Pointer to [**GrafanaOauth**](GrafanaOauth.md) |  | [optional] 
**Message** | **string** |  | 

## Methods

### NewGrafanaConfigsSerializerRespond

`func NewGrafanaConfigsSerializerRespond(message string, ) *GrafanaConfigsSerializerRespond`

NewGrafanaConfigsSerializerRespond instantiates a new GrafanaConfigsSerializerRespond object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrafanaConfigsSerializerRespondWithDefaults

`func NewGrafanaConfigsSerializerRespondWithDefaults() *GrafanaConfigsSerializerRespond`

NewGrafanaConfigsSerializerRespondWithDefaults instantiates a new GrafanaConfigsSerializerRespond object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicReadAccess

`func (o *GrafanaConfigsSerializerRespond) GetPublicReadAccess() bool`

GetPublicReadAccess returns the PublicReadAccess field if non-nil, zero value otherwise.

### GetPublicReadAccessOk

`func (o *GrafanaConfigsSerializerRespond) GetPublicReadAccessOk() (*bool, bool)`

GetPublicReadAccessOk returns a tuple with the PublicReadAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicReadAccess

`func (o *GrafanaConfigsSerializerRespond) SetPublicReadAccess(v bool)`

SetPublicReadAccess sets PublicReadAccess field to given value.

### HasPublicReadAccess

`func (o *GrafanaConfigsSerializerRespond) HasPublicReadAccess() bool`

HasPublicReadAccess returns a boolean if a field has been set.

### GetGenericOauth

`func (o *GrafanaConfigsSerializerRespond) GetGenericOauth() GrafanaOauth`

GetGenericOauth returns the GenericOauth field if non-nil, zero value otherwise.

### GetGenericOauthOk

`func (o *GrafanaConfigsSerializerRespond) GetGenericOauthOk() (*GrafanaOauth, bool)`

GetGenericOauthOk returns a tuple with the GenericOauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericOauth

`func (o *GrafanaConfigsSerializerRespond) SetGenericOauth(v GrafanaOauth)`

SetGenericOauth sets GenericOauth field to given value.

### HasGenericOauth

`func (o *GrafanaConfigsSerializerRespond) HasGenericOauth() bool`

HasGenericOauth returns a boolean if a field has been set.

### GetMessage

`func (o *GrafanaConfigsSerializerRespond) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GrafanaConfigsSerializerRespond) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GrafanaConfigsSerializerRespond) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


