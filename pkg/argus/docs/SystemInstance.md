# SystemInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] [default to ""]
**MetricsUrl** | **string** |  | 
**TargetsUrl** | **string** |  | 
**AlertingUrl** | **string** |  | 
**PushMetricsUrl** | **string** |  | 
**LogsUrl** | **string** |  | 
**LogsPushUrl** | **string** |  | 
**JaegerTracesUrl** | **string** |  | 
**OtlpTracesUrl** | **string** |  | 
**ZipkinSpansUrl** | **string** |  | 
**JaegerUiUrl** | **string** |  | 

## Methods

### NewSystemInstance

`func NewSystemInstance(instance string, metricsUrl string, targetsUrl string, alertingUrl string, pushMetricsUrl string, logsUrl string, logsPushUrl string, jaegerTracesUrl string, otlpTracesUrl string, zipkinSpansUrl string, jaegerUiUrl string, ) *SystemInstance`

NewSystemInstance instantiates a new SystemInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemInstanceWithDefaults

`func NewSystemInstanceWithDefaults() *SystemInstance`

NewSystemInstanceWithDefaults instantiates a new SystemInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *SystemInstance) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *SystemInstance) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *SystemInstance) SetInstance(v string)`

SetInstance sets Instance field to given value.


### GetName

`func (o *SystemInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SystemInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetricsUrl

`func (o *SystemInstance) GetMetricsUrl() string`

GetMetricsUrl returns the MetricsUrl field if non-nil, zero value otherwise.

### GetMetricsUrlOk

`func (o *SystemInstance) GetMetricsUrlOk() (*string, bool)`

GetMetricsUrlOk returns a tuple with the MetricsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsUrl

`func (o *SystemInstance) SetMetricsUrl(v string)`

SetMetricsUrl sets MetricsUrl field to given value.


### GetTargetsUrl

`func (o *SystemInstance) GetTargetsUrl() string`

GetTargetsUrl returns the TargetsUrl field if non-nil, zero value otherwise.

### GetTargetsUrlOk

`func (o *SystemInstance) GetTargetsUrlOk() (*string, bool)`

GetTargetsUrlOk returns a tuple with the TargetsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetsUrl

`func (o *SystemInstance) SetTargetsUrl(v string)`

SetTargetsUrl sets TargetsUrl field to given value.


### GetAlertingUrl

`func (o *SystemInstance) GetAlertingUrl() string`

GetAlertingUrl returns the AlertingUrl field if non-nil, zero value otherwise.

### GetAlertingUrlOk

`func (o *SystemInstance) GetAlertingUrlOk() (*string, bool)`

GetAlertingUrlOk returns a tuple with the AlertingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingUrl

`func (o *SystemInstance) SetAlertingUrl(v string)`

SetAlertingUrl sets AlertingUrl field to given value.


### GetPushMetricsUrl

`func (o *SystemInstance) GetPushMetricsUrl() string`

GetPushMetricsUrl returns the PushMetricsUrl field if non-nil, zero value otherwise.

### GetPushMetricsUrlOk

`func (o *SystemInstance) GetPushMetricsUrlOk() (*string, bool)`

GetPushMetricsUrlOk returns a tuple with the PushMetricsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushMetricsUrl

`func (o *SystemInstance) SetPushMetricsUrl(v string)`

SetPushMetricsUrl sets PushMetricsUrl field to given value.


### GetLogsUrl

`func (o *SystemInstance) GetLogsUrl() string`

GetLogsUrl returns the LogsUrl field if non-nil, zero value otherwise.

### GetLogsUrlOk

`func (o *SystemInstance) GetLogsUrlOk() (*string, bool)`

GetLogsUrlOk returns a tuple with the LogsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsUrl

`func (o *SystemInstance) SetLogsUrl(v string)`

SetLogsUrl sets LogsUrl field to given value.


### GetLogsPushUrl

`func (o *SystemInstance) GetLogsPushUrl() string`

GetLogsPushUrl returns the LogsPushUrl field if non-nil, zero value otherwise.

### GetLogsPushUrlOk

`func (o *SystemInstance) GetLogsPushUrlOk() (*string, bool)`

GetLogsPushUrlOk returns a tuple with the LogsPushUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsPushUrl

`func (o *SystemInstance) SetLogsPushUrl(v string)`

SetLogsPushUrl sets LogsPushUrl field to given value.


### GetJaegerTracesUrl

`func (o *SystemInstance) GetJaegerTracesUrl() string`

GetJaegerTracesUrl returns the JaegerTracesUrl field if non-nil, zero value otherwise.

### GetJaegerTracesUrlOk

`func (o *SystemInstance) GetJaegerTracesUrlOk() (*string, bool)`

GetJaegerTracesUrlOk returns a tuple with the JaegerTracesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJaegerTracesUrl

`func (o *SystemInstance) SetJaegerTracesUrl(v string)`

SetJaegerTracesUrl sets JaegerTracesUrl field to given value.


### GetOtlpTracesUrl

`func (o *SystemInstance) GetOtlpTracesUrl() string`

GetOtlpTracesUrl returns the OtlpTracesUrl field if non-nil, zero value otherwise.

### GetOtlpTracesUrlOk

`func (o *SystemInstance) GetOtlpTracesUrlOk() (*string, bool)`

GetOtlpTracesUrlOk returns a tuple with the OtlpTracesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtlpTracesUrl

`func (o *SystemInstance) SetOtlpTracesUrl(v string)`

SetOtlpTracesUrl sets OtlpTracesUrl field to given value.


### GetZipkinSpansUrl

`func (o *SystemInstance) GetZipkinSpansUrl() string`

GetZipkinSpansUrl returns the ZipkinSpansUrl field if non-nil, zero value otherwise.

### GetZipkinSpansUrlOk

`func (o *SystemInstance) GetZipkinSpansUrlOk() (*string, bool)`

GetZipkinSpansUrlOk returns a tuple with the ZipkinSpansUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipkinSpansUrl

`func (o *SystemInstance) SetZipkinSpansUrl(v string)`

SetZipkinSpansUrl sets ZipkinSpansUrl field to given value.


### GetJaegerUiUrl

`func (o *SystemInstance) GetJaegerUiUrl() string`

GetJaegerUiUrl returns the JaegerUiUrl field if non-nil, zero value otherwise.

### GetJaegerUiUrlOk

`func (o *SystemInstance) GetJaegerUiUrlOk() (*string, bool)`

GetJaegerUiUrlOk returns a tuple with the JaegerUiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJaegerUiUrl

`func (o *SystemInstance) SetJaegerUiUrl(v string)`

SetJaegerUiUrl sets JaegerUiUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


