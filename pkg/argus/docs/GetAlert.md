# GetAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Data** | [**Alert**](Alert.md) |  | 

## Methods

### NewGetAlert

`func NewGetAlert(message string, data Alert, ) *GetAlert`

NewGetAlert instantiates a new GetAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAlertWithDefaults

`func NewGetAlertWithDefaults() *GetAlert`

NewGetAlertWithDefaults instantiates a new GetAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetAlert) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetAlert) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetAlert) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *GetAlert) GetData() Alert`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAlert) GetDataOk() (*Alert, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAlert) SetData(v Alert)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


