# V1InstancesScrapeconfigsUpdateRequestStaticConfigsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Targets** | **[]string** | The targets specified by the static config. | 
**Labels** | Pointer to **map[string]interface{}** | Labels assigned to all metrics scraped from the targets. &#x60;Additional Validators:&#x60; * should not contain more than 5 keys * each key and value should not be longer than 200 characters | [optional] 

## Methods

### NewV1InstancesScrapeconfigsUpdateRequestStaticConfigsInner

`func NewV1InstancesScrapeconfigsUpdateRequestStaticConfigsInner(targets []string, ) *V1InstancesScrapeconfigsUpdateRequestStaticConfigsInner`

NewV1InstancesScrapeconfigsUpdateRequestStaticConfigsInner instantiates a new V1InstancesScrapeconfigsUpdateRequestStaticConfigsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstancesScrapeconfigsUpdateRequestStaticConfigsInnerWithDefaults

`func NewV1InstancesScrapeconfigsUpdateRequestStaticConfigsInnerWithDefaults() *V1InstancesScrapeconfigsUpdateRequestStaticConfigsInner`

NewV1InstancesScrapeconfigsUpdateRequestStaticConfigsInnerWithDefaults instantiates a new V1InstancesScrapeconfigsUpdateRequestStaticConfigsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargets

`func (o *V1InstancesScrapeconfigsUpdateRequestStaticConfigsInner) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *V1InstancesScrapeconfigsUpdateRequestStaticConfigsInner) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *V1InstancesScrapeconfigsUpdateRequestStaticConfigsInner) SetTargets(v []string)`

SetTargets sets Targets field to given value.


### GetLabels

`func (o *V1InstancesScrapeconfigsUpdateRequestStaticConfigsInner) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *V1InstancesScrapeconfigsUpdateRequestStaticConfigsInner) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *V1InstancesScrapeconfigsUpdateRequestStaticConfigsInner) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *V1InstancesScrapeconfigsUpdateRequestStaticConfigsInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


