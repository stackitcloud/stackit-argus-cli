# V1InstancesAlertconfigsUpdateRequestRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Receiver** | **string** | Receiver that should be one item of receivers &#x60;Additional Validators:&#x60; * must be a in name of receivers | 
**GroupBy** | Pointer to **[]string** | The labels by which incoming alerts are grouped together. For example, multiple alerts coming in for cluster&#x3D;A and alertname&#x3D;LatencyHigh would be batched into a single group. To aggregate by all possible labels use the special value &#39;...&#39; as the sole label name, for example: group_by: [&#39;...&#39;]. This effectively disables aggregation entirely, passing through all alerts as-is. This is unlikely to be what you want, unless you have a very low alert volume or your upstream notification system performs its own grouping. | [optional] 
**GroupWait** | Pointer to **string** | How long to initially wait to send a notification for a group of alerts. Allows to wait for an inhibiting alert to arrive or collect more initial alerts for the same group. (Usually ~0s to few minutes.) &#x60;Additional Validators:&#x60; * must be a valid time format | [optional] [default to "30s"]
**GroupInterval** | Pointer to **string** | How long to wait before sending a notification about new alerts that are added to a group of alerts for which an initial notification has already been sent. (Usually ~5m or more.) &#x60;Additional Validators:&#x60; * must be a valid time format | [optional] [default to "5m"]
**RepeatInterval** | Pointer to **string** | How long to wait before sending a notification again if it has already been sent successfully for an alert. (Usually ~3h or more). &#x60;Additional Validators:&#x60; * must be a valid time format | [optional] [default to "4h"]
**Match** | Pointer to **map[string]interface{}** | map of key:value. A set of equality matchers an alert has to fulfill to match the node.  &#x60;Additional Validators:&#x60; * should not contain more than 5 keys * each key and value should not be longer than 200 characters * key and values should only include the characters: a-zA-Z0-9_./@&amp;?:- | [optional] 
**MatchRe** | Pointer to **map[string]interface{}** | map of key:value. A set of regex-matchers an alert has to fulfill to match the node.  &#x60;Additional Validators:&#x60; * should not contain more than 5 keys * each key and value should not be longer than 200 characters | [optional] 
**Matchers** | Pointer to **[]string** | A list of matchers that an alert has to fulfill to match the node. A matcher is a string with a syntax inspired by PromQL and OpenMetrics. The syntax of a matcher consists of three tokens: * A valid Prometheus label name. * One of &#x3D;, !&#x3D;, &#x3D;~, or !~. &#x3D; means equals, !&#x3D; means that the strings are not equal, &#x3D;~ is used for equality of regex expressions and !~ is used for un-equality of regex expressions. They have the same meaning as known from PromQL selectors. * A UTF-8 string, which may be enclosed in double quotes. Before or after each token, there may be any amount of whitespace. &#x60;Additional Validators:&#x60; * should not contain more than 5 keys * each key and value should not be longer than 200 characters | [optional] 
**Routes** | Pointer to [**[]V1InstancesAlertconfigsUpdateRequestRouteRoutesInner**](V1InstancesAlertconfigsUpdateRequestRouteRoutesInner.md) | Zero or more child routes. | [optional] 

## Methods

### NewV1InstancesAlertconfigsUpdateRequestRoute

`func NewV1InstancesAlertconfigsUpdateRequestRoute(receiver string, ) *V1InstancesAlertconfigsUpdateRequestRoute`

NewV1InstancesAlertconfigsUpdateRequestRoute instantiates a new V1InstancesAlertconfigsUpdateRequestRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesAlertconfigsUpdateRequestRouteWithDefaults

`func NewV1InstancesAlertconfigsUpdateRequestRouteWithDefaults() *V1InstancesAlertconfigsUpdateRequestRoute`

NewV1InstancesAlertconfigsUpdateRequestRouteWithDefaults instantiates a new V1InstancesAlertconfigsUpdateRequestRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceiver

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.


### GetGroupBy

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetGroupWait

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) GetGroupWait() string`

GetGroupWait returns the GroupWait field if non-nil, zero value otherwise.

### GetGroupWaitOk

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) GetGroupWaitOk() (*string, bool)`

GetGroupWaitOk returns a tuple with the GroupWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupWait

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) SetGroupWait(v string)`

SetGroupWait sets GroupWait field to given value.

### HasGroupWait

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) HasGroupWait() bool`

HasGroupWait returns a boolean if a field has been set.

### GetGroupInterval

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) GetGroupInterval() string`

GetGroupInterval returns the GroupInterval field if non-nil, zero value otherwise.

### GetGroupIntervalOk

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) GetGroupIntervalOk() (*string, bool)`

GetGroupIntervalOk returns a tuple with the GroupInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupInterval

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) SetGroupInterval(v string)`

SetGroupInterval sets GroupInterval field to given value.

### HasGroupInterval

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) HasGroupInterval() bool`

HasGroupInterval returns a boolean if a field has been set.

### GetRepeatInterval

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) GetRepeatInterval() string`

GetRepeatInterval returns the RepeatInterval field if non-nil, zero value otherwise.

### GetRepeatIntervalOk

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) GetRepeatIntervalOk() (*string, bool)`

GetRepeatIntervalOk returns a tuple with the RepeatInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatInterval

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) SetRepeatInterval(v string)`

SetRepeatInterval sets RepeatInterval field to given value.

### HasRepeatInterval

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) HasRepeatInterval() bool`

HasRepeatInterval returns a boolean if a field has been set.

### GetMatch

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) GetMatch() map[string]interface{}`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) GetMatchOk() (*map[string]interface{}, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) SetMatch(v map[string]interface{})`

SetMatch sets Match field to given value.

### HasMatch

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetMatchRe

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) GetMatchRe() map[string]interface{}`

GetMatchRe returns the MatchRe field if non-nil, zero value otherwise.

### GetMatchReOk

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) GetMatchReOk() (*map[string]interface{}, bool)`

GetMatchReOk returns a tuple with the MatchRe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchRe

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) SetMatchRe(v map[string]interface{})`

SetMatchRe sets MatchRe field to given value.

### HasMatchRe

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) HasMatchRe() bool`

HasMatchRe returns a boolean if a field has been set.

### GetMatchers

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) GetMatchers() []string`

GetMatchers returns the Matchers field if non-nil, zero value otherwise.

### GetMatchersOk

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) GetMatchersOk() (*[]string, bool)`

GetMatchersOk returns a tuple with the Matchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchers

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) SetMatchers(v []string)`

SetMatchers sets Matchers field to given value.

### HasMatchers

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) HasMatchers() bool`

HasMatchers returns a boolean if a field has been set.

### GetRoutes

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) GetRoutes() []V1InstancesAlertconfigsUpdateRequestRouteRoutesInner`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) GetRoutesOk() (*[]V1InstancesAlertconfigsUpdateRequestRouteRoutesInner, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) SetRoutes(v []V1InstancesAlertconfigsUpdateRequestRouteRoutesInner)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *V1InstancesAlertconfigsUpdateRequestRoute) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


