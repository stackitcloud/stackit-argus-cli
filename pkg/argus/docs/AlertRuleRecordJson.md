# AlertRuleRecordJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | Pointer to **string** |  | [optional] 
**Expr** | **string** |  | 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**Record** | Pointer to **string** |  | [optional] 
**For** | Pointer to **string** |  | [optional] 

## Methods

### NewAlertRuleRecordJson

`func NewAlertRuleRecordJson(expr string, ) *AlertRuleRecordJson`

NewAlertRuleRecordJson instantiates a new AlertRuleRecordJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRuleRecordJsonWithDefaults

`func NewAlertRuleRecordJsonWithDefaults() *AlertRuleRecordJson`

NewAlertRuleRecordJsonWithDefaults instantiates a new AlertRuleRecordJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlert

`func (o *AlertRuleRecordJson) GetAlert() string`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *AlertRuleRecordJson) GetAlertOk() (*string, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *AlertRuleRecordJson) SetAlert(v string)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *AlertRuleRecordJson) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetExpr

`func (o *AlertRuleRecordJson) GetExpr() string`

GetExpr returns the Expr field if non-nil, zero value otherwise.

### GetExprOk

`func (o *AlertRuleRecordJson) GetExprOk() (*string, bool)`

GetExprOk returns a tuple with the Expr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpr

`func (o *AlertRuleRecordJson) SetExpr(v string)`

SetExpr sets Expr field to given value.


### GetLabels

`func (o *AlertRuleRecordJson) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AlertRuleRecordJson) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AlertRuleRecordJson) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AlertRuleRecordJson) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAnnotations

`func (o *AlertRuleRecordJson) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *AlertRuleRecordJson) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *AlertRuleRecordJson) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *AlertRuleRecordJson) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetRecord

`func (o *AlertRuleRecordJson) GetRecord() string`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *AlertRuleRecordJson) GetRecordOk() (*string, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *AlertRuleRecordJson) SetRecord(v string)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *AlertRuleRecordJson) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetFor

`func (o *AlertRuleRecordJson) GetFor() string`

GetFor returns the For field if non-nil, zero value otherwise.

### GetForOk

`func (o *AlertRuleRecordJson) GetForOk() (*string, bool)`

GetForOk returns a tuple with the For field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFor

`func (o *AlertRuleRecordJson) SetFor(v string)`

SetFor sets For field to given value.

### HasFor

`func (o *AlertRuleRecordJson) HasFor() bool`

HasFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


