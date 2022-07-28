# RouteSerializer2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Receiver** | Pointer to **string** |  | [optional] 
**GroupBy** | Pointer to **[]string** |  | [optional] 
**GroupWait** | Pointer to **string** |  | [optional] [default to "30s"]
**GroupInterval** | Pointer to **string** |  | [optional] [default to "5m"]
**RepeatInterval** | Pointer to **string** |  | [optional] [default to "4h"]
**Match** | Pointer to **map[string]string** |  | [optional] 
**MatchRe** | Pointer to **map[string]string** |  | [optional] 
**Continue** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewRouteSerializer2

`func NewRouteSerializer2() *RouteSerializer2`

NewRouteSerializer2 instantiates a new RouteSerializer2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteSerializer2WithDefaults

`func NewRouteSerializer2WithDefaults() *RouteSerializer2`

NewRouteSerializer2WithDefaults instantiates a new RouteSerializer2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceiver

`func (o *RouteSerializer2) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *RouteSerializer2) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *RouteSerializer2) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *RouteSerializer2) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetGroupBy

`func (o *RouteSerializer2) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *RouteSerializer2) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *RouteSerializer2) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *RouteSerializer2) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetGroupWait

`func (o *RouteSerializer2) GetGroupWait() string`

GetGroupWait returns the GroupWait field if non-nil, zero value otherwise.

### GetGroupWaitOk

`func (o *RouteSerializer2) GetGroupWaitOk() (*string, bool)`

GetGroupWaitOk returns a tuple with the GroupWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupWait

`func (o *RouteSerializer2) SetGroupWait(v string)`

SetGroupWait sets GroupWait field to given value.

### HasGroupWait

`func (o *RouteSerializer2) HasGroupWait() bool`

HasGroupWait returns a boolean if a field has been set.

### GetGroupInterval

`func (o *RouteSerializer2) GetGroupInterval() string`

GetGroupInterval returns the GroupInterval field if non-nil, zero value otherwise.

### GetGroupIntervalOk

`func (o *RouteSerializer2) GetGroupIntervalOk() (*string, bool)`

GetGroupIntervalOk returns a tuple with the GroupInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupInterval

`func (o *RouteSerializer2) SetGroupInterval(v string)`

SetGroupInterval sets GroupInterval field to given value.

### HasGroupInterval

`func (o *RouteSerializer2) HasGroupInterval() bool`

HasGroupInterval returns a boolean if a field has been set.

### GetRepeatInterval

`func (o *RouteSerializer2) GetRepeatInterval() string`

GetRepeatInterval returns the RepeatInterval field if non-nil, zero value otherwise.

### GetRepeatIntervalOk

`func (o *RouteSerializer2) GetRepeatIntervalOk() (*string, bool)`

GetRepeatIntervalOk returns a tuple with the RepeatInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatInterval

`func (o *RouteSerializer2) SetRepeatInterval(v string)`

SetRepeatInterval sets RepeatInterval field to given value.

### HasRepeatInterval

`func (o *RouteSerializer2) HasRepeatInterval() bool`

HasRepeatInterval returns a boolean if a field has been set.

### GetMatch

`func (o *RouteSerializer2) GetMatch() map[string]string`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *RouteSerializer2) GetMatchOk() (*map[string]string, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *RouteSerializer2) SetMatch(v map[string]string)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *RouteSerializer2) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetMatchRe

`func (o *RouteSerializer2) GetMatchRe() map[string]string`

GetMatchRe returns the MatchRe field if non-nil, zero value otherwise.

### GetMatchReOk

`func (o *RouteSerializer2) GetMatchReOk() (*map[string]string, bool)`

GetMatchReOk returns a tuple with the MatchRe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchRe

`func (o *RouteSerializer2) SetMatchRe(v map[string]string)`

SetMatchRe sets MatchRe field to given value.

### HasMatchRe

`func (o *RouteSerializer2) HasMatchRe() bool`

HasMatchRe returns a boolean if a field has been set.

### GetContinue

`func (o *RouteSerializer2) GetContinue() bool`

GetContinue returns the Continue field if non-nil, zero value otherwise.

### GetContinueOk

`func (o *RouteSerializer2) GetContinueOk() (*bool, bool)`

GetContinueOk returns a tuple with the Continue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinue

`func (o *RouteSerializer2) SetContinue(v bool)`

SetContinue sets Continue field to given value.

### HasContinue

`func (o *RouteSerializer2) HasContinue() bool`

HasContinue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


