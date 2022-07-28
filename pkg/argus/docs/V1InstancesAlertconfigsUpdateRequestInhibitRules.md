# V1InstancesAlertconfigsUpdateRequestInhibitRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceMatch** | Pointer to **map[string]interface{}** | map of key:value. Matchers for which one or more alerts have to exist for the inhibition to take effect. | [optional] 
**SourceMatchRe** | Pointer to **map[string]interface{}** | map of key:value. Regex match | [optional] 
**TargetMatch** | Pointer to **map[string]interface{}** | map of key:value. Matchers that have to be fulfilled in the alerts to be muted. | [optional] 
**TargetMatchRe** | Pointer to **map[string]interface{}** | map of key:value. Matchers that have to be fulfilled in the alerts to be muted. Regex. | [optional] 
**Equal** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV1InstancesAlertconfigsUpdateRequestInhibitRules

`func NewV1InstancesAlertconfigsUpdateRequestInhibitRules() *V1InstancesAlertconfigsUpdateRequestInhibitRules`

NewV1InstancesAlertconfigsUpdateRequestInhibitRules instantiates a new V1InstancesAlertconfigsUpdateRequestInhibitRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesAlertconfigsUpdateRequestInhibitRulesWithDefaults

`func NewV1InstancesAlertconfigsUpdateRequestInhibitRulesWithDefaults() *V1InstancesAlertconfigsUpdateRequestInhibitRules`

NewV1InstancesAlertconfigsUpdateRequestInhibitRulesWithDefaults instantiates a new V1InstancesAlertconfigsUpdateRequestInhibitRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceMatch

`func (o *V1InstancesAlertconfigsUpdateRequestInhibitRules) GetSourceMatch() map[string]interface{}`

GetSourceMatch returns the SourceMatch field if non-nil, zero value otherwise.

### GetSourceMatchOk

`func (o *V1InstancesAlertconfigsUpdateRequestInhibitRules) GetSourceMatchOk() (*map[string]interface{}, bool)`

GetSourceMatchOk returns a tuple with the SourceMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMatch

`func (o *V1InstancesAlertconfigsUpdateRequestInhibitRules) SetSourceMatch(v map[string]interface{})`

SetSourceMatch sets SourceMatch field to given value.

### HasSourceMatch

`func (o *V1InstancesAlertconfigsUpdateRequestInhibitRules) HasSourceMatch() bool`

HasSourceMatch returns a boolean if a field has been set.

### GetSourceMatchRe

`func (o *V1InstancesAlertconfigsUpdateRequestInhibitRules) GetSourceMatchRe() map[string]interface{}`

GetSourceMatchRe returns the SourceMatchRe field if non-nil, zero value otherwise.

### GetSourceMatchReOk

`func (o *V1InstancesAlertconfigsUpdateRequestInhibitRules) GetSourceMatchReOk() (*map[string]interface{}, bool)`

GetSourceMatchReOk returns a tuple with the SourceMatchRe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMatchRe

`func (o *V1InstancesAlertconfigsUpdateRequestInhibitRules) SetSourceMatchRe(v map[string]interface{})`

SetSourceMatchRe sets SourceMatchRe field to given value.

### HasSourceMatchRe

`func (o *V1InstancesAlertconfigsUpdateRequestInhibitRules) HasSourceMatchRe() bool`

HasSourceMatchRe returns a boolean if a field has been set.

### GetTargetMatch

`func (o *V1InstancesAlertconfigsUpdateRequestInhibitRules) GetTargetMatch() map[string]interface{}`

GetTargetMatch returns the TargetMatch field if non-nil, zero value otherwise.

### GetTargetMatchOk

`func (o *V1InstancesAlertconfigsUpdateRequestInhibitRules) GetTargetMatchOk() (*map[string]interface{}, bool)`

GetTargetMatchOk returns a tuple with the TargetMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMatch

`func (o *V1InstancesAlertconfigsUpdateRequestInhibitRules) SetTargetMatch(v map[string]interface{})`

SetTargetMatch sets TargetMatch field to given value.

### HasTargetMatch

`func (o *V1InstancesAlertconfigsUpdateRequestInhibitRules) HasTargetMatch() bool`

HasTargetMatch returns a boolean if a field has been set.

### GetTargetMatchRe

`func (o *V1InstancesAlertconfigsUpdateRequestInhibitRules) GetTargetMatchRe() map[string]interface{}`

GetTargetMatchRe returns the TargetMatchRe field if non-nil, zero value otherwise.

### GetTargetMatchReOk

`func (o *V1InstancesAlertconfigsUpdateRequestInhibitRules) GetTargetMatchReOk() (*map[string]interface{}, bool)`

GetTargetMatchReOk returns a tuple with the TargetMatchRe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMatchRe

`func (o *V1InstancesAlertconfigsUpdateRequestInhibitRules) SetTargetMatchRe(v map[string]interface{})`

SetTargetMatchRe sets TargetMatchRe field to given value.

### HasTargetMatchRe

`func (o *V1InstancesAlertconfigsUpdateRequestInhibitRules) HasTargetMatchRe() bool`

HasTargetMatchRe returns a boolean if a field has been set.

### GetEqual

`func (o *V1InstancesAlertconfigsUpdateRequestInhibitRules) GetEqual() []string`

GetEqual returns the Equal field if non-nil, zero value otherwise.

### GetEqualOk

`func (o *V1InstancesAlertconfigsUpdateRequestInhibitRules) GetEqualOk() (*[]string, bool)`

GetEqualOk returns a tuple with the Equal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEqual

`func (o *V1InstancesAlertconfigsUpdateRequestInhibitRules) SetEqual(v []string)`

SetEqual sets Equal field to given value.

### HasEqual

`func (o *V1InstancesAlertconfigsUpdateRequestInhibitRules) HasEqual() bool`

HasEqual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


