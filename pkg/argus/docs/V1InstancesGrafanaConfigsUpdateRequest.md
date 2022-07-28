# V1InstancesGrafanaConfigsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicReadAccess** | Pointer to **bool** | If it&#39;s true, anyone can access the Grafana dashboards without logging in. If it is wrong, a login is required. | [optional] 
**GenericOauth** | Pointer to [**V1InstancesGrafanaConfigsUpdateRequestGenericOauth**](V1InstancesGrafanaConfigsUpdateRequestGenericOauth.md) |  | [optional] 

## Methods

### NewV1InstancesGrafanaConfigsUpdateRequest

`func NewV1InstancesGrafanaConfigsUpdateRequest() *V1InstancesGrafanaConfigsUpdateRequest`

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

### HasPublicReadAccess

`func (o *V1InstancesGrafanaConfigsUpdateRequest) HasPublicReadAccess() bool`

HasPublicReadAccess returns a boolean if a field has been set.

### GetGenericOauth

`func (o *V1InstancesGrafanaConfigsUpdateRequest) GetGenericOauth() V1InstancesGrafanaConfigsUpdateRequestGenericOauth`

GetGenericOauth returns the GenericOauth field if non-nil, zero value otherwise.

### GetGenericOauthOk

`func (o *V1InstancesGrafanaConfigsUpdateRequest) GetGenericOauthOk() (*V1InstancesGrafanaConfigsUpdateRequestGenericOauth, bool)`

GetGenericOauthOk returns a tuple with the GenericOauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericOauth

`func (o *V1InstancesGrafanaConfigsUpdateRequest) SetGenericOauth(v V1InstancesGrafanaConfigsUpdateRequestGenericOauth)`

SetGenericOauth sets GenericOauth field to given value.

### HasGenericOauth

`func (o *V1InstancesGrafanaConfigsUpdateRequest) HasGenericOauth() bool`

HasGenericOauth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


