# V1InstancesAlertconfigsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Global** | Pointer to [**V1InstancesAlertconfigsUpdateRequestGlobal**](V1InstancesAlertconfigsUpdateRequestGlobal.md) |  | [optional] 
**Route** | [**V1InstancesAlertconfigsUpdateRequestRoute**](V1InstancesAlertconfigsUpdateRequestRoute.md) |  | 
**Receivers** | [**[]V1InstancesAlertconfigsUpdateRequestReceiversInner**](V1InstancesAlertconfigsUpdateRequestReceiversInner.md) | A list of notification receivers. | 
**InhibitRules** | Pointer to [**V1InstancesAlertconfigsUpdateRequestInhibitRules**](V1InstancesAlertconfigsUpdateRequestInhibitRules.md) |  | [optional] 

## Methods

### NewV1InstancesAlertconfigsUpdateRequest

`func NewV1InstancesAlertconfigsUpdateRequest(route V1InstancesAlertconfigsUpdateRequestRoute, receivers []V1InstancesAlertconfigsUpdateRequestReceiversInner, ) *V1InstancesAlertconfigsUpdateRequest`

NewV1InstancesAlertconfigsUpdateRequest instantiates a new V1InstancesAlertconfigsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesAlertconfigsUpdateRequestWithDefaults

`func NewV1InstancesAlertconfigsUpdateRequestWithDefaults() *V1InstancesAlertconfigsUpdateRequest`

NewV1InstancesAlertconfigsUpdateRequestWithDefaults instantiates a new V1InstancesAlertconfigsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobal

`func (o *V1InstancesAlertconfigsUpdateRequest) GetGlobal() V1InstancesAlertconfigsUpdateRequestGlobal`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *V1InstancesAlertconfigsUpdateRequest) GetGlobalOk() (*V1InstancesAlertconfigsUpdateRequestGlobal, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *V1InstancesAlertconfigsUpdateRequest) SetGlobal(v V1InstancesAlertconfigsUpdateRequestGlobal)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *V1InstancesAlertconfigsUpdateRequest) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.

### GetRoute

`func (o *V1InstancesAlertconfigsUpdateRequest) GetRoute() V1InstancesAlertconfigsUpdateRequestRoute`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *V1InstancesAlertconfigsUpdateRequest) GetRouteOk() (*V1InstancesAlertconfigsUpdateRequestRoute, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *V1InstancesAlertconfigsUpdateRequest) SetRoute(v V1InstancesAlertconfigsUpdateRequestRoute)`

SetRoute sets Route field to given value.


### GetReceivers

`func (o *V1InstancesAlertconfigsUpdateRequest) GetReceivers() []V1InstancesAlertconfigsUpdateRequestReceiversInner`

GetReceivers returns the Receivers field if non-nil, zero value otherwise.

### GetReceiversOk

`func (o *V1InstancesAlertconfigsUpdateRequest) GetReceiversOk() (*[]V1InstancesAlertconfigsUpdateRequestReceiversInner, bool)`

GetReceiversOk returns a tuple with the Receivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivers

`func (o *V1InstancesAlertconfigsUpdateRequest) SetReceivers(v []V1InstancesAlertconfigsUpdateRequestReceiversInner)`

SetReceivers sets Receivers field to given value.


### GetInhibitRules

`func (o *V1InstancesAlertconfigsUpdateRequest) GetInhibitRules() V1InstancesAlertconfigsUpdateRequestInhibitRules`

GetInhibitRules returns the InhibitRules field if non-nil, zero value otherwise.

### GetInhibitRulesOk

`func (o *V1InstancesAlertconfigsUpdateRequest) GetInhibitRulesOk() (*V1InstancesAlertconfigsUpdateRequestInhibitRules, bool)`

GetInhibitRulesOk returns a tuple with the InhibitRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInhibitRules

`func (o *V1InstancesAlertconfigsUpdateRequest) SetInhibitRules(v V1InstancesAlertconfigsUpdateRequestInhibitRules)`

SetInhibitRules sets InhibitRules field to given value.

### HasInhibitRules

`func (o *V1InstancesAlertconfigsUpdateRequest) HasInhibitRules() bool`

HasInhibitRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


