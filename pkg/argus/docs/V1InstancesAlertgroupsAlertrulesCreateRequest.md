# V1InstancesAlertgroupsAlertrulesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | **string** | The name of the alert. Must be unique. | 
**Expr** | **string** | The PromQL expression to evaluate. Every evaluation cycle this is evaluated at the current time, and all resultant time series become pending/firing alerts. | 
**For** | Pointer to **string** | Alerts are considered firing once they have been returned for this long. Alerts which have not yet fired for long enough are considered pending. | [optional] [default to "0s"]
**Labels** | **map[string]interface{}** | map of key:value. Labels to add or overwrite for each alert. | 
**Annotations** | **map[string]interface{}** | map of key:value. Annotations to add to each alert. | 

## Methods

### NewV1InstancesAlertgroupsAlertrulesCreateRequest

`func NewV1InstancesAlertgroupsAlertrulesCreateRequest(alert string, expr string, labels map[string]interface{}, annotations map[string]interface{}, ) *V1InstancesAlertgroupsAlertrulesCreateRequest`

NewV1InstancesAlertgroupsAlertrulesCreateRequest instantiates a new V1InstancesAlertgroupsAlertrulesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesAlertgroupsAlertrulesCreateRequestWithDefaults

`func NewV1InstancesAlertgroupsAlertrulesCreateRequestWithDefaults() *V1InstancesAlertgroupsAlertrulesCreateRequest`

NewV1InstancesAlertgroupsAlertrulesCreateRequestWithDefaults instantiates a new V1InstancesAlertgroupsAlertrulesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlert

`func (o *V1InstancesAlertgroupsAlertrulesCreateRequest) GetAlert() string`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *V1InstancesAlertgroupsAlertrulesCreateRequest) GetAlertOk() (*string, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *V1InstancesAlertgroupsAlertrulesCreateRequest) SetAlert(v string)`

SetAlert sets Alert field to given value.


### GetExpr

`func (o *V1InstancesAlertgroupsAlertrulesCreateRequest) GetExpr() string`

GetExpr returns the Expr field if non-nil, zero value otherwise.

### GetExprOk

`func (o *V1InstancesAlertgroupsAlertrulesCreateRequest) GetExprOk() (*string, bool)`

GetExprOk returns a tuple with the Expr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpr

`func (o *V1InstancesAlertgroupsAlertrulesCreateRequest) SetExpr(v string)`

SetExpr sets Expr field to given value.


### GetFor

`func (o *V1InstancesAlertgroupsAlertrulesCreateRequest) GetFor() string`

GetFor returns the For field if non-nil, zero value otherwise.

### GetForOk

`func (o *V1InstancesAlertgroupsAlertrulesCreateRequest) GetForOk() (*string, bool)`

GetForOk returns a tuple with the For field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFor

`func (o *V1InstancesAlertgroupsAlertrulesCreateRequest) SetFor(v string)`

SetFor sets For field to given value.

### HasFor

`func (o *V1InstancesAlertgroupsAlertrulesCreateRequest) HasFor() bool`

HasFor returns a boolean if a field has been set.

### GetLabels

`func (o *V1InstancesAlertgroupsAlertrulesCreateRequest) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *V1InstancesAlertgroupsAlertrulesCreateRequest) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *V1InstancesAlertgroupsAlertrulesCreateRequest) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.


### GetAnnotations

`func (o *V1InstancesAlertgroupsAlertrulesCreateRequest) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *V1InstancesAlertgroupsAlertrulesCreateRequest) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *V1InstancesAlertgroupsAlertrulesCreateRequest) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


