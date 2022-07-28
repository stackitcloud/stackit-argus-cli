# AlertGroupJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Interval** | Pointer to **string** |  | [optional] [default to "60s"]
**Rules** | [**[]AlertRuleRecordJson**](AlertRuleRecordJson.md) |  | 

## Methods

### NewAlertGroupJson

`func NewAlertGroupJson(name string, rules []AlertRuleRecordJson, ) *AlertGroupJson`

NewAlertGroupJson instantiates a new AlertGroupJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertGroupJsonWithDefaults

`func NewAlertGroupJsonWithDefaults() *AlertGroupJson`

NewAlertGroupJsonWithDefaults instantiates a new AlertGroupJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AlertGroupJson) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertGroupJson) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertGroupJson) SetName(v string)`

SetName sets Name field to given value.


### GetInterval

`func (o *AlertGroupJson) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *AlertGroupJson) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *AlertGroupJson) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *AlertGroupJson) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetRules

`func (o *AlertGroupJson) GetRules() []AlertRuleRecordJson`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *AlertGroupJson) GetRulesOk() (*[]AlertRuleRecordJson, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *AlertGroupJson) SetRules(v []AlertRuleRecordJson)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


