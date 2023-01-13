# GetAllAlertRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Data** | [**[]AlertRule**](AlertRule.md) |  | 

## Methods

### NewGetAllAlertRules

`func NewGetAllAlertRules(message string, data []AlertRule, ) *GetAllAlertRules`

NewGetAllAlertRules instantiates a new GetAllAlertRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllAlertRulesWithDefaults

`func NewGetAllAlertRulesWithDefaults() *GetAllAlertRules`

NewGetAllAlertRulesWithDefaults instantiates a new GetAllAlertRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetAllAlertRules) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetAllAlertRules) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetAllAlertRules) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *GetAllAlertRules) GetData() []AlertRule`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAllAlertRules) GetDataOk() (*[]AlertRule, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAllAlertRules) SetData(v []AlertRule)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


