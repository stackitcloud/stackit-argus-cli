# Alert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Route** | [**Route**](Route.md) |  | 
**Receivers** | [**[]Receivers**](Receivers.md) |  | 
**InhibitRules** | Pointer to [**[]InhibitRules**](InhibitRules.md) |  | [optional] 
**Global** | Pointer to [**Global**](Global.md) |  | [optional] 

## Methods

### NewAlert

`func NewAlert(route Route, receivers []Receivers, ) *Alert`

NewAlert instantiates a new Alert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertWithDefaults

`func NewAlertWithDefaults() *Alert`

NewAlertWithDefaults instantiates a new Alert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoute

`func (o *Alert) GetRoute() Route`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *Alert) GetRouteOk() (*Route, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *Alert) SetRoute(v Route)`

SetRoute sets Route field to given value.


### GetReceivers

`func (o *Alert) GetReceivers() []Receivers`

GetReceivers returns the Receivers field if non-nil, zero value otherwise.

### GetReceiversOk

`func (o *Alert) GetReceiversOk() (*[]Receivers, bool)`

GetReceiversOk returns a tuple with the Receivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivers

`func (o *Alert) SetReceivers(v []Receivers)`

SetReceivers sets Receivers field to given value.


### GetInhibitRules

`func (o *Alert) GetInhibitRules() []InhibitRules`

GetInhibitRules returns the InhibitRules field if non-nil, zero value otherwise.

### GetInhibitRulesOk

`func (o *Alert) GetInhibitRulesOk() (*[]InhibitRules, bool)`

GetInhibitRulesOk returns a tuple with the InhibitRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInhibitRules

`func (o *Alert) SetInhibitRules(v []InhibitRules)`

SetInhibitRules sets InhibitRules field to given value.

### HasInhibitRules

`func (o *Alert) HasInhibitRules() bool`

HasInhibitRules returns a boolean if a field has been set.

### GetGlobal

`func (o *Alert) GetGlobal() Global`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *Alert) GetGlobalOk() (*Global, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *Alert) SetGlobal(v Global)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *Alert) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


