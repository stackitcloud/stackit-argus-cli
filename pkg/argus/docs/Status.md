# Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** |  | 
**Description** | **string** |  | 
**InstanceUsable** | **string** |  | 
**UpdateRepeatable** | **string** |  | 

## Methods

### NewStatus

`func NewStatus(state string, description string, instanceUsable string, updateRepeatable string, ) *Status`

NewStatus instantiates a new Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusWithDefaults

`func NewStatusWithDefaults() *Status`

NewStatusWithDefaults instantiates a new Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *Status) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Status) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Status) SetState(v string)`

SetState sets State field to given value.


### GetDescription

`func (o *Status) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Status) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Status) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInstanceUsable

`func (o *Status) GetInstanceUsable() string`

GetInstanceUsable returns the InstanceUsable field if non-nil, zero value otherwise.

### GetInstanceUsableOk

`func (o *Status) GetInstanceUsableOk() (*string, bool)`

GetInstanceUsableOk returns a tuple with the InstanceUsable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUsable

`func (o *Status) SetInstanceUsable(v string)`

SetInstanceUsable sets InstanceUsable field to given value.


### GetUpdateRepeatable

`func (o *Status) GetUpdateRepeatable() string`

GetUpdateRepeatable returns the UpdateRepeatable field if non-nil, zero value otherwise.

### GetUpdateRepeatableOk

`func (o *Status) GetUpdateRepeatableOk() (*string, bool)`

GetUpdateRepeatableOk returns a tuple with the UpdateRepeatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateRepeatable

`func (o *Status) SetUpdateRepeatable(v string)`

SetUpdateRepeatable sets UpdateRepeatable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


