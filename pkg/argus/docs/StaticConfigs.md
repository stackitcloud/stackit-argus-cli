# StaticConfigs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Targets** | **[]string** |  | 
**Labels** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewStaticConfigs

`func NewStaticConfigs(targets []string, ) *StaticConfigs`

NewStaticConfigs instantiates a new StaticConfigs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticConfigsWithDefaults

`func NewStaticConfigsWithDefaults() *StaticConfigs`

NewStaticConfigsWithDefaults instantiates a new StaticConfigs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargets

`func (o *StaticConfigs) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *StaticConfigs) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *StaticConfigs) SetTargets(v []string)`

SetTargets sets Targets field to given value.


### GetLabels

`func (o *StaticConfigs) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *StaticConfigs) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *StaticConfigs) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *StaticConfigs) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


