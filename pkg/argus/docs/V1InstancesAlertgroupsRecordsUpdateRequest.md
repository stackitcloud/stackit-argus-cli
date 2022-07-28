# V1InstancesAlertgroupsRecordsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expr** | **string** | The PromQL expression to evaluate. Every evaluation cycle this is evaluated at the current time, and all resultant time series become pending/firing alerts. | 
**Labels** | Pointer to **map[string]interface{}** | map of key:value. Labels to add or overwrite for each alert. &#x60;Additional Validators:&#x60; * should not contain more than 5 keys * each key and value should not be longer than 200 characters  | [optional] 

## Methods

### NewV1InstancesAlertgroupsRecordsUpdateRequest

`func NewV1InstancesAlertgroupsRecordsUpdateRequest(expr string, ) *V1InstancesAlertgroupsRecordsUpdateRequest`

NewV1InstancesAlertgroupsRecordsUpdateRequest instantiates a new V1InstancesAlertgroupsRecordsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesAlertgroupsRecordsUpdateRequestWithDefaults

`func NewV1InstancesAlertgroupsRecordsUpdateRequestWithDefaults() *V1InstancesAlertgroupsRecordsUpdateRequest`

NewV1InstancesAlertgroupsRecordsUpdateRequestWithDefaults instantiates a new V1InstancesAlertgroupsRecordsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpr

`func (o *V1InstancesAlertgroupsRecordsUpdateRequest) GetExpr() string`

GetExpr returns the Expr field if non-nil, zero value otherwise.

### GetExprOk

`func (o *V1InstancesAlertgroupsRecordsUpdateRequest) GetExprOk() (*string, bool)`

GetExprOk returns a tuple with the Expr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpr

`func (o *V1InstancesAlertgroupsRecordsUpdateRequest) SetExpr(v string)`

SetExpr sets Expr field to given value.


### GetLabels

`func (o *V1InstancesAlertgroupsRecordsUpdateRequest) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *V1InstancesAlertgroupsRecordsUpdateRequest) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *V1InstancesAlertgroupsRecordsUpdateRequest) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *V1InstancesAlertgroupsRecordsUpdateRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


