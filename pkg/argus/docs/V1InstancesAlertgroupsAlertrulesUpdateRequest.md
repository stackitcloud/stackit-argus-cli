# V1InstancesAlertgroupsAlertrulesUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expr** | **string** | The PromQL expression to evaluate. Every evaluation cycle this is evaluated at the current time, and all resultant time series become pending/firing alerts. | 
**For** | Pointer to **string** | Alerts are considered firing once they have been returned for this long. Alerts which have not yet fired for long enough are considered pending. &#x60;Additional Validators:&#x60; * must be a valid time string | [optional] [default to "0s"]
**Labels** | Pointer to **map[string]interface{}** | map of key:value. Labels to add or overwrite for each alert. &#x60;Additional Validators:&#x60; * should not contain more than 5 keys * each key and value should not be longer than 200 characters | [optional] 
**Annotations** | Pointer to **map[string]interface{}** | map of key:value. Annotations to add to each alert. &#x60;Additional Validators:&#x60; * should not contain more than 5 keys * each key and value should not be longer than 200 characters | [optional] 

## Methods

### NewV1InstancesAlertgroupsAlertrulesUpdateRequest

`func NewV1InstancesAlertgroupsAlertrulesUpdateRequest(expr string, ) *V1InstancesAlertgroupsAlertrulesUpdateRequest`

NewV1InstancesAlertgroupsAlertrulesUpdateRequest instantiates a new V1InstancesAlertgroupsAlertrulesUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesAlertgroupsAlertrulesUpdateRequestWithDefaults

`func NewV1InstancesAlertgroupsAlertrulesUpdateRequestWithDefaults() *V1InstancesAlertgroupsAlertrulesUpdateRequest`

NewV1InstancesAlertgroupsAlertrulesUpdateRequestWithDefaults instantiates a new V1InstancesAlertgroupsAlertrulesUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpr

`func (o *V1InstancesAlertgroupsAlertrulesUpdateRequest) GetExpr() string`

GetExpr returns the Expr field if non-nil, zero value otherwise.

### GetExprOk

`func (o *V1InstancesAlertgroupsAlertrulesUpdateRequest) GetExprOk() (*string, bool)`

GetExprOk returns a tuple with the Expr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpr

`func (o *V1InstancesAlertgroupsAlertrulesUpdateRequest) SetExpr(v string)`

SetExpr sets Expr field to given value.


### GetFor

`func (o *V1InstancesAlertgroupsAlertrulesUpdateRequest) GetFor() string`

GetFor returns the For field if non-nil, zero value otherwise.

### GetForOk

`func (o *V1InstancesAlertgroupsAlertrulesUpdateRequest) GetForOk() (*string, bool)`

GetForOk returns a tuple with the For field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFor

`func (o *V1InstancesAlertgroupsAlertrulesUpdateRequest) SetFor(v string)`

SetFor sets For field to given value.

### HasFor

`func (o *V1InstancesAlertgroupsAlertrulesUpdateRequest) HasFor() bool`

HasFor returns a boolean if a field has been set.

### GetLabels

`func (o *V1InstancesAlertgroupsAlertrulesUpdateRequest) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *V1InstancesAlertgroupsAlertrulesUpdateRequest) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *V1InstancesAlertgroupsAlertrulesUpdateRequest) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *V1InstancesAlertgroupsAlertrulesUpdateRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAnnotations

`func (o *V1InstancesAlertgroupsAlertrulesUpdateRequest) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *V1InstancesAlertgroupsAlertrulesUpdateRequest) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *V1InstancesAlertgroupsAlertrulesUpdateRequest) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *V1InstancesAlertgroupsAlertrulesUpdateRequest) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


