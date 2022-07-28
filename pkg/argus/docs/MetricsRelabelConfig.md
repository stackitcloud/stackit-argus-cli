# MetricsRelabelConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceLabels** | **[]string** |  | 
**Separator** | Pointer to **string** |  | [optional] [default to ";"]
**TargetLabel** | Pointer to **string** |  | [optional] 
**Regex** | Pointer to **string** |  | [optional] [default to ".*"]
**Modulus** | Pointer to **int32** |  | [optional] 
**Replacement** | Pointer to **string** |  | [optional] [default to "$1"]
**Action** | Pointer to **string** |  | [optional] [default to "replace"]

## Methods

### NewMetricsRelabelConfig

`func NewMetricsRelabelConfig(sourceLabels []string, ) *MetricsRelabelConfig`

NewMetricsRelabelConfig instantiates a new MetricsRelabelConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsRelabelConfigWithDefaults

`func NewMetricsRelabelConfigWithDefaults() *MetricsRelabelConfig`

NewMetricsRelabelConfigWithDefaults instantiates a new MetricsRelabelConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceLabels

`func (o *MetricsRelabelConfig) GetSourceLabels() []string`

GetSourceLabels returns the SourceLabels field if non-nil, zero value otherwise.

### GetSourceLabelsOk

`func (o *MetricsRelabelConfig) GetSourceLabelsOk() (*[]string, bool)`

GetSourceLabelsOk returns a tuple with the SourceLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLabels

`func (o *MetricsRelabelConfig) SetSourceLabels(v []string)`

SetSourceLabels sets SourceLabels field to given value.


### GetSeparator

`func (o *MetricsRelabelConfig) GetSeparator() string`

GetSeparator returns the Separator field if non-nil, zero value otherwise.

### GetSeparatorOk

`func (o *MetricsRelabelConfig) GetSeparatorOk() (*string, bool)`

GetSeparatorOk returns a tuple with the Separator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparator

`func (o *MetricsRelabelConfig) SetSeparator(v string)`

SetSeparator sets Separator field to given value.

### HasSeparator

`func (o *MetricsRelabelConfig) HasSeparator() bool`

HasSeparator returns a boolean if a field has been set.

### GetTargetLabel

`func (o *MetricsRelabelConfig) GetTargetLabel() string`

GetTargetLabel returns the TargetLabel field if non-nil, zero value otherwise.

### GetTargetLabelOk

`func (o *MetricsRelabelConfig) GetTargetLabelOk() (*string, bool)`

GetTargetLabelOk returns a tuple with the TargetLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLabel

`func (o *MetricsRelabelConfig) SetTargetLabel(v string)`

SetTargetLabel sets TargetLabel field to given value.

### HasTargetLabel

`func (o *MetricsRelabelConfig) HasTargetLabel() bool`

HasTargetLabel returns a boolean if a field has been set.

### GetRegex

`func (o *MetricsRelabelConfig) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *MetricsRelabelConfig) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *MetricsRelabelConfig) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *MetricsRelabelConfig) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### GetModulus

`func (o *MetricsRelabelConfig) GetModulus() int32`

GetModulus returns the Modulus field if non-nil, zero value otherwise.

### GetModulusOk

`func (o *MetricsRelabelConfig) GetModulusOk() (*int32, bool)`

GetModulusOk returns a tuple with the Modulus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModulus

`func (o *MetricsRelabelConfig) SetModulus(v int32)`

SetModulus sets Modulus field to given value.

### HasModulus

`func (o *MetricsRelabelConfig) HasModulus() bool`

HasModulus returns a boolean if a field has been set.

### GetReplacement

`func (o *MetricsRelabelConfig) GetReplacement() string`

GetReplacement returns the Replacement field if non-nil, zero value otherwise.

### GetReplacementOk

`func (o *MetricsRelabelConfig) GetReplacementOk() (*string, bool)`

GetReplacementOk returns a tuple with the Replacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacement

`func (o *MetricsRelabelConfig) SetReplacement(v string)`

SetReplacement sets Replacement field to given value.

### HasReplacement

`func (o *MetricsRelabelConfig) HasReplacement() bool`

HasReplacement returns a boolean if a field has been set.

### GetAction

`func (o *MetricsRelabelConfig) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *MetricsRelabelConfig) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *MetricsRelabelConfig) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *MetricsRelabelConfig) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


