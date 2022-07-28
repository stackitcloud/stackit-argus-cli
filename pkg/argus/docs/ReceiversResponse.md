# ReceiversResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Data** | [**[]Receivers**](Receivers.md) |  | 

## Methods

### NewReceiversResponse

`func NewReceiversResponse(message string, data []Receivers, ) *ReceiversResponse`

NewReceiversResponse instantiates a new ReceiversResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiversResponseWithDefaults

`func NewReceiversResponseWithDefaults() *ReceiversResponse`

NewReceiversResponseWithDefaults instantiates a new ReceiversResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ReceiversResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ReceiversResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ReceiversResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *ReceiversResponse) GetData() []Receivers`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ReceiversResponse) GetDataOk() (*[]Receivers, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ReceiversResponse) SetData(v []Receivers)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


