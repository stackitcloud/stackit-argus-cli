# AlertGroupsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Data** | [**[]AlertGroupJson**](AlertGroupJson.md) |  | 

## Methods

### NewAlertGroupsResponse

`func NewAlertGroupsResponse(message string, data []AlertGroupJson, ) *AlertGroupsResponse`

NewAlertGroupsResponse instantiates a new AlertGroupsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertGroupsResponseWithDefaults

`func NewAlertGroupsResponseWithDefaults() *AlertGroupsResponse`

NewAlertGroupsResponseWithDefaults instantiates a new AlertGroupsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AlertGroupsResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AlertGroupsResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AlertGroupsResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *AlertGroupsResponse) GetData() []AlertGroupJson`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AlertGroupsResponse) GetDataOk() (*[]AlertGroupJson, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AlertGroupsResponse) SetData(v []AlertGroupJson)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


