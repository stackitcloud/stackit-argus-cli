# ProjectInstancesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Instances** | [**[]ProjectInstance**](ProjectInstance.md) |  | 

## Methods

### NewProjectInstancesResponse

`func NewProjectInstancesResponse(message string, instances []ProjectInstance, ) *ProjectInstancesResponse`

NewProjectInstancesResponse instantiates a new ProjectInstancesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectInstancesResponseWithDefaults

`func NewProjectInstancesResponseWithDefaults() *ProjectInstancesResponse`

NewProjectInstancesResponseWithDefaults instantiates a new ProjectInstancesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ProjectInstancesResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProjectInstancesResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProjectInstancesResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetInstances

`func (o *ProjectInstancesResponse) GetInstances() []ProjectInstance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ProjectInstancesResponse) GetInstancesOk() (*[]ProjectInstance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ProjectInstancesResponse) SetInstances(v []ProjectInstance)`

SetInstances sets Instances field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


