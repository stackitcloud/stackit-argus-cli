# V1InstancesAlertgroupsRecordsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Record** | **string** | The name of the record. &#x60;Additional Validators:&#x60; * is the identifier and so unique in the group  | 
**Expr** | **string** | The PromQL expression to evaluate. Every evaluation cycle this is evaluated at the current time, and all resultant time series become pending/firing alerts. | 
**Labels** | Pointer to **map[string]interface{}** | map of key:value. Labels to add or overwrite for each alert. &#x60;Additional Validators:&#x60; * should not contain more than 10 keys * each key and value should not be longer than 200 characters  | [optional] 

## Methods

### NewV1InstancesAlertgroupsRecordsCreateRequest

`func NewV1InstancesAlertgroupsRecordsCreateRequest(record string, expr string, ) *V1InstancesAlertgroupsRecordsCreateRequest`

NewV1InstancesAlertgroupsRecordsCreateRequest instantiates a new V1InstancesAlertgroupsRecordsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesAlertgroupsRecordsCreateRequestWithDefaults

`func NewV1InstancesAlertgroupsRecordsCreateRequestWithDefaults() *V1InstancesAlertgroupsRecordsCreateRequest`

NewV1InstancesAlertgroupsRecordsCreateRequestWithDefaults instantiates a new V1InstancesAlertgroupsRecordsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecord

`func (o *V1InstancesAlertgroupsRecordsCreateRequest) GetRecord() string`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *V1InstancesAlertgroupsRecordsCreateRequest) GetRecordOk() (*string, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *V1InstancesAlertgroupsRecordsCreateRequest) SetRecord(v string)`

SetRecord sets Record field to given value.


### GetExpr

`func (o *V1InstancesAlertgroupsRecordsCreateRequest) GetExpr() string`

GetExpr returns the Expr field if non-nil, zero value otherwise.

### GetExprOk

`func (o *V1InstancesAlertgroupsRecordsCreateRequest) GetExprOk() (*string, bool)`

GetExprOk returns a tuple with the Expr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpr

`func (o *V1InstancesAlertgroupsRecordsCreateRequest) SetExpr(v string)`

SetExpr sets Expr field to given value.


### GetLabels

`func (o *V1InstancesAlertgroupsRecordsCreateRequest) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *V1InstancesAlertgroupsRecordsCreateRequest) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *V1InstancesAlertgroupsRecordsCreateRequest) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *V1InstancesAlertgroupsRecordsCreateRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


