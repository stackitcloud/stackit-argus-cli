# ApiUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**ApiUsers** | **[]string** |  | 

## Methods

### NewApiUser

`func NewApiUser(message string, apiUsers []string, ) *ApiUser`

NewApiUser instantiates a new ApiUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUserWithDefaults

`func NewApiUserWithDefaults() *ApiUser`

NewApiUserWithDefaults instantiates a new ApiUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ApiUser) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiUser) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiUser) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetApiUsers

`func (o *ApiUser) GetApiUsers() []string`

GetApiUsers returns the ApiUsers field if non-nil, zero value otherwise.

### GetApiUsersOk

`func (o *ApiUser) GetApiUsersOk() (*[]string, bool)`

GetApiUsersOk returns a tuple with the ApiUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUsers

`func (o *ApiUser) SetApiUsers(v []string)`

SetApiUsers sets ApiUsers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


