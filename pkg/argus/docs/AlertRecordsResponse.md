# AlertRecordsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Data** | [**[]AlertRecord**](AlertRecord.md) |  | 

## Methods

### NewAlertRecordsResponse

`func NewAlertRecordsResponse(message string, data []AlertRecord, ) *AlertRecordsResponse`

NewAlertRecordsResponse instantiates a new AlertRecordsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRecordsResponseWithDefaults

`func NewAlertRecordsResponseWithDefaults() *AlertRecordsResponse`

NewAlertRecordsResponseWithDefaults instantiates a new AlertRecordsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AlertRecordsResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AlertRecordsResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AlertRecordsResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *AlertRecordsResponse) GetData() []AlertRecord`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AlertRecordsResponse) GetDataOk() (*[]AlertRecord, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AlertRecordsResponse) SetData(v []AlertRecord)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


