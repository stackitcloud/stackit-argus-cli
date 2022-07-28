# V1InstancesAlertgroupsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | Pointer to **string** | How often rules in the group are evaluated. &#x60;Additional Validators:&#x60; * must be a valid time string * should be &gt;&#x3D;60s | [optional] [default to "60s"]
**Rules** | [**[]V1InstancesAlertgroupsCreateRequestRulesInner**](V1InstancesAlertgroupsCreateRequestRulesInner.md) | rules for the alert group | 

## Methods

### NewV1InstancesAlertgroupsUpdateRequest

`func NewV1InstancesAlertgroupsUpdateRequest(rules []V1InstancesAlertgroupsCreateRequestRulesInner, ) *V1InstancesAlertgroupsUpdateRequest`

NewV1InstancesAlertgroupsUpdateRequest instantiates a new V1InstancesAlertgroupsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesAlertgroupsUpdateRequestWithDefaults

`func NewV1InstancesAlertgroupsUpdateRequestWithDefaults() *V1InstancesAlertgroupsUpdateRequest`

NewV1InstancesAlertgroupsUpdateRequestWithDefaults instantiates a new V1InstancesAlertgroupsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *V1InstancesAlertgroupsUpdateRequest) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *V1InstancesAlertgroupsUpdateRequest) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *V1InstancesAlertgroupsUpdateRequest) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *V1InstancesAlertgroupsUpdateRequest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetRules

`func (o *V1InstancesAlertgroupsUpdateRequest) GetRules() []V1InstancesAlertgroupsCreateRequestRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *V1InstancesAlertgroupsUpdateRequest) GetRulesOk() (*[]V1InstancesAlertgroupsCreateRequestRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *V1InstancesAlertgroupsUpdateRequest) SetRules(v []V1InstancesAlertgroupsCreateRequestRulesInner)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


