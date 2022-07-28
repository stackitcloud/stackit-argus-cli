# WebHook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**MsTeams** | Pointer to **bool** |  | [optional] 

## Methods

### NewWebHook

`func NewWebHook(url string, ) *WebHook`

NewWebHook instantiates a new WebHook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebHookWithDefaults

`func NewWebHookWithDefaults() *WebHook`

NewWebHookWithDefaults instantiates a new WebHook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *WebHook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebHook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebHook) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMsTeams

`func (o *WebHook) GetMsTeams() bool`

GetMsTeams returns the MsTeams field if non-nil, zero value otherwise.

### GetMsTeamsOk

`func (o *WebHook) GetMsTeamsOk() (*bool, bool)`

GetMsTeamsOk returns a tuple with the MsTeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsTeams

`func (o *WebHook) SetMsTeams(v bool)`

SetMsTeams sets MsTeams field to given value.

### HasMsTeams

`func (o *WebHook) HasMsTeams() bool`

HasMsTeams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


