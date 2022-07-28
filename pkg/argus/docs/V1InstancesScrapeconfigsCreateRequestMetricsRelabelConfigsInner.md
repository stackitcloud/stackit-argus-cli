# V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceLabels** | Pointer to **[]string** | The source labels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the replace, keep, and drop actions. | [optional] 
**Separator** | Pointer to **string** | Separator placed between concatenated source label values. | [optional] [default to ";"]
**TargetLabel** | Pointer to **string** | Label to which the resulting value is written in a replace action. It is mandatory for replace actions. Regex capture groups are available. | [optional] 
**Regex** | Pointer to **string** | Regular expression against which the extracted value is matched. | [optional] [default to ".*"]
**Modulus** | Pointer to **float32** | Modulus to take of the hash of the source label values. | [optional] 
**Replacement** | Pointer to **string** | Replacement value against which a regex replace is performed if the regular expression matches. Regex capture groups are available. | [optional] [default to "$1"]
**Action** | Pointer to **string** | Choice field [&#39;replace&#39;, &#39;keep&#39;, &#39;drop&#39;, &#39;hashmod&#39;, &#39;labelmap&#39;, &#39;labeldrop&#39;, &#39;labelkeep&#39;] | [optional] [default to "replace"]

## Methods

### NewV1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner

`func NewV1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner() *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner`

NewV1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner instantiates a new V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInnerWithDefaults

`func NewV1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInnerWithDefaults() *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner`

NewV1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInnerWithDefaults instantiates a new V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceLabels

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) GetSourceLabels() []string`

GetSourceLabels returns the SourceLabels field if non-nil, zero value otherwise.

### GetSourceLabelsOk

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) GetSourceLabelsOk() (*[]string, bool)`

GetSourceLabelsOk returns a tuple with the SourceLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLabels

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) SetSourceLabels(v []string)`

SetSourceLabels sets SourceLabels field to given value.

### HasSourceLabels

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) HasSourceLabels() bool`

HasSourceLabels returns a boolean if a field has been set.

### GetSeparator

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) GetSeparator() string`

GetSeparator returns the Separator field if non-nil, zero value otherwise.

### GetSeparatorOk

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) GetSeparatorOk() (*string, bool)`

GetSeparatorOk returns a tuple with the Separator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparator

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) SetSeparator(v string)`

SetSeparator sets Separator field to given value.

### HasSeparator

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) HasSeparator() bool`

HasSeparator returns a boolean if a field has been set.

### GetTargetLabel

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) GetTargetLabel() string`

GetTargetLabel returns the TargetLabel field if non-nil, zero value otherwise.

### GetTargetLabelOk

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) GetTargetLabelOk() (*string, bool)`

GetTargetLabelOk returns a tuple with the TargetLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLabel

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) SetTargetLabel(v string)`

SetTargetLabel sets TargetLabel field to given value.

### HasTargetLabel

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) HasTargetLabel() bool`

HasTargetLabel returns a boolean if a field has been set.

### GetRegex

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### GetModulus

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) GetModulus() float32`

GetModulus returns the Modulus field if non-nil, zero value otherwise.

### GetModulusOk

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) GetModulusOk() (*float32, bool)`

GetModulusOk returns a tuple with the Modulus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModulus

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) SetModulus(v float32)`

SetModulus sets Modulus field to given value.

### HasModulus

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) HasModulus() bool`

HasModulus returns a boolean if a field has been set.

### GetReplacement

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) GetReplacement() string`

GetReplacement returns the Replacement field if non-nil, zero value otherwise.

### GetReplacementOk

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) GetReplacementOk() (*string, bool)`

GetReplacementOk returns a tuple with the Replacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacement

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) SetReplacement(v string)`

SetReplacement sets Replacement field to given value.

### HasReplacement

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) HasReplacement() bool`

HasReplacement returns a boolean if a field has been set.

### GetAction

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *V1InstancesScrapeconfigsCreateRequestMetricsRelabelConfigsInner) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


