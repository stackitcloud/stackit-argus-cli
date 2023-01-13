# AlertGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Interval** | **string** |  | 
**Rules** | [**[]AlertRule**](AlertRule.md) |  | 

## Methods

### NewAlertGroup

`func NewAlertGroup(name string, interval string, rules []AlertRule, ) *AlertGroup`

NewAlertGroup instantiates a new AlertGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertGroupWithDefaults

`func NewAlertGroupWithDefaults() *AlertGroup`

NewAlertGroupWithDefaults instantiates a new AlertGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AlertGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertGroup) SetName(v string)`

SetName sets Name field to given value.


### GetInterval

`func (o *AlertGroup) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *AlertGroup) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *AlertGroup) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetRules

`func (o *AlertGroup) GetRules() []AlertRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *AlertGroup) GetRulesOk() (*[]AlertRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *AlertGroup) SetRules(v []AlertRule)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


