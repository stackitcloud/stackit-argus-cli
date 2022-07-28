# PutAlertGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Data** | [**[]AlertGroup**](AlertGroup.md) |  | 

## Methods

### NewPutAlertGroup

`func NewPutAlertGroup(message string, data []AlertGroup, ) *PutAlertGroup`

NewPutAlertGroup instantiates a new PutAlertGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutAlertGroupWithDefaults

`func NewPutAlertGroupWithDefaults() *PutAlertGroup`

NewPutAlertGroupWithDefaults instantiates a new PutAlertGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *PutAlertGroup) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PutAlertGroup) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PutAlertGroup) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *PutAlertGroup) GetData() []AlertGroup`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PutAlertGroup) GetDataOk() (*[]AlertGroup, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PutAlertGroup) SetData(v []AlertGroup)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


