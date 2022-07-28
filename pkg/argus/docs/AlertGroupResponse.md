# AlertGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Data** | [**AlertGroupJson**](AlertGroupJson.md) |  | 

## Methods

### NewAlertGroupResponse

`func NewAlertGroupResponse(message string, data AlertGroupJson, ) *AlertGroupResponse`

NewAlertGroupResponse instantiates a new AlertGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertGroupResponseWithDefaults

`func NewAlertGroupResponseWithDefaults() *AlertGroupResponse`

NewAlertGroupResponseWithDefaults instantiates a new AlertGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AlertGroupResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AlertGroupResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AlertGroupResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *AlertGroupResponse) GetData() AlertGroupJson`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AlertGroupResponse) GetDataOk() (*AlertGroupJson, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AlertGroupResponse) SetData(v AlertGroupJson)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


