# V1ClustersCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kubeconfig** | **string** | Encoded base64 string of a valid kubeconfig | 
**ClusterName** | Pointer to **string** | Concrete name of the cluster. If you leave the field blank, the cluster name is clusterId{len(clusterId)}. So in a cluster category defined by clusterId many concrete clusters can exist. | [optional] 

## Methods

### NewV1ClustersCreateRequest

`func NewV1ClustersCreateRequest(kubeconfig string, ) *V1ClustersCreateRequest`

NewV1ClustersCreateRequest instantiates a new V1ClustersCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ClustersCreateRequestWithDefaults

`func NewV1ClustersCreateRequestWithDefaults() *V1ClustersCreateRequest`

NewV1ClustersCreateRequestWithDefaults instantiates a new V1ClustersCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKubeconfig

`func (o *V1ClustersCreateRequest) GetKubeconfig() string`

GetKubeconfig returns the Kubeconfig field if non-nil, zero value otherwise.

### GetKubeconfigOk

`func (o *V1ClustersCreateRequest) GetKubeconfigOk() (*string, bool)`

GetKubeconfigOk returns a tuple with the Kubeconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeconfig

`func (o *V1ClustersCreateRequest) SetKubeconfig(v string)`

SetKubeconfig sets Kubeconfig field to given value.


### GetClusterName

`func (o *V1ClustersCreateRequest) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *V1ClustersCreateRequest) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *V1ClustersCreateRequest) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *V1ClustersCreateRequest) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


