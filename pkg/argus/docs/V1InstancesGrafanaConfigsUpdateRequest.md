# V1InstancesGrafanaConfigsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicReadAccess** | **bool** | If it is true anyone can access the grafana dashboards without logging in. If it is false a login is required. | 

## Methods

### NewV1InstancesGrafanaConfigsUpdateRequest

`func NewV1InstancesGrafanaConfigsUpdateRequest(publicReadAccess bool, ) *V1InstancesGrafanaConfigsUpdateRequest`

NewV1InstancesGrafanaConfigsUpdateRequest instantiates a new V1InstancesGrafanaConfigsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesGrafanaConfigsUpdateRequestWithDefaults

`func NewV1InstancesGrafanaConfigsUpdateRequestWithDefaults() *V1InstancesGrafanaConfigsUpdateRequest`

NewV1InstancesGrafanaConfigsUpdateRequestWithDefaults instantiates a new V1InstancesGrafanaConfigsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicReadAccess

`func (o *V1InstancesGrafanaConfigsUpdateRequest) GetPublicReadAccess() bool`

GetPublicReadAccess returns the PublicReadAccess field if non-nil, zero value otherwise.

### GetPublicReadAccessOk

`func (o *V1InstancesGrafanaConfigsUpdateRequest) GetPublicReadAccessOk() (*bool, bool)`

GetPublicReadAccessOk returns a tuple with the PublicReadAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicReadAccess

`func (o *V1InstancesGrafanaConfigsUpdateRequest) SetPublicReadAccess(v bool)`

SetPublicReadAccess sets PublicReadAccess field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


