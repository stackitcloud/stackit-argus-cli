# V1InstancesAlertgroupsCreateRequestRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | **string** | The name of the alert. &#x60;Additional Validators:&#x60; * is the identifier and so unique in the group * should only include the characters: a-zA-Z0-9- | 
**Expr** | **string** | The PromQL expression to evaluate. Every evaluation cycle this is evaluated at the current time, and all resultant time series become pending/firing alerts. | 
**For** | Pointer to **string** | Alerts are considered firing once they have been returned for this long. Alerts which have not yet fired for long enough are considered pending. &#x60;Additional Validators:&#x60; * must be a valid time string | [optional] [default to "0s"]
**Labels** | Pointer to **map[string]interface{}** | map of key:value. Labels to add or overwrite for each alert. &#x60;Additional Validators:&#x60; * should not contain more than 10 keys * each key and value should not be longer than 200 characters | [optional] 
**Annotations** | Pointer to **map[string]interface{}** | map of key:value. Annotations to add to each alert. &#x60;Additional Validators:&#x60; * should not contain more than 5 keys * each key and value should not be longer than 200 characters | [optional] 

## Methods

### NewV1InstancesAlertgroupsCreateRequestRulesInner

`func NewV1InstancesAlertgroupsCreateRequestRulesInner(alert string, expr string, ) *V1InstancesAlertgroupsCreateRequestRulesInner`

NewV1InstancesAlertgroupsCreateRequestRulesInner instantiates a new V1InstancesAlertgroupsCreateRequestRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesAlertgroupsCreateRequestRulesInnerWithDefaults

`func NewV1InstancesAlertgroupsCreateRequestRulesInnerWithDefaults() *V1InstancesAlertgroupsCreateRequestRulesInner`

NewV1InstancesAlertgroupsCreateRequestRulesInnerWithDefaults instantiates a new V1InstancesAlertgroupsCreateRequestRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlert

`func (o *V1InstancesAlertgroupsCreateRequestRulesInner) GetAlert() string`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *V1InstancesAlertgroupsCreateRequestRulesInner) GetAlertOk() (*string, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *V1InstancesAlertgroupsCreateRequestRulesInner) SetAlert(v string)`

SetAlert sets Alert field to given value.


### GetExpr

`func (o *V1InstancesAlertgroupsCreateRequestRulesInner) GetExpr() string`

GetExpr returns the Expr field if non-nil, zero value otherwise.

### GetExprOk

`func (o *V1InstancesAlertgroupsCreateRequestRulesInner) GetExprOk() (*string, bool)`

GetExprOk returns a tuple with the Expr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpr

`func (o *V1InstancesAlertgroupsCreateRequestRulesInner) SetExpr(v string)`

SetExpr sets Expr field to given value.


### GetFor

`func (o *V1InstancesAlertgroupsCreateRequestRulesInner) GetFor() string`

GetFor returns the For field if non-nil, zero value otherwise.

### GetForOk

`func (o *V1InstancesAlertgroupsCreateRequestRulesInner) GetForOk() (*string, bool)`

GetForOk returns a tuple with the For field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFor

`func (o *V1InstancesAlertgroupsCreateRequestRulesInner) SetFor(v string)`

SetFor sets For field to given value.

### HasFor

`func (o *V1InstancesAlertgroupsCreateRequestRulesInner) HasFor() bool`

HasFor returns a boolean if a field has been set.

### GetLabels

`func (o *V1InstancesAlertgroupsCreateRequestRulesInner) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *V1InstancesAlertgroupsCreateRequestRulesInner) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *V1InstancesAlertgroupsCreateRequestRulesInner) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *V1InstancesAlertgroupsCreateRequestRulesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAnnotations

`func (o *V1InstancesAlertgroupsCreateRequestRulesInner) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *V1InstancesAlertgroupsCreateRequestRulesInner) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *V1InstancesAlertgroupsCreateRequestRulesInner) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *V1InstancesAlertgroupsCreateRequestRulesInner) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


