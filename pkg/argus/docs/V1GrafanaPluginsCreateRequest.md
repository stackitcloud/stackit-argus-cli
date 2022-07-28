# V1GrafanaPluginsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the grafana plugin | 
**Description** | **string** | Description of the grafana plugin | 

## Methods

### NewV1GrafanaPluginsCreateRequest

`func NewV1GrafanaPluginsCreateRequest(name string, description string, ) *V1GrafanaPluginsCreateRequest`

NewV1GrafanaPluginsCreateRequest instantiates a new V1GrafanaPluginsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GrafanaPluginsCreateRequestWithDefaults

`func NewV1GrafanaPluginsCreateRequestWithDefaults() *V1GrafanaPluginsCreateRequest`

NewV1GrafanaPluginsCreateRequestWithDefaults instantiates a new V1GrafanaPluginsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V1GrafanaPluginsCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1GrafanaPluginsCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1GrafanaPluginsCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *V1GrafanaPluginsCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1GrafanaPluginsCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1GrafanaPluginsCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


