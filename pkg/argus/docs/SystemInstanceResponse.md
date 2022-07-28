# SystemInstanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Instance** | [**SystemInstance**](SystemInstance.md) |  | 

## Methods

### NewSystemInstanceResponse

`func NewSystemInstanceResponse(message string, instance SystemInstance, ) *SystemInstanceResponse`

NewSystemInstanceResponse instantiates a new SystemInstanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemInstanceResponseWithDefaults

`func NewSystemInstanceResponseWithDefaults() *SystemInstanceResponse`

NewSystemInstanceResponseWithDefaults instantiates a new SystemInstanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SystemInstanceResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SystemInstanceResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SystemInstanceResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetInstance

`func (o *SystemInstanceResponse) GetInstance() SystemInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *SystemInstanceResponse) GetInstanceOk() (*SystemInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *SystemInstanceResponse) SetInstance(v SystemInstance)`

SetInstance sets Instance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


