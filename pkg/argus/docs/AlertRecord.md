# AlertRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Record** | **string** |  | 
**Expr** | **string** |  | 
**Labels** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAlertRecord

`func NewAlertRecord(record string, expr string, ) *AlertRecord`

NewAlertRecord instantiates a new AlertRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRecordWithDefaults

`func NewAlertRecordWithDefaults() *AlertRecord`

NewAlertRecordWithDefaults instantiates a new AlertRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecord

`func (o *AlertRecord) GetRecord() string`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *AlertRecord) GetRecordOk() (*string, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *AlertRecord) SetRecord(v string)`

SetRecord sets Record field to given value.


### GetExpr

`func (o *AlertRecord) GetExpr() string`

GetExpr returns the Expr field if non-nil, zero value otherwise.

### GetExprOk

`func (o *AlertRecord) GetExprOk() (*string, bool)`

GetExprOk returns a tuple with the Expr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpr

`func (o *AlertRecord) SetExpr(v string)`

SetExpr sets Expr field to given value.


### GetLabels

`func (o *AlertRecord) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AlertRecord) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AlertRecord) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AlertRecord) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


