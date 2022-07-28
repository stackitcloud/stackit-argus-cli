# V1InstancesAlertgroupsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the group. Must be unique. | 
**Interval** | **string** | How often rules in the group are evaluated. | 
**Rules** | [**[]V1InstancesAlertgroupsCreateRequestRulesInner**](V1InstancesAlertgroupsCreateRequestRulesInner.md) | rules for the alert group | 

## Methods

### NewV1InstancesAlertgroupsCreateRequest

`func NewV1InstancesAlertgroupsCreateRequest(name string, interval string, rules []V1InstancesAlertgroupsCreateRequestRulesInner, ) *V1InstancesAlertgroupsCreateRequest`

NewV1InstancesAlertgroupsCreateRequest instantiates a new V1InstancesAlertgroupsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesAlertgroupsCreateRequestWithDefaults

`func NewV1InstancesAlertgroupsCreateRequestWithDefaults() *V1InstancesAlertgroupsCreateRequest`

NewV1InstancesAlertgroupsCreateRequestWithDefaults instantiates a new V1InstancesAlertgroupsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V1InstancesAlertgroupsCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1InstancesAlertgroupsCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1InstancesAlertgroupsCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetInterval

`func (o *V1InstancesAlertgroupsCreateRequest) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *V1InstancesAlertgroupsCreateRequest) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *V1InstancesAlertgroupsCreateRequest) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetRules

`func (o *V1InstancesAlertgroupsCreateRequest) GetRules() []V1InstancesAlertgroupsCreateRequestRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *V1InstancesAlertgroupsCreateRequest) GetRulesOk() (*[]V1InstancesAlertgroupsCreateRequestRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *V1InstancesAlertgroupsCreateRequest) SetRules(v []V1InstancesAlertgroupsCreateRequestRulesInner)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


