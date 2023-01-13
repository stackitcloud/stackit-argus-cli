# Route

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Receiver** | **string** |  | 
**GroupBy** | Pointer to **[]string** |  | [optional] 
**GroupWait** | Pointer to **string** |  | [optional] [default to "30s"]
**GroupInterval** | Pointer to **string** |  | [optional] [default to "5m"]
**RepeatInterval** | Pointer to **string** |  | [optional] [default to "4h"]
**Match** | Pointer to **map[string]string** |  | [optional] 
**MatchRe** | Pointer to **map[string]string** |  | [optional] 
**Routes** | Pointer to [**[]RouteSerializer2**](RouteSerializer2.md) |  | [optional] 
**Continue** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewRoute

`func NewRoute(receiver string, ) *Route`

NewRoute instantiates a new Route object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteWithDefaults

`func NewRouteWithDefaults() *Route`

NewRouteWithDefaults instantiates a new Route object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceiver

`func (o *Route) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *Route) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *Route) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.


### GetGroupBy

`func (o *Route) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *Route) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *Route) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *Route) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetGroupWait

`func (o *Route) GetGroupWait() string`

GetGroupWait returns the GroupWait field if non-nil, zero value otherwise.

### GetGroupWaitOk

`func (o *Route) GetGroupWaitOk() (*string, bool)`

GetGroupWaitOk returns a tuple with the GroupWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupWait

`func (o *Route) SetGroupWait(v string)`

SetGroupWait sets GroupWait field to given value.

### HasGroupWait

`func (o *Route) HasGroupWait() bool`

HasGroupWait returns a boolean if a field has been set.

### GetGroupInterval

`func (o *Route) GetGroupInterval() string`

GetGroupInterval returns the GroupInterval field if non-nil, zero value otherwise.

### GetGroupIntervalOk

`func (o *Route) GetGroupIntervalOk() (*string, bool)`

GetGroupIntervalOk returns a tuple with the GroupInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupInterval

`func (o *Route) SetGroupInterval(v string)`

SetGroupInterval sets GroupInterval field to given value.

### HasGroupInterval

`func (o *Route) HasGroupInterval() bool`

HasGroupInterval returns a boolean if a field has been set.

### GetRepeatInterval

`func (o *Route) GetRepeatInterval() string`

GetRepeatInterval returns the RepeatInterval field if non-nil, zero value otherwise.

### GetRepeatIntervalOk

`func (o *Route) GetRepeatIntervalOk() (*string, bool)`

GetRepeatIntervalOk returns a tuple with the RepeatInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatInterval

`func (o *Route) SetRepeatInterval(v string)`

SetRepeatInterval sets RepeatInterval field to given value.

### HasRepeatInterval

`func (o *Route) HasRepeatInterval() bool`

HasRepeatInterval returns a boolean if a field has been set.

### GetMatch

`func (o *Route) GetMatch() map[string]string`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *Route) GetMatchOk() (*map[string]string, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *Route) SetMatch(v map[string]string)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *Route) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetMatchRe

`func (o *Route) GetMatchRe() map[string]string`

GetMatchRe returns the MatchRe field if non-nil, zero value otherwise.

### GetMatchReOk

`func (o *Route) GetMatchReOk() (*map[string]string, bool)`

GetMatchReOk returns a tuple with the MatchRe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchRe

`func (o *Route) SetMatchRe(v map[string]string)`

SetMatchRe sets MatchRe field to given value.

### HasMatchRe

`func (o *Route) HasMatchRe() bool`

HasMatchRe returns a boolean if a field has been set.

### GetRoutes

`func (o *Route) GetRoutes() []RouteSerializer2`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *Route) GetRoutesOk() (*[]RouteSerializer2, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *Route) SetRoutes(v []RouteSerializer2)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *Route) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetContinue

`func (o *Route) GetContinue() bool`

GetContinue returns the Continue field if non-nil, zero value otherwise.

### GetContinueOk

`func (o *Route) GetContinueOk() (*bool, bool)`

GetContinueOk returns a tuple with the Continue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinue

`func (o *Route) SetContinue(v bool)`

SetContinue sets Continue field to given value.

### HasContinue

`func (o *Route) HasContinue() bool`

HasContinue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


