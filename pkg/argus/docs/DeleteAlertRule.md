# DeleteAlertRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Data** | [**[]AlertRule**](AlertRule.md) |  | 

## Methods

### NewDeleteAlertRule

`func NewDeleteAlertRule(message string, data []AlertRule, ) *DeleteAlertRule`

NewDeleteAlertRule instantiates a new DeleteAlertRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteAlertRuleWithDefaults

`func NewDeleteAlertRuleWithDefaults() *DeleteAlertRule`

NewDeleteAlertRuleWithDefaults instantiates a new DeleteAlertRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *DeleteAlertRule) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeleteAlertRule) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeleteAlertRule) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *DeleteAlertRule) GetData() []AlertRule`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeleteAlertRule) GetDataOk() (*[]AlertRule, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeleteAlertRule) SetData(v []AlertRule)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


