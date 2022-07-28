# AlertRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | **string** |  | 
**Expr** | **string** |  | 
**Labels** | **map[string]string** |  | 
**Annotations** | **map[string]string** |  | 
**For** | Pointer to **string** |  | [optional] [default to "0s"]

## Methods

### NewAlertRule

`func NewAlertRule(alert string, expr string, labels map[string]string, annotations map[string]string, ) *AlertRule`

NewAlertRule instantiates a new AlertRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRuleWithDefaults

`func NewAlertRuleWithDefaults() *AlertRule`

NewAlertRuleWithDefaults instantiates a new AlertRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlert

`func (o *AlertRule) GetAlert() string`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *AlertRule) GetAlertOk() (*string, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *AlertRule) SetAlert(v string)`

SetAlert sets Alert field to given value.


### GetExpr

`func (o *AlertRule) GetExpr() string`

GetExpr returns the Expr field if non-nil, zero value otherwise.

### GetExprOk

`func (o *AlertRule) GetExprOk() (*string, bool)`

GetExprOk returns a tuple with the Expr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpr

`func (o *AlertRule) SetExpr(v string)`

SetExpr sets Expr field to given value.


### GetLabels

`func (o *AlertRule) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AlertRule) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AlertRule) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.


### GetAnnotations

`func (o *AlertRule) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *AlertRule) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *AlertRule) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.


### GetFor

`func (o *AlertRule) GetFor() string`

GetFor returns the For field if non-nil, zero value otherwise.

### GetForOk

`func (o *AlertRule) GetForOk() (*string, bool)`

GetForOk returns a tuple with the For field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFor

`func (o *AlertRule) SetFor(v string)`

SetFor sets For field to given value.

### HasFor

`func (o *AlertRule) HasFor() bool`

HasFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


