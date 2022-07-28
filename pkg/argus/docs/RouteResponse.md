# RouteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Data** | [**Route**](Route.md) |  | 

## Methods

### NewRouteResponse

`func NewRouteResponse(message string, data Route, ) *RouteResponse`

NewRouteResponse instantiates a new RouteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteResponseWithDefaults

`func NewRouteResponseWithDefaults() *RouteResponse`

NewRouteResponseWithDefaults instantiates a new RouteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *RouteResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RouteResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RouteResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *RouteResponse) GetData() Route`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RouteResponse) GetDataOk() (*Route, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RouteResponse) SetData(v Route)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


