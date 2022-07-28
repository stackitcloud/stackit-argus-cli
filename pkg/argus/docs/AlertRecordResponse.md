# AlertRecordResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Data** | [**AlertRecord**](AlertRecord.md) |  | 

## Methods

### NewAlertRecordResponse

`func NewAlertRecordResponse(message string, data AlertRecord, ) *AlertRecordResponse`

NewAlertRecordResponse instantiates a new AlertRecordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRecordResponseWithDefaults

`func NewAlertRecordResponseWithDefaults() *AlertRecordResponse`

NewAlertRecordResponseWithDefaults instantiates a new AlertRecordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AlertRecordResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AlertRecordResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AlertRecordResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *AlertRecordResponse) GetData() AlertRecord`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AlertRecordResponse) GetDataOk() (*AlertRecord, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AlertRecordResponse) SetData(v AlertRecord)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


